package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Lsn() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_lsn`: Configure triggers for cgnv6.lsn.global object\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnDelete,

		Schema: map[string]*schema.Schema{
			"trigger_stats_inc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_quota_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Creation Failed",
						},
						"data_sesn_user_quota_exceeded": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data Session User-Quota Exceeded",
						},
						"fullcone_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Full-cone Session Creation Failed",
						},
						"fullcone_self_hairpinning_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Self-Hairpinning Drop",
						},
						"nat_pool_unusable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT Pool Unusable",
						},
						"ha_nat_pool_unusable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA NAT Pool Unusable",
						},
						"ha_nat_pool_batch_type_mismatch": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA NAT Pool Batch Type Mismatch",
						},
						"sip_alg_quota_inc_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG User-Quota Exceeded",
						},
						"sip_alg_alloc_rtp_rtcp_port_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG Alloc RTP/RTCP NAT Ports Failure",
						},
						"sip_alg_alloc_single_port_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG Alloc Single RTP or RTCP NAT Port Failure",
						},
						"sip_alg_create_single_fullcone_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG Create Single RTP or RTCP Full-cone Session Failure",
						},
						"sip_alg_create_rtp_fullcone_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG Create RTP Full-cone Session Failure",
						},
						"sip_alg_create_rtcp_fullcone_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG Create RTCP Full-cone Session Failure",
						},
						"h323_alg_alloc_single_port_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for H323 ALG Alloc Single RTP or RTCP NAT Port Failure",
						},
						"h323_alg_create_single_fullcone_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for H323 ALG Create Single RTP or RTCP Full-cone Session Failure",
						},
						"h323_alg_create_rtp_fullcone_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for H323 ALG Create RTP Full-cone Session Failure",
						},
						"h323_alg_create_rtcp_fullcone_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for H323 ALG Create RTCP Full-cone Session Failure",
						},
						"port_overloading_out_of_memory": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Overloading Out of Memory",
						},
						"port_overloading_inc_overflow": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Overloading Inc Overflow",
						},
						"fullcone_ext_mem_alloc_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for LSN Fullcone Extension Memory Allocate Failure",
						},
						"fullcone_ext_mem_alloc_init_faulure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for LSN Fullcone Extension Initialization Failure",
						},
						"mgcp_alg_create_rtp_fullcone_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for MGCP ALG Create RTP Full-cone Session Failure",
						},
						"mgcp_alg_create_rtcp_fullcone_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for MGCP ALG Create RTCP Full-cone Session Failure",
						},
						"mgcp_alg_port_pair_alloc_from_quota_par": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for MGCP ALG Port Pair Allocated From Quota Partition Error",
						},
						"user_quota_unusable_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Unusable Drop",
						},
						"user_quota_unusable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Marked Unusable",
						},
						"adc_port_allocation_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ADC Port Allocation Failed",
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
						"user_quota_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Creation Failed",
						},
						"data_sesn_user_quota_exceeded": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data Session User-Quota Exceeded",
						},
						"fullcone_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Full-cone Session Creation Failed",
						},
						"fullcone_self_hairpinning_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Self-Hairpinning Drop",
						},
						"nat_pool_unusable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT Pool Unusable",
						},
						"ha_nat_pool_unusable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA NAT Pool Unusable",
						},
						"ha_nat_pool_batch_type_mismatch": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA NAT Pool Batch Type Mismatch",
						},
						"sip_alg_quota_inc_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG User-Quota Exceeded",
						},
						"sip_alg_alloc_rtp_rtcp_port_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG Alloc RTP/RTCP NAT Ports Failure",
						},
						"sip_alg_alloc_single_port_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG Alloc Single RTP or RTCP NAT Port Failure",
						},
						"sip_alg_create_single_fullcone_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG Create Single RTP or RTCP Full-cone Session Failure",
						},
						"sip_alg_create_rtp_fullcone_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG Create RTP Full-cone Session Failure",
						},
						"sip_alg_create_rtcp_fullcone_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG Create RTCP Full-cone Session Failure",
						},
						"h323_alg_alloc_single_port_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for H323 ALG Alloc Single RTP or RTCP NAT Port Failure",
						},
						"h323_alg_create_single_fullcone_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for H323 ALG Create Single RTP or RTCP Full-cone Session Failure",
						},
						"h323_alg_create_rtp_fullcone_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for H323 ALG Create RTP Full-cone Session Failure",
						},
						"h323_alg_create_rtcp_fullcone_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for H323 ALG Create RTCP Full-cone Session Failure",
						},
						"port_overloading_out_of_memory": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Overloading Out of Memory",
						},
						"port_overloading_inc_overflow": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Overloading Inc Overflow",
						},
						"fullcone_ext_mem_alloc_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for LSN Fullcone Extension Memory Allocate Failure",
						},
						"fullcone_ext_mem_alloc_init_faulure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for LSN Fullcone Extension Initialization Failure",
						},
						"mgcp_alg_create_rtp_fullcone_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for MGCP ALG Create RTP Full-cone Session Failure",
						},
						"mgcp_alg_create_rtcp_fullcone_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for MGCP ALG Create RTCP Full-cone Session Failure",
						},
						"mgcp_alg_port_pair_alloc_from_quota_par": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for MGCP ALG Port Pair Allocated From Quota Partition Error",
						},
						"user_quota_unusable_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Unusable Drop",
						},
						"user_quota_unusable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Marked Unusable",
						},
						"adc_port_allocation_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ADC Port Allocation Failed",
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
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Lsn(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Lsn(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Lsn(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Lsn(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc1995(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc1995 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc1995
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.User_quota_failure = in["user_quota_failure"].(int)
		ret.Data_sesn_user_quota_exceeded = in["data_sesn_user_quota_exceeded"].(int)
		ret.Fullcone_failure = in["fullcone_failure"].(int)
		ret.Fullcone_self_hairpinning_drop = in["fullcone_self_hairpinning_drop"].(int)
		ret.Nat_pool_unusable = in["nat_pool_unusable"].(int)
		ret.Ha_nat_pool_unusable = in["ha_nat_pool_unusable"].(int)
		ret.Ha_nat_pool_batch_type_mismatch = in["ha_nat_pool_batch_type_mismatch"].(int)
		ret.Sip_alg_quota_inc_failure = in["sip_alg_quota_inc_failure"].(int)
		ret.Sip_alg_alloc_rtp_rtcp_port_failure = in["sip_alg_alloc_rtp_rtcp_port_failure"].(int)
		ret.Sip_alg_alloc_single_port_failure = in["sip_alg_alloc_single_port_failure"].(int)
		ret.Sip_alg_create_single_fullcone_failure = in["sip_alg_create_single_fullcone_failure"].(int)
		ret.Sip_alg_create_rtp_fullcone_failure = in["sip_alg_create_rtp_fullcone_failure"].(int)
		ret.Sip_alg_create_rtcp_fullcone_failure = in["sip_alg_create_rtcp_fullcone_failure"].(int)
		ret.H323_alg_alloc_single_port_failure = in["h323_alg_alloc_single_port_failure"].(int)
		ret.H323_alg_create_single_fullcone_failure = in["h323_alg_create_single_fullcone_failure"].(int)
		ret.H323_alg_create_rtp_fullcone_failure = in["h323_alg_create_rtp_fullcone_failure"].(int)
		ret.H323_alg_create_rtcp_fullcone_failure = in["h323_alg_create_rtcp_fullcone_failure"].(int)
		ret.Port_overloading_out_of_memory = in["port_overloading_out_of_memory"].(int)
		ret.Port_overloading_inc_overflow = in["port_overloading_inc_overflow"].(int)
		ret.Fullcone_ext_mem_alloc_failure = in["fullcone_ext_mem_alloc_failure"].(int)
		ret.Fullcone_ext_mem_alloc_init_faulure = in["fullcone_ext_mem_alloc_init_faulure"].(int)
		ret.Mgcp_alg_create_rtp_fullcone_failure = in["mgcp_alg_create_rtp_fullcone_failure"].(int)
		ret.Mgcp_alg_create_rtcp_fullcone_failure = in["mgcp_alg_create_rtcp_fullcone_failure"].(int)
		ret.Mgcp_alg_port_pair_alloc_from_quota_par = in["mgcp_alg_port_pair_alloc_from_quota_par"].(int)
		ret.User_quota_unusable_drop = in["user_quota_unusable_drop"].(int)
		ret.User_quota_unusable = in["user_quota_unusable"].(int)
		ret.Adc_port_allocation_failed = in["adc_port_allocation_failed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsRate1996(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsRate1996 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsRate1996
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.User_quota_failure = in["user_quota_failure"].(int)
		ret.Data_sesn_user_quota_exceeded = in["data_sesn_user_quota_exceeded"].(int)
		ret.Fullcone_failure = in["fullcone_failure"].(int)
		ret.Fullcone_self_hairpinning_drop = in["fullcone_self_hairpinning_drop"].(int)
		ret.Nat_pool_unusable = in["nat_pool_unusable"].(int)
		ret.Ha_nat_pool_unusable = in["ha_nat_pool_unusable"].(int)
		ret.Ha_nat_pool_batch_type_mismatch = in["ha_nat_pool_batch_type_mismatch"].(int)
		ret.Sip_alg_quota_inc_failure = in["sip_alg_quota_inc_failure"].(int)
		ret.Sip_alg_alloc_rtp_rtcp_port_failure = in["sip_alg_alloc_rtp_rtcp_port_failure"].(int)
		ret.Sip_alg_alloc_single_port_failure = in["sip_alg_alloc_single_port_failure"].(int)
		ret.Sip_alg_create_single_fullcone_failure = in["sip_alg_create_single_fullcone_failure"].(int)
		ret.Sip_alg_create_rtp_fullcone_failure = in["sip_alg_create_rtp_fullcone_failure"].(int)
		ret.Sip_alg_create_rtcp_fullcone_failure = in["sip_alg_create_rtcp_fullcone_failure"].(int)
		ret.H323_alg_alloc_single_port_failure = in["h323_alg_alloc_single_port_failure"].(int)
		ret.H323_alg_create_single_fullcone_failure = in["h323_alg_create_single_fullcone_failure"].(int)
		ret.H323_alg_create_rtp_fullcone_failure = in["h323_alg_create_rtp_fullcone_failure"].(int)
		ret.H323_alg_create_rtcp_fullcone_failure = in["h323_alg_create_rtcp_fullcone_failure"].(int)
		ret.Port_overloading_out_of_memory = in["port_overloading_out_of_memory"].(int)
		ret.Port_overloading_inc_overflow = in["port_overloading_inc_overflow"].(int)
		ret.Fullcone_ext_mem_alloc_failure = in["fullcone_ext_mem_alloc_failure"].(int)
		ret.Fullcone_ext_mem_alloc_init_faulure = in["fullcone_ext_mem_alloc_init_faulure"].(int)
		ret.Mgcp_alg_create_rtp_fullcone_failure = in["mgcp_alg_create_rtp_fullcone_failure"].(int)
		ret.Mgcp_alg_create_rtcp_fullcone_failure = in["mgcp_alg_create_rtcp_fullcone_failure"].(int)
		ret.Mgcp_alg_port_pair_alloc_from_quota_par = in["mgcp_alg_port_pair_alloc_from_quota_par"].(int)
		ret.User_quota_unusable_drop = in["user_quota_unusable_drop"].(int)
		ret.User_quota_unusable = in["user_quota_unusable"].(int)
		ret.Adc_port_allocation_failed = in["adc_port_allocation_failed"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Lsn(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Lsn {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Lsn
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc1995(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsRate1996(d.Get("trigger_stats_rate").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
