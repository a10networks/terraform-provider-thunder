

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DnssecOper struct {
    
    Oper DnssecOperOper `json:"oper"`

}
type DataDnssecOper struct {
    DtDnssecOper DnssecOper `json:"dnssec"`
}


type DnssecOperOper struct {
    Soa_memory int `json:"soa_memory"`
    Soa_objects int `json:"soa_objects"`
    Dnskey_memory int `json:"dnskey_memory"`
    Dnskey_objects int `json:"dnskey_objects"`
    Ds_memory int `json:"ds_memory"`
    Ds_objects int `json:"ds_objects"`
    Nsec3param_memory int `json:"nsec3param_memory"`
    Nsec3param_objects int `json:"nsec3param_objects"`
    Nsec_memory int `json:"nsec_memory"`
    Nsec_objects int `json:"nsec_objects"`
    Nsec3_memory int `json:"nsec3_memory"`
    Nsec3_objects int `json:"nsec3_objects"`
    Rrsig_memory int `json:"rrsig_memory"`
    Rrsig_objects int `json:"rrsig_objects"`
    A_memory int `json:"a_memory"`
    A_objects int `json:"a_objects"`
    Aaaa_memory int `json:"aaaa_memory"`
    Aaaa_objects int `json:"aaaa_objects"`
    Ptr_memory int `json:"ptr_memory"`
    Ptr_objects int `json:"ptr_objects"`
    Cname_memory int `json:"cname_memory"`
    Cname_objects int `json:"cname_objects"`
    Ns_memory int `json:"ns_memory"`
    Ns_objects int `json:"ns_objects"`
    Mx_memory int `json:"mx_memory"`
    Mx_objects int `json:"mx_objects"`
    Srv_memory int `json:"srv_memory"`
    Srv_objects int `json:"srv_objects"`
    Txt_memory int `json:"txt_memory"`
    Txt_objects int `json:"txt_objects"`
    Zone_memory int `json:"zone_memory"`
    Zone_objects int `json:"zone_objects"`
    Domain_memory int `json:"domain_memory"`
    Domain_objects int `json:"domain_objects"`
    Table_memory int `json:"table_memory"`
    Table_objects int `json:"table_objects"`
    Reference_memory int `json:"reference_memory"`
    Reference_objects int `json:"reference_objects"`
    Array_memory int `json:"array_memory"`
    Array_objects int `json:"array_objects"`
    Rrsig2_memory int `json:"rrsig2_memory"`
    Rrsig2_objects int `json:"rrsig2_objects"`
    Total_memory int `json:"total_memory"`
    Total_objects int `json:"total_objects"`
}

func (p *DnssecOper) GetId() string{
    return "1"
}

func (p *DnssecOper) getPath() string{
    return "dnssec/oper"
}

func (p *DnssecOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDnssecOper,error) {
logger.Println("DnssecOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDnssecOper
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
