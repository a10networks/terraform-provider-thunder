package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerEnableTrapsSystemAppsGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_enable_traps_system_apps_global`: Enable apps-global group traps\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerEnableTrapsSystemAppsGlobalCreate,
		UpdateContext: resourceSnmpServerEnableTrapsSystemAppsGlobalUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsSystemAppsGlobalRead,
		DeleteContext: resourceSnmpServerEnableTrapsSystemAppsGlobalDelete,

		Schema: map[string]*schema.Schema{
			"cps_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable CPS trap",
			},
			"sessions_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sessions threshold trap",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSnmpServerEnableTrapsSystemAppsGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsSystemAppsGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsSystemAppsGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsSystemAppsGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerEnableTrapsSystemAppsGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsSystemAppsGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsSystemAppsGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsSystemAppsGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerEnableTrapsSystemAppsGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsSystemAppsGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsSystemAppsGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerEnableTrapsSystemAppsGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsSystemAppsGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsSystemAppsGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerEnableTrapsSystemAppsGlobal(d *schema.ResourceData) edpt.SnmpServerEnableTrapsSystemAppsGlobal {
	var ret edpt.SnmpServerEnableTrapsSystemAppsGlobal
	ret.Inst.CpsThreshold = d.Get("cps_threshold").(int)
	ret.Inst.SessionsThreshold = d.Get("sessions_threshold").(int)
	//omit uuid
	return ret
}
