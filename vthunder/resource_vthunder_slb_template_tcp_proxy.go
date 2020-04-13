package vthunder

//vThunder resource SlbTemplateTcpProxy

import (
	"log"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbTemplateTcpProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbTemplateTcpProxyCreate,
		Update: resourceSlbTemplateTcpProxyUpdate,
		Read:   resourceSlbTemplateTcpProxyRead,
		Delete: resourceSlbTemplateTcpProxyDelete,
		Schema: map[string]*schema.Schema{
			"disable_sack": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"insert_client_ip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_tcp_timestamps": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"reno": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"reassembly_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"maxburst": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"initial_window_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"mss": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"force_delete_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"down": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"limited_slowstart": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"half_close_idle_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"receive_buffer": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"qos": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"half_open_idle_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"idle_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"init_cwnd": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"alive_if_active": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dynamic_buffer_allocation": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ack_aggressiveness": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"disable_window_scale": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"invalid_rate_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"transmit_buffer": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_abc": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"syn_retries": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"retransmit_retries": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"server_down_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"min_rto": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"reset_rev": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"keepalive_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"force_delete_timeout_100ms": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"reassembly_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"keepalive_probes": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"fin_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"psh_flag_optimization": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"backend_wscale": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"nagle": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"reset_fwd": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"early_retransmit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"del_session_on_server_down": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"timewait": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTemplateTcpProxyCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateTcpProxy (Inside resourceSlbTemplateTcpProxyCreate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateTcpProxy(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateTcpProxy --")
		d.SetId(name)
		go_vthunder.PostSlbTemplateTcpProxy(client.Token, data, client.Host)

		return resourceSlbTemplateTcpProxyRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateTcpProxyRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbTemplateTcpProxy (Inside resourceSlbTemplateTcpProxyRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetSlbTemplateTcpProxy(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbTemplateTcpProxyUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplateTcpProxy   (Inside resourceSlbTemplateTcpProxyUpdate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateTcpProxy(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateTcpProxy ")
		d.SetId(name)
		go_vthunder.PutSlbTemplateTcpProxy(client.Token, name, data, client.Host)

		return resourceSlbTemplateTcpProxyRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateTcpProxyDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateTcpProxyDelete) " + name)
		err := go_vthunder.DeleteSlbTemplateTcpProxy(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbTemplateTcpProxy(d *schema.ResourceData) go_vthunder.TCPProxy {
	var vc go_vthunder.TCPProxy
	var c go_vthunder.TCPProxyInstance
	c.Qos = d.Get("qos").(int)
	c.InitCwnd = d.Get("init_cwnd").(int)
	c.IdleTimeout = d.Get("idle_timeout").(int)
	c.FinTimeout = d.Get("fin_timeout").(int)
	c.HalfOpenIdleTimeout = d.Get("half_open_idle_timeout").(int)
	c.Reno = d.Get("reno").(int)
	c.Down = d.Get("down").(int)
	c.EarlyRetransmit = d.Get("early_retransmit").(int)
	c.ServerDownAction = d.Get("server_down_action").(string)
	c.Timewait = d.Get("timewait").(int)
	c.MinRto = d.Get("min_rto").(int)
	c.DynamicBufferAllocation = d.Get("dynamic_buffer_allocation").(int)
	c.LimitedSlowstart = d.Get("limited_slowstart").(int)
	c.DisableSack = d.Get("disable_sack").(int)
	c.DisableWindowScale = d.Get("disable_window_scale").(int)
	c.AliveIfActive = d.Get("alive_if_active").(int)
	c.Mss = d.Get("mss").(int)
	c.KeepaliveInterval = d.Get("keepalive_interval").(int)
	c.RetransmitRetries = d.Get("retransmit_retries").(int)
	c.InsertClientIP = d.Get("insert_client_ip").(int)
	c.TransmitBuffer = d.Get("transmit_buffer").(int)
	c.Nagle = d.Get("nagle").(int)
	c.ForceDeleteTimeout100Ms = d.Get("force_delete_timeout_100ms").(int)
	c.InitialWindowSize = d.Get("initial_window_size").(int)
	c.KeepaliveProbes = d.Get("keepalive_probes").(int)
	c.PshFlagOptimization = d.Get("psh_flag_optimization").(int)
	c.AckAggressiveness = d.Get("ack_aggressiveness").(string)
	c.BackendWscale = d.Get("backend_wscale").(int)
	c.Disable = d.Get("disable").(int)
	c.ResetRev = d.Get("reset_rev").(int)
	c.Maxburst = d.Get("maxburst").(int)
	c.ReceiveBuffer = d.Get("receive_buffer").(int)
	c.DelSessionOnServerDown = d.Get("del_session_on_server_down").(int)
	c.Name = d.Get("name").(string)
	c.ReassemblyTimeout = d.Get("reassembly_timeout").(int)
	c.ResetFwd = d.Get("reset_fwd").(int)
	c.DisableTCPTimestamps = d.Get("disable_tcp_timestamps").(int)
	c.SynRetries = d.Get("syn_retries").(int)
	c.ForceDeleteTimeout = d.Get("force_delete_timeout").(int)
	c.UserTag = d.Get("user_tag").(string)
	c.ReassemblyLimit = d.Get("reassembly_limit").(int)
	c.InvalidRateLimit = d.Get("invalid_rate_limit").(int)
	c.DisableAbc = d.Get("disable_abc").(int)
	c.HalfCloseIdleTimeout = d.Get("half_close_idle_timeout").(int)
	vc.UUID = c
	return vc
}
