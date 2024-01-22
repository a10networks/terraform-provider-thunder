

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RadiusServerHostIpv6 struct {
	Inst struct {

    Ipv6Addr string `json:"ipv6-addr"`
    Secret RadiusServerHostIpv6Secret `json:"secret"`
    Uuid string `json:"uuid"`

	} `json:"ipv6"`
}


type RadiusServerHostIpv6Secret struct {
    SecretValue string `json:"secret-value"`
    Encrypted string `json:"encrypted"`
    PortCfg RadiusServerHostIpv6SecretPortCfg `json:"port-cfg"`
}


type RadiusServerHostIpv6SecretPortCfg struct {
    AuthPort int `json:"auth-port"`
    AcctPort int `json:"acct-port"`
    Retransmit int `json:"retransmit" dval:"3"`
    Timeout int `json:"timeout" dval:"3"`
}

func (p *RadiusServerHostIpv6) GetId() string{
    return p.Inst.Ipv6Addr
}

func (p *RadiusServerHostIpv6) getPath() string{
    return "radius-server/host/ipv6"
}

func (p *RadiusServerHostIpv6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RadiusServerHostIpv6::Post")
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

func (p *RadiusServerHostIpv6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RadiusServerHostIpv6::Get")
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
func (p *RadiusServerHostIpv6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RadiusServerHostIpv6::Put")
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

func (p *RadiusServerHostIpv6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RadiusServerHostIpv6::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
