

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemGeoLocationEntry struct {
	Inst struct {

    GeoLocnMultipleAddresses []SystemGeoLocationEntryGeoLocnMultipleAddresses `json:"geo-locn-multiple-addresses"`
    GeoLocnObjName string `json:"geo-locn-obj-name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"entry"`
}


type SystemGeoLocationEntryGeoLocnMultipleAddresses struct {
    FirstIpAddress string `json:"first-ip-address"`
    GeolIpv4Mask string `json:"geol-ipv4-mask"`
    IpAddr2 string `json:"ip-addr2"`
    FirstIpv6Address string `json:"first-ipv6-address"`
    GeolIpv6Mask int `json:"geol-ipv6-mask"`
    Ipv6Addr2 string `json:"ipv6-addr2"`
}

func (p *SystemGeoLocationEntry) GetId() string{
    return p.Inst.GeoLocnObjName
}

func (p *SystemGeoLocationEntry) getPath() string{
    return "system/geo-location/entry"
}

func (p *SystemGeoLocationEntry) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemGeoLocationEntry::Post")
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

func (p *SystemGeoLocationEntry) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemGeoLocationEntry::Get")
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
func (p *SystemGeoLocationEntry) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemGeoLocationEntry::Put")
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

func (p *SystemGeoLocationEntry) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemGeoLocationEntry::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
