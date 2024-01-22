package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsInc() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_so_counters_trigger_stats_inc`: Configure stats to trigger packet capture on increment\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsIncCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsIncUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsIncRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsIncDelete,

		Schema: map[string]*schema.Schema{
			"so_pkts_l2redirect_dest_mac_zero_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Destination MAC Address zero Drop",
			},
			"so_pkts_l2redirect_interface_not_up": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L2redirect Intf is not UP",
			},
			"so_pkts_l2redirect_invalid_redirect_inf": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Redirect Table Error due to invalid redirect info",
			},
			"so_pkts_l2redirect_port_retrieval_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L2 redirect pkt port not retrieved",
			},
			"so_pkts_l2redirect_vlan_retrieval_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L2 redirect pkt vlan not retrieved",
			},
			"so_pkts_l3_redirect_chassis_dest_mac_er": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Redirect RX multi-slot Destination MAC Error",
			},
			"so_pkts_l3_redirect_encap_error_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Redirect encap error drop during transmission",
			},
			"so_pkts_l3_redirect_fragmentation_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 redirect encap Fragmentation error",
			},
			"so_pkts_l3_redirect_inner_mac_zero_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Redirect inner mac zero drop during transmission",
			},
			"so_pkts_l3_redirect_invalid_dev_dir": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Redirect Invalid Device direction during transmission",
			},
			"so_pkts_l3_redirect_table_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Redirect Table error Drop",
			},
			"so_pkts_l3_redirect_table_no_entry_foun": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 redirect Table no redirect entry found error",
			},
			"so_pkts_slb_nat_release_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total SLB NAT release failures",
			},
			"so_pkts_slb_nat_reserve_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total SLB NAT reserve failures",
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
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsIncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsIncCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsInc(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsIncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsIncUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsInc(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsIncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsIncDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsInc(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsIncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsIncRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsInc(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsInc(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsInc {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsInc
	ret.Inst.So_pkts_l2redirect_dest_mac_zero_drop = d.Get("so_pkts_l2redirect_dest_mac_zero_drop").(int)
	ret.Inst.So_pkts_l2redirect_interface_not_up = d.Get("so_pkts_l2redirect_interface_not_up").(int)
	ret.Inst.So_pkts_l2redirect_invalid_redirect_inf = d.Get("so_pkts_l2redirect_invalid_redirect_inf").(int)
	ret.Inst.So_pkts_l2redirect_port_retrieval_error = d.Get("so_pkts_l2redirect_port_retrieval_error").(int)
	ret.Inst.So_pkts_l2redirect_vlan_retrieval_error = d.Get("so_pkts_l2redirect_vlan_retrieval_error").(int)
	ret.Inst.So_pkts_l3_redirect_chassis_dest_mac_er = d.Get("so_pkts_l3_redirect_chassis_dest_mac_er").(int)
	ret.Inst.So_pkts_l3_redirect_encap_error_drop = d.Get("so_pkts_l3_redirect_encap_error_drop").(int)
	ret.Inst.So_pkts_l3_redirect_fragmentation_error = d.Get("so_pkts_l3_redirect_fragmentation_error").(int)
	ret.Inst.So_pkts_l3_redirect_inner_mac_zero_drop = d.Get("so_pkts_l3_redirect_inner_mac_zero_drop").(int)
	ret.Inst.So_pkts_l3_redirect_invalid_dev_dir = d.Get("so_pkts_l3_redirect_invalid_dev_dir").(int)
	ret.Inst.So_pkts_l3_redirect_table_error = d.Get("so_pkts_l3_redirect_table_error").(int)
	ret.Inst.So_pkts_l3_redirect_table_no_entry_foun = d.Get("so_pkts_l3_redirect_table_no_entry_foun").(int)
	ret.Inst.So_pkts_slb_nat_release_fail = d.Get("so_pkts_slb_nat_release_fail").(int)
	ret.Inst.So_pkts_slb_nat_reserve_fail = d.Get("so_pkts_slb_nat_reserve_fail").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
