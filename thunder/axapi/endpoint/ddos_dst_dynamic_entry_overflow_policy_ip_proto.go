

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosDstDynamicEntryOverflowPolicyIpProto struct {
	Inst struct {

    Deny int `json:"deny"`
    Glid string `json:"glid"`
    PortNum int `json:"port-num"`
    Template DdosDstDynamicEntryOverflowPolicyIpProtoTemplate `json:"template"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    DefaultAddressType string 

	} `json:"ip-proto"`
}


type DdosDstDynamicEntryOverflowPolicyIpProtoTemplate struct {
    Other string `json:"other"`
}

func (p *DdosDstDynamicEntryOverflowPolicyIpProto) GetId() string{
    return strconv.Itoa(p.Inst.PortNum)
}

func (p *DdosDstDynamicEntryOverflowPolicyIpProto) getPath() string{
    return "ddos/dst/dynamic-entry-overflow-policy/" +p.Inst.DefaultAddressType + "/ip-proto"
}

func (p *DdosDstDynamicEntryOverflowPolicyIpProto) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDynamicEntryOverflowPolicyIpProto::Post")
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

func (p *DdosDstDynamicEntryOverflowPolicyIpProto) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDynamicEntryOverflowPolicyIpProto::Get")
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
func (p *DdosDstDynamicEntryOverflowPolicyIpProto) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDynamicEntryOverflowPolicyIpProto::Put")
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

func (p *DdosDstDynamicEntryOverflowPolicyIpProto) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDynamicEntryOverflowPolicyIpProto::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
