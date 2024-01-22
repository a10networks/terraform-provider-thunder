package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpRibOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ip_rib_oper`: Operational Status for the object rib\n\n__PLACEHOLDER__",
		ReadContext: resourceIpRibOperRead,

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
						"ipv4_routes": {
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

func resourceIpRibOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpRibOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpRibOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		IpRibOperOper := setObjectIpRibOperOper(res)
		d.Set("oper", IpRibOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectIpRibOperOper(ret edpt.DataIpRibOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"description": ret.DtIpRibOper.Oper.Description,
			"total":       ret.DtIpRibOper.Oper.Total,
			"total_paths": ret.DtIpRibOper.Oper.TotalPaths,
			"limit":       ret.DtIpRibOper.Oper.Limit,
			"ipv4_routes": setSliceIpRibOperOperIpv4Routes(ret.DtIpRibOper.Oper.Ipv4Routes),
		},
	}
}

func setSliceIpRibOperOperIpv4Routes(d []edpt.IpRibOperOperIpv4Routes) []map[string]interface{} {
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

func getObjectIpRibOperOper(d []interface{}) edpt.IpRibOperOper {

	count1 := len(d)
	var ret edpt.IpRibOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Description = in["description"].(string)
		ret.Total = in["total"].(int)
		ret.TotalPaths = in["total_paths"].(int)
		ret.Limit = in["limit"].(int)
		ret.Ipv4Routes = getSliceIpRibOperOperIpv4Routes(in["ipv4_routes"].([]interface{}))
	}
	return ret
}

func getSliceIpRibOperOperIpv4Routes(d []interface{}) []edpt.IpRibOperOperIpv4Routes {

	count1 := len(d)
	ret := make([]edpt.IpRibOperOperIpv4Routes, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpRibOperOperIpv4Routes
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

func dataToEndpointIpRibOper(d *schema.ResourceData) edpt.IpRibOper {
	var ret edpt.IpRibOper

	ret.Oper = getObjectIpRibOperOper(d.Get("oper").([]interface{}))
	return ret
}
