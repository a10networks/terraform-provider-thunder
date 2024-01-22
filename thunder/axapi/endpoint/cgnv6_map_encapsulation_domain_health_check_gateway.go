

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6MapEncapsulationDomainHealthCheckGateway struct {
	Inst struct {

    AddressList []Cgnv6MapEncapsulationDomainHealthCheckGatewayAddressList `json:"address-list"`
    Ipv6AddressList []Cgnv6MapEncapsulationDomainHealthCheckGatewayIpv6AddressList `json:"ipv6-address-list"`
    Uuid string `json:"uuid"`
    WithdrawRoute string `json:"withdraw-route" dval:"any-link-failure"`
    Name string 

	} `json:"health-check-gateway"`
}


type Cgnv6MapEncapsulationDomainHealthCheckGatewayAddressList struct {
    Ipv4Gateway string `json:"ipv4-gateway"`
}


type Cgnv6MapEncapsulationDomainHealthCheckGatewayIpv6AddressList struct {
    Ipv6Gateway string `json:"ipv6-gateway"`
}

func (p *Cgnv6MapEncapsulationDomainHealthCheckGateway) GetId() string{
    return "1"
}

func (p *Cgnv6MapEncapsulationDomainHealthCheckGateway) getPath() string{
    return "cgnv6/map/encapsulation/domain/" +p.Inst.Name + "/health-check-gateway"
}

func (p *Cgnv6MapEncapsulationDomainHealthCheckGateway) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6MapEncapsulationDomainHealthCheckGateway::Post")
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

func (p *Cgnv6MapEncapsulationDomainHealthCheckGateway) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6MapEncapsulationDomainHealthCheckGateway::Get")
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
func (p *Cgnv6MapEncapsulationDomainHealthCheckGateway) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6MapEncapsulationDomainHealthCheckGateway::Put")
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

func (p *Cgnv6MapEncapsulationDomainHealthCheckGateway) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6MapEncapsulationDomainHealthCheckGateway::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
