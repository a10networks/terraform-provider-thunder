

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbDnsResponseRateLimitingOper struct {
    
    Oper SlbDnsResponseRateLimitingOperOper `json:"oper"`

}
type DataSlbDnsResponseRateLimitingOper struct {
    DtSlbDnsResponseRateLimitingOper SlbDnsResponseRateLimitingOper `json:"dns-response-rate-limiting"`
}


type SlbDnsResponseRateLimitingOperOper struct {
    Filter_type string `json:"filter_type"`
    Filter_address_v4 string `json:"filter_address_v4"`
    Filter_address_v6 string `json:"filter_address_v6"`
    Filter_fqdn string `json:"filter_fqdn"`
    Filter_debug int `json:"filter_debug"`
    EntryList []SlbDnsResponseRateLimitingOperOperEntryList `json:"entry-list"`
    DnsrrlCpuList []SlbDnsResponseRateLimitingOperOperDnsrrlCpuList `json:"dnsrrl-cpu-list"`
    CpuCount int `json:"cpu-count"`
}


type SlbDnsResponseRateLimitingOperOperEntryList struct {
    EntryAddress string `json:"entry-address"`
    EntryFqdn string `json:"entry-fqdn"`
    EntryHitCount int `json:"entry-hit-count"`
    EntryAge int `json:"entry-age"`
    EntryResponseCredit int `json:"entry-response-credit"`
    EntryAction string `json:"entry-action"`
    EntryOverLimit int `json:"entry-over-limit"`
}


type SlbDnsResponseRateLimitingOperOperDnsrrlCpuList struct {
    Total_created int `json:"total_created"`
    Total_inserted int `json:"total_inserted"`
    Total_withdrew int `json:"total_withdrew"`
    Total_ready_to_free int `json:"total_ready_to_free"`
    Total_freed int `json:"total_freed"`
    Total_logs int `json:"total_logs"`
    Total_overflow_entry_hits int `json:"total_overflow_entry_hits"`
    Total_refill int `json:"total_refill"`
    Total_credit_exceeded int `json:"total_credit_exceeded"`
    Other_thread_refill int `json:"other_thread_refill"`
    Err_entry_create_failed int `json:"err_entry_create_failed"`
    Err_entry_create_oom int `json:"err_entry_create_oom"`
    Err_entry_ext_create_oom int `json:"err_entry_ext_create_oom"`
    Err_entry_insert_failed int `json:"err_entry_insert_failed"`
    Err_vport_fail_match int `json:"err_vport_fail_match"`
}

func (p *SlbDnsResponseRateLimitingOper) GetId() string{
    return "1"
}

func (p *SlbDnsResponseRateLimitingOper) getPath() string{
    return "slb/dns-response-rate-limiting/oper"
}

func (p *SlbDnsResponseRateLimitingOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbDnsResponseRateLimitingOper,error) {
logger.Println("SlbDnsResponseRateLimitingOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbDnsResponseRateLimitingOper
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
