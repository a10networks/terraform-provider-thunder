package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonitoredEntitySecondaryMonTopkOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_monitored_entity_secondary_mon_topk_oper`: Operational Status for the object mon-topk\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityMonitoredEntitySecondaryMonTopkOperRead,

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

func resourceVisibilityMonitoredEntitySecondaryMonTopkOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitoredEntitySecondaryMonTopkOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitoredEntitySecondaryMonTopkOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityMonitoredEntitySecondaryMonTopkOperOper := setObjectVisibilityMonitoredEntitySecondaryMonTopkOperOper(res)
		d.Set("oper", VisibilityMonitoredEntitySecondaryMonTopkOperOper)
		VisibilityMonitoredEntitySecondaryMonTopkOperSources := setObjectVisibilityMonitoredEntitySecondaryMonTopkOperSources(res)
		d.Set("sources", VisibilityMonitoredEntitySecondaryMonTopkOperSources)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityMonitoredEntitySecondaryMonTopkOperOper(ret edpt.DataVisibilityMonitoredEntitySecondaryMonTopkOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ipv4_addr":        ret.DtVisibilityMonitoredEntitySecondaryMonTopkOper.Oper.Ipv4Addr,
			"ipv6_addr":        ret.DtVisibilityMonitoredEntitySecondaryMonTopkOper.Oper.Ipv6Addr,
			"l4_proto":         ret.DtVisibilityMonitoredEntitySecondaryMonTopkOper.Oper.L4Proto,
			"l4_port":          ret.DtVisibilityMonitoredEntitySecondaryMonTopkOper.Oper.L4Port,
			"metric_topk_list": setSliceVisibilityMonitoredEntitySecondaryMonTopkOperOperMetricTopkList(ret.DtVisibilityMonitoredEntitySecondaryMonTopkOper.Oper.MetricTopkList),
		},
	}
}

func setSliceVisibilityMonitoredEntitySecondaryMonTopkOperOperMetricTopkList(d []edpt.VisibilityMonitoredEntitySecondaryMonTopkOperOperMetricTopkList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["metric_name"] = item.MetricName
		in["topk_list"] = setSliceVisibilityMonitoredEntitySecondaryMonTopkOperOperMetricTopkListTopkList(item.TopkList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntitySecondaryMonTopkOperOperMetricTopkListTopkList(d []edpt.VisibilityMonitoredEntitySecondaryMonTopkOperOperMetricTopkListTopkList) []map[string]interface{} {
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

func setObjectVisibilityMonitoredEntitySecondaryMonTopkOperSources(ret edpt.DataVisibilityMonitoredEntitySecondaryMonTopkOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectVisibilityMonitoredEntitySecondaryMonTopkOperSourcesOper(ret.DtVisibilityMonitoredEntitySecondaryMonTopkOper.Sources.Oper),
		},
	}
}

func setObjectVisibilityMonitoredEntitySecondaryMonTopkOperSourcesOper(d edpt.VisibilityMonitoredEntitySecondaryMonTopkOperSourcesOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["ipv4_addr"] = d.Ipv4Addr

	in["ipv6_addr"] = d.Ipv6Addr

	in["l4_proto"] = d.L4Proto

	in["l4_port"] = d.L4Port
	in["metric_topk_list"] = setSliceVisibilityMonitoredEntitySecondaryMonTopkOperSourcesOperMetricTopkList(d.MetricTopkList)
	result = append(result, in)
	return result
}

func setSliceVisibilityMonitoredEntitySecondaryMonTopkOperSourcesOperMetricTopkList(d []edpt.VisibilityMonitoredEntitySecondaryMonTopkOperSourcesOperMetricTopkList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["metric_name"] = item.MetricName
		in["topk_list"] = setSliceVisibilityMonitoredEntitySecondaryMonTopkOperSourcesOperMetricTopkListTopkList(item.TopkList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntitySecondaryMonTopkOperSourcesOperMetricTopkListTopkList(d []edpt.VisibilityMonitoredEntitySecondaryMonTopkOperSourcesOperMetricTopkListTopkList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip_addr"] = item.IpAddr
		in["metric_value"] = item.MetricValue
		result = append(result, in)
	}
	return result
}

func getObjectVisibilityMonitoredEntitySecondaryMonTopkOperOper(d []interface{}) edpt.VisibilityMonitoredEntitySecondaryMonTopkOperOper {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntitySecondaryMonTopkOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4Addr = in["ipv4_addr"].(string)
		ret.Ipv6Addr = in["ipv6_addr"].(string)
		ret.L4Proto = in["l4_proto"].(string)
		ret.L4Port = in["l4_port"].(int)
		ret.MetricTopkList = getSliceVisibilityMonitoredEntitySecondaryMonTopkOperOperMetricTopkList(in["metric_topk_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityMonitoredEntitySecondaryMonTopkOperOperMetricTopkList(d []interface{}) []edpt.VisibilityMonitoredEntitySecondaryMonTopkOperOperMetricTopkList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntitySecondaryMonTopkOperOperMetricTopkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntitySecondaryMonTopkOperOperMetricTopkList
		oi.MetricName = in["metric_name"].(string)
		oi.TopkList = getSliceVisibilityMonitoredEntitySecondaryMonTopkOperOperMetricTopkListTopkList(in["topk_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntitySecondaryMonTopkOperOperMetricTopkListTopkList(d []interface{}) []edpt.VisibilityMonitoredEntitySecondaryMonTopkOperOperMetricTopkListTopkList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntitySecondaryMonTopkOperOperMetricTopkListTopkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntitySecondaryMonTopkOperOperMetricTopkListTopkList
		oi.IpAddr = in["ip_addr"].(string)
		oi.Port = in["port"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.MetricValue = in["metric_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityMonitoredEntitySecondaryMonTopkOperSources(d []interface{}) edpt.VisibilityMonitoredEntitySecondaryMonTopkOperSources {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntitySecondaryMonTopkOperSources
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVisibilityMonitoredEntitySecondaryMonTopkOperSourcesOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityMonitoredEntitySecondaryMonTopkOperSourcesOper(d []interface{}) edpt.VisibilityMonitoredEntitySecondaryMonTopkOperSourcesOper {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntitySecondaryMonTopkOperSourcesOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4Addr = in["ipv4_addr"].(string)
		ret.Ipv6Addr = in["ipv6_addr"].(string)
		ret.L4Proto = in["l4_proto"].(string)
		ret.L4Port = in["l4_port"].(int)
		ret.MetricTopkList = getSliceVisibilityMonitoredEntitySecondaryMonTopkOperSourcesOperMetricTopkList(in["metric_topk_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityMonitoredEntitySecondaryMonTopkOperSourcesOperMetricTopkList(d []interface{}) []edpt.VisibilityMonitoredEntitySecondaryMonTopkOperSourcesOperMetricTopkList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntitySecondaryMonTopkOperSourcesOperMetricTopkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntitySecondaryMonTopkOperSourcesOperMetricTopkList
		oi.MetricName = in["metric_name"].(string)
		oi.TopkList = getSliceVisibilityMonitoredEntitySecondaryMonTopkOperSourcesOperMetricTopkListTopkList(in["topk_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntitySecondaryMonTopkOperSourcesOperMetricTopkListTopkList(d []interface{}) []edpt.VisibilityMonitoredEntitySecondaryMonTopkOperSourcesOperMetricTopkListTopkList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntitySecondaryMonTopkOperSourcesOperMetricTopkListTopkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntitySecondaryMonTopkOperSourcesOperMetricTopkListTopkList
		oi.IpAddr = in["ip_addr"].(string)
		oi.MetricValue = in["metric_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityMonitoredEntitySecondaryMonTopkOper(d *schema.ResourceData) edpt.VisibilityMonitoredEntitySecondaryMonTopkOper {
	var ret edpt.VisibilityMonitoredEntitySecondaryMonTopkOper

	ret.Oper = getObjectVisibilityMonitoredEntitySecondaryMonTopkOperOper(d.Get("oper").([]interface{}))

	ret.Sources = getObjectVisibilityMonitoredEntitySecondaryMonTopkOperSources(d.Get("sources").([]interface{}))
	return ret
}
