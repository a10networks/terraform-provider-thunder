package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6FibSummaryOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ipv6_fib_summary_oper`: Operational Status for the object fib-summary\n\n__PLACEHOLDER__",
		ReadContext: resourceIpv6FibSummaryOperRead,

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

func resourceIpv6FibSummaryOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6FibSummaryOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6FibSummaryOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Ipv6FibSummaryOperOper := setObjectIpv6FibSummaryOperOper(res)
		d.Set("oper", Ipv6FibSummaryOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectIpv6FibSummaryOperOper(ret edpt.DataIpv6FibSummaryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"connected_routes":      ret.DtIpv6FibSummaryOper.Oper.Connected_routes,
			"static_dynamic_routes": ret.DtIpv6FibSummaryOper.Oper.Static_dynamic_routes,
			"total_routes":          ret.DtIpv6FibSummaryOper.Oper.Total_routes,
			"static_dynamic_paths":  ret.DtIpv6FibSummaryOper.Oper.Static_dynamic_paths,
			"total_paths":           ret.DtIpv6FibSummaryOper.Oper.Total_paths,
		},
	}
}

func getObjectIpv6FibSummaryOperOper(d []interface{}) edpt.Ipv6FibSummaryOperOper {

	count1 := len(d)
	var ret edpt.Ipv6FibSummaryOperOper
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

func dataToEndpointIpv6FibSummaryOper(d *schema.ResourceData) edpt.Ipv6FibSummaryOper {
	var ret edpt.Ipv6FibSummaryOper

	ret.Oper = getObjectIpv6FibSummaryOperOper(d.Get("oper").([]interface{}))
	return ret
}
