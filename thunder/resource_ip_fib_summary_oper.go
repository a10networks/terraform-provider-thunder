package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpFibSummaryOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ip_fib_summary_oper`: Operational Status for the object fib-summary\n\n__PLACEHOLDER__",
		ReadContext: resourceIpFibSummaryOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connected_routes": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"static_dynamic_routes": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_routes": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"static_dynamic_paths": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_paths": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceIpFibSummaryOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpFibSummaryOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpFibSummaryOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		IpFibSummaryOperOper := setObjectIpFibSummaryOperOper(res)
		d.Set("oper", IpFibSummaryOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectIpFibSummaryOperOper(ret edpt.DataIpFibSummaryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"connected_routes":      ret.DtIpFibSummaryOper.Oper.Connected_routes,
			"static_dynamic_routes": ret.DtIpFibSummaryOper.Oper.Static_dynamic_routes,
			"total_routes":          ret.DtIpFibSummaryOper.Oper.Total_routes,
			"static_dynamic_paths":  ret.DtIpFibSummaryOper.Oper.Static_dynamic_paths,
			"total_paths":           ret.DtIpFibSummaryOper.Oper.Total_paths,
		},
	}
}

func getObjectIpFibSummaryOperOper(d []interface{}) edpt.IpFibSummaryOperOper {

	count1 := len(d)
	var ret edpt.IpFibSummaryOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Connected_routes = in["connected_routes"].(int)
		ret.Static_dynamic_routes = in["static_dynamic_routes"].(int)
		ret.Total_routes = in["total_routes"].(int)
		ret.Static_dynamic_paths = in["static_dynamic_paths"].(int)
		ret.Total_paths = in["total_paths"].(int)
	}
	return ret
}

func dataToEndpointIpFibSummaryOper(d *schema.ResourceData) edpt.IpFibSummaryOper {
	var ret edpt.IpFibSummaryOper

	ret.Oper = getObjectIpFibSummaryOperOper(d.Get("oper").([]interface{}))
	return ret
}
