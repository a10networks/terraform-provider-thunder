package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMiscelleniousAlbOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_miscellenious_alb_oper`: Operational Status for the object miscellenious-alb\n\n__PLACEHOLDER__",
		ReadContext: resourceMiscelleniousAlbOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uptime": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"reboot_num": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"crash_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceMiscelleniousAlbOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceMiscelleniousAlbOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointMiscelleniousAlbOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		MiscelleniousAlbOperOper := setObjectMiscelleniousAlbOperOper(res)
		d.Set("oper", MiscelleniousAlbOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectMiscelleniousAlbOperOper(ret edpt.DataMiscelleniousAlbOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"uptime":      ret.DtMiscelleniousAlbOper.Oper.Uptime,
			"reboot_num":  ret.DtMiscelleniousAlbOper.Oper.RebootNum,
			"crash_count": ret.DtMiscelleniousAlbOper.Oper.CrashCount,
		},
	}
}

func getObjectMiscelleniousAlbOperOper(d []interface{}) edpt.MiscelleniousAlbOperOper {

	count1 := len(d)
	var ret edpt.MiscelleniousAlbOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Uptime = in["uptime"].(int)
		ret.RebootNum = in["reboot_num"].(int)
		ret.CrashCount = in["crash_count"].(int)
	}
	return ret
}

func dataToEndpointMiscelleniousAlbOper(d *schema.ResourceData) edpt.MiscelleniousAlbOper {
	var ret edpt.MiscelleniousAlbOper

	ret.Oper = getObjectMiscelleniousAlbOperOper(d.Get("oper").([]interface{}))
	return ret
}
