

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SnmpServerSnmpv3User struct {
	Inst struct {

    AuthVal string `json:"auth-val"`
    Encpasswd string `json:"encpasswd"`
    Group string `json:"group"`
    Passwd string `json:"passwd"`
    Priv string `json:"priv"`
    PrivPwEncrypted string `json:"priv-pw-encrypted"`
    PwEncrypted string `json:"pw-encrypted"`
    Username string `json:"username"`
    Uuid string `json:"uuid"`
    V3 string `json:"v3"`

	} `json:"user"`
}

func (p *SnmpServerSnmpv3User) GetId() string{
    return url.QueryEscape(p.Inst.Username)
}

func (p *SnmpServerSnmpv3User) getPath() string{
    return "snmp-server/SNMPv3/user"
}

func (p *SnmpServerSnmpv3User) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerSnmpv3User::Post")
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

func (p *SnmpServerSnmpv3User) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerSnmpv3User::Get")
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
func (p *SnmpServerSnmpv3User) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerSnmpv3User::Put")
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

func (p *SnmpServerSnmpv3User) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerSnmpv3User::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
