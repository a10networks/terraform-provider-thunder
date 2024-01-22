package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbConnectionReuseOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_connection_reuse_oper`: Operational Status for the object connection-reuse\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbConnectionReuseOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connection_reuse_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"current_open": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"current_active": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nbind": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nunbind": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nestab": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ntermi": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ntermi_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"delay_unbind": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"long_resp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"miss_resp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"unbound_data_rcv": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pause_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pause_conn_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"resume_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"not_remove_from_rport": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"cpu_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbConnectionReuseOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbConnectionReuseOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbConnectionReuseOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbConnectionReuseOperOper := setObjectSlbConnectionReuseOperOper(res)
		d.Set("oper", SlbConnectionReuseOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbConnectionReuseOperOper(ret edpt.DataSlbConnectionReuseOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"connection_reuse_cpu_list": setSliceSlbConnectionReuseOperOperConnectionReuseCpuList(ret.DtSlbConnectionReuseOper.Oper.ConnectionReuseCpuList),
			"cpu_count":                 ret.DtSlbConnectionReuseOper.Oper.CpuCount,
		},
	}
}

func setSliceSlbConnectionReuseOperOperConnectionReuseCpuList(d []edpt.SlbConnectionReuseOperOperConnectionReuseCpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["current_open"] = item.Current_open
		in["current_active"] = item.Current_active
		in["nbind"] = item.Nbind
		in["nunbind"] = item.Nunbind
		in["nestab"] = item.Nestab
		in["ntermi"] = item.Ntermi
		in["ntermi_err"] = item.Ntermi_err
		in["delay_unbind"] = item.Delay_unbind
		in["long_resp"] = item.Long_resp
		in["miss_resp"] = item.Miss_resp
		in["unbound_data_rcv"] = item.Unbound_data_rcv
		in["pause_conn"] = item.Pause_conn
		in["pause_conn_fail"] = item.Pause_conn_fail
		in["resume_conn"] = item.Resume_conn
		in["not_remove_from_rport"] = item.Not_remove_from_rport
		result = append(result, in)
	}
	return result
}

func getObjectSlbConnectionReuseOperOper(d []interface{}) edpt.SlbConnectionReuseOperOper {

	count1 := len(d)
	var ret edpt.SlbConnectionReuseOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ConnectionReuseCpuList = getSliceSlbConnectionReuseOperOperConnectionReuseCpuList(in["connection_reuse_cpu_list"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
	}
	return ret
}

func getSliceSlbConnectionReuseOperOperConnectionReuseCpuList(d []interface{}) []edpt.SlbConnectionReuseOperOperConnectionReuseCpuList {

	count1 := len(d)
	ret := make([]edpt.SlbConnectionReuseOperOperConnectionReuseCpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbConnectionReuseOperOperConnectionReuseCpuList
		oi.Current_open = in["current_open"].(int)
		oi.Current_active = in["current_active"].(int)
		oi.Nbind = in["nbind"].(int)
		oi.Nunbind = in["nunbind"].(int)
		oi.Nestab = in["nestab"].(int)
		oi.Ntermi = in["ntermi"].(int)
		oi.Ntermi_err = in["ntermi_err"].(int)
		oi.Delay_unbind = in["delay_unbind"].(int)
		oi.Long_resp = in["long_resp"].(int)
		oi.Miss_resp = in["miss_resp"].(int)
		oi.Unbound_data_rcv = in["unbound_data_rcv"].(int)
		oi.Pause_conn = in["pause_conn"].(int)
		oi.Pause_conn_fail = in["pause_conn_fail"].(int)
		oi.Resume_conn = in["resume_conn"].(int)
		oi.Not_remove_from_rport = in["not_remove_from_rport"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbConnectionReuseOper(d *schema.ResourceData) edpt.SlbConnectionReuseOper {
	var ret edpt.SlbConnectionReuseOper

	ret.Oper = getObjectSlbConnectionReuseOperOper(d.Get("oper").([]interface{}))
	return ret
}
