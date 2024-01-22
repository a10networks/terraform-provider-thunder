

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RbaGroup struct {
	Inst struct {

    Name string `json:"name"`
    PartitionList []RbaGroupPartitionList `json:"partition-list"`
    UserList []RbaGroupUserList `json:"user-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"group"`
}


type RbaGroupPartitionList struct {
    PartitionName string `json:"partition-name"`
    RoleList []RbaGroupPartitionListRoleList `json:"role-list"`
    RuleList []RbaGroupPartitionListRuleList `json:"rule-list"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type RbaGroupPartitionListRoleList struct {
    Role string `json:"role"`
}


type RbaGroupPartitionListRuleList struct {
    Object string `json:"object"`
    Operation string `json:"operation"`
}


type RbaGroupUserList struct {
    User string `json:"user"`
}

func (p *RbaGroup) GetId() string{
    return p.Inst.Name
}

func (p *RbaGroup) getPath() string{
    return "rba/group"
}

func (p *RbaGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RbaGroup::Post")
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

func (p *RbaGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RbaGroup::Get")
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
func (p *RbaGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RbaGroup::Put")
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

func (p *RbaGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RbaGroup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
