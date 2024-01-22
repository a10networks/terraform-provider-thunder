package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_fixed_nat_global_trigger_stats_rate`: Configure stats to trigger packet capture on increment rate\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRateCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRateUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRateRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRateDelete,

		Schema: map[string]*schema.Schema{
			"config_not_found": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fixed NAT Config not Found",
			},
			"dest_rlist_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fixed NAT Dest Rule List Drop",
			},
			"dest_rlist_pass_through": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fixed NAT Dest Rule List Pass-Through",
			},
			"dest_rlist_snat_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fixed NAT Dest Rules List Source NAT Drop",
			},
			"dslite_eif_limit_exceeded": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DS-Lite Endpoint-Independent-Filtering Limit Exceeded",
			},
			"dslite_inbound_filtered": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DS-Lite Endpoint-Dependent Filtering Drop",
			},
			"duration": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
			},
			"fixed_nat_fullcone_self_hairpinning_dro": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Self-Hairpinning Drop",
			},
			"fullcone_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Full-Cone Session Creation Failed",
			},
			"ha_session_user_quota_exceeded": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA Sessions User Quota Exceeded",
			},
			"nat_port_unavailable_icmp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICMP NAT Port Unavailable",
			},
			"nat_port_unavailable_tcp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP NAT Port Unavailable",
			},
			"nat_port_unavailable_udp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP NAT Port Unavailable",
			},
			"nat44_eif_limit_exceeded": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT44 Endpoint-Independent-Filtering Limit Exceeded",
			},
			"nat44_inbound_filtered": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT44 Endpoint-Dependent Filtering Drop",
			},
			"nat64_eif_limit_exceeded": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT64 Endpoint-Independent-Filtering Limit Exceeded",
			},
			"nat64_inbound_filtered": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT64 Endpoint-Dependent Filtering Drop",
			},
			"port_overload_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port overload failed",
			},
			"session_user_quota_exceeded": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Sessions User Quota Exceeded",
			},
			"sixrd_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fixed NAT IPv6 in IPv4 Packet Drop",
			},
			"standby_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fixed NAT LID Standby Drop",
			},
			"threshold_exceeded_by": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
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
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRateRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRateRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRate(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRate {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRate
	ret.Inst.ConfigNotFound = d.Get("config_not_found").(int)
	ret.Inst.DestRlistDrop = d.Get("dest_rlist_drop").(int)
	ret.Inst.DestRlistPassThrough = d.Get("dest_rlist_pass_through").(int)
	ret.Inst.DestRlistSnatDrop = d.Get("dest_rlist_snat_drop").(int)
	ret.Inst.DsliteEifLimitExceeded = d.Get("dslite_eif_limit_exceeded").(int)
	ret.Inst.DsliteInboundFiltered = d.Get("dslite_inbound_filtered").(int)
	ret.Inst.Duration = d.Get("duration").(int)
	ret.Inst.FixedNatFullconeSelfHairpinningDro = d.Get("fixed_nat_fullcone_self_hairpinning_dro").(int)
	ret.Inst.FullconeFailure = d.Get("fullcone_failure").(int)
	ret.Inst.HaSessionUserQuotaExceeded = d.Get("ha_session_user_quota_exceeded").(int)
	ret.Inst.NatPortUnavailableIcmp = d.Get("nat_port_unavailable_icmp").(int)
	ret.Inst.NatPortUnavailableTcp = d.Get("nat_port_unavailable_tcp").(int)
	ret.Inst.NatPortUnavailableUdp = d.Get("nat_port_unavailable_udp").(int)
	ret.Inst.Nat44EifLimitExceeded = d.Get("nat44_eif_limit_exceeded").(int)
	ret.Inst.Nat44InboundFiltered = d.Get("nat44_inbound_filtered").(int)
	ret.Inst.Nat64EifLimitExceeded = d.Get("nat64_eif_limit_exceeded").(int)
	ret.Inst.Nat64InboundFiltered = d.Get("nat64_inbound_filtered").(int)
	ret.Inst.PortOverloadFailed = d.Get("port_overload_failed").(int)
	ret.Inst.SessionUserQuotaExceeded = d.Get("session_user_quota_exceeded").(int)
	ret.Inst.SixrdDrop = d.Get("sixrd_drop").(int)
	ret.Inst.StandbyDrop = d.Get("standby_drop").(int)
	ret.Inst.ThresholdExceededBy = d.Get("threshold_exceeded_by").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
