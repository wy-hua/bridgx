package tests

import (
	"context"
	"testing"

	"github.com/galaxy-future/BridgX/internal/model"
	"github.com/galaxy-future/BridgX/internal/service"
)

func TestRecordOperationLog(t *testing.T) {
	type args struct {
		ctx   context.Context
		oplog service.OperationLog
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "diff",
			args: args{
				ctx: nil,
				oplog: service.OperationLog{
					Category:    service.ClusterManage,
					Action:      service.ActionUpdate,
					Object:      service.CloudCluster,
					ObjectValue: "gf.bridgx.online",
					Operator:    "root",
					Detail: model.Cluster{
						AccountKey: "key",
					},
					Old: model.Cluster{
						ExpectCount: 1,
						ClusterName: "name",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := service.RecordOperationLog(tt.args.ctx, tt.args.oplog); (err != nil) != tt.wantErr {
				t.Errorf("RecordOperationLog() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
