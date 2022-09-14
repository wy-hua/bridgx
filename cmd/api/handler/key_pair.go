package handler

import (
	"net/http"

	"github.com/galaxy-future/BridgX/internal/types"

	"github.com/galaxy-future/BridgX/cmd/api/helper"
	"github.com/galaxy-future/BridgX/cmd/api/request"
	"github.com/galaxy-future/BridgX/cmd/api/response"
	"github.com/galaxy-future/BridgX/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func CreateKeyPair(ctx *gin.Context) {
	req := request.CreateKeyPairRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.MkResponse(ctx, http.StatusBadRequest, response.ParamInvalid, nil)
		return
	}
	err := service.CreateKeyPair(ctx, req.AK, req.Provider, req.RegionId, req.KeyPairName)
	if err != nil {
		response.MkResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.MkResponse(ctx, http.StatusOK, response.Success, nil)
	return
}

func ImportKeyPair(ctx *gin.Context) {
	req := request.ImportKeyPairRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.MkResponse(ctx, http.StatusBadRequest, response.ParamInvalid, nil)
		return
	}
	err := service.ImportKeyPair(ctx, req.AK, req.Provider, req.RegionId, req.KeyPairName, req.PublicKey, req.PrivateKey)
	if err != nil {
		response.MkResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.MkResponse(ctx, http.StatusOK, response.Success, nil)
	return
}

func GetKeyPair(ctx *gin.Context) {
	id := ctx.Param("id")
	keyPair, err := service.GetKeyPair(ctx, cast.ToInt64(id))
	if err != nil {
		response.MkResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	resp := map[string]interface{}{"key_pair": helper.ConvertToKeyPairInfo(keyPair)}
	response.MkResponse(ctx, http.StatusOK, response.Success, resp)
	return
}

func ListKeyPairs(ctx *gin.Context) {
	req := request.ListKeyPairRequest{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.MkResponse(ctx, http.StatusBadRequest, response.ParamInvalid, nil)
		return
	}
	keyPairs, pageRsp, err := service.ListKeyPairs(ctx, req.AK, req.Provider, req.RegionId, types.AdjustPage(req.PageReq))
	if err != nil {
		response.MkResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	nextMarker := ""
	total := 0
	if pageRsp != nil {
		nextMarker = pageRsp.NextMarker
		total = pageRsp.Total
	}
	pager := response.Pager{
		NextMarker: nextMarker,
		PageNumber: req.PageNum,
		PageSize:   req.PageSize,
		Total:      total,
	}
	resp := &response.KeyPairListResponse{
		KeyPairList: helper.ConvertToKeyPairList(keyPairs),
		Pager:       pager,
	}
	response.MkResponse(ctx, http.StatusOK, response.Success, resp)
	return
}
