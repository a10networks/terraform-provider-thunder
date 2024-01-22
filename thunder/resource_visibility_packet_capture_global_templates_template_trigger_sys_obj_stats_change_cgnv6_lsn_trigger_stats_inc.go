package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_lsn_trigger_stats_inc`: Configure stats to trigger packet capture on increment\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsIncCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsIncUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsIncRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsIncDelete,

		Schema: map[string]*schema.Schema{
			"adc_port_allocation_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ADC Port Allocation Failed",
			},
			"data_sesn_user_quota_exceeded": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data Session User-Quota Exceeded",
			},
			"fullcone_ext_mem_alloc_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for LSN Fullcone Extension Memory Allocate Failure",
			},
			"fullcone_ext_mem_alloc_init_faulure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for LSN Fullcone Extension Initialization Failure",
			},
			"fullcone_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Full-cone Session Creation Failed",
			},
			"fullcone_self_hairpinning_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Self-Hairpinning Drop",
			},
			"h323_alg_alloc_single_port_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for H323 ALG Alloc Single RTP or RTCP NAT Port Failure",
			},
			"h323_alg_create_rtcp_fullcone_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for H323 ALG Create RTCP Full-cone Session Failure",
			},
			"h323_alg_create_rtp_fullcone_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for H323 ALG Create RTP Full-cone Session Failure",
			},
			"h323_alg_create_single_fullcone_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for H323 ALG Create Single RTP or RTCP Full-cone Session Failure",
			},
			"ha_nat_pool_batch_type_mismatch": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA NAT Pool Batch Type Mismatch",
			},
			"ha_nat_pool_unusable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA NAT Pool Unusable",
			},
			"mgcp_alg_create_rtcp_fullcone_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for MGCP ALG Create RTCP Full-cone Session Failure",
			},
			"mgcp_alg_create_rtp_fullcone_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for MGCP ALG Create RTP Full-cone Session Failure",
			},
			"mgcp_alg_port_pair_alloc_from_quota_par": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for MGCP ALG Port Pair Allocated From Quota Partition Error",
			},
			"nat_pool_unusable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT Pool Unusable",
			},
			"port_overloading_inc_overflow": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Overloading Inc Overflow",
			},
			"port_overloading_out_of_memory": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Overloading Out of Memory",
			},
			"sip_alg_alloc_rtp_rtcp_port_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG Alloc RTP/RTCP NAT Ports Failure",
			},
			"sip_alg_alloc_single_port_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG Alloc Single RTP or RTCP NAT Port Failure",
			},
			"sip_alg_create_rtcp_fullcone_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG Create RTCP Full-cone Session Failure",
			},
			"sip_alg_create_rtp_fullcone_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG Create RTP Full-cone Session Failure",
			},
			"sip_alg_create_single_fullcone_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG Create Single RTP or RTCP Full-cone Session Failure",
			},
			"sip_alg_quota_inc_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG User-Quota Exceeded",
			},
			"user_quota_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Creation Failed",
			},
			"user_quota_unusable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Marked Unusable",
			},
			"user_quota_unusable_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Unusable Drop",
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
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsIncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsIncCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsIncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsIncUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsIncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsIncDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsIncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsIncRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc
	ret.Inst.Adc_port_allocation_failed = d.Get("adc_port_allocation_failed").(int)
	ret.Inst.Data_sesn_user_quota_exceeded = d.Get("data_sesn_user_quota_exceeded").(int)
	ret.Inst.Fullcone_ext_mem_alloc_failure = d.Get("fullcone_ext_mem_alloc_failure").(int)
	ret.Inst.Fullcone_ext_mem_alloc_init_faulure = d.Get("fullcone_ext_mem_alloc_init_faulure").(int)
	ret.Inst.Fullcone_failure = d.Get("fullcone_failure").(int)
	ret.Inst.Fullcone_self_hairpinning_drop = d.Get("fullcone_self_hairpinning_drop").(int)
	ret.Inst.H323_alg_alloc_single_port_failure = d.Get("h323_alg_alloc_single_port_failure").(int)
	ret.Inst.H323_alg_create_rtcp_fullcone_failure = d.Get("h323_alg_create_rtcp_fullcone_failure").(int)
	ret.Inst.H323_alg_create_rtp_fullcone_failure = d.Get("h323_alg_create_rtp_fullcone_failure").(int)
	ret.Inst.H323_alg_create_single_fullcone_failure = d.Get("h323_alg_create_single_fullcone_failure").(int)
	ret.Inst.Ha_nat_pool_batch_type_mismatch = d.Get("ha_nat_pool_batch_type_mismatch").(int)
	ret.Inst.Ha_nat_pool_unusable = d.Get("ha_nat_pool_unusable").(int)
	ret.Inst.Mgcp_alg_create_rtcp_fullcone_failure = d.Get("mgcp_alg_create_rtcp_fullcone_failure").(int)
	ret.Inst.Mgcp_alg_create_rtp_fullcone_failure = d.Get("mgcp_alg_create_rtp_fullcone_failure").(int)
	ret.Inst.Mgcp_alg_port_pair_alloc_from_quota_par = d.Get("mgcp_alg_port_pair_alloc_from_quota_par").(int)
	ret.Inst.Nat_pool_unusable = d.Get("nat_pool_unusable").(int)
	ret.Inst.Port_overloading_inc_overflow = d.Get("port_overloading_inc_overflow").(int)
	ret.Inst.Port_overloading_out_of_memory = d.Get("port_overloading_out_of_memory").(int)
	ret.Inst.Sip_alg_alloc_rtp_rtcp_port_failure = d.Get("sip_alg_alloc_rtp_rtcp_port_failure").(int)
	ret.Inst.Sip_alg_alloc_single_port_failure = d.Get("sip_alg_alloc_single_port_failure").(int)
	ret.Inst.Sip_alg_create_rtcp_fullcone_failure = d.Get("sip_alg_create_rtcp_fullcone_failure").(int)
	ret.Inst.Sip_alg_create_rtp_fullcone_failure = d.Get("sip_alg_create_rtp_fullcone_failure").(int)
	ret.Inst.Sip_alg_create_single_fullcone_failure = d.Get("sip_alg_create_single_fullcone_failure").(int)
	ret.Inst.Sip_alg_quota_inc_failure = d.Get("sip_alg_quota_inc_failure").(int)
	ret.Inst.User_quota_failure = d.Get("user_quota_failure").(int)
	ret.Inst.User_quota_unusable = d.Get("user_quota_unusable").(int)
	ret.Inst.User_quota_unusable_drop = d.Get("user_quota_unusable_drop").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
