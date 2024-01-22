package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamJwtAuthorizationStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_jwt_authorization_stats`: Statistics for the object jwt-authorization\n\n__PLACEHOLDER__",
		ReadContext: resourceAamJwtAuthorizationStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify JWT authorization template name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"jwt_request": {
							Type: schema.TypeInt, Optional: true, Description: "JWT Request",
						},
						"jwt_authorize_success": {
							Type: schema.TypeInt, Optional: true, Description: "JWT Authorize Success",
						},
						"jwt_authorize_failure": {
							Type: schema.TypeInt, Optional: true, Description: "JWT Authorize Failure",
						},
						"jwt_missing_token": {
							Type: schema.TypeInt, Optional: true, Description: "JWT Missing Token",
						},
						"jwt_missing_claim": {
							Type: schema.TypeInt, Optional: true, Description: "JWT Missing Claim",
						},
						"jwt_token_expired": {
							Type: schema.TypeInt, Optional: true, Description: "JWT Token Expired",
						},
						"jwt_signature_failure": {
							Type: schema.TypeInt, Optional: true, Description: "JWT Signature Failure",
						},
						"jwt_other_error": {
							Type: schema.TypeInt, Optional: true, Description: "JWT Other Error",
						},
					},
				},
			},
		},
	}
}

func resourceAamJwtAuthorizationStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamJwtAuthorizationStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamJwtAuthorizationStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamJwtAuthorizationStatsStats := setObjectAamJwtAuthorizationStatsStats(res)
		d.Set("stats", AamJwtAuthorizationStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamJwtAuthorizationStatsStats(ret edpt.DataAamJwtAuthorizationStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"jwt_request":           ret.DtAamJwtAuthorizationStats.Stats.JwtRequest,
			"jwt_authorize_success": ret.DtAamJwtAuthorizationStats.Stats.JwtAuthorizeSuccess,
			"jwt_authorize_failure": ret.DtAamJwtAuthorizationStats.Stats.JwtAuthorizeFailure,
			"jwt_missing_token":     ret.DtAamJwtAuthorizationStats.Stats.JwtMissingToken,
			"jwt_missing_claim":     ret.DtAamJwtAuthorizationStats.Stats.JwtMissingClaim,
			"jwt_token_expired":     ret.DtAamJwtAuthorizationStats.Stats.JwtTokenExpired,
			"jwt_signature_failure": ret.DtAamJwtAuthorizationStats.Stats.JwtSignatureFailure,
			"jwt_other_error":       ret.DtAamJwtAuthorizationStats.Stats.JwtOtherError,
		},
	}
}

func getObjectAamJwtAuthorizationStatsStats(d []interface{}) edpt.AamJwtAuthorizationStatsStats {

	count1 := len(d)
	var ret edpt.AamJwtAuthorizationStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.JwtRequest = in["jwt_request"].(int)
		ret.JwtAuthorizeSuccess = in["jwt_authorize_success"].(int)
		ret.JwtAuthorizeFailure = in["jwt_authorize_failure"].(int)
		ret.JwtMissingToken = in["jwt_missing_token"].(int)
		ret.JwtMissingClaim = in["jwt_missing_claim"].(int)
		ret.JwtTokenExpired = in["jwt_token_expired"].(int)
		ret.JwtSignatureFailure = in["jwt_signature_failure"].(int)
		ret.JwtOtherError = in["jwt_other_error"].(int)
	}
	return ret
}

func dataToEndpointAamJwtAuthorizationStats(d *schema.ResourceData) edpt.AamJwtAuthorizationStats {
	var ret edpt.AamJwtAuthorizationStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectAamJwtAuthorizationStatsStats(d.Get("stats").([]interface{}))
	return ret
}
