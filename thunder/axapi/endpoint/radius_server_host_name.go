

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type RadiusServerHostName struct {
	Inst struct {

    Hostname string `json:"hostname"`
    Secret RadiusServerHostNameSecret `json:"secret"`
    Uuid string `json:"uuid"`

	} `json:"name"`
}


type RadiusServerHostNameSecret struct {
    SecretValue string `json:"secret-value"`
    Encrypted string `json:"encrypted"`
    PortCfg RadiusServerHostNameSecretPortCfg `json:"port-cfg"`
}


type RadiusServerHostNameSecretPortCfg struct {
    AuthPort int `json:"auth-port"`
    AcctPort int `json:"acct-port"`
    Retransmit int `json:"retransmit" dval:"3"`
    Timeout int `json:"timeout" dval:"3"`
}

func (p *RadiusServerHostName) GetId() string{
    return url.QueryEscape(p.Inst.Hostname)
}

func (p *RadiusServerHostName) getPath() string{
    return "radius-server/host/name"
}

func (p *RadiusServerHostName) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RadiusServerHostName::Post")
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

func (p *RadiusServerHostName) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RadiusServerHostName::Get")
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
func (p *RadiusServerHostName) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RadiusServerHostName::Put")
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

func (p *RadiusServerHostName) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RadiusServerHostName::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
