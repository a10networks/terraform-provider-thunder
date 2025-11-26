package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkSpanningTreeModeInterfaceOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_network_spanning_tree_mode_interface_oper`: Operational Status for the object interface\n\n__PLACEHOLDER__",
		ReadContext: resourceNetworkSpanningTreeModeInterfaceOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"instances": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"instance_num": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"role": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"port_id": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"state": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"in_port_cost": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"ad_in_port_cost": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"des_reg_root": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"des_reg_root_priority": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"des_reg_root_ext_priority": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"des_int_cost": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"des_bridge": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"des_bridge_priority": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"des_bridge_ext_priority": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"des_port": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"disputed": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"port_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"bridge": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"mode": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"enabled": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ex_port_cost": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ad_ex_port_cost": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"des_root": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"des_root_priority": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"des_root_ext_priority": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"des_ext_cost": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"adm_edge_port": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"auto_edge_port": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"oper_edge_port": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"tc_ack": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"oper_p2p": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"adm_p2p": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"rest_role": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"rest_tcn": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"port_hello_time": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bpdu_guard_port": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"bpdu_guard_err": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"net_port": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ba_incon": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"tx_bpdu": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tx_tcn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rx_bpdu": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rx_tcn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"trns_fwd": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"trns_blk": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcvd_bpdu": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"rcvd_stp": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"rcvd_rstp": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"snd_rstp": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"rcvd_ack": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"rcvd_tcn": {
										Type: schema.TypeString, Optional: true, Description: "",
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

func resourceNetworkSpanningTreeModeInterfaceOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkSpanningTreeModeInterfaceOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkSpanningTreeModeInterfaceOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NetworkSpanningTreeModeInterfaceOperOper := setObjectNetworkSpanningTreeModeInterfaceOperOper(res)
		d.Set("oper", NetworkSpanningTreeModeInterfaceOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNetworkSpanningTreeModeInterfaceOperOper(ret edpt.DataNetworkSpanningTreeModeInterfaceOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ports": setSliceNetworkSpanningTreeModeInterfaceOperOperPorts(ret.DtNetworkSpanningTreeModeInterfaceOper.Oper.Ports),
		},
	}
}

func setSliceNetworkSpanningTreeModeInterfaceOperOperPorts(d []edpt.NetworkSpanningTreeModeInterfaceOperOperPorts) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["instances"] = setSliceNetworkSpanningTreeModeInterfaceOperOperPortsInstances(item.Instances)
		in["port_name"] = item.Port_name
		in["bridge"] = item.Bridge
		in["mode"] = item.Mode
		in["enabled"] = item.Enabled
		in["ex_port_cost"] = item.Ex_port_cost
		in["ad_ex_port_cost"] = item.Ad_ex_port_cost
		in["des_root"] = item.Des_root
		in["des_root_priority"] = item.Des_root_priority
		in["des_root_ext_priority"] = item.Des_root_ext_priority
		in["des_ext_cost"] = item.Des_ext_cost
		in["adm_edge_port"] = item.Adm_edge_port
		in["auto_edge_port"] = item.Auto_edge_port
		in["oper_edge_port"] = item.Oper_edge_port
		in["tc_ack"] = item.Tc_ack
		in["oper_p2p"] = item.Oper_p2p
		in["adm_p2p"] = item.Adm_p2p
		in["rest_role"] = item.Rest_role
		in["rest_tcn"] = item.Rest_tcn
		in["port_hello_time"] = item.Port_hello_time
		in["bpdu_guard_port"] = item.Bpdu_guard_port
		in["bpdu_guard_err"] = item.Bpdu_guard_err
		in["net_port"] = item.Net_port
		in["ba_incon"] = item.Ba_incon
		in["tx_bpdu"] = item.Tx_bpdu
		in["tx_tcn"] = item.Tx_tcn
		in["rx_bpdu"] = item.Rx_bpdu
		in["rx_tcn"] = item.Rx_tcn
		in["trns_fwd"] = item.Trns_fwd
		in["trns_blk"] = item.Trns_blk
		in["rcvd_bpdu"] = item.Rcvd_bpdu
		in["rcvd_stp"] = item.Rcvd_stp
		in["rcvd_rstp"] = item.Rcvd_rstp
		in["snd_rstp"] = item.Snd_rstp
		in["rcvd_ack"] = item.Rcvd_ack
		in["rcvd_tcn"] = item.Rcvd_tcn
		result = append(result, in)
	}
	return result
}

func setSliceNetworkSpanningTreeModeInterfaceOperOperPortsInstances(d []edpt.NetworkSpanningTreeModeInterfaceOperOperPortsInstances) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["instance_num"] = item.Instance_num
		in["role"] = item.Role
		in["port_id"] = item.Port_id
		in["state"] = item.State
		in["in_port_cost"] = item.In_port_cost
		in["ad_in_port_cost"] = item.Ad_in_port_cost
		in["des_reg_root"] = item.Des_reg_root
		in["des_reg_root_priority"] = item.Des_reg_root_priority
		in["des_reg_root_ext_priority"] = item.Des_reg_root_ext_priority
		in["des_int_cost"] = item.Des_int_cost
		in["des_bridge"] = item.Des_bridge
		in["des_bridge_priority"] = item.Des_bridge_priority
		in["des_bridge_ext_priority"] = item.Des_bridge_ext_priority
		in["des_port"] = item.Des_port
		in["disputed"] = item.Disputed
		result = append(result, in)
	}
	return result
}

func getObjectNetworkSpanningTreeModeInterfaceOperOper(d []interface{}) edpt.NetworkSpanningTreeModeInterfaceOperOper {

	count1 := len(d)
	var ret edpt.NetworkSpanningTreeModeInterfaceOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ports = getSliceNetworkSpanningTreeModeInterfaceOperOperPorts(in["ports"].([]interface{}))
	}
	return ret
}

func getSliceNetworkSpanningTreeModeInterfaceOperOperPorts(d []interface{}) []edpt.NetworkSpanningTreeModeInterfaceOperOperPorts {

	count1 := len(d)
	ret := make([]edpt.NetworkSpanningTreeModeInterfaceOperOperPorts, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkSpanningTreeModeInterfaceOperOperPorts
		oi.Instances = getSliceNetworkSpanningTreeModeInterfaceOperOperPortsInstances(in["instances"].([]interface{}))
		oi.Port_name = in["port_name"].(string)
		oi.Bridge = in["bridge"].(string)
		oi.Mode = in["mode"].(int)
		oi.Enabled = in["enabled"].(string)
		oi.Ex_port_cost = in["ex_port_cost"].(int)
		oi.Ad_ex_port_cost = in["ad_ex_port_cost"].(int)
		oi.Des_root = in["des_root"].(string)
		oi.Des_root_priority = in["des_root_priority"].(int)
		oi.Des_root_ext_priority = in["des_root_ext_priority"].(int)
		oi.Des_ext_cost = in["des_ext_cost"].(int)
		oi.Adm_edge_port = in["adm_edge_port"].(string)
		oi.Auto_edge_port = in["auto_edge_port"].(string)
		oi.Oper_edge_port = in["oper_edge_port"].(string)
		oi.Tc_ack = in["tc_ack"].(string)
		oi.Oper_p2p = in["oper_p2p"].(string)
		oi.Adm_p2p = in["adm_p2p"].(string)
		oi.Rest_role = in["rest_role"].(string)
		oi.Rest_tcn = in["rest_tcn"].(string)
		oi.Port_hello_time = in["port_hello_time"].(int)
		oi.Bpdu_guard_port = in["bpdu_guard_port"].(string)
		oi.Bpdu_guard_err = in["bpdu_guard_err"].(string)
		oi.Net_port = in["net_port"].(string)
		oi.Ba_incon = in["ba_incon"].(string)
		oi.Tx_bpdu = in["tx_bpdu"].(int)
		oi.Tx_tcn = in["tx_tcn"].(int)
		oi.Rx_bpdu = in["rx_bpdu"].(int)
		oi.Rx_tcn = in["rx_tcn"].(int)
		oi.Trns_fwd = in["trns_fwd"].(int)
		oi.Trns_blk = in["trns_blk"].(int)
		oi.Rcvd_bpdu = in["rcvd_bpdu"].(string)
		oi.Rcvd_stp = in["rcvd_stp"].(string)
		oi.Rcvd_rstp = in["rcvd_rstp"].(string)
		oi.Snd_rstp = in["snd_rstp"].(string)
		oi.Rcvd_ack = in["rcvd_ack"].(string)
		oi.Rcvd_tcn = in["rcvd_tcn"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetworkSpanningTreeModeInterfaceOperOperPortsInstances(d []interface{}) []edpt.NetworkSpanningTreeModeInterfaceOperOperPortsInstances {

	count1 := len(d)
	ret := make([]edpt.NetworkSpanningTreeModeInterfaceOperOperPortsInstances, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkSpanningTreeModeInterfaceOperOperPortsInstances
		oi.Instance_num = in["instance_num"].(int)
		oi.Role = in["role"].(string)
		oi.Port_id = in["port_id"].(string)
		oi.State = in["state"].(string)
		oi.In_port_cost = in["in_port_cost"].(int)
		oi.Ad_in_port_cost = in["ad_in_port_cost"].(int)
		oi.Des_reg_root = in["des_reg_root"].(string)
		oi.Des_reg_root_priority = in["des_reg_root_priority"].(int)
		oi.Des_reg_root_ext_priority = in["des_reg_root_ext_priority"].(int)
		oi.Des_int_cost = in["des_int_cost"].(int)
		oi.Des_bridge = in["des_bridge"].(string)
		oi.Des_bridge_priority = in["des_bridge_priority"].(int)
		oi.Des_bridge_ext_priority = in["des_bridge_ext_priority"].(int)
		oi.Des_port = in["des_port"].(string)
		oi.Disputed = in["disputed"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNetworkSpanningTreeModeInterfaceOper(d *schema.ResourceData) edpt.NetworkSpanningTreeModeInterfaceOper {
	var ret edpt.NetworkSpanningTreeModeInterfaceOper

	ret.Oper = getObjectNetworkSpanningTreeModeInterfaceOperOper(d.Get("oper").([]interface{}))
	return ret
}
