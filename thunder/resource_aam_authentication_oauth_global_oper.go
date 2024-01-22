package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationOauthGlobalOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_oauth_global_oper`: Operational Status for the object global\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationOauthGlobalOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stats_clear_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationOauthGlobalOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationOauthGlobalOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationOauthGlobalOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationOauthGlobalOperOper := setObjectAamAuthenticationOauthGlobalOperOper(res)
		d.Set("oper", AamAuthenticationOauthGlobalOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationOauthGlobalOperOper(ret edpt.DataAamAuthenticationOauthGlobalOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"stats_clear_type": ret.DtAamAuthenticationOauthGlobalOper.Oper.StatsClearType,
		},
	}
}

func getObjectAamAuthenticationOauthGlobalOperOper(d []interface{}) edpt.AamAuthenticationOauthGlobalOperOper {

	count1 := len(d)
	var ret edpt.AamAuthenticationOauthGlobalOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatsClearType = in["stats_clear_type"].(string)
	}
	return ret
}

func dataToEndpointAamAuthenticationOauthGlobalOper(d *schema.ResourceData) edpt.AamAuthenticationOauthGlobalOper {
	var ret edpt.AamAuthenticationOauthGlobalOper

	ret.Oper = getObjectAamAuthenticationOauthGlobalOperOper(d.Get("oper").([]interface{}))
	return ret
}
