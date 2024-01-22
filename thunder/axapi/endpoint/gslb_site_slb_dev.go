

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbSiteSlbDev struct {
	Inst struct {

    AdminPreference int `json:"admin-preference" dval:"100"`
    AutoDetect string `json:"auto-detect" dval:"ip-and-port"`
    AutoMap int `json:"auto-map" dval:"1"`
    ClientIp string `json:"client-ip"`
    DevResolveAs string `json:"dev-resolve-as" dval:"resolve-to-ipv4"`
    DeviceName string `json:"device-name"`
    Domain string `json:"domain"`
    GatewayIpAddr string `json:"gateway-ip-addr"`
    HealthCheckAction string `json:"health-check-action" dval:"health-check"`
    IpAddress string `json:"ip-address"`
    Ipv6Address string `json:"ipv6-address"`
    MaxClient int `json:"max-client" dval:"32768"`
    MsgFormatAcos2x int `json:"msg-format-acos-2x"`
    ProbeTimer int `json:"probe-timer"`
    ProtoAgingFast int `json:"proto-aging-fast" dval:"1"`
    ProtoAgingTime int `json:"proto-aging-time" dval:"60"`
    ProtoCompatible int `json:"proto-compatible"`
    RdtType string `json:"rdt-type"`
    RdtValue int `json:"rdt-value"`
    SessionNumber int `json:"session-number"`
    SessionUtilization int `json:"session-utilization"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VipServer GslbSiteSlbDevVipServer396 `json:"vip-server"`
    SiteName string 

	} `json:"slb-dev"`
}


type GslbSiteSlbDevVipServer396 struct {
    VipServerV4List []GslbSiteSlbDevVipServerVipServerV4List `json:"vip-server-v4-list"`
    VipServerV6List []GslbSiteSlbDevVipServerVipServerV6List `json:"vip-server-v6-list"`
    VipServerNameList []GslbSiteSlbDevVipServerVipServerNameList `json:"vip-server-name-list"`
}


type GslbSiteSlbDevVipServerVipServerV4List struct {
    Ipv4 string `json:"ipv4"`
    Uuid string `json:"uuid"`
    SamplingEnable []GslbSiteSlbDevVipServerVipServerV4ListSamplingEnable `json:"sampling-enable"`
}


type GslbSiteSlbDevVipServerVipServerV4ListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbSiteSlbDevVipServerVipServerV6List struct {
    Ipv6 string `json:"ipv6"`
    Uuid string `json:"uuid"`
    SamplingEnable []GslbSiteSlbDevVipServerVipServerV6ListSamplingEnable `json:"sampling-enable"`
}


type GslbSiteSlbDevVipServerVipServerV6ListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbSiteSlbDevVipServerVipServerNameList struct {
    VipName string `json:"vip-name"`
    Uuid string `json:"uuid"`
    SamplingEnable []GslbSiteSlbDevVipServerVipServerNameListSamplingEnable `json:"sampling-enable"`
}


type GslbSiteSlbDevVipServerVipServerNameListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *GslbSiteSlbDev) GetId() string{
    return p.Inst.DeviceName
}

func (p *GslbSiteSlbDev) getPath() string{
    return "gslb/site/" +p.Inst.SiteName + "/slb-dev"
}

func (p *GslbSiteSlbDev) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSiteSlbDev::Post")
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

func (p *GslbSiteSlbDev) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSiteSlbDev::Get")
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
func (p *GslbSiteSlbDev) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSiteSlbDev::Put")
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

func (p *GslbSiteSlbDev) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSiteSlbDev::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
