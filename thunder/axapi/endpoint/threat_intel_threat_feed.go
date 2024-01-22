

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ThreatIntelThreatFeed struct {
	Inst struct {

    Domain string `json:"domain"`
    Enable int `json:"enable"`
    Encrypted string `json:"encrypted"`
    LogLevel string `json:"log-level" dval:"warning"`
    Port int `json:"port" dval:"443"`
    ProxyAuthType string `json:"proxy-auth-type" dval:"ntlm"`
    ProxyHost string `json:"proxy-host"`
    ProxyPassword int `json:"proxy-password"`
    ProxyPort int `json:"proxy-port"`
    ProxyUsername string `json:"proxy-username"`
    RtuUpdateDisable int `json:"rtu-update-disable"`
    SecretString string `json:"secret-string"`
    Server string `json:"server"`
    ServerTimeout int `json:"server-timeout" dval:"15"`
    Type string `json:"type"`
    UpdateInterval int `json:"update-interval" dval:"120"`
    UseMgmtPort int `json:"use-mgmt-port"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"threat-feed"`
}

func (p *ThreatIntelThreatFeed) GetId() string{
    return p.Inst.Type
}

func (p *ThreatIntelThreatFeed) getPath() string{
    return "threat-intel/threat-feed"
}

func (p *ThreatIntelThreatFeed) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ThreatIntelThreatFeed::Post")
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

func (p *ThreatIntelThreatFeed) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ThreatIntelThreatFeed::Get")
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
func (p *ThreatIntelThreatFeed) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ThreatIntelThreatFeed::Put")
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

func (p *ThreatIntelThreatFeed) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ThreatIntelThreatFeed::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
