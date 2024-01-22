package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDnsCacheZoneManualOverrideAction() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dns_cache_zone_manual_override_action`: DNS Cache FQDN response action set\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDnsCacheZoneManualOverrideActionCreate,
		UpdateContext: resourceDdosDnsCacheZoneManualOverrideActionUpdate,
		ReadContext:   resourceDdosDnsCacheZoneManualOverrideActionRead,
		DeleteContext: resourceDdosDnsCacheZoneManualOverrideActionDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'default': Default; 'forward': Forward to DNS server; 'drop': Drop the request; 'serve-from-cache': Serve DNS records;",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "Specify zone name",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceDdosDnsCacheZoneManualOverrideActionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheZoneManualOverrideActionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheZoneManualOverrideAction(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDnsCacheZoneManualOverrideActionRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDnsCacheZoneManualOverrideActionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheZoneManualOverrideActionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheZoneManualOverrideAction(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDnsCacheZoneManualOverrideActionRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDnsCacheZoneManualOverrideActionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheZoneManualOverrideActionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheZoneManualOverrideAction(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDnsCacheZoneManualOverrideActionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheZoneManualOverrideActionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheZoneManualOverrideAction(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDnsCacheZoneManualOverrideAction(d *schema.ResourceData) edpt.DdosDnsCacheZoneManualOverrideAction {
	var ret edpt.DdosDnsCacheZoneManualOverrideAction
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
