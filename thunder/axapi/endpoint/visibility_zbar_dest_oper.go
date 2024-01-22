

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityZbarDestOper struct {
    
    BadSources VisibilityZbarDestOperBadSources `json:"bad-sources"`
    Oper VisibilityZbarDestOperOper `json:"oper"`

}
type DataVisibilityZbarDestOper struct {
    DtVisibilityZbarDestOper VisibilityZbarDestOper `json:"dest"`
}


type VisibilityZbarDestOperBadSources struct {
    Oper VisibilityZbarDestOperBadSourcesOper `json:"oper"`
}


type VisibilityZbarDestOperBadSourcesOper struct {
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    Port int `json:"port"`
    Protocol string `json:"protocol"`
    MultiBadSrcList []VisibilityZbarDestOperBadSourcesOperMultiBadSrcList `json:"multi-bad-src-list"`
}


type VisibilityZbarDestOperBadSourcesOperMultiBadSrcList struct {
    SrcIp string `json:"src-ip"`
    IndValue int `json:"ind-value"`
    State string `json:"state"`
    DropCnt int `json:"drop-cnt"`
}


type VisibilityZbarDestOperOper struct {
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    Port int `json:"port"`
    Protocol string `json:"protocol"`
    Phase string `json:"phase"`
    TupleCount int `json:"tuple-count"`
    ZbarMultiIndList []VisibilityZbarDestOperOperZbarMultiIndList `json:"zbar-multi-ind-list"`
}


type VisibilityZbarDestOperOperZbarMultiIndList struct {
    IndName string `json:"ind-name"`
    IndTotalCount int `json:"ind-total-count"`
    ZbarIndMin int `json:"zbar-ind-min"`
    ZbarIndMax int `json:"zbar-ind-max"`
    ZbarIndSlotList []VisibilityZbarDestOperOperZbarMultiIndListZbarIndSlotList `json:"zbar-ind-slot-list"`
}


type VisibilityZbarDestOperOperZbarMultiIndListZbarIndSlotList struct {
    SlotId int `json:"slot-id"`
    Start int `json:"start"`
    End int `json:"end"`
    Agg int `json:"agg"`
    Avg int `json:"avg"`
    SrcCount int `json:"src-count"`
    ZbarSlotSrcList []VisibilityZbarDestOperOperZbarMultiIndListZbarIndSlotListZbarSlotSrcList `json:"zbar-slot-src-list"`
}


type VisibilityZbarDestOperOperZbarMultiIndListZbarIndSlotListZbarSlotSrcList struct {
    SrcIp string `json:"src-ip"`
    SrcIndValue int `json:"src-ind-value"`
}

func (p *VisibilityZbarDestOper) GetId() string{
    return "1"
}

func (p *VisibilityZbarDestOper) getPath() string{
    return "visibility/zbar/dest/oper"
}

func (p *VisibilityZbarDestOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityZbarDestOper,error) {
logger.Println("VisibilityZbarDestOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityZbarDestOper
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
