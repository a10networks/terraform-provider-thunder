

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VrrpAVridBladeParametersTrackingOptionsGatewayIpv6Gateway struct {
	Inst struct {

    Ipv6Address string `json:"ipv6-address"`
    PriorityCost int `json:"priority-cost"`
    Uuid string `json:"uuid"`
    VridVal string 

	} `json:"ipv6-gateway"`
}

func (p *VrrpAVridBladeParametersTrackingOptionsGatewayIpv6Gateway) GetId() string{
    return p.Inst.Ipv6Address
}

func (p *VrrpAVridBladeParametersTrackingOptionsGatewayIpv6Gateway) getPath() string{
    return "vrrp-a/vrid/" +p.Inst.VridVal + "/blade-parameters/tracking-options/gateway/ipv6-gateway"
}

func (p *VrrpAVridBladeParametersTrackingOptionsGatewayIpv6Gateway) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAVridBladeParametersTrackingOptionsGatewayIpv6Gateway::Post")
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

func (p *VrrpAVridBladeParametersTrackingOptionsGatewayIpv6Gateway) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAVridBladeParametersTrackingOptionsGatewayIpv6Gateway::Get")
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
func (p *VrrpAVridBladeParametersTrackingOptionsGatewayIpv6Gateway) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAVridBladeParametersTrackingOptionsGatewayIpv6Gateway::Put")
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

func (p *VrrpAVridBladeParametersTrackingOptionsGatewayIpv6Gateway) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAVridBladeParametersTrackingOptionsGatewayIpv6Gateway::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
