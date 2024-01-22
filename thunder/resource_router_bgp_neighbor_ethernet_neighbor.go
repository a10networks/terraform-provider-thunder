package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpNeighborEthernetNeighbor() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_bgp_neighbor_ethernet_neighbor`: Specify an ethernet unnumbered neighbor\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterBgpNeighborEthernetNeighborCreate,
		UpdateContext: resourceRouterBgpNeighborEthernetNeighborUpdate,
		ReadContext:   resourceRouterBgpNeighborEthernetNeighborRead,
		DeleteContext: resourceRouterBgpNeighborEthernetNeighborDelete,

		Schema: map[string]*schema.Schema{
			"ethernet": {
				Type: schema.TypeInt, Required: true, Description: "Ethernet interface number",
			},
			"peer_group_name": {
				Type: schema.TypeString, Optional: true, Description: "",
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
func resourceRouterBgpNeighborEthernetNeighborCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNeighborEthernetNeighborCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNeighborEthernetNeighbor(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpNeighborEthernetNeighborRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterBgpNeighborEthernetNeighborUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNeighborEthernetNeighborUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNeighborEthernetNeighbor(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpNeighborEthernetNeighborRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterBgpNeighborEthernetNeighborDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNeighborEthernetNeighborDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNeighborEthernetNeighbor(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterBgpNeighborEthernetNeighborRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNeighborEthernetNeighborRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNeighborEthernetNeighbor(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointRouterBgpNeighborEthernetNeighbor(d *schema.ResourceData) edpt.RouterBgpNeighborEthernetNeighbor {
	var ret edpt.RouterBgpNeighborEthernetNeighbor
	ret.Inst.Ethernet = d.Get("ethernet").(int)
	ret.Inst.PeerGroupName = d.Get("peer_group_name").(string)
	ret.Inst.Unnumbered = d.Get("unnumbered").(int)
	//omit uuid
	ret.Inst.AsNumber = d.Get("as_number").(string)
	return ret
}
