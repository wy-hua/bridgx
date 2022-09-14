package tests

import (
	"github.com/galaxy-future/BridgX/cmd/api/middleware"
	"github.com/galaxy-future/BridgX/cmd/api/routers"
	"github.com/galaxy-future/BridgX/cmd/scheduler/crond"
	"github.com/galaxy-future/BridgX/cmd/scheduler/monitors"
	"github.com/galaxy-future/BridgX/cmd/scheduler/types"
	"github.com/galaxy-future/BridgX/internal/bcc"
	"github.com/galaxy-future/BridgX/internal/constants"
	"github.com/galaxy-future/BridgX/internal/service"
	"github.com/galaxy-future/BridgX/pkg/cloud"
	"github.com/gin-gonic/gin"
	"os"
	"testing"

	"github.com/galaxy-future/BridgX/config"
	"github.com/galaxy-future/BridgX/internal/cache"
	"github.com/galaxy-future/BridgX/internal/clients"
	"github.com/galaxy-future/BridgX/internal/logs"
)

var r *gin.Engine

const (
	_v1Api = "/api/v1/"
	_Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJuYW1lIjoicm9vdCIsInVzZXJfdHlwZSI6IkFETUlOIiwib3JnX2lkIjoxLCJleHAiOjE2NjMxNTMyNzIsIm5iZiI6MTY2MzEyNDQ2Mn0.1wJ4tiV09S_eUCVxog_uK2Xaivo6tUR_PJP9w4lZ5bA" // JWT token
)

func TestMain(m *testing.M) {
	//因为是相对路径，需要把conf文件copy到tests目录下
	config.MustInit()
	logs.Init()
	clients.MustInit()
	bcc.MustInit(config.GlobalConfig)
	cache.MustInit()
	service.Init(100)
	middleware.Init()
	initScheduler()
	r = routers.Init()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func initScheduler() {
	crond.Init()
	err := Init()
	if err != nil {
		return
	}
	go func() {
		Run()
	}()
}

var schedulers []*types.Scheduler

func Init() error {
	locker, err := clients.NewEtcdClient(config.GlobalConfig.EtcdConfig)
	if err != nil {
		return err
	}
	schedulers = []*types.Scheduler{
		{
			//扫库，查看是否有待执行的Task，分配Task到WorkerPool
			Interval: constants.DefaultTaskMonitorInterval,
			Monitor: &monitors.TaskMonitor{
				LockerClient: locker,
			},
		},
		{
			Interval: constants.DefaultKillExpireRunningTaskInterval,
			Monitor:  &monitors.TaskKiller{},
		},
	}
	return nil
}

func Run() {
	for _, s := range schedulers {
		crond.AddFixedIntervalSecondsJob(s.Interval, s.Monitor)
	}
	crond.Run()
}
func AKGenerator(provider string) (ak string) {
	switch provider {
	case cloud.AlibabaCloud:
		ak = "xxx"
	case cloud.HuaweiCloud:
		ak = "xxx"
	case cloud.TencentCloud:
		ak = "xxx"
	case cloud.BaiduCloud:
		ak = "xxx"
	case cloud.AwsCloud:
		ak = "xxx"
	}
	return
}
