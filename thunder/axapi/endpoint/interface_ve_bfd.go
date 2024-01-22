

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceVeBfd struct {
	Inst struct {

    Authentication InterfaceVeBfdAuthentication `json:"authentication"`
    Demand int `json:"demand"`
    Echo int `json:"echo"`
    IntervalCfg InterfaceVeBfdIntervalCfg `json:"interval-cfg"`
    Uuid string `json:"uuid"`
    Ifnum string 

	} `json:"bfd"`
}


type InterfaceVeBfdAuthentication struct {
    KeyId int `json:"key-id"`
    Method string `json:"method"`
    Password string `json:"password"`
    Encrypted string `json:"encrypted"`
}


type InterfaceVeBfdIntervalCfg struct {
    Interval int `json:"interval"`
    MinRx int `json:"min-rx"`
    Multiplier int `json:"multiplier"`
}

func (p *InterfaceVeBfd) GetId() string{
    return "1"
}

func (p *InterfaceVeBfd) getPath() string{
    return "interface/ve/" +p.Inst.Ifnum + "/bfd"
}

func (p *InterfaceVeBfd) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVeBfd::Post")
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

func (p *InterfaceVeBfd) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVeBfd::Get")
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
func (p *InterfaceVeBfd) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVeBfd::Put")
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

func (p *InterfaceVeBfd) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVeBfd::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
