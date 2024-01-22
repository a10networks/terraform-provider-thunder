

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityZbarDestBadSourcesOper struct {
    
    Oper VisibilityZbarDestBadSourcesOperOper `json:"oper"`

}
type DataVisibilityZbarDestBadSourcesOper struct {
    DtVisibilityZbarDestBadSourcesOper VisibilityZbarDestBadSourcesOper `json:"bad-sources"`
}


type VisibilityZbarDestBadSourcesOperOper struct {
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    Port int `json:"port"`
    Protocol string `json:"protocol"`
    MultiBadSrcList []VisibilityZbarDestBadSourcesOperOperMultiBadSrcList `json:"multi-bad-src-list"`
}


type VisibilityZbarDestBadSourcesOperOperMultiBadSrcList struct {
    SrcIp string `json:"src-ip"`
    IndValue int `json:"ind-value"`
    State string `json:"state"`
    DropCnt int `json:"drop-cnt"`
}

func (p *VisibilityZbarDestBadSourcesOper) GetId() string{
    return "1"
}

func (p *VisibilityZbarDestBadSourcesOper) getPath() string{
    return "visibility/zbar/dest/bad-sources/oper"
}

func (p *VisibilityZbarDestBadSourcesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityZbarDestBadSourcesOper,error) {
logger.Println("VisibilityZbarDestBadSourcesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityZbarDestBadSourcesOper
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
