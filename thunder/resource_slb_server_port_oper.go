package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbServerPortOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_server_port_oper`: Operational Status for the object port\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbServerPortOperRead,

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

func resourceSlbServerPortOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServerPortOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServerPortOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbServerPortOperOper := setObjectSlbServerPortOperOper(res)
		d.Set("oper", SlbServerPortOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbServerPortOperOper(ret edpt.DataSlbServerPortOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"state":                     ret.DtSlbServerPortOper.Oper.State,
			"curr_conn_rate":            ret.DtSlbServerPortOper.Oper.Curr_conn_rate,
			"conn_rate_unit":            ret.DtSlbServerPortOper.Oper.Conn_rate_unit,
			"slow_start_conn_limit":     ret.DtSlbServerPortOper.Oper.Slow_start_conn_limit,
			"curr_observe_rate":         ret.DtSlbServerPortOper.Oper.Curr_observe_rate,
			"down_grace_period_allowed": ret.DtSlbServerPortOper.Oper.Down_grace_period_allowed,
			"current_time":              ret.DtSlbServerPortOper.Oper.Current_time,
			"down_time_grace_period":    ret.DtSlbServerPortOper.Oper.Down_time_grace_period,
			"diameter_enabled":          ret.DtSlbServerPortOper.Oper.Diameter_enabled,
			"es_resp_time":              ret.DtSlbServerPortOper.Oper.Es_resp_time,
			"inband_hm_reassign_num":    ret.DtSlbServerPortOper.Oper.Inband_hm_reassign_num,
			"disable":                   ret.DtSlbServerPortOper.Oper.Disable,
			"hm_key":                    ret.DtSlbServerPortOper.Oper.HmKey,
			"hm_index":                  ret.DtSlbServerPortOper.Oper.HmIndex,
			"soft_down_time":            ret.DtSlbServerPortOper.Oper.Soft_down_time,
			"aflow_conn_limit":          ret.DtSlbServerPortOper.Oper.Aflow_conn_limit,
			"aflow_queue_size":          ret.DtSlbServerPortOper.Oper.Aflow_queue_size,
			"resv_conn":                 ret.DtSlbServerPortOper.Oper.Resv_conn,
			"auto_nat_addr_list":        setSliceSlbServerPortOperOperAutoNatAddrList(ret.DtSlbServerPortOper.Oper.AutoNatAddrList),
			"drs_auto_nat_list":         setSliceSlbServerPortOperOperDrsAutoNatList(ret.DtSlbServerPortOper.Oper.DrsAutoNatList),
			"pool_name":                 ret.DtSlbServerPortOper.Oper.Pool_name,
			"nat_pool_addr_list":        setSliceSlbServerPortOperOperNatPoolAddrList(ret.DtSlbServerPortOper.Oper.NatPoolAddrList),
			"drs_ip_nat_list":           setSliceSlbServerPortOperOperDrsIpNatList(ret.DtSlbServerPortOper.Oper.DrsIpNatList),
		},
	}
}

func setSliceSlbServerPortOperOperAutoNatAddrList(d []edpt.SlbServerPortOperOperAutoNatAddrList) []map[string]interface{} {
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

func setSliceSlbServerPortOperOperDrsAutoNatList(d []edpt.SlbServerPortOperOperDrsAutoNatList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["drs_name"] = item.Drs_name
		in["drs_port"] = item.Drs_port
		in["drs_auto_nat_address_list"] = setSliceSlbServerPortOperOperDrsAutoNatListDrsAutoNatAddressList(item.DrsAutoNatAddressList)
		result = append(result, in)
	}
	return result
}

func setSliceSlbServerPortOperOperDrsAutoNatListDrsAutoNatAddressList(d []edpt.SlbServerPortOperOperDrsAutoNatListDrsAutoNatAddressList) []map[string]interface{} {
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

func setSliceSlbServerPortOperOperNatPoolAddrList(d []edpt.SlbServerPortOperOperNatPoolAddrList) []map[string]interface{} {
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

func setSliceSlbServerPortOperOperDrsIpNatList(d []edpt.SlbServerPortOperOperDrsIpNatList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["drs_name"] = item.Drs_name
		in["drs_port"] = item.Drs_port
		in["pool_name"] = item.Pool_name
		in["nat_pool_addr_list"] = setSliceSlbServerPortOperOperDrsIpNatListNatPoolAddrList(item.NatPoolAddrList)
		result = append(result, in)
	}
	return result
}

func setSliceSlbServerPortOperOperDrsIpNatListNatPoolAddrList(d []edpt.SlbServerPortOperOperDrsIpNatListNatPoolAddrList) []map[string]interface{} {
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

func getObjectSlbServerPortOperOper(d []interface{}) edpt.SlbServerPortOperOper {

	count1 := len(d)
	var ret edpt.SlbServerPortOperOper
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
		ret.AutoNatAddrList = getSliceSlbServerPortOperOperAutoNatAddrList(in["auto_nat_addr_list"].([]interface{}))
		ret.DrsAutoNatList = getSliceSlbServerPortOperOperDrsAutoNatList(in["drs_auto_nat_list"].([]interface{}))
		ret.Pool_name = in["pool_name"].(string)
		ret.NatPoolAddrList = getSliceSlbServerPortOperOperNatPoolAddrList(in["nat_pool_addr_list"].([]interface{}))
		ret.DrsIpNatList = getSliceSlbServerPortOperOperDrsIpNatList(in["drs_ip_nat_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbServerPortOperOperAutoNatAddrList(d []interface{}) []edpt.SlbServerPortOperOperAutoNatAddrList {

	count1 := len(d)
	ret := make([]edpt.SlbServerPortOperOperAutoNatAddrList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerPortOperOperAutoNatAddrList
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

func getSliceSlbServerPortOperOperDrsAutoNatList(d []interface{}) []edpt.SlbServerPortOperOperDrsAutoNatList {

	count1 := len(d)
	ret := make([]edpt.SlbServerPortOperOperDrsAutoNatList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerPortOperOperDrsAutoNatList
		oi.Drs_name = in["drs_name"].(string)
		oi.Drs_port = in["drs_port"].(int)
		oi.DrsAutoNatAddressList = getSliceSlbServerPortOperOperDrsAutoNatListDrsAutoNatAddressList(in["drs_auto_nat_address_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbServerPortOperOperDrsAutoNatListDrsAutoNatAddressList(d []interface{}) []edpt.SlbServerPortOperOperDrsAutoNatListDrsAutoNatAddressList {

	count1 := len(d)
	ret := make([]edpt.SlbServerPortOperOperDrsAutoNatListDrsAutoNatAddressList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerPortOperOperDrsAutoNatListDrsAutoNatAddressList
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

func getSliceSlbServerPortOperOperNatPoolAddrList(d []interface{}) []edpt.SlbServerPortOperOperNatPoolAddrList {

	count1 := len(d)
	ret := make([]edpt.SlbServerPortOperOperNatPoolAddrList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerPortOperOperNatPoolAddrList
		oi.Nat_ip = in["nat_ip"].(string)
		oi.Ports_consumed = in["ports_consumed"].(int)
		oi.Ports_consumed_total = in["ports_consumed_total"].(int)
		oi.Ports_freed_total = in["ports_freed_total"].(int)
		oi.Alloc_failed = in["alloc_failed"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbServerPortOperOperDrsIpNatList(d []interface{}) []edpt.SlbServerPortOperOperDrsIpNatList {

	count1 := len(d)
	ret := make([]edpt.SlbServerPortOperOperDrsIpNatList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerPortOperOperDrsIpNatList
		oi.Drs_name = in["drs_name"].(string)
		oi.Drs_port = in["drs_port"].(int)
		oi.Pool_name = in["pool_name"].(string)
		oi.NatPoolAddrList = getSliceSlbServerPortOperOperDrsIpNatListNatPoolAddrList(in["nat_pool_addr_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbServerPortOperOperDrsIpNatListNatPoolAddrList(d []interface{}) []edpt.SlbServerPortOperOperDrsIpNatListNatPoolAddrList {

	count1 := len(d)
	ret := make([]edpt.SlbServerPortOperOperDrsIpNatListNatPoolAddrList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerPortOperOperDrsIpNatListNatPoolAddrList
		oi.Nat_ip = in["nat_ip"].(string)
		oi.Ports_consumed = in["ports_consumed"].(int)
		oi.Ports_consumed_total = in["ports_consumed_total"].(int)
		oi.Ports_freed_total = in["ports_freed_total"].(int)
		oi.Alloc_failed = in["alloc_failed"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbServerPortOper(d *schema.ResourceData) edpt.SlbServerPortOper {
	var ret edpt.SlbServerPortOper

	ret.Oper = getObjectSlbServerPortOperOper(d.Get("oper").([]interface{}))

	ret.PortNumber = d.Get("port_number").(int)

	ret.Protocol = d.Get("protocol").(string)

	ret.Name = d.Get("name").(string)
	return ret
}
