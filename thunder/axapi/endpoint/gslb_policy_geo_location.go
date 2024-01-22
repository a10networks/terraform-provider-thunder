

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbPolicyGeoLocation struct {
	Inst struct {

    IpMultipleFields []GslbPolicyGeoLocationIpMultipleFields `json:"ip-multiple-fields"`
    Ipv6MultipleFields []GslbPolicyGeoLocationIpv6MultipleFields `json:"ipv6-multiple-fields"`
    Name string `json:"name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    PolicyName string `json:"-"`

	} `json:"geo-location"`
}


type GslbPolicyGeoLocationIpMultipleFields struct {
    IpSub string `json:"ip-sub"`
    IpMaskSub string `json:"ip-mask-sub"`
    IpAddr2Sub string `json:"ip-addr2-sub"`
}


type GslbPolicyGeoLocationIpv6MultipleFields struct {
    Ipv6Sub string `json:"ipv6-sub"`
    Ipv6MaskSub int `json:"ipv6-mask-sub"`
    Ipv6Addr2Sub string `json:"ipv6-addr2-sub"`
}

func (p *GslbPolicyGeoLocation) GetId() string{
    return p.Inst.Name
}

func (p *GslbPolicyGeoLocation) getPath() string{
    return "gslb/policy/"+p.Inst.PolicyName+"/geo-location"
}

func (p *GslbPolicyGeoLocation) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyGeoLocation::Post")
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

func (p *GslbPolicyGeoLocation) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyGeoLocation::Get")
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
func (p *GslbPolicyGeoLocation) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyGeoLocation::Put")
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

func (p *GslbPolicyGeoLocation) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyGeoLocation::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
