package thunder

//Thunder resource SlbTemplateSMTP

import (
	"context"
	"fmt"
	"log"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateSMTP() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbTemplateSMTPCreate,
		UpdateContext: resourceSlbTemplateSMTPUpdate,
		ReadContext:   resourceSlbTemplateSMTPRead,
		DeleteContext: resourceSlbTemplateSMTPDelete,
		Schema: map[string]*schema.Schema{
			"template": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logging": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"command_disable": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disable_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"service_ready_msg": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"client_domain_switching": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"match_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"service_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"switching_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
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
			"server_starttls_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"server_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"client_starttls_type": {
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

func resourceSlbTemplateSMTPCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateSMTP (Inside resourceSlbTemplateSMTPCreate) ")
		name := d.Get("name").(string)
		data := dataToTemplateSMTP(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateSMTP --")
		d.SetId(name)
		err := go_thunder.PostSlbTemplateSMTP(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateSMTPRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateSMTPRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbTemplateSMTP (Inside resourceSlbTemplateSMTPRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetSlbTemplateSMTP(client.Token, name, client.Host)
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

func resourceSlbTemplateSMTPUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplateSMTP   (Inside resourceSlbTemplateSMTPUpdate) ")
		name := d.Get("name").(string)
		data := dataToTemplateSMTP(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateSMTP ")
		d.SetId(name)
		err := go_thunder.PutSlbTemplateSMTP(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateSMTPRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateSMTPDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateSMTPDelete) " + name)
		err := go_thunder.DeleteSlbTemplateSMTP(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return diags
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToTemplateSMTP(d *schema.ResourceData) go_thunder.SMTP {
	var vc go_thunder.SMTP
	var c go_thunder.SMTPInstance

	clientDomainSwitchingCount := d.Get("client_domain_switching.#").(int)
	c.ServiceGroup = make([]go_thunder.ClientDomainSwitching, 0, clientDomainSwitchingCount)
	for i := 0; i < clientDomainSwitchingCount; i++ {
		var cds go_thunder.ClientDomainSwitching
		prefix := fmt.Sprintf("client_domain_switching.%d", i)
		cds.MatchString = d.Get(prefix + ".match_string").(string)
		cds.ServiceGroup = d.Get(prefix + "service_group").(string)
		cds.SwitchingType = d.Get(prefix + "switching_type").(string)
		c.ServiceGroup = append(c.ServiceGroup, cds)
	}
	commandDisableCount := d.Get("command_disable.#").(int)
	c.DisableType = make([]go_thunder.CommandDisable, 0, commandDisableCount)
	for i := 0; i < commandDisableCount; i++ {
		var cd go_thunder.CommandDisable
		prefix := fmt.Sprintf("command_disable.%d", i)
		cd.DisableType = d.Get(prefix + ".disable_type").(string)
		c.DisableType = append(c.DisableType, cd)
	}

	var temp go_thunder.Template1
	temp.Logging = d.Get("template.0.logging").(string)
	c.Logging = temp
	c.Name = d.Get("name").(string)
	c.ServerDomain = d.Get("server_domain").(string)
	c.ServerStarttlsType = d.Get("server_starttls_type").(string)
	c.ClientStarttlsType = d.Get("client_starttls_type").(string)
	c.ServiceReadyMsg = d.Get("service_ready_msg").(string)
	c.UserTag = d.Get("user_tag").(string)

	vc.UUID = c
	return vc
}
