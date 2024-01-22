package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonitoredEntityMonTopkSourcesOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_monitored_entity_mon_topk_sources_oper`: Operational Status for the object sources\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityMonitoredEntityMonTopkSourcesOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"l4_proto": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"l4_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
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

func resourceVisibilityMonitoredEntityMonTopkSourcesOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitoredEntityMonTopkSourcesOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitoredEntityMonTopkSourcesOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityMonitoredEntityMonTopkSourcesOperOper := setObjectVisibilityMonitoredEntityMonTopkSourcesOperOper(res)
		d.Set("oper", VisibilityMonitoredEntityMonTopkSourcesOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityMonitoredEntityMonTopkSourcesOperOper(ret edpt.DataVisibilityMonitoredEntityMonTopkSourcesOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ipv4_addr":        ret.DtVisibilityMonitoredEntityMonTopkSourcesOper.Oper.Ipv4Addr,
			"ipv6_addr":        ret.DtVisibilityMonitoredEntityMonTopkSourcesOper.Oper.Ipv6Addr,
			"l4_proto":         ret.DtVisibilityMonitoredEntityMonTopkSourcesOper.Oper.L4Proto,
			"l4_port":          ret.DtVisibilityMonitoredEntityMonTopkSourcesOper.Oper.L4Port,
			"metric_topk_list": setSliceVisibilityMonitoredEntityMonTopkSourcesOperOperMetricTopkList(ret.DtVisibilityMonitoredEntityMonTopkSourcesOper.Oper.MetricTopkList),
		},
	}
}

func setSliceVisibilityMonitoredEntityMonTopkSourcesOperOperMetricTopkList(d []edpt.VisibilityMonitoredEntityMonTopkSourcesOperOperMetricTopkList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["metric_name"] = item.MetricName
		in["topk_list"] = setSliceVisibilityMonitoredEntityMonTopkSourcesOperOperMetricTopkListTopkList(item.TopkList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntityMonTopkSourcesOperOperMetricTopkListTopkList(d []edpt.VisibilityMonitoredEntityMonTopkSourcesOperOperMetricTopkListTopkList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip_addr"] = item.IpAddr
		in["metric_value"] = item.MetricValue
		result = append(result, in)
	}
	return result
}

func getObjectVisibilityMonitoredEntityMonTopkSourcesOperOper(d []interface{}) edpt.VisibilityMonitoredEntityMonTopkSourcesOperOper {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityMonTopkSourcesOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4Addr = in["ipv4_addr"].(string)
		ret.Ipv6Addr = in["ipv6_addr"].(string)
		ret.L4Proto = in["l4_proto"].(string)
		ret.L4Port = in["l4_port"].(int)
		ret.MetricTopkList = getSliceVisibilityMonitoredEntityMonTopkSourcesOperOperMetricTopkList(in["metric_topk_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityMonitoredEntityMonTopkSourcesOperOperMetricTopkList(d []interface{}) []edpt.VisibilityMonitoredEntityMonTopkSourcesOperOperMetricTopkList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityMonTopkSourcesOperOperMetricTopkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityMonTopkSourcesOperOperMetricTopkList
		oi.MetricName = in["metric_name"].(string)
		oi.TopkList = getSliceVisibilityMonitoredEntityMonTopkSourcesOperOperMetricTopkListTopkList(in["topk_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntityMonTopkSourcesOperOperMetricTopkListTopkList(d []interface{}) []edpt.VisibilityMonitoredEntityMonTopkSourcesOperOperMetricTopkListTopkList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityMonTopkSourcesOperOperMetricTopkListTopkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityMonTopkSourcesOperOperMetricTopkListTopkList
		oi.IpAddr = in["ip_addr"].(string)
		oi.MetricValue = in["metric_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityMonitoredEntityMonTopkSourcesOper(d *schema.ResourceData) edpt.VisibilityMonitoredEntityMonTopkSourcesOper {
	var ret edpt.VisibilityMonitoredEntityMonTopkSourcesOper

	ret.Oper = getObjectVisibilityMonitoredEntityMonTopkSourcesOperOper(d.Get("oper").([]interface{}))
	return ret
}
