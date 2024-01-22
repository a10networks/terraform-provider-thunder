

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ZoneVlan struct {
	Inst struct {

    Uuid string `json:"uuid"`
    VlanList []ZoneVlanVlanList `json:"vlan-list"`
    Name string 

	} `json:"vlan"`
}


type ZoneVlanVlanList struct {
    VlanStart int `json:"vlan-start"`
    VlanEnd int `json:"vlan-end"`
}

func (p *ZoneVlan) GetId() string{
    return "1"
}

func (p *ZoneVlan) getPath() string{
    return "zone/" +p.Inst.Name + "/vlan"
}

func (p *ZoneVlan) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ZoneVlan::Post")
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

func (p *ZoneVlan) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ZoneVlan::Get")
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
func (p *ZoneVlan) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ZoneVlan::Put")
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

func (p *ZoneVlan) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ZoneVlan::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
