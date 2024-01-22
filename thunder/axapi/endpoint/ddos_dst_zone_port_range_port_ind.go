

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortRangePortInd struct {
	Inst struct {

    SamplingEnable []DdosDstZonePortRangePortIndSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`
    Protocol string 
    ZoneName string 
    PortRangeEnd string 
    PortRangeStart string 

	} `json:"port-ind"`
}


type DdosDstZonePortRangePortIndSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *DdosDstZonePortRangePortInd) GetId() string{
    return "1"
}

func (p *DdosDstZonePortRangePortInd) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/port-range/" +p.Inst.PortRangeStart + "+" +p.Inst.PortRangeEnd + "+" +p.Inst.Protocol + "/port-ind"
}

func (p *DdosDstZonePortRangePortInd) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortRangePortInd::Post")
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

func (p *DdosDstZonePortRangePortInd) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortRangePortInd::Get")
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
func (p *DdosDstZonePortRangePortInd) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortRangePortInd::Put")
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

func (p *DdosDstZonePortRangePortInd) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortRangePortInd::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
