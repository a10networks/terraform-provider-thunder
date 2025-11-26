package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceShutdown() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_shutdown`: Shutdown the System\n\n__PLACEHOLDER__",
		CreateContext: resourceShutdownCreate,
		UpdateContext: resourceShutdownUpdate,
		ReadContext:   resourceShutdownRead,
		DeleteContext: resourceShutdownDelete,

		Schema: map[string]*schema.Schema{
			"at": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Shutdown at a specific time/date",
			},
			"cancel": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Cancel Pending Shutdown",
			},
			"day_of_month": {
				Type: schema.TypeInt, Optional: true, Description: "Day of the Month",
			},
			"day_of_month_2": {
				Type: schema.TypeInt, Optional: true, Description: "Day of the Month",
			},
			"in": {
				Type: schema.TypeString, Optional: true, Description: "Delay before Shutdown (Time in hours and minutes)",
			},
			"month": {
				Type: schema.TypeString, Optional: true, Description: "'January': Month of the year; 'February': Month of the year; 'March': Month of the year; 'April': Month of the year; 'May': Month of the year; 'June': Month of the year; 'July': Month of the year; 'August': Month of the year; 'September': Month of the year; 'October': Month of the year; 'November': Month of the year; 'December': Month of the year;",
			},
			"month_2": {
				Type: schema.TypeString, Optional: true, Description: "'January': Month of the year; 'February': Month of the year; 'March': Month of the year; 'April': Month of the year; 'May': Month of the year; 'June': Month of the year; 'July': Month of the year; 'August': Month of the year; 'September': Month of the year; 'October': Month of the year; 'November': Month of the year; 'December': Month of the year;",
			},
			"reason": {
				Type: schema.TypeString, Optional: true, Description: "Reason for Shutdown",
			},
			"reason_2": {
				Type: schema.TypeString, Optional: true, Description: "Reason for Shutdown",
			},
			"reason_3": {
				Type: schema.TypeString, Optional: true, Description: "Reason for Shutdown",
			},
			"time": {
				Type: schema.TypeString, Optional: true, Description: "Time to Shutdown (hh:mm)",
			},
		},
	}
}
func resourceShutdownCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceShutdownCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointShutdown(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceShutdownRead(ctx, d, meta)
	}
	return diags
}

func resourceShutdownUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceShutdownUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointShutdown(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceShutdownRead(ctx, d, meta)
	}
	return diags
}
func resourceShutdownDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceShutdownDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointShutdown(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceShutdownRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceShutdownRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointShutdown(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointShutdown(d *schema.ResourceData) edpt.Shutdown {
	var ret edpt.Shutdown
	ret.Inst.At = d.Get("at").(int)
	ret.Inst.Cancel = d.Get("cancel").(int)
	ret.Inst.DayOfMonth = d.Get("day_of_month").(int)
	ret.Inst.DayOfMonth2 = d.Get("day_of_month_2").(int)
	ret.Inst.In = d.Get("in").(string)
	ret.Inst.Month = d.Get("month").(string)
	ret.Inst.Month2 = d.Get("month_2").(string)
	ret.Inst.Reason = d.Get("reason").(string)
	ret.Inst.Reason2 = d.Get("reason_2").(string)
	ret.Inst.Reason3 = d.Get("reason_3").(string)
	ret.Inst.Time = d.Get("time").(string)
	return ret
}
