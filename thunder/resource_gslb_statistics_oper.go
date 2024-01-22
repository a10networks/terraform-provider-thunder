package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbStatisticsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_statistics_oper`: Operational Status for the object statistics\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbStatisticsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"curr_ssl_ctx": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceGslbStatisticsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbStatisticsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbStatisticsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbStatisticsOperOper := setObjectGslbStatisticsOperOper(res)
		d.Set("oper", GslbStatisticsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbStatisticsOperOper(ret edpt.DataGslbStatisticsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"curr_ssl_ctx": ret.DtGslbStatisticsOper.Oper.CurrSslCtx,
		},
	}
}

func getObjectGslbStatisticsOperOper(d []interface{}) edpt.GslbStatisticsOperOper {

	count1 := len(d)
	var ret edpt.GslbStatisticsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CurrSslCtx = in["curr_ssl_ctx"].(int)
	}
	return ret
}

func dataToEndpointGslbStatisticsOper(d *schema.ResourceData) edpt.GslbStatisticsOper {
	var ret edpt.GslbStatisticsOper

	ret.Oper = getObjectGslbStatisticsOperOper(d.Get("oper").([]interface{}))
	return ret
}
