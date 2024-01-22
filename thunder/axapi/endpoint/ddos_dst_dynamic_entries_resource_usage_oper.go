

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstDynamicEntriesResourceUsageOper struct {
    
    Oper DdosDstDynamicEntriesResourceUsageOperOper `json:"oper"`

}
type DataDdosDstDynamicEntriesResourceUsageOper struct {
    DtDdosDstDynamicEntriesResourceUsageOper DdosDstDynamicEntriesResourceUsageOper `json:"dynamic-entries-resource-usage"`
}


type DdosDstDynamicEntriesResourceUsageOperOper struct {
    Dst_entry_ip_limit int `json:"dst_entry_ip_limit"`
    Dst_entry_ip_allocated int `json:"dst_entry_ip_allocated"`
    Dst_entry_ip_remaining string `json:"dst_entry_ip_remaining"`
    Conn_total_ip int `json:"conn_total_ip"`
    Dst_entry_ipv6_limit int `json:"dst_entry_ipv6_limit"`
    Dst_entry_ipv6_allocated int `json:"dst_entry_ipv6_allocated"`
    Dst_entry_ipv6_remaining string `json:"dst_entry_ipv6_remaining"`
    Conn_total_ipv6 int `json:"conn_total_ipv6"`
}

func (p *DdosDstDynamicEntriesResourceUsageOper) GetId() string{
    return "1"
}

func (p *DdosDstDynamicEntriesResourceUsageOper) getPath() string{
    return "ddos/dst/dynamic-entries-resource-usage/oper"
}

func (p *DdosDstDynamicEntriesResourceUsageOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstDynamicEntriesResourceUsageOper,error) {
logger.Println("DdosDstDynamicEntriesResourceUsageOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstDynamicEntriesResourceUsageOper
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
