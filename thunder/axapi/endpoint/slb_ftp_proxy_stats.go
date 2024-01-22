

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbFtpProxyStats struct {
    
    Stats SlbFtpProxyStatsStats `json:"stats"`

}
type DataSlbFtpProxyStats struct {
    DtSlbFtpProxyStats SlbFtpProxyStats `json:"ftp-proxy"`
}


type SlbFtpProxyStatsStats struct {
    Curr int `json:"curr"`
    Total int `json:"total"`
    Svrsel_fail int `json:"svrsel_fail"`
    No_route int `json:"no_route"`
    Snat_fail int `json:"snat_fail"`
    Feat int `json:"feat"`
    Cc int `json:"cc"`
    Data_ssl int `json:"data_ssl"`
    Line_too_long int `json:"line_too_long"`
    Line_mem_freed int `json:"line_mem_freed"`
    Invalid_start_line int `json:"invalid_start_line"`
    Auth_tls int `json:"auth_tls"`
    Prot int `json:"prot"`
    Pbsz int `json:"pbsz"`
    Pasv int `json:"pasv"`
    Port int `json:"port"`
    Request_dont_care int `json:"request_dont_care"`
    Client_auth_tls int `json:"client_auth_tls"`
    Cant_find_pasv int `json:"cant_find_pasv"`
    Pasv_addr_ne_server int `json:"pasv_addr_ne_server"`
    Smp_create_fail int `json:"smp_create_fail"`
    Data_server_conn_fail int `json:"data_server_conn_fail"`
    Data_send_fail int `json:"data_send_fail"`
    Epsv int `json:"epsv"`
    Cant_find_epsv int `json:"cant_find_epsv"`
    Data_curr int `json:"data_curr"`
    Data_total int `json:"data_total"`
    Auth_unsupported int `json:"auth_unsupported"`
    Adat int `json:"adat"`
    Unsupported_pbsz_value int `json:"unsupported_pbsz_value"`
    Unsupported_prot_value int `json:"unsupported_prot_value"`
    Unsupported_command int `json:"unsupported_command"`
    Control_to_clear int `json:"control_to_clear"`
    Control_to_ssl int `json:"control_to_ssl"`
    Bad_sequence int `json:"bad_sequence"`
    Rsv_persist_conn_fail int `json:"rsv_persist_conn_fail"`
    Smp_v6_fail int `json:"smp_v6_fail"`
    Smp_v4_fail int `json:"smp_v4_fail"`
    Insert_tuple_fail int `json:"insert_tuple_fail"`
    Cl_est_err int `json:"cl_est_err"`
    Ser_connecting_err int `json:"ser_connecting_err"`
    Server_response_err int `json:"server_response_err"`
    Cl_request_err int `json:"cl_request_err"`
    Data_conn_start_err int `json:"data_conn_start_err"`
    Data_serv_connecting_err int `json:"data_serv_connecting_err"`
    Data_serv_connected_err int `json:"data_serv_connected_err"`
    Request int `json:"request"`
    Auth_req int `json:"auth_req"`
    Auth_succ int `json:"auth_succ"`
    Auth_fail int `json:"auth_fail"`
    Fwd_to_internet int `json:"fwd_to_internet"`
    Fwd_to_sg int `json:"fwd_to_sg"`
    Drop int `json:"drop"`
    Ds_succ int `json:"ds_succ"`
    Ds_fail int `json:"ds_fail"`
    Open int `json:"open"`
    Site int `json:"site"`
    User int `json:"user"`
    Pass int `json:"pass"`
    Quit int `json:"quit"`
    Eprt int `json:"eprt"`
    Cant_find_port int `json:"cant_find_port"`
    Cant_find_eprt int `json:"cant_find_eprt"`
}

func (p *SlbFtpProxyStats) GetId() string{
    return "1"
}

func (p *SlbFtpProxyStats) getPath() string{
    return "slb/ftp-proxy/stats"
}

func (p *SlbFtpProxyStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbFtpProxyStats,error) {
logger.Println("SlbFtpProxyStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbFtpProxyStats
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
