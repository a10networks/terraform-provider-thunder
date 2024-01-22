

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type NetworkSpanningTreeModeMstpInstance struct {
	Inst struct {

    InstanceStart int `json:"instance-start"`
    Priority int `json:"priority" dval:"32768"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VlanList []NetworkSpanningTreeModeMstpInstanceVlanList `json:"vlan-list"`

	} `json:"instance"`
}


type NetworkSpanningTreeModeMstpInstanceVlanList struct {
    VlanStart int `json:"vlan-start"`
    VlanEnd int `json:"vlan-end"`
}

func (p *NetworkSpanningTreeModeMstpInstance) GetId() string{
    return strconv.Itoa(p.Inst.InstanceStart)
}

func (p *NetworkSpanningTreeModeMstpInstance) getPath() string{
    return "network/spanning-tree/mode/mstp/instance"
}

func (p *NetworkSpanningTreeModeMstpInstance) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkSpanningTreeModeMstpInstance::Post")
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

func (p *NetworkSpanningTreeModeMstpInstance) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkSpanningTreeModeMstpInstance::Get")
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
func (p *NetworkSpanningTreeModeMstpInstance) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkSpanningTreeModeMstpInstance::Put")
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

func (p *NetworkSpanningTreeModeMstpInstance) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkSpanningTreeModeMstpInstance::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
