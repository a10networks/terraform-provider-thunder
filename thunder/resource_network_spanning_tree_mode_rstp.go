package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkSpanningTreeModeRstp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_spanning_tree_mode_rstp`: Configure spanning tree protocol RSTP\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkSpanningTreeModeRstpCreate,
		UpdateContext: resourceNetworkSpanningTreeModeRstpUpdate,
		ReadContext:   resourceNetworkSpanningTreeModeRstpRead,
		DeleteContext: resourceNetworkSpanningTreeModeRstpDelete,

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
func resourceNetworkSpanningTreeModeRstpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkSpanningTreeModeRstpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkSpanningTreeModeRstp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkSpanningTreeModeRstpRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkSpanningTreeModeRstpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkSpanningTreeModeRstpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkSpanningTreeModeRstp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkSpanningTreeModeRstpRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkSpanningTreeModeRstpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkSpanningTreeModeRstpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkSpanningTreeModeRstp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkSpanningTreeModeRstpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkSpanningTreeModeRstpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkSpanningTreeModeRstp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNetworkSpanningTreeModeRstp(d *schema.ResourceData) edpt.NetworkSpanningTreeModeRstp {
	var ret edpt.NetworkSpanningTreeModeRstp
	ret.Inst.Action = d.Get("action").(int)
	ret.Inst.Priority = d.Get("priority").(int)
	//omit uuid
	ret.Inst.VlanStart = d.Get("vlan_start").(int)
	return ret
}
