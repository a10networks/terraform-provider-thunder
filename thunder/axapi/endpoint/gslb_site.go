

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type GslbSite struct {
	Inst struct {

    ActiveRdt GslbSiteActiveRdt397 `json:"active-rdt"`
    AutoMap int `json:"auto-map" dval:"1"`
    BwCost int `json:"bw-cost"`
    Controller string `json:"controller"`
    Disable int `json:"disable"`
    EasyRdt GslbSiteEasyRdt398 `json:"easy-rdt"`
    IpServerList []GslbSiteIpServerList `json:"ip-server-list"`
    Limit int `json:"limit"`
    MultipleGeoLocations []GslbSiteMultipleGeoLocations `json:"multiple-geo-locations"`
    ProtoAgingFast int `json:"proto-aging-fast"`
    ProtoAgingTime int `json:"proto-aging-time"`
    SiteName string `json:"site-name"`
    SlbDevList []GslbSiteSlbDevList `json:"slb-dev-list"`
    Template string `json:"template"`
    Threshold int `json:"threshold"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Weight int `json:"weight" dval:"1"`

	} `json:"site"`
}


type GslbSiteActiveRdt397 struct {
    AgingTime int `json:"aging-time" dval:"10"`
    SmoothFactor int `json:"smooth-factor" dval:"10"`
    RangeFactor int `json:"range-factor" dval:"25"`
    Limit int `json:"limit" dval:"16383"`
    Mask string `json:"mask" dval:"/32"`
    Ipv6Mask int `json:"ipv6-mask" dval:"128"`
    IgnoreCount int `json:"ignore-count" dval:"5"`
    BindGeoloc int `json:"bind-geoloc"`
    Overlap int `json:"overlap"`
    Uuid string `json:"uuid"`
}


type GslbSiteEasyRdt398 struct {
    AgingTime int `json:"aging-time" dval:"10"`
    SmoothFactor int `json:"smooth-factor" dval:"10"`
    RangeFactor int `json:"range-factor" dval:"25"`
    Limit int `json:"limit" dval:"16383"`
    Mask string `json:"mask" dval:"/32"`
    Ipv6Mask int `json:"ipv6-mask" dval:"128"`
    IgnoreCount int `json:"ignore-count" dval:"5"`
    BindGeoloc int `json:"bind-geoloc"`
    Overlap int `json:"overlap"`
    Uuid string `json:"uuid"`
}


type GslbSiteIpServerList struct {
    IpServerName string `json:"ip-server-name"`
    Uuid string `json:"uuid"`
}


type GslbSiteMultipleGeoLocations struct {
    GeoLocation string `json:"geo-location"`
}


type GslbSiteSlbDevList struct {
    DeviceName string `json:"device-name"`
    IpAddress string `json:"ip-address"`
    Ipv6Address string `json:"ipv6-address"`
    Domain string `json:"domain"`
    DevResolveAs string `json:"dev-resolve-as" dval:"resolve-to-ipv4"`
    AdminPreference int `json:"admin-preference" dval:"100"`
    SessionNumber int `json:"session-number"`
    SessionUtilization int `json:"session-utilization"`
    RdtType string `json:"rdt-type"`
    ClientIp string `json:"client-ip"`
    RdtValue int `json:"rdt-value"`
    ProbeTimer int `json:"probe-timer"`
    AutoDetect string `json:"auto-detect" dval:"ip-and-port"`
    AutoMap int `json:"auto-map" dval:"1"`
    MaxClient int `json:"max-client" dval:"32768"`
    ProtoAgingTime int `json:"proto-aging-time" dval:"60"`
    ProtoAgingFast int `json:"proto-aging-fast" dval:"1"`
    HealthCheckAction string `json:"health-check-action" dval:"health-check"`
    GatewayIpAddr string `json:"gateway-ip-addr"`
    ProtoCompatible int `json:"proto-compatible"`
    MsgFormatAcos2x int `json:"msg-format-acos-2x"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    VipServer GslbSiteSlbDevListVipServer `json:"vip-server"`
}


type GslbSiteSlbDevListVipServer struct {
    VipServerV4List []GslbSiteSlbDevListVipServerVipServerV4List `json:"vip-server-v4-list"`
    VipServerV6List []GslbSiteSlbDevListVipServerVipServerV6List `json:"vip-server-v6-list"`
    VipServerNameList []GslbSiteSlbDevListVipServerVipServerNameList `json:"vip-server-name-list"`
}


type GslbSiteSlbDevListVipServerVipServerV4List struct {
    Ipv4 string `json:"ipv4"`
    Uuid string `json:"uuid"`
    SamplingEnable []GslbSiteSlbDevListVipServerVipServerV4ListSamplingEnable `json:"sampling-enable"`
}


type GslbSiteSlbDevListVipServerVipServerV4ListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbSiteSlbDevListVipServerVipServerV6List struct {
    Ipv6 string `json:"ipv6"`
    Uuid string `json:"uuid"`
    SamplingEnable []GslbSiteSlbDevListVipServerVipServerV6ListSamplingEnable `json:"sampling-enable"`
}


type GslbSiteSlbDevListVipServerVipServerV6ListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbSiteSlbDevListVipServerVipServerNameList struct {
    VipName string `json:"vip-name"`
    Uuid string `json:"uuid"`
    SamplingEnable []GslbSiteSlbDevListVipServerVipServerNameListSamplingEnable `json:"sampling-enable"`
}


type GslbSiteSlbDevListVipServerVipServerNameListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *GslbSite) GetId() string{
    return url.QueryEscape(p.Inst.SiteName)
}

func (p *GslbSite) getPath() string{
    return "gslb/site"
}

func (p *GslbSite) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSite::Post")
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

func (p *GslbSite) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSite::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *GslbSite) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSite::Put")
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

func (p *GslbSite) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSite::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
