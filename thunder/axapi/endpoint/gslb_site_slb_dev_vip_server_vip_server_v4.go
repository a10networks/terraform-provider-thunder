

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbSiteSlbDevVipServerVipServerV4 struct {
	Inst struct {

    Ipv4 string `json:"ipv4"`
    SamplingEnable []GslbSiteSlbDevVipServerVipServerV4SamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`
    SiteName string 
    DeviceName string 

	} `json:"vip-server-v4"`
}


type GslbSiteSlbDevVipServerVipServerV4SamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *GslbSiteSlbDevVipServerVipServerV4) GetId() string{
    return p.Inst.Ipv4
}

func (p *GslbSiteSlbDevVipServerVipServerV4) getPath() string{
    return "gslb/site/" +p.Inst.SiteName + "/slb-dev/" +p.Inst.DeviceName + "/vip-server/vip-server-v4"
}

func (p *GslbSiteSlbDevVipServerVipServerV4) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSiteSlbDevVipServerVipServerV4::Post")
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

func (p *GslbSiteSlbDevVipServerVipServerV4) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSiteSlbDevVipServerVipServerV4::Get")
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
func (p *GslbSiteSlbDevVipServerVipServerV4) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSiteSlbDevVipServerVipServerV4::Put")
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

func (p *GslbSiteSlbDevVipServerVipServerV4) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSiteSlbDevVipServerVipServerV4::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
