package cloud

type Provider interface {
	BatchCreate(m Params, num int) (instanceIds []string, err error)
	ProviderType() string
	GetInstances(ids []string) (instances []Instance, err error)
	GetInstancesByTags(region string, tags []Tag) (instances []Instance, err error)
	GetInstancesByCluster(regionId, clusterName string) (instances []Instance, err error)
	BatchDelete(ids []string, regionId string) error
	StartInstances(ids []string) error
	StopInstances(ids []string) error
	CreateVPC(req CreateVpcRequest) (CreateVpcResponse, error)
	GetVPC(req GetVpcRequest) (GetVpcResponse, error)
	CreateSwitch(req CreateSwitchRequest) (CreateSwitchResponse, error)
	GetSwitch(req GetSwitchRequest) (GetSwitchResponse, error)
	CreateSecurityGroup(req CreateSecurityGroupRequest) (CreateSecurityGroupResponse, error)
	AddIngressSecurityGroupRule(req AddSecurityGroupRuleRequest) error
	AddEgressSecurityGroupRule(req AddSecurityGroupRuleRequest) error
	DescribeSecurityGroups(req DescribeSecurityGroupsRequest) (DescribeSecurityGroupsResponse, error)
	GetRegions() (GetRegionsResponse, error)
	GetZones(req GetZonesRequest) (GetZonesResponse, error)
	DescribeAvailableResource(req DescribeAvailableResourceRequest) (DescribeAvailableResourceResponse, error)
	DescribeInstanceTypes(req DescribeInstanceTypesRequest) (DescribeInstanceTypesResponse, error)
	DescribeImages(req DescribeImagesRequest) (DescribeImagesResponse, error)
	DescribeVpcs(req DescribeVpcsRequest) (DescribeVpcsResponse, error)
	DescribeSwitches(req DescribeSwitchesRequest) (DescribeSwitchesResponse, error)
	DescribeGroupRules(req DescribeGroupRulesRequest) (DescribeGroupRulesResponse, error)
	GetOrders(req GetOrdersRequest) (GetOrdersResponse, error)
	CreateKeyPair(req CreateKeyPairRequest) (CreateKeyPairResponse, error)
	ImportKeyPair(req ImportKeyPairRequest) (ImportKeyPairResponse, error)
	DescribeKeyPairs(req DescribeKeyPairsRequest) (DescribeKeyPairsResponse, error)
}

type ProviderDriverFunc func(keyId ...string) (Provider, error)

var registeredPlugins = map[string]ProviderDriverFunc{}

func RegisterProviderDriver(name string, f ProviderDriverFunc) {
	registeredPlugins[name] = f
}
