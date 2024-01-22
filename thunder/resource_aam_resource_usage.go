package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamResourceUsage() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_resource_usage`: Configure AAM Resource Usage\n\n__PLACEHOLDER__",
		CreateContext: resourceAamResourceUsageCreate,
		UpdateContext: resourceAamResourceUsageUpdate,
		ReadContext:   resourceAamResourceUsageRead,
		DeleteContext: resourceAamResourceUsageDelete,

		Schema: map[string]*schema.Schema{
			"identity_provider_limit": {
				Type: schema.TypeInt, Optional: true, Default: 64, Description: "Total Number of Identity Provider exists in the System",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAamResourceUsageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamResourceUsageCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamResourceUsage(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamResourceUsageRead(ctx, d, meta)
	}
	return diags
}

func resourceAamResourceUsageUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamResourceUsageUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamResourceUsage(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamResourceUsageRead(ctx, d, meta)
	}
	return diags
}
func resourceAamResourceUsageDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamResourceUsageDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamResourceUsage(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamResourceUsageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamResourceUsageRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamResourceUsage(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAamResourceUsage(d *schema.ResourceData) edpt.AamResourceUsage {
	var ret edpt.AamResourceUsage
	ret.Inst.IdentityProviderLimit = d.Get("identity_provider_limit").(int)
	//omit uuid
	return ret
}
