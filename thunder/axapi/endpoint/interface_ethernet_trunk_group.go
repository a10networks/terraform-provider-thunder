

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type InterfaceEthernetTrunkGroup struct {
	Inst struct {

    AdminKey int `json:"admin-key"`
    Mode string `json:"mode" dval:"active"`
    PortPriority int `json:"port-priority"`
    Timeout string `json:"timeout" dval:"long"`
    TrunkNumber int `json:"trunk-number"`
    Type string `json:"type" dval:"static"`
    UdldTimeoutCfg InterfaceEthernetTrunkGroupUdldTimeoutCfg `json:"udld-timeout-cfg"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Ifnum string 

	} `json:"trunk-group"`
}


type InterfaceEthernetTrunkGroupUdldTimeoutCfg struct {
    Fast int `json:"fast" dval:"1000"`
    Slow int `json:"slow"`
}

func (p *InterfaceEthernetTrunkGroup) GetId() string{
    return strconv.Itoa(p.Inst.TrunkNumber)
}

func (p *InterfaceEthernetTrunkGroup) getPath() string{
    return "interface/ethernet/" +p.Inst.Ifnum + "/trunk-group"
}

func (p *InterfaceEthernetTrunkGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetTrunkGroup::Post")
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

func (p *InterfaceEthernetTrunkGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetTrunkGroup::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *InterfaceEthernetTrunkGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetTrunkGroup::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *InterfaceEthernetTrunkGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetTrunkGroup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
