

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceLifBfd struct {
	Inst struct {

    Authentication InterfaceLifBfdAuthentication `json:"authentication"`
    Demand int `json:"demand"`
    Echo int `json:"echo"`
    IntervalCfg InterfaceLifBfdIntervalCfg `json:"interval-cfg"`
    Uuid string `json:"uuid"`
    Ifname string 

	} `json:"bfd"`
}


type InterfaceLifBfdAuthentication struct {
    KeyId int `json:"key-id"`
    Method string `json:"method"`
    Password string `json:"password"`
    Encrypted string `json:"encrypted"`
}


type InterfaceLifBfdIntervalCfg struct {
    Interval int `json:"interval"`
    MinRx int `json:"min-rx"`
    Multiplier int `json:"multiplier"`
}

func (p *InterfaceLifBfd) GetId() string{
    return "1"
}

func (p *InterfaceLifBfd) getPath() string{
    return "interface/lif/" +p.Inst.Ifname + "/bfd"
}

func (p *InterfaceLifBfd) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifBfd::Post")
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

func (p *InterfaceLifBfd) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifBfd::Get")
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
func (p *InterfaceLifBfd) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifBfd::Put")
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

func (p *InterfaceLifBfd) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifBfd::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
