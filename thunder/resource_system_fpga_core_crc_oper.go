package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemFpgaCoreCrcOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_fpga_core_crc_oper`: Operational Status for the object fpga-core-crc\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemFpgaCoreCrcOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_val": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"reboot_val": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"crc_error_present": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemFpgaCoreCrcOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemFpgaCoreCrcOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemFpgaCoreCrcOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemFpgaCoreCrcOperOper := setObjectSystemFpgaCoreCrcOperOper(res)
		d.Set("oper", SystemFpgaCoreCrcOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemFpgaCoreCrcOperOper(ret edpt.DataSystemFpgaCoreCrcOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"enable_val":        ret.DtSystemFpgaCoreCrcOper.Oper.Enable_val,
			"reboot_val":        ret.DtSystemFpgaCoreCrcOper.Oper.Reboot_val,
			"crc_error_present": ret.DtSystemFpgaCoreCrcOper.Oper.Crc_error_present,
		},
	}
}

func getObjectSystemFpgaCoreCrcOperOper(d []interface{}) edpt.SystemFpgaCoreCrcOperOper {

	count1 := len(d)
	var ret edpt.SystemFpgaCoreCrcOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable_val = in["enable_val"].(string)
		ret.Reboot_val = in["reboot_val"].(string)
		ret.Crc_error_present = in["crc_error_present"].(string)
	}
	return ret
}

func dataToEndpointSystemFpgaCoreCrcOper(d *schema.ResourceData) edpt.SystemFpgaCoreCrcOper {
	var ret edpt.SystemFpgaCoreCrcOper

	ret.Oper = getObjectSystemFpgaCoreCrcOperOper(d.Get("oper").([]interface{}))
	return ret
}
