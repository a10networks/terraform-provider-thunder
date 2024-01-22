package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFailSafeConfigOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fail_safe_config_oper`: Operational Status for the object config\n\n__PLACEHOLDER__",
		ReadContext: resourceFailSafeConfigOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sw_error_mon": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"hw_error_mon": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"sw_recovery_timeout": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"hw_recovery_timeout": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dataplane_recovery_timeout": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fpga_mon_enable": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fpga_mon_forced_reboot": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fpga_mon_interval": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fpga_mon_threshold": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"mem_mon": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceFailSafeConfigOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFailSafeConfigOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFailSafeConfigOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FailSafeConfigOperOper := setObjectFailSafeConfigOperOper(res)
		d.Set("oper", FailSafeConfigOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFailSafeConfigOperOper(ret edpt.DataFailSafeConfigOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"sw_error_mon":               ret.DtFailSafeConfigOper.Oper.Sw_error_mon,
			"hw_error_mon":               ret.DtFailSafeConfigOper.Oper.Hw_error_mon,
			"sw_recovery_timeout":        ret.DtFailSafeConfigOper.Oper.Sw_recovery_timeout,
			"hw_recovery_timeout":        ret.DtFailSafeConfigOper.Oper.Hw_recovery_timeout,
			"dataplane_recovery_timeout": ret.DtFailSafeConfigOper.Oper.Dataplane_recovery_timeout,
			"fpga_mon_enable":            ret.DtFailSafeConfigOper.Oper.Fpga_mon_enable,
			"fpga_mon_forced_reboot":     ret.DtFailSafeConfigOper.Oper.Fpga_mon_forced_reboot,
			"fpga_mon_interval":          ret.DtFailSafeConfigOper.Oper.Fpga_mon_interval,
			"fpga_mon_threshold":         ret.DtFailSafeConfigOper.Oper.Fpga_mon_threshold,
			"mem_mon":                    ret.DtFailSafeConfigOper.Oper.Mem_mon,
		},
	}
}

func getObjectFailSafeConfigOperOper(d []interface{}) edpt.FailSafeConfigOperOper {

	count1 := len(d)
	var ret edpt.FailSafeConfigOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Sw_error_mon = in["sw_error_mon"].(string)
		ret.Hw_error_mon = in["hw_error_mon"].(string)
		ret.Sw_recovery_timeout = in["sw_recovery_timeout"].(string)
		ret.Hw_recovery_timeout = in["hw_recovery_timeout"].(string)
		ret.Dataplane_recovery_timeout = in["dataplane_recovery_timeout"].(string)
		ret.Fpga_mon_enable = in["fpga_mon_enable"].(string)
		ret.Fpga_mon_forced_reboot = in["fpga_mon_forced_reboot"].(string)
		ret.Fpga_mon_interval = in["fpga_mon_interval"].(string)
		ret.Fpga_mon_threshold = in["fpga_mon_threshold"].(string)
		ret.Mem_mon = in["mem_mon"].(string)
	}
	return ret
}

func dataToEndpointFailSafeConfigOper(d *schema.ResourceData) edpt.FailSafeConfigOper {
	var ret edpt.FailSafeConfigOper

	ret.Oper = getObjectFailSafeConfigOperOper(d.Get("oper").([]interface{}))
	return ret
}
