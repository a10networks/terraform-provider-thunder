

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VpnGroupListOper struct {
    
    Oper VpnGroupListOperOper `json:"oper"`

}
type DataVpnGroupListOper struct {
    DtVpnGroupListOper VpnGroupListOper `json:"group-list"`
}


type VpnGroupListOperOper struct {
    GroupName string `json:"group-name"`
    GroupList []VpnGroupListOperOperGroupList `json:"group-list"`
}


type VpnGroupListOperOperGroupList struct {
    Name string `json:"Name"`
    IpsecSaName string `json:"Ipsec-sa-name"`
    IkeGatewayName string `json:"Ike-gateway-name"`
    Priority int `json:"Priority"`
    Status string `json:"Status"`
    Role string `json:"Role"`
    IsNewGroup int `json:"Is-new-group"`
    GrpMemberCount int `json:"Grp-member-count"`
}

func (p *VpnGroupListOper) GetId() string{
    return "1"
}

func (p *VpnGroupListOper) getPath() string{
    return "vpn/group-list/oper"
}

func (p *VpnGroupListOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVpnGroupListOper,error) {
logger.Println("VpnGroupListOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVpnGroupListOper
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
