

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type WebService struct {
	Inst struct {

    AutoRedirtDisable int `json:"auto-redirt-disable"`
    AxapiIdle int `json:"axapi-idle" dval:"10"`
    AxapiSessionLimit int `json:"axapi-session-limit" dval:"30"`
    GuiIdle int `json:"gui-idle" dval:"10"`
    GuiSessionLimit int `json:"gui-session-limit" dval:"30"`
    KeepAliveTimeout int `json:"keep-alive-timeout" dval:"30"`
    LoginMessage string `json:"login-message"`
    MaxKeepAliveRequests int `json:"max-keep-alive-requests" dval:"100"`
    MpmMaxConn int `json:"mpm-max-conn"`
    MpmMaxConnPerChild int `json:"mpm-max-conn-per-child"`
    MpmMinSpareConn int `json:"mpm-min-spare-conn"`
    Port int `json:"port" dval:"80"`
    PreLoginMessage string `json:"pre-login-message"`
    PublicApis []WebServicePublicApis `json:"public-apis"`
    Secure WebServiceSecure3676 `json:"secure"`
    SecurePort int `json:"secure-port" dval:"443"`
    SecureServerDisable int `json:"secure-server-disable"`
    ServerDisable int `json:"server-disable"`
    Uuid string `json:"uuid"`

	} `json:"web-service"`
}


type WebServicePublicApis struct {
    ApiUri string `json:"api-uri"`
}


type WebServiceSecure3676 struct {
    Restart int `json:"restart"`
    Wipe int `json:"wipe"`
    Generate WebServiceSecureGenerate3677 `json:"generate"`
    Regenerate WebServiceSecureRegenerate3678 `json:"regenerate"`
    Certificate WebServiceSecureCertificate3679 `json:"certificate"`
    PrivateKey WebServiceSecurePrivateKey3680 `json:"private-key"`
}


type WebServiceSecureGenerate3677 struct {
    DomainName string `json:"domain-name"`
    Country string `json:"country"`
    State string `json:"state"`
}


type WebServiceSecureRegenerate3678 struct {
    DomainName string `json:"domain-name"`
    Country string `json:"country"`
    State string `json:"state"`
}


type WebServiceSecureCertificate3679 struct {
    Load int `json:"load"`
    UseMgmtPort int `json:"use-mgmt-port"`
    FileUrl string `json:"file-url"`
}


type WebServiceSecurePrivateKey3680 struct {
    Load int `json:"load"`
    UseMgmtPort int `json:"use-mgmt-port"`
    FileUrl string `json:"file-url"`
}

func (p *WebService) GetId() string{
    return "1"
}

func (p *WebService) getPath() string{
    return "web-service"
}

func (p *WebService) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("WebService::Post")
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

func (p *WebService) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("WebService::Get")
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
func (p *WebService) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("WebService::Put")
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

func (p *WebService) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("WebService::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
