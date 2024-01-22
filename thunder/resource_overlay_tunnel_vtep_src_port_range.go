package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOverlayTunnelVtepSrcPortRange() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_overlay_tunnel_vtep_src_port_range`: Layer-4 Source Port Range\n\n__PLACEHOLDER__",
		CreateContext: resourceOverlayTunnelVtepSrcPortRangeCreate,
		UpdateContext: resourceOverlayTunnelVtepSrcPortRangeUpdate,
		ReadContext:   resourceOverlayTunnelVtepSrcPortRangeRead,
		DeleteContext: resourceOverlayTunnelVtepSrcPortRangeDelete,

		Schema: map[string]*schema.Schema{
			"max_port": {
				Type: schema.TypeInt, Optional: true, Default: 65535, Description: "Maximum Port Number",
			},
			"min_port": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Minimum Port Number",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"id1": {
				Type: schema.TypeString, Required: true, Description: "Id1",
			},
		},
	}
}
func resourceOverlayTunnelVtepSrcPortRangeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepSrcPortRangeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepSrcPortRange(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepSrcPortRangeRead(ctx, d, meta)
	}
	return diags
}

func resourceOverlayTunnelVtepSrcPortRangeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepSrcPortRangeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepSrcPortRange(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepSrcPortRangeRead(ctx, d, meta)
	}
	return diags
}
func resourceOverlayTunnelVtepSrcPortRangeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepSrcPortRangeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepSrcPortRange(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceOverlayTunnelVtepSrcPortRangeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepSrcPortRangeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepSrcPortRange(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointOverlayTunnelVtepSrcPortRange(d *schema.ResourceData) edpt.OverlayTunnelVtepSrcPortRange {
	var ret edpt.OverlayTunnelVtepSrcPortRange
	ret.Inst.MaxPort = d.Get("max_port").(int)
	ret.Inst.MinPort = d.Get("min_port").(int)
	//omit uuid
	ret.Inst.Id1 = d.Get("id1").(string)
	return ret
}
