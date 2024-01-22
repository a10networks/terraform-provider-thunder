

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetworkBfd struct {
	Inst struct {

    Echo int `json:"echo"`
    Enable int `json:"enable"`
    IntervalCfg NetworkBfdIntervalCfg `json:"interval-cfg"`
    Uuid string `json:"uuid"`

	} `json:"bfd"`
}


type NetworkBfdIntervalCfg struct {
    Interval int `json:"interval" dval:"800"`
    MinRx int `json:"min-rx" dval:"800"`
    Multiplier int `json:"multiplier" dval:"4"`
}

func (p *NetworkBfd) GetId() string{
    return "1"
}

func (p *NetworkBfd) getPath() string{
    return "network/bfd"
}

func (p *NetworkBfd) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkBfd::Post")
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

func (p *NetworkBfd) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkBfd::Get")
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
func (p *NetworkBfd) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkBfd::Put")
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

func (p *NetworkBfd) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkBfd::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
