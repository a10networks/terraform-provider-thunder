package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAutomaticUpdateConfig() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_automatic_update_config`: Configure software update schedule\n\n__PLACEHOLDER__",
		CreateContext: resourceAutomaticUpdateConfigCreate,
		UpdateContext: resourceAutomaticUpdateConfigUpdate,
		ReadContext:   resourceAutomaticUpdateConfigRead,
		DeleteContext: resourceAutomaticUpdateConfigDelete,

		Schema: map[string]*schema.Schema{
			"daily": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Every day",
			},
			"day_time": {
				Type: schema.TypeString, Optional: true, Description: "Time of day to update (hh:mm) in 24 hour local time",
			},
			"debug": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable libcurl debug option",
			},
			"disable_ssl_verify": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable peer server certificate verification",
			},
			"feature_name": {
				Type: schema.TypeString, Required: true, Description: "'app-fw': Application Firewall Configuration; 'ca-bundle': CA Certificate Bundle; 'a10-threat-intel': A10 Threat intel class list; 'central-cert-pin-list': Central updated cert pinning list;",
			},
			"schedule": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"week_day": {
				Type: schema.TypeString, Optional: true, Description: "'Monday': Monday; 'Tuesday': Tuesday; 'Wednesday': Wednesday; 'Thursday': Thursday; 'Friday': Friday; 'Saturday': Saturday; 'Sunday': Sunday;",
			},
			"week_time": {
				Type: schema.TypeString, Optional: true, Description: "Time of day to update (hh:mm) in 24 hour local time",
			},
			"weekly": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Every week",
			},
		},
	}
}
func resourceAutomaticUpdateConfigCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateConfigCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdateConfig(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAutomaticUpdateConfigRead(ctx, d, meta)
	}
	return diags
}

func resourceAutomaticUpdateConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateConfigUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdateConfig(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAutomaticUpdateConfigRead(ctx, d, meta)
	}
	return diags
}
func resourceAutomaticUpdateConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateConfigDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdateConfig(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAutomaticUpdateConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateConfigRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdateConfig(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAutomaticUpdateConfig(d *schema.ResourceData) edpt.AutomaticUpdateConfig {
	var ret edpt.AutomaticUpdateConfig
	ret.Inst.Daily = d.Get("daily").(int)
	ret.Inst.DayTime = d.Get("day_time").(string)
	ret.Inst.Debug = d.Get("debug").(int)
	ret.Inst.DisableSslVerify = d.Get("disable_ssl_verify").(int)
	ret.Inst.FeatureName = d.Get("feature_name").(string)
	ret.Inst.Schedule = d.Get("schedule").(int)
	//omit uuid
	ret.Inst.WeekDay = d.Get("week_day").(string)
	ret.Inst.WeekTime = d.Get("week_time").(string)
	ret.Inst.Weekly = d.Get("weekly").(int)
	return ret
}
