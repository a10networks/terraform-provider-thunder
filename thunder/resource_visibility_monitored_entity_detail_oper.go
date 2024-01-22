package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonitoredEntityDetailOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_monitored_entity_detail_oper`: Operational Status for the object detail\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityMonitoredEntityDetailOperRead,

		Schema: map[string]*schema.Schema{
			"debug": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
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
				},
			},
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
												"threshold": {
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
															"threshold": {
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

func resourceVisibilityMonitoredEntityDetailOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitoredEntityDetailOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitoredEntityDetailOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityMonitoredEntityDetailOperDebug := setObjectVisibilityMonitoredEntityDetailOperDebug(res)
		d.Set("debug", VisibilityMonitoredEntityDetailOperDebug)
		VisibilityMonitoredEntityDetailOperOper := setObjectVisibilityMonitoredEntityDetailOperOper(res)
		d.Set("oper", VisibilityMonitoredEntityDetailOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityMonitoredEntityDetailOperDebug(ret edpt.DataVisibilityMonitoredEntityDetailOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectVisibilityMonitoredEntityDetailOperDebugOper(ret.DtVisibilityMonitoredEntityDetailOper.Debug.Oper),
		},
	}
}

func setObjectVisibilityMonitoredEntityDetailOperDebugOper(d edpt.VisibilityMonitoredEntityDetailOperDebugOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["primary_keys"] = d.PrimaryKeys

	in["all_keys"] = d.AllKeys
	in["mon_entity_list"] = setSliceVisibilityMonitoredEntityDetailOperDebugOperMonEntityList(d.MonEntityList)
	result = append(result, in)
	return result
}

func setSliceVisibilityMonitoredEntityDetailOperDebugOperMonEntityList(d []edpt.VisibilityMonitoredEntityDetailOperDebugOperMonEntityList) []map[string]interface{} {
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
		in["entity_metric_list"] = setSliceVisibilityMonitoredEntityDetailOperDebugOperMonEntityListEntityMetricList(item.EntityMetricList)
		in["sec_entity_list"] = setSliceVisibilityMonitoredEntityDetailOperDebugOperMonEntityListSecEntityList(item.SecEntityList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntityDetailOperDebugOperMonEntityListEntityMetricList(d []edpt.VisibilityMonitoredEntityDetailOperDebugOperMonEntityListEntityMetricList) []map[string]interface{} {
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

func setSliceVisibilityMonitoredEntityDetailOperDebugOperMonEntityListSecEntityList(d []edpt.VisibilityMonitoredEntityDetailOperDebugOperMonEntityListSecEntityList) []map[string]interface{} {
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
		in["entity_metric_list"] = setSliceVisibilityMonitoredEntityDetailOperDebugOperMonEntityListSecEntityListEntityMetricList(item.EntityMetricList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntityDetailOperDebugOperMonEntityListSecEntityListEntityMetricList(d []edpt.VisibilityMonitoredEntityDetailOperDebugOperMonEntityListSecEntityListEntityMetricList) []map[string]interface{} {
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

func setObjectVisibilityMonitoredEntityDetailOperOper(ret edpt.DataVisibilityMonitoredEntityDetailOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"primary_keys":    ret.DtVisibilityMonitoredEntityDetailOper.Oper.PrimaryKeys,
			"all_keys":        ret.DtVisibilityMonitoredEntityDetailOper.Oper.AllKeys,
			"mon_entity_list": setSliceVisibilityMonitoredEntityDetailOperOperMonEntityList(ret.DtVisibilityMonitoredEntityDetailOper.Oper.MonEntityList),
		},
	}
}

func setSliceVisibilityMonitoredEntityDetailOperOperMonEntityList(d []edpt.VisibilityMonitoredEntityDetailOperOperMonEntityList) []map[string]interface{} {
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
		in["entity_metric_list"] = setSliceVisibilityMonitoredEntityDetailOperOperMonEntityListEntityMetricList(item.EntityMetricList)
		in["sec_entity_list"] = setSliceVisibilityMonitoredEntityDetailOperOperMonEntityListSecEntityList(item.SecEntityList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntityDetailOperOperMonEntityListEntityMetricList(d []edpt.VisibilityMonitoredEntityDetailOperOperMonEntityListEntityMetricList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["metric_name"] = item.MetricName
		in["current"] = item.Current
		in["threshold"] = item.Threshold
		in["anomaly"] = item.Anomaly
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntityDetailOperOperMonEntityListSecEntityList(d []edpt.VisibilityMonitoredEntityDetailOperOperMonEntityListSecEntityList) []map[string]interface{} {
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
		in["entity_metric_list"] = setSliceVisibilityMonitoredEntityDetailOperOperMonEntityListSecEntityListEntityMetricList(item.EntityMetricList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntityDetailOperOperMonEntityListSecEntityListEntityMetricList(d []edpt.VisibilityMonitoredEntityDetailOperOperMonEntityListSecEntityListEntityMetricList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["metric_name"] = item.MetricName
		in["current"] = item.Current
		in["threshold"] = item.Threshold
		in["anomaly"] = item.Anomaly
		result = append(result, in)
	}
	return result
}

func getObjectVisibilityMonitoredEntityDetailOperDebug(d []interface{}) edpt.VisibilityMonitoredEntityDetailOperDebug {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityDetailOperDebug
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVisibilityMonitoredEntityDetailOperDebugOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityMonitoredEntityDetailOperDebugOper(d []interface{}) edpt.VisibilityMonitoredEntityDetailOperDebugOper {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityDetailOperDebugOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PrimaryKeys = in["primary_keys"].(int)
		ret.AllKeys = in["all_keys"].(int)
		ret.MonEntityList = getSliceVisibilityMonitoredEntityDetailOperDebugOperMonEntityList(in["mon_entity_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityMonitoredEntityDetailOperDebugOperMonEntityList(d []interface{}) []edpt.VisibilityMonitoredEntityDetailOperDebugOperMonEntityList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityDetailOperDebugOperMonEntityList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityDetailOperDebugOperMonEntityList
		oi.EntityKey = in["entity_key"].(string)
		//omit uuid
		oi.FlatOid = in["flat_oid"].(int)
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.L4Proto = in["l4_proto"].(string)
		oi.L4Port = in["l4_port"].(int)
		oi.Mode = in["mode"].(string)
		oi.HaState = in["ha_state"].(string)
		oi.EntityMetricList = getSliceVisibilityMonitoredEntityDetailOperDebugOperMonEntityListEntityMetricList(in["entity_metric_list"].([]interface{}))
		oi.SecEntityList = getSliceVisibilityMonitoredEntityDetailOperDebugOperMonEntityListSecEntityList(in["sec_entity_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntityDetailOperDebugOperMonEntityListEntityMetricList(d []interface{}) []edpt.VisibilityMonitoredEntityDetailOperDebugOperMonEntityListEntityMetricList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityDetailOperDebugOperMonEntityListEntityMetricList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityDetailOperDebugOperMonEntityListEntityMetricList
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

func getSliceVisibilityMonitoredEntityDetailOperDebugOperMonEntityListSecEntityList(d []interface{}) []edpt.VisibilityMonitoredEntityDetailOperDebugOperMonEntityListSecEntityList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityDetailOperDebugOperMonEntityListSecEntityList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityDetailOperDebugOperMonEntityListSecEntityList
		oi.EntityKey = in["entity_key"].(string)
		//omit uuid
		oi.FlatOid = in["flat_oid"].(int)
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.L4Proto = in["l4_proto"].(string)
		oi.L4Port = in["l4_port"].(int)
		oi.Mode = in["mode"].(string)
		oi.HaState = in["ha_state"].(string)
		oi.EntityMetricList = getSliceVisibilityMonitoredEntityDetailOperDebugOperMonEntityListSecEntityListEntityMetricList(in["entity_metric_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntityDetailOperDebugOperMonEntityListSecEntityListEntityMetricList(d []interface{}) []edpt.VisibilityMonitoredEntityDetailOperDebugOperMonEntityListSecEntityListEntityMetricList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityDetailOperDebugOperMonEntityListSecEntityListEntityMetricList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityDetailOperDebugOperMonEntityListSecEntityListEntityMetricList
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

func getObjectVisibilityMonitoredEntityDetailOperOper(d []interface{}) edpt.VisibilityMonitoredEntityDetailOperOper {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityDetailOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PrimaryKeys = in["primary_keys"].(int)
		ret.AllKeys = in["all_keys"].(int)
		ret.MonEntityList = getSliceVisibilityMonitoredEntityDetailOperOperMonEntityList(in["mon_entity_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityMonitoredEntityDetailOperOperMonEntityList(d []interface{}) []edpt.VisibilityMonitoredEntityDetailOperOperMonEntityList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityDetailOperOperMonEntityList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityDetailOperOperMonEntityList
		oi.EntityKey = in["entity_key"].(string)
		//omit uuid
		oi.FlatOid = in["flat_oid"].(int)
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.L4Proto = in["l4_proto"].(string)
		oi.L4Port = in["l4_port"].(int)
		oi.Mode = in["mode"].(string)
		oi.HaState = in["ha_state"].(string)
		oi.EntityMetricList = getSliceVisibilityMonitoredEntityDetailOperOperMonEntityListEntityMetricList(in["entity_metric_list"].([]interface{}))
		oi.SecEntityList = getSliceVisibilityMonitoredEntityDetailOperOperMonEntityListSecEntityList(in["sec_entity_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntityDetailOperOperMonEntityListEntityMetricList(d []interface{}) []edpt.VisibilityMonitoredEntityDetailOperOperMonEntityListEntityMetricList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityDetailOperOperMonEntityListEntityMetricList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityDetailOperOperMonEntityListEntityMetricList
		oi.MetricName = in["metric_name"].(string)
		oi.Current = in["current"].(string)
		oi.Threshold = in["threshold"].(string)
		oi.Anomaly = in["anomaly"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntityDetailOperOperMonEntityListSecEntityList(d []interface{}) []edpt.VisibilityMonitoredEntityDetailOperOperMonEntityListSecEntityList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityDetailOperOperMonEntityListSecEntityList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityDetailOperOperMonEntityListSecEntityList
		oi.EntityKey = in["entity_key"].(string)
		//omit uuid
		oi.FlatOid = in["flat_oid"].(int)
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.L4Proto = in["l4_proto"].(string)
		oi.L4Port = in["l4_port"].(int)
		oi.Mode = in["mode"].(string)
		oi.HaState = in["ha_state"].(string)
		oi.EntityMetricList = getSliceVisibilityMonitoredEntityDetailOperOperMonEntityListSecEntityListEntityMetricList(in["entity_metric_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntityDetailOperOperMonEntityListSecEntityListEntityMetricList(d []interface{}) []edpt.VisibilityMonitoredEntityDetailOperOperMonEntityListSecEntityListEntityMetricList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityDetailOperOperMonEntityListSecEntityListEntityMetricList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityDetailOperOperMonEntityListSecEntityListEntityMetricList
		oi.MetricName = in["metric_name"].(string)
		oi.Current = in["current"].(string)
		oi.Threshold = in["threshold"].(string)
		oi.Anomaly = in["anomaly"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityMonitoredEntityDetailOper(d *schema.ResourceData) edpt.VisibilityMonitoredEntityDetailOper {
	var ret edpt.VisibilityMonitoredEntityDetailOper

	ret.Debug = getObjectVisibilityMonitoredEntityDetailOperDebug(d.Get("debug").([]interface{}))

	ret.Oper = getObjectVisibilityMonitoredEntityDetailOperOper(d.Get("oper").([]interface{}))
	return ret
}
