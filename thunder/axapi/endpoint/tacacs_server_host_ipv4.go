

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type TacacsServerHostIpv4 struct {
	Inst struct {

    Ipv4Addr string `json:"ipv4-addr"`
    Secret TacacsServerHostIpv4Secret `json:"secret"`
    Uuid string `json:"uuid"`

	} `json:"ipv4"`
}


type TacacsServerHostIpv4Secret struct {
    SecretValue string `json:"secret-value"`
    Encrypted string `json:"encrypted"`
    SourceIp string `json:"source-ip"`
    SourceLoopback int `json:"source-loopback"`
    SourceEth int `json:"source-eth"`
    SourceVe int `json:"source-ve"`
    SourceTrunk int `json:"source-trunk"`
    SourceLif int `json:"source-lif"`
    PortCfg TacacsServerHostIpv4SecretPortCfg `json:"port-cfg"`
}


type TacacsServerHostIpv4SecretPortCfg struct {
    Port int `json:"port" dval:"49"`
    Timeout int `json:"timeout" dval:"12"`
    Monitor int `json:"monitor"`
    Username string `json:"username"`
    Password int `json:"password"`
    PasswordValue string `json:"password-value"`
    Encrypted string `json:"encrypted"`
}

func (p *TacacsServerHostIpv4) GetId() string{
    return p.Inst.Ipv4Addr
}

func (p *TacacsServerHostIpv4) getPath() string{
    return "tacacs-server/host/ipv4"
}

func (p *TacacsServerHostIpv4) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TacacsServerHostIpv4::Post")
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

func (p *TacacsServerHostIpv4) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TacacsServerHostIpv4::Get")
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
func (p *TacacsServerHostIpv4) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TacacsServerHostIpv4::Put")
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

func (p *TacacsServerHostIpv4) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TacacsServerHostIpv4::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
