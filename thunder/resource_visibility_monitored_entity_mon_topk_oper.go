package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonitoredEntityMonTopkOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_monitored_entity_mon_topk_oper`: Operational Status for the object mon-topk\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityMonitoredEntityMonTopkOperRead,

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
												"port": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"protocol": {
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
			"sources": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
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
				},
			},
		},
	}
}

func resourceVisibilityMonitoredEntityMonTopkOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitoredEntityMonTopkOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitoredEntityMonTopkOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityMonitoredEntityMonTopkOperOper := setObjectVisibilityMonitoredEntityMonTopkOperOper(res)
		d.Set("oper", VisibilityMonitoredEntityMonTopkOperOper)
		VisibilityMonitoredEntityMonTopkOperSources := setObjectVisibilityMonitoredEntityMonTopkOperSources(res)
		d.Set("sources", VisibilityMonitoredEntityMonTopkOperSources)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityMonitoredEntityMonTopkOperOper(ret edpt.DataVisibilityMonitoredEntityMonTopkOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"metric_topk_list": setSliceVisibilityMonitoredEntityMonTopkOperOperMetricTopkList(ret.DtVisibilityMonitoredEntityMonTopkOper.Oper.MetricTopkList),
		},
	}
}

func setSliceVisibilityMonitoredEntityMonTopkOperOperMetricTopkList(d []edpt.VisibilityMonitoredEntityMonTopkOperOperMetricTopkList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["metric_name"] = item.MetricName
		in["topk_list"] = setSliceVisibilityMonitoredEntityMonTopkOperOperMetricTopkListTopkList(item.TopkList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntityMonTopkOperOperMetricTopkListTopkList(d []edpt.VisibilityMonitoredEntityMonTopkOperOperMetricTopkListTopkList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip_addr"] = item.IpAddr
		in["port"] = item.Port
		in["protocol"] = item.Protocol
		in["metric_value"] = item.MetricValue
		result = append(result, in)
	}
	return result
}

func setObjectVisibilityMonitoredEntityMonTopkOperSources(ret edpt.DataVisibilityMonitoredEntityMonTopkOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectVisibilityMonitoredEntityMonTopkOperSourcesOper(ret.DtVisibilityMonitoredEntityMonTopkOper.Sources.Oper),
		},
	}
}

func setObjectVisibilityMonitoredEntityMonTopkOperSourcesOper(d edpt.VisibilityMonitoredEntityMonTopkOperSourcesOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["ipv4_addr"] = d.Ipv4Addr

	in["ipv6_addr"] = d.Ipv6Addr

	in["l4_proto"] = d.L4Proto

	in["l4_port"] = d.L4Port
	in["metric_topk_list"] = setSliceVisibilityMonitoredEntityMonTopkOperSourcesOperMetricTopkList(d.MetricTopkList)
	result = append(result, in)
	return result
}

func setSliceVisibilityMonitoredEntityMonTopkOperSourcesOperMetricTopkList(d []edpt.VisibilityMonitoredEntityMonTopkOperSourcesOperMetricTopkList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["metric_name"] = item.MetricName
		in["topk_list"] = setSliceVisibilityMonitoredEntityMonTopkOperSourcesOperMetricTopkListTopkList(item.TopkList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntityMonTopkOperSourcesOperMetricTopkListTopkList(d []edpt.VisibilityMonitoredEntityMonTopkOperSourcesOperMetricTopkListTopkList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip_addr"] = item.IpAddr
		in["metric_value"] = item.MetricValue
		result = append(result, in)
	}
	return result
}

func getObjectVisibilityMonitoredEntityMonTopkOperOper(d []interface{}) edpt.VisibilityMonitoredEntityMonTopkOperOper {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityMonTopkOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MetricTopkList = getSliceVisibilityMonitoredEntityMonTopkOperOperMetricTopkList(in["metric_topk_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityMonitoredEntityMonTopkOperOperMetricTopkList(d []interface{}) []edpt.VisibilityMonitoredEntityMonTopkOperOperMetricTopkList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityMonTopkOperOperMetricTopkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityMonTopkOperOperMetricTopkList
		oi.MetricName = in["metric_name"].(string)
		oi.TopkList = getSliceVisibilityMonitoredEntityMonTopkOperOperMetricTopkListTopkList(in["topk_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntityMonTopkOperOperMetricTopkListTopkList(d []interface{}) []edpt.VisibilityMonitoredEntityMonTopkOperOperMetricTopkListTopkList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityMonTopkOperOperMetricTopkListTopkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityMonTopkOperOperMetricTopkListTopkList
		oi.IpAddr = in["ip_addr"].(string)
		oi.Port = in["port"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.MetricValue = in["metric_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityMonitoredEntityMonTopkOperSources(d []interface{}) edpt.VisibilityMonitoredEntityMonTopkOperSources {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityMonTopkOperSources
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVisibilityMonitoredEntityMonTopkOperSourcesOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityMonitoredEntityMonTopkOperSourcesOper(d []interface{}) edpt.VisibilityMonitoredEntityMonTopkOperSourcesOper {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityMonTopkOperSourcesOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4Addr = in["ipv4_addr"].(string)
		ret.Ipv6Addr = in["ipv6_addr"].(string)
		ret.L4Proto = in["l4_proto"].(string)
		ret.L4Port = in["l4_port"].(int)
		ret.MetricTopkList = getSliceVisibilityMonitoredEntityMonTopkOperSourcesOperMetricTopkList(in["metric_topk_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityMonitoredEntityMonTopkOperSourcesOperMetricTopkList(d []interface{}) []edpt.VisibilityMonitoredEntityMonTopkOperSourcesOperMetricTopkList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityMonTopkOperSourcesOperMetricTopkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityMonTopkOperSourcesOperMetricTopkList
		oi.MetricName = in["metric_name"].(string)
		oi.TopkList = getSliceVisibilityMonitoredEntityMonTopkOperSourcesOperMetricTopkListTopkList(in["topk_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntityMonTopkOperSourcesOperMetricTopkListTopkList(d []interface{}) []edpt.VisibilityMonitoredEntityMonTopkOperSourcesOperMetricTopkListTopkList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityMonTopkOperSourcesOperMetricTopkListTopkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityMonTopkOperSourcesOperMetricTopkListTopkList
		oi.IpAddr = in["ip_addr"].(string)
		oi.MetricValue = in["metric_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityMonitoredEntityMonTopkOper(d *schema.ResourceData) edpt.VisibilityMonitoredEntityMonTopkOper {
	var ret edpt.VisibilityMonitoredEntityMonTopkOper

	ret.Oper = getObjectVisibilityMonitoredEntityMonTopkOperOper(d.Get("oper").([]interface{}))

	ret.Sources = getObjectVisibilityMonitoredEntityMonTopkOperSources(d.Get("sources").([]interface{}))
	return ret
}
