package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBootimageOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_bootimage_oper`: Operational Status for the object bootimage\n\n__PLACEHOLDER__",
		ReadContext: resourceBootimageOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hd_pri": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"hd_sec": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"cf_pri": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"cf_sec": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"hd_default": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"cf_default": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceBootimageOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBootimageOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBootimageOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		BootimageOperOper := setObjectBootimageOperOper(res)
		d.Set("oper", BootimageOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectBootimageOperOper(ret edpt.DataBootimageOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hd_pri":     ret.DtBootimageOper.Oper.HdPri,
			"hd_sec":     ret.DtBootimageOper.Oper.HdSec,
			"cf_pri":     ret.DtBootimageOper.Oper.CfPri,
			"cf_sec":     ret.DtBootimageOper.Oper.CfSec,
			"hd_default": ret.DtBootimageOper.Oper.HdDefault,
			"cf_default": ret.DtBootimageOper.Oper.CfDefault,
		},
	}
}

func getObjectBootimageOperOper(d []interface{}) edpt.BootimageOperOper {

	count1 := len(d)
	var ret edpt.BootimageOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HdPri = in["hd_pri"].(string)
		ret.HdSec = in["hd_sec"].(string)
		ret.CfPri = in["cf_pri"].(string)
		ret.CfSec = in["cf_sec"].(string)
		ret.HdDefault = in["hd_default"].(string)
		ret.CfDefault = in["cf_default"].(string)
	}
	return ret
}

func dataToEndpointBootimageOper(d *schema.ResourceData) edpt.BootimageOper {
	var ret edpt.BootimageOper

	ret.Oper = getObjectBootimageOperOper(d.Get("oper").([]interface{}))
	return ret
}
