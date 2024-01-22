package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbSiteOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_site_oper`: Operational Status for the object site\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbSiteOperRead,

		Schema: map[string]*schema.Schema{
			"ip_server_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_server_name": {
							Type: schema.TypeString, Required: true, Description: "Specify the real server name",
						},
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_server": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ip_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"state": {
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
									"ip_server_port": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"vport": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"vport_state": {
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
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gslb_site": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"state": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"type_last": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"last": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
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
			"site_name": {
				Type: schema.TypeString, Required: true, Description: "Specify GSLB site name",
			},
			"slb_dev_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
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
					},
				},
			},
		},
	}
}

func resourceGslbSiteOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbSiteOperIpServerList := setSliceGslbSiteOperIpServerList(res)
		d.Set("ip_server_list", GslbSiteOperIpServerList)
		GslbSiteOperOper := setObjectGslbSiteOperOper(res)
		d.Set("oper", GslbSiteOperOper)
		GslbSiteOperSlbDevList := setSliceGslbSiteOperSlbDevList(res)
		d.Set("slb_dev_list", GslbSiteOperSlbDevList)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceGslbSiteOperIpServerList(d edpt.DataGslbSiteOper) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtGslbSiteOper.IpServerList {
		in := make(map[string]interface{})
		in["ip_server_name"] = item.IpServerName
		in["oper"] = setObjectGslbSiteOperIpServerListOper(item.Oper)
		result = append(result, in)
	}
	return result
}

func setObjectGslbSiteOperIpServerListOper(d edpt.GslbSiteOperIpServerListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["ip_server"] = d.IpServer

	in["ip_address"] = d.IpAddress

	in["state"] = d.State

	in["service_ip"] = d.ServiceIp

	in["port_count"] = d.PortCount

	in["virtual_server"] = d.VirtualServer

	in["disabled"] = d.Disabled

	in["gslb_protocol"] = d.GslbProtocol

	in["local_protocol"] = d.LocalProtocol

	in["manually_health_check"] = d.ManuallyHealthCheck

	in["use_gslb_state"] = d.Use_gslb_state

	in["dynamic"] = d.Dynamic
	in["ip_server_port"] = setSliceGslbSiteOperIpServerListOperIpServerPort(d.IpServerPort)
	result = append(result, in)
	return result
}

func setSliceGslbSiteOperIpServerListOperIpServerPort(d []edpt.GslbSiteOperIpServerListOperIpServerPort) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["vport"] = item.Vport
		in["vport_state"] = item.VportState
		result = append(result, in)
	}
	return result
}

func setObjectGslbSiteOperOper(ret edpt.DataGslbSiteOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"gslb_site":        ret.DtGslbSiteOper.Oper.GslbSite,
			"state":            ret.DtGslbSiteOper.Oper.State,
			"type_last":        setSliceGslbSiteOperOperTypeLast(ret.DtGslbSiteOper.Oper.TypeLast),
			"client_ldns_list": setSliceGslbSiteOperOperClientLdnsList(ret.DtGslbSiteOper.Oper.ClientLdnsList),
		},
	}
}

func setSliceGslbSiteOperOperTypeLast(d []edpt.GslbSiteOperOperTypeLast) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["type"] = item.Type
		in["last"] = item.Last
		result = append(result, in)
	}
	return result
}

func setSliceGslbSiteOperOperClientLdnsList(d []edpt.GslbSiteOperOperClientLdnsList) []map[string]interface{} {
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

func setSliceGslbSiteOperSlbDevList(d edpt.DataGslbSiteOper) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtGslbSiteOper.SlbDevList {
		in := make(map[string]interface{})
		in["device_name"] = item.DeviceName
		in["oper"] = setObjectGslbSiteOperSlbDevListOper(item.Oper)
		in["vip_server"] = setObjectGslbSiteOperSlbDevListVipServer(item.VipServer)
		result = append(result, in)
	}
	return result
}

func setObjectGslbSiteOperSlbDevListOper(d edpt.GslbSiteOperSlbDevListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["dev_name"] = d.Dev_name

	in["dev_ip"] = d.Dev_ip

	in["dev_attr"] = d.Dev_attr

	in["dev_admin_preference"] = d.Dev_admin_preference

	in["dev_session_num"] = d.Dev_session_num

	in["dev_session_util"] = d.Dev_session_util

	in["dev_gw_state"] = d.Dev_gw_state

	in["dev_ip_cnt"] = d.Dev_ip_cnt

	in["dev_state"] = d.DevState
	in["client_ldns_list"] = setSliceGslbSiteOperSlbDevListOperClientLdnsList(d.ClientLdnsList)
	result = append(result, in)
	return result
}

func setSliceGslbSiteOperSlbDevListOperClientLdnsList(d []edpt.GslbSiteOperSlbDevListOperClientLdnsList) []map[string]interface{} {
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

func setObjectGslbSiteOperSlbDevListVipServer(d edpt.GslbSiteOperSlbDevListVipServer) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectGslbSiteOperSlbDevListVipServerOper(d.Oper)
	in["vip_server_v4_list"] = setSliceGslbSiteOperSlbDevListVipServerVipServerV4List(d.VipServerV4List)
	in["vip_server_v6_list"] = setSliceGslbSiteOperSlbDevListVipServerVipServerV6List(d.VipServerV6List)
	in["vip_server_name_list"] = setSliceGslbSiteOperSlbDevListVipServerVipServerNameList(d.VipServerNameList)
	result = append(result, in)
	return result
}

func setObjectGslbSiteOperSlbDevListVipServerOper(d edpt.GslbSiteOperSlbDevListVipServerOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	result = append(result, in)
	return result
}

func setSliceGslbSiteOperSlbDevListVipServerVipServerV4List(d []edpt.GslbSiteOperSlbDevListVipServerVipServerV4List) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ipv4"] = item.Ipv4
		in["oper"] = setObjectGslbSiteOperSlbDevListVipServerVipServerV4ListOper(item.Oper)
		result = append(result, in)
	}
	return result
}

func setObjectGslbSiteOperSlbDevListVipServerVipServerV4ListOper(d edpt.GslbSiteOperSlbDevListVipServerVipServerV4ListOper) []map[string]interface{} {
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
	in["dev_vip_port_list"] = setSliceGslbSiteOperSlbDevListVipServerVipServerV4ListOperDevVipPortList(d.DevVipPortList)
	result = append(result, in)
	return result
}

func setSliceGslbSiteOperSlbDevListVipServerVipServerV4ListOperDevVipPortList(d []edpt.GslbSiteOperSlbDevListVipServerVipServerV4ListOperDevVipPortList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dev_vip_port_num"] = item.DevVipPortNum
		in["dev_vip_port_state"] = item.DevVipPortState
		result = append(result, in)
	}
	return result
}

func setSliceGslbSiteOperSlbDevListVipServerVipServerV6List(d []edpt.GslbSiteOperSlbDevListVipServerVipServerV6List) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ipv6"] = item.Ipv6
		in["oper"] = setObjectGslbSiteOperSlbDevListVipServerVipServerV6ListOper(item.Oper)
		result = append(result, in)
	}
	return result
}

func setObjectGslbSiteOperSlbDevListVipServerVipServerV6ListOper(d edpt.GslbSiteOperSlbDevListVipServerVipServerV6ListOper) []map[string]interface{} {
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
	in["dev_vip_port_list"] = setSliceGslbSiteOperSlbDevListVipServerVipServerV6ListOperDevVipPortList(d.DevVipPortList)
	result = append(result, in)
	return result
}

func setSliceGslbSiteOperSlbDevListVipServerVipServerV6ListOperDevVipPortList(d []edpt.GslbSiteOperSlbDevListVipServerVipServerV6ListOperDevVipPortList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dev_vip_port_num"] = item.DevVipPortNum
		in["dev_vip_port_state"] = item.DevVipPortState
		result = append(result, in)
	}
	return result
}

func setSliceGslbSiteOperSlbDevListVipServerVipServerNameList(d []edpt.GslbSiteOperSlbDevListVipServerVipServerNameList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["vip_name"] = item.VipName
		in["oper"] = setObjectGslbSiteOperSlbDevListVipServerVipServerNameListOper(item.Oper)
		result = append(result, in)
	}
	return result
}

func setObjectGslbSiteOperSlbDevListVipServerVipServerNameListOper(d edpt.GslbSiteOperSlbDevListVipServerVipServerNameListOper) []map[string]interface{} {
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
	in["dev_vip_port_list"] = setSliceGslbSiteOperSlbDevListVipServerVipServerNameListOperDevVipPortList(d.DevVipPortList)
	result = append(result, in)
	return result
}

func setSliceGslbSiteOperSlbDevListVipServerVipServerNameListOperDevVipPortList(d []edpt.GslbSiteOperSlbDevListVipServerVipServerNameListOperDevVipPortList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dev_vip_port_num"] = item.DevVipPortNum
		in["dev_vip_port_state"] = item.DevVipPortState
		result = append(result, in)
	}
	return result
}

func getSliceGslbSiteOperIpServerList(d []interface{}) []edpt.GslbSiteOperIpServerList {

	count1 := len(d)
	ret := make([]edpt.GslbSiteOperIpServerList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteOperIpServerList
		oi.IpServerName = in["ip_server_name"].(string)
		oi.Oper = getObjectGslbSiteOperIpServerListOper(in["oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbSiteOperIpServerListOper(d []interface{}) edpt.GslbSiteOperIpServerListOper {

	count1 := len(d)
	var ret edpt.GslbSiteOperIpServerListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpServer = in["ip_server"].(string)
		ret.IpAddress = in["ip_address"].(string)
		ret.State = in["state"].(string)
		ret.ServiceIp = in["service_ip"].(string)
		ret.PortCount = in["port_count"].(int)
		ret.VirtualServer = in["virtual_server"].(int)
		ret.Disabled = in["disabled"].(int)
		ret.GslbProtocol = in["gslb_protocol"].(int)
		ret.LocalProtocol = in["local_protocol"].(int)
		ret.ManuallyHealthCheck = in["manually_health_check"].(int)
		ret.Use_gslb_state = in["use_gslb_state"].(int)
		ret.Dynamic = in["dynamic"].(int)
		ret.IpServerPort = getSliceGslbSiteOperIpServerListOperIpServerPort(in["ip_server_port"].([]interface{}))
	}
	return ret
}

func getSliceGslbSiteOperIpServerListOperIpServerPort(d []interface{}) []edpt.GslbSiteOperIpServerListOperIpServerPort {

	count1 := len(d)
	ret := make([]edpt.GslbSiteOperIpServerListOperIpServerPort, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteOperIpServerListOperIpServerPort
		oi.Vport = in["vport"].(int)
		oi.VportState = in["vport_state"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbSiteOperOper(d []interface{}) edpt.GslbSiteOperOper {

	count1 := len(d)
	var ret edpt.GslbSiteOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GslbSite = in["gslb_site"].(string)
		ret.State = in["state"].(string)
		ret.TypeLast = getSliceGslbSiteOperOperTypeLast(in["type_last"].([]interface{}))
		ret.ClientLdnsList = getSliceGslbSiteOperOperClientLdnsList(in["client_ldns_list"].([]interface{}))
	}
	return ret
}

func getSliceGslbSiteOperOperTypeLast(d []interface{}) []edpt.GslbSiteOperOperTypeLast {

	count1 := len(d)
	ret := make([]edpt.GslbSiteOperOperTypeLast, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteOperOperTypeLast
		oi.Type = in["type"].(string)
		oi.Last = in["last"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbSiteOperOperClientLdnsList(d []interface{}) []edpt.GslbSiteOperOperClientLdnsList {

	count1 := len(d)
	ret := make([]edpt.GslbSiteOperOperClientLdnsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteOperOperClientLdnsList
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

func getSliceGslbSiteOperSlbDevList(d []interface{}) []edpt.GslbSiteOperSlbDevList {

	count1 := len(d)
	ret := make([]edpt.GslbSiteOperSlbDevList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteOperSlbDevList
		oi.DeviceName = in["device_name"].(string)
		oi.Oper = getObjectGslbSiteOperSlbDevListOper(in["oper"].([]interface{}))
		oi.VipServer = getObjectGslbSiteOperSlbDevListVipServer(in["vip_server"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbSiteOperSlbDevListOper(d []interface{}) edpt.GslbSiteOperSlbDevListOper {

	count1 := len(d)
	var ret edpt.GslbSiteOperSlbDevListOper
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
		ret.ClientLdnsList = getSliceGslbSiteOperSlbDevListOperClientLdnsList(in["client_ldns_list"].([]interface{}))
	}
	return ret
}

func getSliceGslbSiteOperSlbDevListOperClientLdnsList(d []interface{}) []edpt.GslbSiteOperSlbDevListOperClientLdnsList {

	count1 := len(d)
	ret := make([]edpt.GslbSiteOperSlbDevListOperClientLdnsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteOperSlbDevListOperClientLdnsList
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

func getObjectGslbSiteOperSlbDevListVipServer(d []interface{}) edpt.GslbSiteOperSlbDevListVipServer {

	count1 := len(d)
	var ret edpt.GslbSiteOperSlbDevListVipServer
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectGslbSiteOperSlbDevListVipServerOper(in["oper"].([]interface{}))
		ret.VipServerV4List = getSliceGslbSiteOperSlbDevListVipServerVipServerV4List(in["vip_server_v4_list"].([]interface{}))
		ret.VipServerV6List = getSliceGslbSiteOperSlbDevListVipServerVipServerV6List(in["vip_server_v6_list"].([]interface{}))
		ret.VipServerNameList = getSliceGslbSiteOperSlbDevListVipServerVipServerNameList(in["vip_server_name_list"].([]interface{}))
	}
	return ret
}

func getObjectGslbSiteOperSlbDevListVipServerOper(d []interface{}) edpt.GslbSiteOperSlbDevListVipServerOper {

	var ret edpt.GslbSiteOperSlbDevListVipServerOper
	return ret
}

func getSliceGslbSiteOperSlbDevListVipServerVipServerV4List(d []interface{}) []edpt.GslbSiteOperSlbDevListVipServerVipServerV4List {

	count1 := len(d)
	ret := make([]edpt.GslbSiteOperSlbDevListVipServerVipServerV4List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteOperSlbDevListVipServerVipServerV4List
		oi.Ipv4 = in["ipv4"].(string)
		oi.Oper = getObjectGslbSiteOperSlbDevListVipServerVipServerV4ListOper(in["oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbSiteOperSlbDevListVipServerVipServerV4ListOper(d []interface{}) edpt.GslbSiteOperSlbDevListVipServerVipServerV4ListOper {

	count1 := len(d)
	var ret edpt.GslbSiteOperSlbDevListVipServerVipServerV4ListOper
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
		ret.DevVipPortList = getSliceGslbSiteOperSlbDevListVipServerVipServerV4ListOperDevVipPortList(in["dev_vip_port_list"].([]interface{}))
	}
	return ret
}

func getSliceGslbSiteOperSlbDevListVipServerVipServerV4ListOperDevVipPortList(d []interface{}) []edpt.GslbSiteOperSlbDevListVipServerVipServerV4ListOperDevVipPortList {

	count1 := len(d)
	ret := make([]edpt.GslbSiteOperSlbDevListVipServerVipServerV4ListOperDevVipPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteOperSlbDevListVipServerVipServerV4ListOperDevVipPortList
		oi.DevVipPortNum = in["dev_vip_port_num"].(int)
		oi.DevVipPortState = in["dev_vip_port_state"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbSiteOperSlbDevListVipServerVipServerV6List(d []interface{}) []edpt.GslbSiteOperSlbDevListVipServerVipServerV6List {

	count1 := len(d)
	ret := make([]edpt.GslbSiteOperSlbDevListVipServerVipServerV6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteOperSlbDevListVipServerVipServerV6List
		oi.Ipv6 = in["ipv6"].(string)
		oi.Oper = getObjectGslbSiteOperSlbDevListVipServerVipServerV6ListOper(in["oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbSiteOperSlbDevListVipServerVipServerV6ListOper(d []interface{}) edpt.GslbSiteOperSlbDevListVipServerVipServerV6ListOper {

	count1 := len(d)
	var ret edpt.GslbSiteOperSlbDevListVipServerVipServerV6ListOper
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
		ret.DevVipPortList = getSliceGslbSiteOperSlbDevListVipServerVipServerV6ListOperDevVipPortList(in["dev_vip_port_list"].([]interface{}))
	}
	return ret
}

func getSliceGslbSiteOperSlbDevListVipServerVipServerV6ListOperDevVipPortList(d []interface{}) []edpt.GslbSiteOperSlbDevListVipServerVipServerV6ListOperDevVipPortList {

	count1 := len(d)
	ret := make([]edpt.GslbSiteOperSlbDevListVipServerVipServerV6ListOperDevVipPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteOperSlbDevListVipServerVipServerV6ListOperDevVipPortList
		oi.DevVipPortNum = in["dev_vip_port_num"].(int)
		oi.DevVipPortState = in["dev_vip_port_state"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbSiteOperSlbDevListVipServerVipServerNameList(d []interface{}) []edpt.GslbSiteOperSlbDevListVipServerVipServerNameList {

	count1 := len(d)
	ret := make([]edpt.GslbSiteOperSlbDevListVipServerVipServerNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteOperSlbDevListVipServerVipServerNameList
		oi.VipName = in["vip_name"].(string)
		oi.Oper = getObjectGslbSiteOperSlbDevListVipServerVipServerNameListOper(in["oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbSiteOperSlbDevListVipServerVipServerNameListOper(d []interface{}) edpt.GslbSiteOperSlbDevListVipServerVipServerNameListOper {

	count1 := len(d)
	var ret edpt.GslbSiteOperSlbDevListVipServerVipServerNameListOper
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
		ret.DevVipPortList = getSliceGslbSiteOperSlbDevListVipServerVipServerNameListOperDevVipPortList(in["dev_vip_port_list"].([]interface{}))
	}
	return ret
}

func getSliceGslbSiteOperSlbDevListVipServerVipServerNameListOperDevVipPortList(d []interface{}) []edpt.GslbSiteOperSlbDevListVipServerVipServerNameListOperDevVipPortList {

	count1 := len(d)
	ret := make([]edpt.GslbSiteOperSlbDevListVipServerVipServerNameListOperDevVipPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteOperSlbDevListVipServerVipServerNameListOperDevVipPortList
		oi.DevVipPortNum = in["dev_vip_port_num"].(int)
		oi.DevVipPortState = in["dev_vip_port_state"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbSiteOper(d *schema.ResourceData) edpt.GslbSiteOper {
	var ret edpt.GslbSiteOper

	ret.IpServerList = getSliceGslbSiteOperIpServerList(d.Get("ip_server_list").([]interface{}))

	ret.Oper = getObjectGslbSiteOperOper(d.Get("oper").([]interface{}))

	ret.SiteName = d.Get("site_name").(string)

	ret.SlbDevList = getSliceGslbSiteOperSlbDevList(d.Get("slb_dev_list").([]interface{}))
	return ret
}
