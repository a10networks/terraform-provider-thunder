package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVcsShowdebugOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vcs_showdebug_oper`: Operational Status for the object showdebug\n\n__PLACEHOLDER__",
		ReadContext: resourceVcsShowdebugOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"debugging_switches": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"debugging_buff_size": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceVcsShowdebugOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsShowdebugOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsShowdebugOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VcsShowdebugOperOper := setObjectVcsShowdebugOperOper(res)
		d.Set("oper", VcsShowdebugOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVcsShowdebugOperOper(ret edpt.DataVcsShowdebugOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"debugging_switches":  ret.DtVcsShowdebugOper.Oper.DebuggingSwitches,
			"debugging_buff_size": ret.DtVcsShowdebugOper.Oper.DebuggingBuffSize,
		},
	}
}

func getObjectVcsShowdebugOperOper(d []interface{}) edpt.VcsShowdebugOperOper {

	count1 := len(d)
	var ret edpt.VcsShowdebugOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DebuggingSwitches = in["debugging_switches"].(string)
		ret.DebuggingBuffSize = in["debugging_buff_size"].(int)
	}
	return ret
}

func dataToEndpointVcsShowdebugOper(d *schema.ResourceData) edpt.VcsShowdebugOper {
	var ret edpt.VcsShowdebugOper

	ret.Oper = getObjectVcsShowdebugOperOper(d.Get("oper").([]interface{}))
	return ret
}
