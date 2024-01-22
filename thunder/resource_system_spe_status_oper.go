package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemSpeStatusOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_spe_status_oper`: Operational Status for the object spe-status\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemSpeStatusOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"spe_profile": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"spe_setup_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv4_allowed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipv6_allowed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemSpeStatusOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSpeStatusOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSpeStatusOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemSpeStatusOperOper := setObjectSystemSpeStatusOperOper(res)
		d.Set("oper", SystemSpeStatusOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemSpeStatusOperOper(ret edpt.DataSystemSpeStatusOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"enable":           ret.DtSystemSpeStatusOper.Oper.Enable,
			"spe_profile":      ret.DtSystemSpeStatusOper.Oper.SpeProfile,
			"spe_setup_status": ret.DtSystemSpeStatusOper.Oper.SpeSetupStatus,
			"ipv4_allowed":     ret.DtSystemSpeStatusOper.Oper.Ipv4_allowed,
			"ipv6_allowed":     ret.DtSystemSpeStatusOper.Oper.Ipv6_allowed,
		},
	}
}

func getObjectSystemSpeStatusOperOper(d []interface{}) edpt.SystemSpeStatusOperOper {

	count1 := len(d)
	var ret edpt.SystemSpeStatusOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(string)
		ret.SpeProfile = in["spe_profile"].(string)
		ret.SpeSetupStatus = in["spe_setup_status"].(string)
		ret.Ipv4_allowed = in["ipv4_allowed"].(int)
		ret.Ipv6_allowed = in["ipv6_allowed"].(int)
	}
	return ret
}

func dataToEndpointSystemSpeStatusOper(d *schema.ResourceData) edpt.SystemSpeStatusOper {
	var ret edpt.SystemSpeStatusOper

	ret.Oper = getObjectSystemSpeStatusOperOper(d.Get("oper").([]interface{}))
	return ret
}
