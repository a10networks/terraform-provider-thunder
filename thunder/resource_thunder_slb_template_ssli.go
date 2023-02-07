package thunder

//Thunder resource TemplateSSLI

import (
	"context"
	"log"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateSSLI() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTemplateSSLICreate,
		UpdateContext: resourceTemplateSSLIUpdate,
		ReadContext:   resourceTemplateSSLIRead,
		DeleteContext: resourceTemplateSSLIDelete,
		Schema: map[string]*schema.Schema{
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
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceTemplateSSLICreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating TemplateSSLI (Inside resourceTemplateSSLICreate) ")
		name := d.Get("name").(string)
		data := dataToTemplateSSLI(d)
		logger.Println("[INFO] received formatted data from method data to TemplateSSLI --")
		d.SetId(name)
		err := go_thunder.PostTemplateSSLI(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceTemplateSSLIRead(ctx, d, meta)

	}
	return diags
}

func resourceTemplateSSLIRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading TemplateSSLI (Inside resourceTemplateSSLIRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetTemplateSSLI(client.Token, name, client.Host)
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

func resourceTemplateSSLIUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Modifying TemplateSSLI   (Inside resourceTemplateSSLIUpdate) ")
		name := d.Get("name").(string)
		data := dataToTemplateSSLI(d)
		logger.Println("[INFO] received formatted data from method data to TemplateSSLI ")
		d.SetId(name)
		err := go_thunder.PutTemplateSSLI(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceTemplateSSLIRead(ctx, d, meta)

	}
	return diags
}

func resourceTemplateSSLIDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceTemplateSSLIDelete) " + name)
		err := go_thunder.DeleteTemplateSSLI(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return diags
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToTemplateSSLI(d *schema.ResourceData) go_thunder.SSLI {
	var vc go_thunder.SSLI
	var c go_thunder.SsliInstance
	c.Name = d.Get("name").(string)
	c.Type = d.Get("type").(string)
	c.UserTag = d.Get("user_tag").(string)
	vc.UUID = c
	return vc
}
