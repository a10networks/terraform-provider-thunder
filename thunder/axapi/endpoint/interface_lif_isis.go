

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceLifIsis struct {
	Inst struct {

    Authentication InterfaceLifIsisAuthentication `json:"authentication"`
    BfdCfg InterfaceLifIsisBfdCfg `json:"bfd-cfg"`
    CircuitType string `json:"circuit-type" dval:"level-1-2"`
    CsnpIntervalList []InterfaceLifIsisCsnpIntervalList `json:"csnp-interval-list"`
    HelloIntervalList []InterfaceLifIsisHelloIntervalList `json:"hello-interval-list"`
    HelloIntervalMinimalList []InterfaceLifIsisHelloIntervalMinimalList `json:"hello-interval-minimal-list"`
    HelloMultiplierList []InterfaceLifIsisHelloMultiplierList `json:"hello-multiplier-list"`
    LspInterval int `json:"lsp-interval" dval:"33"`
    MeshGroup InterfaceLifIsisMeshGroup `json:"mesh-group"`
    MetricList []InterfaceLifIsisMetricList `json:"metric-list"`
    Network string `json:"network"`
    Padding int `json:"padding" dval:"1"`
    PasswordList []InterfaceLifIsisPasswordList `json:"password-list"`
    PriorityList []InterfaceLifIsisPriorityList `json:"priority-list"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    Uuid string `json:"uuid"`
    WideMetricList []InterfaceLifIsisWideMetricList `json:"wide-metric-list"`
    Ifname string 

	} `json:"isis"`
}


type InterfaceLifIsisAuthentication struct {
    SendOnlyList []InterfaceLifIsisAuthenticationSendOnlyList `json:"send-only-list"`
    ModeList []InterfaceLifIsisAuthenticationModeList `json:"mode-list"`
    KeyChainList []InterfaceLifIsisAuthenticationKeyChainList `json:"key-chain-list"`
}


type InterfaceLifIsisAuthenticationSendOnlyList struct {
    SendOnly int `json:"send-only"`
    Level string `json:"level"`
}


type InterfaceLifIsisAuthenticationModeList struct {
    Mode string `json:"mode"`
    Level string `json:"level"`
}


type InterfaceLifIsisAuthenticationKeyChainList struct {
    KeyChain string `json:"key-chain"`
    Level string `json:"level"`
}


type InterfaceLifIsisBfdCfg struct {
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
}


type InterfaceLifIsisCsnpIntervalList struct {
    CsnpInterval int `json:"csnp-interval" dval:"10"`
    Level string `json:"level"`
}


type InterfaceLifIsisHelloIntervalList struct {
    HelloInterval int `json:"hello-interval" dval:"10"`
    Level string `json:"level"`
}


type InterfaceLifIsisHelloIntervalMinimalList struct {
    HelloIntervalMinimal int `json:"hello-interval-minimal"`
    Level string `json:"level"`
}


type InterfaceLifIsisHelloMultiplierList struct {
    HelloMultiplier int `json:"hello-multiplier" dval:"3"`
    Level string `json:"level"`
}


type InterfaceLifIsisMeshGroup struct {
    Value int `json:"value"`
    Blocked int `json:"blocked"`
}


type InterfaceLifIsisMetricList struct {
    Metric int `json:"metric" dval:"10"`
    Level string `json:"level"`
}


type InterfaceLifIsisPasswordList struct {
    Password string `json:"password"`
    Level string `json:"level"`
}


type InterfaceLifIsisPriorityList struct {
    Priority int `json:"priority" dval:"64"`
    Level string `json:"level"`
}


type InterfaceLifIsisWideMetricList struct {
    WideMetric int `json:"wide-metric" dval:"10"`
    Level string `json:"level"`
}

func (p *InterfaceLifIsis) GetId() string{
    return "1"
}

func (p *InterfaceLifIsis) getPath() string{
    return "interface/lif/" +p.Inst.Ifname + "/isis"
}

func (p *InterfaceLifIsis) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifIsis::Post")
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

func (p *InterfaceLifIsis) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifIsis::Get")
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
func (p *InterfaceLifIsis) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifIsis::Put")
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

func (p *InterfaceLifIsis) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifIsis::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
