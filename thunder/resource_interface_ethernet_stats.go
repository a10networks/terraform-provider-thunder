package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceEthernetStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_interface_ethernet_stats`: Statistics for the object ethernet\n\n__PLACEHOLDER__",
		ReadContext: resourceInterfaceEthernetStatsRead,

		Schema: map[string]*schema.Schema{
			"ifnum": {
				Type: schema.TypeInt, Required: true, Description: "Ethernet interface number",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"packets_input": {
							Type: schema.TypeInt, Optional: true, Description: "Input packets",
						},
						"bytes_input": {
							Type: schema.TypeInt, Optional: true, Description: "Input bytes",
						},
						"received_broadcasts": {
							Type: schema.TypeInt, Optional: true, Description: "Received broadcasts",
						},
						"received_multicasts": {
							Type: schema.TypeInt, Optional: true, Description: "Received multicasts",
						},
						"received_unicasts": {
							Type: schema.TypeInt, Optional: true, Description: "Received unicasts",
						},
						"input_errors": {
							Type: schema.TypeInt, Optional: true, Description: "Input errors",
						},
						"crc": {
							Type: schema.TypeInt, Optional: true, Description: "CRC",
						},
						"frame": {
							Type: schema.TypeInt, Optional: true, Description: "Frames",
						},
						"runts": {
							Type: schema.TypeInt, Optional: true, Description: "Runts",
						},
						"giants": {
							Type: schema.TypeInt, Optional: true, Description: "Giants",
						},
						"packets_output": {
							Type: schema.TypeInt, Optional: true, Description: "Output packets",
						},
						"bytes_output": {
							Type: schema.TypeInt, Optional: true, Description: "Output bytes",
						},
						"transmitted_broadcasts": {
							Type: schema.TypeInt, Optional: true, Description: "Transmitted broadcasts",
						},
						"transmitted_multicasts": {
							Type: schema.TypeInt, Optional: true, Description: "Transmitted multicasts",
						},
						"transmitted_unicasts": {
							Type: schema.TypeInt, Optional: true, Description: "Transmitted unicasts",
						},
						"output_errors": {
							Type: schema.TypeInt, Optional: true, Description: "Output errors",
						},
						"collisions": {
							Type: schema.TypeInt, Optional: true, Description: "Collisions",
						},
						"giants_output": {
							Type: schema.TypeInt, Optional: true, Description: "Output Giants",
						},
						"rate_pkt_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Packet sent rate packets/sec",
						},
						"rate_byte_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Byte sent rate bits/sec",
						},
						"rate_pkt_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Packet received rate packets/sec",
						},
						"rate_byte_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Byte received rate bits/sec",
						},
						"load_interval": {
							Type: schema.TypeInt, Optional: true, Description: "Load Interval",
						},
						"drops": {
							Type: schema.TypeInt, Optional: true, Description: "Drops",
						},
						"input_utilization": {
							Type: schema.TypeInt, Optional: true, Description: "Input Utilization",
						},
						"output_utilization": {
							Type: schema.TypeInt, Optional: true, Description: "Output Utilization",
						},
					},
				},
			},
		},
	}
}

func resourceInterfaceEthernetStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		InterfaceEthernetStatsStats := setObjectInterfaceEthernetStatsStats(res)
		d.Set("stats", InterfaceEthernetStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectInterfaceEthernetStatsStats(ret edpt.DataInterfaceEthernetStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"packets_input":          ret.DtInterfaceEthernetStats.Stats.Packets_input,
			"bytes_input":            ret.DtInterfaceEthernetStats.Stats.Bytes_input,
			"received_broadcasts":    ret.DtInterfaceEthernetStats.Stats.Received_broadcasts,
			"received_multicasts":    ret.DtInterfaceEthernetStats.Stats.Received_multicasts,
			"received_unicasts":      ret.DtInterfaceEthernetStats.Stats.Received_unicasts,
			"input_errors":           ret.DtInterfaceEthernetStats.Stats.Input_errors,
			"crc":                    ret.DtInterfaceEthernetStats.Stats.Crc,
			"frame":                  ret.DtInterfaceEthernetStats.Stats.Frame,
			"runts":                  ret.DtInterfaceEthernetStats.Stats.Runts,
			"giants":                 ret.DtInterfaceEthernetStats.Stats.Giants,
			"packets_output":         ret.DtInterfaceEthernetStats.Stats.Packets_output,
			"bytes_output":           ret.DtInterfaceEthernetStats.Stats.Bytes_output,
			"transmitted_broadcasts": ret.DtInterfaceEthernetStats.Stats.Transmitted_broadcasts,
			"transmitted_multicasts": ret.DtInterfaceEthernetStats.Stats.Transmitted_multicasts,
			"transmitted_unicasts":   ret.DtInterfaceEthernetStats.Stats.Transmitted_unicasts,
			"output_errors":          ret.DtInterfaceEthernetStats.Stats.Output_errors,
			"collisions":             ret.DtInterfaceEthernetStats.Stats.Collisions,
			"giants_output":          ret.DtInterfaceEthernetStats.Stats.Giants_output,
			"rate_pkt_sent":          ret.DtInterfaceEthernetStats.Stats.Rate_pkt_sent,
			"rate_byte_sent":         ret.DtInterfaceEthernetStats.Stats.Rate_byte_sent,
			"rate_pkt_rcvd":          ret.DtInterfaceEthernetStats.Stats.Rate_pkt_rcvd,
			"rate_byte_rcvd":         ret.DtInterfaceEthernetStats.Stats.Rate_byte_rcvd,
			"load_interval":          ret.DtInterfaceEthernetStats.Stats.Load_interval,
			"drops":                  ret.DtInterfaceEthernetStats.Stats.Drops,
			"input_utilization":      ret.DtInterfaceEthernetStats.Stats.Input_utilization,
			"output_utilization":     ret.DtInterfaceEthernetStats.Stats.Output_utilization,
		},
	}
}

func getObjectInterfaceEthernetStatsStats(d []interface{}) edpt.InterfaceEthernetStatsStats {

	count1 := len(d)
	var ret edpt.InterfaceEthernetStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Packets_input = in["packets_input"].(int)
		ret.Bytes_input = in["bytes_input"].(int)
		ret.Received_broadcasts = in["received_broadcasts"].(int)
		ret.Received_multicasts = in["received_multicasts"].(int)
		ret.Received_unicasts = in["received_unicasts"].(int)
		ret.Input_errors = in["input_errors"].(int)
		ret.Crc = in["crc"].(int)
		ret.Frame = in["frame"].(int)
		ret.Runts = in["runts"].(int)
		ret.Giants = in["giants"].(int)
		ret.Packets_output = in["packets_output"].(int)
		ret.Bytes_output = in["bytes_output"].(int)
		ret.Transmitted_broadcasts = in["transmitted_broadcasts"].(int)
		ret.Transmitted_multicasts = in["transmitted_multicasts"].(int)
		ret.Transmitted_unicasts = in["transmitted_unicasts"].(int)
		ret.Output_errors = in["output_errors"].(int)
		ret.Collisions = in["collisions"].(int)
		ret.Giants_output = in["giants_output"].(int)
		ret.Rate_pkt_sent = in["rate_pkt_sent"].(int)
		ret.Rate_byte_sent = in["rate_byte_sent"].(int)
		ret.Rate_pkt_rcvd = in["rate_pkt_rcvd"].(int)
		ret.Rate_byte_rcvd = in["rate_byte_rcvd"].(int)
		ret.Load_interval = in["load_interval"].(int)
		ret.Drops = in["drops"].(int)
		ret.Input_utilization = in["input_utilization"].(int)
		ret.Output_utilization = in["output_utilization"].(int)
	}
	return ret
}

func dataToEndpointInterfaceEthernetStats(d *schema.ResourceData) edpt.InterfaceEthernetStats {
	var ret edpt.InterfaceEthernetStats

	ret.Ifnum = d.Get("ifnum").(int)

	ret.Stats = getObjectInterfaceEthernetStatsStats(d.Get("stats").([]interface{}))
	return ret
}
