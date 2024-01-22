

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceLoopbackIsis struct {
	Inst struct {

    Authentication InterfaceLoopbackIsisAuthentication `json:"authentication"`
    BfdCfg InterfaceLoopbackIsisBfdCfg `json:"bfd-cfg"`
    CircuitType string `json:"circuit-type" dval:"level-1-2"`
    CsnpIntervalList []InterfaceLoopbackIsisCsnpIntervalList `json:"csnp-interval-list"`
    HelloIntervalList []InterfaceLoopbackIsisHelloIntervalList `json:"hello-interval-list"`
    HelloIntervalMinimalList []InterfaceLoopbackIsisHelloIntervalMinimalList `json:"hello-interval-minimal-list"`
    HelloMultiplierList []InterfaceLoopbackIsisHelloMultiplierList `json:"hello-multiplier-list"`
    LspInterval int `json:"lsp-interval" dval:"33"`
    MeshGroup InterfaceLoopbackIsisMeshGroup `json:"mesh-group"`
    MetricList []InterfaceLoopbackIsisMetricList `json:"metric-list"`
    Padding int `json:"padding" dval:"1"`
    PasswordList []InterfaceLoopbackIsisPasswordList `json:"password-list"`
    PriorityList []InterfaceLoopbackIsisPriorityList `json:"priority-list"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    Uuid string `json:"uuid"`
    WideMetricList []InterfaceLoopbackIsisWideMetricList `json:"wide-metric-list"`
    Ifnum string 

	} `json:"isis"`
}


type InterfaceLoopbackIsisAuthentication struct {
    SendOnlyList []InterfaceLoopbackIsisAuthenticationSendOnlyList `json:"send-only-list"`
    ModeList []InterfaceLoopbackIsisAuthenticationModeList `json:"mode-list"`
    KeyChainList []InterfaceLoopbackIsisAuthenticationKeyChainList `json:"key-chain-list"`
}


type InterfaceLoopbackIsisAuthenticationSendOnlyList struct {
    SendOnly int `json:"send-only"`
    Level string `json:"level"`
}


type InterfaceLoopbackIsisAuthenticationModeList struct {
    Mode string `json:"mode"`
    Level string `json:"level"`
}


type InterfaceLoopbackIsisAuthenticationKeyChainList struct {
    KeyChain string `json:"key-chain"`
    Level string `json:"level"`
}


type InterfaceLoopbackIsisBfdCfg struct {
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
}


type InterfaceLoopbackIsisCsnpIntervalList struct {
    CsnpInterval int `json:"csnp-interval" dval:"10"`
    Level string `json:"level"`
}


type InterfaceLoopbackIsisHelloIntervalList struct {
    HelloInterval int `json:"hello-interval" dval:"10"`
    Level string `json:"level"`
}


type InterfaceLoopbackIsisHelloIntervalMinimalList struct {
    HelloIntervalMinimal int `json:"hello-interval-minimal"`
    Level string `json:"level"`
}


type InterfaceLoopbackIsisHelloMultiplierList struct {
    HelloMultiplier int `json:"hello-multiplier" dval:"3"`
    Level string `json:"level"`
}


type InterfaceLoopbackIsisMeshGroup struct {
    Value int `json:"value"`
    Blocked int `json:"blocked"`
}


type InterfaceLoopbackIsisMetricList struct {
    Metric int `json:"metric" dval:"10"`
    Level string `json:"level"`
}


type InterfaceLoopbackIsisPasswordList struct {
    Password string `json:"password"`
    Level string `json:"level"`
}


type InterfaceLoopbackIsisPriorityList struct {
    Priority int `json:"priority" dval:"64"`
    Level string `json:"level"`
}


type InterfaceLoopbackIsisWideMetricList struct {
    WideMetric int `json:"wide-metric" dval:"10"`
    Level string `json:"level"`
}

func (p *InterfaceLoopbackIsis) GetId() string{
    return "1"
}

func (p *InterfaceLoopbackIsis) getPath() string{
    return "interface/loopback/" +p.Inst.Ifnum + "/isis"
}

func (p *InterfaceLoopbackIsis) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopbackIsis::Post")
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

func (p *InterfaceLoopbackIsis) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopbackIsis::Get")
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
func (p *InterfaceLoopbackIsis) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopbackIsis::Put")
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

func (p *InterfaceLoopbackIsis) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopbackIsis::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
