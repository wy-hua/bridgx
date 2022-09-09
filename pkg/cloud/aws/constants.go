package aws

import (
	"errors"

	"github.com/galaxy-future/BridgX/pkg/cloud"
)

const (
	_maxNumEcsPerOperation = 100
	_pageSize              = 100

	_filterNameLocation     = "location"
	_locationTypeNameRegion = "region"
	_locationTypeNameZone   = "availability-zone"
	_locationTypeNameZoneId = "availability-zone-id"

	_filterNameVpcId           = "vpc-id"
	_filterNameGroupId         = "group-id"
	_filterNameRegionName      = "region-name"
	_filterNameState           = "state"
	_filterNameArchitecture    = "architecture"
	_filterNameInstanceId      = "instance-id"
	_filterNameAttachmentVpcId = "attachment.vpc-id"
	_filterNameInstanceType    = "instance-type"

	_resourceTypeVpc      = "vpc"
	_resourceTypeSubnet   = "subnet"
	_resourceTypeInstance = "instance"
	_resourceTypeEip      = "elastic-ip"

	_tagKeyVpcName    = "vpc-name"
	_tagKeySwitchName = "switch-name"
	_tagKeyEipName    = "eip-name"
)

var (
	_errInstanceIdsEmpty    = errors.New("at least one instance id")
	_errInvalidParameter    = errors.New("invalid parameter")
	_errCodeDryRunOperation = "DryRunOperation"
)

var _imageType = map[string]string{
	cloud.ImageGlobal:  "amazon",
	cloud.ImageShared:  "aws-marketplace",
	cloud.ImagePrivate: "self",
}

var _ecsStatus = map[string]string{
	"pending":       cloud.EcsBuilding,
	"running":       cloud.EcsRunning,
	"shutting-down": cloud.EcsStopped,
	"terminated":    cloud.EcsDeleted,
	"stopping":      cloud.EcsStopping,
	"stopped":       cloud.EcsStopped,
}

var _insTypeStat = map[string]string{
	"normal":    cloud.InsTypeAvailable,
	"promotion": cloud.InsTypeAvailable,
	"":          cloud.InsTypeAvailable,
	"obt":       cloud.InsTypeAvaSoon,
	"abandon":   cloud.InsTypeSellOut,
	"sellout":   cloud.InsTypeSellOut,
}

var _secGrpRuleDirection = map[bool]string{
	false: cloud.SecGroupRuleIn,
	true:  cloud.SecGroupRuleOut,
}

var _vpcStatus = map[string]string{
	"pending":   cloud.VPCStatusPending,
	"available": cloud.VPCStatusAvailable,
}

var _subnetStatus = map[string]string{
	"pending":   cloud.VPCStatusPending,
	"available": cloud.VPCStatusAvailable,
}

var _subnetIsDefault = map[bool]int{
	false: 0,
	true:  1,
}

var _regionLocalName = map[string]string{
	"cn-north-1":     "中国 (北京)",
	"cn-northwest-1": "中国 (宁夏)",
	"us-east-1":      "美国东部 (弗吉尼亚北部)",
	"us-east-2":      "美国东部 (俄亥俄州)",
	"us-west-1":      "美国西部 (加利福尼亚北部)",
	"us-west-2":      "美国西部 (俄勒冈州)",
	"af-south-1":     "非洲 (开普敦)",
	"ap-east-1":      "亚太地区 (香港)",
	"ap-southeast-3": "亚太地区 (雅加达)",
	"ap-south-1":     "亚太地区 (孟买)",
	"ap-northeast-3": "亚太地区 (大阪)",
	"ap-northeast-2": "亚太地区 (首尔)",
	"ap-southeast-1": "亚太地区 (新加坡)",
	"ap-southeast-2": "亚太地区 (悉尼)",
	"ap-northeast-1": "亚太地区 (东京)",
	"ca-central-1":   "加拿大 (中部)",
	"eu-central-1":   "欧洲 (法兰克福)",
	"eu-west-1":      "欧洲 (爱尔兰)",
	"eu-west-2":      "欧洲 (伦敦)",
	"eu-south-1":     "欧洲 (米兰)",
	"eu-west-3":      "欧洲 (巴黎)",
	"eu-north-1":     "欧洲 (斯德哥尔摩)",
	"me-south-1":     "中东 (巴林)",
	"sa-east-1":      "南美洲 (圣保罗)",
}

var _letter = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
