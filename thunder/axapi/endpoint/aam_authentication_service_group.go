

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationServiceGroup struct {
	Inst struct {

    HealthCheck string `json:"health-check"`
    HealthCheckDisable int `json:"health-check-disable"`
    LbMethod string `json:"lb-method"`
    MemberList []AamAuthenticationServiceGroupMemberList `json:"member-list"`
    Name string `json:"name"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
    Protocol string `json:"protocol"`
    SamplingEnable []AamAuthenticationServiceGroupSamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"service-group"`
}


type AamAuthenticationServiceGroupMemberList struct {
    Name string `json:"name"`
    Port int `json:"port"`
    MemberState string `json:"member-state" dval:"enable"`
    MemberPriority int `json:"member-priority"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []AamAuthenticationServiceGroupMemberListSamplingEnable `json:"sampling-enable"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
}


type AamAuthenticationServiceGroupMemberListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type AamAuthenticationServiceGroupSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *AamAuthenticationServiceGroup) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *AamAuthenticationServiceGroup) getPath() string{
    return "aam/authentication/service-group"
}

func (p *AamAuthenticationServiceGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServiceGroup::Post")
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

func (p *AamAuthenticationServiceGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServiceGroup::Get")
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
func (p *AamAuthenticationServiceGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServiceGroup::Put")
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

func (p *AamAuthenticationServiceGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServiceGroup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
