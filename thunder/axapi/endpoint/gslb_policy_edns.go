

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbPolicyEdns struct {
	Inst struct {

    ClientSubnetGeographic int `json:"client-subnet-geographic"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"edns"`
}

func (p *GslbPolicyEdns) GetId() string{
    return "1"
}

func (p *GslbPolicyEdns) getPath() string{
    return "gslb/policy/" +p.Inst.Name + "/edns"
}

func (p *GslbPolicyEdns) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyEdns::Post")
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

func (p *GslbPolicyEdns) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyEdns::Get")
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
func (p *GslbPolicyEdns) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyEdns::Put")
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

func (p *GslbPolicyEdns) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyEdns::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
