

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbResourceUsageOper struct {
    
    Oper SlbResourceUsageOperOper `json:"oper"`

}
type DataSlbResourceUsageOper struct {
    DtSlbResourceUsageOper SlbResourceUsageOper `json:"resource-usage"`
}


type SlbResourceUsageOperOper struct {
    NatPoolAddrMin int `json:"nat-pool-addr-min"`
    NatPoolAddrMax int `json:"nat-pool-addr-max"`
    NatPoolAddrDefault int `json:"nat-pool-addr-default"`
    RealServerMin int `json:"real-server-min"`
    RealServerMax int `json:"real-server-max"`
    RealServerDefault int `json:"real-server-default"`
    RealPortMin int `json:"real-port-min"`
    RealPortMax int `json:"real-port-max"`
    RealPortDefault int `json:"real-port-default"`
    ServiceGroupMin int `json:"service-group-min"`
    ServiceGroupMax int `json:"service-group-max"`
    ServiceGroupDefault int `json:"service-group-default"`
    VirtualPortMin int `json:"virtual-port-min"`
    VirtualPortMax int `json:"virtual-port-max"`
    VirtualPortDefault int `json:"virtual-port-default"`
    VirtualServerMin int `json:"virtual-server-min"`
    VirtualServerMax int `json:"virtual-server-max"`
    VirtualServerDefault int `json:"virtual-server-default"`
    HttpTemplateMin int `json:"http-template-min"`
    HttpTemplateMax int `json:"http-template-max"`
    HttpTemplateDefault int `json:"http-template-default"`
    FixTemplateMin int `json:"fix-template-min"`
    FixTemplateMax int `json:"fix-template-max"`
    FixTemplateDefault int `json:"fix-template-default"`
    ProxyTemplateMin int `json:"proxy-template-min"`
    ProxyTemplateMax int `json:"proxy-template-max"`
    ProxyTemplateDefault int `json:"proxy-template-default"`
    ConnReuseTemplateMin int `json:"conn-reuse-template-min"`
    ConnReuseTemplateMax int `json:"conn-reuse-template-max"`
    ConnReuseTemplateDefault int `json:"conn-reuse-template-default"`
    FastTcpTemplateMin int `json:"fast-tcp-template-min"`
    FastTcpTemplateMax int `json:"fast-tcp-template-max"`
    FastTcpTemplateDefault int `json:"fast-tcp-template-default"`
    FastUdpTemplateMin int `json:"fast-udp-template-min"`
    FastUdpTemplateMax int `json:"fast-udp-template-max"`
    FastUdpTemplateDefault int `json:"fast-udp-template-default"`
    ClientSslTemplateMin int `json:"client-ssl-template-min"`
    ClientSslTemplateMax int `json:"client-ssl-template-max"`
    ClientSslTemplateDefault int `json:"client-ssl-template-default"`
    ServerSslTemplateMin int `json:"server-ssl-template-min"`
    ServerSslTemplateMax int `json:"server-ssl-template-max"`
    ServerSslTemplateDefault int `json:"server-ssl-template-default"`
    LinkCostTemplateMin int `json:"link-cost-template-min"`
    LinkCostTemplateMax int `json:"link-cost-template-max"`
    LinkCostTemplateDefault int `json:"link-cost-template-default"`
    PbslbEntryMin int `json:"pbslb-entry-min"`
    PbslbEntryMax int `json:"pbslb-entry-max"`
    PbslbEntryDefault int `json:"pbslb-entry-default"`
    StreamTemplateMin int `json:"stream-template-min"`
    StreamTemplateMax int `json:"stream-template-max"`
    StreamTemplateDefault int `json:"stream-template-default"`
    PersistCookieTemplateMin int `json:"persist-cookie-template-min"`
    PersistCookieTemplateMax int `json:"persist-cookie-template-max"`
    PersistCookieTemplateDefault int `json:"persist-cookie-template-default"`
    PersistSrcipTemplateMin int `json:"persist-srcip-template-min"`
    PersistSrcipTemplateMax int `json:"persist-srcip-template-max"`
    PersistSrcipTemplateDefault int `json:"persist-srcip-template-default"`
    HealthMonitorCountMin int `json:"health-monitor-count-min"`
    HealthMonitorCountMax int `json:"health-monitor-count-max"`
    HealthMonitorCountDefault int `json:"health-monitor-count-default"`
    PbslbSubnetCountMin int `json:"pbslb-subnet-count-min"`
    PbslbSubnetCountMax int `json:"pbslb-subnet-count-max"`
    PbslbSubnetCountDefault int `json:"pbslb-subnet-count-default"`
    GslbSiteCountMin int `json:"gslb-site-count-min"`
    GslbSiteCountMax int `json:"gslb-site-count-max"`
    GslbSiteCountDefault int `json:"gslb-site-count-default"`
    GslbDeviceCountMin int `json:"gslb-device-count-min"`
    GslbDeviceCountMax int `json:"gslb-device-count-max"`
    GslbDeviceCountDefault int `json:"gslb-device-count-default"`
    GslbServiceIpCountMin int `json:"gslb-service-ip-count-min"`
    GslbServiceIpCountMax int `json:"gslb-service-ip-count-max"`
    GslbServiceIpCountDefault int `json:"gslb-service-ip-count-default"`
    GslbServicePortCountMin int `json:"gslb-service-port-count-min"`
    GslbServicePortCountMax int `json:"gslb-service-port-count-max"`
    GslbServicePortCountDefault int `json:"gslb-service-port-count-default"`
    GslbZoneCountMin int `json:"gslb-zone-count-min"`
    GslbZoneCountMax int `json:"gslb-zone-count-max"`
    GslbZoneCountDefault int `json:"gslb-zone-count-default"`
    GslbServiceCountMin int `json:"gslb-service-count-min"`
    GslbServiceCountMax int `json:"gslb-service-count-max"`
    GslbServiceCountDefault int `json:"gslb-service-count-default"`
    GslbPolicyCountMin int `json:"gslb-policy-count-min"`
    GslbPolicyCountMax int `json:"gslb-policy-count-max"`
    GslbPolicyCountDefault int `json:"gslb-policy-count-default"`
    GslbGeoLocationCountMin int `json:"gslb-geo-location-count-min"`
    GslbGeoLocationCountMax int `json:"gslb-geo-location-count-max"`
    GslbGeoLocationCountDefault int `json:"gslb-geo-location-count-default"`
    GslbIpListCountMin int `json:"gslb-ip-list-count-min"`
    GslbIpListCountMax int `json:"gslb-ip-list-count-max"`
    GslbIpListCountDefault int `json:"gslb-ip-list-count-default"`
    GslbTemplateCountMin int `json:"gslb-template-count-min"`
    GslbTemplateCountMax int `json:"gslb-template-count-max"`
    GslbTemplateCountDefault int `json:"gslb-template-count-default"`
    GslbSvcgroupCountMin int `json:"gslb-svcgroup-count-min"`
    GslbSvcgroupCountMax int `json:"gslb-svcgroup-count-max"`
    GslbSvcgroupCountDefault int `json:"gslb-svcgroup-count-default"`
    CacheTemplateMin int `json:"cache-template-min"`
    CacheTemplateMax int `json:"cache-template-max"`
    CacheTemplateDefault int `json:"cache-template-default"`
    SlbThresholdResUsageDefault int `json:"slb-threshold-res-usage-default"`
    SlbThresholdResUsageMin int `json:"slb-threshold-res-usage-min"`
    SlbThresholdResUsageMax int `json:"slb-threshold-res-usage-max"`
}

func (p *SlbResourceUsageOper) GetId() string{
    return "1"
}

func (p *SlbResourceUsageOper) getPath() string{
    return "slb/resource-usage/oper"
}

func (p *SlbResourceUsageOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbResourceUsageOper,error) {
logger.Println("SlbResourceUsageOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbResourceUsageOper
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
