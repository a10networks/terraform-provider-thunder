

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosSrcDynamicEntriesResourceUsageOper struct {
    
    Oper DdosSrcDynamicEntriesResourceUsageOperOper `json:"oper"`

}
type DataDdosSrcDynamicEntriesResourceUsageOper struct {
    DtDdosSrcDynamicEntriesResourceUsageOper DdosSrcDynamicEntriesResourceUsageOper `json:"dynamic-entries-resource-usage"`
}


type DdosSrcDynamicEntriesResourceUsageOperOper struct {
    Src_entry_ip_limit int `json:"src_entry_ip_limit"`
    Src_entry_ip_allocated int `json:"src_entry_ip_allocated"`
    Src_entry_ip_remaining string `json:"src_entry_ip_remaining"`
    Conn_total_ip int `json:"conn_total_ip"`
    Src_entry_ipv6_limit int `json:"src_entry_ipv6_limit"`
    Src_entry_ipv6_allocated int `json:"src_entry_ipv6_allocated"`
    Src_entry_ipv6_remaining string `json:"src_entry_ipv6_remaining"`
    Conn_total_ipv6 int `json:"conn_total_ipv6"`
}

func (p *DdosSrcDynamicEntriesResourceUsageOper) GetId() string{
    return "1"
}

func (p *DdosSrcDynamicEntriesResourceUsageOper) getPath() string{
    return "ddos/src/dynamic-entries-resource-usage/oper"
}

func (p *DdosSrcDynamicEntriesResourceUsageOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosSrcDynamicEntriesResourceUsageOper,error) {
logger.Println("DdosSrcDynamicEntriesResourceUsageOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosSrcDynamicEntriesResourceUsageOper
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
