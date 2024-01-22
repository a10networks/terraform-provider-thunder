package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbFixOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_fix_oper`: Operational Status for the object fix\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbFixOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fix_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"curr_proxy": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_proxy": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"svrsel_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"noroute": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"snat_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"insert_clientip": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"default_switching": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sender_switching": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"target_switching": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_tls_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_tls_conn": {
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

func resourceSlbFixOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbFixOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbFixOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbFixOperOper := setObjectSlbFixOperOper(res)
		d.Set("oper", SlbFixOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbFixOperOper(ret edpt.DataSlbFixOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"fix_cpu_list": setSliceSlbFixOperOperFixCpuList(ret.DtSlbFixOper.Oper.FixCpuList),
			"cpu_count":    ret.DtSlbFixOper.Oper.CpuCount,
		},
	}
}

func setSliceSlbFixOperOperFixCpuList(d []edpt.SlbFixOperOperFixCpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["curr_proxy"] = item.Curr_proxy
		in["total_proxy"] = item.Total_proxy
		in["svrsel_fail"] = item.Svrsel_fail
		in["noroute"] = item.Noroute
		in["snat_fail"] = item.Snat_fail
		in["client_err"] = item.Client_err
		in["server_err"] = item.Server_err
		in["insert_clientip"] = item.Insert_clientip
		in["default_switching"] = item.Default_switching
		in["sender_switching"] = item.Sender_switching
		in["target_switching"] = item.Target_switching
		in["client_tls_conn"] = item.Client_tls_conn
		in["server_tls_conn"] = item.Server_tls_conn
		result = append(result, in)
	}
	return result
}

func getObjectSlbFixOperOper(d []interface{}) edpt.SlbFixOperOper {

	count1 := len(d)
	var ret edpt.SlbFixOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FixCpuList = getSliceSlbFixOperOperFixCpuList(in["fix_cpu_list"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
	}
	return ret
}

func getSliceSlbFixOperOperFixCpuList(d []interface{}) []edpt.SlbFixOperOperFixCpuList {

	count1 := len(d)
	ret := make([]edpt.SlbFixOperOperFixCpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbFixOperOperFixCpuList
		oi.Curr_proxy = in["curr_proxy"].(int)
		oi.Total_proxy = in["total_proxy"].(int)
		oi.Svrsel_fail = in["svrsel_fail"].(int)
		oi.Noroute = in["noroute"].(int)
		oi.Snat_fail = in["snat_fail"].(int)
		oi.Client_err = in["client_err"].(int)
		oi.Server_err = in["server_err"].(int)
		oi.Insert_clientip = in["insert_clientip"].(int)
		oi.Default_switching = in["default_switching"].(int)
		oi.Sender_switching = in["sender_switching"].(int)
		oi.Target_switching = in["target_switching"].(int)
		oi.Client_tls_conn = in["client_tls_conn"].(int)
		oi.Server_tls_conn = in["server_tls_conn"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbFixOper(d *schema.ResourceData) edpt.SlbFixOper {
	var ret edpt.SlbFixOper

	ret.Oper = getObjectSlbFixOperOper(d.Get("oper").([]interface{}))
	return ret
}
