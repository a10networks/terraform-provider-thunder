

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemResourceAccounting struct {
	Inst struct {

    TemplateList []SystemResourceAccountingTemplateList `json:"template-list"`
    Uuid string `json:"uuid"`

	} `json:"resource-accounting"`
}


type SystemResourceAccountingTemplateList struct {
    Name string `json:"name"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    AppResources SystemResourceAccountingTemplateListAppResources `json:"app-resources"`
    NetworkResources SystemResourceAccountingTemplateListNetworkResources `json:"network-resources"`
    SystemResources SystemResourceAccountingTemplateListSystemResources `json:"system-resources"`
}


type SystemResourceAccountingTemplateListAppResources struct {
    GslbDeviceCfg SystemResourceAccountingTemplateListAppResourcesGslbDeviceCfg `json:"gslb-device-cfg"`
    GslbGeoLocationCfg SystemResourceAccountingTemplateListAppResourcesGslbGeoLocationCfg `json:"gslb-geo-location-cfg"`
    GslbIpListCfg SystemResourceAccountingTemplateListAppResourcesGslbIpListCfg `json:"gslb-ip-list-cfg"`
    GslbPolicyCfg SystemResourceAccountingTemplateListAppResourcesGslbPolicyCfg `json:"gslb-policy-cfg"`
    GslbServiceCfg SystemResourceAccountingTemplateListAppResourcesGslbServiceCfg `json:"gslb-service-cfg"`
    GslbServiceIpCfg SystemResourceAccountingTemplateListAppResourcesGslbServiceIpCfg `json:"gslb-service-ip-cfg"`
    GslbServicePortCfg SystemResourceAccountingTemplateListAppResourcesGslbServicePortCfg `json:"gslb-service-port-cfg"`
    GslbSiteCfg SystemResourceAccountingTemplateListAppResourcesGslbSiteCfg `json:"gslb-site-cfg"`
    GslbSvcGroupCfg SystemResourceAccountingTemplateListAppResourcesGslbSvcGroupCfg `json:"gslb-svc-group-cfg"`
    GslbTemplateCfg SystemResourceAccountingTemplateListAppResourcesGslbTemplateCfg `json:"gslb-template-cfg"`
    GslbZoneCfg SystemResourceAccountingTemplateListAppResourcesGslbZoneCfg `json:"gslb-zone-cfg"`
    HealthMonitorCfg SystemResourceAccountingTemplateListAppResourcesHealthMonitorCfg `json:"health-monitor-cfg"`
    RealPortCfg SystemResourceAccountingTemplateListAppResourcesRealPortCfg `json:"real-port-cfg"`
    RealServerCfg SystemResourceAccountingTemplateListAppResourcesRealServerCfg `json:"real-server-cfg"`
    ServiceGroupCfg SystemResourceAccountingTemplateListAppResourcesServiceGroupCfg `json:"service-group-cfg"`
    VirtualServerCfg SystemResourceAccountingTemplateListAppResourcesVirtualServerCfg `json:"virtual-server-cfg"`
    VirtualPortCfg SystemResourceAccountingTemplateListAppResourcesVirtualPortCfg `json:"virtual-port-cfg"`
    CacheTemplateCfg SystemResourceAccountingTemplateListAppResourcesCacheTemplateCfg `json:"cache-template-cfg"`
    ClientSslTemplateCfg SystemResourceAccountingTemplateListAppResourcesClientSslTemplateCfg `json:"client-ssl-template-cfg"`
    ConnReuseTemplateCfg SystemResourceAccountingTemplateListAppResourcesConnReuseTemplateCfg `json:"conn-reuse-template-cfg"`
    FastTcpTemplateCfg SystemResourceAccountingTemplateListAppResourcesFastTcpTemplateCfg `json:"fast-tcp-template-cfg"`
    FastUdpTemplateCfg SystemResourceAccountingTemplateListAppResourcesFastUdpTemplateCfg `json:"fast-udp-template-cfg"`
    FixTemplateCfg SystemResourceAccountingTemplateListAppResourcesFixTemplateCfg `json:"fix-template-cfg"`
    HttpTemplateCfg SystemResourceAccountingTemplateListAppResourcesHttpTemplateCfg `json:"http-template-cfg"`
    LinkCostTemplateCfg SystemResourceAccountingTemplateListAppResourcesLinkCostTemplateCfg `json:"link-cost-template-cfg"`
    PbslbEntryCfg SystemResourceAccountingTemplateListAppResourcesPbslbEntryCfg `json:"pbslb-entry-cfg"`
    PersistCookieTemplateCfg SystemResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfg `json:"persist-cookie-template-cfg"`
    PersistSrcipTemplateCfg SystemResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfg `json:"persist-srcip-template-cfg"`
    ServerSslTemplateCfg SystemResourceAccountingTemplateListAppResourcesServerSslTemplateCfg `json:"server-ssl-template-cfg"`
    ProxyTemplateCfg SystemResourceAccountingTemplateListAppResourcesProxyTemplateCfg `json:"proxy-template-cfg"`
    StreamTemplateCfg SystemResourceAccountingTemplateListAppResourcesStreamTemplateCfg `json:"stream-template-cfg"`
    Threshold int `json:"threshold"`
    Uuid string `json:"uuid"`
}


type SystemResourceAccountingTemplateListAppResourcesGslbDeviceCfg struct {
    GslbDeviceMax int `json:"gslb-device-max"`
    GslbDeviceMinGuarantee int `json:"gslb-device-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesGslbGeoLocationCfg struct {
    GslbGeoLocationMax int `json:"gslb-geo-location-max"`
    GslbGeoLocationMinGuarantee int `json:"gslb-geo-location-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesGslbIpListCfg struct {
    GslbIpListMax int `json:"gslb-ip-list-max"`
    GslbIpListMinGuarantee int `json:"gslb-ip-list-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesGslbPolicyCfg struct {
    GslbPolicyMax int `json:"gslb-policy-max"`
    GslbPolicyMinGuarantee int `json:"gslb-policy-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesGslbServiceCfg struct {
    GslbServiceMax int `json:"gslb-service-max"`
    GslbServiceMinGuarantee int `json:"gslb-service-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesGslbServiceIpCfg struct {
    GslbServiceIpMax int `json:"gslb-service-ip-max"`
    GslbServiceIpMinGuarantee int `json:"gslb-service-ip-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesGslbServicePortCfg struct {
    GslbServicePortMax int `json:"gslb-service-port-max"`
    GslbServicePortMinGuarantee int `json:"gslb-service-port-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesGslbSiteCfg struct {
    GslbSiteMax int `json:"gslb-site-max"`
    GslbSiteMinGuarantee int `json:"gslb-site-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesGslbSvcGroupCfg struct {
    GslbSvcGroupMax int `json:"gslb-svc-group-max"`
    GslbSvcGroupMinGuarantee int `json:"gslb-svc-group-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesGslbTemplateCfg struct {
    GslbTemplateMax int `json:"gslb-template-max"`
    GslbTemplateMinGuarantee int `json:"gslb-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesGslbZoneCfg struct {
    GslbZoneMax int `json:"gslb-zone-max"`
    GslbZoneMinGuarantee int `json:"gslb-zone-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesHealthMonitorCfg struct {
    HealthMonitorMax int `json:"health-monitor-max"`
    HealthMonitorMinGuarantee int `json:"health-monitor-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesRealPortCfg struct {
    RealPortMax int `json:"real-port-max"`
    RealPortMinGuarantee int `json:"real-port-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesRealServerCfg struct {
    RealServerMax int `json:"real-server-max"`
    RealServerMinGuarantee int `json:"real-server-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesServiceGroupCfg struct {
    ServiceGroupMax int `json:"service-group-max"`
    ServiceGroupMinGuarantee int `json:"service-group-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesVirtualServerCfg struct {
    VirtualServerMax int `json:"virtual-server-max"`
    VirtualServerMinGuarantee int `json:"virtual-server-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesVirtualPortCfg struct {
    VirtualPortMax int `json:"virtual-port-max"`
    VirtualPortMinGuarantee int `json:"virtual-port-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesCacheTemplateCfg struct {
    CacheTemplateMax int `json:"cache-template-max"`
    CacheTemplateMinGuarantee int `json:"cache-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesClientSslTemplateCfg struct {
    ClientSslTemplateMax int `json:"client-ssl-template-max"`
    ClientSslTemplateMinGuarantee int `json:"client-ssl-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesConnReuseTemplateCfg struct {
    ConnReuseTemplateMax int `json:"conn-reuse-template-max"`
    ConnReuseTemplateMinGuarantee int `json:"conn-reuse-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesFastTcpTemplateCfg struct {
    FastTcpTemplateMax int `json:"fast-tcp-template-max"`
    FastTcpTemplateMinGuarantee int `json:"fast-tcp-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesFastUdpTemplateCfg struct {
    FastUdpTemplateMax int `json:"fast-udp-template-max"`
    FastUdpTemplateMinGuarantee int `json:"fast-udp-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesFixTemplateCfg struct {
    FixTemplateMax int `json:"fix-template-max"`
    FixTemplateMinGuarantee int `json:"fix-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesHttpTemplateCfg struct {
    HttpTemplateMax int `json:"http-template-max"`
    HttpTemplateMinGuarantee int `json:"http-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesLinkCostTemplateCfg struct {
    LinkCostTemplateMax int `json:"link-cost-template-max"`
    LinkCostTemplateMinGuarantee int `json:"link-cost-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesPbslbEntryCfg struct {
    PbslbEntryMax int `json:"pbslb-entry-max"`
    PbslbEntryMinGuarantee int `json:"pbslb-entry-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfg struct {
    PersistCookieTemplateMax int `json:"persist-cookie-template-max"`
    PersistCookieTemplateMinGuarantee int `json:"persist-cookie-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfg struct {
    PersistSrcipTemplateMax int `json:"persist-srcip-template-max"`
    PersistSrcipTemplateMinGuarantee int `json:"persist-srcip-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesServerSslTemplateCfg struct {
    ServerSslTemplateMax int `json:"server-ssl-template-max"`
    ServerSslTemplateMinGuarantee int `json:"server-ssl-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesProxyTemplateCfg struct {
    ProxyTemplateMax int `json:"proxy-template-max"`
    ProxyTemplateMinGuarantee int `json:"proxy-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesStreamTemplateCfg struct {
    StreamTemplateMax int `json:"stream-template-max"`
    StreamTemplateMinGuarantee int `json:"stream-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListNetworkResources struct {
    StaticIpv4RouteCfg SystemResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfg `json:"static-ipv4-route-cfg"`
    StaticIpv6RouteCfg SystemResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfg `json:"static-ipv6-route-cfg"`
    Ipv4AclLineCfg SystemResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfg `json:"ipv4-acl-line-cfg"`
    Ipv6AclLineCfg SystemResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfg `json:"ipv6-acl-line-cfg"`
    StaticArpCfg SystemResourceAccountingTemplateListNetworkResourcesStaticArpCfg `json:"static-arp-cfg"`
    StaticNeighborCfg SystemResourceAccountingTemplateListNetworkResourcesStaticNeighborCfg `json:"static-neighbor-cfg"`
    StaticMacCfg SystemResourceAccountingTemplateListNetworkResourcesStaticMacCfg `json:"static-mac-cfg"`
    ObjectGroupCfg SystemResourceAccountingTemplateListNetworkResourcesObjectGroupCfg `json:"object-group-cfg"`
    ObjectGroupClauseCfg SystemResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfg `json:"object-group-clause-cfg"`
    Threshold int `json:"threshold"`
    Uuid string `json:"uuid"`
}


type SystemResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfg struct {
    StaticIpv4RouteMax int `json:"static-ipv4-route-max"`
    StaticIpv4RouteMinGuarantee int `json:"static-ipv4-route-min-guarantee"`
}


type SystemResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfg struct {
    StaticIpv6RouteMax int `json:"static-ipv6-route-max"`
    StaticIpv6RouteMinGuarantee int `json:"static-ipv6-route-min-guarantee"`
}


type SystemResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfg struct {
    Ipv4AclLineMax int `json:"ipv4-acl-line-max"`
    Ipv4AclLineMinGuarantee int `json:"ipv4-acl-line-min-guarantee"`
}


type SystemResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfg struct {
    Ipv6AclLineMax int `json:"ipv6-acl-line-max"`
    Ipv6AclLineMinGuarantee int `json:"ipv6-acl-line-min-guarantee"`
}


type SystemResourceAccountingTemplateListNetworkResourcesStaticArpCfg struct {
    StaticArpMax int `json:"static-arp-max"`
    StaticArpMinGuarantee int `json:"static-arp-min-guarantee"`
}


type SystemResourceAccountingTemplateListNetworkResourcesStaticNeighborCfg struct {
    StaticNeighborMax int `json:"static-neighbor-max"`
    StaticNeighborMinGuarantee int `json:"static-neighbor-min-guarantee"`
}


type SystemResourceAccountingTemplateListNetworkResourcesStaticMacCfg struct {
    StaticMacMax int `json:"static-mac-max"`
    StaticMacMinGuarantee int `json:"static-mac-min-guarantee"`
}


type SystemResourceAccountingTemplateListNetworkResourcesObjectGroupCfg struct {
    ObjectGroupMax int `json:"object-group-max"`
    ObjectGroupMinGuarantee int `json:"object-group-min-guarantee"`
}


type SystemResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfg struct {
    ObjectGroupClauseMax int `json:"object-group-clause-max"`
    ObjectGroupClauseMinGuarantee int `json:"object-group-clause-min-guarantee"`
}


type SystemResourceAccountingTemplateListSystemResources struct {
    BwLimitCfg SystemResourceAccountingTemplateListSystemResourcesBwLimitCfg `json:"bw-limit-cfg"`
    ConcurrentSessionLimitCfg SystemResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfg `json:"concurrent-session-limit-cfg"`
    L4SessionLimitCfg SystemResourceAccountingTemplateListSystemResourcesL4SessionLimitCfg `json:"l4-session-limit-cfg"`
    L4cpsLimitCfg SystemResourceAccountingTemplateListSystemResourcesL4cpsLimitCfg `json:"l4cps-limit-cfg"`
    L7cpsLimitCfg SystemResourceAccountingTemplateListSystemResourcesL7cpsLimitCfg `json:"l7cps-limit-cfg"`
    NatcpsLimitCfg SystemResourceAccountingTemplateListSystemResourcesNatcpsLimitCfg `json:"natcps-limit-cfg"`
    FwcpsLimitCfg SystemResourceAccountingTemplateListSystemResourcesFwcpsLimitCfg `json:"fwcps-limit-cfg"`
    SslThroughputLimitCfg SystemResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfg `json:"ssl-throughput-limit-cfg"`
    SslcpsLimitCfg SystemResourceAccountingTemplateListSystemResourcesSslcpsLimitCfg `json:"sslcps-limit-cfg"`
    Threshold int `json:"threshold"`
    Uuid string `json:"uuid"`
}


type SystemResourceAccountingTemplateListSystemResourcesBwLimitCfg struct {
    BwLimitMax int `json:"bw-limit-max"`
    BwLimitWatermarkDisable int `json:"bw-limit-watermark-disable"`
}


type SystemResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfg struct {
    ConcurrentSessionLimitMax int `json:"concurrent-session-limit-max"`
}


type SystemResourceAccountingTemplateListSystemResourcesL4SessionLimitCfg struct {
    L4SessionLimitMax string `json:"l4-session-limit-max"`
    L4SessionLimitMinGuarantee string `json:"l4-session-limit-min-guarantee" dval:"0"`
}


type SystemResourceAccountingTemplateListSystemResourcesL4cpsLimitCfg struct {
    L4cpsLimitMax int `json:"l4cps-limit-max"`
}


type SystemResourceAccountingTemplateListSystemResourcesL7cpsLimitCfg struct {
    L7cpsLimitMax int `json:"l7cps-limit-max"`
}


type SystemResourceAccountingTemplateListSystemResourcesNatcpsLimitCfg struct {
    NatcpsLimitMax int `json:"natcps-limit-max"`
}


type SystemResourceAccountingTemplateListSystemResourcesFwcpsLimitCfg struct {
    FwcpsLimitMax int `json:"fwcps-limit-max"`
}


type SystemResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfg struct {
    SslThroughputLimitMax int `json:"ssl-throughput-limit-max"`
    SslThroughputLimitWatermarkDisable int `json:"ssl-throughput-limit-watermark-disable"`
}


type SystemResourceAccountingTemplateListSystemResourcesSslcpsLimitCfg struct {
    SslcpsLimitMax int `json:"sslcps-limit-max"`
}

func (p *SystemResourceAccounting) GetId() string{
    return "1"
}

func (p *SystemResourceAccounting) getPath() string{
    return "system/resource-accounting"
}

func (p *SystemResourceAccounting) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemResourceAccounting::Post")
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

func (p *SystemResourceAccounting) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemResourceAccounting::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *SystemResourceAccounting) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemResourceAccounting::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *SystemResourceAccounting) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemResourceAccounting::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
