package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysAuditLogOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_sys_audit_log_oper`: Operational Status for the object sys-audit-log\n\n__PLACEHOLDER__",
		ReadContext: resourceSysAuditLogOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"system_audit_log": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"log_audit_data": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"partitions": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"log_audit_search": {
										Type: schema.TypeString, Optional: true, Description: "",
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

func resourceSysAuditLogOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysAuditLogOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysAuditLogOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SysAuditLogOperOper := setObjectSysAuditLogOperOper(res)
		d.Set("oper", SysAuditLogOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSysAuditLogOperOper(ret edpt.DataSysAuditLogOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"system_audit_log": setSliceSysAuditLogOperOperSystemAuditLog(ret.DtSysAuditLogOper.Oper.SystemAuditLog),
		},
	}
}

func setSliceSysAuditLogOperOperSystemAuditLog(d []edpt.SysAuditLogOperOperSystemAuditLog) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["log_audit_data"] = item.LogAuditData
		in["partitions"] = item.Partitions
		in["log_audit_search"] = item.LogAuditSearch
		result = append(result, in)
	}
	return result
}

func getObjectSysAuditLogOperOper(d []interface{}) edpt.SysAuditLogOperOper {

	count1 := len(d)
	var ret edpt.SysAuditLogOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SystemAuditLog = getSliceSysAuditLogOperOperSystemAuditLog(in["system_audit_log"].([]interface{}))
	}
	return ret
}

func getSliceSysAuditLogOperOperSystemAuditLog(d []interface{}) []edpt.SysAuditLogOperOperSystemAuditLog {

	count1 := len(d)
	ret := make([]edpt.SysAuditLogOperOperSystemAuditLog, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysAuditLogOperOperSystemAuditLog
		oi.LogAuditData = in["log_audit_data"].(string)
		oi.Partitions = in["partitions"].(string)
		oi.LogAuditSearch = in["log_audit_search"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSysAuditLogOper(d *schema.ResourceData) edpt.SysAuditLogOper {
	var ret edpt.SysAuditLogOper

	ret.Oper = getObjectSysAuditLogOperOper(d.Get("oper").([]interface{}))
	return ret
}
