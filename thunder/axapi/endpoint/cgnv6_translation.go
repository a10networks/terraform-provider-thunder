

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6Translation struct {
	Inst struct {

    IcmpTimeout Cgnv6TranslationIcmpTimeout `json:"icmp-timeout"`
    ServiceTimeoutList []Cgnv6TranslationServiceTimeoutList `json:"service-timeout-list"`
    TcpTimeout int `json:"tcp-timeout" dval:"300"`
    UdpTimeout int `json:"udp-timeout" dval:"300"`
    Uuid string `json:"uuid"`

	} `json:"translation"`
}


type Cgnv6TranslationIcmpTimeout struct {
    IcmpTimeoutVal int `json:"icmp-timeout-val"`
    Fast int `json:"fast"`
}


type Cgnv6TranslationServiceTimeoutList struct {
    ServiceType string `json:"service-type"`
    Port int `json:"port"`
    PortEnd int `json:"port-end"`
    TimeoutVal int `json:"timeout-val"`
    Fast int `json:"fast"`
    Uuid string `json:"uuid"`
}

func (p *Cgnv6Translation) GetId() string{
    return "1"
}

func (p *Cgnv6Translation) getPath() string{
    return "cgnv6/translation"
}

func (p *Cgnv6Translation) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Translation::Post")
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

func (p *Cgnv6Translation) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Translation::Get")
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
func (p *Cgnv6Translation) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Translation::Put")
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

func (p *Cgnv6Translation) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Translation::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
