

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbGroupInfoOper struct {
    
    Oper GslbGroupInfoOperOper `json:"oper"`

}
type DataGslbGroupInfoOper struct {
    DtGslbGroupInfoOper GslbGroupInfoOper `json:"group-info"`
}


type GslbGroupInfoOperOper struct {
    MemberList []GslbGroupInfoOperOperMemberList `json:"member-list"`
}


type GslbGroupInfoOperOperMemberList struct {
    GroupName string `json:"group-name"`
    MemberName string `json:"member-name"`
    Is_master int `json:"is_master"`
    SysId int `json:"sys-id"`
    Priority int `json:"priority"`
    Learn int `json:"learn"`
    Passive int `json:"passive"`
    Status string `json:"status"`
    Address string `json:"address"`
    Ipv6Address string `json:"ipv6-address"`
    ConnectSuccess int `json:"connect-success"`
    ConnectFail int `json:"connect-fail"`
    OpenIn int `json:"open-in"`
    OpenOut int `json:"open-out"`
    OpenSuccess int `json:"open-success"`
    SyncSequenceNumber int `json:"sync-sequence-number"`
    UpdateIn int `json:"update-in"`
}

func (p *GslbGroupInfoOper) GetId() string{
    return "1"
}

func (p *GslbGroupInfoOper) getPath() string{
    return "gslb/group-info/oper"
}

func (p *GslbGroupInfoOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbGroupInfoOper,error) {
logger.Println("GslbGroupInfoOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbGroupInfoOper
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
