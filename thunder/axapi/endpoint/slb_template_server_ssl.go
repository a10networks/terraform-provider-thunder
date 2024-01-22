

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateServerSsl struct {
	Inst struct {

    AlertType string `json:"alert-type"`
    CaCerts []SlbTemplateServerSslCaCerts `json:"ca-certs"`
    Certificate SlbTemplateServerSslCertificate1474 `json:"certificate"`
    CipherTemplate string `json:"cipher-template"`
    CipherWithoutPrioList []SlbTemplateServerSslCipherWithoutPrioList `json:"cipher-without-prio-list"`
    CloseNotify int `json:"close-notify"`
    CrlCerts []SlbTemplateServerSslCrlCerts `json:"crl-certs"`
    Dgversion int `json:"dgversion" dval:"31"`
    DhType string `json:"dh-type"`
    EarlyData int `json:"early-data"`
    EcList []SlbTemplateServerSslEcList `json:"ec-list"`
    EnableSsliFtpAlg int `json:"enable-ssli-ftp-alg"`
    EnableTlsAlertLogging int `json:"enable-tls-alert-logging"`
    ForwardProxyEnable int `json:"forward-proxy-enable"`
    HandshakeLoggingEnable int `json:"handshake-logging-enable"`
    Name string `json:"name"`
    OcspStapling int `json:"ocsp-stapling"`
    RenegotiationDisable int `json:"renegotiation-disable"`
    ServerCertificateError []SlbTemplateServerSslServerCertificateError `json:"server-certificate-error"`
    ServerName string `json:"server-name"`
    SessionCacheSize int `json:"session-cache-size"`
    SessionCacheTimeout int `json:"session-cache-timeout"`
    SessionTicketEnable int `json:"session-ticket-enable"`
    SharedPartitionCipherTemplate int `json:"shared-partition-cipher-template"`
    SsliLogging int `json:"ssli-logging"`
    Sslilogging string `json:"sslilogging"`
    TemplateCipherShared string `json:"template-cipher-shared"`
    UseClientSni int `json:"use-client-sni"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Version int `json:"version" dval:"33"`

	} `json:"server-ssl"`
}


type SlbTemplateServerSslCaCerts struct {
    CaCert string `json:"ca-cert"`
    CaCertPartitionShared int `json:"ca-cert-partition-shared"`
    ServerOcspSrvr string `json:"server-ocsp-srvr"`
    ServerOcspSg string `json:"server-ocsp-sg"`
}


type SlbTemplateServerSslCertificate1474 struct {
    Cert string `json:"cert"`
    Key string `json:"key"`
    Passphrase string `json:"passphrase"`
    Encrypted string `json:"encrypted"`
    Shared int `json:"shared"`
    Uuid string `json:"uuid"`
}


type SlbTemplateServerSslCipherWithoutPrioList struct {
    CipherWoPrio string `json:"cipher-wo-prio"`
}


type SlbTemplateServerSslCrlCerts struct {
    Crl string `json:"crl"`
    CrlPartitionShared int `json:"crl-partition-shared"`
}


type SlbTemplateServerSslEcList struct {
    Ec string `json:"ec"`
}


type SlbTemplateServerSslServerCertificateError struct {
    ErrorType string `json:"error-type"`
}

func (p *SlbTemplateServerSsl) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateServerSsl) getPath() string{
    return "slb/template/server-ssl"
}

func (p *SlbTemplateServerSsl) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateServerSsl::Post")
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

func (p *SlbTemplateServerSsl) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateServerSsl::Get")
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
func (p *SlbTemplateServerSsl) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateServerSsl::Put")
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

func (p *SlbTemplateServerSsl) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateServerSsl::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
