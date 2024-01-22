

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6OneToOneSharedPoolGroupOper struct {
    
    Members Cgnv6OneToOneSharedPoolGroupOperMembers `json:"members"`
    Oper Cgnv6OneToOneSharedPoolGroupOperOper `json:"oper"`

}
type DataCgnv6OneToOneSharedPoolGroupOper struct {
    DtCgnv6OneToOneSharedPoolGroupOper Cgnv6OneToOneSharedPoolGroupOper `json:"shared-pool-group"`
}


type Cgnv6OneToOneSharedPoolGroupOperMembers struct {
    Oper Cgnv6OneToOneSharedPoolGroupOperMembersOper `json:"oper"`
}


type Cgnv6OneToOneSharedPoolGroupOperMembersOper struct {
    MemberList []Cgnv6OneToOneSharedPoolGroupOperMembersOperMemberList `json:"member-list"`
    PoolGroupName string `json:"pool-group-name"`
}


type Cgnv6OneToOneSharedPoolGroupOperMembersOperMemberList struct {
    PoolName string `json:"pool-name"`
}


type Cgnv6OneToOneSharedPoolGroupOperOper struct {
    SharedPoolGroupList []Cgnv6OneToOneSharedPoolGroupOperOperSharedPoolGroupList `json:"shared-pool-group-list"`
}


type Cgnv6OneToOneSharedPoolGroupOperOperSharedPoolGroupList struct {
    PoolGroupName string `json:"pool-group-name"`
    Vird int `json:"vird"`
}

func (p *Cgnv6OneToOneSharedPoolGroupOper) GetId() string{
    return "1"
}

func (p *Cgnv6OneToOneSharedPoolGroupOper) getPath() string{
    return "cgnv6/one-to-one/shared-pool-group/oper"
}

func (p *Cgnv6OneToOneSharedPoolGroupOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6OneToOneSharedPoolGroupOper,error) {
logger.Println("Cgnv6OneToOneSharedPoolGroupOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6OneToOneSharedPoolGroupOper
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
