package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpNeighborTrunkNeighbor() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_bgp_neighbor_trunk_neighbor`: Specify a trunk unnumbered neighbor\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterBgpNeighborTrunkNeighborCreate,
		UpdateContext: resourceRouterBgpNeighborTrunkNeighborUpdate,
		ReadContext:   resourceRouterBgpNeighborTrunkNeighborRead,
		DeleteContext: resourceRouterBgpNeighborTrunkNeighborDelete,

		Schema: map[string]*schema.Schema{
			"peer_group_name": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"trunk": {
				Type: schema.TypeInt, Required: true, Description: "Trunk interface number",
			},
			"unnumbered": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"as_number": {
				Type: schema.TypeString, Required: true, Description: "AsNumber",
			},
		},
	}
}
func resourceRouterBgpNeighborTrunkNeighborCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNeighborTrunkNeighborCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNeighborTrunkNeighbor(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpNeighborTrunkNeighborRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterBgpNeighborTrunkNeighborUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNeighborTrunkNeighborUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNeighborTrunkNeighbor(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpNeighborTrunkNeighborRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterBgpNeighborTrunkNeighborDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNeighborTrunkNeighborDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNeighborTrunkNeighbor(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterBgpNeighborTrunkNeighborRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNeighborTrunkNeighborRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNeighborTrunkNeighbor(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointRouterBgpNeighborTrunkNeighbor(d *schema.ResourceData) edpt.RouterBgpNeighborTrunkNeighbor {
	var ret edpt.RouterBgpNeighborTrunkNeighbor
	ret.Inst.PeerGroupName = d.Get("peer_group_name").(string)
	ret.Inst.Trunk = d.Get("trunk").(int)
	ret.Inst.Unnumbered = d.Get("unnumbered").(int)
	//omit uuid
	ret.Inst.AsNumber = d.Get("as_number").(string)
	return ret
}
