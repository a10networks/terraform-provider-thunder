

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SysUtEventActionL1 struct {
	Inst struct {

    Auto int `json:"auto"`
    EthList []SysUtEventActionL1EthList `json:"eth-list"`
    Length int `json:"length"`
    Trunk_list []SysUtEventActionL1Trunk_list `json:"trunk_list"`
    Uuid string `json:"uuid"`
    Value int `json:"value"`
    EventNumber string 
    Direction string 

	} `json:"l1"`
}


type SysUtEventActionL1EthList struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type SysUtEventActionL1Trunk_list struct {
    TrunkStart int `json:"trunk-start"`
    TrunkEnd int `json:"trunk-end"`
}

func (p *SysUtEventActionL1) GetId() string{
    return "1"
}

func (p *SysUtEventActionL1) getPath() string{
    return "sys-ut/event/" +p.Inst.EventNumber + "/action/" +p.Inst.Direction + "/l1"
}

func (p *SysUtEventActionL1) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtEventActionL1::Post")
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

func (p *SysUtEventActionL1) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtEventActionL1::Get")
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
func (p *SysUtEventActionL1) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtEventActionL1::Put")
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

func (p *SysUtEventActionL1) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtEventActionL1::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
