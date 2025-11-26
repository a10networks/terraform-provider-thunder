package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthSourceNatStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_health_source_nat_stats`: Statistics for the object source-nat\n\n__PLACEHOLDER__",
		ReadContext: resourceHealthSourceNatStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"act_recv_from_sby": {
							Type: schema.TypeInt, Optional: true, Description: "Packets received from standby",
						},
						"act_send_to_sby": {
							Type: schema.TypeInt, Optional: true, Description: "Packets sent to standby",
						},
						"sby_recv_from_act": {
							Type: schema.TypeInt, Optional: true, Description: "Packets received from active",
						},
						"sby_send_to_act": {
							Type: schema.TypeInt, Optional: true, Description: "Packets sent to active",
						},
						"sby_recv_from_act_err": {
							Type: schema.TypeInt, Optional: true, Description: "Packets received from active error",
						},
						"recv_from_kernel": {
							Type: schema.TypeInt, Optional: true, Description: "Packets received from kernel",
						},
						"send_to_kernel": {
							Type: schema.TypeInt, Optional: true, Description: "Packets sent to kernel",
						},
						"send_to_kernel_err": {
							Type: schema.TypeInt, Optional: true, Description: "Packets sent to kernel error",
						},
						"sby_no_peer": {
							Type: schema.TypeInt, Optional: true, Description: "Peer not found on standby",
						},
						"dcmsg_err": {
							Type: schema.TypeInt, Optional: true, Description: "DCMSG error",
						},
						"no_slb_object": {
							Type: schema.TypeInt, Optional: true, Description: "SLB object not found",
						},
						"smart_nat_init_port_err": {
							Type: schema.TypeInt, Optional: true, Description: "Smart NAT port initialization error",
						},
						"smart_nat_init_inst_err": {
							Type: schema.TypeInt, Optional: true, Description: "Smart NAT instance initialization error",
						},
						"smart_nat_rserver_route_err": {
							Type: schema.TypeInt, Optional: true, Description: "Smart NAT rserver route update error",
						},
						"smart_nat_rserver_ip_err": {
							Type: schema.TypeInt, Optional: true, Description: "Smart NAT rserver ip update error",
						},
						"nat_resource_err": {
							Type: schema.TypeInt, Optional: true, Description: "NAT resource error",
						},
						"frag_err": {
							Type: schema.TypeInt, Optional: true, Description: "Fragmentation error",
						},
					},
				},
			},
		},
	}
}

func resourceHealthSourceNatStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthSourceNatStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthSourceNatStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		HealthSourceNatStatsStats := setObjectHealthSourceNatStatsStats(res)
		d.Set("stats", HealthSourceNatStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectHealthSourceNatStatsStats(ret edpt.DataHealthSourceNatStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"act_recv_from_sby":           ret.DtHealthSourceNatStats.Stats.Act_recv_from_sby,
			"act_send_to_sby":             ret.DtHealthSourceNatStats.Stats.Act_send_to_sby,
			"sby_recv_from_act":           ret.DtHealthSourceNatStats.Stats.Sby_recv_from_act,
			"sby_send_to_act":             ret.DtHealthSourceNatStats.Stats.Sby_send_to_act,
			"sby_recv_from_act_err":       ret.DtHealthSourceNatStats.Stats.Sby_recv_from_act_err,
			"recv_from_kernel":            ret.DtHealthSourceNatStats.Stats.Recv_from_kernel,
			"send_to_kernel":              ret.DtHealthSourceNatStats.Stats.Send_to_kernel,
			"send_to_kernel_err":          ret.DtHealthSourceNatStats.Stats.Send_to_kernel_err,
			"sby_no_peer":                 ret.DtHealthSourceNatStats.Stats.Sby_no_peer,
			"dcmsg_err":                   ret.DtHealthSourceNatStats.Stats.Dcmsg_err,
			"no_slb_object":               ret.DtHealthSourceNatStats.Stats.No_slb_object,
			"smart_nat_init_port_err":     ret.DtHealthSourceNatStats.Stats.Smart_nat_init_port_err,
			"smart_nat_init_inst_err":     ret.DtHealthSourceNatStats.Stats.Smart_nat_init_inst_err,
			"smart_nat_rserver_route_err": ret.DtHealthSourceNatStats.Stats.Smart_nat_rserver_route_err,
			"smart_nat_rserver_ip_err":    ret.DtHealthSourceNatStats.Stats.Smart_nat_rserver_ip_err,
			"nat_resource_err":            ret.DtHealthSourceNatStats.Stats.Nat_resource_err,
			"frag_err":                    ret.DtHealthSourceNatStats.Stats.Frag_err,
		},
	}
}

func getObjectHealthSourceNatStatsStats(d []interface{}) edpt.HealthSourceNatStatsStats {

	count1 := len(d)
	var ret edpt.HealthSourceNatStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Act_recv_from_sby = in["act_recv_from_sby"].(int)
		ret.Act_send_to_sby = in["act_send_to_sby"].(int)
		ret.Sby_recv_from_act = in["sby_recv_from_act"].(int)
		ret.Sby_send_to_act = in["sby_send_to_act"].(int)
		ret.Sby_recv_from_act_err = in["sby_recv_from_act_err"].(int)
		ret.Recv_from_kernel = in["recv_from_kernel"].(int)
		ret.Send_to_kernel = in["send_to_kernel"].(int)
		ret.Send_to_kernel_err = in["send_to_kernel_err"].(int)
		ret.Sby_no_peer = in["sby_no_peer"].(int)
		ret.Dcmsg_err = in["dcmsg_err"].(int)
		ret.No_slb_object = in["no_slb_object"].(int)
		ret.Smart_nat_init_port_err = in["smart_nat_init_port_err"].(int)
		ret.Smart_nat_init_inst_err = in["smart_nat_init_inst_err"].(int)
		ret.Smart_nat_rserver_route_err = in["smart_nat_rserver_route_err"].(int)
		ret.Smart_nat_rserver_ip_err = in["smart_nat_rserver_ip_err"].(int)
		ret.Nat_resource_err = in["nat_resource_err"].(int)
		ret.Frag_err = in["frag_err"].(int)
	}
	return ret
}

func dataToEndpointHealthSourceNatStats(d *schema.ResourceData) edpt.HealthSourceNatStats {
	var ret edpt.HealthSourceNatStats

	ret.Stats = getObjectHealthSourceNatStatsStats(d.Get("stats").([]interface{}))
	return ret
}
