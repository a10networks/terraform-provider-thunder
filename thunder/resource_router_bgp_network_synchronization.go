package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpNetworkSynchronization() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_bgp_network_synchronization`: help Perform IGP synchronization\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterBgpNetworkSynchronizationCreate,
		UpdateContext: resourceRouterBgpNetworkSynchronizationUpdate,
		ReadContext:   resourceRouterBgpNetworkSynchronizationRead,
		DeleteContext: resourceRouterBgpNetworkSynchronizationDelete,

		Schema: map[string]*schema.Schema{
			"network_synchronization": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Perform IGP synchronization",
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
func resourceRouterBgpNetworkSynchronizationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNetworkSynchronizationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNetworkSynchronization(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpNetworkSynchronizationRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterBgpNetworkSynchronizationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNetworkSynchronizationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNetworkSynchronization(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpNetworkSynchronizationRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterBgpNetworkSynchronizationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNetworkSynchronizationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNetworkSynchronization(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterBgpNetworkSynchronizationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNetworkSynchronizationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNetworkSynchronization(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointRouterBgpNetworkSynchronization(d *schema.ResourceData) edpt.RouterBgpNetworkSynchronization {
	var ret edpt.RouterBgpNetworkSynchronization
	ret.Inst.NetworkSynchronization = d.Get("network_synchronization").(int)
	//omit uuid
	ret.Inst.AsNumber = d.Get("as_number").(string)
	return ret
}
