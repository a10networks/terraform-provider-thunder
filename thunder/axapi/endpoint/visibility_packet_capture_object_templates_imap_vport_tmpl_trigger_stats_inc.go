

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsInc struct {
	Inst struct {

    Auth_unsupported int `json:"auth_unsupported"`
    Bad_sequence int `json:"bad_sequence"`
    Cant_find_epsv int `json:"cant_find_epsv"`
    Cant_find_pasv int `json:"cant_find_pasv"`
    Cl_est_err int `json:"cl_est_err"`
    Cl_request_err int `json:"cl_request_err"`
    Data_send_fail int `json:"data_send_fail"`
    Data_server_conn_fail int `json:"data_server_conn_fail"`
    Insert_tuple_fail int `json:"insert_tuple_fail"`
    Invalid_start_line int `json:"invalid_start_line"`
    Line_too_long int `json:"line_too_long"`
    No_route int `json:"no_route"`
    Rsv_persist_conn_fail int `json:"rsv_persist_conn_fail"`
    Ser_connecting_err int `json:"ser_connecting_err"`
    Server_response_err int `json:"server_response_err"`
    Smp_create_fail int `json:"smp_create_fail"`
    Smp_v4_fail int `json:"smp_v4_fail"`
    Smp_v6_fail int `json:"smp_v6_fail"`
    Snat_fail int `json:"snat_fail"`
    Svrsel_fail int `json:"svrsel_fail"`
    Unsupported_pbsz_value int `json:"unsupported_pbsz_value"`
    Unsupported_prot_value int `json:"unsupported_prot_value"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"trigger-stats-inc"`
}

func (p *VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsInc) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsInc) getPath() string{
    return "visibility/packet-capture/object-templates/imap-vport-tmpl/" +p.Inst.Name + "/trigger-stats-inc"
}

func (p *VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsInc) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsInc::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
    return err
}

func (p *VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsInc) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsInc::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsInc) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsInc::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsInc) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsInc::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
