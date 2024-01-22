

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetflowCommon struct {
	Inst struct {

    MaxPacketQueueTime int `json:"max-packet-queue-time" dval:"50"`
    Nat44_tpl_1001 int `json:"nat44_tpl_1001"`
    ResetTimeOnFlowRecord int `json:"reset-time-on-flow-record"`
    SelectorAlgorithmList []NetflowCommonSelectorAlgorithmList `json:"selector-algorithm-list"`
    Uuid string `json:"uuid"`

	} `json:"common"`
}


type NetflowCommonSelectorAlgorithmList struct {
    AlgorithmName string `json:"algorithm-name"`
    SamplingPopulation int `json:"sampling-population" dval:"1"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}

func (p *NetflowCommon) GetId() string{
    return "1"
}

func (p *NetflowCommon) getPath() string{
    return "netflow/common"
}

func (p *NetflowCommon) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowCommon::Post")
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

func (p *NetflowCommon) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowCommon::Get")
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
func (p *NetflowCommon) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowCommon::Put")
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

func (p *NetflowCommon) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowCommon::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
