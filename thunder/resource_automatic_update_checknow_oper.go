package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAutomaticUpdateChecknowOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_automatic_update_checknow_oper`: Operational Status for the object checknow\n\n__PLACEHOLDER__",
		ReadContext: resourceAutomaticUpdateChecknowOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"feature_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"result": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceAutomaticUpdateChecknowOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateChecknowOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdateChecknowOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AutomaticUpdateChecknowOperOper := setObjectAutomaticUpdateChecknowOperOper(res)
		d.Set("oper", AutomaticUpdateChecknowOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAutomaticUpdateChecknowOperOper(ret edpt.DataAutomaticUpdateChecknowOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"feature_name": ret.DtAutomaticUpdateChecknowOper.Oper.FeatureName,
			"result":       ret.DtAutomaticUpdateChecknowOper.Oper.Result,
		},
	}
}

func getObjectAutomaticUpdateChecknowOperOper(d []interface{}) edpt.AutomaticUpdateChecknowOperOper {

	count1 := len(d)
	var ret edpt.AutomaticUpdateChecknowOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FeatureName = in["feature_name"].(string)
		ret.Result = in["result"].(string)
	}
	return ret
}

func dataToEndpointAutomaticUpdateChecknowOper(d *schema.ResourceData) edpt.AutomaticUpdateChecknowOper {
	var ret edpt.AutomaticUpdateChecknowOper

	ret.Oper = getObjectAutomaticUpdateChecknowOperOper(d.Get("oper").([]interface{}))
	return ret
}
