

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationServerOcsp struct {
	Inst struct {

    InstanceList []AamAuthenticationServerOcspInstanceList `json:"instance-list"`
    SamplingEnable []AamAuthenticationServerOcspSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"ocsp"`
}


type AamAuthenticationServerOcspInstanceList struct {
    Name string `json:"name"`
    Url string `json:"url"`
    ResponderCa string `json:"responder-ca"`
    ResponderCert string `json:"responder-cert"`
    HealthCheck int `json:"health-check"`
    HealthCheckString string `json:"health-check-string"`
    HealthCheckDisable int `json:"health-check-disable"`
    PortHealthCheck string `json:"port-health-check"`
    PortHealthCheckDisable int `json:"port-health-check-disable"`
    HttpVersion int `json:"http-version"`
    VersionType string `json:"version-type"`
    Uuid string `json:"uuid"`
    SamplingEnable []AamAuthenticationServerOcspInstanceListSamplingEnable `json:"sampling-enable"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
}


type AamAuthenticationServerOcspInstanceListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type AamAuthenticationServerOcspSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *AamAuthenticationServerOcsp) GetId() string{
    return "1"
}

func (p *AamAuthenticationServerOcsp) getPath() string{
    return "aam/authentication/server/ocsp"
}

func (p *AamAuthenticationServerOcsp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerOcsp::Post")
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

func (p *AamAuthenticationServerOcsp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerOcsp::Get")
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
func (p *AamAuthenticationServerOcsp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerOcsp::Put")
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

func (p *AamAuthenticationServerOcsp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerOcsp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
