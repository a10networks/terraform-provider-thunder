

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type WebServiceSecure struct {
	Inst struct {

    Certificate WebServiceSecureCertificate3672 `json:"certificate"`
    Generate WebServiceSecureGenerate3673 `json:"generate"`
    PrivateKey WebServiceSecurePrivateKey3674 `json:"private-key"`
    Regenerate WebServiceSecureRegenerate3675 `json:"regenerate"`
    Restart int `json:"restart"`
    Wipe int `json:"wipe"`

	} `json:"secure"`
}


type WebServiceSecureCertificate3672 struct {
    Load int `json:"load"`
    UseMgmtPort int `json:"use-mgmt-port"`
    FileUrl string `json:"file-url"`
}


type WebServiceSecureGenerate3673 struct {
    DomainName string `json:"domain-name"`
    Country string `json:"country"`
    State string `json:"state"`
}


type WebServiceSecurePrivateKey3674 struct {
    Load int `json:"load"`
    UseMgmtPort int `json:"use-mgmt-port"`
    FileUrl string `json:"file-url"`
}


type WebServiceSecureRegenerate3675 struct {
    DomainName string `json:"domain-name"`
    Country string `json:"country"`
    State string `json:"state"`
}

func (p *WebServiceSecure) GetId() string{
    return "1"
}

func (p *WebServiceSecure) getPath() string{
    return "web-service/secure"
}

func (p *WebServiceSecure) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("WebServiceSecure::Post")
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

func (p *WebServiceSecure) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("WebServiceSecure::Get")
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
func (p *WebServiceSecure) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("WebServiceSecure::Put")
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

func (p *WebServiceSecure) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("WebServiceSecure::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
