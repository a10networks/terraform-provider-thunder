

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryPortPortInd struct {
	Inst struct {

    SamplingEnable []DdosDstEntryPortPortIndSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`
    Protocol string 
    PortNum string 
    DstEntryName string 

	} `json:"port-ind"`
}


type DdosDstEntryPortPortIndSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *DdosDstEntryPortPortInd) GetId() string{
    return "1"
}

func (p *DdosDstEntryPortPortInd) getPath() string{
    return "ddos/dst/entry/" +p.Inst.DstEntryName + "/port/" +p.Inst.PortNum + "+" +p.Inst.Protocol + "/port-ind"
}

func (p *DdosDstEntryPortPortInd) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryPortPortInd::Post")
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

func (p *DdosDstEntryPortPortInd) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryPortPortInd::Get")
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
func (p *DdosDstEntryPortPortInd) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryPortPortInd::Put")
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

func (p *DdosDstEntryPortPortInd) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryPortPortInd::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
