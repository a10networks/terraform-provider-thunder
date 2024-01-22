package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonTopkSourcesOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_mon_topk_sources_oper`: Operational Status for the object sources\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityMonTopkSourcesOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metric_topk_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"metric_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"topk_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip_addr": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"metric_value": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
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

func resourceVisibilityMonTopkSourcesOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonTopkSourcesOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonTopkSourcesOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityMonTopkSourcesOperOper := setObjectVisibilityMonTopkSourcesOperOper(res)
		d.Set("oper", VisibilityMonTopkSourcesOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityMonTopkSourcesOperOper(ret edpt.DataVisibilityMonTopkSourcesOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"metric_topk_list": setSliceVisibilityMonTopkSourcesOperOperMetricTopkList(ret.DtVisibilityMonTopkSourcesOper.Oper.MetricTopkList),
		},
	}
}

func setSliceVisibilityMonTopkSourcesOperOperMetricTopkList(d []edpt.VisibilityMonTopkSourcesOperOperMetricTopkList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["metric_name"] = item.MetricName
		in["topk_list"] = setSliceVisibilityMonTopkSourcesOperOperMetricTopkListTopkList(item.TopkList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonTopkSourcesOperOperMetricTopkListTopkList(d []edpt.VisibilityMonTopkSourcesOperOperMetricTopkListTopkList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip_addr"] = item.IpAddr
		in["metric_value"] = item.MetricValue
		result = append(result, in)
	}
	return result
}

func getObjectVisibilityMonTopkSourcesOperOper(d []interface{}) edpt.VisibilityMonTopkSourcesOperOper {

	count1 := len(d)
	var ret edpt.VisibilityMonTopkSourcesOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MetricTopkList = getSliceVisibilityMonTopkSourcesOperOperMetricTopkList(in["metric_topk_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityMonTopkSourcesOperOperMetricTopkList(d []interface{}) []edpt.VisibilityMonTopkSourcesOperOperMetricTopkList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonTopkSourcesOperOperMetricTopkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonTopkSourcesOperOperMetricTopkList
		oi.MetricName = in["metric_name"].(string)
		oi.TopkList = getSliceVisibilityMonTopkSourcesOperOperMetricTopkListTopkList(in["topk_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonTopkSourcesOperOperMetricTopkListTopkList(d []interface{}) []edpt.VisibilityMonTopkSourcesOperOperMetricTopkListTopkList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonTopkSourcesOperOperMetricTopkListTopkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonTopkSourcesOperOperMetricTopkListTopkList
		oi.IpAddr = in["ip_addr"].(string)
		oi.MetricValue = in["metric_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityMonTopkSourcesOper(d *schema.ResourceData) edpt.VisibilityMonTopkSourcesOper {
	var ret edpt.VisibilityMonTopkSourcesOper

	ret.Oper = getObjectVisibilityMonTopkSourcesOperOper(d.Get("oper").([]interface{}))
	return ret
}
