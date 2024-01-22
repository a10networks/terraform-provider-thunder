package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutDebugRoleOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_scaleout_debug_role_oper`: Operational Status for the object role\n\n__PLACEHOLDER__",
		ReadContext: resourceScaleoutDebugRoleOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scaleout_role_action": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceScaleoutDebugRoleOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDebugRoleOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDebugRoleOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ScaleoutDebugRoleOperOper := setObjectScaleoutDebugRoleOperOper(res)
		d.Set("oper", ScaleoutDebugRoleOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectScaleoutDebugRoleOperOper(ret edpt.DataScaleoutDebugRoleOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"scaleout_role_action": ret.DtScaleoutDebugRoleOper.Oper.Scaleout_role_action,
		},
	}
}

func getObjectScaleoutDebugRoleOperOper(d []interface{}) edpt.ScaleoutDebugRoleOperOper {

	count1 := len(d)
	var ret edpt.ScaleoutDebugRoleOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Scaleout_role_action = in["scaleout_role_action"].(string)
	}
	return ret
}

func dataToEndpointScaleoutDebugRoleOper(d *schema.ResourceData) edpt.ScaleoutDebugRoleOper {
	var ret edpt.ScaleoutDebugRoleOper

	ret.Oper = getObjectScaleoutDebugRoleOperOper(d.Get("oper").([]interface{}))
	return ret
}
