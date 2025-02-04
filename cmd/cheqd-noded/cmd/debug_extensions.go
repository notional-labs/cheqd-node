package cmd

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/multiformats/go-multibase"
	"github.com/spf13/cobra"
)

func extendDebug(debugCmd *cobra.Command) *cobra.Command {
	debugCmd.AddCommand(ed25519Cmd(),
		encodingCmd())

	return debugCmd
}

// ed25519Cmd returns cobra Command.
func ed25519Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ed25519",
		Short: "ed25519 tools",
	}

	cmd.AddCommand(ed25519RandomCmd())

	return cmd
}

// ed25519Cmd returns cobra Command.
func ed25519RandomCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "random",
		Short: "Generate random ed25519 keypair",
		RunE: func(cmd *cobra.Command, args []string) error {
			pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
			if err != nil {
				return err
			}

			keyInfo := struct {
				PubKeyBase64    	string `json:"pub_key_base_64"`
				PubKeyMultibase58   string `json:"pub_key_multibase_58"`
				PrivKeyBase64   	string `json:"priv_key_base_64"`
			}{
				PubKeyBase64: base64.StdEncoding.EncodeToString(pubKey),
				PrivKeyBase64: base64.StdEncoding.EncodeToString(privKey),
			}

			keyInfoJson, err := json.Marshal(keyInfo)
			if err != nil {
				return err
			}

			_, err = fmt.Fprintln(cmd.OutOrStdout(), string(keyInfoJson))
			return err
		},
	}

	return cmd
}

// encoding returns cobra Command.
func encodingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "encoding",
		Short: "Encoding tools",
	}

	cmd.AddCommand(base64toMultibase58Cmd())

	return cmd
}

// base64toMultibase58Cmd returns cobra Command.
func base64toMultibase58Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "base64-multibase58 [input]",
		Short: "Convert base64 string to multibase58 string",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			base64Str := args[0]
			bytes, err := base64.StdEncoding.DecodeString(base64Str)
			if err != nil {
				return err
			}

			multibase58Str, err := multibase.Encode(multibase.Base58BTC, bytes)
			if err != nil {
				return err
			}

			_, err = fmt.Fprintln(cmd.OutOrStdout(), multibase58Str)
			return err
		},
	}

	return cmd
}
