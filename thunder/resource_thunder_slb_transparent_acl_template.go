package thunder

//Thunder resource SlbTransparentAclTemplate

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTransparentAclTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbTransparentAclTemplateCreate,
		UpdateContext: resourceSlbTransparentAclTemplateUpdate,
		ReadContext:   resourceSlbTransparentAclTemplateRead,
		DeleteContext: resourceSlbTransparentAclTemplateDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTransparentAclTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTransparentAclTemplate (Inside resourceSlbTransparentAclTemplateCreate) ")
		name := d.Get("name").(string)
		data := dataToSlbTransparentAclTemplate(d)
		logger.Println("[INFO] received formatted data from method data to SlbTransparentAclTemplate --")
		d.SetId(name)
		err := go_thunder.PostSlbTransparentAclTemplate(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTransparentAclTemplateRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTransparentAclTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbTransparentAclTemplate (Inside resourceSlbTransparentAclTemplateRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetSlbTransparentAclTemplate(client.Token, name, client.Host)
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

func resourceSlbTransparentAclTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbTransparentAclTemplateRead(ctx, d, meta)
}

func resourceSlbTransparentAclTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbTransparentAclTemplateRead(ctx, d, meta)
}

func dataToSlbTransparentAclTemplate(d *schema.ResourceData) go_thunder.TransparentAclTemplate {
	var vc go_thunder.TransparentAclTemplate
	var c go_thunder.TransparentAclTemplateInstance
	c.Name = d.Get("name").(string)

	vc.Name = c
	return vc
}
