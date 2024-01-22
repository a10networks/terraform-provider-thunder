

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationServerOcspInstance struct {
	Inst struct {

    HealthCheck int `json:"health-check"`
    HealthCheckDisable int `json:"health-check-disable"`
    HealthCheckString string `json:"health-check-string"`
    HttpVersion int `json:"http-version"`
    Name string `json:"name"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
    PortHealthCheck string `json:"port-health-check"`
    PortHealthCheckDisable int `json:"port-health-check-disable"`
    ResponderCa string `json:"responder-ca"`
    ResponderCert string `json:"responder-cert"`
    SamplingEnable []AamAuthenticationServerOcspInstanceSamplingEnable `json:"sampling-enable"`
    Url string `json:"url"`
    Uuid string `json:"uuid"`
    VersionType string `json:"version-type"`

	} `json:"instance"`
}


type AamAuthenticationServerOcspInstanceSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *AamAuthenticationServerOcspInstance) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *AamAuthenticationServerOcspInstance) getPath() string{
    return "aam/authentication/server/ocsp/instance"
}

func (p *AamAuthenticationServerOcspInstance) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerOcspInstance::Post")
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

func (p *AamAuthenticationServerOcspInstance) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerOcspInstance::Get")
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
func (p *AamAuthenticationServerOcspInstance) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerOcspInstance::Put")
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

func (p *AamAuthenticationServerOcspInstance) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerOcspInstance::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
