

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbPersistOper struct {
    
    Oper SlbPersistOperOper `json:"oper"`

}
type DataSlbPersistOper struct {
    DtSlbPersistOper SlbPersistOper `json:"persist"`
}


type SlbPersistOperOper struct {
    PersistCpuList []SlbPersistOperOperPersistCpuList `json:"persist-cpu-list"`
    CpuCount int `json:"cpu-count"`
}


type SlbPersistOperOperPersistCpuList struct {
    Hash_tbl_trylock_fail int `json:"hash_tbl_trylock_fail"`
    Hash_tbl_create_ok int `json:"hash_tbl_create_ok"`
    Hash_tbl_create_fail int `json:"hash_tbl_create_fail"`
    Hash_tbl_free int `json:"hash_tbl_free"`
    Hash_tbl_rst_updown int `json:"hash_tbl_rst_updown"`
    Hash_tbl_rst_adddel int `json:"hash_tbl_rst_adddel"`
    Url_hash_pri int `json:"url_hash_pri"`
    Url_hash_enqueue int `json:"url_hash_enqueue"`
    Url_hash_sec int `json:"url_hash_sec"`
    Url_hash_fail int `json:"url_hash_fail"`
    Header_hash_pri int `json:"header_hash_pri"`
    Header_hash_enqueue int `json:"header_hash_enqueue"`
    Header_hash_sec int `json:"header_hash_sec"`
    Header_hash_fail int `json:"header_hash_fail"`
    Src_ip int `json:"src_ip"`
    Src_ip_enqueue int `json:"src_ip_enqueue"`
    Src_ip_fail int `json:"src_ip_fail"`
    Src_ip_new_sess_cache int `json:"src_ip_new_sess_cache"`
    Src_ip_new_sess_cache_fail int `json:"src_ip_new_sess_cache_fail"`
    Src_ip_new_sess_sel int `json:"src_ip_new_sess_sel"`
    Src_ip_new_sess_sel_fail int `json:"src_ip_new_sess_sel_fail"`
    Src_ip_hash_pri int `json:"src_ip_hash_pri"`
    Src_ip_hash_enqueue int `json:"src_ip_hash_enqueue"`
    Src_ip_hash_sec int `json:"src_ip_hash_sec"`
    Src_ip_hash_fail int `json:"src_ip_hash_fail"`
    Src_ip_enforce int `json:"src_ip_enforce"`
    Dst_ip int `json:"dst_ip"`
    Dst_ip_enqueue int `json:"dst_ip_enqueue"`
    Dst_ip_fail int `json:"dst_ip_fail"`
    Dst_ip_new_sess_cache int `json:"dst_ip_new_sess_cache"`
    Dst_ip_new_sess_cache_fail int `json:"dst_ip_new_sess_cache_fail"`
    Dst_ip_new_sess_sel int `json:"dst_ip_new_sess_sel"`
    Dst_ip_new_sess_sel_fail int `json:"dst_ip_new_sess_sel_fail"`
    Dst_ip_hash_pri int `json:"dst_ip_hash_pri"`
    Dst_ip_hash_enqueue int `json:"dst_ip_hash_enqueue"`
    Dst_ip_hash_sec int `json:"dst_ip_hash_sec"`
    Dst_ip_hash_fail int `json:"dst_ip_hash_fail"`
    Cssl_sid_not_found int `json:"cssl_sid_not_found"`
    Cssl_sid_match int `json:"cssl_sid_match"`
    Cssl_sid_not_match int `json:"cssl_sid_not_match"`
    Sssl_sid_not_found int `json:"sssl_sid_not_found"`
    Sssl_sid_reset int `json:"sssl_sid_reset"`
    Sssl_sid_match int `json:"sssl_sid_match"`
    Sssl_sid_not_match int `json:"sssl_sid_not_match"`
    Ssl_sid_persist_ok int `json:"ssl_sid_persist_ok"`
    Ssl_sid_persist_fail int `json:"ssl_sid_persist_fail"`
    Ssl_sid_session_ok int `json:"ssl_sid_session_ok"`
    Ssl_sid_session_fail int `json:"ssl_sid_session_fail"`
    Cookie_persist_ok int `json:"cookie_persist_ok"`
    Cookie_persist_fail int `json:"cookie_persist_fail"`
    Cookie_not_found int `json:"cookie_not_found"`
    Cookie_pass_thru int `json:"cookie_pass_thru"`
    Cookie_invalid int `json:"cookie_invalid"`
}

func (p *SlbPersistOper) GetId() string{
    return "1"
}

func (p *SlbPersistOper) getPath() string{
    return "slb/persist/oper"
}

func (p *SlbPersistOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbPersistOper,error) {
logger.Println("SlbPersistOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbPersistOper
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
