package service

import (
	"context"
	"time"

	"github.com/galaxy-future/BridgX/internal/clients"
	"github.com/galaxy-future/BridgX/internal/errs"
	"github.com/galaxy-future/BridgX/internal/model"
	"github.com/galaxy-future/BridgX/pkg/cmp"
	jsoniter "github.com/json-iterator/go"
)

type Category uint8

const (
	categoryBegin Category = iota
	ExpandAndReduce
	ClusterManage
	K8sManage
	AppCenter
	CloudAccountManage
	AccountManage
	categoryEnd
)

func (t Category) IsValid() bool {
	return t > categoryBegin && t < categoryEnd
}

type Action string

const (
	ActionCreate  Action = "CREATE"
	ActionDelete         = "DELETE"
	ActionUpdate         = "UPDATE"
	ActionPublish        = "PUBLISH"
	ActionExpand         = "EXPAND"
	ActionReduce         = "REDUCE"
)

var actionMap = map[Action]bool{
	ActionCreate:  true,
	ActionDelete:  true,
	ActionUpdate:  true,
	ActionPublish: true,
	ActionExpand:  true,
	ActionReduce:  true,
}

func (t Action) IsValid() bool {
	return actionMap[t]
}

type Object string

const (
	CloudCluster Object = "CLOUD_CLUSTER"
	K8sCluster          = "K8S_CLUSTER"
	ComputingRes        = "COMPUTING_RESOURCE"
	RunningEnv          = "RUNNING_ENV"
)

var objectMap = map[Object]bool{
	CloudCluster: true,
	K8sCluster:   true,
	ComputingRes: true,
	RunningEnv:   true,
}

func (t Object) IsValid() bool {
	return objectMap[t]
}

type OperationLog struct {
	Category    Category    `json:"category"`
	Action      Action      `json:"action"`
	Object      Object      `json:"object"`
	ObjectValue string      `json:"object_value"`
	Operator    string      `json:"operator"`
	Detail      interface{} `json:"detail"`
	Old         interface{} `json:"old_value"`
}

func RecordOperationLog(ctx context.Context, oplog OperationLog) (err error) {
	if !oplog.Category.IsValid() || !oplog.Action.IsValid() || !oplog.Object.IsValid() || oplog.Operator == "" {
		return errs.ErrParam
	}

	detail := ""
	if oplog.Detail != nil {
		var diff cmp.DiffResult
		switch oplog.Action {
		case ActionCreate:
			if diff, err = cmp.Diff(nil, oplog.Detail); err != nil {
				return err
			}
		case ActionUpdate:
			if oplog.Old == nil {
				return errs.ErrParam
			}
			if diff, err = cmp.Diff(oplog.Old, oplog.Detail); err != nil {
				return err
			}
		default:
			if detail, err = jsoniter.MarshalToString(oplog.Detail); err != nil {
				return err
			}
		}

		if len(diff.Fields) > 0 {
			if detail, err = jsoniter.MarshalToString(diff.Fields); err != nil {
				return err
			}
		}
	}

	return clients.WriteDBCli.Create(&model.OperationLog{
		Category:    uint8(oplog.Category),
		Action:      string(oplog.Action),
		Object:      string(oplog.Object),
		ObjectValue: oplog.ObjectValue,
		Operator:    oplog.Operator,
		Detail:      detail,
	}).Error
}

type ExtractCondition struct {
	Operators  []string  `json:"operators" form:"operators"`
	Actions    []string  `json:"actions" form:"actions"`
	TimeStart  time.Time `json:"time_start" form:"time_start"`
	TimeEnd    time.Time `json:"time_end" form:"time_end"`
	PageNumber int       `json:"page_number" form:"page_number"`
	PageSize   int       `json:"page_size" form:"page_size"`
}

func ExtractLogs(ctx context.Context, conds ExtractCondition) ([]model.OperationLog, int64, error) {
	query := clients.ReadDBCli.WithContext(ctx).Model(OperationLog{})
	if len(conds.Operators) > 0 {
		query.Where("operator IN (?)", conds.Operators)
	}
	if len(conds.Actions) > 0 {
		query.Where("action IN (?)", conds.Actions)
	}
	if !conds.TimeStart.IsZero() {
		query.Where("create_at >= ?", conds.TimeStart)
	}
	if !conds.TimeEnd.IsZero() {
		query.Where("create_at < ?", conds.TimeEnd)
	}

	var logs []model.OperationLog
	count, err := model.QueryWhere(query, conds.PageNumber, conds.PageSize, &logs, "", true)
	if err != nil {
		return nil, 0, err
	}
	return logs, count, nil
}
