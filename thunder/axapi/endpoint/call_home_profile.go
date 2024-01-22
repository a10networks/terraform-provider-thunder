

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type CallHomeProfile struct {
	Inst struct {

    Action string `json:"action" dval:"register"`
    ExportPolicy string `json:"export-policy" dval:"restrictive"`
    Ipv4 string `json:"ipv4"`
    Ipv6 string `json:"ipv6"`
    Name string `json:"name"`
    Port int `json:"port"`
    Time int `json:"time"`
    UseMgmtPort int `json:"use-mgmt-port"`
    Uuid string `json:"uuid"`

	} `json:"profile"`
}

func (p *CallHomeProfile) GetId() string{
    return "1"
}

func (p *CallHomeProfile) getPath() string{
    return "call-home/profile"
}

func (p *CallHomeProfile) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("CallHomeProfile::Post")
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

func (p *CallHomeProfile) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("CallHomeProfile::Get")
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
func (p *CallHomeProfile) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("CallHomeProfile::Put")
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

func (p *CallHomeProfile) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("CallHomeProfile::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
