

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SysUtTemplateL1 struct {
	Inst struct {

    Auto int `json:"auto"`
    Drop int `json:"drop"`
    EthList []SysUtTemplateL1EthList `json:"eth-list"`
    Length int `json:"length"`
    Trunk_list []SysUtTemplateL1Trunk_list `json:"trunk_list"`
    Uuid string `json:"uuid"`
    Value int `json:"value"`
    Name string 

	} `json:"l1"`
}


type SysUtTemplateL1EthList struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type SysUtTemplateL1Trunk_list struct {
    TrunkStart int `json:"trunk-start"`
    TrunkEnd int `json:"trunk-end"`
}

func (p *SysUtTemplateL1) GetId() string{
    return "1"
}

func (p *SysUtTemplateL1) getPath() string{
    return "sys-ut/template/" +p.Inst.Name + "/l1"
}

func (p *SysUtTemplateL1) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplateL1::Post")
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

func (p *SysUtTemplateL1) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplateL1::Get")
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
func (p *SysUtTemplateL1) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplateL1::Put")
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

func (p *SysUtTemplateL1) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplateL1::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
