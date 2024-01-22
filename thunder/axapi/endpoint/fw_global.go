

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwGlobal struct {
	Inst struct {

    AlgProcessing string `json:"alg-processing" dval:"honor-rule-set"`
    AllowNonSynSessionCreate int `json:"allow-non-syn-session-create"`
    DisableAppList []FwGlobalDisableAppList `json:"disable-app-list"`
    DisableApplicationMetrics int `json:"disable-application-metrics"`
    DisableIpFwSessions int `json:"disable-ip-fw-sessions"`
    DisableUndeterminedRuleLogs int `json:"disable-undetermined-rule-logs"`
    DsrModeSupport string `json:"dsr-mode-support"`
    ExtendedMatching string `json:"extended-matching"`
    InboundRefresh string `json:"inbound-refresh" dval:"enable"`
    InboundRefreshFullCone string `json:"inbound-refresh-full-cone" dval:"enable"`
    ListenOnPortTimeout int `json:"listen-on-port-timeout" dval:"2"`
    NatipDdosProtection string `json:"natip-ddos-protection" dval:"enable"`
    PermitDefaultAction string `json:"permit-default-action"`
    RespondToUserMac int `json:"respond-to-user-mac"`
    SamplingEnable []FwGlobalSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"global"`
}


type FwGlobalDisableAppList struct {
    DisableApplicationProtocol string `json:"disable-application-protocol"`
    DisableApplicationCategory string `json:"disable-application-category"`
}


type FwGlobalSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *FwGlobal) GetId() string{
    return "1"
}

func (p *FwGlobal) getPath() string{
    return "fw/global"
}

func (p *FwGlobal) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwGlobal::Post")
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

func (p *FwGlobal) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwGlobal::Get")
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
func (p *FwGlobal) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwGlobal::Put")
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

func (p *FwGlobal) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwGlobal::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
