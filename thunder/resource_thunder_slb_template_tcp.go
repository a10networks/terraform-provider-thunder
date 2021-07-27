package thunder

//Thunder resource Tcp

import (
	"context"
	"log"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateTcp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTcpCreate,
		UpdateContext: resourceTcpUpdate,
		ReadContext:   resourceTcpRead,
		DeleteContext: resourceTcpDelete,

		Schema: map[string]*schema.Schema{
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
			"reset_fwd": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"del_session_on_server_down": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceTcpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Creating tcp (Inside resourceTcpCreate    " + name)
		v := dataToTcp(name, d)
		d.SetId(name)
		err := go_thunder.PostTcp(client.Token, v, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTcpRead(ctx, d, meta)
	}
	return diags
}

func resourceTcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading tcp (Inside resourceTcpRead)")

	if client.Host != "" {
		client := meta.(Thunder)

		name := d.Id()

		logger.Println("[INFO] Fetching tcp Read" + name)

		tcp, err := go_thunder.GetTcp(client.Token, name, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if tcp == nil {
			logger.Println("[INFO] No tcp found " + name)
			d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceTcpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Modifying tcp (Inside resourceTcpUpdate    " + name)
		v := dataToTcp(name, d)
		d.SetId(name)
		err := go_thunder.PutTcp(client.Token, name, v, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTcpRead(ctx, d, meta)
	}
	return diags
}

func resourceTcpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger := util.GetLoggerInstance()

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting tcp (Inside resourceTcpDelete) " + name)

		err := go_thunder.DeleteTcp(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete tcp  (%s) (%v)", name, err)
			return diags
		}
		d.SetId("")
		return nil
	}
	return nil
}

//utility method to instantiate tcp structure
func dataToTcp(name string, d *schema.ResourceData) go_thunder.TCP {
	var s go_thunder.TCP

	var sInstance go_thunder.TCPInstance

	sInstance.IdleTimeout = d.Get("idle_timeout").(int)
	sInstance.InsertClientIP = d.Get("insert_client_ip").(int)
	sInstance.LanFastAck = d.Get("lan_fast_ack").(int)
	sInstance.ResetFwd = d.Get("reset_fwd").(int)
	sInstance.ResetRev = d.Get("reset_rev").(int)
	sInstance.ResetFollowFin = d.Get("reset_follow_fin").(int)
	sInstance.DelSessionOnServerDown = d.Get("del_session_on_server_down").(int)

	s.Name = sInstance

	return s
}
