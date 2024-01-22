package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemViewShowMonitorOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_view_show_monitor_oper`: Operational Status for the object show-monitor\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemViewShowMonitorOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disk_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"mem_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ctrl_cpu": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"data_cpu": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"buff_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"buff_drop": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"warn_temp": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"spm0": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"spm1": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"spm2": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"spm3": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"spm4": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"smp0": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"smp1": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"smp2": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"smp3": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"smp4": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemViewShowMonitorOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemViewShowMonitorOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemViewShowMonitorOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemViewShowMonitorOperOper := setObjectSystemViewShowMonitorOperOper(res)
		d.Set("oper", SystemViewShowMonitorOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemViewShowMonitorOperOper(ret edpt.DataSystemViewShowMonitorOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"disk_value": ret.DtSystemViewShowMonitorOper.Oper.DiskValue,
			"mem_value":  ret.DtSystemViewShowMonitorOper.Oper.MemValue,
			"ctrl_cpu":   ret.DtSystemViewShowMonitorOper.Oper.CtrlCpu,
			"data_cpu":   ret.DtSystemViewShowMonitorOper.Oper.DataCpu,
			"buff_value": ret.DtSystemViewShowMonitorOper.Oper.BuffValue,
			"buff_drop":  ret.DtSystemViewShowMonitorOper.Oper.BuffDrop,
			"warn_temp":  ret.DtSystemViewShowMonitorOper.Oper.WarnTemp,
			"spm0":       ret.DtSystemViewShowMonitorOper.Oper.Spm0,
			"spm1":       ret.DtSystemViewShowMonitorOper.Oper.Spm1,
			"spm2":       ret.DtSystemViewShowMonitorOper.Oper.Spm2,
			"spm3":       ret.DtSystemViewShowMonitorOper.Oper.Spm3,
			"spm4":       ret.DtSystemViewShowMonitorOper.Oper.Spm4,
			"smp0":       ret.DtSystemViewShowMonitorOper.Oper.Smp0,
			"smp1":       ret.DtSystemViewShowMonitorOper.Oper.Smp1,
			"smp2":       ret.DtSystemViewShowMonitorOper.Oper.Smp2,
			"smp3":       ret.DtSystemViewShowMonitorOper.Oper.Smp3,
			"smp4":       ret.DtSystemViewShowMonitorOper.Oper.Smp4,
		},
	}
}

func getObjectSystemViewShowMonitorOperOper(d []interface{}) edpt.SystemViewShowMonitorOperOper {

	count1 := len(d)
	var ret edpt.SystemViewShowMonitorOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DiskValue = in["disk_value"].(int)
		ret.MemValue = in["mem_value"].(int)
		ret.CtrlCpu = in["ctrl_cpu"].(int)
		ret.DataCpu = in["data_cpu"].(int)
		ret.BuffValue = in["buff_value"].(int)
		ret.BuffDrop = in["buff_drop"].(int)
		ret.WarnTemp = in["warn_temp"].(int)
		ret.Spm0 = in["spm0"].(int)
		ret.Spm1 = in["spm1"].(int)
		ret.Spm2 = in["spm2"].(int)
		ret.Spm3 = in["spm3"].(int)
		ret.Spm4 = in["spm4"].(int)
		ret.Smp0 = in["smp0"].(int)
		ret.Smp1 = in["smp1"].(int)
		ret.Smp2 = in["smp2"].(int)
		ret.Smp3 = in["smp3"].(int)
		ret.Smp4 = in["smp4"].(int)
	}
	return ret
}

func dataToEndpointSystemViewShowMonitorOper(d *schema.ResourceData) edpt.SystemViewShowMonitorOper {
	var ret edpt.SystemViewShowMonitorOper

	ret.Oper = getObjectSystemViewShowMonitorOperOper(d.Get("oper").([]interface{}))
	return ret
}
