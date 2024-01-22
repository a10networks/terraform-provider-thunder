

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbStatisticsOper struct {
    
    Oper GslbStatisticsOperOper `json:"oper"`

}
type DataGslbStatisticsOper struct {
    DtGslbStatisticsOper GslbStatisticsOper `json:"statistics"`
}


type GslbStatisticsOperOper struct {
    CurrSslCtx int `json:"curr-ssl-ctx"`
}

func (p *GslbStatisticsOper) GetId() string{
    return "1"
}

func (p *GslbStatisticsOper) getPath() string{
    return "gslb/statistics/oper"
}

func (p *GslbStatisticsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbStatisticsOper,error) {
logger.Println("GslbStatisticsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbStatisticsOper
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
