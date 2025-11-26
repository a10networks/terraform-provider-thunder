package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAccessLogOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_access_log_oper`: Operational Status for the object access-log\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAccessLogOperRead,

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
						"access_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"host_name": {
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
									"policy": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"action": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"vip_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"vip_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"host": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"uri": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"web_category": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ssl_status": {
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

func resourceAamAccessLogOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAccessLogOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAccessLogOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAccessLogOperOper := setObjectAamAccessLogOperOper(res)
		d.Set("oper", AamAccessLogOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAccessLogOperOper(ret edpt.DataAamAccessLogOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"tail":          ret.DtAamAccessLogOper.Oper.Tail,
			"last":          ret.DtAamAccessLogOper.Oper.Last,
			"top":           ret.DtAamAccessLogOper.Oper.Top,
			"target":        ret.DtAamAccessLogOper.Oper.Target,
			"access_status": ret.DtAamAccessLogOper.Oper.AccessStatus,
			"host_name":     ret.DtAamAccessLogOper.Oper.HostName,
			"total":         ret.DtAamAccessLogOper.Oper.Total,
			"log_list":      setSliceAamAccessLogOperOperLogList(ret.DtAamAccessLogOper.Oper.LogList),
		},
	}
}

func setSliceAamAccessLogOperOperLogList(d []edpt.AamAccessLogOperOperLogList) []map[string]interface{} {
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
		in["policy"] = item.Policy
		in["action"] = item.Action
		in["vip_name"] = item.VipName
		in["vip_port"] = item.VipPort
		in["host"] = item.Host
		in["uri"] = item.Uri
		in["web_category"] = item.WebCategory
		in["ssl_status"] = item.SslStatus
		in["counter"] = item.Counter
		result = append(result, in)
	}
	return result
}

func getObjectAamAccessLogOperOper(d []interface{}) edpt.AamAccessLogOperOper {

	count1 := len(d)
	var ret edpt.AamAccessLogOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tail = in["tail"].(int)
		ret.Last = in["last"].(int)
		ret.Top = in["top"].(int)
		ret.Target = in["target"].(string)
		ret.AccessStatus = in["access_status"].(string)
		ret.HostName = in["host_name"].(string)
		ret.Total = in["total"].(int)
		ret.LogList = getSliceAamAccessLogOperOperLogList(in["log_list"].([]interface{}))
	}
	return ret
}

func getSliceAamAccessLogOperOperLogList(d []interface{}) []edpt.AamAccessLogOperOperLogList {

	count1 := len(d)
	ret := make([]edpt.AamAccessLogOperOperLogList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAccessLogOperOperLogList
		oi.RecordId = in["record_id"].(int)
		oi.Time = in["time"].(string)
		oi.UserName = in["user_name"].(string)
		oi.UserDomain = in["user_domain"].(string)
		oi.ClientIp = in["client_ip"].(string)
		oi.DestinationIp = in["destination_ip"].(string)
		oi.SourcePort = in["source_port"].(int)
		oi.DestinationPort = in["destination_port"].(int)
		oi.Policy = in["policy"].(string)
		oi.Action = in["action"].(string)
		oi.VipName = in["vip_name"].(string)
		oi.VipPort = in["vip_port"].(int)
		oi.Host = in["host"].(string)
		oi.Uri = in["uri"].(string)
		oi.WebCategory = in["web_category"].(string)
		oi.SslStatus = in["ssl_status"].(string)
		oi.Counter = in["counter"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAccessLogOper(d *schema.ResourceData) edpt.AamAccessLogOper {
	var ret edpt.AamAccessLogOper

	ret.Oper = getObjectAamAccessLogOperOper(d.Get("oper").([]interface{}))
	return ret
}
