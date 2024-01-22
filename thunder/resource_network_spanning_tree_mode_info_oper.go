package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkSpanningTreeModeInfoOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_network_spanning_tree_mode_info_oper`: Operational Status for the object info\n\n__PLACEHOLDER__",
		ReadContext: resourceNetworkSpanningTreeModeInfoOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"packets_input": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"packets_output": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"instances": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"instance_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fwd_state_port": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"fwd_state_trunk": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"blk_state_port": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"blk_state_trunk": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"bridge_id": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"bridge_priority": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bridge_ext_priority": {
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
									"regional_root": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"regional_root_priority": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"regional_root_ext_priority": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"root_port": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"path_cost": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"int_path_cost": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"max_age": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"int_max_age": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"root_fwd_delay": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bridge_fwd_delay": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tx_hold_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"max_hops": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bridge_hello_time": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"age_time": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"time_since_topo_change": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"topo_change_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"topo_change_port": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"last_topo_change_port": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"vlans": {
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

func resourceNetworkSpanningTreeModeInfoOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkSpanningTreeModeInfoOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkSpanningTreeModeInfoOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NetworkSpanningTreeModeInfoOperOper := setObjectNetworkSpanningTreeModeInfoOperOper(res)
		d.Set("oper", NetworkSpanningTreeModeInfoOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNetworkSpanningTreeModeInfoOperOper(ret edpt.DataNetworkSpanningTreeModeInfoOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"mode":           ret.DtNetworkSpanningTreeModeInfoOper.Oper.Mode,
			"packets_input":  ret.DtNetworkSpanningTreeModeInfoOper.Oper.Packets_input,
			"packets_output": ret.DtNetworkSpanningTreeModeInfoOper.Oper.Packets_output,
			"instances":      setSliceNetworkSpanningTreeModeInfoOperOperInstances(ret.DtNetworkSpanningTreeModeInfoOper.Oper.Instances),
		},
	}
}

func setSliceNetworkSpanningTreeModeInfoOperOperInstances(d []edpt.NetworkSpanningTreeModeInfoOperOperInstances) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["instance_num"] = item.Instance_num
		in["fwd_state_port"] = item.Fwd_state_port
		in["fwd_state_trunk"] = item.Fwd_state_trunk
		in["blk_state_port"] = item.Blk_state_port
		in["blk_state_trunk"] = item.Blk_state_trunk
		in["bridge_id"] = item.Bridge_id
		in["bridge_priority"] = item.Bridge_priority
		in["bridge_ext_priority"] = item.Bridge_ext_priority
		in["des_root"] = item.Des_root
		in["des_root_priority"] = item.Des_root_priority
		in["des_root_ext_priority"] = item.Des_root_ext_priority
		in["regional_root"] = item.Regional_root
		in["regional_root_priority"] = item.Regional_root_priority
		in["regional_root_ext_priority"] = item.Regional_root_ext_priority
		in["root_port"] = item.Root_port
		in["path_cost"] = item.Path_cost
		in["int_path_cost"] = item.Int_path_cost
		in["max_age"] = item.Max_age
		in["int_max_age"] = item.Int_max_age
		in["root_fwd_delay"] = item.Root_fwd_delay
		in["bridge_fwd_delay"] = item.Bridge_fwd_delay
		in["tx_hold_count"] = item.Tx_hold_count
		in["max_hops"] = item.Max_hops
		in["bridge_hello_time"] = item.Bridge_hello_time
		in["age_time"] = item.Age_time
		in["time_since_topo_change"] = item.Time_since_topo_change
		in["topo_change_count"] = item.Topo_change_count
		in["topo_change_port"] = item.Topo_change_port
		in["last_topo_change_port"] = item.Last_topo_change_port
		in["vlans"] = item.Vlans
		result = append(result, in)
	}
	return result
}

func getObjectNetworkSpanningTreeModeInfoOperOper(d []interface{}) edpt.NetworkSpanningTreeModeInfoOperOper {

	count1 := len(d)
	var ret edpt.NetworkSpanningTreeModeInfoOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mode = in["mode"].(string)
		ret.Packets_input = in["packets_input"].(int)
		ret.Packets_output = in["packets_output"].(int)
		ret.Instances = getSliceNetworkSpanningTreeModeInfoOperOperInstances(in["instances"].([]interface{}))
	}
	return ret
}

func getSliceNetworkSpanningTreeModeInfoOperOperInstances(d []interface{}) []edpt.NetworkSpanningTreeModeInfoOperOperInstances {

	count1 := len(d)
	ret := make([]edpt.NetworkSpanningTreeModeInfoOperOperInstances, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkSpanningTreeModeInfoOperOperInstances
		oi.Instance_num = in["instance_num"].(int)
		oi.Fwd_state_port = in["fwd_state_port"].(string)
		oi.Fwd_state_trunk = in["fwd_state_trunk"].(string)
		oi.Blk_state_port = in["blk_state_port"].(string)
		oi.Blk_state_trunk = in["blk_state_trunk"].(string)
		oi.Bridge_id = in["bridge_id"].(string)
		oi.Bridge_priority = in["bridge_priority"].(int)
		oi.Bridge_ext_priority = in["bridge_ext_priority"].(int)
		oi.Des_root = in["des_root"].(string)
		oi.Des_root_priority = in["des_root_priority"].(int)
		oi.Des_root_ext_priority = in["des_root_ext_priority"].(int)
		oi.Regional_root = in["regional_root"].(string)
		oi.Regional_root_priority = in["regional_root_priority"].(int)
		oi.Regional_root_ext_priority = in["regional_root_ext_priority"].(int)
		oi.Root_port = in["root_port"].(string)
		oi.Path_cost = in["path_cost"].(int)
		oi.Int_path_cost = in["int_path_cost"].(int)
		oi.Max_age = in["max_age"].(int)
		oi.Int_max_age = in["int_max_age"].(int)
		oi.Root_fwd_delay = in["root_fwd_delay"].(int)
		oi.Bridge_fwd_delay = in["bridge_fwd_delay"].(int)
		oi.Tx_hold_count = in["tx_hold_count"].(int)
		oi.Max_hops = in["max_hops"].(int)
		oi.Bridge_hello_time = in["bridge_hello_time"].(int)
		oi.Age_time = in["age_time"].(int)
		oi.Time_since_topo_change = in["time_since_topo_change"].(int)
		oi.Topo_change_count = in["topo_change_count"].(int)
		oi.Topo_change_port = in["topo_change_port"].(string)
		oi.Last_topo_change_port = in["last_topo_change_port"].(string)
		oi.Vlans = in["vlans"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNetworkSpanningTreeModeInfoOper(d *schema.ResourceData) edpt.NetworkSpanningTreeModeInfoOper {
	var ret edpt.NetworkSpanningTreeModeInfoOper

	ret.Oper = getObjectNetworkSpanningTreeModeInfoOperOper(d.Get("oper").([]interface{}))
	return ret
}
