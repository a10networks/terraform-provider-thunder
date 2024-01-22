

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemResourceAccountingTemplateAppResources struct {
	Inst struct {

    CacheTemplateCfg SystemResourceAccountingTemplateAppResourcesCacheTemplateCfg `json:"cache-template-cfg"`
    ClientSslTemplateCfg SystemResourceAccountingTemplateAppResourcesClientSslTemplateCfg `json:"client-ssl-template-cfg"`
    ConnReuseTemplateCfg SystemResourceAccountingTemplateAppResourcesConnReuseTemplateCfg `json:"conn-reuse-template-cfg"`
    FastTcpTemplateCfg SystemResourceAccountingTemplateAppResourcesFastTcpTemplateCfg `json:"fast-tcp-template-cfg"`
    FastUdpTemplateCfg SystemResourceAccountingTemplateAppResourcesFastUdpTemplateCfg `json:"fast-udp-template-cfg"`
    FixTemplateCfg SystemResourceAccountingTemplateAppResourcesFixTemplateCfg `json:"fix-template-cfg"`
    GslbDeviceCfg SystemResourceAccountingTemplateAppResourcesGslbDeviceCfg `json:"gslb-device-cfg"`
    GslbGeoLocationCfg SystemResourceAccountingTemplateAppResourcesGslbGeoLocationCfg `json:"gslb-geo-location-cfg"`
    GslbIpListCfg SystemResourceAccountingTemplateAppResourcesGslbIpListCfg `json:"gslb-ip-list-cfg"`
    GslbPolicyCfg SystemResourceAccountingTemplateAppResourcesGslbPolicyCfg `json:"gslb-policy-cfg"`
    GslbServiceCfg SystemResourceAccountingTemplateAppResourcesGslbServiceCfg `json:"gslb-service-cfg"`
    GslbServiceIpCfg SystemResourceAccountingTemplateAppResourcesGslbServiceIpCfg `json:"gslb-service-ip-cfg"`
    GslbServicePortCfg SystemResourceAccountingTemplateAppResourcesGslbServicePortCfg `json:"gslb-service-port-cfg"`
    GslbSiteCfg SystemResourceAccountingTemplateAppResourcesGslbSiteCfg `json:"gslb-site-cfg"`
    GslbSvcGroupCfg SystemResourceAccountingTemplateAppResourcesGslbSvcGroupCfg `json:"gslb-svc-group-cfg"`
    GslbTemplateCfg SystemResourceAccountingTemplateAppResourcesGslbTemplateCfg `json:"gslb-template-cfg"`
    GslbZoneCfg SystemResourceAccountingTemplateAppResourcesGslbZoneCfg `json:"gslb-zone-cfg"`
    HealthMonitorCfg SystemResourceAccountingTemplateAppResourcesHealthMonitorCfg `json:"health-monitor-cfg"`
    HttpTemplateCfg SystemResourceAccountingTemplateAppResourcesHttpTemplateCfg `json:"http-template-cfg"`
    LinkCostTemplateCfg SystemResourceAccountingTemplateAppResourcesLinkCostTemplateCfg `json:"link-cost-template-cfg"`
    PbslbEntryCfg SystemResourceAccountingTemplateAppResourcesPbslbEntryCfg `json:"pbslb-entry-cfg"`
    PersistCookieTemplateCfg SystemResourceAccountingTemplateAppResourcesPersistCookieTemplateCfg `json:"persist-cookie-template-cfg"`
    PersistSrcipTemplateCfg SystemResourceAccountingTemplateAppResourcesPersistSrcipTemplateCfg `json:"persist-srcip-template-cfg"`
    ProxyTemplateCfg SystemResourceAccountingTemplateAppResourcesProxyTemplateCfg `json:"proxy-template-cfg"`
    RealPortCfg SystemResourceAccountingTemplateAppResourcesRealPortCfg `json:"real-port-cfg"`
    RealServerCfg SystemResourceAccountingTemplateAppResourcesRealServerCfg `json:"real-server-cfg"`
    ServerSslTemplateCfg SystemResourceAccountingTemplateAppResourcesServerSslTemplateCfg `json:"server-ssl-template-cfg"`
    ServiceGroupCfg SystemResourceAccountingTemplateAppResourcesServiceGroupCfg `json:"service-group-cfg"`
    StreamTemplateCfg SystemResourceAccountingTemplateAppResourcesStreamTemplateCfg `json:"stream-template-cfg"`
    Threshold int `json:"threshold"`
    Uuid string `json:"uuid"`
    VirtualPortCfg SystemResourceAccountingTemplateAppResourcesVirtualPortCfg `json:"virtual-port-cfg"`
    VirtualServerCfg SystemResourceAccountingTemplateAppResourcesVirtualServerCfg `json:"virtual-server-cfg"`
    Name string 

	} `json:"app-resources"`
}


type SystemResourceAccountingTemplateAppResourcesCacheTemplateCfg struct {
    CacheTemplateMax int `json:"cache-template-max"`
    CacheTemplateMinGuarantee int `json:"cache-template-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesClientSslTemplateCfg struct {
    ClientSslTemplateMax int `json:"client-ssl-template-max"`
    ClientSslTemplateMinGuarantee int `json:"client-ssl-template-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesConnReuseTemplateCfg struct {
    ConnReuseTemplateMax int `json:"conn-reuse-template-max"`
    ConnReuseTemplateMinGuarantee int `json:"conn-reuse-template-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesFastTcpTemplateCfg struct {
    FastTcpTemplateMax int `json:"fast-tcp-template-max"`
    FastTcpTemplateMinGuarantee int `json:"fast-tcp-template-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesFastUdpTemplateCfg struct {
    FastUdpTemplateMax int `json:"fast-udp-template-max"`
    FastUdpTemplateMinGuarantee int `json:"fast-udp-template-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesFixTemplateCfg struct {
    FixTemplateMax int `json:"fix-template-max"`
    FixTemplateMinGuarantee int `json:"fix-template-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesGslbDeviceCfg struct {
    GslbDeviceMax int `json:"gslb-device-max"`
    GslbDeviceMinGuarantee int `json:"gslb-device-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesGslbGeoLocationCfg struct {
    GslbGeoLocationMax int `json:"gslb-geo-location-max"`
    GslbGeoLocationMinGuarantee int `json:"gslb-geo-location-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesGslbIpListCfg struct {
    GslbIpListMax int `json:"gslb-ip-list-max"`
    GslbIpListMinGuarantee int `json:"gslb-ip-list-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesGslbPolicyCfg struct {
    GslbPolicyMax int `json:"gslb-policy-max"`
    GslbPolicyMinGuarantee int `json:"gslb-policy-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesGslbServiceCfg struct {
    GslbServiceMax int `json:"gslb-service-max"`
    GslbServiceMinGuarantee int `json:"gslb-service-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesGslbServiceIpCfg struct {
    GslbServiceIpMax int `json:"gslb-service-ip-max"`
    GslbServiceIpMinGuarantee int `json:"gslb-service-ip-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesGslbServicePortCfg struct {
    GslbServicePortMax int `json:"gslb-service-port-max"`
    GslbServicePortMinGuarantee int `json:"gslb-service-port-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesGslbSiteCfg struct {
    GslbSiteMax int `json:"gslb-site-max"`
    GslbSiteMinGuarantee int `json:"gslb-site-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesGslbSvcGroupCfg struct {
    GslbSvcGroupMax int `json:"gslb-svc-group-max"`
    GslbSvcGroupMinGuarantee int `json:"gslb-svc-group-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesGslbTemplateCfg struct {
    GslbTemplateMax int `json:"gslb-template-max"`
    GslbTemplateMinGuarantee int `json:"gslb-template-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesGslbZoneCfg struct {
    GslbZoneMax int `json:"gslb-zone-max"`
    GslbZoneMinGuarantee int `json:"gslb-zone-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesHealthMonitorCfg struct {
    HealthMonitorMax int `json:"health-monitor-max"`
    HealthMonitorMinGuarantee int `json:"health-monitor-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesHttpTemplateCfg struct {
    HttpTemplateMax int `json:"http-template-max"`
    HttpTemplateMinGuarantee int `json:"http-template-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesLinkCostTemplateCfg struct {
    LinkCostTemplateMax int `json:"link-cost-template-max"`
    LinkCostTemplateMinGuarantee int `json:"link-cost-template-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesPbslbEntryCfg struct {
    PbslbEntryMax int `json:"pbslb-entry-max"`
    PbslbEntryMinGuarantee int `json:"pbslb-entry-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesPersistCookieTemplateCfg struct {
    PersistCookieTemplateMax int `json:"persist-cookie-template-max"`
    PersistCookieTemplateMinGuarantee int `json:"persist-cookie-template-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesPersistSrcipTemplateCfg struct {
    PersistSrcipTemplateMax int `json:"persist-srcip-template-max"`
    PersistSrcipTemplateMinGuarantee int `json:"persist-srcip-template-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesProxyTemplateCfg struct {
    ProxyTemplateMax int `json:"proxy-template-max"`
    ProxyTemplateMinGuarantee int `json:"proxy-template-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesRealPortCfg struct {
    RealPortMax int `json:"real-port-max"`
    RealPortMinGuarantee int `json:"real-port-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesRealServerCfg struct {
    RealServerMax int `json:"real-server-max"`
    RealServerMinGuarantee int `json:"real-server-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesServerSslTemplateCfg struct {
    ServerSslTemplateMax int `json:"server-ssl-template-max"`
    ServerSslTemplateMinGuarantee int `json:"server-ssl-template-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesServiceGroupCfg struct {
    ServiceGroupMax int `json:"service-group-max"`
    ServiceGroupMinGuarantee int `json:"service-group-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesStreamTemplateCfg struct {
    StreamTemplateMax int `json:"stream-template-max"`
    StreamTemplateMinGuarantee int `json:"stream-template-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesVirtualPortCfg struct {
    VirtualPortMax int `json:"virtual-port-max"`
    VirtualPortMinGuarantee int `json:"virtual-port-min-guarantee"`
}


type SystemResourceAccountingTemplateAppResourcesVirtualServerCfg struct {
    VirtualServerMax int `json:"virtual-server-max"`
    VirtualServerMinGuarantee int `json:"virtual-server-min-guarantee"`
}

func (p *SystemResourceAccountingTemplateAppResources) GetId() string{
    return "1"
}

func (p *SystemResourceAccountingTemplateAppResources) getPath() string{
    return "system/resource-accounting/template/" +p.Inst.Name + "/app-resources"
}

func (p *SystemResourceAccountingTemplateAppResources) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemResourceAccountingTemplateAppResources::Post")
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

func (p *SystemResourceAccountingTemplateAppResources) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemResourceAccountingTemplateAppResources::Get")
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
func (p *SystemResourceAccountingTemplateAppResources) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemResourceAccountingTemplateAppResources::Put")
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

func (p *SystemResourceAccountingTemplateAppResources) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemResourceAccountingTemplateAppResources::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
