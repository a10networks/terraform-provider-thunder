package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSoCounters() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_so_counters`: Show scaleout statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceSoCountersCreate,
		UpdateContext: resourceSoCountersUpdate,
		ReadContext:   resourceSoCountersRead,
		DeleteContext: resourceSoCountersDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'so_pkts_rcvd': Total data packets received; 'so_redirect_pkts_sent': Total packets redirected out of node; 'so_pkts_dropped': Total packets dropped; 'so_redirected_pkts_rcvd': Total redirected packets received on node; 'so_fw_shadow_session_created': FW Shadow Session created; 'so_pkts_traffic_map_not_found_drop': Traffic MAP Not Found Drop; 'so_slb_pkts_redirect_conn_aged_out': Total SLB redirect conns aged out; 'so_pkts_scaleout_not_active_drop': Scaleout Not Active Drop; 'so_pkts_slb_nat_reserve_fail': Total SLB NAT reserve failures; 'so_pkts_slb_nat_release_fail': Total SLB NAT release failures; 'so_pkts_dest_mac_mismatch_drop': Destination MAC Mistmatch Drop; 'so_pkts_l2redirect_dest_mac_zero_drop': Destination MAC Address zero Drop; 'so_pkts_l2redirect_interface_not_up': L2redirect Intf is not UP; 'so_pkts_l2redirect_invalid_redirect_info_error': Redirect Table Error due to invalid redirect info; 'so_pkts_l3_redirect_encap_error_drop': L3 Redirect encap error drop during transmission; 'so_pkts_l3_redirect_inner_mac_zero_drop': L3 Redirect inner mac zero drop during transmission; 'so_pkts_l3_redirect_decap_vlan_sanity_drop': L3 Redirect Decap VLAN Sanity Drop during receipt; 'so_pkts_l3_redirect_decap_non_ipv4_vxlan_drop': L3 Redirect received non ipv4 VXLAN packet; 'so_pkts_l3_redirect_decap_rx_encap_params_drop': L3 Redirect decap Rx encap params error Drop; 'so_pkts_l3_redirect_table_error': L3 Redirect Table error Drop; 'so_pkts_l3_redirect_rcvd_in_l2_mode_drop': Recevied l3 redirect packets in L2 mode Drop; 'so_pkts_l3_redirect_fragmentation_error': L3 redirect encap Fragmentation error; 'so_pkts_l3_redirect_table_no_entry_found': L3 redirect Table no redirect entry found error; 'so_pkts_l3_redirect_invalid_dev_dir': L3 Redirect Invalid Device direction during transmission; 'so_pkts_l3_redirect_chassis_dest_mac_error': L3 Redirect RX multi-slot Destination MAC Error; 'so_pkts_l3_redirect_encap_ipv4_jumbo_frag_drop': L3 Redirect ipv4 packet after encap more than max jumbo size; 'so_pkts_l3_redirect_encap_ipv6_jumbo_frag_drop': L3 Redirect tx ipv6 packet after encap more than max jumbo size; 'so_pkts_l3_redirect_too_large_pkts_in_drop': Received L3 Redirected fragmented packets too large; 'so_sync_fw_shadow_session_create': Sent Sync message for FW Shadow session creation; 'so_sync_fw_shadow_session_delete': Sent Sync message for FW Shadow session deletion; 'so_sync_fw_shadow_ext': Sync FW Shadow extension creation/updation; 'so_sync_shadow_stats_to_active': Sync Shadow session stats from shadow to active; 'so_fw_internal_rule_count': FW internal rule count; 'so_hc_registration_done': Scaleout stats block registered with HC; 'so_hc_deregistration_done': Scaleout stats block de-registered with HC; 'so_pkts_l2redirect_vlan_retrieval_error': L2 redirect pkt vlan not retrieved; 'so_pkts_l2redirect_port_retrieval_error': L2 redirect pkt port not retrieved; 'so_slb_shadow_session_created': SLB Shadow Session created; 'so_sync_slb_shadow_session_create': Sent Sync message for SLB Shadow session creation; 'so_sync_slb_shadow_session_delete': Sent Sync message for SLB Shadow session deletion;",
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
func resourceSoCountersCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSoCountersCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSoCounters(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSoCountersRead(ctx, d, meta)
	}
	return diags
}

func resourceSoCountersUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSoCountersUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSoCounters(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSoCountersRead(ctx, d, meta)
	}
	return diags
}
func resourceSoCountersDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSoCountersDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSoCounters(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSoCountersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSoCountersRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSoCounters(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSoCountersSamplingEnable(d []interface{}) []edpt.SoCountersSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SoCountersSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SoCountersSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSoCounters(d *schema.ResourceData) edpt.SoCounters {
	var ret edpt.SoCounters
	ret.Inst.SamplingEnable = getSliceSoCountersSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
