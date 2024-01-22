

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceTrunkSpanningTree struct {
	Inst struct {

    AdminEdge int `json:"admin-edge"`
    AutoEdge int `json:"auto-edge" dval:"1"`
    InstanceList []InterfaceTrunkSpanningTreeInstanceList `json:"instance-list"`
    PathCost int `json:"path-cost"`
    Uuid string `json:"uuid"`
    Ifnum string 

	} `json:"spanning-tree"`
}


type InterfaceTrunkSpanningTreeInstanceList struct {
    InstanceStart int `json:"instance-start"`
    MstpPathCost int `json:"mstp-path-cost"`
}

func (p *InterfaceTrunkSpanningTree) GetId() string{
    return "1"
}

func (p *InterfaceTrunkSpanningTree) getPath() string{
    return "interface/trunk/" +p.Inst.Ifnum + "/spanning-tree"
}

func (p *InterfaceTrunkSpanningTree) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkSpanningTree::Post")
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

func (p *InterfaceTrunkSpanningTree) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkSpanningTree::Get")
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
func (p *InterfaceTrunkSpanningTree) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkSpanningTree::Put")
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

func (p *InterfaceTrunkSpanningTree) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkSpanningTree::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
