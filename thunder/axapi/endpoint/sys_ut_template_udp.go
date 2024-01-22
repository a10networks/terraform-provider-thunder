

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SysUtTemplateUdp struct {
	Inst struct {

    Checksum string `json:"checksum" dval:"valid"`
    DestPort int `json:"dest-port"`
    DestPortValue int `json:"dest-port-value"`
    Length int `json:"length"`
    NatPool string `json:"nat-pool"`
    SrcPortRange []SysUtTemplateUdpSrcPortRange `json:"src-port-range"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"udp"`
}


type SysUtTemplateUdpSrcPortRange struct {
    SrcPortStart int `json:"src-port-start"`
    SrcPortEnd int `json:"src-port-end"`
}

func (p *SysUtTemplateUdp) GetId() string{
    return "1"
}

func (p *SysUtTemplateUdp) getPath() string{
    return "sys-ut/template/" +p.Inst.Name + "/udp"
}

func (p *SysUtTemplateUdp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplateUdp::Post")
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

func (p *SysUtTemplateUdp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplateUdp::Get")
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
func (p *SysUtTemplateUdp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplateUdp::Put")
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

func (p *SysUtTemplateUdp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplateUdp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
