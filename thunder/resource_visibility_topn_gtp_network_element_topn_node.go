package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityTopnGtpNetworkElementTopnNode() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_topn_gtp_network_element_topn_node`: Activate gtp-network-element-topn template for fw.gtp.network-element\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityTopnGtpNetworkElementTopnNodeCreate,
		UpdateContext: resourceVisibilityTopnGtpNetworkElementTopnNodeUpdate,
		ReadContext:   resourceVisibilityTopnGtpNetworkElementTopnNodeRead,
		DeleteContext: resourceVisibilityTopnGtpNetworkElementTopnNodeDelete,

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
func resourceVisibilityTopnGtpNetworkElementTopnNodeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnGtpNetworkElementTopnNodeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnGtpNetworkElementTopnNode(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityTopnGtpNetworkElementTopnNodeRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityTopnGtpNetworkElementTopnNodeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnGtpNetworkElementTopnNodeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnGtpNetworkElementTopnNode(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityTopnGtpNetworkElementTopnNodeRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityTopnGtpNetworkElementTopnNodeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnGtpNetworkElementTopnNodeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnGtpNetworkElementTopnNode(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityTopnGtpNetworkElementTopnNodeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnGtpNetworkElementTopnNodeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnGtpNetworkElementTopnNode(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityTopnGtpNetworkElementTopnNode(d *schema.ResourceData) edpt.VisibilityTopnGtpNetworkElementTopnNode {
	var ret edpt.VisibilityTopnGtpNetworkElementTopnNode
	ret.Inst.Activate = d.Get("activate").(string)
	//omit uuid
	return ret
}
