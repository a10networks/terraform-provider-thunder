package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceManagementStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_interface_management_stats`: Statistics for the object management\n\n__PLACEHOLDER__",
		ReadContext: resourceInterfaceManagementStatsRead,

		Schema: map[string]*schema.Schema{
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
						"input_err_short": {
							Type: schema.TypeInt, Optional: true, Description: "Runts",
						},
						"input_err_long": {
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
					},
				},
			},
		},
	}
}

func resourceInterfaceManagementStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceManagementStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceManagementStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		InterfaceManagementStatsStats := setObjectInterfaceManagementStatsStats(res)
		d.Set("stats", InterfaceManagementStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectInterfaceManagementStatsStats(ret edpt.DataInterfaceManagementStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"packets_input":          ret.DtInterfaceManagementStats.Stats.Packets_input,
			"bytes_input":            ret.DtInterfaceManagementStats.Stats.Bytes_input,
			"received_broadcasts":    ret.DtInterfaceManagementStats.Stats.Received_broadcasts,
			"received_multicasts":    ret.DtInterfaceManagementStats.Stats.Received_multicasts,
			"received_unicasts":      ret.DtInterfaceManagementStats.Stats.Received_unicasts,
			"input_errors":           ret.DtInterfaceManagementStats.Stats.Input_errors,
			"crc":                    ret.DtInterfaceManagementStats.Stats.Crc,
			"frame":                  ret.DtInterfaceManagementStats.Stats.Frame,
			"input_err_short":        ret.DtInterfaceManagementStats.Stats.Input_err_short,
			"input_err_long":         ret.DtInterfaceManagementStats.Stats.Input_err_long,
			"packets_output":         ret.DtInterfaceManagementStats.Stats.Packets_output,
			"bytes_output":           ret.DtInterfaceManagementStats.Stats.Bytes_output,
			"transmitted_broadcasts": ret.DtInterfaceManagementStats.Stats.Transmitted_broadcasts,
			"transmitted_multicasts": ret.DtInterfaceManagementStats.Stats.Transmitted_multicasts,
			"transmitted_unicasts":   ret.DtInterfaceManagementStats.Stats.Transmitted_unicasts,
			"output_errors":          ret.DtInterfaceManagementStats.Stats.Output_errors,
			"collisions":             ret.DtInterfaceManagementStats.Stats.Collisions,
		},
	}
}

func getObjectInterfaceManagementStatsStats(d []interface{}) edpt.InterfaceManagementStatsStats {

	count1 := len(d)
	var ret edpt.InterfaceManagementStatsStats
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
		ret.Input_err_short = in["input_err_short"].(int)
		ret.Input_err_long = in["input_err_long"].(int)
		ret.Packets_output = in["packets_output"].(int)
		ret.Bytes_output = in["bytes_output"].(int)
		ret.Transmitted_broadcasts = in["transmitted_broadcasts"].(int)
		ret.Transmitted_multicasts = in["transmitted_multicasts"].(int)
		ret.Transmitted_unicasts = in["transmitted_unicasts"].(int)
		ret.Output_errors = in["output_errors"].(int)
		ret.Collisions = in["collisions"].(int)
	}
	return ret
}

func dataToEndpointInterfaceManagementStats(d *schema.ResourceData) edpt.InterfaceManagementStats {
	var ret edpt.InterfaceManagementStats

	ret.Stats = getObjectInterfaceManagementStatsStats(d.Get("stats").([]interface{}))
	return ret
}
