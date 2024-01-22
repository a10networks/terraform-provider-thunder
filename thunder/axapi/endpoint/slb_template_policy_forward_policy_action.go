

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplatePolicyForwardPolicyAction struct {
	Inst struct {

    Action1 string `json:"action1"`
    DropMessage string `json:"drop-message"`
    DropRedirectUrl string `json:"drop-redirect-url"`
    DropResponseCode int `json:"drop-response-code"`
    FakeSg string `json:"fake-sg"`
    FallBack string `json:"fall-back"`
    FallBackSnat string `json:"fall-back-snat"`
    FallBackSnatPtOnly int `json:"fall-back-snat-pt-only"`
    ForwardSnat string `json:"forward-snat"`
    ForwardSnatPtOnly int `json:"forward-snat-pt-only"`
    HttpStatusCode string `json:"http-status-code" dval:"302"`
    Log int `json:"log"`
    Name string `json:"name"`
    ProxyChaining int `json:"proxy-chaining"`
    ProxyChainingBypass int `json:"proxy-chaining-bypass"`
    RealSg string `json:"real-sg"`
    SamplingEnable []SlbTemplatePolicyForwardPolicyActionSamplingEnable `json:"sampling-enable"`
    SupportCertFetch int `json:"support-cert-fetch"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"action"`
}


type SlbTemplatePolicyForwardPolicyActionSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SlbTemplatePolicyForwardPolicyAction) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplatePolicyForwardPolicyAction) getPath() string{
    return "slb/template/policy/"+p.Inst.Name+"/forward-policy/action"
}

func (p *SlbTemplatePolicyForwardPolicyAction) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicyAction::Post")
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

func (p *SlbTemplatePolicyForwardPolicyAction) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicyAction::Get")
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
func (p *SlbTemplatePolicyForwardPolicyAction) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicyAction::Put")
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

func (p *SlbTemplatePolicyForwardPolicyAction) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicyAction::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
