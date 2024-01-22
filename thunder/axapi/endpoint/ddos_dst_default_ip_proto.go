

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosDstDefaultIpProto struct {
	Inst struct {

    Deny int `json:"deny"`
    Glid string `json:"glid"`
    PortNum int `json:"port-num"`
    Template DdosDstDefaultIpProtoTemplate `json:"template"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    DefaultAddressType string 

	} `json:"ip-proto"`
}


type DdosDstDefaultIpProtoTemplate struct {
    Other string `json:"other"`
}

func (p *DdosDstDefaultIpProto) GetId() string{
    return strconv.Itoa(p.Inst.PortNum)
}

func (p *DdosDstDefaultIpProto) getPath() string{
    return "ddos/dst/default/" +p.Inst.DefaultAddressType + "/ip-proto"
}

func (p *DdosDstDefaultIpProto) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDefaultIpProto::Post")
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

func (p *DdosDstDefaultIpProto) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDefaultIpProto::Get")
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
func (p *DdosDstDefaultIpProto) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDefaultIpProto::Put")
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

func (p *DdosDstDefaultIpProto) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDefaultIpProto::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
