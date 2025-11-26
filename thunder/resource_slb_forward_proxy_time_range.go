package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbForwardProxyTimeRange() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_forward_proxy_time_range`: time-range condition for forward-proxy\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbForwardProxyTimeRangeCreate,
		UpdateContext: resourceSlbForwardProxyTimeRangeUpdate,
		ReadContext:   resourceSlbForwardProxyTimeRangeRead,
		DeleteContext: resourceSlbForwardProxyTimeRangeDelete,

		Schema: map[string]*schema.Schema{
			"daily_begin": {
				Type: schema.TypeString, Optional: true, Description: "Daily enables after this time (HH:MM, 24-hour notation)",
			},
			"daily_end": {
				Type: schema.TypeString, Optional: true, Description: "Daily enables until thie time (HH:MM, 24-hour notation, inclusive)",
			},
			"days_of_month_begin": {
				Type: schema.TypeInt, Optional: true, Description: "Recurring enables after this day per month",
			},
			"days_of_month_end": {
				Type: schema.TypeInt, Optional: true, Description: "Recurring enables until this day per month (inclusive)",
			},
			"months_of_year_begin": {
				Type: schema.TypeInt, Optional: true, Description: "Recurring enables after this month per year",
			},
			"months_of_year_end": {
				Type: schema.TypeInt, Optional: true, Description: "Recurring enables until this month per year (inclusive)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name of this object",
			},
			"use_utc_time": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use UTC+0 time",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"weekday_fr": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Recurring enables on Friday",
			},
			"weekday_mo": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Recurring enables on Monday",
			},
			"weekday_sa": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Recurring enables on Saturday",
			},
			"weekday_su": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Recurring enables on Sunday",
			},
			"weekday_th": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Recurring enables on Thursday",
			},
			"weekday_tu": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Recurring enables on Tuesday",
			},
			"weekday_we": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Recurring enables on Wednesday",
			},
		},
	}
}
func resourceSlbForwardProxyTimeRangeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbForwardProxyTimeRangeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbForwardProxyTimeRange(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbForwardProxyTimeRangeRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbForwardProxyTimeRangeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbForwardProxyTimeRangeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbForwardProxyTimeRange(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbForwardProxyTimeRangeRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbForwardProxyTimeRangeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbForwardProxyTimeRangeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbForwardProxyTimeRange(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbForwardProxyTimeRangeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbForwardProxyTimeRangeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbForwardProxyTimeRange(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbForwardProxyTimeRange(d *schema.ResourceData) edpt.SlbForwardProxyTimeRange {
	var ret edpt.SlbForwardProxyTimeRange
	ret.Inst.DailyBegin = d.Get("daily_begin").(string)
	ret.Inst.DailyEnd = d.Get("daily_end").(string)
	ret.Inst.DaysOfMonthBegin = d.Get("days_of_month_begin").(int)
	ret.Inst.DaysOfMonthEnd = d.Get("days_of_month_end").(int)
	ret.Inst.MonthsOfYearBegin = d.Get("months_of_year_begin").(int)
	ret.Inst.MonthsOfYearEnd = d.Get("months_of_year_end").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UseUtcTime = d.Get("use_utc_time").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.WeekdayFr = d.Get("weekday_fr").(int)
	ret.Inst.WeekdayMo = d.Get("weekday_mo").(int)
	ret.Inst.WeekdaySa = d.Get("weekday_sa").(int)
	ret.Inst.WeekdaySu = d.Get("weekday_su").(int)
	ret.Inst.WeekdayTh = d.Get("weekday_th").(int)
	ret.Inst.WeekdayTu = d.Get("weekday_tu").(int)
	ret.Inst.WeekdayWe = d.Get("weekday_we").(int)
	return ret
}
