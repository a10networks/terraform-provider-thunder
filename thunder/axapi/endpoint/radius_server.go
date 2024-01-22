

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RadiusServer struct {
	Inst struct {

    DefaultPrivilegeReadWrite int `json:"default-privilege-read-write"`
    Host RadiusServerHost1093 `json:"host"`
    Uuid string `json:"uuid"`

	} `json:"radius-server"`
}


type RadiusServerHost1093 struct {
    Ipv4List []RadiusServerHostIpv4List `json:"ipv4-list"`
    Ipv6List []RadiusServerHostIpv6List `json:"ipv6-list"`
    NameList []RadiusServerHostNameList `json:"name-list"`
}


type RadiusServerHostIpv4List struct {
    Ipv4Addr string `json:"ipv4-addr"`
    Secret RadiusServerHostIpv4ListSecret `json:"secret"`
    Uuid string `json:"uuid"`
}


type RadiusServerHostIpv4ListSecret struct {
    SecretValue string `json:"secret-value"`
    Encrypted string `json:"encrypted"`
    PortCfg RadiusServerHostIpv4ListSecretPortCfg `json:"port-cfg"`
}


type RadiusServerHostIpv4ListSecretPortCfg struct {
    AuthPort int `json:"auth-port"`
    AcctPort int `json:"acct-port"`
    Retransmit int `json:"retransmit" dval:"3"`
    Timeout int `json:"timeout" dval:"3"`
}


type RadiusServerHostIpv6List struct {
    Ipv6Addr string `json:"ipv6-addr"`
    Secret RadiusServerHostIpv6ListSecret `json:"secret"`
    Uuid string `json:"uuid"`
}


type RadiusServerHostIpv6ListSecret struct {
    SecretValue string `json:"secret-value"`
    Encrypted string `json:"encrypted"`
    PortCfg RadiusServerHostIpv6ListSecretPortCfg `json:"port-cfg"`
}


type RadiusServerHostIpv6ListSecretPortCfg struct {
    AuthPort int `json:"auth-port"`
    AcctPort int `json:"acct-port"`
    Retransmit int `json:"retransmit" dval:"3"`
    Timeout int `json:"timeout" dval:"3"`
}


type RadiusServerHostNameList struct {
    Hostname string `json:"hostname"`
    Secret RadiusServerHostNameListSecret `json:"secret"`
    Uuid string `json:"uuid"`
}


type RadiusServerHostNameListSecret struct {
    SecretValue string `json:"secret-value"`
    Encrypted string `json:"encrypted"`
    PortCfg RadiusServerHostNameListSecretPortCfg `json:"port-cfg"`
}


type RadiusServerHostNameListSecretPortCfg struct {
    AuthPort int `json:"auth-port"`
    AcctPort int `json:"acct-port"`
    Retransmit int `json:"retransmit" dval:"3"`
    Timeout int `json:"timeout" dval:"3"`
}

func (p *RadiusServer) GetId() string{
    return "1"
}

func (p *RadiusServer) getPath() string{
    return "radius-server"
}

func (p *RadiusServer) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RadiusServer::Post")
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

func (p *RadiusServer) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RadiusServer::Get")
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
func (p *RadiusServer) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RadiusServer::Put")
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

func (p *RadiusServer) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RadiusServer::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
