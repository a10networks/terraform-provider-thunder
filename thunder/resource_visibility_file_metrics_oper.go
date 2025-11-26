package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityFileMetricsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_file_metrics_oper`: Operational Status for the object metrics\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityFileMetricsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"monitor_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"pri_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"pri_ipv4_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"pri_ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"pri_l4_proto": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"pri_l4_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sec_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"sec_ipv4_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"sec_ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"sec_l4_proto": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sec_l4_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"file_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"proc_metric_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"metric_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"metric_attr_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"metric_attr_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"metric_attr_value": {
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

func resourceVisibilityFileMetricsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityFileMetricsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityFileMetricsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityFileMetricsOperOper := setObjectVisibilityFileMetricsOperOper(res)
		d.Set("oper", VisibilityFileMetricsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityFileMetricsOperOper(ret edpt.DataVisibilityFileMetricsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"monitor_type":     ret.DtVisibilityFileMetricsOper.Oper.MonitorType,
			"pri_type":         ret.DtVisibilityFileMetricsOper.Oper.PriType,
			"pri_ipv4_addr":    ret.DtVisibilityFileMetricsOper.Oper.PriIpv4Addr,
			"pri_ipv6_addr":    ret.DtVisibilityFileMetricsOper.Oper.PriIpv6Addr,
			"pri_l4_proto":     ret.DtVisibilityFileMetricsOper.Oper.PriL4Proto,
			"pri_l4_port":      ret.DtVisibilityFileMetricsOper.Oper.PriL4Port,
			"sec_type":         ret.DtVisibilityFileMetricsOper.Oper.SecType,
			"sec_ipv4_addr":    ret.DtVisibilityFileMetricsOper.Oper.SecIpv4Addr,
			"sec_ipv6_addr":    ret.DtVisibilityFileMetricsOper.Oper.SecIpv6Addr,
			"sec_l4_proto":     ret.DtVisibilityFileMetricsOper.Oper.SecL4Proto,
			"sec_l4_port":      ret.DtVisibilityFileMetricsOper.Oper.SecL4Port,
			"file_name":        ret.DtVisibilityFileMetricsOper.Oper.FileName,
			"proc_metric_list": setSliceVisibilityFileMetricsOperOperProcMetricList(ret.DtVisibilityFileMetricsOper.Oper.ProcMetricList),
		},
	}
}

func setSliceVisibilityFileMetricsOperOperProcMetricList(d []edpt.VisibilityFileMetricsOperOperProcMetricList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["metric_name"] = item.MetricName
		in["metric_attr_list"] = setSliceVisibilityFileMetricsOperOperProcMetricListMetricAttrList(item.MetricAttrList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityFileMetricsOperOperProcMetricListMetricAttrList(d []edpt.VisibilityFileMetricsOperOperProcMetricListMetricAttrList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["metric_attr_name"] = item.MetricAttrName
		in["metric_attr_value"] = item.MetricAttrValue
		result = append(result, in)
	}
	return result
}

func getObjectVisibilityFileMetricsOperOper(d []interface{}) edpt.VisibilityFileMetricsOperOper {

	count1 := len(d)
	var ret edpt.VisibilityFileMetricsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MonitorType = in["monitor_type"].(string)
		ret.PriType = in["pri_type"].(string)
		ret.PriIpv4Addr = in["pri_ipv4_addr"].(string)
		ret.PriIpv6Addr = in["pri_ipv6_addr"].(string)
		ret.PriL4Proto = in["pri_l4_proto"].(int)
		ret.PriL4Port = in["pri_l4_port"].(int)
		ret.SecType = in["sec_type"].(string)
		ret.SecIpv4Addr = in["sec_ipv4_addr"].(string)
		ret.SecIpv6Addr = in["sec_ipv6_addr"].(string)
		ret.SecL4Proto = in["sec_l4_proto"].(int)
		ret.SecL4Port = in["sec_l4_port"].(int)
		ret.FileName = in["file_name"].(string)
		ret.ProcMetricList = getSliceVisibilityFileMetricsOperOperProcMetricList(in["proc_metric_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityFileMetricsOperOperProcMetricList(d []interface{}) []edpt.VisibilityFileMetricsOperOperProcMetricList {

	count1 := len(d)
	ret := make([]edpt.VisibilityFileMetricsOperOperProcMetricList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityFileMetricsOperOperProcMetricList
		oi.MetricName = in["metric_name"].(string)
		oi.MetricAttrList = getSliceVisibilityFileMetricsOperOperProcMetricListMetricAttrList(in["metric_attr_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityFileMetricsOperOperProcMetricListMetricAttrList(d []interface{}) []edpt.VisibilityFileMetricsOperOperProcMetricListMetricAttrList {

	count1 := len(d)
	ret := make([]edpt.VisibilityFileMetricsOperOperProcMetricListMetricAttrList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityFileMetricsOperOperProcMetricListMetricAttrList
		oi.MetricAttrName = in["metric_attr_name"].(string)
		oi.MetricAttrValue = in["metric_attr_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityFileMetricsOper(d *schema.ResourceData) edpt.VisibilityFileMetricsOper {
	var ret edpt.VisibilityFileMetricsOper

	ret.Oper = getObjectVisibilityFileMetricsOperOper(d.Get("oper").([]interface{}))
	return ret
}
