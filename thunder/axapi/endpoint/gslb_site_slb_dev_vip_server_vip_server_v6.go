

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbSiteSlbDevVipServerVipServerV6 struct {
	Inst struct {

    Ipv6 string `json:"ipv6"`
    SamplingEnable []GslbSiteSlbDevVipServerVipServerV6SamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`
    SiteName string 
    DeviceName string 

	} `json:"vip-server-v6"`
}


type GslbSiteSlbDevVipServerVipServerV6SamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *GslbSiteSlbDevVipServerVipServerV6) GetId() string{
    return p.Inst.Ipv6
}

func (p *GslbSiteSlbDevVipServerVipServerV6) getPath() string{
    return "gslb/site/" +p.Inst.SiteName + "/slb-dev/" +p.Inst.DeviceName + "/vip-server/vip-server-v6"
}

func (p *GslbSiteSlbDevVipServerVipServerV6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSiteSlbDevVipServerVipServerV6::Post")
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

func (p *GslbSiteSlbDevVipServerVipServerV6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSiteSlbDevVipServerVipServerV6::Get")
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
func (p *GslbSiteSlbDevVipServerVipServerV6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSiteSlbDevVipServerVipServerV6::Put")
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

func (p *GslbSiteSlbDevVipServerVipServerV6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSiteSlbDevVipServerVipServerV6::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
