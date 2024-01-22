package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_aam_jwt_authorization_tmpl_trigger_stats_rate`: Configure stats to triggers packet capture on increment\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRateCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRateUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRateRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRateDelete,

		Schema: map[string]*schema.Schema{
			"duration": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
			},
			"jwt_authorize_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Authorize Failure",
			},
			"jwt_missing_claim": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Missing Claim",
			},
			"jwt_missing_token": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Missing Token",
			},
			"jwt_other_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Other Error",
			},
			"jwt_signature_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Signature Failure",
			},
			"jwt_token_expired": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Token Expired",
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
func resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRateRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRateRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRate(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRate {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsRate
	ret.Inst.Duration = d.Get("duration").(int)
	ret.Inst.JwtAuthorizeFailure = d.Get("jwt_authorize_failure").(int)
	ret.Inst.JwtMissingClaim = d.Get("jwt_missing_claim").(int)
	ret.Inst.JwtMissingToken = d.Get("jwt_missing_token").(int)
	ret.Inst.JwtOtherError = d.Get("jwt_other_error").(int)
	ret.Inst.JwtSignatureFailure = d.Get("jwt_signature_failure").(int)
	ret.Inst.JwtTokenExpired = d.Get("jwt_token_expired").(int)
	ret.Inst.ThresholdExceededBy = d.Get("threshold_exceeded_by").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
