

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwFullConeSessionOper struct {
    
    Oper FwFullConeSessionOperOper `json:"oper"`

}
type DataFwFullConeSessionOper struct {
    DtFwFullConeSessionOper FwFullConeSessionOper `json:"full-cone-session"`
}


type FwFullConeSessionOperOper struct {
    SessionList []FwFullConeSessionOperOperSessionList `json:"session-list"`
    Total int `json:"total"`
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
}


type FwFullConeSessionOperOperSessionList struct {
    Protocol string `json:"protocol"`
    InsideAddress string `json:"inside-address"`
    InsidePort int `json:"inside-port"`
    Outbound int `json:"outbound"`
    Inbound int `json:"inbound"`
    Cpu int `json:"cpu"`
    Age string `json:"age"`
}

func (p *FwFullConeSessionOper) GetId() string{
    return "1"
}

func (p *FwFullConeSessionOper) getPath() string{
    return "fw/full-cone-session/oper"
}

func (p *FwFullConeSessionOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwFullConeSessionOper,error) {
logger.Println("FwFullConeSessionOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwFullConeSessionOper
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
