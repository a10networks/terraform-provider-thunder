package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRrdSlbVirtualServerOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_rrd_slb_virtual_server_oper`: Operational Status for the object slb-virtual-server\n\n__PLACEHOLDER__",
		ReadContext: resourceRrdSlbVirtualServerOperRead,

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
						"slb_virtual_server_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"slb_server_statistics": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"time": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"in_pkts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"in_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"out_pkts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"out_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cur_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"p_conn": {
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

func resourceRrdSlbVirtualServerOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRrdSlbVirtualServerOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRrdSlbVirtualServerOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		RrdSlbVirtualServerOperOper := setObjectRrdSlbVirtualServerOperOper(res)
		d.Set("oper", RrdSlbVirtualServerOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectRrdSlbVirtualServerOperOper(ret edpt.DataRrdSlbVirtualServerOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"start_time":              ret.DtRrdSlbVirtualServerOper.Oper.StartTime,
			"end_time":                ret.DtRrdSlbVirtualServerOper.Oper.EndTime,
			"slb_virtual_server_name": ret.DtRrdSlbVirtualServerOper.Oper.SlbVirtualServerName,
			"slb_server_statistics":   setSliceRrdSlbVirtualServerOperOperSlbServerStatistics(ret.DtRrdSlbVirtualServerOper.Oper.SlbServerStatistics),
		},
	}
}

func setSliceRrdSlbVirtualServerOperOperSlbServerStatistics(d []edpt.RrdSlbVirtualServerOperOperSlbServerStatistics) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["time"] = item.Time
		in["in_pkts"] = item.In_pkts
		in["in_bytes"] = item.In_bytes
		in["out_pkts"] = item.Out_pkts
		in["out_bytes"] = item.Out_bytes
		in["cur_conn"] = item.Cur_conn
		in["p_conn"] = item.P_conn
		result = append(result, in)
	}
	return result
}

func getObjectRrdSlbVirtualServerOperOper(d []interface{}) edpt.RrdSlbVirtualServerOperOper {

	count1 := len(d)
	var ret edpt.RrdSlbVirtualServerOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StartTime = in["start_time"].(int)
		ret.EndTime = in["end_time"].(int)
		ret.SlbVirtualServerName = in["slb_virtual_server_name"].(string)
		ret.SlbServerStatistics = getSliceRrdSlbVirtualServerOperOperSlbServerStatistics(in["slb_server_statistics"].([]interface{}))
	}
	return ret
}

func getSliceRrdSlbVirtualServerOperOperSlbServerStatistics(d []interface{}) []edpt.RrdSlbVirtualServerOperOperSlbServerStatistics {

	count1 := len(d)
	ret := make([]edpt.RrdSlbVirtualServerOperOperSlbServerStatistics, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RrdSlbVirtualServerOperOperSlbServerStatistics
		oi.Time = in["time"].(int)
		oi.In_pkts = in["in_pkts"].(int)
		oi.In_bytes = in["in_bytes"].(int)
		oi.Out_pkts = in["out_pkts"].(int)
		oi.Out_bytes = in["out_bytes"].(int)
		oi.Cur_conn = in["cur_conn"].(int)
		oi.P_conn = in["p_conn"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRrdSlbVirtualServerOper(d *schema.ResourceData) edpt.RrdSlbVirtualServerOper {
	var ret edpt.RrdSlbVirtualServerOper

	ret.Oper = getObjectRrdSlbVirtualServerOperOper(d.Get("oper").([]interface{}))
	return ret
}
