package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOverlayTunnelOptionsSrcPortRange() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_overlay_tunnel_options_src_port_range`: Source Port Range\n\n__PLACEHOLDER__",
		CreateContext: resourceOverlayTunnelOptionsSrcPortRangeCreate,
		UpdateContext: resourceOverlayTunnelOptionsSrcPortRangeUpdate,
		ReadContext:   resourceOverlayTunnelOptionsSrcPortRangeRead,
		DeleteContext: resourceOverlayTunnelOptionsSrcPortRangeDelete,

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
		},
	}
}
func resourceOverlayTunnelOptionsSrcPortRangeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelOptionsSrcPortRangeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelOptionsSrcPortRange(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelOptionsSrcPortRangeRead(ctx, d, meta)
	}
	return diags
}

func resourceOverlayTunnelOptionsSrcPortRangeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelOptionsSrcPortRangeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelOptionsSrcPortRange(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelOptionsSrcPortRangeRead(ctx, d, meta)
	}
	return diags
}
func resourceOverlayTunnelOptionsSrcPortRangeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelOptionsSrcPortRangeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelOptionsSrcPortRange(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceOverlayTunnelOptionsSrcPortRangeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelOptionsSrcPortRangeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelOptionsSrcPortRange(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointOverlayTunnelOptionsSrcPortRange(d *schema.ResourceData) edpt.OverlayTunnelOptionsSrcPortRange {
	var ret edpt.OverlayTunnelOptionsSrcPortRange
	ret.Inst.MaxPort = d.Get("max_port").(int)
	ret.Inst.MinPort = d.Get("min_port").(int)
	//omit uuid
	return ret
}
