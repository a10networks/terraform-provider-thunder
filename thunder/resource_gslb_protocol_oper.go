package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbProtocolOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_protocol_oper`: Operational Status for the object protocol\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbProtocolOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"session_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol_info": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"state": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"session_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"connection_succeeded": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"connection_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"open_packet_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"open_packet_received": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"open_session_succeeded": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"open_session_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sessions_dropped": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"retry": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"update_packet_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"update_packet_received": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"keepalive_packet_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"keepalive_packet_received": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"notify_packet_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"notify_packet_received": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"message_header_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"secure_negotiation_success": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"secure_negotiation_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ssl_handshake_success": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ssl_handshake_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"secure_state": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"secure_config": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceGslbProtocolOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbProtocolOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbProtocolOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbProtocolOperOper := setObjectGslbProtocolOperOper(res)
		d.Set("oper", GslbProtocolOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbProtocolOperOper(ret edpt.DataGslbProtocolOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"session_list": setSliceGslbProtocolOperOperSessionList(ret.DtGslbProtocolOper.Oper.SessionList),
		},
	}
}

func setSliceGslbProtocolOperOperSessionList(d []edpt.GslbProtocolOperOperSessionList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["protocol_info"] = item.ProtocolInfo
		in["state"] = item.State
		in["session_id"] = item.SessionId
		in["connection_succeeded"] = item.ConnectionSucceeded
		in["connection_failed"] = item.ConnectionFailed
		in["open_packet_sent"] = item.OpenPacketSent
		in["open_packet_received"] = item.OpenPacketReceived
		in["open_session_succeeded"] = item.OpenSessionSucceeded
		in["open_session_failed"] = item.OpenSessionFailed
		in["sessions_dropped"] = item.SessionsDropped
		in["retry"] = item.Retry
		in["update_packet_sent"] = item.UpdatePacketSent
		in["update_packet_received"] = item.UpdatePacketReceived
		in["keepalive_packet_sent"] = item.KeepalivePacketSent
		in["keepalive_packet_received"] = item.KeepalivePacketReceived
		in["notify_packet_sent"] = item.NotifyPacketSent
		in["notify_packet_received"] = item.NotifyPacketReceived
		in["message_header_error"] = item.MessageHeaderError
		in["secure_negotiation_success"] = item.SecureNegotiationSuccess
		in["secure_negotiation_fail"] = item.SecureNegotiationFail
		in["ssl_handshake_success"] = item.SslHandshakeSuccess
		in["ssl_handshake_fail"] = item.SslHandshakeFail
		in["secure_state"] = item.Secure_state
		in["secure_config"] = item.Secure_config
		result = append(result, in)
	}
	return result
}

func getObjectGslbProtocolOperOper(d []interface{}) edpt.GslbProtocolOperOper {

	count1 := len(d)
	var ret edpt.GslbProtocolOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionList = getSliceGslbProtocolOperOperSessionList(in["session_list"].([]interface{}))
	}
	return ret
}

func getSliceGslbProtocolOperOperSessionList(d []interface{}) []edpt.GslbProtocolOperOperSessionList {

	count1 := len(d)
	ret := make([]edpt.GslbProtocolOperOperSessionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbProtocolOperOperSessionList
		oi.ProtocolInfo = in["protocol_info"].(string)
		oi.State = in["state"].(string)
		oi.SessionId = in["session_id"].(int)
		oi.ConnectionSucceeded = in["connection_succeeded"].(int)
		oi.ConnectionFailed = in["connection_failed"].(int)
		oi.OpenPacketSent = in["open_packet_sent"].(int)
		oi.OpenPacketReceived = in["open_packet_received"].(int)
		oi.OpenSessionSucceeded = in["open_session_succeeded"].(int)
		oi.OpenSessionFailed = in["open_session_failed"].(int)
		oi.SessionsDropped = in["sessions_dropped"].(int)
		oi.Retry = in["retry"].(int)
		oi.UpdatePacketSent = in["update_packet_sent"].(int)
		oi.UpdatePacketReceived = in["update_packet_received"].(int)
		oi.KeepalivePacketSent = in["keepalive_packet_sent"].(int)
		oi.KeepalivePacketReceived = in["keepalive_packet_received"].(int)
		oi.NotifyPacketSent = in["notify_packet_sent"].(int)
		oi.NotifyPacketReceived = in["notify_packet_received"].(int)
		oi.MessageHeaderError = in["message_header_error"].(int)
		oi.SecureNegotiationSuccess = in["secure_negotiation_success"].(int)
		oi.SecureNegotiationFail = in["secure_negotiation_fail"].(int)
		oi.SslHandshakeSuccess = in["ssl_handshake_success"].(int)
		oi.SslHandshakeFail = in["ssl_handshake_fail"].(int)
		oi.Secure_state = in["secure_state"].(string)
		oi.Secure_config = in["secure_config"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbProtocolOper(d *schema.ResourceData) edpt.GslbProtocolOper {
	var ret edpt.GslbProtocolOper

	ret.Oper = getObjectGslbProtocolOperOper(d.Get("oper").([]interface{}))
	return ret
}
