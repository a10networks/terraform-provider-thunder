

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDst struct {
	Inst struct {

    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Protocol string `json:"protocol"`
    Template DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstTemplate `json:"template"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    SrcBasedPolicyName string 
    ClassListName string 
    DstEntryName string 

	} `json:"l4-type-src-dst"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstTemplate struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Other string `json:"other"`
    TemplateIcmpV4 string `json:"template-icmp-v4"`
    TemplateIcmpV6 string `json:"template-icmp-v6"`
}

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDst) GetId() string{
    return p.Inst.Protocol
}

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDst) getPath() string{
    return "ddos/dst/entry/" +p.Inst.DstEntryName + "/src-dst-pair-policy/" +p.Inst.SrcBasedPolicyName + "/policy-class-list/" +p.Inst.ClassListName + "/l4-type-src-dst"
}

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDst) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDst::Post")
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

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDst) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDst::Get")
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
func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDst) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDst::Put")
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

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDst) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDst::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
