package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbServerServiceOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_server_service_oper`: Operational Status for the object service\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbServerServiceOperRead,

		Schema: map[string]*schema.Schema{
			"label": {
				Type: schema.TypeString, Required: true, Description: "Service Label",
			},
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
			"server_name": {
				Type: schema.TypeString, Required: true, Description: "Server_name",
			},
		},
	}
}

func resourceSlbServerServiceOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServerServiceOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServerServiceOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbServerServiceOperOper := setObjectSlbServerServiceOperOper(res)
		d.Set("oper", SlbServerServiceOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbServerServiceOperOper(ret edpt.DataSlbServerServiceOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"state":                     ret.DtSlbServerServiceOper.Oper.State,
			"curr_conn_rate":            ret.DtSlbServerServiceOper.Oper.Curr_conn_rate,
			"conn_rate_unit":            ret.DtSlbServerServiceOper.Oper.Conn_rate_unit,
			"slow_start_conn_limit":     ret.DtSlbServerServiceOper.Oper.Slow_start_conn_limit,
			"curr_observe_rate":         ret.DtSlbServerServiceOper.Oper.Curr_observe_rate,
			"down_grace_period_allowed": ret.DtSlbServerServiceOper.Oper.Down_grace_period_allowed,
			"current_time":              ret.DtSlbServerServiceOper.Oper.Current_time,
			"down_time_grace_period":    ret.DtSlbServerServiceOper.Oper.Down_time_grace_period,
			"diameter_enabled":          ret.DtSlbServerServiceOper.Oper.Diameter_enabled,
			"es_resp_time":              ret.DtSlbServerServiceOper.Oper.Es_resp_time,
			"inband_hm_reassign_num":    ret.DtSlbServerServiceOper.Oper.Inband_hm_reassign_num,
			"disable":                   ret.DtSlbServerServiceOper.Oper.Disable,
			"hm_key":                    ret.DtSlbServerServiceOper.Oper.HmKey,
			"hm_index":                  ret.DtSlbServerServiceOper.Oper.HmIndex,
			"soft_down_time":            ret.DtSlbServerServiceOper.Oper.Soft_down_time,
			"aflow_conn_limit":          ret.DtSlbServerServiceOper.Oper.Aflow_conn_limit,
			"aflow_queue_size":          ret.DtSlbServerServiceOper.Oper.Aflow_queue_size,
			"resv_conn":                 ret.DtSlbServerServiceOper.Oper.Resv_conn,
			"auto_nat_addr_list":        setSliceSlbServerServiceOperOperAutoNatAddrList(ret.DtSlbServerServiceOper.Oper.AutoNatAddrList),
			"drs_auto_nat_list":         setSliceSlbServerServiceOperOperDrsAutoNatList(ret.DtSlbServerServiceOper.Oper.DrsAutoNatList),
			"pool_name":                 ret.DtSlbServerServiceOper.Oper.Pool_name,
			"nat_pool_addr_list":        setSliceSlbServerServiceOperOperNatPoolAddrList(ret.DtSlbServerServiceOper.Oper.NatPoolAddrList),
			"drs_ip_nat_list":           setSliceSlbServerServiceOperOperDrsIpNatList(ret.DtSlbServerServiceOper.Oper.DrsIpNatList),
		},
	}
}

func setSliceSlbServerServiceOperOperAutoNatAddrList(d []edpt.SlbServerServiceOperOperAutoNatAddrList) []map[string]interface{} {
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

func setSliceSlbServerServiceOperOperDrsAutoNatList(d []edpt.SlbServerServiceOperOperDrsAutoNatList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["drs_name"] = item.Drs_name
		in["drs_port"] = item.Drs_port
		in["drs_auto_nat_address_list"] = setSliceSlbServerServiceOperOperDrsAutoNatListDrsAutoNatAddressList(item.DrsAutoNatAddressList)
		result = append(result, in)
	}
	return result
}

func setSliceSlbServerServiceOperOperDrsAutoNatListDrsAutoNatAddressList(d []edpt.SlbServerServiceOperOperDrsAutoNatListDrsAutoNatAddressList) []map[string]interface{} {
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

func setSliceSlbServerServiceOperOperNatPoolAddrList(d []edpt.SlbServerServiceOperOperNatPoolAddrList) []map[string]interface{} {
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

func setSliceSlbServerServiceOperOperDrsIpNatList(d []edpt.SlbServerServiceOperOperDrsIpNatList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["drs_name"] = item.Drs_name
		in["drs_port"] = item.Drs_port
		in["pool_name"] = item.Pool_name
		in["nat_pool_addr_list"] = setSliceSlbServerServiceOperOperDrsIpNatListNatPoolAddrList(item.NatPoolAddrList)
		result = append(result, in)
	}
	return result
}

func setSliceSlbServerServiceOperOperDrsIpNatListNatPoolAddrList(d []edpt.SlbServerServiceOperOperDrsIpNatListNatPoolAddrList) []map[string]interface{} {
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

func getObjectSlbServerServiceOperOper(d []interface{}) edpt.SlbServerServiceOperOper {

	count1 := len(d)
	var ret edpt.SlbServerServiceOperOper
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
		ret.AutoNatAddrList = getSliceSlbServerServiceOperOperAutoNatAddrList(in["auto_nat_addr_list"].([]interface{}))
		ret.DrsAutoNatList = getSliceSlbServerServiceOperOperDrsAutoNatList(in["drs_auto_nat_list"].([]interface{}))
		ret.Pool_name = in["pool_name"].(string)
		ret.NatPoolAddrList = getSliceSlbServerServiceOperOperNatPoolAddrList(in["nat_pool_addr_list"].([]interface{}))
		ret.DrsIpNatList = getSliceSlbServerServiceOperOperDrsIpNatList(in["drs_ip_nat_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbServerServiceOperOperAutoNatAddrList(d []interface{}) []edpt.SlbServerServiceOperOperAutoNatAddrList {

	count1 := len(d)
	ret := make([]edpt.SlbServerServiceOperOperAutoNatAddrList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerServiceOperOperAutoNatAddrList
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

func getSliceSlbServerServiceOperOperDrsAutoNatList(d []interface{}) []edpt.SlbServerServiceOperOperDrsAutoNatList {

	count1 := len(d)
	ret := make([]edpt.SlbServerServiceOperOperDrsAutoNatList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerServiceOperOperDrsAutoNatList
		oi.Drs_name = in["drs_name"].(string)
		oi.Drs_port = in["drs_port"].(int)
		oi.DrsAutoNatAddressList = getSliceSlbServerServiceOperOperDrsAutoNatListDrsAutoNatAddressList(in["drs_auto_nat_address_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbServerServiceOperOperDrsAutoNatListDrsAutoNatAddressList(d []interface{}) []edpt.SlbServerServiceOperOperDrsAutoNatListDrsAutoNatAddressList {

	count1 := len(d)
	ret := make([]edpt.SlbServerServiceOperOperDrsAutoNatListDrsAutoNatAddressList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerServiceOperOperDrsAutoNatListDrsAutoNatAddressList
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

func getSliceSlbServerServiceOperOperNatPoolAddrList(d []interface{}) []edpt.SlbServerServiceOperOperNatPoolAddrList {

	count1 := len(d)
	ret := make([]edpt.SlbServerServiceOperOperNatPoolAddrList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerServiceOperOperNatPoolAddrList
		oi.Nat_ip = in["nat_ip"].(string)
		oi.Ports_consumed = in["ports_consumed"].(int)
		oi.Ports_consumed_total = in["ports_consumed_total"].(int)
		oi.Ports_freed_total = in["ports_freed_total"].(int)
		oi.Alloc_failed = in["alloc_failed"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbServerServiceOperOperDrsIpNatList(d []interface{}) []edpt.SlbServerServiceOperOperDrsIpNatList {

	count1 := len(d)
	ret := make([]edpt.SlbServerServiceOperOperDrsIpNatList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerServiceOperOperDrsIpNatList
		oi.Drs_name = in["drs_name"].(string)
		oi.Drs_port = in["drs_port"].(int)
		oi.Pool_name = in["pool_name"].(string)
		oi.NatPoolAddrList = getSliceSlbServerServiceOperOperDrsIpNatListNatPoolAddrList(in["nat_pool_addr_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbServerServiceOperOperDrsIpNatListNatPoolAddrList(d []interface{}) []edpt.SlbServerServiceOperOperDrsIpNatListNatPoolAddrList {

	count1 := len(d)
	ret := make([]edpt.SlbServerServiceOperOperDrsIpNatListNatPoolAddrList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerServiceOperOperDrsIpNatListNatPoolAddrList
		oi.Nat_ip = in["nat_ip"].(string)
		oi.Ports_consumed = in["ports_consumed"].(int)
		oi.Ports_consumed_total = in["ports_consumed_total"].(int)
		oi.Ports_freed_total = in["ports_freed_total"].(int)
		oi.Alloc_failed = in["alloc_failed"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbServerServiceOper(d *schema.ResourceData) edpt.SlbServerServiceOper {
	var ret edpt.SlbServerServiceOper

	ret.Label = d.Get("label").(string)

	ret.Oper = getObjectSlbServerServiceOperOper(d.Get("oper").([]interface{}))

	ret.PortNumber = d.Get("port_number").(int)

	ret.Protocol = d.Get("protocol").(string)

	ret.Server_name = d.Get("server_name").(string)
	return ret
}
