package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6ServerPortOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_server_port_oper`: Operational Status for the object port\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6ServerPortOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"curr_conn_rate": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"conn_rate_unit": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"slow_start_conn_limit": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"curr_observe_rate": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"down_grace_period_allowed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"current_time": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"down_time_grace_period": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"diameter_enabled": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"es_resp_time": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"inband_hm_reassign_num": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"disable": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"hm_key": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"hm_index": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"soft_down_time": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"aflow_conn_limit": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"aflow_queue_size": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"resv_conn": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"auto_nat_addr_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auto_nat_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"vrid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ha_group_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ip_rr": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ports_consumed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ports_consumed_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ports_freed_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"alloc_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"drs_auto_nat_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"drs_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"drs_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"drs_auto_nat_address_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"auto_nat_ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"vrid": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"ha_group_id": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"ip_rr": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"ports_consumed": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"ports_consumed_total": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"ports_freed_total": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"alloc_failed": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
						"pool_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"nat_pool_addr_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"nat_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ports_consumed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ports_consumed_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ports_freed_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"alloc_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"drs_ip_nat_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"drs_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"drs_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pool_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"nat_pool_addr_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"nat_ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ports_consumed": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"ports_consumed_total": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"ports_freed_total": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"alloc_failed": {
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
			"port_number": {
				Type: schema.TypeInt, Required: true, Description: "Port Number",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'tcp': TCP Port; 'udp': UDP Port;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceCgnv6ServerPortOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6ServerPortOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6ServerPortOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6ServerPortOperOper := setObjectCgnv6ServerPortOperOper(res)
		d.Set("oper", Cgnv6ServerPortOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6ServerPortOperOper(ret edpt.DataCgnv6ServerPortOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"state":                     ret.DtCgnv6ServerPortOper.Oper.State,
			"curr_conn_rate":            ret.DtCgnv6ServerPortOper.Oper.Curr_conn_rate,
			"conn_rate_unit":            ret.DtCgnv6ServerPortOper.Oper.Conn_rate_unit,
			"slow_start_conn_limit":     ret.DtCgnv6ServerPortOper.Oper.Slow_start_conn_limit,
			"curr_observe_rate":         ret.DtCgnv6ServerPortOper.Oper.Curr_observe_rate,
			"down_grace_period_allowed": ret.DtCgnv6ServerPortOper.Oper.Down_grace_period_allowed,
			"current_time":              ret.DtCgnv6ServerPortOper.Oper.Current_time,
			"down_time_grace_period":    ret.DtCgnv6ServerPortOper.Oper.Down_time_grace_period,
			"diameter_enabled":          ret.DtCgnv6ServerPortOper.Oper.Diameter_enabled,
			"es_resp_time":              ret.DtCgnv6ServerPortOper.Oper.Es_resp_time,
			"inband_hm_reassign_num":    ret.DtCgnv6ServerPortOper.Oper.Inband_hm_reassign_num,
			"disable":                   ret.DtCgnv6ServerPortOper.Oper.Disable,
			"hm_key":                    ret.DtCgnv6ServerPortOper.Oper.HmKey,
			"hm_index":                  ret.DtCgnv6ServerPortOper.Oper.HmIndex,
			"soft_down_time":            ret.DtCgnv6ServerPortOper.Oper.Soft_down_time,
			"aflow_conn_limit":          ret.DtCgnv6ServerPortOper.Oper.Aflow_conn_limit,
			"aflow_queue_size":          ret.DtCgnv6ServerPortOper.Oper.Aflow_queue_size,
			"resv_conn":                 ret.DtCgnv6ServerPortOper.Oper.Resv_conn,
			"auto_nat_addr_list":        setSliceCgnv6ServerPortOperOperAutoNatAddrList(ret.DtCgnv6ServerPortOper.Oper.AutoNatAddrList),
			"drs_auto_nat_list":         setSliceCgnv6ServerPortOperOperDrsAutoNatList(ret.DtCgnv6ServerPortOper.Oper.DrsAutoNatList),
			"pool_name":                 ret.DtCgnv6ServerPortOper.Oper.Pool_name,
			"nat_pool_addr_list":        setSliceCgnv6ServerPortOperOperNatPoolAddrList(ret.DtCgnv6ServerPortOper.Oper.NatPoolAddrList),
			"drs_ip_nat_list":           setSliceCgnv6ServerPortOperOperDrsIpNatList(ret.DtCgnv6ServerPortOper.Oper.DrsIpNatList),
		},
	}
}

func setSliceCgnv6ServerPortOperOperAutoNatAddrList(d []edpt.Cgnv6ServerPortOperOperAutoNatAddrList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["auto_nat_ip"] = item.Auto_nat_ip
		in["vrid"] = item.Vrid
		in["ha_group_id"] = item.Ha_group_id
		in["ip_rr"] = item.Ip_rr
		in["ports_consumed"] = item.Ports_consumed
		in["ports_consumed_total"] = item.Ports_consumed_total
		in["ports_freed_total"] = item.Ports_freed_total
		in["alloc_failed"] = item.Alloc_failed
		result = append(result, in)
	}
	return result
}

func setSliceCgnv6ServerPortOperOperDrsAutoNatList(d []edpt.Cgnv6ServerPortOperOperDrsAutoNatList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["drs_name"] = item.Drs_name
		in["drs_port"] = item.Drs_port
		in["drs_auto_nat_address_list"] = setSliceCgnv6ServerPortOperOperDrsAutoNatListDrsAutoNatAddressList(item.DrsAutoNatAddressList)
		result = append(result, in)
	}
	return result
}

func setSliceCgnv6ServerPortOperOperDrsAutoNatListDrsAutoNatAddressList(d []edpt.Cgnv6ServerPortOperOperDrsAutoNatListDrsAutoNatAddressList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["auto_nat_ip"] = item.Auto_nat_ip
		in["vrid"] = item.Vrid
		in["ha_group_id"] = item.Ha_group_id
		in["ip_rr"] = item.Ip_rr
		in["ports_consumed"] = item.Ports_consumed
		in["ports_consumed_total"] = item.Ports_consumed_total
		in["ports_freed_total"] = item.Ports_freed_total
		in["alloc_failed"] = item.Alloc_failed
		result = append(result, in)
	}
	return result
}

func setSliceCgnv6ServerPortOperOperNatPoolAddrList(d []edpt.Cgnv6ServerPortOperOperNatPoolAddrList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["nat_ip"] = item.Nat_ip
		in["ports_consumed"] = item.Ports_consumed
		in["ports_consumed_total"] = item.Ports_consumed_total
		in["ports_freed_total"] = item.Ports_freed_total
		in["alloc_failed"] = item.Alloc_failed
		result = append(result, in)
	}
	return result
}

func setSliceCgnv6ServerPortOperOperDrsIpNatList(d []edpt.Cgnv6ServerPortOperOperDrsIpNatList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["drs_name"] = item.Drs_name
		in["drs_port"] = item.Drs_port
		in["pool_name"] = item.Pool_name
		in["nat_pool_addr_list"] = setSliceCgnv6ServerPortOperOperDrsIpNatListNatPoolAddrList(item.NatPoolAddrList)
		result = append(result, in)
	}
	return result
}

func setSliceCgnv6ServerPortOperOperDrsIpNatListNatPoolAddrList(d []edpt.Cgnv6ServerPortOperOperDrsIpNatListNatPoolAddrList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["nat_ip"] = item.Nat_ip
		in["ports_consumed"] = item.Ports_consumed
		in["ports_consumed_total"] = item.Ports_consumed_total
		in["ports_freed_total"] = item.Ports_freed_total
		in["alloc_failed"] = item.Alloc_failed
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6ServerPortOperOper(d []interface{}) edpt.Cgnv6ServerPortOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6ServerPortOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.Curr_conn_rate = in["curr_conn_rate"].(int)
		ret.Conn_rate_unit = in["conn_rate_unit"].(string)
		ret.Slow_start_conn_limit = in["slow_start_conn_limit"].(int)
		ret.Curr_observe_rate = in["curr_observe_rate"].(int)
		ret.Down_grace_period_allowed = in["down_grace_period_allowed"].(int)
		ret.Current_time = in["current_time"].(int)
		ret.Down_time_grace_period = in["down_time_grace_period"].(int)
		ret.Diameter_enabled = in["diameter_enabled"].(int)
		ret.Es_resp_time = in["es_resp_time"].(int)
		ret.Inband_hm_reassign_num = in["inband_hm_reassign_num"].(int)
		ret.Disable = in["disable"].(int)
		ret.HmKey = in["hm_key"].(int)
		ret.HmIndex = in["hm_index"].(int)
		ret.Soft_down_time = in["soft_down_time"].(int)
		ret.Aflow_conn_limit = in["aflow_conn_limit"].(int)
		ret.Aflow_queue_size = in["aflow_queue_size"].(int)
		ret.Resv_conn = in["resv_conn"].(int)
		ret.AutoNatAddrList = getSliceCgnv6ServerPortOperOperAutoNatAddrList(in["auto_nat_addr_list"].([]interface{}))
		ret.DrsAutoNatList = getSliceCgnv6ServerPortOperOperDrsAutoNatList(in["drs_auto_nat_list"].([]interface{}))
		ret.Pool_name = in["pool_name"].(string)
		ret.NatPoolAddrList = getSliceCgnv6ServerPortOperOperNatPoolAddrList(in["nat_pool_addr_list"].([]interface{}))
		ret.DrsIpNatList = getSliceCgnv6ServerPortOperOperDrsIpNatList(in["drs_ip_nat_list"].([]interface{}))
	}
	return ret
}

func getSliceCgnv6ServerPortOperOperAutoNatAddrList(d []interface{}) []edpt.Cgnv6ServerPortOperOperAutoNatAddrList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6ServerPortOperOperAutoNatAddrList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6ServerPortOperOperAutoNatAddrList
		oi.Auto_nat_ip = in["auto_nat_ip"].(string)
		oi.Vrid = in["vrid"].(int)
		oi.Ha_group_id = in["ha_group_id"].(int)
		oi.Ip_rr = in["ip_rr"].(int)
		oi.Ports_consumed = in["ports_consumed"].(int)
		oi.Ports_consumed_total = in["ports_consumed_total"].(int)
		oi.Ports_freed_total = in["ports_freed_total"].(int)
		oi.Alloc_failed = in["alloc_failed"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6ServerPortOperOperDrsAutoNatList(d []interface{}) []edpt.Cgnv6ServerPortOperOperDrsAutoNatList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6ServerPortOperOperDrsAutoNatList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6ServerPortOperOperDrsAutoNatList
		oi.Drs_name = in["drs_name"].(string)
		oi.Drs_port = in["drs_port"].(int)
		oi.DrsAutoNatAddressList = getSliceCgnv6ServerPortOperOperDrsAutoNatListDrsAutoNatAddressList(in["drs_auto_nat_address_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6ServerPortOperOperDrsAutoNatListDrsAutoNatAddressList(d []interface{}) []edpt.Cgnv6ServerPortOperOperDrsAutoNatListDrsAutoNatAddressList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6ServerPortOperOperDrsAutoNatListDrsAutoNatAddressList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6ServerPortOperOperDrsAutoNatListDrsAutoNatAddressList
		oi.Auto_nat_ip = in["auto_nat_ip"].(string)
		oi.Vrid = in["vrid"].(int)
		oi.Ha_group_id = in["ha_group_id"].(int)
		oi.Ip_rr = in["ip_rr"].(int)
		oi.Ports_consumed = in["ports_consumed"].(int)
		oi.Ports_consumed_total = in["ports_consumed_total"].(int)
		oi.Ports_freed_total = in["ports_freed_total"].(int)
		oi.Alloc_failed = in["alloc_failed"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6ServerPortOperOperNatPoolAddrList(d []interface{}) []edpt.Cgnv6ServerPortOperOperNatPoolAddrList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6ServerPortOperOperNatPoolAddrList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6ServerPortOperOperNatPoolAddrList
		oi.Nat_ip = in["nat_ip"].(string)
		oi.Ports_consumed = in["ports_consumed"].(int)
		oi.Ports_consumed_total = in["ports_consumed_total"].(int)
		oi.Ports_freed_total = in["ports_freed_total"].(int)
		oi.Alloc_failed = in["alloc_failed"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6ServerPortOperOperDrsIpNatList(d []interface{}) []edpt.Cgnv6ServerPortOperOperDrsIpNatList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6ServerPortOperOperDrsIpNatList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6ServerPortOperOperDrsIpNatList
		oi.Drs_name = in["drs_name"].(string)
		oi.Drs_port = in["drs_port"].(int)
		oi.Pool_name = in["pool_name"].(string)
		oi.NatPoolAddrList = getSliceCgnv6ServerPortOperOperDrsIpNatListNatPoolAddrList(in["nat_pool_addr_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6ServerPortOperOperDrsIpNatListNatPoolAddrList(d []interface{}) []edpt.Cgnv6ServerPortOperOperDrsIpNatListNatPoolAddrList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6ServerPortOperOperDrsIpNatListNatPoolAddrList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6ServerPortOperOperDrsIpNatListNatPoolAddrList
		oi.Nat_ip = in["nat_ip"].(string)
		oi.Ports_consumed = in["ports_consumed"].(int)
		oi.Ports_consumed_total = in["ports_consumed_total"].(int)
		oi.Ports_freed_total = in["ports_freed_total"].(int)
		oi.Alloc_failed = in["alloc_failed"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6ServerPortOper(d *schema.ResourceData) edpt.Cgnv6ServerPortOper {
	var ret edpt.Cgnv6ServerPortOper

	ret.Oper = getObjectCgnv6ServerPortOperOper(d.Get("oper").([]interface{}))

	ret.PortNumber = d.Get("port_number").(int)

	ret.Protocol = d.Get("protocol").(string)

	ret.Name = d.Get("name").(string)
	return ret
}
