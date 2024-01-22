package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbGenericProxy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_generic_proxy`: Configure Generic Proxy\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbGenericProxyCreate,
		UpdateContext: resourceSlbGenericProxyUpdate,
		ReadContext:   resourceSlbGenericProxyRead,
		DeleteContext: resourceSlbGenericProxyDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'num': Number; 'curr': Current; 'total': Total; 'svrsel_fail': Number of server selection failed; 'no_route': Number of no routes; 'snat_fail': Number of snat failures; 'client_fail': Number of client failures; 'server_fail': Number of server failures; 'no_sess': Number of no sessions; 'user_session': Number of user sessions; 'acr_out': Number of ACRs out; 'acr_in': Number of ACRs in; 'aca_out': Number of ACAs out; 'aca_in': Number of ACAs in; 'cea_out': Number of CEAs out; 'cea_in': Number of CEAs in; 'cer_out': Number of CERs out; 'cer_in': Number of CERs in; 'dwr_out': Number of DWRs out; 'dwr_in': Number of DWRs in; 'dwa_out': Number of DWAs out; 'dwa_in': Number of DWAs in; 'str_out': Number of STRs out; 'str_in': Number of STRs in; 'sta_out': Number of STAs out; 'sta_in': Number of STAs in; 'asr_out': Number of ASRs out; 'asr_in': Number of ASRs in; 'asa_out': Number of ASAs out; 'asa_in': Number of ASAs in; 'other_out': Number of other messages out; 'other_in': Number of other messages in; 'total_http_req_enter_gen': Total number of HTTP requests enter generic proxy; 'mismatch_fwd_id': Diameter mismatch fwd session id; 'mismatch_rev_id': Diameter mismatch rev session id; 'unkwn_cmd_code': Diameter unkown cmd code; 'no_session_id': Diameter no session id avp; 'no_fwd_tuple': Diameter no fwd tuple matched; 'no_rev_tuple': Diameter no rev tuple matched; 'dcmsg_fwd_in': Diameter cross cpu fwd in; 'dcmsg_fwd_out': Diameter cross cpu fwd out; 'dcmsg_rev_in': Diameter cross cpu rev in; 'dcmsg_rev_out': Diameter cross cpu rev out; 'dcmsg_error': Diameter cross cpu error; 'retry_client_request': Diameter retry client request; 'retry_client_request_fail': Diameter retry client request fail; 'reply_unknown_session_id': Reply with unknown session ID error info; 'ccr_out': Number of CCRs out; 'ccr_in': Number of CCRs in; 'cca_out': Number of CCAs out; 'cca_in': Number of CCAs in; 'ccr_i': Number of CCRs initial; 'ccr_u': Number of CCRs update; 'ccr_t': Number of CCRs terminate; 'cca_t': Number of CCAs terminate; 'terminate_on_cca_t': Diameter terminate on cca_t; 'forward_unknown_session_id': Forward server side message with unknown session id; 'update_latest_server': Update to the latest server that used a session id; 'client_select_fail': Fail to select client; 'close_conn_when_vport_down': Close client conn when virtual port is down; 'invalid_avp': AVP value contains illegal chars; 'reselect_fwd_tuple': Original client tuple does not exist so reselect another one; 'reselect_fwd_tuple_other_cpu': Original client tuple does not exist so reselect another one on other CPUs; 'reselect_rev_tuple': Original server tuple does not exist so reselect another one; 'conn_closed_by_client': Client initiates TCP close/reset; 'conn_closed_by_server': Server initiates TCP close/reset; 'reply_invalid_avp_value': Reply with invalid AVP error info; 'reply_unable_to_deliver': Reply with unable to deliver error info; 'reply_error_info_fail': Fail to reply error info to peer; 'dpr_out': Number of DPRs out; 'dpr_in': Number of DPRs in; 'dpa_out': Number of DPAs out; 'dpa_in': Number of DPAs in;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbGenericProxyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbGenericProxyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbGenericProxy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbGenericProxyRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbGenericProxyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbGenericProxyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbGenericProxy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbGenericProxyRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbGenericProxyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbGenericProxyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbGenericProxy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbGenericProxyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbGenericProxyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbGenericProxy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbGenericProxySamplingEnable(d []interface{}) []edpt.SlbGenericProxySamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbGenericProxySamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbGenericProxySamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbGenericProxy(d *schema.ResourceData) edpt.SlbGenericProxy {
	var ret edpt.SlbGenericProxy
	ret.Inst.SamplingEnable = getSliceSlbGenericProxySamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
