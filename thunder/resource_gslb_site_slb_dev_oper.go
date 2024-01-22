package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbSiteSlbDevOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_site_slb_dev_oper`: Operational Status for the object slb-dev\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbSiteSlbDevOperRead,

		Schema: map[string]*schema.Schema{
			"device_name": {
				Type: schema.TypeString, Required: true, Description: "Specify SLB device name",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dev_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dev_ip": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dev_attr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dev_admin_preference": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dev_session_num": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dev_session_util": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dev_gw_state": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dev_ip_cnt": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dev_state": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"client_ldns_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"client_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"age": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"rdt_sample1": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rdt_sample2": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rdt_sample3": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rdt_sample4": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rdt_sample5": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rdt_sample6": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rdt_sample7": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rdt_sample8": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"vip_server": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{},
							},
						},
						"vip_server_v4_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4": {
										Type: schema.TypeString, Required: true, Description: "Specify IP address",
									},
									"oper": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dev_vip_addr": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"dev_vip_state": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"node_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"service_ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"port_count": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"virtual_server": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"disabled": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"gslb_protocol": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"local_protocol": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"manually_health_check": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"use_gslb_state": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"dynamic": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"hits": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"recent": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"dev_vip_port_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dev_vip_port_num": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"dev_vip_port_state": {
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
						"vip_server_v6_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6": {
										Type: schema.TypeString, Required: true, Description: "Specify IP address (IPv6 address)",
									},
									"oper": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dev_vip_addr": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"dev_vip_state": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"node_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"service_ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"port_count": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"virtual_server": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"disabled": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"gslb_protocol": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"local_protocol": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"manually_health_check": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"use_gslb_state": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"dynamic": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"hits": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"recent": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"dev_vip_port_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dev_vip_port_num": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"dev_vip_port_state": {
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
						"vip_server_name_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vip_name": {
										Type: schema.TypeString, Required: true, Description: "Specify a VIP name for the SLB device",
									},
									"oper": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dev_vip_addr": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"dev_vip_state": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"node_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"service_ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"port_count": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"virtual_server": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"disabled": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"gslb_protocol": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"local_protocol": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"manually_health_check": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"use_gslb_state": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"dynamic": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"hits": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"recent": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"dev_vip_port_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dev_vip_port_num": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"dev_vip_port_state": {
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
			"site_name": {
				Type: schema.TypeString, Required: true, Description: "SiteName",
			},
		},
	}
}

func resourceGslbSiteSlbDevOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteSlbDevOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteSlbDevOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbSiteSlbDevOperOper := setObjectGslbSiteSlbDevOperOper(res)
		d.Set("oper", GslbSiteSlbDevOperOper)
		GslbSiteSlbDevOperVipServer := setObjectGslbSiteSlbDevOperVipServer(res)
		d.Set("vip_server", GslbSiteSlbDevOperVipServer)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbSiteSlbDevOperOper(ret edpt.DataGslbSiteSlbDevOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"dev_name":             ret.DtGslbSiteSlbDevOper.Oper.Dev_name,
			"dev_ip":               ret.DtGslbSiteSlbDevOper.Oper.Dev_ip,
			"dev_attr":             ret.DtGslbSiteSlbDevOper.Oper.Dev_attr,
			"dev_admin_preference": ret.DtGslbSiteSlbDevOper.Oper.Dev_admin_preference,
			"dev_session_num":      ret.DtGslbSiteSlbDevOper.Oper.Dev_session_num,
			"dev_session_util":     ret.DtGslbSiteSlbDevOper.Oper.Dev_session_util,
			"dev_gw_state":         ret.DtGslbSiteSlbDevOper.Oper.Dev_gw_state,
			"dev_ip_cnt":           ret.DtGslbSiteSlbDevOper.Oper.Dev_ip_cnt,
			"dev_state":            ret.DtGslbSiteSlbDevOper.Oper.DevState,
			"client_ldns_list":     setSliceGslbSiteSlbDevOperOperClientLdnsList(ret.DtGslbSiteSlbDevOper.Oper.ClientLdnsList),
		},
	}
}

func setSliceGslbSiteSlbDevOperOperClientLdnsList(d []edpt.GslbSiteSlbDevOperOperClientLdnsList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["client_ip"] = item.ClientIp
		in["age"] = item.Age
		in["type"] = item.Type
		in["rdt_sample1"] = item.RdtSample1
		in["rdt_sample2"] = item.RdtSample2
		in["rdt_sample3"] = item.RdtSample3
		in["rdt_sample4"] = item.RdtSample4
		in["rdt_sample5"] = item.RdtSample5
		in["rdt_sample6"] = item.RdtSample6
		in["rdt_sample7"] = item.RdtSample7
		in["rdt_sample8"] = item.RdtSample8
		result = append(result, in)
	}
	return result
}

func setObjectGslbSiteSlbDevOperVipServer(ret edpt.DataGslbSiteSlbDevOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper":                 setObjectGslbSiteSlbDevOperVipServerOper(ret.DtGslbSiteSlbDevOper.VipServer.Oper),
			"vip_server_v4_list":   setSliceGslbSiteSlbDevOperVipServerVipServerV4List(ret.DtGslbSiteSlbDevOper.VipServer.VipServerV4List),
			"vip_server_v6_list":   setSliceGslbSiteSlbDevOperVipServerVipServerV6List(ret.DtGslbSiteSlbDevOper.VipServer.VipServerV6List),
			"vip_server_name_list": setSliceGslbSiteSlbDevOperVipServerVipServerNameList(ret.DtGslbSiteSlbDevOper.VipServer.VipServerNameList),
		},
	}
}

func setObjectGslbSiteSlbDevOperVipServerOper(d edpt.GslbSiteSlbDevOperVipServerOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	result = append(result, in)
	return result
}

func setSliceGslbSiteSlbDevOperVipServerVipServerV4List(d []edpt.GslbSiteSlbDevOperVipServerVipServerV4List) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ipv4"] = item.Ipv4
		in["oper"] = setObjectGslbSiteSlbDevOperVipServerVipServerV4ListOper(item.Oper)
		result = append(result, in)
	}
	return result
}

func setObjectGslbSiteSlbDevOperVipServerVipServerV4ListOper(d edpt.GslbSiteSlbDevOperVipServerVipServerV4ListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["dev_vip_addr"] = d.Dev_vip_addr

	in["dev_vip_state"] = d.Dev_vip_state

	in["node_name"] = d.NodeName

	in["service_ip"] = d.ServiceIp

	in["port_count"] = d.PortCount

	in["virtual_server"] = d.VirtualServer

	in["disabled"] = d.Disabled

	in["gslb_protocol"] = d.GslbProtocol

	in["local_protocol"] = d.LocalProtocol

	in["manually_health_check"] = d.ManuallyHealthCheck

	in["use_gslb_state"] = d.Use_gslb_state

	in["dynamic"] = d.Dynamic

	in["hits"] = d.Hits

	in["recent"] = d.Recent
	in["dev_vip_port_list"] = setSliceGslbSiteSlbDevOperVipServerVipServerV4ListOperDevVipPortList(d.DevVipPortList)
	result = append(result, in)
	return result
}

func setSliceGslbSiteSlbDevOperVipServerVipServerV4ListOperDevVipPortList(d []edpt.GslbSiteSlbDevOperVipServerVipServerV4ListOperDevVipPortList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dev_vip_port_num"] = item.DevVipPortNum
		in["dev_vip_port_state"] = item.DevVipPortState
		result = append(result, in)
	}
	return result
}

func setSliceGslbSiteSlbDevOperVipServerVipServerV6List(d []edpt.GslbSiteSlbDevOperVipServerVipServerV6List) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ipv6"] = item.Ipv6
		in["oper"] = setObjectGslbSiteSlbDevOperVipServerVipServerV6ListOper(item.Oper)
		result = append(result, in)
	}
	return result
}

func setObjectGslbSiteSlbDevOperVipServerVipServerV6ListOper(d edpt.GslbSiteSlbDevOperVipServerVipServerV6ListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["dev_vip_addr"] = d.Dev_vip_addr

	in["dev_vip_state"] = d.Dev_vip_state

	in["node_name"] = d.NodeName

	in["service_ip"] = d.ServiceIp

	in["port_count"] = d.PortCount

	in["virtual_server"] = d.VirtualServer

	in["disabled"] = d.Disabled

	in["gslb_protocol"] = d.GslbProtocol

	in["local_protocol"] = d.LocalProtocol

	in["manually_health_check"] = d.ManuallyHealthCheck

	in["use_gslb_state"] = d.Use_gslb_state

	in["dynamic"] = d.Dynamic

	in["hits"] = d.Hits

	in["recent"] = d.Recent
	in["dev_vip_port_list"] = setSliceGslbSiteSlbDevOperVipServerVipServerV6ListOperDevVipPortList(d.DevVipPortList)
	result = append(result, in)
	return result
}

func setSliceGslbSiteSlbDevOperVipServerVipServerV6ListOperDevVipPortList(d []edpt.GslbSiteSlbDevOperVipServerVipServerV6ListOperDevVipPortList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dev_vip_port_num"] = item.DevVipPortNum
		in["dev_vip_port_state"] = item.DevVipPortState
		result = append(result, in)
	}
	return result
}

func setSliceGslbSiteSlbDevOperVipServerVipServerNameList(d []edpt.GslbSiteSlbDevOperVipServerVipServerNameList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["vip_name"] = item.VipName
		in["oper"] = setObjectGslbSiteSlbDevOperVipServerVipServerNameListOper(item.Oper)
		result = append(result, in)
	}
	return result
}

func setObjectGslbSiteSlbDevOperVipServerVipServerNameListOper(d edpt.GslbSiteSlbDevOperVipServerVipServerNameListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["dev_vip_addr"] = d.Dev_vip_addr

	in["dev_vip_state"] = d.Dev_vip_state

	in["node_name"] = d.NodeName

	in["service_ip"] = d.ServiceIp

	in["port_count"] = d.PortCount

	in["virtual_server"] = d.VirtualServer

	in["disabled"] = d.Disabled

	in["gslb_protocol"] = d.GslbProtocol

	in["local_protocol"] = d.LocalProtocol

	in["manually_health_check"] = d.ManuallyHealthCheck

	in["use_gslb_state"] = d.Use_gslb_state

	in["dynamic"] = d.Dynamic

	in["hits"] = d.Hits

	in["recent"] = d.Recent
	in["dev_vip_port_list"] = setSliceGslbSiteSlbDevOperVipServerVipServerNameListOperDevVipPortList(d.DevVipPortList)
	result = append(result, in)
	return result
}

func setSliceGslbSiteSlbDevOperVipServerVipServerNameListOperDevVipPortList(d []edpt.GslbSiteSlbDevOperVipServerVipServerNameListOperDevVipPortList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dev_vip_port_num"] = item.DevVipPortNum
		in["dev_vip_port_state"] = item.DevVipPortState
		result = append(result, in)
	}
	return result
}

func getObjectGslbSiteSlbDevOperOper(d []interface{}) edpt.GslbSiteSlbDevOperOper {

	count1 := len(d)
	var ret edpt.GslbSiteSlbDevOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dev_name = in["dev_name"].(string)
		ret.Dev_ip = in["dev_ip"].(string)
		ret.Dev_attr = in["dev_attr"].(string)
		ret.Dev_admin_preference = in["dev_admin_preference"].(int)
		ret.Dev_session_num = in["dev_session_num"].(int)
		ret.Dev_session_util = in["dev_session_util"].(int)
		ret.Dev_gw_state = in["dev_gw_state"].(string)
		ret.Dev_ip_cnt = in["dev_ip_cnt"].(int)
		ret.DevState = in["dev_state"].(string)
		ret.ClientLdnsList = getSliceGslbSiteSlbDevOperOperClientLdnsList(in["client_ldns_list"].([]interface{}))
	}
	return ret
}

func getSliceGslbSiteSlbDevOperOperClientLdnsList(d []interface{}) []edpt.GslbSiteSlbDevOperOperClientLdnsList {

	count1 := len(d)
	ret := make([]edpt.GslbSiteSlbDevOperOperClientLdnsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteSlbDevOperOperClientLdnsList
		oi.ClientIp = in["client_ip"].(string)
		oi.Age = in["age"].(int)
		oi.Type = in["type"].(string)
		oi.RdtSample1 = in["rdt_sample1"].(int)
		oi.RdtSample2 = in["rdt_sample2"].(int)
		oi.RdtSample3 = in["rdt_sample3"].(int)
		oi.RdtSample4 = in["rdt_sample4"].(int)
		oi.RdtSample5 = in["rdt_sample5"].(int)
		oi.RdtSample6 = in["rdt_sample6"].(int)
		oi.RdtSample7 = in["rdt_sample7"].(int)
		oi.RdtSample8 = in["rdt_sample8"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbSiteSlbDevOperVipServer(d []interface{}) edpt.GslbSiteSlbDevOperVipServer {

	count1 := len(d)
	var ret edpt.GslbSiteSlbDevOperVipServer
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectGslbSiteSlbDevOperVipServerOper(in["oper"].([]interface{}))
		ret.VipServerV4List = getSliceGslbSiteSlbDevOperVipServerVipServerV4List(in["vip_server_v4_list"].([]interface{}))
		ret.VipServerV6List = getSliceGslbSiteSlbDevOperVipServerVipServerV6List(in["vip_server_v6_list"].([]interface{}))
		ret.VipServerNameList = getSliceGslbSiteSlbDevOperVipServerVipServerNameList(in["vip_server_name_list"].([]interface{}))
	}
	return ret
}

func getObjectGslbSiteSlbDevOperVipServerOper(d []interface{}) edpt.GslbSiteSlbDevOperVipServerOper {

	var ret edpt.GslbSiteSlbDevOperVipServerOper
	return ret
}

func getSliceGslbSiteSlbDevOperVipServerVipServerV4List(d []interface{}) []edpt.GslbSiteSlbDevOperVipServerVipServerV4List {

	count1 := len(d)
	ret := make([]edpt.GslbSiteSlbDevOperVipServerVipServerV4List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteSlbDevOperVipServerVipServerV4List
		oi.Ipv4 = in["ipv4"].(string)
		oi.Oper = getObjectGslbSiteSlbDevOperVipServerVipServerV4ListOper(in["oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbSiteSlbDevOperVipServerVipServerV4ListOper(d []interface{}) edpt.GslbSiteSlbDevOperVipServerVipServerV4ListOper {

	count1 := len(d)
	var ret edpt.GslbSiteSlbDevOperVipServerVipServerV4ListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dev_vip_addr = in["dev_vip_addr"].(string)
		ret.Dev_vip_state = in["dev_vip_state"].(string)
		ret.NodeName = in["node_name"].(string)
		ret.ServiceIp = in["service_ip"].(string)
		ret.PortCount = in["port_count"].(int)
		ret.VirtualServer = in["virtual_server"].(int)
		ret.Disabled = in["disabled"].(int)
		ret.GslbProtocol = in["gslb_protocol"].(int)
		ret.LocalProtocol = in["local_protocol"].(int)
		ret.ManuallyHealthCheck = in["manually_health_check"].(int)
		ret.Use_gslb_state = in["use_gslb_state"].(int)
		ret.Dynamic = in["dynamic"].(int)
		ret.Hits = in["hits"].(int)
		ret.Recent = in["recent"].(int)
		ret.DevVipPortList = getSliceGslbSiteSlbDevOperVipServerVipServerV4ListOperDevVipPortList(in["dev_vip_port_list"].([]interface{}))
	}
	return ret
}

func getSliceGslbSiteSlbDevOperVipServerVipServerV4ListOperDevVipPortList(d []interface{}) []edpt.GslbSiteSlbDevOperVipServerVipServerV4ListOperDevVipPortList {

	count1 := len(d)
	ret := make([]edpt.GslbSiteSlbDevOperVipServerVipServerV4ListOperDevVipPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteSlbDevOperVipServerVipServerV4ListOperDevVipPortList
		oi.DevVipPortNum = in["dev_vip_port_num"].(int)
		oi.DevVipPortState = in["dev_vip_port_state"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbSiteSlbDevOperVipServerVipServerV6List(d []interface{}) []edpt.GslbSiteSlbDevOperVipServerVipServerV6List {

	count1 := len(d)
	ret := make([]edpt.GslbSiteSlbDevOperVipServerVipServerV6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteSlbDevOperVipServerVipServerV6List
		oi.Ipv6 = in["ipv6"].(string)
		oi.Oper = getObjectGslbSiteSlbDevOperVipServerVipServerV6ListOper(in["oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbSiteSlbDevOperVipServerVipServerV6ListOper(d []interface{}) edpt.GslbSiteSlbDevOperVipServerVipServerV6ListOper {

	count1 := len(d)
	var ret edpt.GslbSiteSlbDevOperVipServerVipServerV6ListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dev_vip_addr = in["dev_vip_addr"].(string)
		ret.Dev_vip_state = in["dev_vip_state"].(string)
		ret.NodeName = in["node_name"].(string)
		ret.ServiceIp = in["service_ip"].(string)
		ret.PortCount = in["port_count"].(int)
		ret.VirtualServer = in["virtual_server"].(int)
		ret.Disabled = in["disabled"].(int)
		ret.GslbProtocol = in["gslb_protocol"].(int)
		ret.LocalProtocol = in["local_protocol"].(int)
		ret.ManuallyHealthCheck = in["manually_health_check"].(int)
		ret.Use_gslb_state = in["use_gslb_state"].(int)
		ret.Dynamic = in["dynamic"].(int)
		ret.Hits = in["hits"].(int)
		ret.Recent = in["recent"].(int)
		ret.DevVipPortList = getSliceGslbSiteSlbDevOperVipServerVipServerV6ListOperDevVipPortList(in["dev_vip_port_list"].([]interface{}))
	}
	return ret
}

func getSliceGslbSiteSlbDevOperVipServerVipServerV6ListOperDevVipPortList(d []interface{}) []edpt.GslbSiteSlbDevOperVipServerVipServerV6ListOperDevVipPortList {

	count1 := len(d)
	ret := make([]edpt.GslbSiteSlbDevOperVipServerVipServerV6ListOperDevVipPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteSlbDevOperVipServerVipServerV6ListOperDevVipPortList
		oi.DevVipPortNum = in["dev_vip_port_num"].(int)
		oi.DevVipPortState = in["dev_vip_port_state"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbSiteSlbDevOperVipServerVipServerNameList(d []interface{}) []edpt.GslbSiteSlbDevOperVipServerVipServerNameList {

	count1 := len(d)
	ret := make([]edpt.GslbSiteSlbDevOperVipServerVipServerNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteSlbDevOperVipServerVipServerNameList
		oi.VipName = in["vip_name"].(string)
		oi.Oper = getObjectGslbSiteSlbDevOperVipServerVipServerNameListOper(in["oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbSiteSlbDevOperVipServerVipServerNameListOper(d []interface{}) edpt.GslbSiteSlbDevOperVipServerVipServerNameListOper {

	count1 := len(d)
	var ret edpt.GslbSiteSlbDevOperVipServerVipServerNameListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dev_vip_addr = in["dev_vip_addr"].(string)
		ret.Dev_vip_state = in["dev_vip_state"].(string)
		ret.NodeName = in["node_name"].(string)
		ret.ServiceIp = in["service_ip"].(string)
		ret.PortCount = in["port_count"].(int)
		ret.VirtualServer = in["virtual_server"].(int)
		ret.Disabled = in["disabled"].(int)
		ret.GslbProtocol = in["gslb_protocol"].(int)
		ret.LocalProtocol = in["local_protocol"].(int)
		ret.ManuallyHealthCheck = in["manually_health_check"].(int)
		ret.Use_gslb_state = in["use_gslb_state"].(int)
		ret.Dynamic = in["dynamic"].(int)
		ret.Hits = in["hits"].(int)
		ret.Recent = in["recent"].(int)
		ret.DevVipPortList = getSliceGslbSiteSlbDevOperVipServerVipServerNameListOperDevVipPortList(in["dev_vip_port_list"].([]interface{}))
	}
	return ret
}

func getSliceGslbSiteSlbDevOperVipServerVipServerNameListOperDevVipPortList(d []interface{}) []edpt.GslbSiteSlbDevOperVipServerVipServerNameListOperDevVipPortList {

	count1 := len(d)
	ret := make([]edpt.GslbSiteSlbDevOperVipServerVipServerNameListOperDevVipPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteSlbDevOperVipServerVipServerNameListOperDevVipPortList
		oi.DevVipPortNum = in["dev_vip_port_num"].(int)
		oi.DevVipPortState = in["dev_vip_port_state"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbSiteSlbDevOper(d *schema.ResourceData) edpt.GslbSiteSlbDevOper {
	var ret edpt.GslbSiteSlbDevOper

	ret.DeviceName = d.Get("device_name").(string)

	ret.Oper = getObjectGslbSiteSlbDevOperOper(d.Get("oper").([]interface{}))

	ret.VipServer = getObjectGslbSiteSlbDevOperVipServer(d.Get("vip_server").([]interface{}))

	ret.SiteName = d.Get("site_name").(string)
	return ret
}
