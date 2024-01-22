

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbRcCacheGlobalOper struct {
    
    Oper SlbRcCacheGlobalOperOper `json:"oper"`

}
type DataSlbRcCacheGlobalOper struct {
    DtSlbRcCacheGlobalOper SlbRcCacheGlobalOper `json:"rc-cache-global"`
}


type SlbRcCacheGlobalOperOper struct {
    EntryList []SlbRcCacheGlobalOperOperEntryList `json:"entry-list"`
    Cache_hits int `json:"cache_hits"`
    Cache_miss int `json:"cache_miss"`
    Memory_used string `json:"memory_used"`
    Hit_ratio string `json:"hit_ratio"`
    Ratio304_ string `json:"ratio304_"`
    Bytes_served string `json:"bytes_served"`
    Total_request int `json:"total_request"`
    No_cache_requests int `json:"no_cache_requests"`
    Cacheable_requests int `json:"cacheable_requests"`
    Ims_requests int `json:"ims_requests"`
    Resp_server_304_not_modified int `json:"resp_server_304_not_modified"`
    Resp_server_200_ok_chunk int `json:"resp_server_200_ok_chunk"`
    Resp_server_no_cache_response int `json:"resp_server_no_cache_response"`
    Resp_server_200_ok_cont int `json:"resp_server_200_ok_cont"`
    Resp_server_other int `json:"resp_server_other"`
    Resp_cache_304_not_modified int `json:"resp_cache_304_not_modified"`
    Resp_cache_200_ok_gzip int `json:"resp_cache_200_ok_gzip"`
    Resp_cache_other int `json:"resp_cache_other"`
    Resp_cache_200_ok_no_comp int `json:"resp_cache_200_ok_no_comp"`
    Resp_cache_200_ok_deflate int `json:"resp_cache_200_ok_deflate"`
    Entries_cached int `json:"entries_cached"`
    Entries_aged int `json:"entries_aged"`
    Entries_create_fail int `json:"entries_create_fail"`
    Entries_replaced int `json:"entries_replaced"`
    Entries_cleaned int `json:"entries_cleaned"`
    Revalidation_success int `json:"revalidation_success"`
    Revalidation_failure int `json:"revalidation_failure"`
    Policy_uri_nocache int `json:"policy_uri_nocache"`
    Polic_uri_invalidate int `json:"polic_uri_invalidate"`
    Policy_content_small int `json:"policy_content_small"`
    Policy_uri_cache int `json:"policy_uri_cache"`
    Policy_content_big int `json:"policy_content_big"`
    Virtual_server string `json:"virtual_server"`
    Virtual_port int `json:"virtual_port"`
    Display_detail int `json:"display_detail"`
    Uri_name string `json:"uri_name"`
    ReplacementList []SlbRcCacheGlobalOperOperReplacementList `json:"replacement-list"`
}


type SlbRcCacheGlobalOperOperEntryList struct {
    Host string `json:"host"`
    Url string `json:"url"`
    Bytes int `json:"bytes"`
    Type string `json:"type"`
    Status string `json:"status"`
    Expires string `json:"expires"`
    Host1 string `json:"host1"`
    Url1 string `json:"url1"`
    Bytes1 string `json:"bytes1"`
    Response_hdr_len string `json:"response_hdr_len"`
    Status_code string `json:"status_code"`
    Etag string `json:"etag"`
    Cache_control string `json:"cache_control"`
    Date string `json:"date"`
    Last_modified string `json:"last_modified"`
    Time_elapsed string `json:"time_elapsed"`
    Age string `json:"age"`
    Expires1 string `json:"expires1"`
    Hits string `json:"hits"`
    Misses string `json:"misses"`
    Concurrent_readers string `json:"concurrent_readers"`
    Content_encoding string `json:"content_encoding"`
    Http_version string `json:"http_version"`
    Response_chunked_encoding string `json:"response_chunked_encoding"`
    Weak_etag string `json:"weak_etag"`
    Full_response_cache string `json:"full_response_cache"`
    Http_request_method string `json:"http_request_method"`
    Vserver_name string `json:"vserver_name"`
    Vport string `json:"vport"`
    Memory_configured string `json:"memory_configured"`
    Memory_used string `json:"memory_used"`
    Memory_used_locally string `json:"memory_used_locally"`
    Percent_used string `json:"percent_used"`
    Partition string `json:"partition"`
}


type SlbRcCacheGlobalOperOperReplacementList struct {
    One256th int `json:"one-256th"`
    One128th int `json:"one-128th"`
    One64th int `json:"one-64th"`
    One32th int `json:"one-32th"`
    One16th int `json:"one-16th"`
    One8th int `json:"one-8th"`
    One4th int `json:"one-4th"`
    One2th int `json:"one-2th"`
    One int `json:"one"`
    Two int `json:"two"`
    Four int `json:"four"`
    Eight int `json:"eight"`
    Sixteen int `json:"sixteen"`
    ThirtyTwo int `json:"thirty-two"`
    SixtyFour int `json:"sixty-four"`
    OneTwentyEight int `json:"one-twenty-eight"`
}

func (p *SlbRcCacheGlobalOper) GetId() string{
    return "1"
}

func (p *SlbRcCacheGlobalOper) getPath() string{
    return "slb/rc-cache-global/oper"
}

func (p *SlbRcCacheGlobalOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbRcCacheGlobalOper,error) {
logger.Println("SlbRcCacheGlobalOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbRcCacheGlobalOper
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
