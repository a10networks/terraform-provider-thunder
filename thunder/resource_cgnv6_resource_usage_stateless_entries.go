package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6ResourceUsageStatelessEntries() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_resource_usage_stateless_entries`: Helper size for CGN Stateless Technologies\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6ResourceUsageStatelessEntriesCreate,
		UpdateContext: resourceCgnv6ResourceUsageStatelessEntriesUpdate,
		ReadContext:   resourceCgnv6ResourceUsageStatelessEntriesRead,
		DeleteContext: resourceCgnv6ResourceUsageStatelessEntriesDelete,

		Schema: map[string]*schema.Schema{
			"l4_session_count": {
				Type: schema.TypeInt, Optional: true, Description: "Helper size for CGN Stateless Technologies",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6ResourceUsageStatelessEntriesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6ResourceUsageStatelessEntriesCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6ResourceUsageStatelessEntries(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6ResourceUsageStatelessEntriesRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6ResourceUsageStatelessEntriesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6ResourceUsageStatelessEntriesUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6ResourceUsageStatelessEntries(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6ResourceUsageStatelessEntriesRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6ResourceUsageStatelessEntriesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6ResourceUsageStatelessEntriesDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6ResourceUsageStatelessEntries(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6ResourceUsageStatelessEntriesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6ResourceUsageStatelessEntriesRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6ResourceUsageStatelessEntries(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6ResourceUsageStatelessEntries(d *schema.ResourceData) edpt.Cgnv6ResourceUsageStatelessEntries {
	var ret edpt.Cgnv6ResourceUsageStatelessEntries
	ret.Inst.L4SessionCount = d.Get("l4_session_count").(int)
	//omit uuid
	return ret
}
