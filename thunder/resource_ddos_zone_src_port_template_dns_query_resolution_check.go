package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneSrcPortTemplateDnsQueryResolutionCheck() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_src_port_template_dns_query_resolution_check`: Mitigation options against external DNS server (SYM mode only)\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneSrcPortTemplateDnsQueryResolutionCheckCreate,
		UpdateContext: resourceDdosZoneSrcPortTemplateDnsQueryResolutionCheckUpdate,
		ReadContext:   resourceDdosZoneSrcPortTemplateDnsQueryResolutionCheckRead,
		DeleteContext: resourceDdosZoneSrcPortTemplateDnsQueryResolutionCheckDelete,

		Schema: map[string]*schema.Schema{
			"big_response_action": {
				Type: schema.TypeString, Optional: true, Default: "default", Description: "'default': Default, No action for future connections; 'blacklist-src': Blacklist the external server for future connections;",
			},
			"big_response_size": {
				Type: schema.TypeInt, Optional: true, Description: "Max DNS response size (in Bytes)",
			},
			"domain_lockup_action": {
				Type: schema.TypeString, Optional: true, Default: "default", Description: "'default': Default, No action for future connections; 'blacklist-src': Blacklist the external server for future connections;",
			},
			"session_timeout_value": {
				Type: schema.TypeInt, Optional: true, Description: "max session timeout (secs) between DNS external server and Protected object",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceDdosZoneSrcPortTemplateDnsQueryResolutionCheckCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneSrcPortTemplateDnsQueryResolutionCheckCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneSrcPortTemplateDnsQueryResolutionCheck(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneSrcPortTemplateDnsQueryResolutionCheckRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneSrcPortTemplateDnsQueryResolutionCheckUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneSrcPortTemplateDnsQueryResolutionCheckUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneSrcPortTemplateDnsQueryResolutionCheck(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneSrcPortTemplateDnsQueryResolutionCheckRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneSrcPortTemplateDnsQueryResolutionCheckDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneSrcPortTemplateDnsQueryResolutionCheckDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneSrcPortTemplateDnsQueryResolutionCheck(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneSrcPortTemplateDnsQueryResolutionCheckRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneSrcPortTemplateDnsQueryResolutionCheckRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneSrcPortTemplateDnsQueryResolutionCheck(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosZoneSrcPortTemplateDnsQueryResolutionCheck(d *schema.ResourceData) edpt.DdosZoneSrcPortTemplateDnsQueryResolutionCheck {
	var ret edpt.DdosZoneSrcPortTemplateDnsQueryResolutionCheck
	ret.Inst.BigResponseAction = d.Get("big_response_action").(string)
	ret.Inst.BigResponseSize = d.Get("big_response_size").(int)
	ret.Inst.DomainLockupAction = d.Get("domain_lockup_action").(string)
	ret.Inst.SessionTimeoutValue = d.Get("session_timeout_value").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
