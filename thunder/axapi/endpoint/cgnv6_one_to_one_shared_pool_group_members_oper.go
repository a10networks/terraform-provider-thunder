

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6OneToOneSharedPoolGroupMembersOper struct {
    
    Oper Cgnv6OneToOneSharedPoolGroupMembersOperOper `json:"oper"`

}
type DataCgnv6OneToOneSharedPoolGroupMembersOper struct {
    DtCgnv6OneToOneSharedPoolGroupMembersOper Cgnv6OneToOneSharedPoolGroupMembersOper `json:"members"`
}


type Cgnv6OneToOneSharedPoolGroupMembersOperOper struct {
    MemberList []Cgnv6OneToOneSharedPoolGroupMembersOperOperMemberList `json:"member-list"`
    PoolGroupName string `json:"pool-group-name"`
}


type Cgnv6OneToOneSharedPoolGroupMembersOperOperMemberList struct {
    PoolName string `json:"pool-name"`
}

func (p *Cgnv6OneToOneSharedPoolGroupMembersOper) GetId() string{
    return "1"
}

func (p *Cgnv6OneToOneSharedPoolGroupMembersOper) getPath() string{
    return "cgnv6/one-to-one/shared-pool-group/members/oper"
}

func (p *Cgnv6OneToOneSharedPoolGroupMembersOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6OneToOneSharedPoolGroupMembersOper,error) {
logger.Println("Cgnv6OneToOneSharedPoolGroupMembersOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6OneToOneSharedPoolGroupMembersOper
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
