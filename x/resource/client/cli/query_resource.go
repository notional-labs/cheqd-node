package cli

//func CmdGetDid() *cobra.Command {
//	cmd := &cobra.Command{
//		Use:   "did [id]",
//		Short: "Query a did",
//		Args:  cobra.ExactArgs(1),
//		RunE: func(cmd *cobra.Command, args []string) error {
//			clientCtx := client.GetClientContextFromCmd(cmd)
//
//			queryClient := types.NewQueryClient(clientCtx)
//
//			did := args[0]
//			params := &types.QueryGetDidRequest{
//				Id: did,
//			}
//
//			resp, err := queryClient.Did(context.Background(), params)
//			if err != nil {
//				return err
//			}
//
//			return clientCtx.PrintProto(resp)
//		},
//	}
//
//	flags.AddQueryFlagsToCmd(cmd)
//
//	return cmd
//}
