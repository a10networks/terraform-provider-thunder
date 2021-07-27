package thunder

//Thunder resource Overlay tunnel Options

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOverlayTunnelOptions() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOverlayTunnelOptionsCreate,
		UpdateContext: resourceOverlayTunnelOptionsUpdate,
		ReadContext:   resourceOverlayTunnelOptionsRead,
		DeleteContext: resourceOverlayTunnelOptionsDelete,

		Schema: map[string]*schema.Schema{
			"tcp_mss_adjust_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"nvgre_disable_flow_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ip_dscp_preserve": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"vxlan_dest_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"gateway_mac": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"nvgre_key_mode_lower24": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}

}

func resourceOverlayTunnelOptionsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	logger.Println("[INFO] Creating options (Inside resourceOverlayTunnelOptionsCreate)")

	if client.Host != "" {
		d.SetId("1")
		vc := dataToOptions(d)
		err := go_thunder.PostOptions(client.Token, vc, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelOptionsRead(ctx, d, meta)
	}
	return diags
}

func resourceOverlayTunnelOptionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading options (Inside resourceOverlayTunnelOptionsRead)")

	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {

		vc, err := go_thunder.GetOptions(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if vc == nil {
			logger.Println("[INFO] No options found")
			return nil
		}

		return diags
	}
	return nil
}

func resourceOverlayTunnelOptionsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceOverlayTunnelOptionsRead(ctx, d, meta)
}

func resourceOverlayTunnelOptionsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceOverlayTunnelOptionsRead(ctx, d, meta)
}

//Utility method to instantiate OverlayTunnelOptions Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToOptions(d *schema.ResourceData) go_thunder.Options {
	var vc go_thunder.Options

	var c go_thunder.OptionsInstance

	c.NvgreKeyModeLower24 = d.Get("nvgre_key_mode_lower24").(int)
	c.UUID = d.Get("uuid").(string)
	c.TCPMssAdjustDisable = d.Get("tcp_mss_adjust_disable").(int)
	c.GatewayMac = d.Get("gateway_mac").(string)
	c.IPDscpPreserve = d.Get("ip_dscp_preserve").(int)
	c.NvgreDisableFlowID = d.Get("nvgre_disable_flow_id").(int)
	c.VxlanDestPort = d.Get("vxlan_dest_port").(int)

	vc.UUID = c

	return vc
}
