package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamResourceUsageOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_resource_usage_oper`: Operational Status for the object resource-usage\n\n__PLACEHOLDER__",
		ReadContext: resourceAamResourceUsageOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"idp_limit_current": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"idp_limit_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"idp_limit_minimum": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"idp_limit_maximum": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceAamResourceUsageOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamResourceUsageOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamResourceUsageOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamResourceUsageOperOper := setObjectAamResourceUsageOperOper(res)
		d.Set("oper", AamResourceUsageOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamResourceUsageOperOper(ret edpt.DataAamResourceUsageOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"idp_limit_current": ret.DtAamResourceUsageOper.Oper.IdpLimitCurrent,
			"idp_limit_default": ret.DtAamResourceUsageOper.Oper.IdpLimitDefault,
			"idp_limit_minimum": ret.DtAamResourceUsageOper.Oper.IdpLimitMinimum,
			"idp_limit_maximum": ret.DtAamResourceUsageOper.Oper.IdpLimitMaximum,
		},
	}
}

func getObjectAamResourceUsageOperOper(d []interface{}) edpt.AamResourceUsageOperOper {

	count1 := len(d)
	var ret edpt.AamResourceUsageOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IdpLimitCurrent = in["idp_limit_current"].(int)
		ret.IdpLimitDefault = in["idp_limit_default"].(int)
		ret.IdpLimitMinimum = in["idp_limit_minimum"].(int)
		ret.IdpLimitMaximum = in["idp_limit_maximum"].(int)
	}
	return ret
}

func dataToEndpointAamResourceUsageOper(d *schema.ResourceData) edpt.AamResourceUsageOper {
	var ret edpt.AamResourceUsageOper

	ret.Oper = getObjectAamResourceUsageOperOper(d.Get("oper").([]interface{}))
	return ret
}
