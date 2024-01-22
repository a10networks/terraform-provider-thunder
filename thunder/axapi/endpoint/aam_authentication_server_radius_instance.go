

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationServerRadiusInstance struct {
	Inst struct {

    AccountingPort int `json:"accounting-port" dval:"1813"`
    AcctPortHm string `json:"acct-port-hm"`
    AcctPortHmDisable int `json:"acct-port-hm-disable"`
    AuthType string `json:"auth-type"`
    Encrypted string `json:"encrypted"`
    HealthCheck int `json:"health-check"`
    HealthCheckDisable int `json:"health-check-disable"`
    HealthCheckString string `json:"health-check-string"`
    Host AamAuthenticationServerRadiusInstanceHost `json:"host"`
    Interval int `json:"interval" dval:"3"`
    Name string `json:"name"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
    Port int `json:"port" dval:"1812"`
    PortHm string `json:"port-hm"`
    PortHmDisable int `json:"port-hm-disable"`
    Retry int `json:"retry" dval:"5"`
    SamplingEnable []AamAuthenticationServerRadiusInstanceSamplingEnable `json:"sampling-enable"`
    Secret int `json:"secret"`
    SecretString string `json:"secret-string"`
    Uuid string `json:"uuid"`

	} `json:"instance"`
}


type AamAuthenticationServerRadiusInstanceHost struct {
    Hostip string `json:"hostip"`
    Hostipv6 string `json:"hostipv6"`
}


type AamAuthenticationServerRadiusInstanceSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *AamAuthenticationServerRadiusInstance) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *AamAuthenticationServerRadiusInstance) getPath() string{
    return "aam/authentication/server/radius/instance"
}

func (p *AamAuthenticationServerRadiusInstance) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerRadiusInstance::Post")
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

func (p *AamAuthenticationServerRadiusInstance) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerRadiusInstance::Get")
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
func (p *AamAuthenticationServerRadiusInstance) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerRadiusInstance::Put")
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

func (p *AamAuthenticationServerRadiusInstance) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerRadiusInstance::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
