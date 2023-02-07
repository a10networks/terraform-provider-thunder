package thunder

//Thunder resource TemplateConnectionReuse

import (
	"context"
	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"util"
)

func resourceTemplateConnectionReuse() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTemplateConnectionReuseCreate,
		UpdateContext: resourceTemplateConnectionReuseUpdate,
		ReadContext:   resourceTemplateConnectionReuseRead,
		DeleteContext: resourceTemplateConnectionReuseDelete,
		Schema: map[string]*schema.Schema{
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"keep_alive_conn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"num_conn_per_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"preopen": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"limit_per_server": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceTemplateConnectionReuseCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating TemplateConnectionReuse (Inside resourceTemplateConnectionReuseCreate) ")
		name := d.Get("name").(string)
		data := dataToTemplateConnectionReuse(d)
		logger.Println("[INFO] received formatted data from method data to TemplateConnectionReuse --")
		d.SetId(name)
		err := go_thunder.PostTemplateConnectionReuse(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceTemplateConnectionReuseRead(ctx, d, meta)

	}
	return diags
}

func resourceTemplateConnectionReuseRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading TemplateConnectionReuse (Inside resourceTemplateConnectionReuseRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetTemplateConnectionReuse(client.Token, name, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return diags
	}
	return nil
}

func resourceTemplateConnectionReuseUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Modifying TemplateConnectionReuse   (Inside resourceTemplateConnectionReuseUpdate) ")
		name := d.Get("name").(string)
		data := dataToTemplateConnectionReuse(d)
		logger.Println("[INFO] received formatted data from method data to TemplateConnectionReuse ")
		d.SetId(name)
		err := go_thunder.PutTemplateConnectionReuse(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceTemplateConnectionReuseRead(ctx, d, meta)

	}
	return diags
}

func resourceTemplateConnectionReuseDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceTemplateConnectionReuseDelete) " + name)
		err := go_thunder.DeleteTemplateConnectionReuse(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return diags
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToTemplateConnectionReuse(d *schema.ResourceData) go_thunder.Connection_Reuse {
	var vc go_thunder.Connection_Reuse
	var c go_thunder.ConnectionReuseInstance
	c.KeepAliveConn = d.Get("keep_alive_conn").(int)
	c.LimitPerServer = d.Get("limit_per_server").(int)
	c.Name = d.Get("name").(string)
	c.NumConnPerPort = d.Get("num_conn_per_port").(int)
	c.Preopen = d.Get("preopen").(int)
	c.Timeout = d.Get("timeout").(int)
	c.UserTag = d.Get("user_tag").(string)
	vc.Name = c
	return vc
}
