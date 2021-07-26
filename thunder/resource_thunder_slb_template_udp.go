package thunder

//Thunder resource Udp

import (
	"context"
	"log"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateUdp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUdpCreate,
		UpdateContext: resourceUdpUpdate,
		ReadContext:   resourceUdpRead,
		DeleteContext: resourceUdpDelete,

		Schema: map[string]*schema.Schema{
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
			"re_select_if_server_down": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"immediate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
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
			"stateless_conn_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceUdpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Creating udp (Inside resourceUdpCreate    " + name)
		v := dataToUdp(name, d)
		d.SetId(name)
		err := go_thunder.PostUdp(client.Token, v, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceUdpRead(ctx, d, meta)
	}
	return diags
}

func resourceUdpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading udp (Inside resourceUdpRead)")

	if client.Host != "" {
		client := meta.(Thunder)

		name := d.Id()

		logger.Println("[INFO] Fetching udp Read" + name)

		udp, err := go_thunder.GetUdp(client.Token, name, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if udp == nil {
			logger.Println("[INFO] No udp found " + name)
			d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceUdpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Modifying udp (Inside resourceUdpUpdate    " + name)
		v := dataToUdp(name, d)
		d.SetId(name)
		err := go_thunder.PutUdp(client.Token, name, v, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceUdpRead(ctx, d, meta)
	}
	return diags
}

func resourceUdpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger := util.GetLoggerInstance()

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting udp (Inside resourceUdpDelete) " + name)

		err := go_thunder.DeleteUdp(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete udp  (%s) (%v)", name, err)
			return diags
		}
		d.SetId("")
		return nil
	}
	return nil
}

//utility method to instantiate udp structure
func dataToUdp(name string, d *schema.ResourceData) go_thunder.UDP {
	var s go_thunder.UDP

	var sInstance go_thunder.UDPInstance

	sInstance.Qos = d.Get("qos").(int)
	sInstance.Name = d.Get("name").(string)
	sInstance.StatelessConnTimeout = d.Get("stateless_conn_timeout").(int)
	sInstance.IdleTimeout = d.Get("idle_timeout").(int)
	sInstance.ReSelectIfServerDown = d.Get("re_select_if_server_down").(int)
	sInstance.Immediate = d.Get("immediate").(int)
	sInstance.UserTag = d.Get("user_tag").(string)

	s.Name = sInstance

	return s
}
