

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RbaUserPartition struct {
	Inst struct {

    PartitionName string `json:"partition-name"`
    RoleList []RbaUserPartitionRoleList `json:"role-list"`
    RuleList []RbaUserPartitionRuleList `json:"rule-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"partition"`
}


type RbaUserPartitionRoleList struct {
    Role string `json:"role"`
}


type RbaUserPartitionRuleList struct {
    Object string `json:"object"`
    Operation string `json:"operation"`
}

func (p *RbaUserPartition) GetId() string{
    return p.Inst.PartitionName
}

func (p *RbaUserPartition) getPath() string{
    return "rba/user/" +p.Inst.Name + "/partition"
}

func (p *RbaUserPartition) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RbaUserPartition::Post")
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

func (p *RbaUserPartition) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RbaUserPartition::Get")
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
func (p *RbaUserPartition) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RbaUserPartition::Put")
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

func (p *RbaUserPartition) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RbaUserPartition::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
