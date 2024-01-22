

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbSiteSlbDevVipServerVipServerName struct {
	Inst struct {

    SamplingEnable []GslbSiteSlbDevVipServerVipServerNameSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`
    VipName string `json:"vip-name"`
    SiteName string 
    DeviceName string 

	} `json:"vip-server-name"`
}


type GslbSiteSlbDevVipServerVipServerNameSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *GslbSiteSlbDevVipServerVipServerName) GetId() string{
    return p.Inst.VipName
}

func (p *GslbSiteSlbDevVipServerVipServerName) getPath() string{
    return "gslb/site/" +p.Inst.SiteName + "/slb-dev/" +p.Inst.DeviceName + "/vip-server/vip-server-name"
}

func (p *GslbSiteSlbDevVipServerVipServerName) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSiteSlbDevVipServerVipServerName::Post")
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

func (p *GslbSiteSlbDevVipServerVipServerName) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSiteSlbDevVipServerVipServerName::Get")
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
func (p *GslbSiteSlbDevVipServerVipServerName) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSiteSlbDevVipServerVipServerName::Put")
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

func (p *GslbSiteSlbDevVipServerVipServerName) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSiteSlbDevVipServerVipServerName::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
