

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneServiceDnsSrvRecord struct {
	Inst struct {

    Port int `json:"port"`
    Priority int `json:"priority"`
    SamplingEnable []GslbZoneServiceDnsSrvRecordSamplingEnable `json:"sampling-enable"`
    SrvName string `json:"srv-name"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    Weight int `json:"weight" dval:"10"`
    ServicePort string 
    ServiceName string 
    Name string 

	} `json:"dns-srv-record"`
}


type GslbZoneServiceDnsSrvRecordSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *GslbZoneServiceDnsSrvRecord) GetId() string{
    return p.Inst.SrvName+"+"+strconv.Itoa(p.Inst.Port)
}

func (p *GslbZoneServiceDnsSrvRecord) getPath() string{
    return "gslb/zone/" +p.Inst.Name + "/service/" +p.Inst.ServicePort + "+" +p.Inst.ServiceName + "/dns-srv-record"
}

func (p *GslbZoneServiceDnsSrvRecord) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsSrvRecord::Post")
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

func (p *GslbZoneServiceDnsSrvRecord) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsSrvRecord::Get")
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
func (p *GslbZoneServiceDnsSrvRecord) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsSrvRecord::Put")
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

func (p *GslbZoneServiceDnsSrvRecord) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsSrvRecord::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
