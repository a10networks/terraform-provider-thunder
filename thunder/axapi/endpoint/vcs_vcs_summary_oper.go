

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VcsVcsSummaryOper struct {
    
    Oper VcsVcsSummaryOperOper `json:"oper"`

}
type DataVcsVcsSummaryOper struct {
    DtVcsVcsSummaryOper VcsVcsSummaryOper `json:"vcs-summary"`
}


type VcsVcsSummaryOperOper struct {
    VcsEnabled string `json:"vcs-enabled"`
    ChassisId int `json:"chassis-id"`
    MulticastAddr string `json:"multicast-addr"`
    MulticastPort int `json:"multicast-port"`
    Version string `json:"version"`
    VmasterMaintenanceInterval int `json:"vmaster-maintenance-interval"`
    VmasterMaintenanceLeft int `json:"vmaster-maintenance-left"`
    VcsHandshakeCompletedList []VcsVcsSummaryOperOperVcsHandshakeCompletedList `json:"vcs-handshake-completed-list"`
    FloatingIpv4List []VcsVcsSummaryOperOperFloatingIpv4List `json:"floating-ipv4-list"`
    FloatingIpv6List []VcsVcsSummaryOperOperFloatingIpv6List `json:"floating-ipv6-list"`
    MemberList []VcsVcsSummaryOperOperMemberList `json:"member-list"`
}


type VcsVcsSummaryOperOperVcsHandshakeCompletedList struct {
    VcsHandshakeCompletedId int `json:"vcs-handshake-completed-id"`
    VcsHandshakeCompleted int `json:"vcs-handshake-completed"`
}


type VcsVcsSummaryOperOperFloatingIpv4List struct {
    FloatingIpv4 string `json:"floating-ipv4"`
    FloatingIpv4Mask string `json:"floating-ipv4-mask"`
}


type VcsVcsSummaryOperOperFloatingIpv6List struct {
    FloatingIpv6 string `json:"floating-ipv6"`
    FloatingIpv6Prefix int `json:"floating-ipv6-prefix"`
}


type VcsVcsSummaryOperOperMemberList struct {
    Id1 int `json:"id1"`
    State string `json:"state"`
    Priority int `json:"priority"`
    IpList []VcsVcsSummaryOperOperMemberListIpList `json:"ip-list"`
    Port int `json:"port"`
    Location string `json:"location"`
}


type VcsVcsSummaryOperOperMemberListIpList struct {
    Ip string `json:"ip"`
}

func (p *VcsVcsSummaryOper) GetId() string{
    return "1"
}

func (p *VcsVcsSummaryOper) getPath() string{
    return "vcs/vcs-summary/oper"
}

func (p *VcsVcsSummaryOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVcsVcsSummaryOper,error) {
logger.Println("VcsVcsSummaryOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVcsVcsSummaryOper
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
