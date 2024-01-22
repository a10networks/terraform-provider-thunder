

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SysUtEventActionTcp struct {
	Inst struct {

    AckSeqNumber string `json:"ack-seq-number" dval:"valid"`
    Checksum string `json:"checksum" dval:"valid"`
    DestPort int `json:"dest-port"`
    DestPortValue int `json:"dest-port-value"`
    Flags SysUtEventActionTcpFlags1526 `json:"flags"`
    NatPool string `json:"nat-pool"`
    Options SysUtEventActionTcpOptions1527 `json:"options"`
    SeqNumber string `json:"seq-number" dval:"valid"`
    SrcPort int `json:"src-port"`
    Urgent string `json:"urgent" dval:"valid"`
    Uuid string `json:"uuid"`
    Window string `json:"window" dval:"valid"`
    EventNumber string 
    Direction string 

	} `json:"tcp"`
}


type SysUtEventActionTcpFlags1526 struct {
    Syn int `json:"syn"`
    Ack int `json:"ack"`
    Fin int `json:"fin"`
    Rst int `json:"rst"`
    Psh int `json:"psh"`
    Ece int `json:"ece"`
    Urg int `json:"urg"`
    Cwr int `json:"cwr"`
    Uuid string `json:"uuid"`
}


type SysUtEventActionTcpOptions1527 struct {
    Mss int `json:"mss"`
    Wscale int `json:"wscale"`
    SackType string `json:"sack-type"`
    TimeStampEnable int `json:"time-stamp-enable"`
    Nop int `json:"nop"`
    Uuid string `json:"uuid"`
}

func (p *SysUtEventActionTcp) GetId() string{
    return "1"
}

func (p *SysUtEventActionTcp) getPath() string{
    return "sys-ut/event/" +p.Inst.EventNumber + "/action/" +p.Inst.Direction + "/tcp"
}

func (p *SysUtEventActionTcp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtEventActionTcp::Post")
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

func (p *SysUtEventActionTcp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtEventActionTcp::Get")
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
func (p *SysUtEventActionTcp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtEventActionTcp::Put")
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

func (p *SysUtEventActionTcp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtEventActionTcp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
