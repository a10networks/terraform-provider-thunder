

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityZbarTruplesOper struct {
    
    Oper VisibilityZbarTruplesOperOper `json:"oper"`

}
type DataVisibilityZbarTruplesOper struct {
    DtVisibilityZbarTruplesOper VisibilityZbarTruplesOper `json:"truples"`
}


type VisibilityZbarTruplesOperOper struct {
    SrcIpv4Addr string `json:"src-ipv4-addr"`
    SrcIpv6Addr string `json:"src-ipv6-addr"`
    Phase string `json:"phase"`
    DestIpv4Addr string `json:"dest-ipv4-addr"`
    DestIpv6Addr string `json:"dest-ipv6-addr"`
    Port int `json:"port"`
    Protocol string `json:"protocol"`
    ZbarMultiIndList []VisibilityZbarTruplesOperOperZbarMultiIndList `json:"zbar-multi-ind-list"`
}


type VisibilityZbarTruplesOperOperZbarMultiIndList struct {
    IndName string `json:"ind-name"`
    IndTotalCount int `json:"ind-total-count"`
    IndBandId int `json:"ind-band-id"`
    ZbarIndSecList []VisibilityZbarTruplesOperOperZbarMultiIndListZbarIndSecList `json:"zbar-ind-sec-list"`
}


type VisibilityZbarTruplesOperOperZbarMultiIndListZbarIndSecList struct {
    Second string `json:"second"`
    Value int `json:"value"`
}

func (p *VisibilityZbarTruplesOper) GetId() string{
    return "1"
}

func (p *VisibilityZbarTruplesOper) getPath() string{
    return "visibility/zbar/truples/oper"
}

func (p *VisibilityZbarTruplesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityZbarTruplesOper,error) {
logger.Println("VisibilityZbarTruplesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityZbarTruplesOper
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
