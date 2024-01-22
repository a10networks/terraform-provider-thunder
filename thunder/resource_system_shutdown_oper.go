package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemShutdownOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_shutdown_oper`: Operational Status for the object shutdown\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemShutdownOperRead,

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

func resourceSystemShutdownOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemShutdownOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemShutdownOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemShutdownOperOper := setObjectSystemShutdownOperOper(res)
		d.Set("oper", SystemShutdownOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemShutdownOperOper(ret edpt.DataSystemShutdownOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"boot_scheduled": ret.DtSystemShutdownOper.Oper.BootScheduled,
			"boot_now":       ret.DtSystemShutdownOper.Oper.BootNow,
			"epoch_time":     ret.DtSystemShutdownOper.Oper.EpochTime,
			"hour":           ret.DtSystemShutdownOper.Oper.Hour,
			"min":            ret.DtSystemShutdownOper.Oper.Min,
			"uname":          ret.DtSystemShutdownOper.Oper.Uname,
			"hostname":       ret.DtSystemShutdownOper.Oper.Hostname,
			"reason":         ret.DtSystemShutdownOper.Oper.Reason,
		},
	}
}

func getObjectSystemShutdownOperOper(d []interface{}) edpt.SystemShutdownOperOper {

	count1 := len(d)
	var ret edpt.SystemShutdownOperOper
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

func dataToEndpointSystemShutdownOper(d *schema.ResourceData) edpt.SystemShutdownOper {
	var ret edpt.SystemShutdownOper

	ret.Oper = getObjectSystemShutdownOperOper(d.Get("oper").([]interface{}))
	return ret
}
