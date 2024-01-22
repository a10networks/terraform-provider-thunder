package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemAppsGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_apps_global`: Global application generic configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemAppsGlobalCreate,
		UpdateContext: resourceSystemAppsGlobalUpdate,
		ReadContext:   resourceSystemAppsGlobalRead,
		DeleteContext: resourceSystemAppsGlobalDelete,

		Schema: map[string]*schema.Schema{
			"cps_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set threshold for the total Connections Per Second across the system (Enter threshold number)",
			},
			"log_session_on_established": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send TCP session creation log on completion of 3-way handshake",
			},
			"msl_time": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "Configure maximum session life, default is 2 seconds (1-39 seconds, default is 2 seconds)",
			},
			"sessions_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set threshold for total sessions across the system (Enter threshold number)",
			},
			"timer_wheel_walk_limit": {
				Type: schema.TypeInt, Optional: true, Default: 100, Description: "Set timer wheel walk limit (0-1024, 0 is unlimited, default is 100)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemAppsGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAppsGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAppsGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemAppsGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemAppsGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAppsGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAppsGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemAppsGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemAppsGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAppsGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAppsGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemAppsGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAppsGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAppsGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemAppsGlobal(d *schema.ResourceData) edpt.SystemAppsGlobal {
	var ret edpt.SystemAppsGlobal
	ret.Inst.CpsThreshold = d.Get("cps_threshold").(int)
	ret.Inst.LogSessionOnEstablished = d.Get("log_session_on_established").(int)
	ret.Inst.MslTime = d.Get("msl_time").(int)
	ret.Inst.SessionsThreshold = d.Get("sessions_threshold").(int)
	ret.Inst.TimerWheelWalkLimit = d.Get("timer_wheel_walk_limit").(int)
	//omit uuid
	return ret
}
