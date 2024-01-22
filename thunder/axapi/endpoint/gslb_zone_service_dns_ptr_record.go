

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneServiceDnsPtrRecord struct {
	Inst struct {

    PtrName string `json:"ptr-name"`
    SamplingEnable []GslbZoneServiceDnsPtrRecordSamplingEnable `json:"sampling-enable"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    ServiceName string 
    ServicePort string 
    Name string 

	} `json:"dns-ptr-record"`
}


type GslbZoneServiceDnsPtrRecordSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *GslbZoneServiceDnsPtrRecord) GetId() string{
    return p.Inst.PtrName
}

func (p *GslbZoneServiceDnsPtrRecord) getPath() string{
    return "gslb/zone/" +p.Inst.Name + "/service/" +p.Inst.ServicePort + "+" +p.Inst.ServiceName + "/dns-ptr-record"
}

func (p *GslbZoneServiceDnsPtrRecord) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsPtrRecord::Post")
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

func (p *GslbZoneServiceDnsPtrRecord) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsPtrRecord::Get")
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
func (p *GslbZoneServiceDnsPtrRecord) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsPtrRecord::Put")
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

func (p *GslbZoneServiceDnsPtrRecord) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsPtrRecord::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
