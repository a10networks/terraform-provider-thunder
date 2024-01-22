

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbSiteIpServer struct {
	Inst struct {

    IpServerName string `json:"ip-server-name"`
    Uuid string `json:"uuid"`
    SiteName string 

	} `json:"ip-server"`
}

func (p *GslbSiteIpServer) GetId() string{
    return p.Inst.IpServerName
}

func (p *GslbSiteIpServer) getPath() string{
    return "gslb/site/" +p.Inst.SiteName + "/ip-server"
}

func (p *GslbSiteIpServer) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSiteIpServer::Post")
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

func (p *GslbSiteIpServer) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSiteIpServer::Get")
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
func (p *GslbSiteIpServer) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSiteIpServer::Put")
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

func (p *GslbSiteIpServer) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSiteIpServer::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
