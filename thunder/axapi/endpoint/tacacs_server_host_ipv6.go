

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type TacacsServerHostIpv6 struct {
	Inst struct {

    Ipv6Addr string `json:"ipv6-addr"`
    Secret TacacsServerHostIpv6Secret `json:"secret"`
    Uuid string `json:"uuid"`

	} `json:"ipv6"`
}


type TacacsServerHostIpv6Secret struct {
    SecretValue string `json:"secret-value"`
    Encrypted string `json:"encrypted"`
    SourceIpv6 string `json:"source-ipv6"`
    SourceLoopback int `json:"source-loopback"`
    SourceEth int `json:"source-eth"`
    SourceVe int `json:"source-ve"`
    SourceTrunk int `json:"source-trunk"`
    SourceLif int `json:"source-lif"`
    PortCfg TacacsServerHostIpv6SecretPortCfg `json:"port-cfg"`
}


type TacacsServerHostIpv6SecretPortCfg struct {
    Port int `json:"port" dval:"49"`
    Timeout int `json:"timeout" dval:"12"`
    Monitor int `json:"monitor"`
    Username string `json:"username"`
    Password int `json:"password"`
    PasswordValue string `json:"password-value"`
    Encrypted string `json:"encrypted"`
}

func (p *TacacsServerHostIpv6) GetId() string{
    return p.Inst.Ipv6Addr
}

func (p *TacacsServerHostIpv6) getPath() string{
    return "tacacs-server/host/ipv6"
}

func (p *TacacsServerHostIpv6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TacacsServerHostIpv6::Post")
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

func (p *TacacsServerHostIpv6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TacacsServerHostIpv6::Get")
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
func (p *TacacsServerHostIpv6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TacacsServerHostIpv6::Put")
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

func (p *TacacsServerHostIpv6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TacacsServerHostIpv6::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
