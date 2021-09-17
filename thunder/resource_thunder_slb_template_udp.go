package thunder

//Thunder resource SlbTemplateUdp

import (
	"context"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceSlbTemplateUdp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbTemplateUdpCreate,
		UpdateContext: resourceSlbTemplateUdpUpdate,
		ReadContext:   resourceSlbTemplateUdpRead,
		DeleteContext: resourceSlbTemplateUdpDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"idle_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"qos": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"stateless_conn_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"immediate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"short": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"age": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"re_select_if_server_down": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_clear_session": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"radius_lb_method_hash_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"avp": {
				Type:        schema.TypeString,
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

func resourceSlbTemplateUdpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateUdp (Inside resourceSlbTemplateUdpCreate) ")
		name1 := d.Get("name").(string)
		data := dataToSlbTemplateUdp(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateUdp --")
		d.SetId(name1)
		err := go_thunder.PostSlbTemplateUdp(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateUdpRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateUdpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbTemplateUdp (Inside resourceSlbTemplateUdpRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSlbTemplateUdp(client.Token, name1, client.Host)
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

func resourceSlbTemplateUdpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying SlbTemplateUdp   (Inside resourceSlbTemplateUdpUpdate) ")
		data := dataToSlbTemplateUdp(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateUdp ")
		err := go_thunder.PutSlbTemplateUdp(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateUdpRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateUdpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateUdpDelete) " + name1)
		err := go_thunder.DeleteSlbTemplateUdp(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToSlbTemplateUdp(d *schema.ResourceData) go_thunder.SlbTemplateUdp {
	var vc go_thunder.SlbTemplateUdp
	var c go_thunder.SlbTemplateUDPInstance
	c.SlbTemplateUDPInstanceName = d.Get("name").(string)
	c.SlbTemplateUDPInstanceIdleTimeout = d.Get("idle_timeout").(int)
	c.SlbTemplateUDPInstanceQos = d.Get("qos").(int)
	c.SlbTemplateUDPInstanceStatelessConnTimeout = d.Get("stateless_conn_timeout").(int)
	c.SlbTemplateUDPInstanceImmediate = d.Get("immediate").(int)
	c.SlbTemplateUDPInstanceShort = d.Get("short").(int)
	c.SlbTemplateUDPInstanceAge = d.Get("age").(int)
	c.SlbTemplateUDPInstanceReSelectIfServerDown = d.Get("re_select_if_server_down").(int)
	c.SlbTemplateUDPInstanceDisableClearSession = d.Get("disable_clear_session").(int)
	c.SlbTemplateUDPInstanceRadiusLbMethodHashType = d.Get("radius_lb_method_hash_type").(string)
	c.SlbTemplateUDPInstanceAvp = d.Get("avp").(string)
	c.SlbTemplateUDPInstanceUserTag = d.Get("user_tag").(string)

	vc.SlbTemplateUDPInstanceName = c
	return vc
}
