package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationSamlIdentityProviderStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_saml_identity_provider_stats`: Statistics for the object identity-provider\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationSamlIdentityProviderStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "SAML authentication identity provider name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"valid_status": {
							Type: schema.TypeInt, Optional: true, Description: "Valid IdP status or not",
						},
						"md_state": {
							Type: schema.TypeInt, Optional: true, Description: "Metadata State",
						},
						"md_update": {
							Type: schema.TypeInt, Optional: true, Description: "Metadata Update Success Count",
						},
						"md_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Metadata Update Fail Count",
						},
						"acs_state": {
							Type: schema.TypeInt, Optional: true, Description: "ACS State",
						},
						"acs_req": {
							Type: schema.TypeInt, Optional: true, Description: "ACS Request Total Count",
						},
						"acs_pass": {
							Type: schema.TypeInt, Optional: true, Description: "ACS Pass Count",
						},
						"acs_fail": {
							Type: schema.TypeInt, Optional: true, Description: "ACS Fail Count",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationSamlIdentityProviderStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationSamlIdentityProviderStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationSamlIdentityProviderStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationSamlIdentityProviderStatsStats := setObjectAamAuthenticationSamlIdentityProviderStatsStats(res)
		d.Set("stats", AamAuthenticationSamlIdentityProviderStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationSamlIdentityProviderStatsStats(ret edpt.DataAamAuthenticationSamlIdentityProviderStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"valid_status": ret.DtAamAuthenticationSamlIdentityProviderStats.Stats.ValidStatus,
			"md_state":     ret.DtAamAuthenticationSamlIdentityProviderStats.Stats.MdState,
			"md_update":    ret.DtAamAuthenticationSamlIdentityProviderStats.Stats.MdUpdate,
			"md_fail":      ret.DtAamAuthenticationSamlIdentityProviderStats.Stats.MdFail,
			"acs_state":    ret.DtAamAuthenticationSamlIdentityProviderStats.Stats.AcsState,
			"acs_req":      ret.DtAamAuthenticationSamlIdentityProviderStats.Stats.AcsReq,
			"acs_pass":     ret.DtAamAuthenticationSamlIdentityProviderStats.Stats.AcsPass,
			"acs_fail":     ret.DtAamAuthenticationSamlIdentityProviderStats.Stats.AcsFail,
		},
	}
}

func getObjectAamAuthenticationSamlIdentityProviderStatsStats(d []interface{}) edpt.AamAuthenticationSamlIdentityProviderStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationSamlIdentityProviderStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ValidStatus = in["valid_status"].(int)
		ret.MdState = in["md_state"].(int)
		ret.MdUpdate = in["md_update"].(int)
		ret.MdFail = in["md_fail"].(int)
		ret.AcsState = in["acs_state"].(int)
		ret.AcsReq = in["acs_req"].(int)
		ret.AcsPass = in["acs_pass"].(int)
		ret.AcsFail = in["acs_fail"].(int)
	}
	return ret
}

func dataToEndpointAamAuthenticationSamlIdentityProviderStats(d *schema.ResourceData) edpt.AamAuthenticationSamlIdentityProviderStats {
	var ret edpt.AamAuthenticationSamlIdentityProviderStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectAamAuthenticationSamlIdentityProviderStatsStats(d.Get("stats").([]interface{}))
	return ret
}
