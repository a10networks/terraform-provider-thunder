package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Pcp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_pcp`: Configure triggers for cgnv6.pcp object\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpDelete,

		Schema: map[string]*schema.Schema{
			"trigger_stats_inc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pkt_not_request_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packet Not a PCP Request",
						},
						"pkt_too_short_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packet Too Short",
						},
						"noroute_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response No Route",
						},
						"unsupported_version": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PCP version",
						},
						"not_authorized": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for PCP Request Not Authorized",
						},
						"malform_request": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for PCP Request Malformed",
						},
						"unsupp_opcode": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PCP Opcode",
						},
						"unsupp_option": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PCP Option",
						},
						"malform_option": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for PCP Option Malformed",
						},
						"no_resources": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No System or NAT Resources",
						},
						"unsupp_protocol": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported Mapping Protocol",
						},
						"cannot_provide_suggest": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cannot Provide Suggested Port When PREFER_FAILURE",
						},
						"address_mismatch": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for PCP Client Address Mismatch",
						},
						"excessive_remote_peers": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Excessive Remote Peers",
						},
						"pkt_not_from_nat_inside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packet Dropped For Not Coming From NAT Inside",
						},
						"l4_process_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3/L4 Process Error",
						},
						"internal_error_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Internal Error",
						},
						"unsol_ance_sent_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsolicited Announce Send Failure",
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
						"pkt_not_request_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packet Not a PCP Request",
						},
						"pkt_too_short_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packet Too Short",
						},
						"noroute_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response No Route",
						},
						"unsupported_version": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PCP version",
						},
						"not_authorized": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for PCP Request Not Authorized",
						},
						"malform_request": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for PCP Request Malformed",
						},
						"unsupp_opcode": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PCP Opcode",
						},
						"unsupp_option": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PCP Option",
						},
						"malform_option": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for PCP Option Malformed",
						},
						"no_resources": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No System or NAT Resources",
						},
						"unsupp_protocol": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported Mapping Protocol",
						},
						"cannot_provide_suggest": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cannot Provide Suggested Port When PREFER_FAILURE",
						},
						"address_mismatch": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for PCP Client Address Mismatch",
						},
						"excessive_remote_peers": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Excessive Remote Peers",
						},
						"pkt_not_from_nat_inside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packet Dropped For Not Coming From NAT Inside",
						},
						"l4_process_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3/L4 Process Error",
						},
						"internal_error_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Internal Error",
						},
						"unsol_ance_sent_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsolicited Announce Send Failure",
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
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Pcp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Pcp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Pcp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Pcp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpTriggerStatsInc1999(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpTriggerStatsInc1999 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpTriggerStatsInc1999
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PktNotRequestDrop = in["pkt_not_request_drop"].(int)
		ret.PktTooShortDrop = in["pkt_too_short_drop"].(int)
		ret.NorouteDrop = in["noroute_drop"].(int)
		ret.UnsupportedVersion = in["unsupported_version"].(int)
		ret.NotAuthorized = in["not_authorized"].(int)
		ret.MalformRequest = in["malform_request"].(int)
		ret.UnsuppOpcode = in["unsupp_opcode"].(int)
		ret.UnsuppOption = in["unsupp_option"].(int)
		ret.MalformOption = in["malform_option"].(int)
		ret.NoResources = in["no_resources"].(int)
		ret.UnsuppProtocol = in["unsupp_protocol"].(int)
		ret.CannotProvideSuggest = in["cannot_provide_suggest"].(int)
		ret.AddressMismatch = in["address_mismatch"].(int)
		ret.ExcessiveRemotePeers = in["excessive_remote_peers"].(int)
		ret.PktNotFromNatInside = in["pkt_not_from_nat_inside"].(int)
		ret.L4ProcessError = in["l4_process_error"].(int)
		ret.InternalErrorDrop = in["internal_error_drop"].(int)
		ret.Unsol_ance_sent_fail = in["unsol_ance_sent_fail"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpTriggerStatsRate2000(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpTriggerStatsRate2000 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpTriggerStatsRate2000
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.PktNotRequestDrop = in["pkt_not_request_drop"].(int)
		ret.PktTooShortDrop = in["pkt_too_short_drop"].(int)
		ret.NorouteDrop = in["noroute_drop"].(int)
		ret.UnsupportedVersion = in["unsupported_version"].(int)
		ret.NotAuthorized = in["not_authorized"].(int)
		ret.MalformRequest = in["malform_request"].(int)
		ret.UnsuppOpcode = in["unsupp_opcode"].(int)
		ret.UnsuppOption = in["unsupp_option"].(int)
		ret.MalformOption = in["malform_option"].(int)
		ret.NoResources = in["no_resources"].(int)
		ret.UnsuppProtocol = in["unsupp_protocol"].(int)
		ret.CannotProvideSuggest = in["cannot_provide_suggest"].(int)
		ret.AddressMismatch = in["address_mismatch"].(int)
		ret.ExcessiveRemotePeers = in["excessive_remote_peers"].(int)
		ret.PktNotFromNatInside = in["pkt_not_from_nat_inside"].(int)
		ret.L4ProcessError = in["l4_process_error"].(int)
		ret.InternalErrorDrop = in["internal_error_drop"].(int)
		ret.Unsol_ance_sent_fail = in["unsol_ance_sent_fail"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Pcp(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Pcp {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Pcp
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpTriggerStatsInc1999(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpTriggerStatsRate2000(d.Get("trigger_stats_rate").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
