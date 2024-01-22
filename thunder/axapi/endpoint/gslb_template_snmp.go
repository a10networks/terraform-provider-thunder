

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbTemplateSnmp struct {
	Inst struct {

    AuthKey string `json:"auth-key"`
    AuthProto string `json:"auth-proto" dval:"md5"`
    Community string `json:"community"`
    ContextEngineId string `json:"context-engine-id"`
    ContextName string `json:"context-name"`
    Host string `json:"host"`
    Interface int `json:"interface"`
    Interval int `json:"interval" dval:"3"`
    Oid string `json:"oid"`
    Port int `json:"port" dval:"161"`
    PrivKey string `json:"priv-key"`
    PrivProto string `json:"priv-proto" dval:"des"`
    SecurityEngineId string `json:"security-engine-id"`
    SecurityLevel string `json:"security-level" dval:"no-auth"`
    SnmpName string `json:"snmp-name"`
    UserTag string `json:"user-tag"`
    Username string `json:"username"`
    Uuid string `json:"uuid"`
    Version string `json:"version" dval:"v3"`

	} `json:"snmp"`
}

func (p *GslbTemplateSnmp) GetId() string{
    return p.Inst.SnmpName
}

func (p *GslbTemplateSnmp) getPath() string{
    return "gslb/template/snmp"
}

func (p *GslbTemplateSnmp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbTemplateSnmp::Post")
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

func (p *GslbTemplateSnmp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbTemplateSnmp::Get")
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
func (p *GslbTemplateSnmp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbTemplateSnmp::Put")
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

func (p *GslbTemplateSnmp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbTemplateSnmp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
