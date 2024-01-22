package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwStatusOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_status_oper`: Operational Status for the object status\n\n__PLACEHOLDER__",
		ReadContext: resourceFwStatusOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"current_active_rule_set": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"previous_successful_compilation_attempt": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"previous_successful_compilation_duration": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"most_recent_compilation_attempt": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"most_recent_compilation_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"internal": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceFwStatusOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwStatusOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwStatusOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwStatusOperOper := setObjectFwStatusOperOper(res)
		d.Set("oper", FwStatusOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwStatusOperOper(ret edpt.DataFwStatusOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"current_active_rule_set":                  ret.DtFwStatusOper.Oper.CurrentActiveRuleSet,
			"previous_successful_compilation_attempt":  ret.DtFwStatusOper.Oper.PreviousSuccessfulCompilationAttempt,
			"previous_successful_compilation_duration": ret.DtFwStatusOper.Oper.PreviousSuccessfulCompilationDuration,
			"most_recent_compilation_attempt":          ret.DtFwStatusOper.Oper.MostRecentCompilationAttempt,
			"most_recent_compilation_status":           ret.DtFwStatusOper.Oper.MostRecentCompilationStatus,
			"internal":                                 ret.DtFwStatusOper.Oper.Internal,
		},
	}
}

func getObjectFwStatusOperOper(d []interface{}) edpt.FwStatusOperOper {

	count1 := len(d)
	var ret edpt.FwStatusOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CurrentActiveRuleSet = in["current_active_rule_set"].(string)
		ret.PreviousSuccessfulCompilationAttempt = in["previous_successful_compilation_attempt"].(string)
		ret.PreviousSuccessfulCompilationDuration = in["previous_successful_compilation_duration"].(string)
		ret.MostRecentCompilationAttempt = in["most_recent_compilation_attempt"].(string)
		ret.MostRecentCompilationStatus = in["most_recent_compilation_status"].(string)
		ret.Internal = in["internal"].(int)
	}
	return ret
}

func dataToEndpointFwStatusOper(d *schema.ResourceData) edpt.FwStatusOper {
	var ret edpt.FwStatusOper

	ret.Oper = getObjectFwStatusOperOper(d.Get("oper").([]interface{}))
	return ret
}
