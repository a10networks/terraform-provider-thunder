package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemDomainListSettingsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_domain_list_settings_oper`: Operational Status for the object domain-list-settings\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemDomainListSettingsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total_entry_num": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"inited_entry_num": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"in_prog_entry_num": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"exception_entry_num": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"serial_updated": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"xfr_start_fail": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"xfr_timeout_fail": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"xfr_parse_fail": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemDomainListSettingsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDomainListSettingsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDomainListSettingsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemDomainListSettingsOperOper := setObjectSystemDomainListSettingsOperOper(res)
		d.Set("oper", SystemDomainListSettingsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemDomainListSettingsOperOper(ret edpt.DataSystemDomainListSettingsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total_entry_num":     ret.DtSystemDomainListSettingsOper.Oper.TotalEntryNum,
			"inited_entry_num":    ret.DtSystemDomainListSettingsOper.Oper.InitedEntryNum,
			"in_prog_entry_num":   ret.DtSystemDomainListSettingsOper.Oper.InProgEntryNum,
			"exception_entry_num": ret.DtSystemDomainListSettingsOper.Oper.ExceptionEntryNum,
			"serial_updated":      ret.DtSystemDomainListSettingsOper.Oper.SerialUpdated,
			"xfr_start_fail":      ret.DtSystemDomainListSettingsOper.Oper.XfrStartFail,
			"xfr_timeout_fail":    ret.DtSystemDomainListSettingsOper.Oper.XfrTimeoutFail,
			"xfr_parse_fail":      ret.DtSystemDomainListSettingsOper.Oper.XfrParseFail,
		},
	}
}

func getObjectSystemDomainListSettingsOperOper(d []interface{}) edpt.SystemDomainListSettingsOperOper {

	count1 := len(d)
	var ret edpt.SystemDomainListSettingsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TotalEntryNum = in["total_entry_num"].(int)
		ret.InitedEntryNum = in["inited_entry_num"].(int)
		ret.InProgEntryNum = in["in_prog_entry_num"].(int)
		ret.ExceptionEntryNum = in["exception_entry_num"].(int)
		ret.SerialUpdated = in["serial_updated"].(int)
		ret.XfrStartFail = in["xfr_start_fail"].(int)
		ret.XfrTimeoutFail = in["xfr_timeout_fail"].(int)
		ret.XfrParseFail = in["xfr_parse_fail"].(int)
	}
	return ret
}

func dataToEndpointSystemDomainListSettingsOper(d *schema.ResourceData) edpt.SystemDomainListSettingsOper {
	var ret edpt.SystemDomainListSettingsOper

	ret.Oper = getObjectSystemDomainListSettingsOperOper(d.Get("oper").([]interface{}))
	return ret
}
