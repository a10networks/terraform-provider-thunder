

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityTopnCgnv6NatPoolTopnNode struct {
	Inst struct {

    Activate string `json:"activate"`
    Uuid string `json:"uuid"`

	} `json:"cgnv6-nat-pool-topn-node"`
}

func (p *VisibilityTopnCgnv6NatPoolTopnNode) GetId() string{
    return "1"
}

func (p *VisibilityTopnCgnv6NatPoolTopnNode) getPath() string{
    return "visibility/topn/cgnv6-nat-pool-topn-node"
}

func (p *VisibilityTopnCgnv6NatPoolTopnNode) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityTopnCgnv6NatPoolTopnNode::Post")
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

func (p *VisibilityTopnCgnv6NatPoolTopnNode) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityTopnCgnv6NatPoolTopnNode::Get")
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
func (p *VisibilityTopnCgnv6NatPoolTopnNode) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityTopnCgnv6NatPoolTopnNode::Put")
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

func (p *VisibilityTopnCgnv6NatPoolTopnNode) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityTopnCgnv6NatPoolTopnNode::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
