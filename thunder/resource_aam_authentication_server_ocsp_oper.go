package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationServerOcspOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_server_ocsp_oper`: Operational Status for the object ocsp\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationServerOcspOperRead,

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

func resourceAamAuthenticationServerOcspOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerOcspOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerOcspOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationServerOcspOperOper := setObjectAamAuthenticationServerOcspOperOper(res)
		d.Set("oper", AamAuthenticationServerOcspOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationServerOcspOperOper(ret edpt.DataAamAuthenticationServerOcspOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"stats_clear_type": ret.DtAamAuthenticationServerOcspOper.Oper.StatsClearType,
			"name":             ret.DtAamAuthenticationServerOcspOper.Oper.Name,
		},
	}
}

func getObjectAamAuthenticationServerOcspOperOper(d []interface{}) edpt.AamAuthenticationServerOcspOperOper {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerOcspOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatsClearType = in["stats_clear_type"].(string)
		ret.Name = in["name"].(string)
	}
	return ret
}

func dataToEndpointAamAuthenticationServerOcspOper(d *schema.ResourceData) edpt.AamAuthenticationServerOcspOper {
	var ret edpt.AamAuthenticationServerOcspOper

	ret.Oper = getObjectAamAuthenticationServerOcspOperOper(d.Get("oper").([]interface{}))
	return ret
}
