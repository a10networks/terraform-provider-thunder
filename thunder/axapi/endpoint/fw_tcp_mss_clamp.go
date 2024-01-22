

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwTcpMssClamp struct {
	Inst struct {

    Min int `json:"min" dval:"456"`
    MssClampType string `json:"mss-clamp-type"`
    MssSubtract int `json:"mss-subtract"`
    MssValue int `json:"mss-value"`
    Uuid string `json:"uuid"`

	} `json:"mss-clamp"`
}

func (p *FwTcpMssClamp) GetId() string{
    return "1"
}

func (p *FwTcpMssClamp) getPath() string{
    return "fw/tcp/mss-clamp"
}

func (p *FwTcpMssClamp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwTcpMssClamp::Post")
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

func (p *FwTcpMssClamp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwTcpMssClamp::Get")
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
func (p *FwTcpMssClamp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwTcpMssClamp::Put")
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

func (p *FwTcpMssClamp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwTcpMssClamp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
