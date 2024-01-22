

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemGeoLocation struct {
	Inst struct {

    EntryList []SystemGeoLocationEntryList `json:"entry-list"`
    GeoLocationGeolite2Asn int `json:"geo-location-geolite2-asn"`
    GeoLocationGeolite2City int `json:"geo-location-geolite2-city"`
    GeoLocationGeolite2Country int `json:"geo-location-geolite2-country"`
    GeoLocationIana int `json:"geo-location-iana" dval:"1"`
    GeoLocationIanaSystem int `json:"geo-location-iana-system"`
    Geolite2AsnIncludeIpv6 int `json:"geolite2-asn-include-ipv6"`
    Geolite2CityIncludeIpv6 int `json:"geolite2-city-include-ipv6"`
    Geolite2CountryIncludeIpv6 int `json:"geolite2-country-include-ipv6"`
    GeolocLoadFileList []SystemGeoLocationGeolocLoadFileList `json:"geoloc-load-file-list"`
    Uuid string `json:"uuid"`

	} `json:"geo-location"`
}


type SystemGeoLocationEntryList struct {
    GeoLocnObjName string `json:"geo-locn-obj-name"`
    GeoLocnMultipleAddresses []SystemGeoLocationEntryListGeoLocnMultipleAddresses `json:"geo-locn-multiple-addresses"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type SystemGeoLocationEntryListGeoLocnMultipleAddresses struct {
    FirstIpAddress string `json:"first-ip-address"`
    GeolIpv4Mask string `json:"geol-ipv4-mask"`
    IpAddr2 string `json:"ip-addr2"`
    FirstIpv6Address string `json:"first-ipv6-address"`
    GeolIpv6Mask int `json:"geol-ipv6-mask"`
    Ipv6Addr2 string `json:"ipv6-addr2"`
}


type SystemGeoLocationGeolocLoadFileList struct {
    GeoLocationLoadFilename string `json:"geo-location-load-filename"`
    GeoLocationLoadFileIncludeIpv6 int `json:"geo-location-load-file-include-ipv6"`
    TemplateName string `json:"template-name"`
}

func (p *SystemGeoLocation) GetId() string{
    return "1"
}

func (p *SystemGeoLocation) getPath() string{
    return "system/geo-location"
}

func (p *SystemGeoLocation) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemGeoLocation::Post")
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

func (p *SystemGeoLocation) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemGeoLocation::Get")
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
func (p *SystemGeoLocation) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemGeoLocation::Put")
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

func (p *SystemGeoLocation) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemGeoLocation::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
