package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6FibOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ipv6_fib_oper`: Operational Status for the object fib\n\n__PLACEHOLDER__",
		ReadContext: resourceIpv6FibOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total_routes": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_paths": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipv6_fib": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"prefixlen": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nexthop": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"interface": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"distance": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceIpv6FibOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6FibOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6FibOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Ipv6FibOperOper := setObjectIpv6FibOperOper(res)
		d.Set("oper", Ipv6FibOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectIpv6FibOperOper(ret edpt.DataIpv6FibOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total_routes": ret.DtIpv6FibOper.Oper.Total_routes,
			"total_paths":  ret.DtIpv6FibOper.Oper.Total_paths,
			"ipv6_fib":     setSliceIpv6FibOperOperIpv6Fib(ret.DtIpv6FibOper.Oper.Ipv6Fib),
		},
	}
}

func setSliceIpv6FibOperOperIpv6Fib(d []edpt.Ipv6FibOperOperIpv6Fib) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["prefix"] = item.Prefix
		in["prefixlen"] = item.Prefixlen
		in["nexthop"] = item.Nexthop
		in["interface"] = item.Interface
		in["distance"] = item.Distance
		result = append(result, in)
	}
	return result
}

func getObjectIpv6FibOperOper(d []interface{}) edpt.Ipv6FibOperOper {

	count1 := len(d)
	var ret edpt.Ipv6FibOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Total_routes = in["total_routes"].(int)
		ret.Total_paths = in["total_paths"].(int)
		ret.Ipv6Fib = getSliceIpv6FibOperOperIpv6Fib(in["ipv6_fib"].([]interface{}))
	}
	return ret
}

func getSliceIpv6FibOperOperIpv6Fib(d []interface{}) []edpt.Ipv6FibOperOperIpv6Fib {

	count1 := len(d)
	ret := make([]edpt.Ipv6FibOperOperIpv6Fib, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Ipv6FibOperOperIpv6Fib
		oi.Prefix = in["prefix"].(string)
		oi.Prefixlen = in["prefixlen"].(int)
		oi.Nexthop = in["nexthop"].(string)
		oi.Interface = in["interface"].(string)
		oi.Distance = in["distance"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpv6FibOper(d *schema.ResourceData) edpt.Ipv6FibOper {
	var ret edpt.Ipv6FibOper

	ret.Oper = getObjectIpv6FibOperOper(d.Get("oper").([]interface{}))
	return ret
}
