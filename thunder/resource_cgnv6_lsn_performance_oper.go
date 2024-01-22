package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnPerformanceOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_performance_oper`: Operational Status for the object performance\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnPerformanceOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"full_cone_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"user_quotas": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6LsnPerformanceOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnPerformanceOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnPerformanceOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnPerformanceOperOper := setObjectCgnv6LsnPerformanceOperOper(res)
		d.Set("oper", Cgnv6LsnPerformanceOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnPerformanceOperOper(ret edpt.DataCgnv6LsnPerformanceOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"data_sessions":      ret.DtCgnv6LsnPerformanceOper.Oper.DataSessions,
			"full_cone_sessions": ret.DtCgnv6LsnPerformanceOper.Oper.FullConeSessions,
			"user_quotas":        ret.DtCgnv6LsnPerformanceOper.Oper.UserQuotas,
		},
	}
}

func getObjectCgnv6LsnPerformanceOperOper(d []interface{}) edpt.Cgnv6LsnPerformanceOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6LsnPerformanceOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DataSessions = in["data_sessions"].(int)
		ret.FullConeSessions = in["full_cone_sessions"].(int)
		ret.UserQuotas = in["user_quotas"].(int)
	}
	return ret
}

func dataToEndpointCgnv6LsnPerformanceOper(d *schema.ResourceData) edpt.Cgnv6LsnPerformanceOper {
	var ret edpt.Cgnv6LsnPerformanceOper

	ret.Oper = getObjectCgnv6LsnPerformanceOperOper(d.Get("oper").([]interface{}))
	return ret
}
