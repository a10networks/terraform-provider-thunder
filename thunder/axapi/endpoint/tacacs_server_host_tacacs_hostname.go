

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type TacacsServerHostTacacsHostname struct {
	Inst struct {

    Hostname string `json:"hostname"`
    Secret TacacsServerHostTacacsHostnameSecret `json:"secret"`
    Uuid string `json:"uuid"`

	} `json:"tacacs-hostname"`
}


type TacacsServerHostTacacsHostnameSecret struct {
    SecretValue string `json:"secret-value"`
    Encrypted string `json:"encrypted"`
    SourceIp string `json:"source-ip"`
    SourceIpv6 string `json:"source-ipv6"`
    SourceLoopback int `json:"source-loopback"`
    SourceEth int `json:"source-eth"`
    SourceVe int `json:"source-ve"`
    SourceTrunk int `json:"source-trunk"`
    SourceLif int `json:"source-lif"`
    PortCfg TacacsServerHostTacacsHostnameSecretPortCfg `json:"port-cfg"`
}


type TacacsServerHostTacacsHostnameSecretPortCfg struct {
    Port int `json:"port" dval:"49"`
    Timeout int `json:"timeout" dval:"12"`
    Monitor int `json:"monitor"`
    Username string `json:"username"`
    Password int `json:"password"`
    PasswordValue string `json:"password-value"`
    Encrypted string `json:"encrypted"`
}

func (p *TacacsServerHostTacacsHostname) GetId() string{
    return url.QueryEscape(p.Inst.Hostname)
}

func (p *TacacsServerHostTacacsHostname) getPath() string{
    return "tacacs-server/host/tacacs-hostname"
}

func (p *TacacsServerHostTacacsHostname) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TacacsServerHostTacacsHostname::Post")
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

func (p *TacacsServerHostTacacsHostname) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TacacsServerHostTacacsHostname::Get")
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
func (p *TacacsServerHostTacacsHostname) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TacacsServerHostTacacsHostname::Put")
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

func (p *TacacsServerHostTacacsHostname) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TacacsServerHostTacacsHostname::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
