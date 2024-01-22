

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugIsis struct {
	Inst struct {

    All int `json:"all"`
    Bfd int `json:"bfd"`
    Events int `json:"events"`
    Ifsm int `json:"ifsm"`
    Lsp int `json:"lsp"`
    Nfsm int `json:"nfsm"`
    Nsm int `json:"nsm"`
    Pdu int `json:"pdu"`
    Spf int `json:"spf"`
    Uuid string `json:"uuid"`

	} `json:"isis"`
}

func (p *DebugIsis) GetId() string{
    return "1"
}

func (p *DebugIsis) getPath() string{
    return "debug/isis"
}

func (p *DebugIsis) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugIsis::Post")
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

func (p *DebugIsis) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugIsis::Get")
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
func (p *DebugIsis) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugIsis::Put")
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

func (p *DebugIsis) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugIsis::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
