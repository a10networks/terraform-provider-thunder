

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AccessListSummaryOper struct {
    
    Oper AccessListSummaryOperOper `json:"oper"`

}
type DataAccessListSummaryOper struct {
    DtAccessListSummaryOper AccessListSummaryOper `json:"summary"`
}


type AccessListSummaryOperOper struct {
    AclList []AccessListSummaryOperOperAclList `json:"acl-list"`
}


type AccessListSummaryOperOperAclList struct {
    Id1 int `json:"id1"`
    Name string `json:"name"`
    MgmtPktHitCount int `json:"mgmt-pkt-hit-count"`
    V6flag int `json:"v6flag"`
    RuleList []AccessListSummaryOperOperAclListRuleList `json:"rule-list"`
}


type AccessListSummaryOperOperAclListRuleList struct {
    SequenceNum int `json:"sequence-num"`
    Action string `json:"action"`
    Remark string `json:"remark"`
    Proto string `json:"proto"`
    IcmpType int `json:"icmp-type"`
    IcmpCode int `json:"icmp-code"`
    SvcObjId string `json:"svc-obj-id"`
    GeoLocationName string `json:"geo-location-name"`
    SrcObjId string `json:"src-obj-id"`
    SrcHost string `json:"src-host"`
    SrcHostMask string `json:"src-host-mask"`
    SrcPortStart int `json:"src-port-start"`
    SrcPortEnd int `json:"src-port-end"`
    DstObjId string `json:"dst-obj-id"`
    DstHost string `json:"dst-host"`
    DstHostMask string `json:"dst-host-mask"`
    DstPortStart int `json:"dst-port-start"`
    DstPortEnd int `json:"dst-port-end"`
    Eth int `json:"eth"`
    Trunk int `json:"trunk"`
    VlanId int `json:"vlan-id"`
    TcpEstablished int `json:"tcp-established"`
    Dscp int `json:"dscp"`
    IpFrag int `json:"ip-frag"`
    Log int `json:"log"`
    LogTransparentSessOnly int `json:"log-transparent-sess-only"`
    DataPlaneHits int `json:"data-plane-hits"`
}

func (p *AccessListSummaryOper) GetId() string{
    return "1"
}

func (p *AccessListSummaryOper) getPath() string{
    return "access-list/summary/oper"
}

func (p *AccessListSummaryOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAccessListSummaryOper,error) {
logger.Println("AccessListSummaryOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAccessListSummaryOper
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
