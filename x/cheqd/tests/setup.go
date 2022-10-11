package tests

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/cheqd/cheqd-node/x/cheqd/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/multiformats/go-multibase"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/cheqd/cheqd-node/x/cheqd/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type TestSetup struct {
	Cdc codec.Codec

	SdkCtx sdk.Context
	StdCtx context.Context

	Keeper      keeper.Keeper
	MsgServer   types.MsgServer
	QueryServer types.QueryServer
}

func Setup() TestSetup {
	// Init Codec
	ir := codectypes.NewInterfaceRegistry()
	types.RegisterInterfaces(ir)
	cdc := codec.NewProtoCodec(ir)

	// Init KVSore
	db := dbm.NewMemDB()

	dbStore := store.NewCommitMultiStore(db)
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	dbStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, nil)

	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	newKeeper := keeper.NewKeeper(cdc, storeKey)

	// Create Tx
	txBytes := make([]byte, 28)
	_, _ = rand.Read(txBytes)

	// Create context
	blockTime, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00.000Z")
	ctx := sdk.NewContext(dbStore,
		tmproto.Header{ChainID: "test", Time: blockTime},
		false, log.NewNopLogger()).WithTxBytes(txBytes)

	msgServer := keeper.NewMsgServer(*newKeeper)
	queryServer := keeper.NewQueryServer(*newKeeper)

	setup := TestSetup{
		Cdc: cdc,

		SdkCtx: ctx,
		StdCtx: sdk.WrapSDKContext(ctx),

		Keeper:      *newKeeper,
		MsgServer:   msgServer,
		QueryServer: queryServer,
	}

	setup.Keeper.SetDidNamespace(&ctx, DID_NAMESPACE)
	return setup
}

func (s *TestSetup) BuildMsgCreateDidPayload(did string, keyId string, key ed25519.PublicKey) *types.MsgCreateDidPayload {
	return &types.MsgCreateDidPayload{
		Id: did,
		VerificationMethod: []*types.VerificationMethod{
			{
				Id:                 keyId,
				Type:               types.Ed25519VerificationKey2020,
				Controller:         did,
				PublicKeyMultibase: MustEncodeBase58(key),
			},
		},
		Authentication: []string{keyId},
	}
}

func (s *TestSetup) CreateToUpdateDid(did *types.MsgCreateDidPayload) *types.MsgUpdateDidPayload {
	return &types.MsgUpdateDidPayload{
		Id:                   did.Id,
		Controller:           did.Controller,
		VerificationMethod:   did.VerificationMethod,
		Authentication:       did.Authentication,
		AssertionMethod:      did.AssertionMethod,
		CapabilityInvocation: did.CapabilityInvocation,
		CapabilityDelegation: did.CapabilityDelegation,
		KeyAgreement:         did.KeyAgreement,
		AlsoKnownAs:          did.AlsoKnownAs,
		Service:              did.Service,
		Context:              did.Context,
	}
}

// TODO: Remove
func (s *TestSetup) WrapCreateRequest(payload *types.MsgCreateDidPayload, keys map[string]ed25519.PrivateKey) *types.MsgCreateDid {
	var signatures []*types.SignInfo
	signingInput := payload.GetSignBytes()

	for privKeyId, privKey := range keys {
		signature := base64.StdEncoding.EncodeToString(ed25519.Sign(privKey, signingInput))
		signatures = append(signatures, &types.SignInfo{
			VerificationMethodId: privKeyId,
			Signature:            signature,
		})
	}

	return &types.MsgCreateDid{
		Payload:    payload,
		Signatures: signatures,
	}
}

func (s *TestSetup) CreateDid(payload *types.MsgCreateDidPayload, signInputs []SignInput) (*types.MsgCreateDidResponse, error) {
	signBytes := payload.GetSignBytes()
	var signatures []*types.SignInfo

	for _, input := range signInputs {
		signature := ed25519.Sign(input.Key, signBytes)

		signatures = append(signatures, &types.SignInfo{
			VerificationMethodId: input.VerificationMethodId,
			Signature:            base64.StdEncoding.EncodeToString(signature),
		})
	}

	msg := &types.MsgCreateDid{
		Payload:    payload,
		Signatures: signatures,
	}

	return s.MsgServer.CreateDid(s.StdCtx, msg)
}

func (s *TestSetup) CreateTestDid() (string, KeyPair, string) {
	did := GenerateDID(Base58_16chars)
	keypair := GenerateKeyPair()
	keyId := did + "#key-1"

	msg := types.MsgCreateDidPayload{
		Id: did,
		VerificationMethod: []*types.VerificationMethod{
			{
				Id:                 keyId,
				Type:               types.Ed25519VerificationKey2020,
				Controller:         did,
				PublicKeyMultibase: MustEncodeBase58(keypair.Public),
			},
		},
		Authentication: []string{keyId},
	}

	signatures := []SignInput{
		{
			VerificationMethodId: keyId,
			Key:                  keypair.Private,
		},
	}

	_, err := s.CreateDid(&msg, signatures)
	if err != nil {
		panic(err)
	}

	return did, keypair, keyId
}

func (s *TestSetup) QueryDid(did string) (*types.QueryGetDidResponse, error) {
	req := &types.QueryGetDidRequest{
		Id: did,
	}

	return s.QueryServer.Did(s.StdCtx, req)
}

func (s *TestSetup) WrapUpdateRequest(payload *types.MsgUpdateDidPayload, keys []SignInput) *types.MsgUpdateDid {
	var signatures []*types.SignInfo
	signingInput := payload.GetSignBytes()

	for _, skey := range keys {
		signature := base64.StdEncoding.EncodeToString(ed25519.Sign(skey.Key, signingInput))
		signatures = append(signatures, &types.SignInfo{
			VerificationMethodId: skey.VerificationMethodId,
			Signature:            signature,
		})
	}

	return &types.MsgUpdateDid{
		Payload:    payload,
		Signatures: signatures,
	}
}

func (s *TestSetup) InitDid(did string) (map[string]ed25519.PrivateKey, *types.MsgCreateDidPayload, error) {
	pubKey, privKey, _ := ed25519.GenerateKey(rand.Reader)
	keyId := did + "#key-1"

	// add new Did
	didMsg := s.BuildMsgCreateDidPayload(did, keyId, pubKey)

	keys := map[string]ed25519.PrivateKey{keyId: privKey}

	_, err := s.MsgServer.CreateDid(s.StdCtx, s.WrapCreateRequest(didMsg, keys))
	if err != nil {
		return nil, nil, err
	}

	return keys, didMsg, nil
}

func (s *TestSetup) SendUpdateDid(msg *types.MsgUpdateDidPayload, keys []SignInput) (*types.Did, error) {
	req := types.QueryGetDidRequest{
		Id: msg.Id,
	}

	state, _ := s.QueryServer.Did(s.StdCtx, &req)
	if len(msg.VersionId) == 0 {
		msg.VersionId = state.Metadata.VersionId
	}

	_, err := s.MsgServer.UpdateDid(s.StdCtx, s.WrapUpdateRequest(msg, keys))
	if err != nil {
		return nil, err
	}

	req = types.QueryGetDidRequest{
		Id: msg.Id,
	}

	updated, _ := s.QueryServer.Did(s.StdCtx, &req)
	return updated.Did, nil
}

func (s *TestSetup) SendCreateDid(msg *types.MsgCreateDidPayload, keys map[string]ed25519.PrivateKey) (*types.Did, error) {
	_, err := s.MsgServer.CreateDid(s.StdCtx, s.WrapCreateRequest(msg, keys))
	if err != nil {
		return nil, err
	}

	req := types.QueryGetDidRequest{
		Id: msg.Id,
	}

	created, _ := s.QueryServer.Did(s.StdCtx, &req)
	return created.Did, nil
}

func ConcatKeys(dst map[string]ed25519.PrivateKey, src map[string]ed25519.PrivateKey) map[string]ed25519.PrivateKey {
	for k, v := range src {
		dst[k] = v
	}

	return dst
}

func MapToListOfSignerKeys(mp map[string]ed25519.PrivateKey) []SignInput {
	rlist := []SignInput{}
	for k, v := range mp {
		rlist = append(rlist, SignInput{
			VerificationMethodId: k,
			Key:                  v,
		})
	}
	return rlist
}

func (s TestSetup) CreateTestDIDs(keys map[string]KeyPair) error {
	testDIDs := []struct {
		signers []string
		msg     *types.MsgCreateDidPayload
	}{
		{
			signers: []string{AliceKey1},
			msg: &types.MsgCreateDidPayload{
				Id:             AliceDID,
				Authentication: []string{AliceKey1},
				VerificationMethod: []*types.VerificationMethod{
					{
						Id:         AliceKey1,
						Type:       types.Ed25519VerificationKey2020,
						Controller: AliceDID,
					},
				},
			},
		},
		{
			signers: []string{BobKey2},
			msg: &types.MsgCreateDidPayload{
				Id: BobDID,
				Authentication: []string{
					BobKey1,
					BobKey2,
					BobKey3,
				},
				CapabilityDelegation: []string{
					BobKey4,
				},
				VerificationMethod: []*types.VerificationMethod{
					{
						Id:         BobKey1,
						Type:       types.Ed25519VerificationKey2020,
						Controller: BobDID,
					},
					{
						Id:         BobKey2,
						Type:       types.Ed25519VerificationKey2020,
						Controller: BobDID,
					},
					{
						Id:         BobKey3,
						Type:       types.Ed25519VerificationKey2020,
						Controller: BobDID,
					},
					{
						Id:         BobKey4,
						Type:       types.Ed25519VerificationKey2020,
						Controller: BobDID,
					},
				},
			},
		},
		{
			signers: []string{CharlieKey2, BobKey2},
			msg: &types.MsgCreateDidPayload{
				Id: CharlieDID,
				Authentication: []string{
					CharlieKey1,
					CharlieKey2,
					CharlieKey3,
				},
				VerificationMethod: []*types.VerificationMethod{
					{
						Id:         CharlieKey1,
						Type:       types.Ed25519VerificationKey2020,
						Controller: BobDID,
					},
					{
						Id:         CharlieKey2,
						Type:       types.Ed25519VerificationKey2020,
						Controller: BobDID,
					},
					{
						Id:         CharlieKey3,
						Type:       types.Ed25519VerificationKey2020,
						Controller: BobDID,
					},
				},
			},
		},
	}

	for _, prefilled := range testDIDs {
		msg := prefilled.msg

		for _, vm := range msg.VerificationMethod {
			encoded, err := multibase.Encode(multibase.Base58BTC, keys[vm.Id].Public)
			if err != nil {
				return err
			}
			vm.PublicKeyMultibase = encoded
		}

		signerKeys := map[string]ed25519.PrivateKey{}
		for _, signer := range prefilled.signers {
			signerKeys[signer] = keys[signer].Private
		}

		_, err := s.SendCreateDid(msg, signerKeys)
		if err != nil {
			return err
		}
	}

	return nil
}
