package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCounters() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_so_counters`: Configure triggers for so-counters object\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersDelete,

		Schema: map[string]*schema.Schema{
			"trigger_stats_inc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"so_pkts_slb_nat_reserve_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total SLB NAT reserve failures",
						},
						"so_pkts_slb_nat_release_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total SLB NAT release failures",
						},
						"so_pkts_l2redirect_dest_mac_zero_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Destination MAC Address zero Drop",
						},
						"so_pkts_l2redirect_interface_not_up": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L2redirect Intf is not UP",
						},
						"so_pkts_l2redirect_invalid_redirect_inf": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Redirect Table Error due to invalid redirect info",
						},
						"so_pkts_l3_redirect_encap_error_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Redirect encap error drop during transmission",
						},
						"so_pkts_l3_redirect_inner_mac_zero_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Redirect inner mac zero drop during transmission",
						},
						"so_pkts_l3_redirect_table_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Redirect Table error Drop",
						},
						"so_pkts_l3_redirect_fragmentation_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 redirect encap Fragmentation error",
						},
						"so_pkts_l3_redirect_table_no_entry_foun": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 redirect Table no redirect entry found error",
						},
						"so_pkts_l3_redirect_invalid_dev_dir": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Redirect Invalid Device direction during transmission",
						},
						"so_pkts_l3_redirect_chassis_dest_mac_er": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Redirect RX multi-slot Destination MAC Error",
						},
						"so_pkts_l2redirect_vlan_retrieval_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L2 redirect pkt vlan not retrieved",
						},
						"so_pkts_l2redirect_port_retrieval_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L2 redirect pkt port not retrieved",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"trigger_stats_rate": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"threshold_exceeded_by": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
						},
						"duration": {
							Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
						},
						"so_pkts_slb_nat_reserve_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total SLB NAT reserve failures",
						},
						"so_pkts_slb_nat_release_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total SLB NAT release failures",
						},
						"so_pkts_l2redirect_dest_mac_zero_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Destination MAC Address zero Drop",
						},
						"so_pkts_l2redirect_interface_not_up": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L2redirect Intf is not UP",
						},
						"so_pkts_l2redirect_invalid_redirect_inf": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Redirect Table Error due to invalid redirect info",
						},
						"so_pkts_l3_redirect_encap_error_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Redirect encap error drop during transmission",
						},
						"so_pkts_l3_redirect_inner_mac_zero_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Redirect inner mac zero drop during transmission",
						},
						"so_pkts_l3_redirect_table_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Redirect Table error Drop",
						},
						"so_pkts_l3_redirect_fragmentation_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 redirect encap Fragmentation error",
						},
						"so_pkts_l3_redirect_table_no_entry_foun": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 redirect Table no redirect entry found error",
						},
						"so_pkts_l3_redirect_invalid_dev_dir": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Redirect Invalid Device direction during transmission",
						},
						"so_pkts_l3_redirect_chassis_dest_mac_er": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Redirect RX multi-slot Destination MAC Error",
						},
						"so_pkts_l2redirect_vlan_retrieval_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L2 redirect pkt vlan not retrieved",
						},
						"so_pkts_l2redirect_port_retrieval_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L2 redirect pkt port not retrieved",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCounters(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCounters(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCounters(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCounters(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsInc2087(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsInc2087 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsInc2087
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.So_pkts_slb_nat_reserve_fail = in["so_pkts_slb_nat_reserve_fail"].(int)
		ret.So_pkts_slb_nat_release_fail = in["so_pkts_slb_nat_release_fail"].(int)
		ret.So_pkts_l2redirect_dest_mac_zero_drop = in["so_pkts_l2redirect_dest_mac_zero_drop"].(int)
		ret.So_pkts_l2redirect_interface_not_up = in["so_pkts_l2redirect_interface_not_up"].(int)
		ret.So_pkts_l2redirect_invalid_redirect_inf = in["so_pkts_l2redirect_invalid_redirect_inf"].(int)
		ret.So_pkts_l3_redirect_encap_error_drop = in["so_pkts_l3_redirect_encap_error_drop"].(int)
		ret.So_pkts_l3_redirect_inner_mac_zero_drop = in["so_pkts_l3_redirect_inner_mac_zero_drop"].(int)
		ret.So_pkts_l3_redirect_table_error = in["so_pkts_l3_redirect_table_error"].(int)
		ret.So_pkts_l3_redirect_fragmentation_error = in["so_pkts_l3_redirect_fragmentation_error"].(int)
		ret.So_pkts_l3_redirect_table_no_entry_foun = in["so_pkts_l3_redirect_table_no_entry_foun"].(int)
		ret.So_pkts_l3_redirect_invalid_dev_dir = in["so_pkts_l3_redirect_invalid_dev_dir"].(int)
		ret.So_pkts_l3_redirect_chassis_dest_mac_er = in["so_pkts_l3_redirect_chassis_dest_mac_er"].(int)
		ret.So_pkts_l2redirect_vlan_retrieval_error = in["so_pkts_l2redirect_vlan_retrieval_error"].(int)
		ret.So_pkts_l2redirect_port_retrieval_error = in["so_pkts_l2redirect_port_retrieval_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsRate2088(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsRate2088 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsRate2088
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.So_pkts_slb_nat_reserve_fail = in["so_pkts_slb_nat_reserve_fail"].(int)
		ret.So_pkts_slb_nat_release_fail = in["so_pkts_slb_nat_release_fail"].(int)
		ret.So_pkts_l2redirect_dest_mac_zero_drop = in["so_pkts_l2redirect_dest_mac_zero_drop"].(int)
		ret.So_pkts_l2redirect_interface_not_up = in["so_pkts_l2redirect_interface_not_up"].(int)
		ret.So_pkts_l2redirect_invalid_redirect_inf = in["so_pkts_l2redirect_invalid_redirect_inf"].(int)
		ret.So_pkts_l3_redirect_encap_error_drop = in["so_pkts_l3_redirect_encap_error_drop"].(int)
		ret.So_pkts_l3_redirect_inner_mac_zero_drop = in["so_pkts_l3_redirect_inner_mac_zero_drop"].(int)
		ret.So_pkts_l3_redirect_table_error = in["so_pkts_l3_redirect_table_error"].(int)
		ret.So_pkts_l3_redirect_fragmentation_error = in["so_pkts_l3_redirect_fragmentation_error"].(int)
		ret.So_pkts_l3_redirect_table_no_entry_foun = in["so_pkts_l3_redirect_table_no_entry_foun"].(int)
		ret.So_pkts_l3_redirect_invalid_dev_dir = in["so_pkts_l3_redirect_invalid_dev_dir"].(int)
		ret.So_pkts_l3_redirect_chassis_dest_mac_er = in["so_pkts_l3_redirect_chassis_dest_mac_er"].(int)
		ret.So_pkts_l2redirect_vlan_retrieval_error = in["so_pkts_l2redirect_vlan_retrieval_error"].(int)
		ret.So_pkts_l2redirect_port_retrieval_error = in["so_pkts_l2redirect_port_retrieval_error"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCounters(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCounters {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCounters
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsInc2087(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsRate2088(d.Get("trigger_stats_rate").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
