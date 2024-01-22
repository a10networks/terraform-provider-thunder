

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneDnsMxRecord struct {
	Inst struct {

    MxName string `json:"mx-name"`
    Priority int `json:"priority"`
    SamplingEnable []GslbZoneDnsMxRecordSamplingEnable `json:"sampling-enable"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"dns-mx-record"`
}


type GslbZoneDnsMxRecordSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *GslbZoneDnsMxRecord) GetId() string{
    return p.Inst.MxName
}

func (p *GslbZoneDnsMxRecord) getPath() string{
    return "gslb/zone/" +p.Inst.Name + "/dns-mx-record"
}

func (p *GslbZoneDnsMxRecord) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneDnsMxRecord::Post")
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

func (p *GslbZoneDnsMxRecord) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneDnsMxRecord::Get")
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
func (p *GslbZoneDnsMxRecord) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneDnsMxRecord::Put")
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

func (p *GslbZoneDnsMxRecord) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneDnsMxRecord::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
