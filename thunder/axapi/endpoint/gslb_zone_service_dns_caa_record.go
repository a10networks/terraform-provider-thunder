

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneServiceDnsCaaRecord struct {
	Inst struct {

    CriticalFlag int `json:"critical-flag"`
    PropertyTag string `json:"property-tag"`
    Rdata string `json:"rdata"`
    SamplingEnable []GslbZoneServiceDnsCaaRecordSamplingEnable `json:"sampling-enable"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    ServicePort string 
    ServiceName string 
    Name string 

	} `json:"dns-caa-record"`
}


type GslbZoneServiceDnsCaaRecordSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *GslbZoneServiceDnsCaaRecord) GetId() string{
    return strconv.Itoa(p.Inst.CriticalFlag)+"+"+url.QueryEscape(p.Inst.PropertyTag)+"+"+url.QueryEscape(p.Inst.Rdata)
}

func (p *GslbZoneServiceDnsCaaRecord) getPath() string{
    return "gslb/zone/" +p.Inst.Name + "/service/" +p.Inst.ServicePort + "+" +p.Inst.ServiceName + "/dns-caa-record"
}

func (p *GslbZoneServiceDnsCaaRecord) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsCaaRecord::Post")
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

func (p *GslbZoneServiceDnsCaaRecord) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsCaaRecord::Get")
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
func (p *GslbZoneServiceDnsCaaRecord) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsCaaRecord::Put")
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

func (p *GslbZoneServiceDnsCaaRecord) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsCaaRecord::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
