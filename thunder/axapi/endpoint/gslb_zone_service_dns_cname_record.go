

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneServiceDnsCnameRecord struct {
	Inst struct {

    AdminPreference int `json:"admin-preference" dval:"100"`
    AliasName string `json:"alias-name"`
    AsBackup int `json:"as-backup"`
    SamplingEnable []GslbZoneServiceDnsCnameRecordSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`
    Weight int `json:"weight" dval:"1"`
    ServiceName string 
    ServicePort string 
    Name string 

	} `json:"dns-cname-record"`
}


type GslbZoneServiceDnsCnameRecordSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *GslbZoneServiceDnsCnameRecord) GetId() string{
    return p.Inst.AliasName
}

func (p *GslbZoneServiceDnsCnameRecord) getPath() string{
    return "gslb/zone/" +p.Inst.Name + "/service/" +p.Inst.ServicePort + "+" +p.Inst.ServiceName + "/dns-cname-record"
}

func (p *GslbZoneServiceDnsCnameRecord) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsCnameRecord::Post")
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

func (p *GslbZoneServiceDnsCnameRecord) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsCnameRecord::Get")
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
func (p *GslbZoneServiceDnsCnameRecord) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsCnameRecord::Put")
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

func (p *GslbZoneServiceDnsCnameRecord) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsCnameRecord::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
