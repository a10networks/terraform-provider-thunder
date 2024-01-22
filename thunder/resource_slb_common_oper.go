package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbCommonOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_common_oper`: Operational Status for the object common\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbCommonOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server_auto_reselect": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbCommonOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbCommonOperOper := setObjectSlbCommonOperOper(res)
		d.Set("oper", SlbCommonOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbCommonOperOper(ret edpt.DataSlbCommonOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"server_auto_reselect": ret.DtSlbCommonOper.Oper.ServerAutoReselect,
		},
	}
}

func getObjectSlbCommonOperOper(d []interface{}) edpt.SlbCommonOperOper {

	count1 := len(d)
	var ret edpt.SlbCommonOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ServerAutoReselect = in["server_auto_reselect"].(int)
	}
	return ret
}

func dataToEndpointSlbCommonOper(d *schema.ResourceData) edpt.SlbCommonOper {
	var ret edpt.SlbCommonOper

	ret.Oper = getObjectSlbCommonOperOper(d.Get("oper").([]interface{}))
	return ret
}
