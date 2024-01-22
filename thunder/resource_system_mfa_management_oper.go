package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemMfaManagementOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_mfa_management_oper`: Operational Status for the object mfa-management\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemMfaManagementOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemMfaManagementOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMfaManagementOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMfaManagementOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemMfaManagementOperOper := setObjectSystemMfaManagementOperOper(res)
		d.Set("oper", SystemMfaManagementOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemMfaManagementOperOper(ret edpt.DataSystemMfaManagementOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"enable": ret.DtSystemMfaManagementOper.Oper.Enable,
		},
	}
}

func getObjectSystemMfaManagementOperOper(d []interface{}) edpt.SystemMfaManagementOperOper {

	count1 := len(d)
	var ret edpt.SystemMfaManagementOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
	}
	return ret
}

func dataToEndpointSystemMfaManagementOper(d *schema.ResourceData) edpt.SystemMfaManagementOper {
	var ret edpt.SystemMfaManagementOper

	ret.Oper = getObjectSystemMfaManagementOperOper(d.Get("oper").([]interface{}))
	return ret
}
