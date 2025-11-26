package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceClockSet() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_clock_set`: Set the Time and Date\n\n__PLACEHOLDER__",
		CreateContext: resourceClockSetCreate,
		UpdateContext: resourceClockSetUpdate,
		ReadContext:   resourceClockSetRead,
		DeleteContext: resourceClockSetDelete,

		Schema: map[string]*schema.Schema{
			"time_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"time": {
							Type: schema.TypeString, Optional: true, Description: "Current Time",
						},
						"month": {
							Type: schema.TypeString, Optional: true, Description: "'January': Month of the year; 'February': Month of the year; 'March': Month of the year; 'April': Month of the year; 'May': Month of the year; 'June': Month of the year; 'July': Month of the year; 'August': Month of the year; 'September': Month of the year; 'October': Month of the year; 'November': Month of the year; 'December': Month of the year;",
						},
						"day_of_month_2": {
							Type: schema.TypeInt, Optional: true, Description: "Day of the Month",
						},
						"month_2": {
							Type: schema.TypeString, Optional: true, Description: "'January': Month of the year; 'February': Month of the year; 'March': Month of the year; 'April': Month of the year; 'May': Month of the year; 'June': Month of the year; 'July': Month of the year; 'August': Month of the year; 'September': Month of the year; 'October': Month of the year; 'November': Month of the year; 'December': Month of the year;",
						},
						"day_of_month": {
							Type: schema.TypeInt, Optional: true, Description: "Day of the Month",
						},
						"year": {
							Type: schema.TypeInt, Optional: true, Description: "Year",
						},
					},
				},
			},
		},
	}
}
func resourceClockSetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceClockSetCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointClockSet(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceClockSetRead(ctx, d, meta)
	}
	return diags
}

func resourceClockSetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceClockSetUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointClockSet(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceClockSetRead(ctx, d, meta)
	}
	return diags
}
func resourceClockSetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceClockSetDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointClockSet(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceClockSetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceClockSetRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointClockSet(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectClockSetTimeCfg(d []interface{}) edpt.ClockSetTimeCfg {

	count1 := len(d)
	var ret edpt.ClockSetTimeCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Time = in["time"].(string)
		ret.Month = in["month"].(string)
		ret.DayOfMonth2 = in["day_of_month_2"].(int)
		ret.Month2 = in["month_2"].(string)
		ret.DayOfMonth = in["day_of_month"].(int)
		ret.Year = in["year"].(int)
	}
	return ret
}

func dataToEndpointClockSet(d *schema.ResourceData) edpt.ClockSet {
	var ret edpt.ClockSet
	ret.Inst.TimeCfg = getObjectClockSetTimeCfg(d.Get("time_cfg").([]interface{}))
	return ret
}
