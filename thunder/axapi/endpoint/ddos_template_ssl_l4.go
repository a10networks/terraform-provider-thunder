

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosTemplateSslL4 struct {
	Inst struct {

    Action string `json:"action" dval:"drop"`
    AllowNonTls int `json:"allow-non-tls"`
    AuthConfigCfg DdosTemplateSslL4AuthConfigCfg `json:"auth-config-cfg"`
    CertCfg DdosTemplateSslL4CertCfg `json:"cert-cfg"`
    Disable int `json:"disable"`
    MultiPuThresholdDistribution DdosTemplateSslL4MultiPuThresholdDistribution `json:"multi-pu-threshold-distribution"`
    Renegotiation int `json:"renegotiation"`
    RequestRateLimit int `json:"request-rate-limit"`
    ServerNameList []DdosTemplateSslL4ServerNameList `json:"server-name-list"`
    SslL4TmplName string `json:"ssl-l4-tmpl-name"`
    SslTrafficCheck DdosTemplateSslL4SslTrafficCheck299 `json:"ssl-traffic-check"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"ssl-l4"`
}


type DdosTemplateSslL4AuthConfigCfg struct {
    Timeout int `json:"timeout" dval:"5"`
    Trials int `json:"trials" dval:"5"`
    AuthHandshakeFailAction string `json:"auth-handshake-fail-action"`
}


type DdosTemplateSslL4CertCfg struct {
    Cert string `json:"cert"`
    Key string `json:"key"`
    KeyPassphrase string `json:"key-passphrase"`
    KeyEncrypted string `json:"key-encrypted"`
}


type DdosTemplateSslL4MultiPuThresholdDistribution struct {
    MultiPuThresholdDistributionValue int `json:"multi-pu-threshold-distribution-value"`
    MultiPuThresholdDistributionDisable string `json:"multi-pu-threshold-distribution-disable"`
}


type DdosTemplateSslL4ServerNameList struct {
    ServerName string `json:"server-name"`
    ServerCert string `json:"server-cert"`
    ServerKey string `json:"server-key"`
    ServerPassphrase string `json:"server-passphrase"`
    ServerEncrypted string `json:"server-encrypted"`
}


type DdosTemplateSslL4SslTrafficCheck299 struct {
    HeaderInspection int `json:"header-inspection"`
    HeaderAction string `json:"header-action"`
    CheckResumedConnection int `json:"check-resumed-connection"`
    Uuid string `json:"uuid"`
}

func (p *DdosTemplateSslL4) GetId() string{
    return url.QueryEscape(p.Inst.SslL4TmplName)
}

func (p *DdosTemplateSslL4) getPath() string{
    return "ddos/template/ssl-l4"
}

func (p *DdosTemplateSslL4) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateSslL4::Post")
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

func (p *DdosTemplateSslL4) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateSslL4::Get")
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
func (p *DdosTemplateSslL4) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateSslL4::Put")
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

func (p *DdosTemplateSslL4) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateSslL4::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
