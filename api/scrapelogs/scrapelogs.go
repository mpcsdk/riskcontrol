// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package scrapelogs

import (
	"context"
	
	"riskcontral/api/scrapelogs/v1"
)

type IScrapelogsV1 interface {
	NftCnt(ctx context.Context, req *v1.NftCntReq) (res *v1.NftCntRes, err error)
	FtCnt(ctx context.Context, req *v1.FtCntReq) (res *v1.FtCntRes, err error)
}


