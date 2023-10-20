package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSyslogOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_syslog_oper`: Operational Status for the object syslog\n\n__PLACEHOLDER__",
		ReadContext: resourceSyslogOperRead,
		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"system_log": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"log_data": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"add_slot_info": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"log_data_search": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"next_msg_idx": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSyslogOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSyslogOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSyslogOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		opers := setObjectSyslogOperOper(res)
		d.SetId(obj.GetId())
		d.Set("oper", opers)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSyslogOperOper(res edpt.Syslogg) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["system_log"] = setSliceSyslogOperOperSystemLog(res.DataSyslog.Oper.SystemLog)
	in["next_msg_idx"] = res.DataSyslog.Oper.NextMsgIdx
	result = append(result, in)
	return result
}

func setSliceSyslogOperOperSystemLog(d []edpt.SyslogOperOperSystemLog) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["log_data"] = item.LogData
		in["add_slot_info"] = item.AddSlotInfo
		in["log_data_search"] = item.LogDataSearch
		result = append(result, in)
	}
	return result
}

func getObjectSyslogOperOper(d []interface{}) edpt.SyslogOperOper {
	count := len(d)
	var ret edpt.SyslogOperOper
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.SystemLog = getSliceSyslogOperOperSystemLog(in["system_log"].([]interface{}))
		ret.NextMsgIdx = in["next_msg_idx"].(int)
	}
	return ret
}

func getSliceSyslogOperOperSystemLog(d []interface{}) []edpt.SyslogOperOperSystemLog {
	count := len(d)
	ret := make([]edpt.SyslogOperOperSystemLog, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SyslogOperOperSystemLog
		oi.LogData = in["log_data"].(string)
		oi.AddSlotInfo = in["add_slot_info"].(int)
		oi.LogDataSearch = in["log_data_search"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSyslogOper(d *schema.ResourceData) edpt.SyslogOper {
	var ret edpt.SyslogOper
	ret.Oper = getObjectSyslogOperOper(d.Get("oper").([]interface{}))
	return ret
}
