package baidu

import (
	"github.com/galaxy-future/BridgX/pkg/cloud"
)

var _inEcsChargeType = map[string]string{
	cloud.InstanceChargeTypePrePaid:  "Prepaid",
	cloud.InstanceChargeTypePostPaid: "Postpaid",
}

var _imageType = map[string]string{
	cloud.ImageGlobal:  "System",
	cloud.ImageShared:  "Sharing",
	cloud.ImagePrivate: "Custom",
}

//out
var _ecsChargeType = map[string]string{
	"Prepaid":  cloud.InstanceChargeTypePrePaid,
	"Postpaid": cloud.InstanceChargeTypePostPaid,
	"prepay":   cloud.InstanceChargeTypePrePaid,
	"postpay":  cloud.InstanceChargeTypePostPaid,
}

var _insTypeChargeType = map[string]string{
	"Prepaid":  cloud.InsTypeChargeTypePrePaid,
	"Postpaid": cloud.InsTypeChargeTypePostPaid,
	"both":     cloud.InsTypeChargeTypeAll,
}

var _ecsStatus = map[string]string{
	"Starting":           cloud.EcsStarting,
	"Running":            cloud.EcsRunning,
	"Stopping":           cloud.EcsStopping,
	"Stopped":            cloud.EcsStopped,
	"Deleted":            cloud.EcsDeleted,
	"Scaling":            cloud.EcsStarting,
	"Expired":            cloud.EcsAbnormal,
	"Error":              cloud.EcsAbnormal,
	"SnapshotProcessing": cloud.EcsStarting,
	"ImageProcessing":    cloud.EcsStarting,
	"Recharging":         cloud.EcsStarting,
}
