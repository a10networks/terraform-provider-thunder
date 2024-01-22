

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SysUtTemplateL3 struct {
	Inst struct {

    Checksum string `json:"checksum" dval:"valid"`
    IpList []SysUtTemplateL3IpList `json:"ip-list"`
    Protocol int `json:"protocol"`
    Ttl int `json:"ttl"`
    Type string `json:"type"`
    Uuid string `json:"uuid"`
    Value int `json:"value"`
    Name string 

	} `json:"l3"`
}


type SysUtTemplateL3IpList struct {
    SrcDst string `json:"src-dst"`
    Ipv4StartAddress string `json:"ipv4-start-address"`
    Ipv4EndAddress string `json:"ipv4-end-address"`
    Ipv6StartAddress string `json:"ipv6-start-address"`
    Ipv6EndAddress string `json:"ipv6-end-address"`
    VirtualServer string `json:"virtual-server"`
    NatPool string `json:"nat-pool"`
    Ethernet int `json:"ethernet"`
    Ve int `json:"ve"`
    Trunk int `json:"trunk"`
    Uuid string `json:"uuid"`
}

func (p *SysUtTemplateL3) GetId() string{
    return "1"
}

func (p *SysUtTemplateL3) getPath() string{
    return "sys-ut/template/" +p.Inst.Name + "/l3"
}

func (p *SysUtTemplateL3) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplateL3::Post")
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

func (p *SysUtTemplateL3) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplateL3::Get")
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
func (p *SysUtTemplateL3) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplateL3::Put")
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

func (p *SysUtTemplateL3) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplateL3::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
