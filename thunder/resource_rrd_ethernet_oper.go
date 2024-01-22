package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRrdEthernetOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_rrd_ethernet_oper`: Operational Status for the object ethernet\n\n__PLACEHOLDER__",
		ReadContext: resourceRrdEthernetOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"start_time": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"end_time": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ethernet_index": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ethernet_statistics": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"time": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rx_packets": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rx_bits": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rx_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rx_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tx_packets": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tx_bits": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tx_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tx_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
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

func resourceRrdEthernetOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRrdEthernetOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRrdEthernetOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		RrdEthernetOperOper := setObjectRrdEthernetOperOper(res)
		d.Set("oper", RrdEthernetOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectRrdEthernetOperOper(ret edpt.DataRrdEthernetOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"start_time":          ret.DtRrdEthernetOper.Oper.StartTime,
			"end_time":            ret.DtRrdEthernetOper.Oper.EndTime,
			"ethernet_index":      ret.DtRrdEthernetOper.Oper.EthernetIndex,
			"ethernet_statistics": setSliceRrdEthernetOperOperEthernetStatistics(ret.DtRrdEthernetOper.Oper.EthernetStatistics),
		},
	}
}

func setSliceRrdEthernetOperOperEthernetStatistics(d []edpt.RrdEthernetOperOperEthernetStatistics) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["time"] = item.Time
		in["rx_packets"] = item.Rx_packets
		in["rx_bits"] = item.Rx_bits
		in["rx_error"] = item.Rx_error
		in["rx_drop"] = item.Rx_drop
		in["tx_packets"] = item.Tx_packets
		in["tx_bits"] = item.Tx_bits
		in["tx_error"] = item.Tx_error
		in["tx_drop"] = item.Tx_drop
		result = append(result, in)
	}
	return result
}

func getObjectRrdEthernetOperOper(d []interface{}) edpt.RrdEthernetOperOper {

	count1 := len(d)
	var ret edpt.RrdEthernetOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StartTime = in["start_time"].(int)
		ret.EndTime = in["end_time"].(int)
		ret.EthernetIndex = in["ethernet_index"].(int)
		ret.EthernetStatistics = getSliceRrdEthernetOperOperEthernetStatistics(in["ethernet_statistics"].([]interface{}))
	}
	return ret
}

func getSliceRrdEthernetOperOperEthernetStatistics(d []interface{}) []edpt.RrdEthernetOperOperEthernetStatistics {

	count1 := len(d)
	ret := make([]edpt.RrdEthernetOperOperEthernetStatistics, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RrdEthernetOperOperEthernetStatistics
		oi.Time = in["time"].(int)
		oi.Rx_packets = in["rx_packets"].(int)
		oi.Rx_bits = in["rx_bits"].(int)
		oi.Rx_error = in["rx_error"].(int)
		oi.Rx_drop = in["rx_drop"].(int)
		oi.Tx_packets = in["tx_packets"].(int)
		oi.Tx_bits = in["tx_bits"].(int)
		oi.Tx_error = in["tx_error"].(int)
		oi.Tx_drop = in["tx_drop"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRrdEthernetOper(d *schema.ResourceData) edpt.RrdEthernetOper {
	var ret edpt.RrdEthernetOper

	ret.Oper = getObjectRrdEthernetOperOper(d.Get("oper").([]interface{}))
	return ret
}
