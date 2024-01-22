package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHmDplane() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_hm_dplane`: Configure hm-dplane\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbHmDplaneCreate,
		UpdateContext: resourceSlbHmDplaneUpdate,
		ReadContext:   resourceSlbHmDplaneRead,
		DeleteContext: resourceSlbHmDplaneDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'curr_entries': Current HM Entries; 'total_created': Total HM Entries Created; 'total_inserted': Total HM entries inserted; 'total_ready_to_free': Total HM entries ready to free; 'total_freed': Total HM entries freed; 'err_entry_create_failed': Entry Creation Failure; 'err_entry_create_oom': Entry creation out of memory; 'err_entry_insert_failed': Entry insert failed; 'total_tcp_err': Total TCP errors in health-checks sent; 'err_smart_nat_alloc': Error creating smart-nat instance; 'err_smart_nat_port_alloc': Error obtaining smart-nat source port; 'err_l4_sess_alloc': Error allocating L4 session for HM; 'err_hm_tcp_conn_sent': Error in initiating TCP connection for HM; 'hm_tcp_conn_sent': Total TCP connections sent for HM; 'entry_deleted': Entry deleted; 'err_entry_create_vip_failed': Error in creating HM internal VIP; 'total_match_resp_code': Total HTTP received response with match response code; 'total_match_default_resp_code': Total HTTP received response with match 200 response code; 'total_maintenance_received': Total maintenace response received; 'total_wrong_status_received': Total HTTP received response with wrong response code; 'err_no_hm_entry': Error no HM entry found; 'err_ssl_cert_name_mismatch': Error SSL cert name mismatch; 'err_server_syn_timeout': Error SSL server SYN timeout; 'err_http2_callback': Error HTTP2 callback; 'err_l7_sess_process_tcp_estab_failed': L7 session process TCP established failed; 'err_l7_sess_process_tcp_data_failed': L7 session process TCP data failed; 'err_http2_ver_mismatch': Error HTTP2 version mismatch; 'smart_nat_alloc': Total smart-nat allocation successful; 'smart_nat_release': Total smart-nat release successful; 'smart_nat_alloc_failed': Total smart-nat allocation failed; 'smart_nat_release_failed': Total smart-nat release failed; 'total_server_quic_conn': Total start server QUIC connections; 'total_server_quic_conn_err': Total start server QUIC connections error;",
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
func resourceSlbHmDplaneCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHmDplaneCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHmDplane(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbHmDplaneRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbHmDplaneUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHmDplaneUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHmDplane(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbHmDplaneRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbHmDplaneDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHmDplaneDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHmDplane(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbHmDplaneRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHmDplaneRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHmDplane(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbHmDplaneSamplingEnable(d []interface{}) []edpt.SlbHmDplaneSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbHmDplaneSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbHmDplaneSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbHmDplane(d *schema.ResourceData) edpt.SlbHmDplane {
	var ret edpt.SlbHmDplane
	ret.Inst.SamplingEnable = getSliceSlbHmDplaneSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
