package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugDnsRrl() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_dns_rrl`: Debug DNS Response-Rate-Limiting\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugDnsRrlCreate,
		UpdateContext: resourceDebugDnsRrlUpdate,
		ReadContext:   resourceDebugDnsRrlRead,
		DeleteContext: resourceDebugDnsRrlDelete,

		Schema: map[string]*schema.Schema{
			"dumy": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dummy",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugDnsRrlCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugDnsRrlCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugDnsRrl(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugDnsRrlRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugDnsRrlUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugDnsRrlUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugDnsRrl(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugDnsRrlRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugDnsRrlDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugDnsRrlDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugDnsRrl(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugDnsRrlRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugDnsRrlRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugDnsRrl(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugDnsRrl(d *schema.ResourceData) edpt.DebugDnsRrl {
	var ret edpt.DebugDnsRrl
	ret.Inst.Dumy = d.Get("dumy").(int)
	//omit uuid
	return ret
}
