package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthLogOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_auth_log_oper`: Operational Status for the object auth-log\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthLogOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tail": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"last": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"top": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"target": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"auth_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"total": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"log_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"record_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"time": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"user_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"user_domain": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"client_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"destination_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"source_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"destination_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"auth_result": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"counter": {
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

func resourceAamAuthLogOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthLogOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthLogOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthLogOperOper := setObjectAamAuthLogOperOper(res)
		d.Set("oper", AamAuthLogOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthLogOperOper(ret edpt.DataAamAuthLogOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"tail":        ret.DtAamAuthLogOper.Oper.Tail,
			"last":        ret.DtAamAuthLogOper.Oper.Last,
			"top":         ret.DtAamAuthLogOper.Oper.Top,
			"target":      ret.DtAamAuthLogOper.Oper.Target,
			"auth_status": ret.DtAamAuthLogOper.Oper.AuthStatus,
			"total":       ret.DtAamAuthLogOper.Oper.Total,
			"log_list":    setSliceAamAuthLogOperOperLogList(ret.DtAamAuthLogOper.Oper.LogList),
		},
	}
}

func setSliceAamAuthLogOperOperLogList(d []edpt.AamAuthLogOperOperLogList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["record_id"] = item.RecordId
		in["time"] = item.Time
		in["user_name"] = item.UserName
		in["user_domain"] = item.UserDomain
		in["client_ip"] = item.ClientIp
		in["destination_ip"] = item.DestinationIp
		in["source_port"] = item.SourcePort
		in["destination_port"] = item.DestinationPort
		in["auth_result"] = item.AuthResult
		in["counter"] = item.Counter
		result = append(result, in)
	}
	return result
}

func getObjectAamAuthLogOperOper(d []interface{}) edpt.AamAuthLogOperOper {

	count1 := len(d)
	var ret edpt.AamAuthLogOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tail = in["tail"].(int)
		ret.Last = in["last"].(int)
		ret.Top = in["top"].(int)
		ret.Target = in["target"].(string)
		ret.AuthStatus = in["auth_status"].(string)
		ret.Total = in["total"].(int)
		ret.LogList = getSliceAamAuthLogOperOperLogList(in["log_list"].([]interface{}))
	}
	return ret
}

func getSliceAamAuthLogOperOperLogList(d []interface{}) []edpt.AamAuthLogOperOperLogList {

	count1 := len(d)
	ret := make([]edpt.AamAuthLogOperOperLogList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthLogOperOperLogList
		oi.RecordId = in["record_id"].(int)
		oi.Time = in["time"].(string)
		oi.UserName = in["user_name"].(string)
		oi.UserDomain = in["user_domain"].(string)
		oi.ClientIp = in["client_ip"].(string)
		oi.DestinationIp = in["destination_ip"].(string)
		oi.SourcePort = in["source_port"].(int)
		oi.DestinationPort = in["destination_port"].(int)
		oi.AuthResult = in["auth_result"].(string)
		oi.Counter = in["counter"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthLogOper(d *schema.ResourceData) edpt.AamAuthLogOper {
	var ret edpt.AamAuthLogOper

	ret.Oper = getObjectAamAuthLogOperOper(d.Get("oper").([]interface{}))
	return ret
}
