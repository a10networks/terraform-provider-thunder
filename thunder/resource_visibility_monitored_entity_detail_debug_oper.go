package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonitoredEntityDetailDebugOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_monitored_entity_detail_debug_oper`: Operational Status for the object debug\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityMonitoredEntityDetailDebugOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"primary_keys": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"all_keys": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"mon_entity_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"entity_key": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "",
									},
									"flat_oid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
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
									"mode": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ha_state": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"entity_metric_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"metric_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"current": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"min": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"max": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"mean": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"threshold": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"std_dev": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"anomaly": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"sec_entity_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"entity_key": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "",
												},
												"flat_oid": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
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
												"mode": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ha_state": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"entity_metric_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"metric_name": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"current": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"min": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"max": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"mean": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"threshold": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"std_dev": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"anomaly": {
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

func resourceVisibilityMonitoredEntityDetailDebugOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitoredEntityDetailDebugOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitoredEntityDetailDebugOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityMonitoredEntityDetailDebugOperOper := setObjectVisibilityMonitoredEntityDetailDebugOperOper(res)
		d.Set("oper", VisibilityMonitoredEntityDetailDebugOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityMonitoredEntityDetailDebugOperOper(ret edpt.DataVisibilityMonitoredEntityDetailDebugOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"primary_keys":    ret.DtVisibilityMonitoredEntityDetailDebugOper.Oper.PrimaryKeys,
			"all_keys":        ret.DtVisibilityMonitoredEntityDetailDebugOper.Oper.AllKeys,
			"mon_entity_list": setSliceVisibilityMonitoredEntityDetailDebugOperOperMonEntityList(ret.DtVisibilityMonitoredEntityDetailDebugOper.Oper.MonEntityList),
		},
	}
}

func setSliceVisibilityMonitoredEntityDetailDebugOperOperMonEntityList(d []edpt.VisibilityMonitoredEntityDetailDebugOperOperMonEntityList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["entity_key"] = item.EntityKey
		//omit uuid
		in["flat_oid"] = item.FlatOid
		in["ipv4_addr"] = item.Ipv4Addr
		in["ipv6_addr"] = item.Ipv6Addr
		in["l4_proto"] = item.L4Proto
		in["l4_port"] = item.L4Port
		in["mode"] = item.Mode
		in["ha_state"] = item.HaState
		in["entity_metric_list"] = setSliceVisibilityMonitoredEntityDetailDebugOperOperMonEntityListEntityMetricList(item.EntityMetricList)
		in["sec_entity_list"] = setSliceVisibilityMonitoredEntityDetailDebugOperOperMonEntityListSecEntityList(item.SecEntityList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntityDetailDebugOperOperMonEntityListEntityMetricList(d []edpt.VisibilityMonitoredEntityDetailDebugOperOperMonEntityListEntityMetricList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["metric_name"] = item.MetricName
		in["current"] = item.Current
		in["min"] = item.Min
		in["max"] = item.Max
		in["mean"] = item.Mean
		in["threshold"] = item.Threshold
		in["std_dev"] = item.StdDev
		in["anomaly"] = item.Anomaly
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntityDetailDebugOperOperMonEntityListSecEntityList(d []edpt.VisibilityMonitoredEntityDetailDebugOperOperMonEntityListSecEntityList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["entity_key"] = item.EntityKey
		//omit uuid
		in["flat_oid"] = item.FlatOid
		in["ipv4_addr"] = item.Ipv4Addr
		in["ipv6_addr"] = item.Ipv6Addr
		in["l4_proto"] = item.L4Proto
		in["l4_port"] = item.L4Port
		in["mode"] = item.Mode
		in["ha_state"] = item.HaState
		in["entity_metric_list"] = setSliceVisibilityMonitoredEntityDetailDebugOperOperMonEntityListSecEntityListEntityMetricList(item.EntityMetricList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntityDetailDebugOperOperMonEntityListSecEntityListEntityMetricList(d []edpt.VisibilityMonitoredEntityDetailDebugOperOperMonEntityListSecEntityListEntityMetricList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["metric_name"] = item.MetricName
		in["current"] = item.Current
		in["min"] = item.Min
		in["max"] = item.Max
		in["mean"] = item.Mean
		in["threshold"] = item.Threshold
		in["std_dev"] = item.StdDev
		in["anomaly"] = item.Anomaly
		result = append(result, in)
	}
	return result
}

func getObjectVisibilityMonitoredEntityDetailDebugOperOper(d []interface{}) edpt.VisibilityMonitoredEntityDetailDebugOperOper {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityDetailDebugOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PrimaryKeys = in["primary_keys"].(int)
		ret.AllKeys = in["all_keys"].(int)
		ret.MonEntityList = getSliceVisibilityMonitoredEntityDetailDebugOperOperMonEntityList(in["mon_entity_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityMonitoredEntityDetailDebugOperOperMonEntityList(d []interface{}) []edpt.VisibilityMonitoredEntityDetailDebugOperOperMonEntityList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityDetailDebugOperOperMonEntityList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityDetailDebugOperOperMonEntityList
		oi.EntityKey = in["entity_key"].(string)
		//omit uuid
		oi.FlatOid = in["flat_oid"].(int)
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.L4Proto = in["l4_proto"].(string)
		oi.L4Port = in["l4_port"].(int)
		oi.Mode = in["mode"].(string)
		oi.HaState = in["ha_state"].(string)
		oi.EntityMetricList = getSliceVisibilityMonitoredEntityDetailDebugOperOperMonEntityListEntityMetricList(in["entity_metric_list"].([]interface{}))
		oi.SecEntityList = getSliceVisibilityMonitoredEntityDetailDebugOperOperMonEntityListSecEntityList(in["sec_entity_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntityDetailDebugOperOperMonEntityListEntityMetricList(d []interface{}) []edpt.VisibilityMonitoredEntityDetailDebugOperOperMonEntityListEntityMetricList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityDetailDebugOperOperMonEntityListEntityMetricList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityDetailDebugOperOperMonEntityListEntityMetricList
		oi.MetricName = in["metric_name"].(string)
		oi.Current = in["current"].(string)
		oi.Min = in["min"].(string)
		oi.Max = in["max"].(string)
		oi.Mean = in["mean"].(string)
		oi.Threshold = in["threshold"].(string)
		oi.StdDev = in["std_dev"].(string)
		oi.Anomaly = in["anomaly"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntityDetailDebugOperOperMonEntityListSecEntityList(d []interface{}) []edpt.VisibilityMonitoredEntityDetailDebugOperOperMonEntityListSecEntityList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityDetailDebugOperOperMonEntityListSecEntityList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityDetailDebugOperOperMonEntityListSecEntityList
		oi.EntityKey = in["entity_key"].(string)
		//omit uuid
		oi.FlatOid = in["flat_oid"].(int)
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.L4Proto = in["l4_proto"].(string)
		oi.L4Port = in["l4_port"].(int)
		oi.Mode = in["mode"].(string)
		oi.HaState = in["ha_state"].(string)
		oi.EntityMetricList = getSliceVisibilityMonitoredEntityDetailDebugOperOperMonEntityListSecEntityListEntityMetricList(in["entity_metric_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntityDetailDebugOperOperMonEntityListSecEntityListEntityMetricList(d []interface{}) []edpt.VisibilityMonitoredEntityDetailDebugOperOperMonEntityListSecEntityListEntityMetricList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityDetailDebugOperOperMonEntityListSecEntityListEntityMetricList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityDetailDebugOperOperMonEntityListSecEntityListEntityMetricList
		oi.MetricName = in["metric_name"].(string)
		oi.Current = in["current"].(string)
		oi.Min = in["min"].(string)
		oi.Max = in["max"].(string)
		oi.Mean = in["mean"].(string)
		oi.Threshold = in["threshold"].(string)
		oi.StdDev = in["std_dev"].(string)
		oi.Anomaly = in["anomaly"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityMonitoredEntityDetailDebugOper(d *schema.ResourceData) edpt.VisibilityMonitoredEntityDetailDebugOper {
	var ret edpt.VisibilityMonitoredEntityDetailDebugOper

	ret.Oper = getObjectVisibilityMonitoredEntityDetailDebugOperOper(d.Get("oper").([]interface{}))
	return ret
}
