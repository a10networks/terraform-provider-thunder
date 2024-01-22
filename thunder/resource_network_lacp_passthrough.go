package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkLacpPassthrough() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_lacp_passthrough`: Peer ports to forward received lacp packets\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkLacpPassthroughCreate,
		UpdateContext: resourceNetworkLacpPassthroughUpdate,
		ReadContext:   resourceNetworkLacpPassthroughRead,
		DeleteContext: resourceNetworkLacpPassthroughDelete,

		Schema: map[string]*schema.Schema{
			"peer_from": {
				Type: schema.TypeInt, Required: true, Description: "Peer member to forward received LACP packets",
			},
			"peer_to": {
				Type: schema.TypeInt, Required: true, Description: "Peer member to forward received LACP packets",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceNetworkLacpPassthroughCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkLacpPassthroughCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkLacpPassthrough(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkLacpPassthroughRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkLacpPassthroughUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkLacpPassthroughUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkLacpPassthrough(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkLacpPassthroughRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkLacpPassthroughDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkLacpPassthroughDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkLacpPassthrough(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkLacpPassthroughRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkLacpPassthroughRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkLacpPassthrough(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNetworkLacpPassthrough(d *schema.ResourceData) edpt.NetworkLacpPassthrough {
	var ret edpt.NetworkLacpPassthrough
	ret.Inst.PeerFrom = d.Get("peer_from").(int)
	ret.Inst.PeerTo = d.Get("peer_to").(int)
	//omit uuid
	return ret
}
