package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type SystemResourceAccountingTemplate struct {
	Inst struct {
		AppResources SystemResourceAccountingTemplateAppResources1702 `json:"app-resources"`

		Name string `json:"name"`

		NetworkResources SystemResourceAccountingTemplateNetworkResources1734 `json:"network-resources"`

		SystemResources SystemResourceAccountingTemplateSystemResources1744 `json:"system-resources"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"template"`
}

type SystemResourceAccountingTemplateAppResources1702 struct {
	GslbDeviceCfg            SystemResourceAccountingTemplateAppResourcesGslbDeviceCfg1703            `json:"gslb-device-cfg"`
	GslbGeoLocationCfg       SystemResourceAccountingTemplateAppResourcesGslbGeoLocationCfg1704       `json:"gslb-geo-location-cfg"`
	GslbIpListCfg            SystemResourceAccountingTemplateAppResourcesGslbIpListCfg1705            `json:"gslb-ip-list-cfg"`
	GslbPolicyCfg            SystemResourceAccountingTemplateAppResourcesGslbPolicyCfg1706            `json:"gslb-policy-cfg"`
	GslbServiceCfg           SystemResourceAccountingTemplateAppResourcesGslbServiceCfg1707           `json:"gslb-service-cfg"`
	GslbServiceIpCfg         SystemResourceAccountingTemplateAppResourcesGslbServiceIpCfg1708         `json:"gslb-service-ip-cfg"`
	GslbServicePortCfg       SystemResourceAccountingTemplateAppResourcesGslbServicePortCfg1709       `json:"gslb-service-port-cfg"`
	GslbSiteCfg              SystemResourceAccountingTemplateAppResourcesGslbSiteCfg1710              `json:"gslb-site-cfg"`
	GslbSvcGroupCfg          SystemResourceAccountingTemplateAppResourcesGslbSvcGroupCfg1711          `json:"gslb-svc-group-cfg"`
	GslbTemplateCfg          SystemResourceAccountingTemplateAppResourcesGslbTemplateCfg1712          `json:"gslb-template-cfg"`
	GslbZoneCfg              SystemResourceAccountingTemplateAppResourcesGslbZoneCfg1713              `json:"gslb-zone-cfg"`
	HealthMonitorCfg         SystemResourceAccountingTemplateAppResourcesHealthMonitorCfg1714         `json:"health-monitor-cfg"`
	RealPortCfg              SystemResourceAccountingTemplateAppResourcesRealPortCfg1715              `json:"real-port-cfg"`
	RealServerCfg            SystemResourceAccountingTemplateAppResourcesRealServerCfg1716            `json:"real-server-cfg"`
	ServiceGroupCfg          SystemResourceAccountingTemplateAppResourcesServiceGroupCfg1717          `json:"service-group-cfg"`
	VirtualServerCfg         SystemResourceAccountingTemplateAppResourcesVirtualServerCfg1718         `json:"virtual-server-cfg"`
	VirtualPortCfg           SystemResourceAccountingTemplateAppResourcesVirtualPortCfg1719           `json:"virtual-port-cfg"`
	CacheTemplateCfg         SystemResourceAccountingTemplateAppResourcesCacheTemplateCfg1720         `json:"cache-template-cfg"`
	ClientSslTemplateCfg     SystemResourceAccountingTemplateAppResourcesClientSslTemplateCfg1721     `json:"client-ssl-template-cfg"`
	ConnReuseTemplateCfg     SystemResourceAccountingTemplateAppResourcesConnReuseTemplateCfg1722     `json:"conn-reuse-template-cfg"`
	FastTcpTemplateCfg       SystemResourceAccountingTemplateAppResourcesFastTcpTemplateCfg1723       `json:"fast-tcp-template-cfg"`
	FastUdpTemplateCfg       SystemResourceAccountingTemplateAppResourcesFastUdpTemplateCfg1724       `json:"fast-udp-template-cfg"`
	FixTemplateCfg           SystemResourceAccountingTemplateAppResourcesFixTemplateCfg1725           `json:"fix-template-cfg"`
	HttpTemplateCfg          SystemResourceAccountingTemplateAppResourcesHttpTemplateCfg1726          `json:"http-template-cfg"`
	LinkCostTemplateCfg      SystemResourceAccountingTemplateAppResourcesLinkCostTemplateCfg1727      `json:"link-cost-template-cfg"`
	PbslbEntryCfg            SystemResourceAccountingTemplateAppResourcesPbslbEntryCfg1728            `json:"pbslb-entry-cfg"`
	PersistCookieTemplateCfg SystemResourceAccountingTemplateAppResourcesPersistCookieTemplateCfg1729 `json:"persist-cookie-template-cfg"`
	PersistSrcipTemplateCfg  SystemResourceAccountingTemplateAppResourcesPersistSrcipTemplateCfg1730  `json:"persist-srcip-template-cfg"`
	ServerSslTemplateCfg     SystemResourceAccountingTemplateAppResourcesServerSslTemplateCfg1731     `json:"server-ssl-template-cfg"`
	ProxyTemplateCfg         SystemResourceAccountingTemplateAppResourcesProxyTemplateCfg1732         `json:"proxy-template-cfg"`
	StreamTemplateCfg        SystemResourceAccountingTemplateAppResourcesStreamTemplateCfg1733        `json:"stream-template-cfg"`
	Threshold                int                                                                      `json:"threshold"`
	Uuid                     string                                                                   `json:"uuid"`
}

type SystemResourceAccountingTemplateAppResourcesGslbDeviceCfg1703 struct {
	GslbDeviceMax          int `json:"gslb-device-max"`
	GslbDeviceMinGuarantee int `json:"gslb-device-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesGslbGeoLocationCfg1704 struct {
	GslbGeoLocationMax          int `json:"gslb-geo-location-max"`
	GslbGeoLocationMinGuarantee int `json:"gslb-geo-location-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesGslbIpListCfg1705 struct {
	GslbIpListMax          int `json:"gslb-ip-list-max"`
	GslbIpListMinGuarantee int `json:"gslb-ip-list-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesGslbPolicyCfg1706 struct {
	GslbPolicyMax          int `json:"gslb-policy-max"`
	GslbPolicyMinGuarantee int `json:"gslb-policy-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesGslbServiceCfg1707 struct {
	GslbServiceMax          int `json:"gslb-service-max"`
	GslbServiceMinGuarantee int `json:"gslb-service-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesGslbServiceIpCfg1708 struct {
	GslbServiceIpMax          int `json:"gslb-service-ip-max"`
	GslbServiceIpMinGuarantee int `json:"gslb-service-ip-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesGslbServicePortCfg1709 struct {
	GslbServicePortMax          int `json:"gslb-service-port-max"`
	GslbServicePortMinGuarantee int `json:"gslb-service-port-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesGslbSiteCfg1710 struct {
	GslbSiteMax          int `json:"gslb-site-max"`
	GslbSiteMinGuarantee int `json:"gslb-site-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesGslbSvcGroupCfg1711 struct {
	GslbSvcGroupMax          int `json:"gslb-svc-group-max"`
	GslbSvcGroupMinGuarantee int `json:"gslb-svc-group-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesGslbTemplateCfg1712 struct {
	GslbTemplateMax          int `json:"gslb-template-max"`
	GslbTemplateMinGuarantee int `json:"gslb-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesGslbZoneCfg1713 struct {
	GslbZoneMax          int `json:"gslb-zone-max"`
	GslbZoneMinGuarantee int `json:"gslb-zone-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesHealthMonitorCfg1714 struct {
	HealthMonitorMax          int `json:"health-monitor-max"`
	HealthMonitorMinGuarantee int `json:"health-monitor-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesRealPortCfg1715 struct {
	RealPortMax          int `json:"real-port-max"`
	RealPortMinGuarantee int `json:"real-port-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesRealServerCfg1716 struct {
	RealServerMax          int `json:"real-server-max"`
	RealServerMinGuarantee int `json:"real-server-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesServiceGroupCfg1717 struct {
	ServiceGroupMax          int `json:"service-group-max"`
	ServiceGroupMinGuarantee int `json:"service-group-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesVirtualServerCfg1718 struct {
	VirtualServerMax          int `json:"virtual-server-max"`
	VirtualServerMinGuarantee int `json:"virtual-server-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesVirtualPortCfg1719 struct {
	VirtualPortMax          int `json:"virtual-port-max"`
	VirtualPortMinGuarantee int `json:"virtual-port-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesCacheTemplateCfg1720 struct {
	CacheTemplateMax          int `json:"cache-template-max"`
	CacheTemplateMinGuarantee int `json:"cache-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesClientSslTemplateCfg1721 struct {
	ClientSslTemplateMax          int `json:"client-ssl-template-max"`
	ClientSslTemplateMinGuarantee int `json:"client-ssl-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesConnReuseTemplateCfg1722 struct {
	ConnReuseTemplateMax          int `json:"conn-reuse-template-max"`
	ConnReuseTemplateMinGuarantee int `json:"conn-reuse-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesFastTcpTemplateCfg1723 struct {
	FastTcpTemplateMax          int `json:"fast-tcp-template-max"`
	FastTcpTemplateMinGuarantee int `json:"fast-tcp-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesFastUdpTemplateCfg1724 struct {
	FastUdpTemplateMax          int `json:"fast-udp-template-max"`
	FastUdpTemplateMinGuarantee int `json:"fast-udp-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesFixTemplateCfg1725 struct {
	FixTemplateMax          int `json:"fix-template-max"`
	FixTemplateMinGuarantee int `json:"fix-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesHttpTemplateCfg1726 struct {
	HttpTemplateMax          int `json:"http-template-max"`
	HttpTemplateMinGuarantee int `json:"http-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesLinkCostTemplateCfg1727 struct {
	LinkCostTemplateMax          int `json:"link-cost-template-max"`
	LinkCostTemplateMinGuarantee int `json:"link-cost-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesPbslbEntryCfg1728 struct {
	PbslbEntryMax          int `json:"pbslb-entry-max"`
	PbslbEntryMinGuarantee int `json:"pbslb-entry-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesPersistCookieTemplateCfg1729 struct {
	PersistCookieTemplateMax          int `json:"persist-cookie-template-max"`
	PersistCookieTemplateMinGuarantee int `json:"persist-cookie-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesPersistSrcipTemplateCfg1730 struct {
	PersistSrcipTemplateMax          int `json:"persist-srcip-template-max"`
	PersistSrcipTemplateMinGuarantee int `json:"persist-srcip-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesServerSslTemplateCfg1731 struct {
	ServerSslTemplateMax          int `json:"server-ssl-template-max"`
	ServerSslTemplateMinGuarantee int `json:"server-ssl-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesProxyTemplateCfg1732 struct {
	ProxyTemplateMax          int `json:"proxy-template-max"`
	ProxyTemplateMinGuarantee int `json:"proxy-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesStreamTemplateCfg1733 struct {
	StreamTemplateMax          int `json:"stream-template-max"`
	StreamTemplateMinGuarantee int `json:"stream-template-min-guarantee"`
}

type SystemResourceAccountingTemplateNetworkResources1734 struct {
	StaticIpv4RouteCfg   SystemResourceAccountingTemplateNetworkResourcesStaticIpv4RouteCfg1735   `json:"static-ipv4-route-cfg"`
	StaticIpv6RouteCfg   SystemResourceAccountingTemplateNetworkResourcesStaticIpv6RouteCfg1736   `json:"static-ipv6-route-cfg"`
	Ipv4AclLineCfg       SystemResourceAccountingTemplateNetworkResourcesIpv4AclLineCfg1737       `json:"ipv4-acl-line-cfg"`
	Ipv6AclLineCfg       SystemResourceAccountingTemplateNetworkResourcesIpv6AclLineCfg1738       `json:"ipv6-acl-line-cfg"`
	StaticArpCfg         SystemResourceAccountingTemplateNetworkResourcesStaticArpCfg1739         `json:"static-arp-cfg"`
	StaticNeighborCfg    SystemResourceAccountingTemplateNetworkResourcesStaticNeighborCfg1740    `json:"static-neighbor-cfg"`
	StaticMacCfg         SystemResourceAccountingTemplateNetworkResourcesStaticMacCfg1741         `json:"static-mac-cfg"`
	ObjectGroupCfg       SystemResourceAccountingTemplateNetworkResourcesObjectGroupCfg1742       `json:"object-group-cfg"`
	ObjectGroupClauseCfg SystemResourceAccountingTemplateNetworkResourcesObjectGroupClauseCfg1743 `json:"object-group-clause-cfg"`
	Threshold            int                                                                      `json:"threshold"`
	Uuid                 string                                                                   `json:"uuid"`
}

type SystemResourceAccountingTemplateNetworkResourcesStaticIpv4RouteCfg1735 struct {
	StaticIpv4RouteMax          int `json:"static-ipv4-route-max"`
	StaticIpv4RouteMinGuarantee int `json:"static-ipv4-route-min-guarantee"`
}

type SystemResourceAccountingTemplateNetworkResourcesStaticIpv6RouteCfg1736 struct {
	StaticIpv6RouteMax          int `json:"static-ipv6-route-max"`
	StaticIpv6RouteMinGuarantee int `json:"static-ipv6-route-min-guarantee"`
}

type SystemResourceAccountingTemplateNetworkResourcesIpv4AclLineCfg1737 struct {
	Ipv4AclLineMax          int `json:"ipv4-acl-line-max"`
	Ipv4AclLineMinGuarantee int `json:"ipv4-acl-line-min-guarantee"`
}

type SystemResourceAccountingTemplateNetworkResourcesIpv6AclLineCfg1738 struct {
	Ipv6AclLineMax          int `json:"ipv6-acl-line-max"`
	Ipv6AclLineMinGuarantee int `json:"ipv6-acl-line-min-guarantee"`
}

type SystemResourceAccountingTemplateNetworkResourcesStaticArpCfg1739 struct {
	StaticArpMax          int `json:"static-arp-max"`
	StaticArpMinGuarantee int `json:"static-arp-min-guarantee"`
}

type SystemResourceAccountingTemplateNetworkResourcesStaticNeighborCfg1740 struct {
	StaticNeighborMax          int `json:"static-neighbor-max"`
	StaticNeighborMinGuarantee int `json:"static-neighbor-min-guarantee"`
}

type SystemResourceAccountingTemplateNetworkResourcesStaticMacCfg1741 struct {
	StaticMacMax          int `json:"static-mac-max"`
	StaticMacMinGuarantee int `json:"static-mac-min-guarantee"`
}

type SystemResourceAccountingTemplateNetworkResourcesObjectGroupCfg1742 struct {
	ObjectGroupMax          int `json:"object-group-max"`
	ObjectGroupMinGuarantee int `json:"object-group-min-guarantee"`
}

type SystemResourceAccountingTemplateNetworkResourcesObjectGroupClauseCfg1743 struct {
	ObjectGroupClauseMax          int `json:"object-group-clause-max"`
	ObjectGroupClauseMinGuarantee int `json:"object-group-clause-min-guarantee"`
}

type SystemResourceAccountingTemplateSystemResources1744 struct {
	BwLimitCfg                SystemResourceAccountingTemplateSystemResourcesBwLimitCfg1745                `json:"bw-limit-cfg"`
	ConcurrentSessionLimitCfg SystemResourceAccountingTemplateSystemResourcesConcurrentSessionLimitCfg1746 `json:"concurrent-session-limit-cfg"`
	L4SessionLimitCfg         SystemResourceAccountingTemplateSystemResourcesL4SessionLimitCfg1747         `json:"l4-session-limit-cfg"`
	L4cpsLimitCfg             SystemResourceAccountingTemplateSystemResourcesL4cpsLimitCfg1748             `json:"l4cps-limit-cfg"`
	L7cpsLimitCfg             SystemResourceAccountingTemplateSystemResourcesL7cpsLimitCfg1749             `json:"l7cps-limit-cfg"`
	NatcpsLimitCfg            SystemResourceAccountingTemplateSystemResourcesNatcpsLimitCfg1750            `json:"natcps-limit-cfg"`
	FwcpsLimitCfg             SystemResourceAccountingTemplateSystemResourcesFwcpsLimitCfg1751             `json:"fwcps-limit-cfg"`
	SslThroughputLimitCfg     SystemResourceAccountingTemplateSystemResourcesSslThroughputLimitCfg1752     `json:"ssl-throughput-limit-cfg"`
	SslcpsLimitCfg            SystemResourceAccountingTemplateSystemResourcesSslcpsLimitCfg1753            `json:"sslcps-limit-cfg"`
	Threshold                 int                                                                          `json:"threshold"`
	Uuid                      string                                                                       `json:"uuid"`
}

type SystemResourceAccountingTemplateSystemResourcesBwLimitCfg1745 struct {
	BwLimitMax              int `json:"bw-limit-max"`
	BwLimitWatermarkDisable int `json:"bw-limit-watermark-disable"`
}

type SystemResourceAccountingTemplateSystemResourcesConcurrentSessionLimitCfg1746 struct {
	ConcurrentSessionLimitMax int `json:"concurrent-session-limit-max"`
}

type SystemResourceAccountingTemplateSystemResourcesL4SessionLimitCfg1747 struct {
	L4SessionLimitMax          string `json:"l4-session-limit-max"`
	L4SessionLimitMinGuarantee string `json:"l4-session-limit-min-guarantee" dval:"0"`
}

type SystemResourceAccountingTemplateSystemResourcesL4cpsLimitCfg1748 struct {
	L4cpsLimitMax int `json:"l4cps-limit-max"`
}

type SystemResourceAccountingTemplateSystemResourcesL7cpsLimitCfg1749 struct {
	L7cpsLimitMax int `json:"l7cps-limit-max"`
}

type SystemResourceAccountingTemplateSystemResourcesNatcpsLimitCfg1750 struct {
	NatcpsLimitMax int `json:"natcps-limit-max"`
}

type SystemResourceAccountingTemplateSystemResourcesFwcpsLimitCfg1751 struct {
	FwcpsLimitMax int `json:"fwcps-limit-max"`
}

type SystemResourceAccountingTemplateSystemResourcesSslThroughputLimitCfg1752 struct {
	SslThroughputLimitMax              int `json:"ssl-throughput-limit-max"`
	SslThroughputLimitWatermarkDisable int `json:"ssl-throughput-limit-watermark-disable"`
}

type SystemResourceAccountingTemplateSystemResourcesSslcpsLimitCfg1753 struct {
	SslcpsLimitMax int `json:"sslcps-limit-max"`
}

func (p *SystemResourceAccountingTemplate) GetId() string {
	return p.Inst.Name
}

func (p *SystemResourceAccountingTemplate) getPath() string {
	return "system/resource-accounting/template"
}

func (p *SystemResourceAccountingTemplate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemResourceAccountingTemplate::Post")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	_, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
	return err
}

func (p *SystemResourceAccountingTemplate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemResourceAccountingTemplate::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}
func (p *SystemResourceAccountingTemplate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemResourceAccountingTemplate::Put")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload: " + string(payloadBytes))
	_, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
	return err
}

func (p *SystemResourceAccountingTemplate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemResourceAccountingTemplate::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
