

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpFibSummaryOper struct {
    
    Oper IpFibSummaryOperOper `json:"oper"`

}
type DataIpFibSummaryOper struct {
    DtIpFibSummaryOper IpFibSummaryOper `json:"fib-summary"`
}


type IpFibSummaryOperOper struct {
    Connected_routes int `json:"connected_routes"`
    Static_dynamic_routes int `json:"static_dynamic_routes"`
    Total_routes int `json:"total_routes"`
    Static_dynamic_paths int `json:"static_dynamic_paths"`
    Total_paths int `json:"total_paths"`
}

func (p *IpFibSummaryOper) GetId() string{
    return "1"
}

func (p *IpFibSummaryOper) getPath() string{
    return "ip/fib-summary/oper"
}

func (p *IpFibSummaryOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataIpFibSummaryOper,error) {
logger.Println("IpFibSummaryOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataIpFibSummaryOper
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
