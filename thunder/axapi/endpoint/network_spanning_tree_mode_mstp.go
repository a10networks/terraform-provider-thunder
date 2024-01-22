

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetworkSpanningTreeModeMstp struct {
	Inst struct {

    Action int `json:"action"`
    InstanceList []NetworkSpanningTreeModeMstpInstanceList `json:"instance-list"`
    Priority int `json:"priority" dval:"32768"`
    RevisionCfg NetworkSpanningTreeModeMstpRevisionCfg `json:"revision-cfg"`
    Uuid string `json:"uuid"`

	} `json:"mstp"`
}


type NetworkSpanningTreeModeMstpInstanceList struct {
    InstanceStart int `json:"instance-start"`
    VlanList []NetworkSpanningTreeModeMstpInstanceListVlanList `json:"vlan-list"`
    Priority int `json:"priority" dval:"32768"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type NetworkSpanningTreeModeMstpInstanceListVlanList struct {
    VlanStart int `json:"vlan-start"`
    VlanEnd int `json:"vlan-end"`
}


type NetworkSpanningTreeModeMstpRevisionCfg struct {
    Revision int `json:"revision"`
    Name string `json:"name"`
}

func (p *NetworkSpanningTreeModeMstp) GetId() string{
    return "1"
}

func (p *NetworkSpanningTreeModeMstp) getPath() string{
    return "network/spanning-tree/mode/mstp"
}

func (p *NetworkSpanningTreeModeMstp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkSpanningTreeModeMstp::Post")
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

func (p *NetworkSpanningTreeModeMstp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkSpanningTreeModeMstp::Get")
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
func (p *NetworkSpanningTreeModeMstp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkSpanningTreeModeMstp::Put")
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

func (p *NetworkSpanningTreeModeMstp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkSpanningTreeModeMstp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
