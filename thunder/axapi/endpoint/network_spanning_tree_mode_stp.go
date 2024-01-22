

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetworkSpanningTreeModeStp struct {
	Inst struct {

    Action int `json:"action"`
    Priority int `json:"priority" dval:"32768"`
    Uuid string `json:"uuid"`
    VlanStart int `json:"vlan-start"`

	} `json:"stp"`
}

func (p *NetworkSpanningTreeModeStp) GetId() string{
    return "1"
}

func (p *NetworkSpanningTreeModeStp) getPath() string{
    return "network/spanning-tree/mode/stp"
}

func (p *NetworkSpanningTreeModeStp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkSpanningTreeModeStp::Post")
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

func (p *NetworkSpanningTreeModeStp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkSpanningTreeModeStp::Get")
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
func (p *NetworkSpanningTreeModeStp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkSpanningTreeModeStp::Put")
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

func (p *NetworkSpanningTreeModeStp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkSpanningTreeModeStp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
