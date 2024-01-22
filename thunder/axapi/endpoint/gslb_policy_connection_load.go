

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbPolicyConnectionLoad struct {
	Inst struct {

    ConnectionLoadEnable int `json:"connection-load-enable"`
    ConnectionLoadFailBreak int `json:"connection-load-fail-break"`
    ConnectionLoadInterval int `json:"connection-load-interval" dval:"5"`
    ConnectionLoadLimit int `json:"connection-load-limit"`
    ConnectionLoadSamples int `json:"connection-load-samples" dval:"5"`
    Limit int `json:"limit"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"connection-load"`
}

func (p *GslbPolicyConnectionLoad) GetId() string{
    return "1"
}

func (p *GslbPolicyConnectionLoad) getPath() string{
    return "gslb/policy/" +p.Inst.Name + "/connection-load"
}

func (p *GslbPolicyConnectionLoad) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyConnectionLoad::Post")
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

func (p *GslbPolicyConnectionLoad) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyConnectionLoad::Get")
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
func (p *GslbPolicyConnectionLoad) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyConnectionLoad::Put")
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

func (p *GslbPolicyConnectionLoad) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyConnectionLoad::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
