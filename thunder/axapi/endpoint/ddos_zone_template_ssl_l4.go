

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneTemplateSslL4 struct {
	Inst struct {

    AllowNonTls int `json:"allow-non-tls"`
    AuthHandshake DdosZoneTemplateSslL4AuthHandshake `json:"auth-handshake"`
    Disable int `json:"disable"`
    Dst DdosZoneTemplateSslL4Dst `json:"dst"`
    MultiPuThresholdDistribution DdosZoneTemplateSslL4MultiPuThresholdDistribution `json:"multi-pu-threshold-distribution"`
    Renegotiation DdosZoneTemplateSslL4Renegotiation `json:"renegotiation"`
    Src DdosZoneTemplateSslL4Src `json:"src"`
    SslL4TmplName string `json:"ssl-l4-tmpl-name"`
    SslTrafficCheck DdosZoneTemplateSslL4SslTrafficCheck317 `json:"ssl-traffic-check"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"ssl-l4"`
}


type DdosZoneTemplateSslL4AuthHandshake struct {
    AuthHandshakeTimeout int `json:"auth-handshake-timeout" dval:"5"`
    AuthHandshakeTrials int `json:"auth-handshake-trials" dval:"5"`
    CertCfg DdosZoneTemplateSslL4AuthHandshakeCertCfg `json:"cert-cfg"`
    ServerNameList []DdosZoneTemplateSslL4AuthHandshakeServerNameList `json:"server-name-list"`
    AuthHandshakePassActionListName string `json:"auth-handshake-pass-action-list-name"`
    AuthHandshakePassAction string `json:"auth-handshake-pass-action"`
    AuthHandshakeFailActionListName string `json:"auth-handshake-fail-action-list-name"`
    AuthHandshakeFailAction string `json:"auth-handshake-fail-action"`
}


type DdosZoneTemplateSslL4AuthHandshakeCertCfg struct {
    Cert string `json:"cert"`
    Key string `json:"key"`
    KeyPassphrase string `json:"key-passphrase"`
    KeyEncrypted string `json:"key-encrypted"`
}


type DdosZoneTemplateSslL4AuthHandshakeServerNameList struct {
    ServerName string `json:"server-name"`
    ServerCert string `json:"server-cert"`
    ServerKey string `json:"server-key"`
    ServerPassphrase string `json:"server-passphrase"`
    ServerEncrypted string `json:"server-encrypted"`
}


type DdosZoneTemplateSslL4Dst struct {
    RateLimit DdosZoneTemplateSslL4DstRateLimit `json:"rate-limit"`
}


type DdosZoneTemplateSslL4DstRateLimit struct {
    Request DdosZoneTemplateSslL4DstRateLimitRequest `json:"request"`
}


type DdosZoneTemplateSslL4DstRateLimitRequest struct {
    DstRequestRateLimit int `json:"dst-request-rate-limit"`
    DstRequestRateLimitActionListName string `json:"dst-request-rate-limit-action-list-name"`
    DstRequestRateLimitAction string `json:"dst-request-rate-limit-action"`
}


type DdosZoneTemplateSslL4MultiPuThresholdDistribution struct {
    MultiPuThresholdDistributionValue int `json:"multi-pu-threshold-distribution-value"`
    MultiPuThresholdDistributionDisable string `json:"multi-pu-threshold-distribution-disable"`
}


type DdosZoneTemplateSslL4Renegotiation struct {
    NumRenegotiation int `json:"num-renegotiation"`
    SslL4RenegActionListName string `json:"ssl-l4-reneg-action-list-name"`
    SslL4RenegAction string `json:"ssl-l4-reneg-action"`
}


type DdosZoneTemplateSslL4Src struct {
    RateLimit DdosZoneTemplateSslL4SrcRateLimit `json:"rate-limit"`
}


type DdosZoneTemplateSslL4SrcRateLimit struct {
    Request DdosZoneTemplateSslL4SrcRateLimitRequest `json:"request"`
}


type DdosZoneTemplateSslL4SrcRateLimitRequest struct {
    SrcRequestRateLimit int `json:"src-request-rate-limit"`
    SrcRequestRateLimitActionListName string `json:"src-request-rate-limit-action-list-name"`
    SrcRequestRateLimitAction string `json:"src-request-rate-limit-action"`
}


type DdosZoneTemplateSslL4SslTrafficCheck317 struct {
    HeaderInspection int `json:"header-inspection"`
    HeaderAction string `json:"header-action"`
    CheckResumedConnection int `json:"check-resumed-connection"`
    Uuid string `json:"uuid"`
}

func (p *DdosZoneTemplateSslL4) GetId() string{
    return url.QueryEscape(p.Inst.SslL4TmplName)
}

func (p *DdosZoneTemplateSslL4) getPath() string{
    return "ddos/zone-template/ssl-l4"
}

func (p *DdosZoneTemplateSslL4) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateSslL4::Post")
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

func (p *DdosZoneTemplateSslL4) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateSslL4::Get")
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
func (p *DdosZoneTemplateSslL4) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateSslL4::Put")
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

func (p *DdosZoneTemplateSslL4) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateSslL4::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
