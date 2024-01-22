

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosDstInterfaceIpv6IpProto struct {
	Inst struct {

    Deny int `json:"deny"`
    Glid string `json:"glid"`
    PortNum int `json:"port-num"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Addr string 

	} `json:"ip-proto"`
}

func (p *DdosDstInterfaceIpv6IpProto) GetId() string{
    return strconv.Itoa(p.Inst.PortNum)
}

func (p *DdosDstInterfaceIpv6IpProto) getPath() string{
    return "ddos/dst/interface-ipv6/" +p.Inst.Addr + "/ip-proto"
}

func (p *DdosDstInterfaceIpv6IpProto) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstInterfaceIpv6IpProto::Post")
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

func (p *DdosDstInterfaceIpv6IpProto) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstInterfaceIpv6IpProto::Get")
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
func (p *DdosDstInterfaceIpv6IpProto) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstInterfaceIpv6IpProto::Put")
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

func (p *DdosDstInterfaceIpv6IpProto) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstInterfaceIpv6IpProto::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
