

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbEcmpOper struct {
    
    Oper SlbEcmpOperOper `json:"oper"`

}
type DataSlbEcmpOper struct {
    DtSlbEcmpOper SlbEcmpOper `json:"ecmp"`
}


type SlbEcmpOperOper struct {
    TotalPath int `json:"total-path"`
    EcmpPath int `json:"ecmp-path"`
    Filter_type string `json:"filter_type"`
    SourceAddrV4 string `json:"source-addr-v4"`
    DestAddrV4 string `json:"dest-addr-v4"`
    SourceAddrV6 string `json:"source-addr-v6"`
    DestAddrV6 string `json:"dest-addr-v6"`
    DstPort int `json:"dst-port"`
    SrcPort int `json:"src-port"`
}

func (p *SlbEcmpOper) GetId() string{
    return "1"
}

func (p *SlbEcmpOper) getPath() string{
    return "slb/ecmp/oper"
}

func (p *SlbEcmpOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbEcmpOper,error) {
logger.Println("SlbEcmpOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbEcmpOper
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
