

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplatePersistSourceIp struct {
	Inst struct {

    DontHonorConnRules int `json:"dont-honor-conn-rules"`
    EnforceHigherPriority int `json:"enforce-higher-priority"`
    HashPersist int `json:"hash-persist"`
    InclDstIp int `json:"incl-dst-ip"`
    InclSport int `json:"incl-sport"`
    MatchType int `json:"match-type"`
    Name string `json:"name"`
    Netmask string `json:"netmask" dval:"255.255.255.255"`
    Netmask6 int `json:"netmask6" dval:"128"`
    PrimaryPort int `json:"primary-port"`
    ScanAllMembers int `json:"scan-all-members"`
    Server int `json:"server"`
    ServiceGroup int `json:"service-group"`
    Timeout int `json:"timeout" dval:"5"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"source-ip"`
}

func (p *SlbTemplatePersistSourceIp) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplatePersistSourceIp) getPath() string{
    return "slb/template/persist/source-ip"
}

func (p *SlbTemplatePersistSourceIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePersistSourceIp::Post")
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

func (p *SlbTemplatePersistSourceIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePersistSourceIp::Get")
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
func (p *SlbTemplatePersistSourceIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePersistSourceIp::Put")
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

func (p *SlbTemplatePersistSourceIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePersistSourceIp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
