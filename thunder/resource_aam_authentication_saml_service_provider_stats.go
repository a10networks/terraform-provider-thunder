package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationSamlServiceProviderStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_saml_service_provider_stats`: Statistics for the object service-provider\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationSamlServiceProviderStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify SAML authentication service provider name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sp_metadata_export_req": {
							Type: schema.TypeInt, Optional: true, Description: "Metadata Export Request",
						},
						"sp_metadata_export_success": {
							Type: schema.TypeInt, Optional: true, Description: "Metadata Export Success",
						},
						"login_auth_req": {
							Type: schema.TypeInt, Optional: true, Description: "Login Authentication Request",
						},
						"login_auth_resp": {
							Type: schema.TypeInt, Optional: true, Description: "Login Authentication Response",
						},
						"acs_req": {
							Type: schema.TypeInt, Optional: true, Description: "SAML Single-Sign-On Request",
						},
						"acs_success": {
							Type: schema.TypeInt, Optional: true, Description: "SAML Single-Sign-On Success",
						},
						"acs_authz_fail": {
							Type: schema.TypeInt, Optional: true, Description: "SAML Single-Sign-On Authorization Fail",
						},
						"acs_error": {
							Type: schema.TypeInt, Optional: true, Description: "SAML Single-Sign-On Error",
						},
						"slo_req": {
							Type: schema.TypeInt, Optional: true, Description: "Single Logout Request",
						},
						"slo_success": {
							Type: schema.TypeInt, Optional: true, Description: "Single Logout Success",
						},
						"slo_error": {
							Type: schema.TypeInt, Optional: true, Description: "Single Logout Error",
						},
						"sp_slo_req": {
							Type: schema.TypeInt, Optional: true, Description: "SP-initiated Single Logout Request",
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
						"other_error": {
							Type: schema.TypeInt, Optional: true, Description: "Other Error",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationSamlServiceProviderStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationSamlServiceProviderStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationSamlServiceProviderStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationSamlServiceProviderStatsStats := setObjectAamAuthenticationSamlServiceProviderStatsStats(res)
		d.Set("stats", AamAuthenticationSamlServiceProviderStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationSamlServiceProviderStatsStats(ret edpt.DataAamAuthenticationSamlServiceProviderStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"sp_metadata_export_req":     ret.DtAamAuthenticationSamlServiceProviderStats.Stats.SpMetadataExportReq,
			"sp_metadata_export_success": ret.DtAamAuthenticationSamlServiceProviderStats.Stats.SpMetadataExportSuccess,
			"login_auth_req":             ret.DtAamAuthenticationSamlServiceProviderStats.Stats.LoginAuthReq,
			"login_auth_resp":            ret.DtAamAuthenticationSamlServiceProviderStats.Stats.LoginAuthResp,
			"acs_req":                    ret.DtAamAuthenticationSamlServiceProviderStats.Stats.AcsReq,
			"acs_success":                ret.DtAamAuthenticationSamlServiceProviderStats.Stats.AcsSuccess,
			"acs_authz_fail":             ret.DtAamAuthenticationSamlServiceProviderStats.Stats.AcsAuthzFail,
			"acs_error":                  ret.DtAamAuthenticationSamlServiceProviderStats.Stats.AcsError,
			"slo_req":                    ret.DtAamAuthenticationSamlServiceProviderStats.Stats.SloReq,
			"slo_success":                ret.DtAamAuthenticationSamlServiceProviderStats.Stats.SloSuccess,
			"slo_error":                  ret.DtAamAuthenticationSamlServiceProviderStats.Stats.SloError,
			"sp_slo_req":                 ret.DtAamAuthenticationSamlServiceProviderStats.Stats.SpSloReq,
			"glo_slo_success":            ret.DtAamAuthenticationSamlServiceProviderStats.Stats.GloSloSuccess,
			"loc_slo_success":            ret.DtAamAuthenticationSamlServiceProviderStats.Stats.LocSloSuccess,
			"par_slo_success":            ret.DtAamAuthenticationSamlServiceProviderStats.Stats.ParSloSuccess,
			"other_error":                ret.DtAamAuthenticationSamlServiceProviderStats.Stats.OtherError,
		},
	}
}

func getObjectAamAuthenticationSamlServiceProviderStatsStats(d []interface{}) edpt.AamAuthenticationSamlServiceProviderStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationSamlServiceProviderStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
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
		ret.OtherError = in["other_error"].(int)
	}
	return ret
}

func dataToEndpointAamAuthenticationSamlServiceProviderStats(d *schema.ResourceData) edpt.AamAuthenticationSamlServiceProviderStats {
	var ret edpt.AamAuthenticationSamlServiceProviderStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectAamAuthenticationSamlServiceProviderStatsStats(d.Get("stats").([]interface{}))
	return ret
}
