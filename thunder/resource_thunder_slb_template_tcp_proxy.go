package thunder

//Thunder resource SlbTemplateTcpProxy

import (
	"context"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceSlbTemplateTcpProxy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbTemplateTcpProxyCreate,
		UpdateContext: resourceSlbTemplateTcpProxyUpdate,
		ReadContext:   resourceSlbTemplateTcpProxyRead,
		DeleteContext: resourceSlbTemplateTcpProxyDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ack_aggressiveness": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"backend_wscale": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dynamic_buffer_allocation": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"fin_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"force_delete_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"force_delete_timeout_100ms": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"alive_if_active": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"idle_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"server_down_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"half_open_idle_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"half_close_idle_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"init_cwnd": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"initial_window_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"keepalive_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"keepalive_probes": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"mss": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"psh_flag_optimization": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"nagle": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"qos": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"receive_buffer": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"reno": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"transmit_buffer": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"reset_fwd": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"reset_rev": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"down": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"del_session_on_server_down": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"retransmit_retries": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"insert_client_ip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"syn_retries": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"timewait": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_tcp_timestamps": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_window_scale": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_sack": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"invalid_rate_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_abc": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"reassembly_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"reassembly_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"min_rto": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"limited_slowstart": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"early_retransmit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"maxburst": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"proxy_header": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proxy_header_action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"version": {
							Type:        schema.TypeString,
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
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTemplateTcpProxyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateTcpProxy (Inside resourceSlbTemplateTcpProxyCreate) ")
		name1 := d.Get("name").(string)
		data := dataToSlbTemplateTcpProxy(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateTcpProxy --")
		d.SetId(name1)
		err := go_thunder.PostSlbTemplateTcpProxy(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateTcpProxyRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateTcpProxyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbTemplateTcpProxy (Inside resourceSlbTemplateTcpProxyRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSlbTemplateTcpProxy(client.Token, name1, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return diags
}

func resourceSlbTemplateTcpProxyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying SlbTemplateTcpProxy   (Inside resourceSlbTemplateTcpProxyUpdate) ")
		data := dataToSlbTemplateTcpProxy(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateTcpProxy ")
		err := go_thunder.PutSlbTemplateTcpProxy(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateTcpProxyRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateTcpProxyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateTcpProxyDelete) " + name1)
		err := go_thunder.DeleteSlbTemplateTcpProxy(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToSlbTemplateTcpProxy(d *schema.ResourceData) go_thunder.SlbTemplateTcpProxy {
	var vc go_thunder.SlbTemplateTcpProxy
	var c go_thunder.SlbTemplateTCPProxyInstance
	c.SlbTemplateTCPProxyInstanceName = d.Get("name").(string)
	c.SlbTemplateTCPProxyInstanceAckAggressiveness = d.Get("ack_aggressiveness").(string)
	c.SlbTemplateTCPProxyInstanceBackendWscale = d.Get("backend_wscale").(int)
	c.SlbTemplateTCPProxyInstanceDynamicBufferAllocation = d.Get("dynamic_buffer_allocation").(int)
	c.SlbTemplateTCPProxyInstanceFinTimeout = d.Get("fin_timeout").(int)
	c.SlbTemplateTCPProxyInstanceForceDeleteTimeout = d.Get("force_delete_timeout").(int)
	c.SlbTemplateTCPProxyInstanceForceDeleteTimeout100Ms = d.Get("force_delete_timeout_100ms").(int)
	c.SlbTemplateTCPProxyInstanceAliveIfActive = d.Get("alive_if_active").(int)
	c.SlbTemplateTCPProxyInstanceIdleTimeout = d.Get("idle_timeout").(int)
	c.SlbTemplateTCPProxyInstanceServerDownAction = d.Get("server_down_action").(string)
	c.SlbTemplateTCPProxyInstanceHalfOpenIdleTimeout = d.Get("half_open_idle_timeout").(int)
	c.SlbTemplateTCPProxyInstanceHalfCloseIdleTimeout = d.Get("half_close_idle_timeout").(int)
	c.SlbTemplateTCPProxyInstanceInitCwnd = d.Get("init_cwnd").(int)
	c.SlbTemplateTCPProxyInstanceInitialWindowSize = d.Get("initial_window_size").(int)
	c.SlbTemplateTCPProxyInstanceKeepaliveInterval = d.Get("keepalive_interval").(int)
	c.SlbTemplateTCPProxyInstanceKeepaliveProbes = d.Get("keepalive_probes").(int)
	c.SlbTemplateTCPProxyInstanceMss = d.Get("mss").(int)
	c.SlbTemplateTCPProxyInstancePshFlagOptimization = d.Get("psh_flag_optimization").(int)
	c.SlbTemplateTCPProxyInstanceNagle = d.Get("nagle").(int)
	c.SlbTemplateTCPProxyInstanceQos = d.Get("qos").(int)
	c.SlbTemplateTCPProxyInstanceReceiveBuffer = d.Get("receive_buffer").(int)
	c.SlbTemplateTCPProxyInstanceReno = d.Get("reno").(int)
	c.SlbTemplateTCPProxyInstanceTransmitBuffer = d.Get("transmit_buffer").(int)
	c.SlbTemplateTCPProxyInstanceResetFwd = d.Get("reset_fwd").(int)
	c.SlbTemplateTCPProxyInstanceResetRev = d.Get("reset_rev").(int)
	c.SlbTemplateTCPProxyInstanceDisable = d.Get("disable").(int)
	c.SlbTemplateTCPProxyInstanceDown = d.Get("down").(int)
	c.SlbTemplateTCPProxyInstanceDelSessionOnServerDown = d.Get("del_session_on_server_down").(int)
	c.SlbTemplateTCPProxyInstanceRetransmitRetries = d.Get("retransmit_retries").(int)
	c.SlbTemplateTCPProxyInstanceInsertClientIP = d.Get("insert_client_ip").(int)
	c.SlbTemplateTCPProxyInstanceSynRetries = d.Get("syn_retries").(int)
	c.SlbTemplateTCPProxyInstanceTimewait = d.Get("timewait").(int)
	c.SlbTemplateTCPProxyInstanceDisableTCPTimestamps = d.Get("disable_tcp_timestamps").(int)
	c.SlbTemplateTCPProxyInstanceDisableWindowScale = d.Get("disable_window_scale").(int)
	c.SlbTemplateTCPProxyInstanceDisableSack = d.Get("disable_sack").(int)
	c.SlbTemplateTCPProxyInstanceInvalidRateLimit = d.Get("invalid_rate_limit").(int)
	c.SlbTemplateTCPProxyInstanceDisableAbc = d.Get("disable_abc").(int)
	c.SlbTemplateTCPProxyInstanceReassemblyTimeout = d.Get("reassembly_timeout").(int)
	c.SlbTemplateTCPProxyInstanceReassemblyLimit = d.Get("reassembly_limit").(int)
	c.SlbTemplateTCPProxyInstanceMinRto = d.Get("min_rto").(int)
	c.SlbTemplateTCPProxyInstanceLimitedSlowstart = d.Get("limited_slowstart").(int)
	c.SlbTemplateTCPProxyInstanceEarlyRetransmit = d.Get("early_retransmit").(int)
	c.SlbTemplateTCPProxyInstanceMaxburst = d.Get("maxburst").(int)

	var obj1 go_thunder.SlbTemplateTCPProxyInstanceProxyHeader
	prefix1 := "proxy_header.0."
	obj1.SlbTemplateTCPProxyInstanceProxyHeaderProxyHeaderAction = d.Get(prefix1 + "proxy_header_action").(string)
	obj1.SlbTemplateTCPProxyInstanceProxyHeaderVersion = d.Get(prefix1 + "version").(string)

	c.SlbTemplateTCPProxyInstanceProxyHeaderProxyHeaderAction = obj1

	c.SlbTemplateTCPProxyInstanceUserTag = d.Get("user_tag").(string)

	vc.SlbTemplateTCPProxyInstanceName = c
	return vc
}
