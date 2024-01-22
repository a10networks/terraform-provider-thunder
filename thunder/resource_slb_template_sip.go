package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateSip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_sip`: SIP Template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateSipCreate,
		UpdateContext: resourceSlbTemplateSipUpdate,
		ReadContext:   resourceSlbTemplateSipRead,
		DeleteContext: resourceSlbTemplateSipDelete,

		Schema: map[string]*schema.Schema{
			"acl_id": {
				Type: schema.TypeInt, Optional: true, Description: "ACL id",
			},
			"acl_name_value": {
				Type: schema.TypeString, Optional: true, Description: "IPv4 Access List Name",
			},
			"alg_dest_nat": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Translate VIP to real server IP in SIP message when destination NAT is used",
			},
			"alg_source_nat": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Translate source IP to NAT IP in SIP message when source NAT is used",
			},
			"call_id_persist_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable call-ID persistence",
			},
			"client_keep_alive": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Respond client keep-alive packet directly instead of forwarding to server",
			},
			"client_request_header": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_request_header_erase": {
							Type: schema.TypeString, Optional: true, Description: "Erase a SIP header (Header Name)",
						},
						"client_request_erase_all": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Erase all headers",
						},
						"client_request_header_insert": {
							Type: schema.TypeString, Optional: true, Description: "Insert a SIP header (Header Content (Format: \"name:value\"))",
						},
						"insert_condition_client_request": {
							Type: schema.TypeString, Optional: true, Description: "'insert-if-not-exist': Only insert the header when it does not exist; 'insert-always': Always insert the header even when there is a header with the same name;",
						},
					},
				},
			},
			"client_response_header": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_response_header_erase": {
							Type: schema.TypeString, Optional: true, Description: "Erase a SIP header (Header Name)",
						},
						"client_response_erase_all": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Erase all headers",
						},
						"client_response_header_insert": {
							Type: schema.TypeString, Optional: true, Description: "Insert a SIP header (Header Content (Format: \"name:value\"))",
						},
						"insert_condition_client_response": {
							Type: schema.TypeString, Optional: true, Description: "'insert-if-not-exist': Only insert the header when it does not exist; 'insert-always': Always insert the header even when there is a header with the same name;",
						},
					},
				},
			},
			"dialog_aware": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Permit system processes dialog session",
			},
			"drop_when_client_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop current SIP message when select client fail",
			},
			"drop_when_server_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop current SIP message when select server fail",
			},
			"exclude_translation": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"translation_value": {
							Type: schema.TypeString, Optional: true, Description: "'start-line': SIP request line or status line; 'header': SIP message headers; 'body': SIP message body;",
						},
						"header_string": {
							Type: schema.TypeString, Optional: true, Description: "SIP header name",
						},
					},
				},
			},
			"failed_client_selection": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Define action when select client fail",
			},
			"failed_client_selection_message": {
				Type: schema.TypeString, Optional: true, Description: "Send SIP message (includs status code) to server when select client fail(Format: 3 digits(1XX~6XX) space reason)",
			},
			"failed_server_selection": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Define action when select server fail",
			},
			"failed_server_selection_message": {
				Type: schema.TypeString, Optional: true, Description: "Send SIP message (includs status code) to client when select server fail(Format: 3 digits(1XX~6XX) space reason)",
			},
			"insert_client_ip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Insert Client IP address into SIP header",
			},
			"interval": {
				Type: schema.TypeInt, Optional: true, Default: 30, Description: "The interval of keep-alive packet for each persist connection (second)",
			},
			"keep_server_ip_if_match_acl": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use Real Server IP for addresses matching the ACL for a Call-Id",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "SIP Template Name",
			},
			"pstn_gw": {
				Type: schema.TypeString, Optional: true, Default: "pstn", Description: "configure pstn gw host name for tel: uri translate to sip: uri (Hostname String, default is \"pstn\")",
			},
			"server_keep_alive": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send server keep-alive packet for every persist connection when enable conn-reuse",
			},
			"server_request_header": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server_request_header_erase": {
							Type: schema.TypeString, Optional: true, Description: "Erase a SIP header (Header Name)",
						},
						"server_request_erase_all": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Erase all headers",
						},
						"server_request_header_insert": {
							Type: schema.TypeString, Optional: true, Description: "Insert a SIP header (Header Content (Format: \"name:value\"))",
						},
						"insert_condition_server_request": {
							Type: schema.TypeString, Optional: true, Description: "'insert-if-not-exist': Only insert the header when it does not exist; 'insert-always': Always insert the header even when there is a header with the same name;",
						},
					},
				},
			},
			"server_response_header": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server_response_header_erase": {
							Type: schema.TypeString, Optional: true, Description: "Erase a SIP header (Header Name)",
						},
						"server_response_erase_all": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Erase all headers",
						},
						"server_response_header_insert": {
							Type: schema.TypeString, Optional: true, Description: "Insert a SIP header (Header Content (Format: \"name:value\"))",
						},
						"insert_condition_server_response": {
							Type: schema.TypeString, Optional: true, Description: "'insert-if-not-exist': Only insert the header when it does not exist; 'insert-always': Always insert the header even when there is a header with the same name;",
						},
					},
				},
			},
			"server_selection_per_request": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Force server selection on every SIP request",
			},
			"service_group": {
				Type: schema.TypeString, Optional: true, Description: "service group name",
			},
			"smp_call_id_rtp_session": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Create the across cpu call-id rtp session",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 30, Description: "Time in minutes",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbTemplateSipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateSipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateSip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateSipRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateSipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateSipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateSip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateSipRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateSipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateSipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateSip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateSipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateSipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateSip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateSipClientRequestHeader(d []interface{}) []edpt.SlbTemplateSipClientRequestHeader {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateSipClientRequestHeader, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateSipClientRequestHeader
		oi.ClientRequestHeaderErase = in["client_request_header_erase"].(string)
		oi.ClientRequestEraseAll = in["client_request_erase_all"].(int)
		oi.ClientRequestHeaderInsert = in["client_request_header_insert"].(string)
		oi.InsertConditionClientRequest = in["insert_condition_client_request"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateSipClientResponseHeader(d []interface{}) []edpt.SlbTemplateSipClientResponseHeader {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateSipClientResponseHeader, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateSipClientResponseHeader
		oi.ClientResponseHeaderErase = in["client_response_header_erase"].(string)
		oi.ClientResponseEraseAll = in["client_response_erase_all"].(int)
		oi.ClientResponseHeaderInsert = in["client_response_header_insert"].(string)
		oi.InsertConditionClientResponse = in["insert_condition_client_response"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateSipExcludeTranslation(d []interface{}) []edpt.SlbTemplateSipExcludeTranslation {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateSipExcludeTranslation, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateSipExcludeTranslation
		oi.TranslationValue = in["translation_value"].(string)
		oi.HeaderString = in["header_string"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateSipServerRequestHeader(d []interface{}) []edpt.SlbTemplateSipServerRequestHeader {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateSipServerRequestHeader, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateSipServerRequestHeader
		oi.ServerRequestHeaderErase = in["server_request_header_erase"].(string)
		oi.ServerRequestEraseAll = in["server_request_erase_all"].(int)
		oi.ServerRequestHeaderInsert = in["server_request_header_insert"].(string)
		oi.InsertConditionServerRequest = in["insert_condition_server_request"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateSipServerResponseHeader(d []interface{}) []edpt.SlbTemplateSipServerResponseHeader {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateSipServerResponseHeader, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateSipServerResponseHeader
		oi.ServerResponseHeaderErase = in["server_response_header_erase"].(string)
		oi.ServerResponseEraseAll = in["server_response_erase_all"].(int)
		oi.ServerResponseHeaderInsert = in["server_response_header_insert"].(string)
		oi.InsertConditionServerResponse = in["insert_condition_server_response"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateSip(d *schema.ResourceData) edpt.SlbTemplateSip {
	var ret edpt.SlbTemplateSip
	ret.Inst.AclId = d.Get("acl_id").(int)
	ret.Inst.AclNameValue = d.Get("acl_name_value").(string)
	ret.Inst.AlgDestNat = d.Get("alg_dest_nat").(int)
	ret.Inst.AlgSourceNat = d.Get("alg_source_nat").(int)
	ret.Inst.CallIdPersistDisable = d.Get("call_id_persist_disable").(int)
	ret.Inst.ClientKeepAlive = d.Get("client_keep_alive").(int)
	ret.Inst.ClientRequestHeader = getSliceSlbTemplateSipClientRequestHeader(d.Get("client_request_header").([]interface{}))
	ret.Inst.ClientResponseHeader = getSliceSlbTemplateSipClientResponseHeader(d.Get("client_response_header").([]interface{}))
	ret.Inst.DialogAware = d.Get("dialog_aware").(int)
	ret.Inst.DropWhenClientFail = d.Get("drop_when_client_fail").(int)
	ret.Inst.DropWhenServerFail = d.Get("drop_when_server_fail").(int)
	ret.Inst.ExcludeTranslation = getSliceSlbTemplateSipExcludeTranslation(d.Get("exclude_translation").([]interface{}))
	ret.Inst.FailedClientSelection = d.Get("failed_client_selection").(int)
	ret.Inst.FailedClientSelectionMessage = d.Get("failed_client_selection_message").(string)
	ret.Inst.FailedServerSelection = d.Get("failed_server_selection").(int)
	ret.Inst.FailedServerSelectionMessage = d.Get("failed_server_selection_message").(string)
	ret.Inst.InsertClientIp = d.Get("insert_client_ip").(int)
	ret.Inst.Interval = d.Get("interval").(int)
	ret.Inst.KeepServerIpIfMatchAcl = d.Get("keep_server_ip_if_match_acl").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PstnGw = d.Get("pstn_gw").(string)
	ret.Inst.ServerKeepAlive = d.Get("server_keep_alive").(int)
	ret.Inst.ServerRequestHeader = getSliceSlbTemplateSipServerRequestHeader(d.Get("server_request_header").([]interface{}))
	ret.Inst.ServerResponseHeader = getSliceSlbTemplateSipServerResponseHeader(d.Get("server_response_header").([]interface{}))
	ret.Inst.ServerSelectionPerRequest = d.Get("server_selection_per_request").(int)
	ret.Inst.ServiceGroup = d.Get("service_group").(string)
	ret.Inst.SmpCallIdRtpSession = d.Get("smp_call_id_rtp_session").(int)
	ret.Inst.Timeout = d.Get("timeout").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
