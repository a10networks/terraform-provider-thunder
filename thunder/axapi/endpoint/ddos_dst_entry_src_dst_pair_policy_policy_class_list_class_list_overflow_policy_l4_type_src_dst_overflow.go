

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflow struct {
	Inst struct {

    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Protocol string `json:"protocol"`
    Template DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflowTemplate `json:"template"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ClassListName string 
    DstEntryName string 
    SrcBasedPolicyName string 
    DummyName string 

	} `json:"l4-type-src-dst-overflow"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflowTemplate struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Other string `json:"other"`
    TemplateIcmpV4 string `json:"template-icmp-v4"`
    TemplateIcmpV6 string `json:"template-icmp-v6"`
}

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflow) GetId() string{
    return p.Inst.Protocol
}

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflow) getPath() string{
    return "ddos/dst/entry/" +p.Inst.DstEntryName + "/src-dst-pair-policy/" +p.Inst.SrcBasedPolicyName + "/policy-class-list/" +p.Inst.ClassListName + "/class-list-overflow-policy/" +p.Inst.DummyName + "/l4-type-src-dst-overflow"
}

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflow) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflow::Post")
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

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflow) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflow::Get")
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
func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflow) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflow::Put")
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

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflow) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflow::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
