

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NgWaf struct {
	Inst struct {

    Cpu NgWafCpu1076 `json:"cpu"`
    CustomPage NgWafCustomPage1077 `json:"custom-page"`
    CustomSignals NgWafCustomSignals1078 `json:"custom-signals"`
    StatsList []NgWafStatsList `json:"stats-list"`
    Status NgWafStatus1079 `json:"status"`
    Uuid string `json:"uuid"`

	} `json:"ng-waf"`
}


type NgWafCpu1076 struct {
    Uuid string `json:"uuid"`
}


type NgWafCustomPage1077 struct {
    Uuid string `json:"uuid"`
}


type NgWafCustomSignals1078 struct {
    Uuid string `json:"uuid"`
}


type NgWafStatsList struct {
    Name string `json:"name"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type NgWafStatus1079 struct {
    Uuid string `json:"uuid"`
}

func (p *NgWaf) GetId() string{
    return "1"
}

func (p *NgWaf) getPath() string{
    return "ng-waf"
}

func (p *NgWaf) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NgWaf::Post")
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

func (p *NgWaf) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NgWaf::Get")
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
func (p *NgWaf) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NgWaf::Put")
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

func (p *NgWaf) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NgWaf::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
