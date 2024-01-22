

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbPolicyCapacity struct {
	Inst struct {

    CapacityEnable int `json:"capacity-enable"`
    CapacityFailBreak int `json:"capacity-fail-break"`
    Threshold int `json:"threshold" dval:"90"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"capacity"`
}

func (p *GslbPolicyCapacity) GetId() string{
    return "1"
}

func (p *GslbPolicyCapacity) getPath() string{
    return "gslb/policy/" +p.Inst.Name + "/capacity"
}

func (p *GslbPolicyCapacity) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyCapacity::Post")
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

func (p *GslbPolicyCapacity) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyCapacity::Get")
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
func (p *GslbPolicyCapacity) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyCapacity::Put")
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

func (p *GslbPolicyCapacity) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyCapacity::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
