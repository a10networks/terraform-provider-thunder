

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6NatInsideSourceStatisticsOper struct {
    
    Oper Cgnv6NatInsideSourceStatisticsOperOper `json:"oper"`

}
type DataCgnv6NatInsideSourceStatisticsOper struct {
    DtCgnv6NatInsideSourceStatisticsOper Cgnv6NatInsideSourceStatisticsOper `json:"statistics"`
}


type Cgnv6NatInsideSourceStatisticsOperOper struct {
    StaticList []Cgnv6NatInsideSourceStatisticsOperOperStaticList `json:"static-list"`
    Total int `json:"total"`
}


type Cgnv6NatInsideSourceStatisticsOperOperStaticList struct {
    SrcAddress string `json:"src-address"`
    PortUsage int `json:"port-usage"`
    TotalUsed int `json:"total-used"`
    TotalFreed int `json:"total-freed"`
    NatAddress string `json:"nat-address"`
}

func (p *Cgnv6NatInsideSourceStatisticsOper) GetId() string{
    return "1"
}

func (p *Cgnv6NatInsideSourceStatisticsOper) getPath() string{
    return "cgnv6/nat/inside/source/statistics/oper"
}

func (p *Cgnv6NatInsideSourceStatisticsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6NatInsideSourceStatisticsOper,error) {
logger.Println("Cgnv6NatInsideSourceStatisticsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6NatInsideSourceStatisticsOper
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
