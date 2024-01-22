

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceTrunkIsis struct {
	Inst struct {

    Authentication InterfaceTrunkIsisAuthentication `json:"authentication"`
    BfdCfg InterfaceTrunkIsisBfdCfg `json:"bfd-cfg"`
    CircuitType string `json:"circuit-type" dval:"level-1-2"`
    CsnpIntervalList []InterfaceTrunkIsisCsnpIntervalList `json:"csnp-interval-list"`
    HelloIntervalList []InterfaceTrunkIsisHelloIntervalList `json:"hello-interval-list"`
    HelloIntervalMinimalList []InterfaceTrunkIsisHelloIntervalMinimalList `json:"hello-interval-minimal-list"`
    HelloMultiplierList []InterfaceTrunkIsisHelloMultiplierList `json:"hello-multiplier-list"`
    LspInterval int `json:"lsp-interval" dval:"33"`
    MeshGroup InterfaceTrunkIsisMeshGroup `json:"mesh-group"`
    MetricList []InterfaceTrunkIsisMetricList `json:"metric-list"`
    Network string `json:"network"`
    Padding int `json:"padding" dval:"1"`
    PasswordList []InterfaceTrunkIsisPasswordList `json:"password-list"`
    PriorityList []InterfaceTrunkIsisPriorityList `json:"priority-list"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    Uuid string `json:"uuid"`
    WideMetricList []InterfaceTrunkIsisWideMetricList `json:"wide-metric-list"`
    Ifnum string 

	} `json:"isis"`
}


type InterfaceTrunkIsisAuthentication struct {
    SendOnlyList []InterfaceTrunkIsisAuthenticationSendOnlyList `json:"send-only-list"`
    ModeList []InterfaceTrunkIsisAuthenticationModeList `json:"mode-list"`
    KeyChainList []InterfaceTrunkIsisAuthenticationKeyChainList `json:"key-chain-list"`
}


type InterfaceTrunkIsisAuthenticationSendOnlyList struct {
    SendOnly int `json:"send-only"`
    Level string `json:"level"`
}


type InterfaceTrunkIsisAuthenticationModeList struct {
    Mode string `json:"mode"`
    Level string `json:"level"`
}


type InterfaceTrunkIsisAuthenticationKeyChainList struct {
    KeyChain string `json:"key-chain"`
    Level string `json:"level"`
}


type InterfaceTrunkIsisBfdCfg struct {
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
}


type InterfaceTrunkIsisCsnpIntervalList struct {
    CsnpInterval int `json:"csnp-interval" dval:"10"`
    Level string `json:"level"`
}


type InterfaceTrunkIsisHelloIntervalList struct {
    HelloInterval int `json:"hello-interval" dval:"10"`
    Level string `json:"level"`
}


type InterfaceTrunkIsisHelloIntervalMinimalList struct {
    HelloIntervalMinimal int `json:"hello-interval-minimal"`
    Level string `json:"level"`
}


type InterfaceTrunkIsisHelloMultiplierList struct {
    HelloMultiplier int `json:"hello-multiplier" dval:"3"`
    Level string `json:"level"`
}


type InterfaceTrunkIsisMeshGroup struct {
    Value int `json:"value"`
    Blocked int `json:"blocked"`
}


type InterfaceTrunkIsisMetricList struct {
    Metric int `json:"metric" dval:"10"`
    Level string `json:"level"`
}


type InterfaceTrunkIsisPasswordList struct {
    Password string `json:"password"`
    Level string `json:"level"`
}


type InterfaceTrunkIsisPriorityList struct {
    Priority int `json:"priority" dval:"64"`
    Level string `json:"level"`
}


type InterfaceTrunkIsisWideMetricList struct {
    WideMetric int `json:"wide-metric" dval:"10"`
    Level string `json:"level"`
}

func (p *InterfaceTrunkIsis) GetId() string{
    return "1"
}

func (p *InterfaceTrunkIsis) getPath() string{
    return "interface/trunk/" +p.Inst.Ifnum + "/isis"
}

func (p *InterfaceTrunkIsis) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkIsis::Post")
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

func (p *InterfaceTrunkIsis) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkIsis::Get")
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
func (p *InterfaceTrunkIsis) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkIsis::Put")
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

func (p *InterfaceTrunkIsis) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkIsis::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
