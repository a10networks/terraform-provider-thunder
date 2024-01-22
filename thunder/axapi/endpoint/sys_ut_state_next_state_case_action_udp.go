

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SysUtStateNextStateCaseActionUdp struct {
	Inst struct {

    Checksum string `json:"checksum" dval:"valid"`
    DestPort int `json:"dest-port"`
    DestPortValue int `json:"dest-port-value"`
    Length int `json:"length"`
    NatPool string `json:"nat-pool"`
    SrcPort int `json:"src-port"`
    Uuid string `json:"uuid"`
    CaseNumber string 
    Name string 
    Direction string 

	} `json:"udp"`
}

func (p *SysUtStateNextStateCaseActionUdp) GetId() string{
    return "1"
}

func (p *SysUtStateNextStateCaseActionUdp) getPath() string{
    return "sys-ut/state/" +p.Inst.Name + "/next-state/" +p.Inst.Name + "/case/" +p.Inst.CaseNumber + "/action/" +p.Inst.Direction + "/udp"
}

func (p *SysUtStateNextStateCaseActionUdp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtStateNextStateCaseActionUdp::Post")
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

func (p *SysUtStateNextStateCaseActionUdp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtStateNextStateCaseActionUdp::Get")
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
func (p *SysUtStateNextStateCaseActionUdp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtStateNextStateCaseActionUdp::Put")
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

func (p *SysUtStateNextStateCaseActionUdp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtStateNextStateCaseActionUdp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
