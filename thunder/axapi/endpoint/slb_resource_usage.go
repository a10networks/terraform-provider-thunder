

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbResourceUsage struct {
	Inst struct {

    CacheTemplateCount int `json:"cache-template-count"`
    ClientSslTemplateCount int `json:"client-ssl-template-count"`
    ConnReuseTemplateCount int `json:"conn-reuse-template-count"`
    FastTcpTemplateCount int `json:"fast-tcp-template-count"`
    FastUdpTemplateCount int `json:"fast-udp-template-count"`
    FixTemplateCount int `json:"fix-template-count"`
    GslbDeviceCount int `json:"gslb-device-count"`
    GslbGeoLocationCount int `json:"gslb-geo-location-count"`
    GslbIpListCount int `json:"gslb-ip-list-count"`
    GslbPolicyCount int `json:"gslb-policy-count"`
    GslbServiceCount int `json:"gslb-service-count"`
    GslbServiceIpCount int `json:"gslb-service-ip-count"`
    GslbServicePortCount int `json:"gslb-service-port-count"`
    GslbSiteCount int `json:"gslb-site-count"`
    GslbSvcGroupCount int `json:"gslb-svc-group-count"`
    GslbTemplateCount int `json:"gslb-template-count"`
    GslbZoneCount int `json:"gslb-zone-count"`
    HealthMonitorCount int `json:"health-monitor-count"`
    HttpTemplateCount int `json:"http-template-count"`
    LinkCostTemplateCount int `json:"link-cost-template-count"`
    NatPoolAddrCount int `json:"nat-pool-addr-count"`
    PbslbEntryCount int `json:"pbslb-entry-count"`
    PbslbSubnetCount int `json:"pbslb-subnet-count"`
    PersistCookieTemplateCount int `json:"persist-cookie-template-count"`
    PersistSrcipTemplateCount int `json:"persist-srcip-template-count"`
    ProxyTemplateCount int `json:"proxy-template-count"`
    RealPortCount int `json:"real-port-count"`
    RealServerCount int `json:"real-server-count"`
    ServerSslTemplateCount int `json:"server-ssl-template-count"`
    ServiceGroupCount int `json:"service-group-count"`
    SlbThresholdResUsagePercent int `json:"slb-threshold-res-usage-percent"`
    StreamTemplateCount int `json:"stream-template-count"`
    Uuid string `json:"uuid"`
    VirtualPortCount int `json:"virtual-port-count"`
    VirtualServerCount int `json:"virtual-server-count"`

	} `json:"resource-usage"`
}

func (p *SlbResourceUsage) GetId() string{
    return "1"
}

func (p *SlbResourceUsage) getPath() string{
    return "slb/resource-usage"
}

func (p *SlbResourceUsage) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbResourceUsage::Post")
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

func (p *SlbResourceUsage) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbResourceUsage::Get")
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
func (p *SlbResourceUsage) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbResourceUsage::Put")
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

func (p *SlbResourceUsage) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbResourceUsage::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
