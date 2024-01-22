

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosSrcGeoLocationL4Type struct {
	Inst struct {

    Action string `json:"action"`
    Glid string `json:"glid"`
    Protocol string `json:"protocol"`
    Template DdosSrcGeoLocationL4TypeTemplate `json:"template"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    GeolocationName string 

	} `json:"l4-type"`
}


type DdosSrcGeoLocationL4TypeTemplate struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Other string `json:"other"`
    TemplateIcmpV4 string `json:"template-icmp-v4"`
}

func (p *DdosSrcGeoLocationL4Type) GetId() string{
    return p.Inst.Protocol
}

func (p *DdosSrcGeoLocationL4Type) getPath() string{
    return "ddos/src/geo-location/" +p.Inst.GeolocationName + "/l4-type"
}

func (p *DdosSrcGeoLocationL4Type) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcGeoLocationL4Type::Post")
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

func (p *DdosSrcGeoLocationL4Type) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcGeoLocationL4Type::Get")
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
func (p *DdosSrcGeoLocationL4Type) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcGeoLocationL4Type::Put")
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

func (p *DdosSrcGeoLocationL4Type) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcGeoLocationL4Type::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
