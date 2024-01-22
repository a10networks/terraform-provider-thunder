package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwSystemStatusOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_system_status_oper`: Operational Status for the object system-status\n\n__PLACEHOLDER__",
		ReadContext: resourceFwSystemStatusOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_sessions_used": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"data_sessions_free": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"smp_sessions_used": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"smp_sessions_free": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"radius_entries_used": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"radius_entries_free": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceFwSystemStatusOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwSystemStatusOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwSystemStatusOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwSystemStatusOperOper := setObjectFwSystemStatusOperOper(res)
		d.Set("oper", FwSystemStatusOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwSystemStatusOperOper(ret edpt.DataFwSystemStatusOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"data_sessions_used":  ret.DtFwSystemStatusOper.Oper.DataSessionsUsed,
			"data_sessions_free":  ret.DtFwSystemStatusOper.Oper.DataSessionsFree,
			"smp_sessions_used":   ret.DtFwSystemStatusOper.Oper.SmpSessionsUsed,
			"smp_sessions_free":   ret.DtFwSystemStatusOper.Oper.SmpSessionsFree,
			"radius_entries_used": ret.DtFwSystemStatusOper.Oper.RadiusEntriesUsed,
			"radius_entries_free": ret.DtFwSystemStatusOper.Oper.RadiusEntriesFree,
		},
	}
}

func getObjectFwSystemStatusOperOper(d []interface{}) edpt.FwSystemStatusOperOper {

	count1 := len(d)
	var ret edpt.FwSystemStatusOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DataSessionsUsed = in["data_sessions_used"].(int)
		ret.DataSessionsFree = in["data_sessions_free"].(int)
		ret.SmpSessionsUsed = in["smp_sessions_used"].(int)
		ret.SmpSessionsFree = in["smp_sessions_free"].(int)
		ret.RadiusEntriesUsed = in["radius_entries_used"].(int)
		ret.RadiusEntriesFree = in["radius_entries_free"].(int)
	}
	return ret
}

func dataToEndpointFwSystemStatusOper(d *schema.ResourceData) edpt.FwSystemStatusOper {
	var ret edpt.FwSystemStatusOper

	ret.Oper = getObjectFwSystemStatusOperOper(d.Get("oper").([]interface{}))
	return ret
}
