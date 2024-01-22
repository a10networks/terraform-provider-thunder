

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbLinkProbeEntryOper struct {
    
    Oper SlbLinkProbeEntryOperOper `json:"oper"`

}
type DataSlbLinkProbeEntryOper struct {
    DtSlbLinkProbeEntryOper SlbLinkProbeEntryOper `json:"entry"`
}


type SlbLinkProbeEntryOperOper struct {
    EntryList []SlbLinkProbeEntryOperOperEntryList `json:"entry-list"`
}


type SlbLinkProbeEntryOperOperEntryList struct {
    Rserver_name string `json:"rserver_name"`
    Sg_name string `json:"sg_name"`
    Probe_template_name string `json:"probe_template_name"`
    Domain_name string `json:"domain_name"`
    Url string `json:"url"`
    Target_type int `json:"target_type"`
    Ip_type int `json:"ip_type"`
    Ipv4_addr string `json:"ipv4_addr"`
    Ipv6_addr string `json:"ipv6_addr"`
    Ref_count int `json:"ref_count"`
    Data_cpu_id int `json:"data_cpu_id"`
    Curr_probe_count int `json:"curr_probe_count"`
    Test_interval int `json:"test_interval"`
    Probe_interval int `json:"probe_interval"`
    Probes_per_test int `json:"probes_per_test"`
    Rtt_method int `json:"rtt_method"`
    Last_status_code int `json:"last_status_code"`
    Rtt_avg int `json:"rtt_avg"`
    Rtt_samples1 int `json:"rtt_samples1"`
    Rtt_samples2 int `json:"rtt_samples2"`
    Rtt_samples3 int `json:"rtt_samples3"`
    Rtt_samples4 int `json:"rtt_samples4"`
    Rtt_samples5 int `json:"rtt_samples5"`
    Rtt_samples6 int `json:"rtt_samples6"`
    Rtt_samples7 int `json:"rtt_samples7"`
    Rtt_samples8 int `json:"rtt_samples8"`
    Rtt_samples9 int `json:"rtt_samples9"`
    Rtt_samples10 int `json:"rtt_samples10"`
}

func (p *SlbLinkProbeEntryOper) GetId() string{
    return "1"
}

func (p *SlbLinkProbeEntryOper) getPath() string{
    return "slb/link-probe/entry/oper"
}

func (p *SlbLinkProbeEntryOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbLinkProbeEntryOper,error) {
logger.Println("SlbLinkProbeEntryOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbLinkProbeEntryOper
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
