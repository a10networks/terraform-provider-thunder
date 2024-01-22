package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemViewHotfixOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_view_hotfix_oper`: Operational Status for the object hotfix\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemViewHotfixOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"applied": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"copied": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemViewHotfixOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemViewHotfixOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemViewHotfixOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemViewHotfixOperOper := setObjectSystemViewHotfixOperOper(res)
		d.Set("oper", SystemViewHotfixOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemViewHotfixOperOper(ret edpt.DataSystemViewHotfixOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"applied": ret.DtSystemViewHotfixOper.Oper.Applied,
			"copied":  ret.DtSystemViewHotfixOper.Oper.Copied,
		},
	}
}

func getObjectSystemViewHotfixOperOper(d []interface{}) edpt.SystemViewHotfixOperOper {

	count1 := len(d)
	var ret edpt.SystemViewHotfixOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Applied = in["applied"].(string)
		ret.Copied = in["copied"].(string)
	}
	return ret
}

func dataToEndpointSystemViewHotfixOper(d *schema.ResourceData) edpt.SystemViewHotfixOper {
	var ret edpt.SystemViewHotfixOper

	ret.Oper = getObjectSystemViewHotfixOperOper(d.Get("oper").([]interface{}))
	return ret
}
