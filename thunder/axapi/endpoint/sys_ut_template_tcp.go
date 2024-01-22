

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SysUtTemplateTcp struct {
	Inst struct {

    AckSeqNumber string `json:"ack-seq-number" dval:"valid"`
    Checksum string `json:"checksum" dval:"valid"`
    DestPort int `json:"dest-port"`
    DestPortValue int `json:"dest-port-value"`
    Flags SysUtTemplateTcpFlags1553 `json:"flags"`
    NatPool string `json:"nat-pool"`
    Options SysUtTemplateTcpOptions1554 `json:"options"`
    SeqNumber string `json:"seq-number" dval:"valid"`
    SrcPortRange []SysUtTemplateTcpSrcPortRange `json:"src-port-range"`
    Urgent string `json:"urgent" dval:"valid"`
    Uuid string `json:"uuid"`
    Window string `json:"window" dval:"valid"`
    Name string 

	} `json:"tcp"`
}


type SysUtTemplateTcpFlags1553 struct {
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


type SysUtTemplateTcpOptions1554 struct {
    Mss int `json:"mss"`
    Wscale int `json:"wscale"`
    SackType string `json:"sack-type"`
    TimeStampEnable int `json:"time-stamp-enable"`
    Nop int `json:"nop"`
    Uuid string `json:"uuid"`
}


type SysUtTemplateTcpSrcPortRange struct {
    SrcPortStart int `json:"src-port-start"`
    SrcPortEnd int `json:"src-port-end"`
}

func (p *SysUtTemplateTcp) GetId() string{
    return "1"
}

func (p *SysUtTemplateTcp) getPath() string{
    return "sys-ut/template/" +p.Inst.Name + "/tcp"
}

func (p *SysUtTemplateTcp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplateTcp::Post")
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

func (p *SysUtTemplateTcp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplateTcp::Get")
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
func (p *SysUtTemplateTcp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplateTcp::Put")
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

func (p *SysUtTemplateTcp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplateTcp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
