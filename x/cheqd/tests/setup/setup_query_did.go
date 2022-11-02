package setup

import "github.com/cheqd/cheqd-node/x/cheqd/types"

func (s *TestSetup) QueryDid(did string) (*types.QueryGetDidDocResponse, error) {
	req := &types.QueryGetDidDocRequest{
		Id: did,
	}

	return s.QueryServer.DidDoc(s.StdCtx, req)
}
