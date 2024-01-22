package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6RibOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ipv6_rib_oper`: Operational Status for the object rib\n\n__PLACEHOLDER__",
		ReadContext: resourceIpv6RibOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"total": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_paths": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"limit": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipv6_routes": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"prefixlen": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"subtype": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"distance": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"metric": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nexthop": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"interface": {
										Type: schema.TypeString, Optional: true, Description: "",
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

func resourceIpv6RibOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RibOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RibOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Ipv6RibOperOper := setObjectIpv6RibOperOper(res)
		d.Set("oper", Ipv6RibOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectIpv6RibOperOper(ret edpt.DataIpv6RibOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"description": ret.DtIpv6RibOper.Oper.Description,
			"total":       ret.DtIpv6RibOper.Oper.Total,
			"total_paths": ret.DtIpv6RibOper.Oper.TotalPaths,
			"limit":       ret.DtIpv6RibOper.Oper.Limit,
			"ipv6_routes": setSliceIpv6RibOperOperIpv6Routes(ret.DtIpv6RibOper.Oper.Ipv6Routes),
		},
	}
}

func setSliceIpv6RibOperOperIpv6Routes(d []edpt.Ipv6RibOperOperIpv6Routes) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["prefix"] = item.Prefix
		in["prefixlen"] = item.Prefixlen
		in["type"] = item.Type
		in["subtype"] = item.Subtype
		in["distance"] = item.Distance
		in["metric"] = item.Metric
		in["nexthop"] = item.Nexthop
		in["interface"] = item.Interface
		result = append(result, in)
	}
	return result
}

func getObjectIpv6RibOperOper(d []interface{}) edpt.Ipv6RibOperOper {

	count1 := len(d)
	var ret edpt.Ipv6RibOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Description = in["description"].(string)
		ret.Total = in["total"].(int)
		ret.TotalPaths = in["total_paths"].(int)
		ret.Limit = in["limit"].(int)
		ret.Ipv6Routes = getSliceIpv6RibOperOperIpv6Routes(in["ipv6_routes"].([]interface{}))
	}
	return ret
}

func getSliceIpv6RibOperOperIpv6Routes(d []interface{}) []edpt.Ipv6RibOperOperIpv6Routes {

	count1 := len(d)
	ret := make([]edpt.Ipv6RibOperOperIpv6Routes, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Ipv6RibOperOperIpv6Routes
		oi.Prefix = in["prefix"].(string)
		oi.Prefixlen = in["prefixlen"].(int)
		oi.Type = in["type"].(string)
		oi.Subtype = in["subtype"].(string)
		oi.Distance = in["distance"].(int)
		oi.Metric = in["metric"].(int)
		oi.Nexthop = in["nexthop"].(string)
		oi.Interface = in["interface"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpv6RibOper(d *schema.ResourceData) edpt.Ipv6RibOper {
	var ret edpt.Ipv6RibOper

	ret.Oper = getObjectIpv6RibOperOper(d.Get("oper").([]interface{}))
	return ret
}
