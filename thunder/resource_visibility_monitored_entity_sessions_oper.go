package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonitoredEntitySessionsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_monitored_entity_sessions_oper`: Operational Status for the object sessions\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityMonitoredEntitySessionsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mon_entity_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"entity_key": {
										Type: schema.TypeString, Optional: true, Description: "",
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
									"session_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"proto": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"fwd_src_ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"fwd_src_port": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"fwd_dst_ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"fwd_dst_port": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"rev_src_ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"rev_src_port": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"rev_dst_ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"rev_dst_port": {
													Type: schema.TypeInt, Optional: true, Description: "",
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
												"session_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"proto": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"fwd_src_ip": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"fwd_src_port": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"fwd_dst_ip": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"fwd_dst_port": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"rev_src_ip": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"rev_src_port": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"rev_dst_ip": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"rev_dst_port": {
																Type: schema.TypeInt, Optional: true, Description: "",
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

func resourceVisibilityMonitoredEntitySessionsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitoredEntitySessionsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitoredEntitySessionsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityMonitoredEntitySessionsOperOper := setObjectVisibilityMonitoredEntitySessionsOperOper(res)
		d.Set("oper", VisibilityMonitoredEntitySessionsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityMonitoredEntitySessionsOperOper(ret edpt.DataVisibilityMonitoredEntitySessionsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"mon_entity_list": setSliceVisibilityMonitoredEntitySessionsOperOperMonEntityList(ret.DtVisibilityMonitoredEntitySessionsOper.Oper.MonEntityList),
		},
	}
}

func setSliceVisibilityMonitoredEntitySessionsOperOperMonEntityList(d []edpt.VisibilityMonitoredEntitySessionsOperOperMonEntityList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["entity_key"] = item.EntityKey
		in["ipv4_addr"] = item.Ipv4Addr
		in["ipv6_addr"] = item.Ipv6Addr
		in["l4_proto"] = item.L4Proto
		in["l4_port"] = item.L4Port
		in["session_list"] = setSliceVisibilityMonitoredEntitySessionsOperOperMonEntityListSessionList(item.SessionList)
		in["sec_entity_list"] = setSliceVisibilityMonitoredEntitySessionsOperOperMonEntityListSecEntityList(item.SecEntityList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntitySessionsOperOperMonEntityListSessionList(d []edpt.VisibilityMonitoredEntitySessionsOperOperMonEntityListSessionList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["proto"] = item.Proto
		in["fwd_src_ip"] = item.FwdSrcIp
		in["fwd_src_port"] = item.FwdSrcPort
		in["fwd_dst_ip"] = item.FwdDstIp
		in["fwd_dst_port"] = item.FwdDstPort
		in["rev_src_ip"] = item.RevSrcIp
		in["rev_src_port"] = item.RevSrcPort
		in["rev_dst_ip"] = item.RevDstIp
		in["rev_dst_port"] = item.RevDstPort
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntitySessionsOperOperMonEntityListSecEntityList(d []edpt.VisibilityMonitoredEntitySessionsOperOperMonEntityListSecEntityList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["entity_key"] = item.EntityKey
		in["ipv4_addr"] = item.Ipv4Addr
		in["ipv6_addr"] = item.Ipv6Addr
		in["l4_proto"] = item.L4Proto
		in["l4_port"] = item.L4Port
		in["session_list"] = setSliceVisibilityMonitoredEntitySessionsOperOperMonEntityListSecEntityListSessionList(item.SessionList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntitySessionsOperOperMonEntityListSecEntityListSessionList(d []edpt.VisibilityMonitoredEntitySessionsOperOperMonEntityListSecEntityListSessionList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["proto"] = item.Proto
		in["fwd_src_ip"] = item.FwdSrcIp
		in["fwd_src_port"] = item.FwdSrcPort
		in["fwd_dst_ip"] = item.FwdDstIp
		in["fwd_dst_port"] = item.FwdDstPort
		in["rev_src_ip"] = item.RevSrcIp
		in["rev_src_port"] = item.RevSrcPort
		in["rev_dst_ip"] = item.RevDstIp
		in["rev_dst_port"] = item.RevDstPort
		result = append(result, in)
	}
	return result
}

func getObjectVisibilityMonitoredEntitySessionsOperOper(d []interface{}) edpt.VisibilityMonitoredEntitySessionsOperOper {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntitySessionsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MonEntityList = getSliceVisibilityMonitoredEntitySessionsOperOperMonEntityList(in["mon_entity_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityMonitoredEntitySessionsOperOperMonEntityList(d []interface{}) []edpt.VisibilityMonitoredEntitySessionsOperOperMonEntityList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntitySessionsOperOperMonEntityList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntitySessionsOperOperMonEntityList
		oi.EntityKey = in["entity_key"].(string)
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.L4Proto = in["l4_proto"].(string)
		oi.L4Port = in["l4_port"].(int)
		oi.SessionList = getSliceVisibilityMonitoredEntitySessionsOperOperMonEntityListSessionList(in["session_list"].([]interface{}))
		oi.SecEntityList = getSliceVisibilityMonitoredEntitySessionsOperOperMonEntityListSecEntityList(in["sec_entity_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntitySessionsOperOperMonEntityListSessionList(d []interface{}) []edpt.VisibilityMonitoredEntitySessionsOperOperMonEntityListSessionList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntitySessionsOperOperMonEntityListSessionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntitySessionsOperOperMonEntityListSessionList
		oi.Proto = in["proto"].(string)
		oi.FwdSrcIp = in["fwd_src_ip"].(string)
		oi.FwdSrcPort = in["fwd_src_port"].(int)
		oi.FwdDstIp = in["fwd_dst_ip"].(string)
		oi.FwdDstPort = in["fwd_dst_port"].(int)
		oi.RevSrcIp = in["rev_src_ip"].(string)
		oi.RevSrcPort = in["rev_src_port"].(int)
		oi.RevDstIp = in["rev_dst_ip"].(string)
		oi.RevDstPort = in["rev_dst_port"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntitySessionsOperOperMonEntityListSecEntityList(d []interface{}) []edpt.VisibilityMonitoredEntitySessionsOperOperMonEntityListSecEntityList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntitySessionsOperOperMonEntityListSecEntityList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntitySessionsOperOperMonEntityListSecEntityList
		oi.EntityKey = in["entity_key"].(string)
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.L4Proto = in["l4_proto"].(string)
		oi.L4Port = in["l4_port"].(int)
		oi.SessionList = getSliceVisibilityMonitoredEntitySessionsOperOperMonEntityListSecEntityListSessionList(in["session_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntitySessionsOperOperMonEntityListSecEntityListSessionList(d []interface{}) []edpt.VisibilityMonitoredEntitySessionsOperOperMonEntityListSecEntityListSessionList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntitySessionsOperOperMonEntityListSecEntityListSessionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntitySessionsOperOperMonEntityListSecEntityListSessionList
		oi.Proto = in["proto"].(string)
		oi.FwdSrcIp = in["fwd_src_ip"].(string)
		oi.FwdSrcPort = in["fwd_src_port"].(int)
		oi.FwdDstIp = in["fwd_dst_ip"].(string)
		oi.FwdDstPort = in["fwd_dst_port"].(int)
		oi.RevSrcIp = in["rev_src_ip"].(string)
		oi.RevSrcPort = in["rev_src_port"].(int)
		oi.RevDstIp = in["rev_dst_ip"].(string)
		oi.RevDstPort = in["rev_dst_port"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityMonitoredEntitySessionsOper(d *schema.ResourceData) edpt.VisibilityMonitoredEntitySessionsOper {
	var ret edpt.VisibilityMonitoredEntitySessionsOper

	ret.Oper = getObjectVisibilityMonitoredEntitySessionsOperOper(d.Get("oper").([]interface{}))
	return ret
}
