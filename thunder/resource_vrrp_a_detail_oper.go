package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpADetailOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vrrp_a_detail_oper`: Operational Status for the object detail\n\n__PLACEHOLDER__",
		ReadContext: resourceVrrpADetailOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dup_id_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"bad_group_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"set_id_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"vrrp_version_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"time_inaccurate_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ip_pools_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ip6_pools_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"err_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"err_devid": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"err_parid": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"lock_try": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"l2_no_route": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"peer_info_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"peer_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"peer_pkt_recv": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"peer_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"peer_id": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"vrid": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"missing_heartbeat": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"peer_port_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"eth": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"vrrp_pkt_recv": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"eth_miss": {
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
						"local_info_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"switch_to_active": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"switch_to_standby": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"vrid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"vrrp_pkt_send": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"local_eth_send_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"eth_port": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"eth_pkt_send": {
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
	}
}

func resourceVrrpADetailOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpADetailOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpADetailOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VrrpADetailOperOper := setObjectVrrpADetailOperOper(res)
		d.Set("oper", VrrpADetailOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVrrpADetailOperOper(ret edpt.DataVrrpADetailOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"dup_id_rcv":            ret.DtVrrpADetailOper.Oper.Dup_id_rcv,
			"bad_group_rcv":         ret.DtVrrpADetailOper.Oper.Bad_group_rcv,
			"set_id_mismatch":       ret.DtVrrpADetailOper.Oper.Set_id_mismatch,
			"vrrp_version_mismatch": ret.DtVrrpADetailOper.Oper.Vrrp_version_mismatch,
			"time_inaccurate_count": ret.DtVrrpADetailOper.Oper.Time_inaccurate_count,
			"ip_pools_exceeded":     ret.DtVrrpADetailOper.Oper.Ip_pools_exceeded,
			"ip6_pools_exceeded":    ret.DtVrrpADetailOper.Oper.Ip6_pools_exceeded,
			"err_port":              ret.DtVrrpADetailOper.Oper.Err_port,
			"err_devid":             ret.DtVrrpADetailOper.Oper.Err_devid,
			"err_parid":             ret.DtVrrpADetailOper.Oper.Err_parid,
			"lock_try":              ret.DtVrrpADetailOper.Oper.Lock_try,
			"l2_no_route":           ret.DtVrrpADetailOper.Oper.L2_no_route,
			"peer_info_list":        setSliceVrrpADetailOperOperPeer_info_list(ret.DtVrrpADetailOper.Oper.Peer_info_list),
			"local_info_list":       setSliceVrrpADetailOperOperLocal_info_list(ret.DtVrrpADetailOper.Oper.Local_info_list),
		},
	}
}

func setSliceVrrpADetailOperOperPeer_info_list(d []edpt.VrrpADetailOperOperPeer_info_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["peer_ip"] = item.Peer_ip
		in["peer_pkt_recv"] = item.Peer_pkt_recv
		in["peer_list"] = setSliceVrrpADetailOperOperPeer_info_listPeer_list(item.Peer_list)
		result = append(result, in)
	}
	return result
}

func setSliceVrrpADetailOperOperPeer_info_listPeer_list(d []edpt.VrrpADetailOperOperPeer_info_listPeer_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["peer_id"] = item.Peer_id
		in["vrid"] = item.Vrid
		in["missing_heartbeat"] = item.Missing_heartbeat
		in["peer_port_list"] = setSliceVrrpADetailOperOperPeer_info_listPeer_listPeer_port_list(item.Peer_port_list)
		result = append(result, in)
	}
	return result
}

func setSliceVrrpADetailOperOperPeer_info_listPeer_listPeer_port_list(d []edpt.VrrpADetailOperOperPeer_info_listPeer_listPeer_port_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["eth"] = item.Eth
		in["vrrp_pkt_recv"] = item.Vrrp_pkt_recv
		in["eth_miss"] = item.Eth_miss
		result = append(result, in)
	}
	return result
}

func setSliceVrrpADetailOperOperLocal_info_list(d []edpt.VrrpADetailOperOperLocal_info_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["switch_to_active"] = item.Switch_to_active
		in["switch_to_standby"] = item.Switch_to_standby
		in["vrid"] = item.Vrid
		in["vrrp_pkt_send"] = item.Vrrp_pkt_send
		in["local_eth_send_list"] = setSliceVrrpADetailOperOperLocal_info_listLocal_eth_send_list(item.Local_eth_send_list)
		result = append(result, in)
	}
	return result
}

func setSliceVrrpADetailOperOperLocal_info_listLocal_eth_send_list(d []edpt.VrrpADetailOperOperLocal_info_listLocal_eth_send_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["eth_port"] = item.Eth_port
		in["eth_pkt_send"] = item.Eth_pkt_send
		result = append(result, in)
	}
	return result
}

func getObjectVrrpADetailOperOper(d []interface{}) edpt.VrrpADetailOperOper {

	count1 := len(d)
	var ret edpt.VrrpADetailOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dup_id_rcv = in["dup_id_rcv"].(int)
		ret.Bad_group_rcv = in["bad_group_rcv"].(int)
		ret.Set_id_mismatch = in["set_id_mismatch"].(int)
		ret.Vrrp_version_mismatch = in["vrrp_version_mismatch"].(int)
		ret.Time_inaccurate_count = in["time_inaccurate_count"].(int)
		ret.Ip_pools_exceeded = in["ip_pools_exceeded"].(int)
		ret.Ip6_pools_exceeded = in["ip6_pools_exceeded"].(int)
		ret.Err_port = in["err_port"].(int)
		ret.Err_devid = in["err_devid"].(int)
		ret.Err_parid = in["err_parid"].(int)
		ret.Lock_try = in["lock_try"].(int)
		ret.L2_no_route = in["l2_no_route"].(int)
		ret.Peer_info_list = getSliceVrrpADetailOperOperPeer_info_list(in["peer_info_list"].([]interface{}))
		ret.Local_info_list = getSliceVrrpADetailOperOperLocal_info_list(in["local_info_list"].([]interface{}))
	}
	return ret
}

func getSliceVrrpADetailOperOperPeer_info_list(d []interface{}) []edpt.VrrpADetailOperOperPeer_info_list {

	count1 := len(d)
	ret := make([]edpt.VrrpADetailOperOperPeer_info_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpADetailOperOperPeer_info_list
		oi.Peer_ip = in["peer_ip"].(string)
		oi.Peer_pkt_recv = in["peer_pkt_recv"].(int)
		oi.Peer_list = getSliceVrrpADetailOperOperPeer_info_listPeer_list(in["peer_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpADetailOperOperPeer_info_listPeer_list(d []interface{}) []edpt.VrrpADetailOperOperPeer_info_listPeer_list {

	count1 := len(d)
	ret := make([]edpt.VrrpADetailOperOperPeer_info_listPeer_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpADetailOperOperPeer_info_listPeer_list
		oi.Peer_id = in["peer_id"].(int)
		oi.Vrid = in["vrid"].(int)
		oi.Missing_heartbeat = in["missing_heartbeat"].(int)
		oi.Peer_port_list = getSliceVrrpADetailOperOperPeer_info_listPeer_listPeer_port_list(in["peer_port_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpADetailOperOperPeer_info_listPeer_listPeer_port_list(d []interface{}) []edpt.VrrpADetailOperOperPeer_info_listPeer_listPeer_port_list {

	count1 := len(d)
	ret := make([]edpt.VrrpADetailOperOperPeer_info_listPeer_listPeer_port_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpADetailOperOperPeer_info_listPeer_listPeer_port_list
		oi.Eth = in["eth"].(int)
		oi.Vrrp_pkt_recv = in["vrrp_pkt_recv"].(int)
		oi.Eth_miss = in["eth_miss"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpADetailOperOperLocal_info_list(d []interface{}) []edpt.VrrpADetailOperOperLocal_info_list {

	count1 := len(d)
	ret := make([]edpt.VrrpADetailOperOperLocal_info_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpADetailOperOperLocal_info_list
		oi.Switch_to_active = in["switch_to_active"].(int)
		oi.Switch_to_standby = in["switch_to_standby"].(int)
		oi.Vrid = in["vrid"].(int)
		oi.Vrrp_pkt_send = in["vrrp_pkt_send"].(int)
		oi.Local_eth_send_list = getSliceVrrpADetailOperOperLocal_info_listLocal_eth_send_list(in["local_eth_send_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpADetailOperOperLocal_info_listLocal_eth_send_list(d []interface{}) []edpt.VrrpADetailOperOperLocal_info_listLocal_eth_send_list {

	count1 := len(d)
	ret := make([]edpt.VrrpADetailOperOperLocal_info_listLocal_eth_send_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpADetailOperOperLocal_info_listLocal_eth_send_list
		oi.Eth_port = in["eth_port"].(int)
		oi.Eth_pkt_send = in["eth_pkt_send"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVrrpADetailOper(d *schema.ResourceData) edpt.VrrpADetailOper {
	var ret edpt.VrrpADetailOper

	ret.Oper = getObjectVrrpADetailOperOper(d.Get("oper").([]interface{}))
	return ret
}
