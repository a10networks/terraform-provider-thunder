package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonitoredEntitySecondaryMonTopkSourcesOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_monitored_entity_secondary_mon_topk_sources_oper`: Operational Status for the object sources\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityMonitoredEntitySecondaryMonTopkSourcesOperRead,

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

func resourceVisibilityMonitoredEntitySecondaryMonTopkSourcesOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitoredEntitySecondaryMonTopkSourcesOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitoredEntitySecondaryMonTopkSourcesOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityMonitoredEntitySecondaryMonTopkSourcesOperOper := setObjectVisibilityMonitoredEntitySecondaryMonTopkSourcesOperOper(res)
		d.Set("oper", VisibilityMonitoredEntitySecondaryMonTopkSourcesOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityMonitoredEntitySecondaryMonTopkSourcesOperOper(ret edpt.DataVisibilityMonitoredEntitySecondaryMonTopkSourcesOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ipv4_addr":        ret.DtVisibilityMonitoredEntitySecondaryMonTopkSourcesOper.Oper.Ipv4Addr,
			"ipv6_addr":        ret.DtVisibilityMonitoredEntitySecondaryMonTopkSourcesOper.Oper.Ipv6Addr,
			"l4_proto":         ret.DtVisibilityMonitoredEntitySecondaryMonTopkSourcesOper.Oper.L4Proto,
			"l4_port":          ret.DtVisibilityMonitoredEntitySecondaryMonTopkSourcesOper.Oper.L4Port,
			"metric_topk_list": setSliceVisibilityMonitoredEntitySecondaryMonTopkSourcesOperOperMetricTopkList(ret.DtVisibilityMonitoredEntitySecondaryMonTopkSourcesOper.Oper.MetricTopkList),
		},
	}
}

func setSliceVisibilityMonitoredEntitySecondaryMonTopkSourcesOperOperMetricTopkList(d []edpt.VisibilityMonitoredEntitySecondaryMonTopkSourcesOperOperMetricTopkList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["metric_name"] = item.MetricName
		in["topk_list"] = setSliceVisibilityMonitoredEntitySecondaryMonTopkSourcesOperOperMetricTopkListTopkList(item.TopkList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntitySecondaryMonTopkSourcesOperOperMetricTopkListTopkList(d []edpt.VisibilityMonitoredEntitySecondaryMonTopkSourcesOperOperMetricTopkListTopkList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip_addr"] = item.IpAddr
		in["metric_value"] = item.MetricValue
		result = append(result, in)
	}
	return result
}

func getObjectVisibilityMonitoredEntitySecondaryMonTopkSourcesOperOper(d []interface{}) edpt.VisibilityMonitoredEntitySecondaryMonTopkSourcesOperOper {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntitySecondaryMonTopkSourcesOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4Addr = in["ipv4_addr"].(string)
		ret.Ipv6Addr = in["ipv6_addr"].(string)
		ret.L4Proto = in["l4_proto"].(string)
		ret.L4Port = in["l4_port"].(int)
		ret.MetricTopkList = getSliceVisibilityMonitoredEntitySecondaryMonTopkSourcesOperOperMetricTopkList(in["metric_topk_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityMonitoredEntitySecondaryMonTopkSourcesOperOperMetricTopkList(d []interface{}) []edpt.VisibilityMonitoredEntitySecondaryMonTopkSourcesOperOperMetricTopkList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntitySecondaryMonTopkSourcesOperOperMetricTopkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntitySecondaryMonTopkSourcesOperOperMetricTopkList
		oi.MetricName = in["metric_name"].(string)
		oi.TopkList = getSliceVisibilityMonitoredEntitySecondaryMonTopkSourcesOperOperMetricTopkListTopkList(in["topk_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntitySecondaryMonTopkSourcesOperOperMetricTopkListTopkList(d []interface{}) []edpt.VisibilityMonitoredEntitySecondaryMonTopkSourcesOperOperMetricTopkListTopkList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntitySecondaryMonTopkSourcesOperOperMetricTopkListTopkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntitySecondaryMonTopkSourcesOperOperMetricTopkListTopkList
		oi.IpAddr = in["ip_addr"].(string)
		oi.MetricValue = in["metric_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityMonitoredEntitySecondaryMonTopkSourcesOper(d *schema.ResourceData) edpt.VisibilityMonitoredEntitySecondaryMonTopkSourcesOper {
	var ret edpt.VisibilityMonitoredEntitySecondaryMonTopkSourcesOper

	ret.Oper = getObjectVisibilityMonitoredEntitySecondaryMonTopkSourcesOperOper(d.Get("oper").([]interface{}))
	return ret
}
