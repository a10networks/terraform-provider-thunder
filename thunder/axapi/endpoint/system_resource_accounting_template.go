package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type SystemResourceAccountingTemplate struct {
	Inst struct {
		AppResources SystemResourceAccountingTemplateAppResources1701 `json:"app-resources"`

		Name string `json:"name"`

		NetworkResources SystemResourceAccountingTemplateNetworkResources1733 `json:"network-resources"`

		SystemResources SystemResourceAccountingTemplateSystemResources1743 `json:"system-resources"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"template"`
}

type SystemResourceAccountingTemplateAppResources1701 struct {
	GslbDeviceCfg            SystemResourceAccountingTemplateAppResourcesGslbDeviceCfg1702            `json:"gslb-device-cfg"`
	GslbGeoLocationCfg       SystemResourceAccountingTemplateAppResourcesGslbGeoLocationCfg1703       `json:"gslb-geo-location-cfg"`
	GslbIpListCfg            SystemResourceAccountingTemplateAppResourcesGslbIpListCfg1704            `json:"gslb-ip-list-cfg"`
	GslbPolicyCfg            SystemResourceAccountingTemplateAppResourcesGslbPolicyCfg1705            `json:"gslb-policy-cfg"`
	GslbServiceCfg           SystemResourceAccountingTemplateAppResourcesGslbServiceCfg1706           `json:"gslb-service-cfg"`
	GslbServiceIpCfg         SystemResourceAccountingTemplateAppResourcesGslbServiceIpCfg1707         `json:"gslb-service-ip-cfg"`
	GslbServicePortCfg       SystemResourceAccountingTemplateAppResourcesGslbServicePortCfg1708       `json:"gslb-service-port-cfg"`
	GslbSiteCfg              SystemResourceAccountingTemplateAppResourcesGslbSiteCfg1709              `json:"gslb-site-cfg"`
	GslbSvcGroupCfg          SystemResourceAccountingTemplateAppResourcesGslbSvcGroupCfg1710          `json:"gslb-svc-group-cfg"`
	GslbTemplateCfg          SystemResourceAccountingTemplateAppResourcesGslbTemplateCfg1711          `json:"gslb-template-cfg"`
	GslbZoneCfg              SystemResourceAccountingTemplateAppResourcesGslbZoneCfg1712              `json:"gslb-zone-cfg"`
	HealthMonitorCfg         SystemResourceAccountingTemplateAppResourcesHealthMonitorCfg1713         `json:"health-monitor-cfg"`
	RealPortCfg              SystemResourceAccountingTemplateAppResourcesRealPortCfg1714              `json:"real-port-cfg"`
	RealServerCfg            SystemResourceAccountingTemplateAppResourcesRealServerCfg1715            `json:"real-server-cfg"`
	ServiceGroupCfg          SystemResourceAccountingTemplateAppResourcesServiceGroupCfg1716          `json:"service-group-cfg"`
	VirtualServerCfg         SystemResourceAccountingTemplateAppResourcesVirtualServerCfg1717         `json:"virtual-server-cfg"`
	VirtualPortCfg           SystemResourceAccountingTemplateAppResourcesVirtualPortCfg1718           `json:"virtual-port-cfg"`
	CacheTemplateCfg         SystemResourceAccountingTemplateAppResourcesCacheTemplateCfg1719         `json:"cache-template-cfg"`
	ClientSslTemplateCfg     SystemResourceAccountingTemplateAppResourcesClientSslTemplateCfg1720     `json:"client-ssl-template-cfg"`
	ConnReuseTemplateCfg     SystemResourceAccountingTemplateAppResourcesConnReuseTemplateCfg1721     `json:"conn-reuse-template-cfg"`
	FastTcpTemplateCfg       SystemResourceAccountingTemplateAppResourcesFastTcpTemplateCfg1722       `json:"fast-tcp-template-cfg"`
	FastUdpTemplateCfg       SystemResourceAccountingTemplateAppResourcesFastUdpTemplateCfg1723       `json:"fast-udp-template-cfg"`
	FixTemplateCfg           SystemResourceAccountingTemplateAppResourcesFixTemplateCfg1724           `json:"fix-template-cfg"`
	HttpTemplateCfg          SystemResourceAccountingTemplateAppResourcesHttpTemplateCfg1725          `json:"http-template-cfg"`
	LinkCostTemplateCfg      SystemResourceAccountingTemplateAppResourcesLinkCostTemplateCfg1726      `json:"link-cost-template-cfg"`
	PbslbEntryCfg            SystemResourceAccountingTemplateAppResourcesPbslbEntryCfg1727            `json:"pbslb-entry-cfg"`
	PersistCookieTemplateCfg SystemResourceAccountingTemplateAppResourcesPersistCookieTemplateCfg1728 `json:"persist-cookie-template-cfg"`
	PersistSrcipTemplateCfg  SystemResourceAccountingTemplateAppResourcesPersistSrcipTemplateCfg1729  `json:"persist-srcip-template-cfg"`
	ServerSslTemplateCfg     SystemResourceAccountingTemplateAppResourcesServerSslTemplateCfg1730     `json:"server-ssl-template-cfg"`
	ProxyTemplateCfg         SystemResourceAccountingTemplateAppResourcesProxyTemplateCfg1731         `json:"proxy-template-cfg"`
	StreamTemplateCfg        SystemResourceAccountingTemplateAppResourcesStreamTemplateCfg1732        `json:"stream-template-cfg"`
	Threshold                int                                                                      `json:"threshold"`
	Uuid                     string                                                                   `json:"uuid"`
}

type SystemResourceAccountingTemplateAppResourcesGslbDeviceCfg1702 struct {
	GslbDeviceMax          int `json:"gslb-device-max"`
	GslbDeviceMinGuarantee int `json:"gslb-device-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesGslbGeoLocationCfg1703 struct {
	GslbGeoLocationMax          int `json:"gslb-geo-location-max"`
	GslbGeoLocationMinGuarantee int `json:"gslb-geo-location-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesGslbIpListCfg1704 struct {
	GslbIpListMax          int `json:"gslb-ip-list-max"`
	GslbIpListMinGuarantee int `json:"gslb-ip-list-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesGslbPolicyCfg1705 struct {
	GslbPolicyMax          int `json:"gslb-policy-max"`
	GslbPolicyMinGuarantee int `json:"gslb-policy-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesGslbServiceCfg1706 struct {
	GslbServiceMax          int `json:"gslb-service-max"`
	GslbServiceMinGuarantee int `json:"gslb-service-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesGslbServiceIpCfg1707 struct {
	GslbServiceIpMax          int `json:"gslb-service-ip-max"`
	GslbServiceIpMinGuarantee int `json:"gslb-service-ip-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesGslbServicePortCfg1708 struct {
	GslbServicePortMax          int `json:"gslb-service-port-max"`
	GslbServicePortMinGuarantee int `json:"gslb-service-port-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesGslbSiteCfg1709 struct {
	GslbSiteMax          int `json:"gslb-site-max"`
	GslbSiteMinGuarantee int `json:"gslb-site-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesGslbSvcGroupCfg1710 struct {
	GslbSvcGroupMax          int `json:"gslb-svc-group-max"`
	GslbSvcGroupMinGuarantee int `json:"gslb-svc-group-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesGslbTemplateCfg1711 struct {
	GslbTemplateMax          int `json:"gslb-template-max"`
	GslbTemplateMinGuarantee int `json:"gslb-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesGslbZoneCfg1712 struct {
	GslbZoneMax          int `json:"gslb-zone-max"`
	GslbZoneMinGuarantee int `json:"gslb-zone-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesHealthMonitorCfg1713 struct {
	HealthMonitorMax          int `json:"health-monitor-max"`
	HealthMonitorMinGuarantee int `json:"health-monitor-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesRealPortCfg1714 struct {
	RealPortMax          int `json:"real-port-max"`
	RealPortMinGuarantee int `json:"real-port-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesRealServerCfg1715 struct {
	RealServerMax          int `json:"real-server-max"`
	RealServerMinGuarantee int `json:"real-server-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesServiceGroupCfg1716 struct {
	ServiceGroupMax          int `json:"service-group-max"`
	ServiceGroupMinGuarantee int `json:"service-group-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesVirtualServerCfg1717 struct {
	VirtualServerMax          int `json:"virtual-server-max"`
	VirtualServerMinGuarantee int `json:"virtual-server-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesVirtualPortCfg1718 struct {
	VirtualPortMax          int `json:"virtual-port-max"`
	VirtualPortMinGuarantee int `json:"virtual-port-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesCacheTemplateCfg1719 struct {
	CacheTemplateMax          int `json:"cache-template-max"`
	CacheTemplateMinGuarantee int `json:"cache-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesClientSslTemplateCfg1720 struct {
	ClientSslTemplateMax          int `json:"client-ssl-template-max"`
	ClientSslTemplateMinGuarantee int `json:"client-ssl-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesConnReuseTemplateCfg1721 struct {
	ConnReuseTemplateMax          int `json:"conn-reuse-template-max"`
	ConnReuseTemplateMinGuarantee int `json:"conn-reuse-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesFastTcpTemplateCfg1722 struct {
	FastTcpTemplateMax          int `json:"fast-tcp-template-max"`
	FastTcpTemplateMinGuarantee int `json:"fast-tcp-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesFastUdpTemplateCfg1723 struct {
	FastUdpTemplateMax          int `json:"fast-udp-template-max"`
	FastUdpTemplateMinGuarantee int `json:"fast-udp-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesFixTemplateCfg1724 struct {
	FixTemplateMax          int `json:"fix-template-max"`
	FixTemplateMinGuarantee int `json:"fix-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesHttpTemplateCfg1725 struct {
	HttpTemplateMax          int `json:"http-template-max"`
	HttpTemplateMinGuarantee int `json:"http-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesLinkCostTemplateCfg1726 struct {
	LinkCostTemplateMax          int `json:"link-cost-template-max"`
	LinkCostTemplateMinGuarantee int `json:"link-cost-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesPbslbEntryCfg1727 struct {
	PbslbEntryMax          int `json:"pbslb-entry-max"`
	PbslbEntryMinGuarantee int `json:"pbslb-entry-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesPersistCookieTemplateCfg1728 struct {
	PersistCookieTemplateMax          int `json:"persist-cookie-template-max"`
	PersistCookieTemplateMinGuarantee int `json:"persist-cookie-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesPersistSrcipTemplateCfg1729 struct {
	PersistSrcipTemplateMax          int `json:"persist-srcip-template-max"`
	PersistSrcipTemplateMinGuarantee int `json:"persist-srcip-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesServerSslTemplateCfg1730 struct {
	ServerSslTemplateMax          int `json:"server-ssl-template-max"`
	ServerSslTemplateMinGuarantee int `json:"server-ssl-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesProxyTemplateCfg1731 struct {
	ProxyTemplateMax          int `json:"proxy-template-max"`
	ProxyTemplateMinGuarantee int `json:"proxy-template-min-guarantee"`
}

type SystemResourceAccountingTemplateAppResourcesStreamTemplateCfg1732 struct {
	StreamTemplateMax          int `json:"stream-template-max"`
	StreamTemplateMinGuarantee int `json:"stream-template-min-guarantee"`
}

type SystemResourceAccountingTemplateNetworkResources1733 struct {
	StaticIpv4RouteCfg   SystemResourceAccountingTemplateNetworkResourcesStaticIpv4RouteCfg1734   `json:"static-ipv4-route-cfg"`
	StaticIpv6RouteCfg   SystemResourceAccountingTemplateNetworkResourcesStaticIpv6RouteCfg1735   `json:"static-ipv6-route-cfg"`
	Ipv4AclLineCfg       SystemResourceAccountingTemplateNetworkResourcesIpv4AclLineCfg1736       `json:"ipv4-acl-line-cfg"`
	Ipv6AclLineCfg       SystemResourceAccountingTemplateNetworkResourcesIpv6AclLineCfg1737       `json:"ipv6-acl-line-cfg"`
	StaticArpCfg         SystemResourceAccountingTemplateNetworkResourcesStaticArpCfg1738         `json:"static-arp-cfg"`
	StaticNeighborCfg    SystemResourceAccountingTemplateNetworkResourcesStaticNeighborCfg1739    `json:"static-neighbor-cfg"`
	StaticMacCfg         SystemResourceAccountingTemplateNetworkResourcesStaticMacCfg1740         `json:"static-mac-cfg"`
	ObjectGroupCfg       SystemResourceAccountingTemplateNetworkResourcesObjectGroupCfg1741       `json:"object-group-cfg"`
	ObjectGroupClauseCfg SystemResourceAccountingTemplateNetworkResourcesObjectGroupClauseCfg1742 `json:"object-group-clause-cfg"`
	Threshold            int                                                                      `json:"threshold"`
	Uuid                 string                                                                   `json:"uuid"`
}

type SystemResourceAccountingTemplateNetworkResourcesStaticIpv4RouteCfg1734 struct {
	StaticIpv4RouteMax          int `json:"static-ipv4-route-max"`
	StaticIpv4RouteMinGuarantee int `json:"static-ipv4-route-min-guarantee"`
}

type SystemResourceAccountingTemplateNetworkResourcesStaticIpv6RouteCfg1735 struct {
	StaticIpv6RouteMax          int `json:"static-ipv6-route-max"`
	StaticIpv6RouteMinGuarantee int `json:"static-ipv6-route-min-guarantee"`
}

type SystemResourceAccountingTemplateNetworkResourcesIpv4AclLineCfg1736 struct {
	Ipv4AclLineMax          int `json:"ipv4-acl-line-max"`
	Ipv4AclLineMinGuarantee int `json:"ipv4-acl-line-min-guarantee"`
}

type SystemResourceAccountingTemplateNetworkResourcesIpv6AclLineCfg1737 struct {
	Ipv6AclLineMax          int `json:"ipv6-acl-line-max"`
	Ipv6AclLineMinGuarantee int `json:"ipv6-acl-line-min-guarantee"`
}

type SystemResourceAccountingTemplateNetworkResourcesStaticArpCfg1738 struct {
	StaticArpMax          int `json:"static-arp-max"`
	StaticArpMinGuarantee int `json:"static-arp-min-guarantee"`
}

type SystemResourceAccountingTemplateNetworkResourcesStaticNeighborCfg1739 struct {
	StaticNeighborMax          int `json:"static-neighbor-max"`
	StaticNeighborMinGuarantee int `json:"static-neighbor-min-guarantee"`
}

type SystemResourceAccountingTemplateNetworkResourcesStaticMacCfg1740 struct {
	StaticMacMax          int `json:"static-mac-max"`
	StaticMacMinGuarantee int `json:"static-mac-min-guarantee"`
}

type SystemResourceAccountingTemplateNetworkResourcesObjectGroupCfg1741 struct {
	ObjectGroupMax          int `json:"object-group-max"`
	ObjectGroupMinGuarantee int `json:"object-group-min-guarantee"`
}

type SystemResourceAccountingTemplateNetworkResourcesObjectGroupClauseCfg1742 struct {
	ObjectGroupClauseMax          int `json:"object-group-clause-max"`
	ObjectGroupClauseMinGuarantee int `json:"object-group-clause-min-guarantee"`
}

type SystemResourceAccountingTemplateSystemResources1743 struct {
	BwLimitCfg                SystemResourceAccountingTemplateSystemResourcesBwLimitCfg1744                `json:"bw-limit-cfg"`
	ConcurrentSessionLimitCfg SystemResourceAccountingTemplateSystemResourcesConcurrentSessionLimitCfg1745 `json:"concurrent-session-limit-cfg"`
	L4SessionLimitCfg         SystemResourceAccountingTemplateSystemResourcesL4SessionLimitCfg1746         `json:"l4-session-limit-cfg"`
	L4cpsLimitCfg             SystemResourceAccountingTemplateSystemResourcesL4cpsLimitCfg1747             `json:"l4cps-limit-cfg"`
	L7cpsLimitCfg             SystemResourceAccountingTemplateSystemResourcesL7cpsLimitCfg1748             `json:"l7cps-limit-cfg"`
	NatcpsLimitCfg            SystemResourceAccountingTemplateSystemResourcesNatcpsLimitCfg1749            `json:"natcps-limit-cfg"`
	FwcpsLimitCfg             SystemResourceAccountingTemplateSystemResourcesFwcpsLimitCfg1750             `json:"fwcps-limit-cfg"`
	SslThroughputLimitCfg     SystemResourceAccountingTemplateSystemResourcesSslThroughputLimitCfg1751     `json:"ssl-throughput-limit-cfg"`
	SslcpsLimitCfg            SystemResourceAccountingTemplateSystemResourcesSslcpsLimitCfg1752            `json:"sslcps-limit-cfg"`
	Threshold                 int                                                                          `json:"threshold"`
	Uuid                      string                                                                       `json:"uuid"`
}

type SystemResourceAccountingTemplateSystemResourcesBwLimitCfg1744 struct {
	BwLimitMax              int `json:"bw-limit-max"`
	BwLimitWatermarkDisable int `json:"bw-limit-watermark-disable"`
}

type SystemResourceAccountingTemplateSystemResourcesConcurrentSessionLimitCfg1745 struct {
	ConcurrentSessionLimitMax int `json:"concurrent-session-limit-max"`
}

type SystemResourceAccountingTemplateSystemResourcesL4SessionLimitCfg1746 struct {
	L4SessionLimitMax          string `json:"l4-session-limit-max"`
	L4SessionLimitMinGuarantee string `json:"l4-session-limit-min-guarantee" dval:"0"`
}

type SystemResourceAccountingTemplateSystemResourcesL4cpsLimitCfg1747 struct {
	L4cpsLimitMax int `json:"l4cps-limit-max"`
}

type SystemResourceAccountingTemplateSystemResourcesL7cpsLimitCfg1748 struct {
	L7cpsLimitMax int `json:"l7cps-limit-max"`
}

type SystemResourceAccountingTemplateSystemResourcesNatcpsLimitCfg1749 struct {
	NatcpsLimitMax int `json:"natcps-limit-max"`
}

type SystemResourceAccountingTemplateSystemResourcesFwcpsLimitCfg1750 struct {
	FwcpsLimitMax int `json:"fwcps-limit-max"`
}

type SystemResourceAccountingTemplateSystemResourcesSslThroughputLimitCfg1751 struct {
	SslThroughputLimitMax              int `json:"ssl-throughput-limit-max"`
	SslThroughputLimitWatermarkDisable int `json:"ssl-throughput-limit-watermark-disable"`
}

type SystemResourceAccountingTemplateSystemResourcesSslcpsLimitCfg1752 struct {
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
