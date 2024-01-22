

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryL4TypePortInd struct {
	Inst struct {

    SamplingEnable []DdosDstEntryL4TypePortIndSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`
    Protocol string 
    DstEntryName string 

	} `json:"port-ind"`
}


type DdosDstEntryL4TypePortIndSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *DdosDstEntryL4TypePortInd) GetId() string{
    return "1"
}

func (p *DdosDstEntryL4TypePortInd) getPath() string{
    return "ddos/dst/entry/" +p.Inst.DstEntryName + "/l4-type/" +p.Inst.Protocol + "/port-ind"
}

func (p *DdosDstEntryL4TypePortInd) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryL4TypePortInd::Post")
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

func (p *DdosDstEntryL4TypePortInd) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryL4TypePortInd::Get")
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
func (p *DdosDstEntryL4TypePortInd) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryL4TypePortInd::Put")
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

func (p *DdosDstEntryL4TypePortInd) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryL4TypePortInd::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
