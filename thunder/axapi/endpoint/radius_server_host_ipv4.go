

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RadiusServerHostIpv4 struct {
	Inst struct {

    Ipv4Addr string `json:"ipv4-addr"`
    Secret RadiusServerHostIpv4Secret `json:"secret"`
    Uuid string `json:"uuid"`

	} `json:"ipv4"`
}


type RadiusServerHostIpv4Secret struct {
    SecretValue string `json:"secret-value"`
    Encrypted string `json:"encrypted"`
    PortCfg RadiusServerHostIpv4SecretPortCfg `json:"port-cfg"`
}


type RadiusServerHostIpv4SecretPortCfg struct {
    AuthPort int `json:"auth-port"`
    AcctPort int `json:"acct-port"`
    Retransmit int `json:"retransmit" dval:"3"`
    Timeout int `json:"timeout" dval:"3"`
}

func (p *RadiusServerHostIpv4) GetId() string{
    return p.Inst.Ipv4Addr
}

func (p *RadiusServerHostIpv4) getPath() string{
    return "radius-server/host/ipv4"
}

func (p *RadiusServerHostIpv4) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RadiusServerHostIpv4::Post")
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

func (p *RadiusServerHostIpv4) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RadiusServerHostIpv4::Get")
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
func (p *RadiusServerHostIpv4) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RadiusServerHostIpv4::Put")
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

func (p *RadiusServerHostIpv4) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RadiusServerHostIpv4::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
