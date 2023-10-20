package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceClockShowOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_clock_show_oper`: Operational Status for the object show\n\n__PLACEHOLDER__",
		ReadContext: resourceClockShowOperRead,
		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_type": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"time": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"timezone": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"offset": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"day": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"date": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceClockShowOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceClockShowOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointClockShowOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		opers := setObjectClockShowOperOper(res)
		d.SetId(obj.GetId())
		d.Set("oper", opers)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectClockShowOperOper(res edpt.ClockShoww) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["source_type"] = res.DataClockShow.Oper.SourceType
	in["time"] = res.DataClockShow.Oper.Time
	in["timezone"] = res.DataClockShow.Oper.Timezone
	in["offset"] = res.DataClockShow.Oper.Offset
	in["day"] = res.DataClockShow.Oper.Day
	in["date"] = res.DataClockShow.Oper.Date
	result = append(result, in)
	return result
}

func getObjectClockShowOperOper(d []interface{}) edpt.ClockShowOperOper {
	count := len(d)
	var ret edpt.ClockShowOperOper
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.SourceType = in["source_type"].(int)
		ret.Time = in["time"].(string)
		ret.Timezone = in["timezone"].(string)
		ret.Offset = in["offset"].(string)
		ret.Day = in["day"].(string)
		ret.Date = in["date"].(string)
	}
	return ret
}

func dataToEndpointClockShowOper(d *schema.ResourceData) edpt.ClockShowOper {
	var ret edpt.ClockShowOper
	ret.Oper = getObjectClockShowOperOper(d.Get("oper").([]interface{}))
	return ret
}
