package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkSpanningTreeModeStp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_spanning_tree_mode_stp`: Configure spanning tree protocol STP\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkSpanningTreeModeStpCreate,
		UpdateContext: resourceNetworkSpanningTreeModeStpUpdate,
		ReadContext:   resourceNetworkSpanningTreeModeStpRead,
		DeleteContext: resourceNetworkSpanningTreeModeStpDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable spanning tree RSTP mode",
			},
			"priority": {
				Type: schema.TypeInt, Optional: true, Default: 32768, Description: "set bridge priority",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vlan_start": {
				Type: schema.TypeInt, Optional: true, Description: "VLAN ID",
			},
		},
	}
}
func resourceNetworkSpanningTreeModeStpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkSpanningTreeModeStpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkSpanningTreeModeStp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkSpanningTreeModeStpRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkSpanningTreeModeStpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkSpanningTreeModeStpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkSpanningTreeModeStp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkSpanningTreeModeStpRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkSpanningTreeModeStpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkSpanningTreeModeStpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkSpanningTreeModeStp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkSpanningTreeModeStpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkSpanningTreeModeStpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkSpanningTreeModeStp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNetworkSpanningTreeModeStp(d *schema.ResourceData) edpt.NetworkSpanningTreeModeStp {
	var ret edpt.NetworkSpanningTreeModeStp
	ret.Inst.Action = d.Get("action").(int)
	ret.Inst.Priority = d.Get("priority").(int)
	//omit uuid
	ret.Inst.VlanStart = d.Get("vlan_start").(int)
	return ret
}
