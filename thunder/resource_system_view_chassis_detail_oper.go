package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemViewChassisDetailOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_view_chassis_detail_oper`: Operational Status for the object chassis-detail\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemViewChassisDetailOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"platform_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"single_brd_mode_forced": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"single_brd_mode_fallback": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"is_now_single_board": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemViewChassisDetailOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemViewChassisDetailOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemViewChassisDetailOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemViewChassisDetailOperOper := setObjectSystemViewChassisDetailOperOper(res)
		d.Set("oper", SystemViewChassisDetailOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemViewChassisDetailOperOper(ret edpt.DataSystemViewChassisDetailOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"platform_name":            ret.DtSystemViewChassisDetailOper.Oper.PlatformName,
			"single_brd_mode_forced":   ret.DtSystemViewChassisDetailOper.Oper.SingleBrdModeForced,
			"single_brd_mode_fallback": ret.DtSystemViewChassisDetailOper.Oper.SingleBrdModeFallback,
			"is_now_single_board":      ret.DtSystemViewChassisDetailOper.Oper.Is_now_single_board,
		},
	}
}

func getObjectSystemViewChassisDetailOperOper(d []interface{}) edpt.SystemViewChassisDetailOperOper {

	count1 := len(d)
	var ret edpt.SystemViewChassisDetailOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PlatformName = in["platform_name"].(string)
		ret.SingleBrdModeForced = in["single_brd_mode_forced"].(string)
		ret.SingleBrdModeFallback = in["single_brd_mode_fallback"].(string)
		ret.Is_now_single_board = in["is_now_single_board"].(string)
	}
	return ret
}

func dataToEndpointSystemViewChassisDetailOper(d *schema.ResourceData) edpt.SystemViewChassisDetailOper {
	var ret edpt.SystemViewChassisDetailOper

	ret.Oper = getObjectSystemViewChassisDetailOperOper(d.Get("oper").([]interface{}))
	return ret
}
