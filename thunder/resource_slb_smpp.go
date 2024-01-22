package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSmpp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_smpp`: Configure SMPP\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbSmppCreate,
		UpdateContext: resourceSlbSmppUpdate,
		ReadContext:   resourceSlbSmppRead,
		DeleteContext: resourceSlbSmppDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'msg_proxy_current': Curr SMPP Proxy; 'msg_proxy_total': Total SMPP Proxy; 'msg_proxy_mem_allocd': msg_proxy_mem_allocd; 'msg_proxy_mem_cached': msg_proxy_mem_cached; 'msg_proxy_mem_freed': msg_proxy_mem_freed; 'msg_proxy_client_recv': Client message rcvd; 'msg_proxy_client_send_success': Sent to server; 'msg_proxy_client_incomplete': Incomplete; 'msg_proxy_client_drop': AX responds directly; 'msg_proxy_client_connection': Connecting server; 'msg_proxy_client_fail': Number of SMPP messages received from client but failed to forward to server; 'msg_proxy_client_fail_parse': msg_proxy_client_fail_parse; 'msg_proxy_client_fail_process': msg_proxy_client_fail_process; 'msg_proxy_client_fail_snat': msg_proxy_client_fail_snat; 'msg_proxy_client_exceed_tmp_buff': msg_proxy_client_exceed_tmp_buff; 'msg_proxy_client_fail_send_pkt': msg_proxy_client_fail_send_pkt; 'msg_proxy_client_fail_start_server_Conn': msg_proxy_client_fail_start_server_Conn; 'msg_proxy_server_recv': Server message rcvd; 'msg_proxy_server_send_success': Sent to client; 'msg_proxy_server_incomplete': Incomplete; 'msg_proxy_server_drop': Number of the packet AX drop; 'msg_proxy_server_fail': Number of SMPP messages received from server but failed to forward to client; 'msg_proxy_server_fail_parse': msg_proxy_server_fail_parse; 'msg_proxy_server_fail_process': msg_proxy_server_fail_process; 'msg_proxy_server_fail_selec_connt': msg_proxy_server_fail_selec_connt; 'msg_proxy_server_fail_snat': msg_proxy_server_fail_snat; 'msg_proxy_server_exceed_tmp_buff': msg_proxy_server_exceed_tmp_buff; 'msg_proxy_server_fail_send_pkt': msg_proxy_server_fail_send_pkt; 'msg_proxy_create_server_conn': Server conn created; 'msg_proxy_start_server_conn': Number of server connection created successfully; 'msg_proxy_fail_start_server_conn': Number of server connection created failed; 'msg_proxy_server_conn_fail_snat': msg_proxy_server_conn_fail_snat; 'msg_proxy_fail_construct_server_conn': msg_proxy_fail_construct_server_conn; 'msg_proxy_fail_reserve_pconn': msg_proxy_fail_reserve_pconn; 'msg_proxy_start_server_conn_failed': msg_proxy_start_server_conn_failed; 'msg_proxy_server_conn_already_exists': msg_proxy_server_conn_already_exists; 'msg_proxy_fail_insert_server_conn': msg_proxy_fail_insert_server_conn; 'msg_proxy_parse_msg_fail': msg_proxy_parse_msg_fail; 'msg_proxy_process_msg_fail': msg_proxy_process_msg_fail; 'msg_proxy_no_vport': msg_proxy_no_vport; 'msg_proxy_fail_select_server': msg_proxy_fail_select_server; 'msg_proxy_fail_alloc_mem': msg_proxy_fail_alloc_mem; 'msg_proxy_unexpected_err': msg_proxy_unexpected_err; 'msg_proxy_l7_cpu_failed': msg_proxy_l7_cpu_failed; 'msg_proxy_l4_to_l7': msg_proxy_l4_to_l7; 'msg_proxy_l4_from_l7': msg_proxy_l4_from_l7; 'msg_proxy_to_l4_send_pkt': msg_proxy_to_l4_send_pkt; 'msg_proxy_l4_from_l4_send': msg_proxy_l4_from_l4_send; 'msg_proxy_l7_to_L4': msg_proxy_l7_to_L4; 'msg_proxy_mag_back': msg_proxy_mag_back; 'msg_proxy_fail_dcmsg': msg_proxy_fail_dcmsg; 'msg_proxy_deprecated_conn': msg_proxy_deprecated_conn; 'msg_proxy_hold_msg': msg_proxy_hold_msg; 'msg_proxy_split_pkt': msg_proxy_split_pkt; 'msg_proxy_pipline_msg': msg_proxy_pipline_msg; 'msg_proxy_client_reset': msg_proxy_client_reset; 'msg_proxy_server_reset': msg_proxy_server_reset; 'payload_allocd': payload_allocd; 'payload_freed': payload_freed; 'pkt_too_small': pkt_too_small; 'invalid_seq': invalid_seq; 'AX_response_directly': Number of packet which AX responds directly; 'select_client_conn': Client conn selection; 'select_client_by_req': Select by request; 'select_client_from_list': Select by roundbin; 'select_client_by_conn': Select by conn; 'select_client_fail': Select failed; 'select_server_conn': Server conn selection; 'select_server_by_req': Select by request; 'select_server_from_list': Select by roundbin; 'select_server_by_conn': Select server conn by client conn; 'select_server_fail': Fail to select server conn; 'bind_conn': bind_conn; 'unbind_conn': unbind_conn; 'enquire_link_recv': enquire_link_recv; 'enquire_link_resp_recv': enquire_link_resp_recv; 'enquire_link_send': enquire_link_send; 'enquire_link_resp_send': enquire_link_resp_send; 'client_conn_put_in_list': client_conn_put_in_list; 'client_conn_get_from_list': client_conn_get_from_list; 'server_conn_put_in_list': server_conn_put_in_list; 'server_conn_get_from_list': server_conn_get_from_list; 'server_conn_fail_bind': server_conn_fail_bind; 'single_msg': single_msg; 'fail_bind_msg': fail_bind_msg;",
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
func resourceSlbSmppCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSmppCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSmpp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSmppRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbSmppUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSmppUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSmpp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSmppRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbSmppDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSmppDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSmpp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbSmppRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSmppRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSmpp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbSmppSamplingEnable(d []interface{}) []edpt.SlbSmppSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbSmppSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSmppSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSmpp(d *schema.ResourceData) edpt.SlbSmpp {
	var ret edpt.SlbSmpp
	ret.Inst.SamplingEnable = getSliceSlbSmppSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
