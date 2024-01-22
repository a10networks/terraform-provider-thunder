

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationServerRadius struct {
	Inst struct {

    InstanceList []AamAuthenticationServerRadiusInstanceList `json:"instance-list"`
    SamplingEnable []AamAuthenticationServerRadiusSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"radius"`
}


type AamAuthenticationServerRadiusInstanceList struct {
    Name string `json:"name"`
    Host AamAuthenticationServerRadiusInstanceListHost `json:"host"`
    Secret int `json:"secret"`
    SecretString string `json:"secret-string"`
    Encrypted string `json:"encrypted"`
    Port int `json:"port" dval:"1812"`
    PortHm string `json:"port-hm"`
    PortHmDisable int `json:"port-hm-disable"`
    Interval int `json:"interval" dval:"3"`
    Retry int `json:"retry" dval:"5"`
    HealthCheck int `json:"health-check"`
    HealthCheckString string `json:"health-check-string"`
    HealthCheckDisable int `json:"health-check-disable"`
    AccountingPort int `json:"accounting-port" dval:"1813"`
    AcctPortHm string `json:"acct-port-hm"`
    AcctPortHmDisable int `json:"acct-port-hm-disable"`
    AuthType string `json:"auth-type"`
    Uuid string `json:"uuid"`
    SamplingEnable []AamAuthenticationServerRadiusInstanceListSamplingEnable `json:"sampling-enable"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
}


type AamAuthenticationServerRadiusInstanceListHost struct {
    Hostip string `json:"hostip"`
    Hostipv6 string `json:"hostipv6"`
}


type AamAuthenticationServerRadiusInstanceListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type AamAuthenticationServerRadiusSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *AamAuthenticationServerRadius) GetId() string{
    return "1"
}

func (p *AamAuthenticationServerRadius) getPath() string{
    return "aam/authentication/server/radius"
}

func (p *AamAuthenticationServerRadius) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerRadius::Post")
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

func (p *AamAuthenticationServerRadius) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerRadius::Get")
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
func (p *AamAuthenticationServerRadius) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerRadius::Put")
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

func (p *AamAuthenticationServerRadius) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerRadius::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
