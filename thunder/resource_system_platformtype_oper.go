package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemPlatformtypeOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_platformtype_oper`: Operational Status for the object platformtype\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemPlatformtypeOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"platform_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"platform_info": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"platform_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"platform_axv": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"platform_dpdk": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"platform_lxc": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemPlatformtypeOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPlatformtypeOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPlatformtypeOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemPlatformtypeOperOper := setObjectSystemPlatformtypeOperOper(res)
		d.Set("oper", SystemPlatformtypeOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemPlatformtypeOperOper(ret edpt.DataSystemPlatformtypeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"platform_type": ret.DtSystemPlatformtypeOper.Oper.PlatformType,
			"platform_info": ret.DtSystemPlatformtypeOper.Oper.PlatformInfo,
			"platform_id":   ret.DtSystemPlatformtypeOper.Oper.PlatformId,
			"platform_axv":  ret.DtSystemPlatformtypeOper.Oper.PlatformAxv,
			"platform_dpdk": ret.DtSystemPlatformtypeOper.Oper.PlatformDpdk,
			"platform_lxc":  ret.DtSystemPlatformtypeOper.Oper.PlatformLxc,
		},
	}
}

func getObjectSystemPlatformtypeOperOper(d []interface{}) edpt.SystemPlatformtypeOperOper {

	count1 := len(d)
	var ret edpt.SystemPlatformtypeOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PlatformType = in["platform_type"].(string)
		ret.PlatformInfo = in["platform_info"].(string)
		ret.PlatformId = in["platform_id"].(int)
		ret.PlatformAxv = in["platform_axv"].(int)
		ret.PlatformDpdk = in["platform_dpdk"].(int)
		ret.PlatformLxc = in["platform_lxc"].(int)
	}
	return ret
}

func dataToEndpointSystemPlatformtypeOper(d *schema.ResourceData) edpt.SystemPlatformtypeOper {
	var ret edpt.SystemPlatformtypeOper

	ret.Oper = getObjectSystemPlatformtypeOperOper(d.Get("oper").([]interface{}))
	return ret
}
