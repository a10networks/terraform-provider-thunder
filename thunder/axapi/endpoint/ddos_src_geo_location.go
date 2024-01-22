

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosSrcGeoLocation struct {
	Inst struct {

    AppTypeList []DdosSrcGeoLocationAppTypeList `json:"app-type-list"`
    Bypass int `json:"bypass"`
    Description string `json:"description"`
    GeolocationName string `json:"geolocation-name"`
    Glid string `json:"glid"`
    L4TypeList []DdosSrcGeoLocationL4TypeList `json:"l4-type-list"`
    LogPeriodic int `json:"log-periodic"`
    Template DdosSrcGeoLocationTemplate `json:"template"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"geo-location"`
}


type DdosSrcGeoLocationAppTypeList struct {
    Protocol string `json:"protocol"`
    Template DdosSrcGeoLocationAppTypeListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosSrcGeoLocationAppTypeListTemplate struct {
    SslL4 string `json:"ssl-l4"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    Sip string `json:"sip"`
}


type DdosSrcGeoLocationL4TypeList struct {
    Protocol string `json:"protocol"`
    Action string `json:"action"`
    Glid string `json:"glid"`
    Template DdosSrcGeoLocationL4TypeListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosSrcGeoLocationL4TypeListTemplate struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Other string `json:"other"`
    TemplateIcmpV4 string `json:"template-icmp-v4"`
}


type DdosSrcGeoLocationTemplate struct {
    Logging string `json:"logging"`
}

func (p *DdosSrcGeoLocation) GetId() string{
    return url.QueryEscape(p.Inst.GeolocationName)
}

func (p *DdosSrcGeoLocation) getPath() string{
    return "ddos/src/geo-location"
}

func (p *DdosSrcGeoLocation) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcGeoLocation::Post")
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

func (p *DdosSrcGeoLocation) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcGeoLocation::Get")
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
func (p *DdosSrcGeoLocation) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcGeoLocation::Put")
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

func (p *DdosSrcGeoLocation) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcGeoLocation::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
