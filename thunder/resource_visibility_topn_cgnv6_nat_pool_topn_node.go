package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityTopnCgnv6NatPoolTopnNode() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_topn_cgnv6_nat_pool_topn_node`: Activate cgnv6-nat-pool-topn template for cgnv6.nat.pool\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityTopnCgnv6NatPoolTopnNodeCreate,
		UpdateContext: resourceVisibilityTopnCgnv6NatPoolTopnNodeUpdate,
		ReadContext:   resourceVisibilityTopnCgnv6NatPoolTopnNodeRead,
		DeleteContext: resourceVisibilityTopnCgnv6NatPoolTopnNodeDelete,

		Schema: map[string]*schema.Schema{
			"activate": {
				Type: schema.TypeString, Optional: true, Description: "Name of the templated to be activated",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVisibilityTopnCgnv6NatPoolTopnNodeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnCgnv6NatPoolTopnNodeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnCgnv6NatPoolTopnNode(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityTopnCgnv6NatPoolTopnNodeRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityTopnCgnv6NatPoolTopnNodeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnCgnv6NatPoolTopnNodeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnCgnv6NatPoolTopnNode(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityTopnCgnv6NatPoolTopnNodeRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityTopnCgnv6NatPoolTopnNodeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnCgnv6NatPoolTopnNodeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnCgnv6NatPoolTopnNode(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityTopnCgnv6NatPoolTopnNodeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnCgnv6NatPoolTopnNodeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnCgnv6NatPoolTopnNode(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityTopnCgnv6NatPoolTopnNode(d *schema.ResourceData) edpt.VisibilityTopnCgnv6NatPoolTopnNode {
	var ret edpt.VisibilityTopnCgnv6NatPoolTopnNode
	ret.Inst.Activate = d.Get("activate").(string)
	//omit uuid
	return ret
}
