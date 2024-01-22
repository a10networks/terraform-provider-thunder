

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpTcp struct {
	Inst struct {

    SynCookie IpTcpSynCookie `json:"syn-cookie"`
    Uuid string `json:"uuid"`

	} `json:"tcp"`
}


type IpTcpSynCookie struct {
    Threshold int `json:"threshold" dval:"4"`
    SackEnable int `json:"sack-enable"`
}

func (p *IpTcp) GetId() string{
    return "1"
}

func (p *IpTcp) getPath() string{
    return "ip/tcp"
}

func (p *IpTcp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpTcp::Post")
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

func (p *IpTcp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpTcp::Get")
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
func (p *IpTcp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpTcp::Put")
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

func (p *IpTcp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpTcp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
