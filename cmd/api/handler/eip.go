package handler

import (
	"net/http"

	"github.com/galaxy-future/BridgX/cmd/api/response"
	"github.com/galaxy-future/BridgX/internal/service"
	"github.com/galaxy-future/BridgX/pkg/cloud"
	"github.com/gin-gonic/gin"
)

type CreateEipReq struct {
	service.Eip
	Num int `json:"num"`
}

func CreateEip(ctx *gin.Context) {
	req := CreateEipReq{}
	if err := ctx.Bind(&req); err != nil {
		response.MkResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	if req.Charge == nil || req.Num < 1 {
		response.MkResponse(ctx, http.StatusBadRequest, response.ParamInvalid, nil)
		return
	}

	if err := req.CreateEip(ctx, req.Num); err != nil {
		response.MkResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.MkResponse(ctx, http.StatusOK, response.Success, nil)
	return
}

type DescribeEipRsp struct {
	List []cloud.Eip `json:"list"`
	response.Pager
}

func DescribeEip(ctx *gin.Context) {
	req := service.Eip{}
	if err := ctx.Bind(&req); err != nil {
		response.MkResponse(ctx, http.StatusBadRequest, response.ParamInvalid, nil)
		return
	}
	pageNumber, pageSize := getPager(ctx)

	rsp, err := req.DescribeEip(ctx, pageNumber, pageSize)
	if err != nil {
		response.MkResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	ret := DescribeEipRsp{
		List: rsp.List,
		Pager: response.Pager{
			PageNumber: pageNumber,
			PageSize:   pageSize,
			Total:      rsp.TotalCount,
		},
	}
	response.MkResponse(ctx, http.StatusOK, response.Success, ret)
	return
}

func BindEip(ctx *gin.Context) {
	req := service.Eip{}
	if err := ctx.Bind(&req); err != nil {
		response.MkResponse(ctx, http.StatusBadRequest, response.ParamInvalid, nil)
		return
	}
	if req.Id == "" || req.InstanceId == "" {
		response.MkResponse(ctx, http.StatusBadRequest, response.ParamInvalid, nil)
		return
	}

	if err := req.BindEip(ctx); err != nil {
		response.MkResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.MkResponse(ctx, http.StatusOK, response.Success, nil)
	return
}

func UnBindEip(ctx *gin.Context) {
	req := service.Eip{}
	if err := ctx.Bind(&req); err != nil {
		response.MkResponse(ctx, http.StatusBadRequest, response.ParamInvalid, nil)
		return
	}
	if req.Id == "" || req.InstanceId == "" {
		response.MkResponse(ctx, http.StatusBadRequest, response.ParamInvalid, nil)
		return
	}

	if err := req.UnBindEip(ctx); err != nil {
		response.MkResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.MkResponse(ctx, http.StatusOK, response.Success, nil)
	return
}

func ConvertPublicIp2Eip(ctx *gin.Context) {
	req := service.Eip{}
	if err := ctx.Bind(&req); err != nil {
		response.MkResponse(ctx, http.StatusBadRequest, response.ParamInvalid, nil)
		return
	}
	if req.InstanceId == "" {
		response.MkResponse(ctx, http.StatusBadRequest, response.ParamInvalid, nil)
		return
	}

	if err := req.ConvertPublicIp2Eip(ctx); err != nil {
		response.MkResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.MkResponse(ctx, http.StatusOK, response.Success, nil)
	return
}
