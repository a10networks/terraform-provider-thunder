

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateDnsClassListLid struct {
	Inst struct {

    ActionValue string `json:"action-value"`
    ConnRateLimit int `json:"conn-rate-limit"`
    Dns SlbTemplateDnsClassListLidDns `json:"dns"`
    Lidnum int `json:"lidnum"`
    Lockout int `json:"lockout"`
    Log int `json:"log"`
    LogInterval int `json:"log-interval"`
    OverLimitAction int `json:"over-limit-action"`
    Per int `json:"per"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"lid"`
}


type SlbTemplateDnsClassListLidDns struct {
    CacheAction string `json:"cache-action" dval:"cache-disable"`
    Ttl int `json:"ttl"`
    Weight int `json:"weight"`
    HonorServerResponseTtl int `json:"honor-server-response-ttl"`
}

func (p *SlbTemplateDnsClassListLid) GetId() string{
    return strconv.Itoa(p.Inst.Lidnum)
}

func (p *SlbTemplateDnsClassListLid) getPath() string{
    return "slb/template/dns/" +p.Inst.Name + "/class-list/lid"
}

func (p *SlbTemplateDnsClassListLid) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsClassListLid::Post")
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

func (p *SlbTemplateDnsClassListLid) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsClassListLid::Get")
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
func (p *SlbTemplateDnsClassListLid) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsClassListLid::Put")
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

func (p *SlbTemplateDnsClassListLid) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsClassListLid::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
