

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryPortRangePatternRecognition struct {
	Inst struct {

    Algorithm string `json:"algorithm"`
    FilterInactiveThreshold int `json:"filter-inactive-threshold"`
    FilterThreshold int `json:"filter-threshold"`
    Mode string `json:"mode"`
    Sensitivity string `json:"sensitivity"`
    Uuid string `json:"uuid"`
    Protocol string 
    PortRangeEnd string 
    PortRangeStart string 
    DstEntryName string 

	} `json:"pattern-recognition"`
}

func (p *DdosDstEntryPortRangePatternRecognition) GetId() string{
    return "1"
}

func (p *DdosDstEntryPortRangePatternRecognition) getPath() string{
    return "ddos/dst/entry/" +p.Inst.DstEntryName + "/port-range/" +p.Inst.PortRangeStart + "+" +p.Inst.PortRangeEnd + "+" +p.Inst.Protocol + "/pattern-recognition"
}

func (p *DdosDstEntryPortRangePatternRecognition) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryPortRangePatternRecognition::Post")
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

func (p *DdosDstEntryPortRangePatternRecognition) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryPortRangePatternRecognition::Get")
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
func (p *DdosDstEntryPortRangePatternRecognition) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryPortRangePatternRecognition::Put")
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

func (p *DdosDstEntryPortRangePatternRecognition) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryPortRangePatternRecognition::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
