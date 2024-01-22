package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOverlayTunnelOptions() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_overlay_tunnel_options`: Global overlay-tunnel configuration options\n\n__PLACEHOLDER__",
		CreateContext: resourceOverlayTunnelOptionsCreate,
		UpdateContext: resourceOverlayTunnelOptionsUpdate,
		ReadContext:   resourceOverlayTunnelOptionsRead,
		DeleteContext: resourceOverlayTunnelOptionsDelete,

		Schema: map[string]*schema.Schema{
			"fragmentation_mode_inner": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the inner-fragmentation",
			},
			"gateway_mac": {
				Type: schema.TypeString, Optional: true, Description: "MAC to be used with Gateway segment Id (MAC Address for the Gateway segment)",
			},
			"ip_dscp_preserve": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Copy DSCP bits from inner IP to outer IP header",
			},
			"nvgre_disable_flow_id": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Flow-ID computation for NVGRE",
			},
			"nvgre_key_mode_lower24": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use the lower 24-bits of the GRE key as the VSID",
			},
			"src_port_range": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"min_port": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Minimum Port Number",
						},
						"max_port": {
							Type: schema.TypeInt, Optional: true, Default: 65535, Description: "Maximum Port Number",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"tcp_mss_adjust_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable TCP MSS adjustment in SYN packet for tunnels",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vxlan_dest_port": {
				Type: schema.TypeInt, Optional: true, Description: "VXLAN UDP Destination Port (UDP Port Number (default 4789))",
			},
		},
	}
}
func resourceOverlayTunnelOptionsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelOptionsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelOptions(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelOptionsRead(ctx, d, meta)
	}
	return diags
}

func resourceOverlayTunnelOptionsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelOptionsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelOptions(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelOptionsRead(ctx, d, meta)
	}
	return diags
}
func resourceOverlayTunnelOptionsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelOptionsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelOptions(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceOverlayTunnelOptionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelOptionsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelOptions(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectOverlayTunnelOptionsSrcPortRange1080(d []interface{}) edpt.OverlayTunnelOptionsSrcPortRange1080 {

	count1 := len(d)
	var ret edpt.OverlayTunnelOptionsSrcPortRange1080
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MinPort = in["min_port"].(int)
		ret.MaxPort = in["max_port"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointOverlayTunnelOptions(d *schema.ResourceData) edpt.OverlayTunnelOptions {
	var ret edpt.OverlayTunnelOptions
	ret.Inst.FragmentationModeInner = d.Get("fragmentation_mode_inner").(int)
	ret.Inst.GatewayMac = d.Get("gateway_mac").(string)
	ret.Inst.IpDscpPreserve = d.Get("ip_dscp_preserve").(int)
	ret.Inst.NvgreDisableFlowId = d.Get("nvgre_disable_flow_id").(int)
	ret.Inst.NvgreKeyModeLower24 = d.Get("nvgre_key_mode_lower24").(int)
	ret.Inst.SrcPortRange = getObjectOverlayTunnelOptionsSrcPortRange1080(d.Get("src_port_range").([]interface{}))
	ret.Inst.TcpMssAdjustDisable = d.Get("tcp_mss_adjust_disable").(int)
	//omit uuid
	ret.Inst.VxlanDestPort = d.Get("vxlan_dest_port").(int)
	return ret
}
