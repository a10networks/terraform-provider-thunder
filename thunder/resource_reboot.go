package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceReboot() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_reboot`: Reboot the System\n\n__PLACEHOLDER__",
		CreateContext: resourceRebootCreate,
		UpdateContext: resourceRebootUpdate,
		ReadContext:   resourceRebootRead,
		DeleteContext: resourceRebootDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reboot all devices when VCS is enabled, or only this device itself if VCS is not enabled",
			},
			"at": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reboot at a Specific time/date",
			},
			"cancel": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Cancel Pending Reboot",
			},
			"day_of_month": {
				Type: schema.TypeInt, Optional: true, Description: "Day of the Month",
			},
			"day_of_month_2": {
				Type: schema.TypeInt, Optional: true, Description: "Day of the Month",
			},
			"device": {
				Type: schema.TypeInt, Optional: true, Description: "Reboot a specific device when VCS is enabled (device id)",
			},
			"in": {
				Type: schema.TypeString, Optional: true, Description: "Reboot after a time interval (Time in hours and minutes)",
			},
			"month": {
				Type: schema.TypeString, Optional: true, Description: "'January': Month of the year; 'February': Month of the year; 'March': Month of the year; 'April': Month of the year; 'May': Month of the year; 'June': Month of the year; 'July': Month of the year; 'August': Month of the year; 'September': Month of the year; 'October': Month of the year; 'November': Month of the year; 'December': Month of the year;",
			},
			"month_2": {
				Type: schema.TypeString, Optional: true, Description: "'January': Month of the year; 'February': Month of the year; 'March': Month of the year; 'April': Month of the year; 'May': Month of the year; 'June': Month of the year; 'July': Month of the r; 'August': Month of the year; 'September': Month of the year; 'October': Month of the year; 'November': Month of the year; 'December': Month of the year;",
			},
			"reason": {
				Type: schema.TypeString, Optional: true, Description: "Reason for Reboot",
			},
			"reason_2": {
				Type: schema.TypeString, Optional: true, Description: "Reason for Reboot",
			},
			"reason_3": {
				Type: schema.TypeString, Optional: true, Description: "Reason for Reboot",
			},
			"time": {
				Type: schema.TypeString, Optional: true, Description: "Time to Reboot (hh:mm)",
			},
		},
	}
}
func resourceRebootCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRebootCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointReboot(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRebootRead(ctx, d, meta)
	}
	return diags
}

func resourceRebootUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRebootUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointReboot(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRebootRead(ctx, d, meta)
	}
	return diags
}
func resourceRebootDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRebootDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointReboot(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRebootRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRebootRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointReboot(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointReboot(d *schema.ResourceData) edpt.Reboot {
	var ret edpt.Reboot
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.At = d.Get("at").(int)
	ret.Inst.Cancel = d.Get("cancel").(int)
	ret.Inst.DayOfMonth = d.Get("day_of_month").(int)
	ret.Inst.DayOfMonth2 = d.Get("day_of_month_2").(int)
	ret.Inst.Device = d.Get("device").(int)
	ret.Inst.In = d.Get("in").(string)
	ret.Inst.Month = d.Get("month").(string)
	ret.Inst.Month2 = d.Get("month_2").(string)
	ret.Inst.Reason = d.Get("reason").(string)
	ret.Inst.Reason2 = d.Get("reason_2").(string)
	ret.Inst.Reason3 = d.Get("reason_3").(string)
	ret.Inst.Time = d.Get("time").(string)
	return ret
}
