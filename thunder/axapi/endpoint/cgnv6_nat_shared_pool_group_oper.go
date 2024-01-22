

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6NatSharedPoolGroupOper struct {
    
    Members Cgnv6NatSharedPoolGroupOperMembers `json:"members"`
    Oper Cgnv6NatSharedPoolGroupOperOper `json:"oper"`

}
type DataCgnv6NatSharedPoolGroupOper struct {
    DtCgnv6NatSharedPoolGroupOper Cgnv6NatSharedPoolGroupOper `json:"shared-pool-group"`
}


type Cgnv6NatSharedPoolGroupOperMembers struct {
    Oper Cgnv6NatSharedPoolGroupOperMembersOper `json:"oper"`
}


type Cgnv6NatSharedPoolGroupOperMembersOper struct {
    MemberList []Cgnv6NatSharedPoolGroupOperMembersOperMemberList `json:"member-list"`
    PoolGroupName string `json:"pool-group-name"`
}


type Cgnv6NatSharedPoolGroupOperMembersOperMemberList struct {
    PoolName string `json:"pool-name"`
}


type Cgnv6NatSharedPoolGroupOperOper struct {
    SharedPoolGroupList []Cgnv6NatSharedPoolGroupOperOperSharedPoolGroupList `json:"shared-pool-group-list"`
}


type Cgnv6NatSharedPoolGroupOperOperSharedPoolGroupList struct {
    PoolGroupName string `json:"pool-group-name"`
    Vird int `json:"vird"`
}

func (p *Cgnv6NatSharedPoolGroupOper) GetId() string{
    return "1"
}

func (p *Cgnv6NatSharedPoolGroupOper) getPath() string{
    return "cgnv6/nat/shared-pool-group/oper"
}

func (p *Cgnv6NatSharedPoolGroupOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6NatSharedPoolGroupOper,error) {
logger.Println("Cgnv6NatSharedPoolGroupOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6NatSharedPoolGroupOper
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
