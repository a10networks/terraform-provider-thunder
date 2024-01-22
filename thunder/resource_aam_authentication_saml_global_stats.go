package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationSamlGlobalStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_saml_global_stats`: Statistics for the object global\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationSamlGlobalStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"requests_to_a10saml": {
							Type: schema.TypeInt, Optional: true, Description: "Total Request to A10 SAML Service",
						},
						"responses_from_a10saml": {
							Type: schema.TypeInt, Optional: true, Description: "Total Response from A10 SAML Service",
						},
						"sp_metadata_export_req": {
							Type: schema.TypeInt, Optional: true, Description: "Total Metadata Export Request",
						},
						"sp_metadata_export_success": {
							Type: schema.TypeInt, Optional: true, Description: "Toal Metadata Export Success",
						},
						"login_auth_req": {
							Type: schema.TypeInt, Optional: true, Description: "Total Login Authentication Request",
						},
						"login_auth_resp": {
							Type: schema.TypeInt, Optional: true, Description: "Total Login Authentication Response",
						},
						"acs_req": {
							Type: schema.TypeInt, Optional: true, Description: "Total SAML Single-Sign-On Request",
						},
						"acs_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total SAML Single-Sign-On Success",
						},
						"acs_authz_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Total SAML Single-Sign-On Authorization Fail",
						},
						"acs_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total SAML Single-Sign-On Error",
						},
						"slo_req": {
							Type: schema.TypeInt, Optional: true, Description: "Total Single Logout Request",
						},
						"slo_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total Single Logout Success",
						},
						"slo_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total Single Logout Error",
						},
						"sp_slo_req": {
							Type: schema.TypeInt, Optional: true, Description: "Total SP-initiated Single Logout Request",
						},
						"glo_slo_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total Global Logout Success",
						},
						"loc_slo_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total Local Logout Success",
						},
						"par_slo_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total Partial Logout Success",
						},
						"relay_req": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"relay_success": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"relay_fail": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"relay_error": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"other_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total Other Error",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationSamlGlobalStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationSamlGlobalStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationSamlGlobalStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationSamlGlobalStatsStats := setObjectAamAuthenticationSamlGlobalStatsStats(res)
		d.Set("stats", AamAuthenticationSamlGlobalStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationSamlGlobalStatsStats(ret edpt.DataAamAuthenticationSamlGlobalStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"requests_to_a10saml":        ret.DtAamAuthenticationSamlGlobalStats.Stats.RequestsToA10saml,
			"responses_from_a10saml":     ret.DtAamAuthenticationSamlGlobalStats.Stats.ResponsesFromA10saml,
			"sp_metadata_export_req":     ret.DtAamAuthenticationSamlGlobalStats.Stats.SpMetadataExportReq,
			"sp_metadata_export_success": ret.DtAamAuthenticationSamlGlobalStats.Stats.SpMetadataExportSuccess,
			"login_auth_req":             ret.DtAamAuthenticationSamlGlobalStats.Stats.LoginAuthReq,
			"login_auth_resp":            ret.DtAamAuthenticationSamlGlobalStats.Stats.LoginAuthResp,
			"acs_req":                    ret.DtAamAuthenticationSamlGlobalStats.Stats.AcsReq,
			"acs_success":                ret.DtAamAuthenticationSamlGlobalStats.Stats.AcsSuccess,
			"acs_authz_fail":             ret.DtAamAuthenticationSamlGlobalStats.Stats.AcsAuthzFail,
			"acs_error":                  ret.DtAamAuthenticationSamlGlobalStats.Stats.AcsError,
			"slo_req":                    ret.DtAamAuthenticationSamlGlobalStats.Stats.SloReq,
			"slo_success":                ret.DtAamAuthenticationSamlGlobalStats.Stats.SloSuccess,
			"slo_error":                  ret.DtAamAuthenticationSamlGlobalStats.Stats.SloError,
			"sp_slo_req":                 ret.DtAamAuthenticationSamlGlobalStats.Stats.SpSloReq,
			"glo_slo_success":            ret.DtAamAuthenticationSamlGlobalStats.Stats.GloSloSuccess,
			"loc_slo_success":            ret.DtAamAuthenticationSamlGlobalStats.Stats.LocSloSuccess,
			"par_slo_success":            ret.DtAamAuthenticationSamlGlobalStats.Stats.ParSloSuccess,
			"relay_req":                  ret.DtAamAuthenticationSamlGlobalStats.Stats.RelayReq,
			"relay_success":              ret.DtAamAuthenticationSamlGlobalStats.Stats.RelaySuccess,
			"relay_fail":                 ret.DtAamAuthenticationSamlGlobalStats.Stats.RelayFail,
			"relay_error":                ret.DtAamAuthenticationSamlGlobalStats.Stats.RelayError,
			"other_error":                ret.DtAamAuthenticationSamlGlobalStats.Stats.OtherError,
		},
	}
}

func getObjectAamAuthenticationSamlGlobalStatsStats(d []interface{}) edpt.AamAuthenticationSamlGlobalStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationSamlGlobalStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RequestsToA10saml = in["requests_to_a10saml"].(int)
		ret.ResponsesFromA10saml = in["responses_from_a10saml"].(int)
		ret.SpMetadataExportReq = in["sp_metadata_export_req"].(int)
		ret.SpMetadataExportSuccess = in["sp_metadata_export_success"].(int)
		ret.LoginAuthReq = in["login_auth_req"].(int)
		ret.LoginAuthResp = in["login_auth_resp"].(int)
		ret.AcsReq = in["acs_req"].(int)
		ret.AcsSuccess = in["acs_success"].(int)
		ret.AcsAuthzFail = in["acs_authz_fail"].(int)
		ret.AcsError = in["acs_error"].(int)
		ret.SloReq = in["slo_req"].(int)
		ret.SloSuccess = in["slo_success"].(int)
		ret.SloError = in["slo_error"].(int)
		ret.SpSloReq = in["sp_slo_req"].(int)
		ret.GloSloSuccess = in["glo_slo_success"].(int)
		ret.LocSloSuccess = in["loc_slo_success"].(int)
		ret.ParSloSuccess = in["par_slo_success"].(int)
		ret.RelayReq = in["relay_req"].(int)
		ret.RelaySuccess = in["relay_success"].(int)
		ret.RelayFail = in["relay_fail"].(int)
		ret.RelayError = in["relay_error"].(int)
		ret.OtherError = in["other_error"].(int)
	}
	return ret
}

func dataToEndpointAamAuthenticationSamlGlobalStats(d *schema.ResourceData) edpt.AamAuthenticationSamlGlobalStats {
	var ret edpt.AamAuthenticationSamlGlobalStats

	ret.Stats = getObjectAamAuthenticationSamlGlobalStatsStats(d.Get("stats").([]interface{}))
	return ret
}
