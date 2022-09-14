package types

import "github.com/galaxy-future/BridgX/internal/constants"

type PageReq struct {
	Marker   string `json:"marker" form:"marker"`
	PageNum  int    `json:"page_number" form:"page_number"`
	PageSize int    `json:"page_size" form:"page_size"`
}

func AdjustPage(page PageReq) PageReq {
	if page.PageNum < 1 {
		page.PageNum = 1
	}

	if page.PageSize < 1 || page.PageSize > constants.DefaultPageSize {
		page.PageSize = constants.DefaultPageSize
	}
	return page
}

type PageRsp struct {
	NextMarker string `json:"next_marker"`
	Total      int    `json:"total"`
}
