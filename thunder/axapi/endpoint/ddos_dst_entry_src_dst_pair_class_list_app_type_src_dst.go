

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntrySrcDstPairClassListAppTypeSrcDst struct {
	Inst struct {

    Protocol string `json:"protocol"`
    Template DdosDstEntrySrcDstPairClassListAppTypeSrcDstTemplate `json:"template"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ClassListName string 
    DstEntryName string 

	} `json:"app-type-src-dst"`
}


type DdosDstEntrySrcDstPairClassListAppTypeSrcDstTemplate struct {
    SslL4 string `json:"ssl-l4"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    Sip string `json:"sip"`
}

func (p *DdosDstEntrySrcDstPairClassListAppTypeSrcDst) GetId() string{
    return p.Inst.Protocol
}

func (p *DdosDstEntrySrcDstPairClassListAppTypeSrcDst) getPath() string{
    return "ddos/dst/entry/" +p.Inst.DstEntryName + "/src-dst-pair-class-list/" +p.Inst.ClassListName + "/app-type-src-dst"
}

func (p *DdosDstEntrySrcDstPairClassListAppTypeSrcDst) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairClassListAppTypeSrcDst::Post")
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

func (p *DdosDstEntrySrcDstPairClassListAppTypeSrcDst) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairClassListAppTypeSrcDst::Get")
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
func (p *DdosDstEntrySrcDstPairClassListAppTypeSrcDst) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairClassListAppTypeSrcDst::Put")
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

func (p *DdosDstEntrySrcDstPairClassListAppTypeSrcDst) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairClassListAppTypeSrcDst::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
