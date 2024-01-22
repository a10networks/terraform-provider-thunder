

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type TacacsServer struct {
	Inst struct {

    Host TacacsServerHost1896 `json:"host"`
    Interval int `json:"interval" dval:"60"`
    Monitor int `json:"monitor"`
    Uuid string `json:"uuid"`

	} `json:"tacacs-server"`
}


type TacacsServerHost1896 struct {
    Ipv4List []TacacsServerHostIpv4List `json:"ipv4-list"`
    Ipv6List []TacacsServerHostIpv6List `json:"ipv6-list"`
    TacacsHostnameList []TacacsServerHostTacacsHostnameList `json:"tacacs-hostname-list"`
}


type TacacsServerHostIpv4List struct {
    Ipv4Addr string `json:"ipv4-addr"`
    Secret TacacsServerHostIpv4ListSecret `json:"secret"`
    Uuid string `json:"uuid"`
}


type TacacsServerHostIpv4ListSecret struct {
    SecretValue string `json:"secret-value"`
    Encrypted string `json:"encrypted"`
    SourceIp string `json:"source-ip"`
    SourceLoopback int `json:"source-loopback"`
    SourceEth int `json:"source-eth"`
    SourceVe int `json:"source-ve"`
    SourceTrunk int `json:"source-trunk"`
    SourceLif int `json:"source-lif"`
    PortCfg TacacsServerHostIpv4ListSecretPortCfg `json:"port-cfg"`
}


type TacacsServerHostIpv4ListSecretPortCfg struct {
    Port int `json:"port" dval:"49"`
    Timeout int `json:"timeout" dval:"12"`
    Monitor int `json:"monitor"`
    Username string `json:"username"`
    Password int `json:"password"`
    PasswordValue string `json:"password-value"`
    Encrypted string `json:"encrypted"`
}


type TacacsServerHostIpv6List struct {
    Ipv6Addr string `json:"ipv6-addr"`
    Secret TacacsServerHostIpv6ListSecret `json:"secret"`
    Uuid string `json:"uuid"`
}


type TacacsServerHostIpv6ListSecret struct {
    SecretValue string `json:"secret-value"`
    Encrypted string `json:"encrypted"`
    SourceIpv6 string `json:"source-ipv6"`
    SourceLoopback int `json:"source-loopback"`
    SourceEth int `json:"source-eth"`
    SourceVe int `json:"source-ve"`
    SourceTrunk int `json:"source-trunk"`
    SourceLif int `json:"source-lif"`
    PortCfg TacacsServerHostIpv6ListSecretPortCfg `json:"port-cfg"`
}


type TacacsServerHostIpv6ListSecretPortCfg struct {
    Port int `json:"port" dval:"49"`
    Timeout int `json:"timeout" dval:"12"`
    Monitor int `json:"monitor"`
    Username string `json:"username"`
    Password int `json:"password"`
    PasswordValue string `json:"password-value"`
    Encrypted string `json:"encrypted"`
}


type TacacsServerHostTacacsHostnameList struct {
    Hostname string `json:"hostname"`
    Secret TacacsServerHostTacacsHostnameListSecret `json:"secret"`
    Uuid string `json:"uuid"`
}


type TacacsServerHostTacacsHostnameListSecret struct {
    SecretValue string `json:"secret-value"`
    Encrypted string `json:"encrypted"`
    SourceIp string `json:"source-ip"`
    SourceIpv6 string `json:"source-ipv6"`
    SourceLoopback int `json:"source-loopback"`
    SourceEth int `json:"source-eth"`
    SourceVe int `json:"source-ve"`
    SourceTrunk int `json:"source-trunk"`
    SourceLif int `json:"source-lif"`
    PortCfg TacacsServerHostTacacsHostnameListSecretPortCfg `json:"port-cfg"`
}


type TacacsServerHostTacacsHostnameListSecretPortCfg struct {
    Port int `json:"port" dval:"49"`
    Timeout int `json:"timeout" dval:"12"`
    Monitor int `json:"monitor"`
    Username string `json:"username"`
    Password int `json:"password"`
    PasswordValue string `json:"password-value"`
    Encrypted string `json:"encrypted"`
}

func (p *TacacsServer) GetId() string{
    return "1"
}

func (p *TacacsServer) getPath() string{
    return "tacacs-server"
}

func (p *TacacsServer) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TacacsServer::Post")
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

func (p *TacacsServer) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TacacsServer::Get")
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
func (p *TacacsServer) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TacacsServer::Put")
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

func (p *TacacsServer) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TacacsServer::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
