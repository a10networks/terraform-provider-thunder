package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmpl() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_aam_jwt_authorization_tmpl`: Configure template for aam.jwt-authorization\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplDelete,

		Schema: map[string]*schema.Schema{
			"capture_config": {
				Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
			},
			"trigger_stats_inc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"jwt_authorize_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Authorize Failure",
						},
						"jwt_missing_token": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Missing Token",
						},
						"jwt_missing_claim": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Missing Claim",
						},
						"jwt_token_expired": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Token Expired",
						},
						"jwt_signature_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Signature Failure",
						},
						"jwt_other_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Other Error",
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
						"jwt_authorize_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Authorize Failure",
						},
						"jwt_missing_token": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Missing Token",
						},
						"jwt_missing_claim": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Missing Claim",
						},
						"jwt_token_expired": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Token Expired",
						},
						"jwt_signature_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Signature Failure",
						},
						"jwt_other_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Other Error",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"trigger_stats_severity": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
						},
						"error_alert": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
						},
						"error_warning": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
						},
						"error_critical": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
						},
						"drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
						},
						"drop_alert": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
						},
						"drop_warning": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
						},
						"drop_critical": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmpl(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmpl(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmpl(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmpl(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsInc2663(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsInc2663 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsInc2663
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.JwtAuthorizeFailure = in["jwt_authorize_failure"].(int)
		ret.JwtMissingToken = in["jwt_missing_token"].(int)
		ret.JwtMissingClaim = in["jwt_missing_claim"].(int)
		ret.JwtTokenExpired = in["jwt_token_expired"].(int)
		ret.JwtSignatureFailure = in["jwt_signature_failure"].(int)
		ret.JwtOtherError = in["jwt_other_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRate2664(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRate2664 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRate2664
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.JwtAuthorizeFailure = in["jwt_authorize_failure"].(int)
		ret.JwtMissingToken = in["jwt_missing_token"].(int)
		ret.JwtMissingClaim = in["jwt_missing_claim"].(int)
		ret.JwtTokenExpired = in["jwt_token_expired"].(int)
		ret.JwtSignatureFailure = in["jwt_signature_failure"].(int)
		ret.JwtOtherError = in["jwt_other_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsSeverity2665(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsSeverity2665 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsSeverity2665
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmpl(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmpl {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmpl
	ret.Inst.CaptureConfig = d.Get("capture_config").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsInc2663(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRate2664(d.Get("trigger_stats_rate").([]interface{}))
	ret.Inst.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsSeverity2665(d.Get("trigger_stats_severity").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
