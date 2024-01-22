package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationServerWindowsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_server_windows_oper`: Operational Status for the object windows\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationServerWindowsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stats_clear_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationServerWindowsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerWindowsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerWindowsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationServerWindowsOperOper := setObjectAamAuthenticationServerWindowsOperOper(res)
		d.Set("oper", AamAuthenticationServerWindowsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationServerWindowsOperOper(ret edpt.DataAamAuthenticationServerWindowsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"stats_clear_type": ret.DtAamAuthenticationServerWindowsOper.Oper.StatsClearType,
			"name":             ret.DtAamAuthenticationServerWindowsOper.Oper.Name,
		},
	}
}

func getObjectAamAuthenticationServerWindowsOperOper(d []interface{}) edpt.AamAuthenticationServerWindowsOperOper {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerWindowsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatsClearType = in["stats_clear_type"].(string)
		ret.Name = in["name"].(string)
	}
	return ret
}

func dataToEndpointAamAuthenticationServerWindowsOper(d *schema.ResourceData) edpt.AamAuthenticationServerWindowsOper {
	var ret edpt.AamAuthenticationServerWindowsOper

	ret.Oper = getObjectAamAuthenticationServerWindowsOperOper(d.Get("oper").([]interface{}))
	return ret
}
