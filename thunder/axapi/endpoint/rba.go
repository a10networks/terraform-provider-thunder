

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Rba struct {
	Inst struct {

    Action string `json:"action" dval:"enable"`
    GroupList []RbaGroupList `json:"group-list"`
    RoleList []RbaRoleList `json:"role-list"`
    UserList []RbaUserList `json:"user-list"`
    Uuid string `json:"uuid"`

	} `json:"rba"`
}


type RbaGroupList struct {
    Name string `json:"name"`
    UserList []RbaGroupListUserList `json:"user-list"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    PartitionList []RbaGroupListPartitionList `json:"partition-list"`
}


type RbaGroupListUserList struct {
    User string `json:"user"`
}


type RbaGroupListPartitionList struct {
    PartitionName string `json:"partition-name"`
    RoleList []RbaGroupListPartitionListRoleList `json:"role-list"`
    RuleList []RbaGroupListPartitionListRuleList `json:"rule-list"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type RbaGroupListPartitionListRoleList struct {
    Role string `json:"role"`
}


type RbaGroupListPartitionListRuleList struct {
    Object string `json:"object"`
    Operation string `json:"operation"`
}


type RbaRoleList struct {
    Name string `json:"name"`
    DefaultPrivilege string `json:"default-privilege" dval:"no-access"`
    PartitionOnly int `json:"partition-only"`
    RuleList []RbaRoleListRuleList `json:"rule-list"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type RbaRoleListRuleList struct {
    Object string `json:"object"`
    Operation string `json:"operation"`
}


type RbaUserList struct {
    Name string `json:"name"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    PartitionList []RbaUserListPartitionList `json:"partition-list"`
}


type RbaUserListPartitionList struct {
    PartitionName string `json:"partition-name"`
    RoleList []RbaUserListPartitionListRoleList `json:"role-list"`
    RuleList []RbaUserListPartitionListRuleList `json:"rule-list"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type RbaUserListPartitionListRoleList struct {
    Role string `json:"role"`
}


type RbaUserListPartitionListRuleList struct {
    Object string `json:"object"`
    Operation string `json:"operation"`
}

func (p *Rba) GetId() string{
    return "1"
}

func (p *Rba) getPath() string{
    return "rba"
}

func (p *Rba) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Rba::Post")
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

func (p *Rba) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Rba::Get")
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
func (p *Rba) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Rba::Put")
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

func (p *Rba) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Rba::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
