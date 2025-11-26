package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemDomainListSettings() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_domain_list_settings`: Configure global domain classification list settings\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemDomainListSettingsCreate,
		UpdateContext: resourceSystemDomainListSettingsUpdate,
		ReadContext:   resourceSystemDomainListSettingsRead,
		DeleteContext: resourceSystemDomainListSettingsDelete,

		Schema: map[string]*schema.Schema{
			"concurrent_task": {
				Type: schema.TypeInt, Optional: true, Default: 6, Description: "Configure max concurrent AXFR task (Default 6)",
			},
			"domain_list_per_group": {
				Type: schema.TypeString, Optional: true, Default: "16", Description: "'16': Allow 16 domain-list per group (Default); '32': Allow 32 domain-list per group;",
			},
			"polling_interval": {
				Type: schema.TypeString, Optional: true, Default: "10-second", Description: "'1-second': Set interval to 1 second; '5-second': Set interval to 5 seconds; '10-second': Set interval to 10 seconds (Default);",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemDomainListSettingsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDomainListSettingsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDomainListSettings(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemDomainListSettingsRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemDomainListSettingsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDomainListSettingsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDomainListSettings(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemDomainListSettingsRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemDomainListSettingsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDomainListSettingsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDomainListSettings(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemDomainListSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDomainListSettingsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDomainListSettings(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemDomainListSettings(d *schema.ResourceData) edpt.SystemDomainListSettings {
	var ret edpt.SystemDomainListSettings
	ret.Inst.ConcurrentTask = d.Get("concurrent_task").(int)
	ret.Inst.DomainListPerGroup = d.Get("domain_list_per_group").(string)
	ret.Inst.PollingInterval = d.Get("polling_interval").(string)
	//omit uuid
	return ret
}
