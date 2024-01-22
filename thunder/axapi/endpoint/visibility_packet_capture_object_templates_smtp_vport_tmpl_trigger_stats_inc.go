

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsInc struct {
	Inst struct {

    Forward_req_data_fail int `json:"forward_req_data_fail"`
    Forward_req_fail int `json:"forward_req_fail"`
    Get_all_headers_fail int `json:"get_all_headers_fail"`
    Insert_resonse_line_fail int `json:"insert_resonse_line_fail"`
    Line_extend_fail int `json:"line_extend_fail"`
    Line_table_extend_fail int `json:"line_table_extend_fail"`
    Line_too_long int `json:"line_too_long"`
    No_proxy int `json:"no_proxy"`
    Parse_req_fail int `json:"parse_req_fail"`
    Parse_request_line_fail int `json:"parse_request_line_fail"`
    Parse_resonse_line_fail int `json:"parse_resonse_line_fail"`
    Read_request_line_fail int `json:"read_request_line_fail"`
    Recv_server_unknow_reply_code int `json:"recv_server_unknow_reply_code"`
    Remove_resonse_line_fail int `json:"remove_resonse_line_fail"`
    Send_client_service_not_ready int `json:"send_client_service_not_ready"`
    Server_select_fail int `json:"server_select_fail"`
    Server_starttls_fail int `json:"server_STARTTLS_fail"`
    Snat_fail int `json:"snat_fail"`
    Too_many_headers int `json:"too_many_headers"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"trigger-stats-inc"`
}

func (p *VisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsInc) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsInc) getPath() string{
    return "visibility/packet-capture/object-templates/smtp-vport-tmpl/" +p.Inst.Name + "/trigger-stats-inc"
}

func (p *VisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsInc) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsInc::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsInc) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsInc::Get")
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
func (p *VisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsInc) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsInc::Put")
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

func (p *VisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsInc) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsInc::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
