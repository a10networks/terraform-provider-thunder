

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplatePersistDestinationIp struct {
	Inst struct {

    DontHonorConnRules int `json:"dont-honor-conn-rules"`
    HashPersist int `json:"hash-persist"`
    MatchType int `json:"match-type"`
    Name string `json:"name"`
    Netmask string `json:"netmask" dval:"255.255.255.255"`
    Netmask6 int `json:"netmask6" dval:"128"`
    ScanAllMembers int `json:"scan-all-members"`
    Server int `json:"server"`
    ServiceGroup int `json:"service-group"`
    Timeout int `json:"timeout" dval:"5"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"destination-ip"`
}

func (p *SlbTemplatePersistDestinationIp) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplatePersistDestinationIp) getPath() string{
    return "slb/template/persist/destination-ip"
}

func (p *SlbTemplatePersistDestinationIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePersistDestinationIp::Post")
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

func (p *SlbTemplatePersistDestinationIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePersistDestinationIp::Get")
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
func (p *SlbTemplatePersistDestinationIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePersistDestinationIp::Put")
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

func (p *SlbTemplatePersistDestinationIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePersistDestinationIp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
