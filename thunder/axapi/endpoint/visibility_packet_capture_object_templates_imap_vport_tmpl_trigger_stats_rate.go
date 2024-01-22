

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRate struct {
	Inst struct {

    Auth_unsupported int `json:"auth_unsupported"`
    Bad_sequence int `json:"bad_sequence"`
    Cant_find_epsv int `json:"cant_find_epsv"`
    Cant_find_pasv int `json:"cant_find_pasv"`
    Cl_est_err int `json:"cl_est_err"`
    Cl_request_err int `json:"cl_request_err"`
    Data_send_fail int `json:"data_send_fail"`
    Data_server_conn_fail int `json:"data_server_conn_fail"`
    Duration int `json:"duration" dval:"60"`
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
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Unsupported_pbsz_value int `json:"unsupported_pbsz_value"`
    Unsupported_prot_value int `json:"unsupported_prot_value"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"trigger-stats-rate"`
}

func (p *VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRate) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRate) getPath() string{
    return "visibility/packet-capture/object-templates/imap-vport-tmpl/" +p.Inst.Name + "/trigger-stats-rate"
}

func (p *VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRate::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRate::Get")
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
func (p *VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRate::Put")
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

func (p *VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
