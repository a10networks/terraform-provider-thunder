

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbPolicyActiveRdt struct {
	Inst struct {

    Controller int `json:"controller"`
    Difference int `json:"difference"`
    Enable int `json:"enable"`
    FailBreak int `json:"fail-break"`
    IgnoreId int `json:"ignore-id"`
    KeepTracking int `json:"keep-tracking"`
    Limit int `json:"limit" dval:"16383"`
    ProtoRdtEnable int `json:"proto-rdt-enable"`
    Samples int `json:"samples" dval:"5"`
    SingleShot int `json:"single-shot"`
    Skip int `json:"skip" dval:"3"`
    Timeout int `json:"timeout" dval:"3"`
    Tolerance int `json:"tolerance" dval:"10"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"active-rdt"`
}

func (p *GslbPolicyActiveRdt) GetId() string{
    return "1"
}

func (p *GslbPolicyActiveRdt) getPath() string{
    return "gslb/policy/" +p.Inst.Name + "/active-rdt"
}

func (p *GslbPolicyActiveRdt) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyActiveRdt::Post")
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

func (p *GslbPolicyActiveRdt) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyActiveRdt::Get")
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
func (p *GslbPolicyActiveRdt) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyActiveRdt::Put")
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

func (p *GslbPolicyActiveRdt) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyActiveRdt::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
