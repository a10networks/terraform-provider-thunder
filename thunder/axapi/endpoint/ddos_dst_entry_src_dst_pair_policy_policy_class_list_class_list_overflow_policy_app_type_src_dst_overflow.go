

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflow struct {
	Inst struct {

    Protocol string `json:"protocol"`
    Template DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflowTemplate `json:"template"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ClassListName string 
    DstEntryName string 
    SrcBasedPolicyName string 
    DummyName string 

	} `json:"app-type-src-dst-overflow"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflowTemplate struct {
    SslL4 string `json:"ssl-l4"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    Sip string `json:"sip"`
}

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflow) GetId() string{
    return p.Inst.Protocol
}

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflow) getPath() string{
    return "ddos/dst/entry/" +p.Inst.DstEntryName + "/src-dst-pair-policy/" +p.Inst.SrcBasedPolicyName + "/policy-class-list/" +p.Inst.ClassListName + "/class-list-overflow-policy/" +p.Inst.DummyName + "/app-type-src-dst-overflow"
}

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflow) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflow::Post")
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

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflow) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflow::Get")
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
func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflow) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflow::Put")
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

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflow) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflow::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
