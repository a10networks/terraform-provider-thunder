

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6MapTranslationDomainHealthCheckGateway struct {
	Inst struct {

    AddressList []Cgnv6MapTranslationDomainHealthCheckGatewayAddressList `json:"address-list"`
    Ipv6AddressList []Cgnv6MapTranslationDomainHealthCheckGatewayIpv6AddressList `json:"ipv6-address-list"`
    Uuid string `json:"uuid"`
    WithdrawRoute string `json:"withdraw-route" dval:"any-link-failure"`
    Name string 

	} `json:"health-check-gateway"`
}


type Cgnv6MapTranslationDomainHealthCheckGatewayAddressList struct {
    Ipv4Gateway string `json:"ipv4-gateway"`
}


type Cgnv6MapTranslationDomainHealthCheckGatewayIpv6AddressList struct {
    Ipv6Gateway string `json:"ipv6-gateway"`
}

func (p *Cgnv6MapTranslationDomainHealthCheckGateway) GetId() string{
    return "1"
}

func (p *Cgnv6MapTranslationDomainHealthCheckGateway) getPath() string{
    return "cgnv6/map/translation/domain/" +p.Inst.Name + "/health-check-gateway"
}

func (p *Cgnv6MapTranslationDomainHealthCheckGateway) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6MapTranslationDomainHealthCheckGateway::Post")
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

func (p *Cgnv6MapTranslationDomainHealthCheckGateway) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6MapTranslationDomainHealthCheckGateway::Get")
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
func (p *Cgnv6MapTranslationDomainHealthCheckGateway) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6MapTranslationDomainHealthCheckGateway::Put")
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

func (p *Cgnv6MapTranslationDomainHealthCheckGateway) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6MapTranslationDomainHealthCheckGateway::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
