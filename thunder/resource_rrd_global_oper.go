package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRrdGlobalOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_rrd_global_oper`: Operational Status for the object global\n\n__PLACEHOLDER__",
		ReadContext: resourceRrdGlobalOperRead,

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
						"global_data": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"time": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_est_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_half_open": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"udp_conn_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ip_conn_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"other_conn_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"snat_tcp_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"snat_udp_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"free_buff_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"curr_free_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"conn_get_cnt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"conn_free_cnt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"syn_half_open": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"conn_smp_alloc": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"conn_smp_free": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"conn_smp_aged": {
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

func resourceRrdGlobalOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRrdGlobalOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRrdGlobalOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		RrdGlobalOperOper := setObjectRrdGlobalOperOper(res)
		d.Set("oper", RrdGlobalOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectRrdGlobalOperOper(ret edpt.DataRrdGlobalOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"start_time":  ret.DtRrdGlobalOper.Oper.StartTime,
			"end_time":    ret.DtRrdGlobalOper.Oper.EndTime,
			"global_data": setSliceRrdGlobalOperOperGlobalData(ret.DtRrdGlobalOper.Oper.GlobalData),
		},
	}
}

func setSliceRrdGlobalOperOperGlobalData(d []edpt.RrdGlobalOperOperGlobalData) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["time"] = item.Time
		in["tcp_est_count"] = item.Tcp_est_count
		in["tcp_half_open"] = item.Tcp_half_open
		in["udp_conn_count"] = item.Udp_conn_count
		in["ip_conn_count"] = item.Ip_conn_count
		in["other_conn_count"] = item.Other_conn_count
		in["snat_tcp_count"] = item.Snat_tcp_count
		in["snat_udp_count"] = item.Snat_udp_count
		in["free_buff_count"] = item.Free_buff_count
		in["curr_free_conn"] = item.Curr_free_conn
		in["conn_get_cnt"] = item.Conn_get_cnt
		in["conn_free_cnt"] = item.Conn_free_cnt
		in["syn_half_open"] = item.Syn_half_open
		in["conn_smp_alloc"] = item.Conn_smp_alloc
		in["conn_smp_free"] = item.Conn_smp_free
		in["conn_smp_aged"] = item.Conn_smp_aged
		result = append(result, in)
	}
	return result
}

func getObjectRrdGlobalOperOper(d []interface{}) edpt.RrdGlobalOperOper {

	count1 := len(d)
	var ret edpt.RrdGlobalOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StartTime = in["start_time"].(int)
		ret.EndTime = in["end_time"].(int)
		ret.GlobalData = getSliceRrdGlobalOperOperGlobalData(in["global_data"].([]interface{}))
	}
	return ret
}

func getSliceRrdGlobalOperOperGlobalData(d []interface{}) []edpt.RrdGlobalOperOperGlobalData {

	count1 := len(d)
	ret := make([]edpt.RrdGlobalOperOperGlobalData, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RrdGlobalOperOperGlobalData
		oi.Time = in["time"].(int)
		oi.Tcp_est_count = in["tcp_est_count"].(int)
		oi.Tcp_half_open = in["tcp_half_open"].(int)
		oi.Udp_conn_count = in["udp_conn_count"].(int)
		oi.Ip_conn_count = in["ip_conn_count"].(int)
		oi.Other_conn_count = in["other_conn_count"].(int)
		oi.Snat_tcp_count = in["snat_tcp_count"].(int)
		oi.Snat_udp_count = in["snat_udp_count"].(int)
		oi.Free_buff_count = in["free_buff_count"].(int)
		oi.Curr_free_conn = in["curr_free_conn"].(int)
		oi.Conn_get_cnt = in["conn_get_cnt"].(int)
		oi.Conn_free_cnt = in["conn_free_cnt"].(int)
		oi.Syn_half_open = in["syn_half_open"].(int)
		oi.Conn_smp_alloc = in["conn_smp_alloc"].(int)
		oi.Conn_smp_free = in["conn_smp_free"].(int)
		oi.Conn_smp_aged = in["conn_smp_aged"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRrdGlobalOper(d *schema.ResourceData) edpt.RrdGlobalOper {
	var ret edpt.RrdGlobalOper

	ret.Oper = getObjectRrdGlobalOperOper(d.Get("oper").([]interface{}))
	return ret
}
