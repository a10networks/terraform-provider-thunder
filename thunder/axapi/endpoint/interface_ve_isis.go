

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceVeIsis struct {
	Inst struct {

    Authentication InterfaceVeIsisAuthentication `json:"authentication"`
    BfdCfg InterfaceVeIsisBfdCfg `json:"bfd-cfg"`
    CircuitType string `json:"circuit-type" dval:"level-1-2"`
    CsnpIntervalList []InterfaceVeIsisCsnpIntervalList `json:"csnp-interval-list"`
    HelloIntervalList []InterfaceVeIsisHelloIntervalList `json:"hello-interval-list"`
    HelloIntervalMinimalList []InterfaceVeIsisHelloIntervalMinimalList `json:"hello-interval-minimal-list"`
    HelloMultiplierList []InterfaceVeIsisHelloMultiplierList `json:"hello-multiplier-list"`
    LspInterval int `json:"lsp-interval" dval:"33"`
    MeshGroup InterfaceVeIsisMeshGroup `json:"mesh-group"`
    MetricList []InterfaceVeIsisMetricList `json:"metric-list"`
    Network string `json:"network"`
    Padding int `json:"padding" dval:"1"`
    PasswordList []InterfaceVeIsisPasswordList `json:"password-list"`
    PriorityList []InterfaceVeIsisPriorityList `json:"priority-list"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    Uuid string `json:"uuid"`
    WideMetricList []InterfaceVeIsisWideMetricList `json:"wide-metric-list"`
    Ifnum string 

	} `json:"isis"`
}


type InterfaceVeIsisAuthentication struct {
    SendOnlyList []InterfaceVeIsisAuthenticationSendOnlyList `json:"send-only-list"`
    ModeList []InterfaceVeIsisAuthenticationModeList `json:"mode-list"`
    KeyChainList []InterfaceVeIsisAuthenticationKeyChainList `json:"key-chain-list"`
}


type InterfaceVeIsisAuthenticationSendOnlyList struct {
    SendOnly int `json:"send-only"`
    Level string `json:"level"`
}


type InterfaceVeIsisAuthenticationModeList struct {
    Mode string `json:"mode"`
    Level string `json:"level"`
}


type InterfaceVeIsisAuthenticationKeyChainList struct {
    KeyChain string `json:"key-chain"`
    Level string `json:"level"`
}


type InterfaceVeIsisBfdCfg struct {
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
}


type InterfaceVeIsisCsnpIntervalList struct {
    CsnpInterval int `json:"csnp-interval" dval:"10"`
    Level string `json:"level"`
}


type InterfaceVeIsisHelloIntervalList struct {
    HelloInterval int `json:"hello-interval" dval:"10"`
    Level string `json:"level"`
}


type InterfaceVeIsisHelloIntervalMinimalList struct {
    HelloIntervalMinimal int `json:"hello-interval-minimal"`
    Level string `json:"level"`
}


type InterfaceVeIsisHelloMultiplierList struct {
    HelloMultiplier int `json:"hello-multiplier" dval:"3"`
    Level string `json:"level"`
}


type InterfaceVeIsisMeshGroup struct {
    Value int `json:"value"`
    Blocked int `json:"blocked"`
}


type InterfaceVeIsisMetricList struct {
    Metric int `json:"metric" dval:"10"`
    Level string `json:"level"`
}


type InterfaceVeIsisPasswordList struct {
    Password string `json:"password"`
    Level string `json:"level"`
}


type InterfaceVeIsisPriorityList struct {
    Priority int `json:"priority" dval:"64"`
    Level string `json:"level"`
}


type InterfaceVeIsisWideMetricList struct {
    WideMetric int `json:"wide-metric" dval:"10"`
    Level string `json:"level"`
}

func (p *InterfaceVeIsis) GetId() string{
    return "1"
}

func (p *InterfaceVeIsis) getPath() string{
    return "interface/ve/" +p.Inst.Ifnum + "/isis"
}

func (p *InterfaceVeIsis) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVeIsis::Post")
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

func (p *InterfaceVeIsis) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVeIsis::Get")
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
func (p *InterfaceVeIsis) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVeIsis::Put")
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

func (p *InterfaceVeIsis) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVeIsis::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
