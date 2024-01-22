

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VrrpAVridBladeParametersTrackingOptionsGatewayIpv4Gateway struct {
	Inst struct {

    IpAddress string `json:"ip-address"`
    PriorityCost int `json:"priority-cost"`
    Uuid string `json:"uuid"`
    VridVal string 

	} `json:"ipv4-gateway"`
}

func (p *VrrpAVridBladeParametersTrackingOptionsGatewayIpv4Gateway) GetId() string{
    return p.Inst.IpAddress
}

func (p *VrrpAVridBladeParametersTrackingOptionsGatewayIpv4Gateway) getPath() string{
    return "vrrp-a/vrid/" +p.Inst.VridVal + "/blade-parameters/tracking-options/gateway/ipv4-gateway"
}

func (p *VrrpAVridBladeParametersTrackingOptionsGatewayIpv4Gateway) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAVridBladeParametersTrackingOptionsGatewayIpv4Gateway::Post")
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

func (p *VrrpAVridBladeParametersTrackingOptionsGatewayIpv4Gateway) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAVridBladeParametersTrackingOptionsGatewayIpv4Gateway::Get")
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
func (p *VrrpAVridBladeParametersTrackingOptionsGatewayIpv4Gateway) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAVridBladeParametersTrackingOptionsGatewayIpv4Gateway::Put")
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

func (p *VrrpAVridBladeParametersTrackingOptionsGatewayIpv4Gateway) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAVridBladeParametersTrackingOptionsGatewayIpv4Gateway::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
