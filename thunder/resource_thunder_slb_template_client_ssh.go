package thunder

//Thunder resource TemplateClientSsh

import (
	"context"
	"log"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateClientSsh() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTemplateClientSshCreate,
		UpdateContext: resourceTemplateClientSshUpdate,
		ReadContext:   resourceTemplateClientSshRead,
		DeleteContext: resourceTemplateClientSshDelete,
		Schema: map[string]*schema.Schema{
			"forward_proxy_hostkey": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"encrypted": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"passphrase": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"forward_proxy_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceTemplateClientSshCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating TemplateClientSsh (Inside resourceTemplateClientSshCreate) ")
		name := d.Get("name").(string)
		data := dataToTemplateClientSsh(d)
		logger.Println("[INFO] received formatted data from method data to TemplateClientSsh --")
		d.SetId(name)
		err := go_thunder.PostTemplateClientSsh(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceTemplateClientSshRead(ctx, d, meta)

	}
	return diags
}

func resourceTemplateClientSshRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading TemplateClientSsh (Inside resourceTemplateClientSshRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetTemplateClientSsh(client.Token, name, client.Host)
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

func resourceTemplateClientSshUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Modifying TemplateClientSsh   (Inside resourceTemplateClientSshUpdate) ")
		name := d.Get("name").(string)
		data := dataToTemplateClientSsh(d)
		logger.Println("[INFO] received V from method data to TemplateClientSsh ")
		d.SetId(name)
		err := go_thunder.PutTemplateClientSsh(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceTemplateClientSshRead(ctx, d, meta)

	}
	return diags
}

func resourceTemplateClientSshDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceTemplateClientSshDelete) " + name)
		err := go_thunder.DeleteTemplateClientSsh(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return diags
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToTemplateClientSsh(d *schema.ResourceData) go_thunder.ClientSSH {
	var vc go_thunder.ClientSSH
	var c go_thunder.ClientSSHInstance

	c.Encrypted = d.Get("encrypted").(string)
	c.ForwardProxyEnable = d.Get("forward_proxy_enable").(int)
	c.ForwardProxyHostkey = d.Get("forward_proxy_hostkey").(string)
	c.Name = d.Get("name").(string)
	c.Passphrase = d.Get("passphrase").(string)
	c.UserTag = d.Get("user_tag").(string)

	vc.UUID = c
	return vc
}
