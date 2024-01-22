

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6ResourceUsageOper struct {
    
    Oper Cgnv6ResourceUsageOperOper `json:"oper"`

}
type DataCgnv6ResourceUsageOper struct {
    DtCgnv6ResourceUsageOper Cgnv6ResourceUsageOper `json:"resource-usage"`
}


type Cgnv6ResourceUsageOperOper struct {
    LsnNatAddrCountMin int `json:"lsn-nat-addr-count-min"`
    LsnNatAddrCountMax int `json:"lsn-nat-addr-count-max"`
    LsnNatAddrCountDefault int `json:"lsn-nat-addr-count-default"`
    FixedNatIpAddrCountMin int `json:"fixed-nat-ip-addr-count-min"`
    FixedNatIpAddrCountMax int `json:"fixed-nat-ip-addr-count-max"`
    FixedNatIpAddrCountDefault int `json:"fixed-nat-ip-addr-count-default"`
    FixedNatInsideUserCountMin int `json:"fixed-nat-inside-user-count-min"`
    FixedNatInsideUserCountMax int `json:"fixed-nat-inside-user-count-max"`
    FixedNatInsideUserCountDefault int `json:"fixed-nat-inside-user-count-default"`
    RadiusTableSizeMin int `json:"radius-table-size-min"`
    RadiusTableSizeMax int `json:"radius-table-size-max"`
    RadiusTableSizeDefault int `json:"radius-table-size-default"`
}

func (p *Cgnv6ResourceUsageOper) GetId() string{
    return "1"
}

func (p *Cgnv6ResourceUsageOper) getPath() string{
    return "cgnv6/resource-usage/oper"
}

func (p *Cgnv6ResourceUsageOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6ResourceUsageOper,error) {
logger.Println("Cgnv6ResourceUsageOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6ResourceUsageOper
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
