package thunder

//Thunder resource SlbTemplateTcp

import (
	"context"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceSlbTemplateTcp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbTemplateTcpCreate,
		UpdateContext: resourceSlbTemplateTcpUpdate,
		ReadContext:   resourceSlbTemplateTcpRead,
		DeleteContext: resourceSlbTemplateTcpDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"logging": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"idle_timeout": {
				Type:        schema.TypeInt,
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
			"initial_window_size": {
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
			"qos": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"insert_client_ip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"lan_fast_ack": {
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
			"reset_follow_fin": {
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
			"re_select_if_server_down": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"del_session_on_server_down": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
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

func resourceSlbTemplateTcpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateTcp (Inside resourceSlbTemplateTcpCreate) ")
		name1 := d.Get("name").(string)
		data := dataToSlbTemplateTcp(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateTcp --")
		d.SetId(name1)
		err := go_thunder.PostSlbTemplateTcp(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateTcpRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateTcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbTemplateTcp (Inside resourceSlbTemplateTcpRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSlbTemplateTcp(client.Token, name1, client.Host)
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

func resourceSlbTemplateTcpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying SlbTemplateTcp   (Inside resourceSlbTemplateTcpUpdate) ")
		data := dataToSlbTemplateTcp(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateTcp ")
		err := go_thunder.PutSlbTemplateTcp(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateTcpRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateTcpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateTcpDelete) " + name1)
		err := go_thunder.DeleteSlbTemplateTcp(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToSlbTemplateTcp(d *schema.ResourceData) go_thunder.SlbTemplateTcp {
	var vc go_thunder.SlbTemplateTcp
	var c go_thunder.SlbTemplateTCPInstance
	c.SlbTemplateTCPInstanceName = d.Get("name").(string)
	c.SlbTemplateTCPInstanceLogging = d.Get("logging").(string)
	c.SlbTemplateTCPInstanceIdleTimeout = d.Get("idle_timeout").(int)
	c.SlbTemplateTCPInstanceHalfOpenIdleTimeout = d.Get("half_open_idle_timeout").(int)
	c.SlbTemplateTCPInstanceHalfCloseIdleTimeout = d.Get("half_close_idle_timeout").(int)
	c.SlbTemplateTCPInstanceInitialWindowSize = d.Get("initial_window_size").(int)
	c.SlbTemplateTCPInstanceForceDeleteTimeout = d.Get("force_delete_timeout").(int)
	c.SlbTemplateTCPInstanceForceDeleteTimeout100Ms = d.Get("force_delete_timeout_100ms").(int)
	c.SlbTemplateTCPInstanceAliveIfActive = d.Get("alive_if_active").(int)
	c.SlbTemplateTCPInstanceQos = d.Get("qos").(int)
	c.SlbTemplateTCPInstanceInsertClientIP = d.Get("insert_client_ip").(int)
	c.SlbTemplateTCPInstanceLanFastAck = d.Get("lan_fast_ack").(int)
	c.SlbTemplateTCPInstanceResetFwd = d.Get("reset_fwd").(int)
	c.SlbTemplateTCPInstanceResetRev = d.Get("reset_rev").(int)
	c.SlbTemplateTCPInstanceResetFollowFin = d.Get("reset_follow_fin").(int)
	c.SlbTemplateTCPInstanceDisable = d.Get("disable").(int)
	c.SlbTemplateTCPInstanceDown = d.Get("down").(int)
	c.SlbTemplateTCPInstanceReSelectIfServerDown = d.Get("re_select_if_server_down").(int)
	c.SlbTemplateTCPInstanceDelSessionOnServerDown = d.Get("del_session_on_server_down").(int)
	c.SlbTemplateTCPInstanceUserTag = d.Get("user_tag").(string)

	vc.SlbTemplateTCPInstanceName = c
	return vc
}
