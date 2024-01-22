

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbPolicyGeoLocationMatch struct {
	Inst struct {

    GeoTypeOverlap string `json:"geo-type-overlap"`
    MatchFirst string `json:"match-first" dval:"global"`
    Overlap int `json:"overlap"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"geo-location-match"`
}

func (p *GslbPolicyGeoLocationMatch) GetId() string{
    return "1"
}

func (p *GslbPolicyGeoLocationMatch) getPath() string{
    return "gslb/policy/" +p.Inst.Name + "/geo-location-match"
}

func (p *GslbPolicyGeoLocationMatch) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyGeoLocationMatch::Post")
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

func (p *GslbPolicyGeoLocationMatch) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyGeoLocationMatch::Get")
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
func (p *GslbPolicyGeoLocationMatch) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyGeoLocationMatch::Put")
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

func (p *GslbPolicyGeoLocationMatch) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyGeoLocationMatch::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
