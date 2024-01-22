package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlmEnterpriseHaHost() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_glm_enterprise_ha_host`: High Avalibility ELM host\n\n__PLACEHOLDER__",
		CreateContext: resourceGlmEnterpriseHaHostCreate,
		UpdateContext: resourceGlmEnterpriseHaHostUpdate,
		ReadContext:   resourceGlmEnterpriseHaHostRead,
		DeleteContext: resourceGlmEnterpriseHaHostDelete,

		Schema: map[string]*schema.Schema{
			"host_entry": {
				Type: schema.TypeString, Required: true, Description: "Enter the ELM hostname, IP or [IPV6]",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceGlmEnterpriseHaHostCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlmEnterpriseHaHostCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlmEnterpriseHaHost(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGlmEnterpriseHaHostRead(ctx, d, meta)
	}
	return diags
}

func resourceGlmEnterpriseHaHostUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlmEnterpriseHaHostUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlmEnterpriseHaHost(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGlmEnterpriseHaHostRead(ctx, d, meta)
	}
	return diags
}
func resourceGlmEnterpriseHaHostDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlmEnterpriseHaHostDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlmEnterpriseHaHost(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGlmEnterpriseHaHostRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlmEnterpriseHaHostRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlmEnterpriseHaHost(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointGlmEnterpriseHaHost(d *schema.ResourceData) edpt.GlmEnterpriseHaHost {
	var ret edpt.GlmEnterpriseHaHost
	ret.Inst.HostEntry = d.Get("host_entry").(string)
	//omit uuid
	return ret
}
