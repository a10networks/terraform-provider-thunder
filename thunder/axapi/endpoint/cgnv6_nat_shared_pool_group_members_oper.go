

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6NatSharedPoolGroupMembersOper struct {
    
    Oper Cgnv6NatSharedPoolGroupMembersOperOper `json:"oper"`

}
type DataCgnv6NatSharedPoolGroupMembersOper struct {
    DtCgnv6NatSharedPoolGroupMembersOper Cgnv6NatSharedPoolGroupMembersOper `json:"members"`
}


type Cgnv6NatSharedPoolGroupMembersOperOper struct {
    MemberList []Cgnv6NatSharedPoolGroupMembersOperOperMemberList `json:"member-list"`
    PoolGroupName string `json:"pool-group-name"`
}


type Cgnv6NatSharedPoolGroupMembersOperOperMemberList struct {
    PoolName string `json:"pool-name"`
}

func (p *Cgnv6NatSharedPoolGroupMembersOper) GetId() string{
    return "1"
}

func (p *Cgnv6NatSharedPoolGroupMembersOper) getPath() string{
    return "cgnv6/nat/shared-pool-group/members/oper"
}

func (p *Cgnv6NatSharedPoolGroupMembersOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6NatSharedPoolGroupMembersOper,error) {
logger.Println("Cgnv6NatSharedPoolGroupMembersOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6NatSharedPoolGroupMembersOper
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
