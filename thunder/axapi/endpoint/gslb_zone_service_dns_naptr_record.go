

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneServiceDnsNaptrRecord struct {
	Inst struct {

    Flag string `json:"flag"`
    NaptrTarget string `json:"naptr-target"`
    Order int `json:"order"`
    Preference int `json:"preference"`
    Regexp int `json:"regexp"`
    SamplingEnable []GslbZoneServiceDnsNaptrRecordSamplingEnable `json:"sampling-enable"`
    ServiceProto string `json:"service-proto"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    ServicePort string 
    ServiceName string 
    Name string 

	} `json:"dns-naptr-record"`
}


type GslbZoneServiceDnsNaptrRecordSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *GslbZoneServiceDnsNaptrRecord) GetId() string{
    return url.QueryEscape(p.Inst.NaptrTarget)+"+"+url.QueryEscape(p.Inst.ServiceProto)+"+"+p.Inst.Flag
}

func (p *GslbZoneServiceDnsNaptrRecord) getPath() string{
    return "gslb/zone/" +p.Inst.Name + "/service/" +p.Inst.ServicePort + "+" +p.Inst.ServiceName + "/dns-naptr-record"
}

func (p *GslbZoneServiceDnsNaptrRecord) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsNaptrRecord::Post")
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

func (p *GslbZoneServiceDnsNaptrRecord) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsNaptrRecord::Get")
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
func (p *GslbZoneServiceDnsNaptrRecord) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsNaptrRecord::Put")
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

func (p *GslbZoneServiceDnsNaptrRecord) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsNaptrRecord::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
