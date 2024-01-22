

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryIpProto struct {
	Inst struct {

    Deny int `json:"deny"`
    EspInspect DdosDstEntryIpProtoEspInspect `json:"esp-inspect"`
    Glid string `json:"glid"`
    GlidExceedAction DdosDstEntryIpProtoGlidExceedAction `json:"glid-exceed-action"`
    IpFilteringPolicy string `json:"ip-filtering-policy"`
    IpFilteringPolicyOper DdosDstEntryIpProtoIpFilteringPolicyOper157 `json:"ip-filtering-policy-oper"`
    PortNum int `json:"port-num"`
    SetCounterBaseVal int `json:"set-counter-base-val"`
    Template DdosDstEntryIpProtoTemplate `json:"template"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    DstEntryName string 

	} `json:"ip-proto"`
}


type DdosDstEntryIpProtoEspInspect struct {
    AuthAlgorithm string `json:"auth-algorithm"`
    EncryptAlgorithm string `json:"encrypt-algorithm"`
    Mode string `json:"mode"`
}


type DdosDstEntryIpProtoGlidExceedAction struct {
    StatelessEncapActionCfg DdosDstEntryIpProtoGlidExceedActionStatelessEncapActionCfg `json:"stateless-encap-action-cfg"`
}


type DdosDstEntryIpProtoGlidExceedActionStatelessEncapActionCfg struct {
    StatelessEncapAction string `json:"stateless-encap-action"`
    EncapTemplate string `json:"encap-template"`
}


type DdosDstEntryIpProtoIpFilteringPolicyOper157 struct {
    Uuid string `json:"uuid"`
}


type DdosDstEntryIpProtoTemplate struct {
    Other string `json:"other"`
}

func (p *DdosDstEntryIpProto) GetId() string{
    return strconv.Itoa(p.Inst.PortNum)
}

func (p *DdosDstEntryIpProto) getPath() string{
    return "ddos/dst/entry/" +p.Inst.DstEntryName + "/ip-proto"
}

func (p *DdosDstEntryIpProto) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryIpProto::Post")
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

func (p *DdosDstEntryIpProto) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryIpProto::Get")
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
func (p *DdosDstEntryIpProto) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryIpProto::Put")
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

func (p *DdosDstEntryIpProto) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryIpProto::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
