package vthunder

//vThunder resource Vrrp common

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
	"strconv"
)

func resourceVrrpCommon() *schema.Resource {
	return &schema.Resource{
		Create: resourceVrrpCommonCreate,
		Update: resourceVrrpCommonUpdate,
		Read:   resourceVrrpCommonRead,
		Delete: resourceVrrpCommonDelete,

		Schema: map[string]*schema.Schema{
			"preemption_delay": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_l4_packet_on_standby": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"set_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_default_vrid": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"hostid_append_to_vrid": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hostid_append_to_vrid_default": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"hostid_append_to_vrid_value": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"device_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"get_ready_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dead_timer": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"arp_retry": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"hello_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"track_event_delay": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"restart_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"inline_mode_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"preferred_trunk": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"preferred_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"inline_mode": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		},
	}

}

func resourceVrrpCommonCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	logger.Println("[INFO] Creating vrrp common (Inside resourceVrrpCommonCreate)")

	if client.Host != "" {
		name := d.Get("device_id").(int)
		vc := dataToVrrpCommon(d)
		d.SetId(strconv.Itoa(name))
		go_vthunder.PostVrrpCommon(client.Token, vc, client.Host)

		return resourceVrrpCommonRead(d, meta)
	}
	return nil
}

func resourceVrrpCommonRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading vrrp common (Inside resourceVrrpCommonRead)")

	client := meta.(vThunder)

	if client.Host != "" {
		
		name := d.Id()

		vc, err := go_vthunder.GetVrrpCommon(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No vrrp common found" + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceVrrpCommonUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceVrrpCommonRead(d, meta)
}

func resourceVrrpCommonDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceVrrpCommonRead(d, meta)
}

//Utility method to instantiate Vrrp Common Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToVrrpCommon(d *schema.ResourceData) go_vthunder.CommonInstance {
	var vc go_vthunder.CommonInstance

	var c go_vthunder.Common
	c.ForwardL4PacketOnStandby = d.Get("forward_l4_packet_on_standby").(int)
	c.GetReadyTime = d.Get("get_ready_time").(int)
	c.HelloInterval = d.Get("hello_interval").(int)
	c.UUID = d.Get("uuid").(string)
	c.PreemptionDelay = d.Get("preemption_delay").(int)
	c.SetID = d.Get("set_id").(int)
	c.DeviceID = d.Get("device_id").(int)
	c.ArpRetry = d.Get("arp_retry").(int)
	c.DeadTimer = d.Get("dead_timer").(int)
	c.DisableDefaultVrid = d.Get("disable_default_vrid").(int)
	c.TrackEventDelay = d.Get("track_event_delay").(int)
	c.Action = d.Get("action").(string)
	c.RestartTime = d.Get("restart_time").(int)

	var cfg go_vthunder.InlineModeCfg
	cfg.InlineMode = d.Get("inline_mode_cfg.0.inline_mode").(int)

	c.InlineMode = cfg

	var host_vrid go_vthunder.HostidAppendToVrid
	host_vrid.HostidAppendToVridDefault = d.Get("hostid_append_to_vrid.0.hostid_append_to_vrid_default").(int)
	c.HostidAppendToVridValue = host_vrid

	vc.UUID = c

	return vc
}
