package thunder

//Thunder resource TemplateImap_POP3

import (
	"context"
	"log"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateImap_POP3() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTemplateImap_POP3Create,
		UpdateContext: resourceTemplateImap_POP3Update,
		ReadContext:   resourceTemplateImap_POP3Read,
		DeleteContext: resourceTemplateImap_POP3Delete,
		Schema: map[string]*schema.Schema{
			"starttls": {
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
			"logindisabled": {
				Type:        schema.TypeInt,
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

func resourceTemplateImap_POP3Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating TemplateImap_POP3 (Inside resourceTemplateImap_POP3Create) ")
		name := d.Get("name").(string)
		data := dataToTemplateImapPOP(d)
		logger.Println("[INFO] received formatted data from method data to TemplateImap_POP3 --")
		d.SetId(name)
		err := go_thunder.PostTemplateImap_POP3(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceTemplateImap_POP3Read(ctx, d, meta)

	}
	return diags
}

func resourceTemplateImap_POP3Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading TemplateImap_POP3 (Inside resourceTemplateImap_POP3Read)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetTemplateImap_POP3(client.Token, name, client.Host)
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

func resourceTemplateImap_POP3Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Modifying TemplateImap_POP3   (Inside resourceTemplateImap_POP3Update) ")
		name := d.Get("name").(string)
		data := dataToTemplateImapPOP(d)
		logger.Println("[INFO] received formatted data from method data to TemplateImap_POP3 ")
		d.SetId(name)
		err := go_thunder.PutTemplateImap_POP3(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceTemplateImap_POP3Read(ctx, d, meta)

	}
	return diags
}

func resourceTemplateImap_POP3Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceTemplateImap_POP3Delete) " + name)
		err := go_thunder.DeleteTemplateImap_POP3(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return diags
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToTemplateImapPOP(d *schema.ResourceData) go_thunder.Imap_pop3 {
	var vc go_thunder.Imap_pop3
	var c go_thunder.ImapPop3Instance

	c.Name = d.Get("name").(string)
	c.Logindisabled = d.Get("logindisabled").(int)
	c.Starttls = d.Get("starttls").(string)
	c.UserTag = d.Get("user_tag").(string)

	vc.UUID = c
	return vc
}
