package vthunder

//vThunder resource Overlay tunnel Options

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceOverlayTunnelOptions() *schema.Resource {
	return &schema.Resource{
		Create: resourceOverlayTunnelOptionsCreate,
		Update: resourceOverlayTunnelOptionsUpdate,
		Read:   resourceOverlayTunnelOptionsRead,
		Delete: resourceOverlayTunnelOptionsDelete,

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

func resourceOverlayTunnelOptionsCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	logger.Println("[INFO] Creating options (Inside resourceOverlayTunnelOptionsCreate)")

	if client.Host != "" {
		d.SetId("1")
		vc := dataToOptions(d)
		go_vthunder.PostOptions(client.Token, vc, client.Host)

		return resourceOverlayTunnelOptionsRead(d, meta)
	}
	return nil
}

func resourceOverlayTunnelOptionsRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading options (Inside resourceOverlayTunnelOptionsRead)")

	client := meta.(vThunder)

	if client.Host != "" {

		vc, err := go_vthunder.GetOptions(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No options found")
			return nil
		}

		return err
	}
	return nil
}

func resourceOverlayTunnelOptionsUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceOverlayTunnelOptionsRead(d, meta)
}

func resourceOverlayTunnelOptionsDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceOverlayTunnelOptionsRead(d, meta)
}

//Utility method to instantiate OverlayTunnelOptions Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToOptions(d *schema.ResourceData) go_vthunder.Options {
	var vc go_vthunder.Options

	var c go_vthunder.OptionsInstance

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
