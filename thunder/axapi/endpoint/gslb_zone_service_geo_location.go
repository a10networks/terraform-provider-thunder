

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneServiceGeoLocation struct {
	Inst struct {

    Action int `json:"action"`
    ActionType string `json:"action-type"`
    Alias []GslbZoneServiceGeoLocationAlias `json:"alias"`
    ForwardType string `json:"forward-type"`
    GeoName string `json:"geo-name"`
    Policy string `json:"policy"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ServiceName string 
    ServicePort string 
    Name string 

	} `json:"geo-location"`
}


type GslbZoneServiceGeoLocationAlias struct {
    Alias string `json:"alias"`
}

func (p *GslbZoneServiceGeoLocation) GetId() string{
    return p.Inst.GeoName
}

func (p *GslbZoneServiceGeoLocation) getPath() string{
    return "gslb/zone/" +p.Inst.Name + "/service/" +p.Inst.ServicePort + "+" +p.Inst.ServiceName + "/geo-location"
}

func (p *GslbZoneServiceGeoLocation) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceGeoLocation::Post")
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

func (p *GslbZoneServiceGeoLocation) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceGeoLocation::Get")
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
func (p *GslbZoneServiceGeoLocation) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceGeoLocation::Put")
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

func (p *GslbZoneServiceGeoLocation) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceGeoLocation::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
