

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosDetectionAgent struct {
	Inst struct {

    AgentName string `json:"agent-name"`
    AgentType string `json:"agent-type"`
    AgentV4Addr string `json:"agent-v4-addr"`
    AgentV6Addr string `json:"agent-v6-addr"`
    Netflow DdosDetectionAgentNetflow128 `json:"netflow"`
    SamplingEnable []DdosDetectionAgentSamplingEnable `json:"sampling-enable"`
    Sflow DdosDetectionAgentSflow129 `json:"sflow"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"agent"`
}


type DdosDetectionAgentNetflow128 struct {
    NetflowSamplesCollection string `json:"netflow-samples-collection" dval:"enable"`
    NetflowSamplingRate int `json:"netflow-sampling-rate" dval:"1"`
    ActiveTimeout int `json:"active-timeout"`
    InactiveTimeout int `json:"inactive-timeout"`
    Uuid string `json:"uuid"`
}


type DdosDetectionAgentSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type DdosDetectionAgentSflow129 struct {
    SflowPktSamplesCollection string `json:"sflow-pkt-samples-collection" dval:"enable"`
    Uuid string `json:"uuid"`
}

func (p *DdosDetectionAgent) GetId() string{
    return url.QueryEscape(p.Inst.AgentName)
}

func (p *DdosDetectionAgent) getPath() string{
    return "ddos/detection/agent"
}

func (p *DdosDetectionAgent) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDetectionAgent::Post")
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

func (p *DdosDetectionAgent) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDetectionAgent::Get")
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
func (p *DdosDetectionAgent) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDetectionAgent::Put")
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

func (p *DdosDetectionAgent) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDetectionAgent::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
