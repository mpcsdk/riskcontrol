// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

type (
	INrpcClient interface {
		Flush()
	}
)

var (
	localNrpcClient INrpcClient
)

func NrpcClient() INrpcClient {
	if localNrpcClient == nil {
		panic("implement not found for interface INrpcClient, forgot register?")
	}
	return localNrpcClient
}

func RegisterNrpcClient(i INrpcClient) {
	localNrpcClient = i
}
