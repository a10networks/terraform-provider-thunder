package thunder

//Thunder resource SlbTransperentTcpTemplate

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTransperentTcpTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbTransperentTcpTemplateCreate,
		UpdateContext: resourceSlbTransperentTcpTemplateUpdate,
		ReadContext:   resourceSlbTransperentTcpTemplateRead,
		DeleteContext: resourceSlbTransperentTcpTemplateDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTransperentTcpTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTransperentTcpTemplate (Inside resourceSlbTransperentTcpTemplateCreate) ")
		name := d.Get("name").(string)
		data := dataToSlbTransperentTcpTemplate(d)
		logger.Println("[INFO] received formatted data from method data to SlbTransperentTcpTemplate --")
		d.SetId(name)
		err := go_thunder.PostSlbTransperentTcpTemplate(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTransperentTcpTemplateRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTransperentTcpTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbTransperentTcpTemplate (Inside resourceSlbTransperentTcpTemplateRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetSlbTransperentTcpTemplate(client.Token, name, client.Host)
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

func resourceSlbTransperentTcpTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbTransperentTcpTemplateRead(ctx, d, meta)
}

func resourceSlbTransperentTcpTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbTransperentTcpTemplateRead(ctx, d, meta)
}

func dataToSlbTransperentTcpTemplate(d *schema.ResourceData) go_thunder.TransperentTcpTemplate {
	var vc go_thunder.TransperentTcpTemplate
	var c go_thunder.TransperentTcpTemplateInstance
	c.Name = d.Get("name").(string)

	vc.Name = c
	return vc
}
