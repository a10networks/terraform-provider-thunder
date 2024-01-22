package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDnsCacheFqdnManualOverrideAction() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dns_cache_fqdn_manual_override_action`: DNS Cache FQDN response action set\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDnsCacheFqdnManualOverrideActionCreate,
		UpdateContext: resourceDdosDnsCacheFqdnManualOverrideActionUpdate,
		ReadContext:   resourceDdosDnsCacheFqdnManualOverrideActionRead,
		DeleteContext: resourceDdosDnsCacheFqdnManualOverrideActionDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'default': Default; 'forward': Forward to DNS server; 'drop': Drop the request; 'serve-from-cache': Serve DNS records;",
			},
			"fqdn_name": {
				Type: schema.TypeString, Required: true, Description: "Specify fqdn name",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceDdosDnsCacheFqdnManualOverrideActionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheFqdnManualOverrideActionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheFqdnManualOverrideAction(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDnsCacheFqdnManualOverrideActionRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDnsCacheFqdnManualOverrideActionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheFqdnManualOverrideActionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheFqdnManualOverrideAction(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDnsCacheFqdnManualOverrideActionRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDnsCacheFqdnManualOverrideActionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheFqdnManualOverrideActionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheFqdnManualOverrideAction(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDnsCacheFqdnManualOverrideActionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheFqdnManualOverrideActionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheFqdnManualOverrideAction(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDnsCacheFqdnManualOverrideAction(d *schema.ResourceData) edpt.DdosDnsCacheFqdnManualOverrideAction {
	var ret edpt.DdosDnsCacheFqdnManualOverrideAction
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.FqdnName = d.Get("fqdn_name").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
