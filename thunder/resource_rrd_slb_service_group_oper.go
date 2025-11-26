package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRrdSlbServiceGroupOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_rrd_slb_service_group_oper`: Operational Status for the object slb-service-group\n\n__PLACEHOLDER__",
		ReadContext: resourceRrdSlbServiceGroupOperRead,

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
						"slb_service_group_name": {
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

func resourceRrdSlbServiceGroupOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRrdSlbServiceGroupOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRrdSlbServiceGroupOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		RrdSlbServiceGroupOperOper := setObjectRrdSlbServiceGroupOperOper(res)
		d.Set("oper", RrdSlbServiceGroupOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectRrdSlbServiceGroupOperOper(ret edpt.DataRrdSlbServiceGroupOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"start_time":             ret.DtRrdSlbServiceGroupOper.Oper.StartTime,
			"end_time":               ret.DtRrdSlbServiceGroupOper.Oper.EndTime,
			"slb_service_group_name": ret.DtRrdSlbServiceGroupOper.Oper.SlbServiceGroupName,
			"slb_server_statistics":  setSliceRrdSlbServiceGroupOperOperSlbServerStatistics(ret.DtRrdSlbServiceGroupOper.Oper.SlbServerStatistics),
		},
	}
}

func setSliceRrdSlbServiceGroupOperOperSlbServerStatistics(d []edpt.RrdSlbServiceGroupOperOperSlbServerStatistics) []map[string]interface{} {
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

func getObjectRrdSlbServiceGroupOperOper(d []interface{}) edpt.RrdSlbServiceGroupOperOper {

	count1 := len(d)
	var ret edpt.RrdSlbServiceGroupOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StartTime = in["start_time"].(int)
		ret.EndTime = in["end_time"].(int)
		ret.SlbServiceGroupName = in["slb_service_group_name"].(string)
		ret.SlbServerStatistics = getSliceRrdSlbServiceGroupOperOperSlbServerStatistics(in["slb_server_statistics"].([]interface{}))
	}
	return ret
}

func getSliceRrdSlbServiceGroupOperOperSlbServerStatistics(d []interface{}) []edpt.RrdSlbServiceGroupOperOperSlbServerStatistics {

	count1 := len(d)
	ret := make([]edpt.RrdSlbServiceGroupOperOperSlbServerStatistics, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RrdSlbServiceGroupOperOperSlbServerStatistics
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

func dataToEndpointRrdSlbServiceGroupOper(d *schema.ResourceData) edpt.RrdSlbServiceGroupOper {
	var ret edpt.RrdSlbServiceGroupOper

	ret.Oper = getObjectRrdSlbServiceGroupOperOper(d.Get("oper").([]interface{}))
	return ret
}
