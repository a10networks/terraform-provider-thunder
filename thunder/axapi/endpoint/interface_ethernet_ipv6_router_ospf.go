

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceEthernetIpv6RouterOspf struct {
	Inst struct {

    AreaList []InterfaceEthernetIpv6RouterOspfAreaList `json:"area-list"`
    Uuid string `json:"uuid"`
    Ifnum string 

	} `json:"ospf"`
}


type InterfaceEthernetIpv6RouterOspfAreaList struct {
    AreaIdNum int `json:"area-id-num"`
    AreaIdAddr string `json:"area-id-addr"`
    Tag string `json:"tag"`
    InstanceId int `json:"instance-id"`
}

func (p *InterfaceEthernetIpv6RouterOspf) GetId() string{
    return "1"
}

func (p *InterfaceEthernetIpv6RouterOspf) getPath() string{
    return "interface/ethernet/" +p.Inst.Ifnum + "/ipv6/router/ospf"
}

func (p *InterfaceEthernetIpv6RouterOspf) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetIpv6RouterOspf::Post")
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

func (p *InterfaceEthernetIpv6RouterOspf) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetIpv6RouterOspf::Get")
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
func (p *InterfaceEthernetIpv6RouterOspf) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetIpv6RouterOspf::Put")
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

func (p *InterfaceEthernetIpv6RouterOspf) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetIpv6RouterOspf::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
