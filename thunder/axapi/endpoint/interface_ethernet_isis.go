

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceEthernetIsis struct {
	Inst struct {

    Authentication InterfaceEthernetIsisAuthentication `json:"authentication"`
    BfdCfg InterfaceEthernetIsisBfdCfg `json:"bfd-cfg"`
    CircuitType string `json:"circuit-type" dval:"level-1-2"`
    CsnpIntervalList []InterfaceEthernetIsisCsnpIntervalList `json:"csnp-interval-list"`
    HelloIntervalList []InterfaceEthernetIsisHelloIntervalList `json:"hello-interval-list"`
    HelloIntervalMinimalList []InterfaceEthernetIsisHelloIntervalMinimalList `json:"hello-interval-minimal-list"`
    HelloMultiplierList []InterfaceEthernetIsisHelloMultiplierList `json:"hello-multiplier-list"`
    LspInterval int `json:"lsp-interval" dval:"33"`
    MeshGroup InterfaceEthernetIsisMeshGroup `json:"mesh-group"`
    MetricList []InterfaceEthernetIsisMetricList `json:"metric-list"`
    Network string `json:"network"`
    Padding int `json:"padding" dval:"1"`
    PasswordList []InterfaceEthernetIsisPasswordList `json:"password-list"`
    PriorityList []InterfaceEthernetIsisPriorityList `json:"priority-list"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    Uuid string `json:"uuid"`
    WideMetricList []InterfaceEthernetIsisWideMetricList `json:"wide-metric-list"`
    Ifnum string 

	} `json:"isis"`
}


type InterfaceEthernetIsisAuthentication struct {
    SendOnlyList []InterfaceEthernetIsisAuthenticationSendOnlyList `json:"send-only-list"`
    ModeList []InterfaceEthernetIsisAuthenticationModeList `json:"mode-list"`
    KeyChainList []InterfaceEthernetIsisAuthenticationKeyChainList `json:"key-chain-list"`
}


type InterfaceEthernetIsisAuthenticationSendOnlyList struct {
    SendOnly int `json:"send-only"`
    Level string `json:"level"`
}


type InterfaceEthernetIsisAuthenticationModeList struct {
    Mode string `json:"mode"`
    Level string `json:"level"`
}


type InterfaceEthernetIsisAuthenticationKeyChainList struct {
    KeyChain string `json:"key-chain"`
    Level string `json:"level"`
}


type InterfaceEthernetIsisBfdCfg struct {
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
}


type InterfaceEthernetIsisCsnpIntervalList struct {
    CsnpInterval int `json:"csnp-interval" dval:"10"`
    Level string `json:"level"`
}


type InterfaceEthernetIsisHelloIntervalList struct {
    HelloInterval int `json:"hello-interval" dval:"10"`
    Level string `json:"level"`
}


type InterfaceEthernetIsisHelloIntervalMinimalList struct {
    HelloIntervalMinimal int `json:"hello-interval-minimal"`
    Level string `json:"level"`
}


type InterfaceEthernetIsisHelloMultiplierList struct {
    HelloMultiplier int `json:"hello-multiplier" dval:"3"`
    Level string `json:"level"`
}


type InterfaceEthernetIsisMeshGroup struct {
    Value int `json:"value"`
    Blocked int `json:"blocked"`
}


type InterfaceEthernetIsisMetricList struct {
    Metric int `json:"metric" dval:"10"`
    Level string `json:"level"`
}


type InterfaceEthernetIsisPasswordList struct {
    Password string `json:"password"`
    Level string `json:"level"`
}


type InterfaceEthernetIsisPriorityList struct {
    Priority int `json:"priority" dval:"64"`
    Level string `json:"level"`
}


type InterfaceEthernetIsisWideMetricList struct {
    WideMetric int `json:"wide-metric" dval:"10"`
    Level string `json:"level"`
}

func (p *InterfaceEthernetIsis) GetId() string{
    return "1"
}

func (p *InterfaceEthernetIsis) getPath() string{
    return "interface/ethernet/" +p.Inst.Ifnum + "/isis"
}

func (p *InterfaceEthernetIsis) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetIsis::Post")
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

func (p *InterfaceEthernetIsis) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetIsis::Get")
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
func (p *InterfaceEthernetIsis) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetIsis::Put")
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

func (p *InterfaceEthernetIsis) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetIsis::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
