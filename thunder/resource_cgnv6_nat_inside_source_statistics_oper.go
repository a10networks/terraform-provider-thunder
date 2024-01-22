package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6NatInsideSourceStatisticsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_nat_inside_source_statistics_oper`: Operational Status for the object statistics\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6NatInsideSourceStatisticsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"static_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"src_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"port_usage": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_used": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_freed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nat_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"total": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6NatInsideSourceStatisticsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatInsideSourceStatisticsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatInsideSourceStatisticsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6NatInsideSourceStatisticsOperOper := setObjectCgnv6NatInsideSourceStatisticsOperOper(res)
		d.Set("oper", Cgnv6NatInsideSourceStatisticsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6NatInsideSourceStatisticsOperOper(ret edpt.DataCgnv6NatInsideSourceStatisticsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"static_list": setSliceCgnv6NatInsideSourceStatisticsOperOperStaticList(ret.DtCgnv6NatInsideSourceStatisticsOper.Oper.StaticList),
			"total":       ret.DtCgnv6NatInsideSourceStatisticsOper.Oper.Total,
		},
	}
}

func setSliceCgnv6NatInsideSourceStatisticsOperOperStaticList(d []edpt.Cgnv6NatInsideSourceStatisticsOperOperStaticList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["src_address"] = item.SrcAddress
		in["port_usage"] = item.PortUsage
		in["total_used"] = item.TotalUsed
		in["total_freed"] = item.TotalFreed
		in["nat_address"] = item.NatAddress
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6NatInsideSourceStatisticsOperOper(d []interface{}) edpt.Cgnv6NatInsideSourceStatisticsOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6NatInsideSourceStatisticsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticList = getSliceCgnv6NatInsideSourceStatisticsOperOperStaticList(in["static_list"].([]interface{}))
		ret.Total = in["total"].(int)
	}
	return ret
}

func getSliceCgnv6NatInsideSourceStatisticsOperOperStaticList(d []interface{}) []edpt.Cgnv6NatInsideSourceStatisticsOperOperStaticList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6NatInsideSourceStatisticsOperOperStaticList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6NatInsideSourceStatisticsOperOperStaticList
		oi.SrcAddress = in["src_address"].(string)
		oi.PortUsage = in["port_usage"].(int)
		oi.TotalUsed = in["total_used"].(int)
		oi.TotalFreed = in["total_freed"].(int)
		oi.NatAddress = in["nat_address"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6NatInsideSourceStatisticsOper(d *schema.ResourceData) edpt.Cgnv6NatInsideSourceStatisticsOper {
	var ret edpt.Cgnv6NatInsideSourceStatisticsOper

	ret.Oper = getObjectCgnv6NatInsideSourceStatisticsOperOper(d.Get("oper").([]interface{}))
	return ret
}
