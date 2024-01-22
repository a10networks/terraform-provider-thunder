package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpFibOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ip_fib_oper`: Operational Status for the object fib\n\n__PLACEHOLDER__",
		ReadContext: resourceIpFibOperRead,

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
						"ipv4_fib": {
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

func resourceIpFibOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpFibOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpFibOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		IpFibOperOper := setObjectIpFibOperOper(res)
		d.Set("oper", IpFibOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectIpFibOperOper(ret edpt.DataIpFibOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total_routes": ret.DtIpFibOper.Oper.Total_routes,
			"total_paths":  ret.DtIpFibOper.Oper.Total_paths,
			"ipv4_fib":     setSliceIpFibOperOperIpv4Fib(ret.DtIpFibOper.Oper.Ipv4Fib),
		},
	}
}

func setSliceIpFibOperOperIpv4Fib(d []edpt.IpFibOperOperIpv4Fib) []map[string]interface{} {
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

func getObjectIpFibOperOper(d []interface{}) edpt.IpFibOperOper {

	count1 := len(d)
	var ret edpt.IpFibOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Total_routes = in["total_routes"].(int)
		ret.Total_paths = in["total_paths"].(int)
		ret.Ipv4Fib = getSliceIpFibOperOperIpv4Fib(in["ipv4_fib"].([]interface{}))
	}
	return ret
}

func getSliceIpFibOperOperIpv4Fib(d []interface{}) []edpt.IpFibOperOperIpv4Fib {

	count1 := len(d)
	ret := make([]edpt.IpFibOperOperIpv4Fib, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpFibOperOperIpv4Fib
		oi.Prefix = in["prefix"].(string)
		oi.Prefixlen = in["prefixlen"].(int)
		oi.Nexthop = in["nexthop"].(string)
		oi.Interface = in["interface"].(string)
		oi.Distance = in["distance"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpFibOper(d *schema.ResourceData) edpt.IpFibOper {
	var ret edpt.IpFibOper

	ret.Oper = getObjectIpFibOperOper(d.Get("oper").([]interface{}))
	return ret
}
