package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbServerOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_server_oper`: Operational Status for the object server\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbServerOperRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Server Name",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"creation_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dns_update_time": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"server_ttl": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"srv_gateway_arp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"is_autocreate": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"slow_start_conn_limit": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"curr_conn_rate": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"conn_rate_unit": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"curr_observe_rate": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"disable": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"weight": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"drs_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"drs_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"drs_host": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"drs_server_ipv6_addr": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"drs_state": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"drs_creation_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"drs_dns_update_time": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"drs_server_ttl": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"drs_srv_gateway_arp": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"drs_is_autocreate": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"drs_slow_start_conn_limit": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"drs_curr_conn_rate": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"drs_conn_rate_unit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"drs_curr_observe_rate": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"drs_disable": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"drs_curr_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"drs_curr_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"drs_tot_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"drs_tot_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"drs_tot_req_suc": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"drs_tot_fwd_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"drs_tot_fwd_pkts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"drs_tot_rev_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"drs_tot_rev_pkts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"drs_peak_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"drs_weight": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"port_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_number": {
							Type: schema.TypeInt, Required: true, Description: "Port Number",
						},
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'tcp': TCP Port; 'udp': UDP Port;",
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
					},
				},
			},
		},
	}
}

func resourceSlbServerOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServerOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServerOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbServerOperOper := setObjectSlbServerOperOper(res)
		d.Set("oper", SlbServerOperOper)
		SlbServerOperPortList := setSliceSlbServerOperPortList(res)
		d.Set("port_list", SlbServerOperPortList)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbServerOperOper(ret edpt.DataSlbServerOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"state":                 ret.DtSlbServerOper.Oper.State,
			"creation_type":         ret.DtSlbServerOper.Oper.Creation_type,
			"dns_update_time":       ret.DtSlbServerOper.Oper.Dns_update_time,
			"server_ttl":            ret.DtSlbServerOper.Oper.Server_ttl,
			"srv_gateway_arp":       ret.DtSlbServerOper.Oper.Srv_gateway_arp,
			"is_autocreate":         ret.DtSlbServerOper.Oper.Is_autocreate,
			"slow_start_conn_limit": ret.DtSlbServerOper.Oper.Slow_start_conn_limit,
			"curr_conn_rate":        ret.DtSlbServerOper.Oper.Curr_conn_rate,
			"conn_rate_unit":        ret.DtSlbServerOper.Oper.Conn_rate_unit,
			"curr_observe_rate":     ret.DtSlbServerOper.Oper.Curr_observe_rate,
			"disable":               ret.DtSlbServerOper.Oper.Disable,
			"weight":                ret.DtSlbServerOper.Oper.Weight,
			"drs_list":              setSliceSlbServerOperOperDrsList(ret.DtSlbServerOper.Oper.DrsList),
		},
	}
}

func setSliceSlbServerOperOperDrsList(d []edpt.SlbServerOperOperDrsList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["drs_name"] = item.DrsName
		in["drs_host"] = item.DrsHost
		in["drs_server_ipv6_addr"] = item.DrsServerIpv6Addr
		in["drs_state"] = item.DrsState
		in["drs_creation_type"] = item.DrsCreation_type
		in["drs_dns_update_time"] = item.DrsDns_update_time
		in["drs_server_ttl"] = item.DrsServer_ttl
		in["drs_srv_gateway_arp"] = item.DrsSrv_gateway_arp
		in["drs_is_autocreate"] = item.DrsIs_autocreate
		in["drs_slow_start_conn_limit"] = item.DrsSlow_start_conn_limit
		in["drs_curr_conn_rate"] = item.DrsCurr_conn_rate
		in["drs_conn_rate_unit"] = item.DrsConn_rate_unit
		in["drs_curr_observe_rate"] = item.DrsCurr_observe_rate
		in["drs_disable"] = item.DrsDisable
		in["drs_curr_conn"] = item.DrsCurrConn
		in["drs_curr_req"] = item.DrsCurrReq
		in["drs_tot_conn"] = item.DrsTotConn
		in["drs_tot_req"] = item.DrsTotReq
		in["drs_tot_req_suc"] = item.DrsTotReqSuc
		in["drs_tot_fwd_bytes"] = item.DrsTotFwdBytes
		in["drs_tot_fwd_pkts"] = item.DrsTotFwdPkts
		in["drs_tot_rev_bytes"] = item.DrsTotRevBytes
		in["drs_tot_rev_pkts"] = item.DrsTotRevPkts
		in["drs_peak_conn"] = item.DrsPeakConn
		in["drs_weight"] = item.DrsWeight
		result = append(result, in)
	}
	return result
}

func setSliceSlbServerOperPortList(d edpt.DataSlbServerOper) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtSlbServerOper.PortList {
		in := make(map[string]interface{})
		in["port_number"] = item.PortNumber
		in["protocol"] = item.Protocol
		in["oper"] = setObjectSlbServerOperPortListOper(item.Oper)
		result = append(result, in)
	}
	return result
}

func setObjectSlbServerOperPortListOper(d edpt.SlbServerOperPortListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["state"] = d.State

	in["curr_conn_rate"] = d.Curr_conn_rate

	in["conn_rate_unit"] = d.Conn_rate_unit

	in["slow_start_conn_limit"] = d.Slow_start_conn_limit

	in["curr_observe_rate"] = d.Curr_observe_rate

	in["down_grace_period_allowed"] = d.Down_grace_period_allowed

	in["current_time"] = d.Current_time

	in["down_time_grace_period"] = d.Down_time_grace_period

	in["diameter_enabled"] = d.Diameter_enabled

	in["es_resp_time"] = d.Es_resp_time

	in["inband_hm_reassign_num"] = d.Inband_hm_reassign_num

	in["disable"] = d.Disable

	in["hm_key"] = d.HmKey

	in["hm_index"] = d.HmIndex

	in["soft_down_time"] = d.Soft_down_time

	in["aflow_conn_limit"] = d.Aflow_conn_limit

	in["aflow_queue_size"] = d.Aflow_queue_size

	in["resv_conn"] = d.Resv_conn
	in["auto_nat_addr_list"] = setSliceSlbServerOperPortListOperAutoNatAddrList(d.AutoNatAddrList)
	in["drs_auto_nat_list"] = setSliceSlbServerOperPortListOperDrsAutoNatList(d.DrsAutoNatList)

	in["pool_name"] = d.Pool_name
	in["nat_pool_addr_list"] = setSliceSlbServerOperPortListOperNatPoolAddrList(d.NatPoolAddrList)
	in["drs_ip_nat_list"] = setSliceSlbServerOperPortListOperDrsIpNatList(d.DrsIpNatList)
	result = append(result, in)
	return result
}

func setSliceSlbServerOperPortListOperAutoNatAddrList(d []edpt.SlbServerOperPortListOperAutoNatAddrList) []map[string]interface{} {
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

func setSliceSlbServerOperPortListOperDrsAutoNatList(d []edpt.SlbServerOperPortListOperDrsAutoNatList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["drs_name"] = item.Drs_name
		in["drs_port"] = item.Drs_port
		in["drs_auto_nat_address_list"] = setSliceSlbServerOperPortListOperDrsAutoNatListDrsAutoNatAddressList(item.DrsAutoNatAddressList)
		result = append(result, in)
	}
	return result
}

func setSliceSlbServerOperPortListOperDrsAutoNatListDrsAutoNatAddressList(d []edpt.SlbServerOperPortListOperDrsAutoNatListDrsAutoNatAddressList) []map[string]interface{} {
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

func setSliceSlbServerOperPortListOperNatPoolAddrList(d []edpt.SlbServerOperPortListOperNatPoolAddrList) []map[string]interface{} {
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

func setSliceSlbServerOperPortListOperDrsIpNatList(d []edpt.SlbServerOperPortListOperDrsIpNatList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["drs_name"] = item.Drs_name
		in["drs_port"] = item.Drs_port
		in["pool_name"] = item.Pool_name
		in["nat_pool_addr_list"] = setSliceSlbServerOperPortListOperDrsIpNatListNatPoolAddrList(item.NatPoolAddrList)
		result = append(result, in)
	}
	return result
}

func setSliceSlbServerOperPortListOperDrsIpNatListNatPoolAddrList(d []edpt.SlbServerOperPortListOperDrsIpNatListNatPoolAddrList) []map[string]interface{} {
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

func getObjectSlbServerOperOper(d []interface{}) edpt.SlbServerOperOper {

	count1 := len(d)
	var ret edpt.SlbServerOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.Creation_type = in["creation_type"].(string)
		ret.Dns_update_time = in["dns_update_time"].(string)
		ret.Server_ttl = in["server_ttl"].(int)
		ret.Srv_gateway_arp = in["srv_gateway_arp"].(string)
		ret.Is_autocreate = in["is_autocreate"].(int)
		ret.Slow_start_conn_limit = in["slow_start_conn_limit"].(int)
		ret.Curr_conn_rate = in["curr_conn_rate"].(int)
		ret.Conn_rate_unit = in["conn_rate_unit"].(string)
		ret.Curr_observe_rate = in["curr_observe_rate"].(int)
		ret.Disable = in["disable"].(int)
		ret.Weight = in["weight"].(int)
		ret.DrsList = getSliceSlbServerOperOperDrsList(in["drs_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbServerOperOperDrsList(d []interface{}) []edpt.SlbServerOperOperDrsList {

	count1 := len(d)
	ret := make([]edpt.SlbServerOperOperDrsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerOperOperDrsList
		oi.DrsName = in["drs_name"].(string)
		oi.DrsHost = in["drs_host"].(string)
		oi.DrsServerIpv6Addr = in["drs_server_ipv6_addr"].(string)
		oi.DrsState = in["drs_state"].(string)
		oi.DrsCreation_type = in["drs_creation_type"].(string)
		oi.DrsDns_update_time = in["drs_dns_update_time"].(string)
		oi.DrsServer_ttl = in["drs_server_ttl"].(int)
		oi.DrsSrv_gateway_arp = in["drs_srv_gateway_arp"].(string)
		oi.DrsIs_autocreate = in["drs_is_autocreate"].(int)
		oi.DrsSlow_start_conn_limit = in["drs_slow_start_conn_limit"].(int)
		oi.DrsCurr_conn_rate = in["drs_curr_conn_rate"].(int)
		oi.DrsConn_rate_unit = in["drs_conn_rate_unit"].(string)
		oi.DrsCurr_observe_rate = in["drs_curr_observe_rate"].(int)
		oi.DrsDisable = in["drs_disable"].(int)
		oi.DrsCurrConn = in["drs_curr_conn"].(int)
		oi.DrsCurrReq = in["drs_curr_req"].(int)
		oi.DrsTotConn = in["drs_tot_conn"].(int)
		oi.DrsTotReq = in["drs_tot_req"].(int)
		oi.DrsTotReqSuc = in["drs_tot_req_suc"].(int)
		oi.DrsTotFwdBytes = in["drs_tot_fwd_bytes"].(int)
		oi.DrsTotFwdPkts = in["drs_tot_fwd_pkts"].(int)
		oi.DrsTotRevBytes = in["drs_tot_rev_bytes"].(int)
		oi.DrsTotRevPkts = in["drs_tot_rev_pkts"].(int)
		oi.DrsPeakConn = in["drs_peak_conn"].(int)
		oi.DrsWeight = in["drs_weight"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbServerOperPortList(d []interface{}) []edpt.SlbServerOperPortList {

	count1 := len(d)
	ret := make([]edpt.SlbServerOperPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerOperPortList
		oi.PortNumber = in["port_number"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Oper = getObjectSlbServerOperPortListOper(in["oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbServerOperPortListOper(d []interface{}) edpt.SlbServerOperPortListOper {

	count1 := len(d)
	var ret edpt.SlbServerOperPortListOper
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
		ret.AutoNatAddrList = getSliceSlbServerOperPortListOperAutoNatAddrList(in["auto_nat_addr_list"].([]interface{}))
		ret.DrsAutoNatList = getSliceSlbServerOperPortListOperDrsAutoNatList(in["drs_auto_nat_list"].([]interface{}))
		ret.Pool_name = in["pool_name"].(string)
		ret.NatPoolAddrList = getSliceSlbServerOperPortListOperNatPoolAddrList(in["nat_pool_addr_list"].([]interface{}))
		ret.DrsIpNatList = getSliceSlbServerOperPortListOperDrsIpNatList(in["drs_ip_nat_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbServerOperPortListOperAutoNatAddrList(d []interface{}) []edpt.SlbServerOperPortListOperAutoNatAddrList {

	count1 := len(d)
	ret := make([]edpt.SlbServerOperPortListOperAutoNatAddrList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerOperPortListOperAutoNatAddrList
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

func getSliceSlbServerOperPortListOperDrsAutoNatList(d []interface{}) []edpt.SlbServerOperPortListOperDrsAutoNatList {

	count1 := len(d)
	ret := make([]edpt.SlbServerOperPortListOperDrsAutoNatList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerOperPortListOperDrsAutoNatList
		oi.Drs_name = in["drs_name"].(string)
		oi.Drs_port = in["drs_port"].(int)
		oi.DrsAutoNatAddressList = getSliceSlbServerOperPortListOperDrsAutoNatListDrsAutoNatAddressList(in["drs_auto_nat_address_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbServerOperPortListOperDrsAutoNatListDrsAutoNatAddressList(d []interface{}) []edpt.SlbServerOperPortListOperDrsAutoNatListDrsAutoNatAddressList {

	count1 := len(d)
	ret := make([]edpt.SlbServerOperPortListOperDrsAutoNatListDrsAutoNatAddressList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerOperPortListOperDrsAutoNatListDrsAutoNatAddressList
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

func getSliceSlbServerOperPortListOperNatPoolAddrList(d []interface{}) []edpt.SlbServerOperPortListOperNatPoolAddrList {

	count1 := len(d)
	ret := make([]edpt.SlbServerOperPortListOperNatPoolAddrList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerOperPortListOperNatPoolAddrList
		oi.Nat_ip = in["nat_ip"].(string)
		oi.Ports_consumed = in["ports_consumed"].(int)
		oi.Ports_consumed_total = in["ports_consumed_total"].(int)
		oi.Ports_freed_total = in["ports_freed_total"].(int)
		oi.Alloc_failed = in["alloc_failed"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbServerOperPortListOperDrsIpNatList(d []interface{}) []edpt.SlbServerOperPortListOperDrsIpNatList {

	count1 := len(d)
	ret := make([]edpt.SlbServerOperPortListOperDrsIpNatList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerOperPortListOperDrsIpNatList
		oi.Drs_name = in["drs_name"].(string)
		oi.Drs_port = in["drs_port"].(int)
		oi.Pool_name = in["pool_name"].(string)
		oi.NatPoolAddrList = getSliceSlbServerOperPortListOperDrsIpNatListNatPoolAddrList(in["nat_pool_addr_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbServerOperPortListOperDrsIpNatListNatPoolAddrList(d []interface{}) []edpt.SlbServerOperPortListOperDrsIpNatListNatPoolAddrList {

	count1 := len(d)
	ret := make([]edpt.SlbServerOperPortListOperDrsIpNatListNatPoolAddrList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerOperPortListOperDrsIpNatListNatPoolAddrList
		oi.Nat_ip = in["nat_ip"].(string)
		oi.Ports_consumed = in["ports_consumed"].(int)
		oi.Ports_consumed_total = in["ports_consumed_total"].(int)
		oi.Ports_freed_total = in["ports_freed_total"].(int)
		oi.Alloc_failed = in["alloc_failed"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbServerOper(d *schema.ResourceData) edpt.SlbServerOper {
	var ret edpt.SlbServerOper

	ret.Name = d.Get("name").(string)

	ret.Oper = getObjectSlbServerOperOper(d.Get("oper").([]interface{}))

	ret.PortList = getSliceSlbServerOperPortList(d.Get("port_list").([]interface{}))
	return ret
}
