package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemRebootOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_reboot_oper`: Operational Status for the object reboot\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemRebootOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"boot_scheduled": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"boot_now": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"epoch_time": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"hour": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"min": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"uname": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"hostname": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"reason": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemRebootOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemRebootOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemRebootOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemRebootOperOper := setObjectSystemRebootOperOper(res)
		d.Set("oper", SystemRebootOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemRebootOperOper(ret edpt.DataSystemRebootOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"boot_scheduled": ret.DtSystemRebootOper.Oper.BootScheduled,
			"boot_now":       ret.DtSystemRebootOper.Oper.BootNow,
			"epoch_time":     ret.DtSystemRebootOper.Oper.EpochTime,
			"hour":           ret.DtSystemRebootOper.Oper.Hour,
			"min":            ret.DtSystemRebootOper.Oper.Min,
			"uname":          ret.DtSystemRebootOper.Oper.Uname,
			"hostname":       ret.DtSystemRebootOper.Oper.Hostname,
			"reason":         ret.DtSystemRebootOper.Oper.Reason,
		},
	}
}

func getObjectSystemRebootOperOper(d []interface{}) edpt.SystemRebootOperOper {

	count1 := len(d)
	var ret edpt.SystemRebootOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.BootScheduled = in["boot_scheduled"].(int)
		ret.BootNow = in["boot_now"].(int)
		ret.EpochTime = in["epoch_time"].(string)
		ret.Hour = in["hour"].(string)
		ret.Min = in["min"].(string)
		ret.Uname = in["uname"].(string)
		ret.Hostname = in["hostname"].(string)
		ret.Reason = in["reason"].(string)
	}
	return ret
}

func dataToEndpointSystemRebootOper(d *schema.ResourceData) edpt.SystemRebootOper {
	var ret edpt.SystemRebootOper

	ret.Oper = getObjectSystemRebootOperOper(d.Get("oper").([]interface{}))
	return ret
}
