

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SysUtStateNextStateCaseActionL3 struct {
	Inst struct {

    Checksum string `json:"checksum" dval:"valid"`
    IpList []SysUtStateNextStateCaseActionL3IpList `json:"ip-list"`
    Protocol int `json:"protocol"`
    Ttl int `json:"ttl"`
    Type string `json:"type"`
    Uuid string `json:"uuid"`
    Value int `json:"value"`
    CaseNumber string 
    Name string 
    Direction string 

	} `json:"l3"`
}


type SysUtStateNextStateCaseActionL3IpList struct {
    SrcDst string `json:"src-dst"`
    Ipv4Address string `json:"ipv4-address"`
    Ipv6Address string `json:"ipv6-address"`
    VirtualServer string `json:"virtual-server"`
    NatPool string `json:"nat-pool"`
    Ethernet int `json:"ethernet"`
    Ve int `json:"ve"`
    Trunk int `json:"trunk"`
    Uuid string `json:"uuid"`
}

func (p *SysUtStateNextStateCaseActionL3) GetId() string{
    return "1"
}

func (p *SysUtStateNextStateCaseActionL3) getPath() string{
    return "sys-ut/state/" +p.Inst.Name + "/next-state/" +p.Inst.Name + "/case/" +p.Inst.CaseNumber + "/action/" +p.Inst.Direction + "/l3"
}

func (p *SysUtStateNextStateCaseActionL3) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtStateNextStateCaseActionL3::Post")
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

func (p *SysUtStateNextStateCaseActionL3) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtStateNextStateCaseActionL3::Get")
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
func (p *SysUtStateNextStateCaseActionL3) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtStateNextStateCaseActionL3::Put")
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

func (p *SysUtStateNextStateCaseActionL3) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtStateNextStateCaseActionL3::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
