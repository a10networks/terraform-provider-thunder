package thunder

//Thunder resource SlbTemplateSmtp

import (
	"context"
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceSlbTemplateSmtp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbTemplateSmtpCreate,
		UpdateContext: resourceSlbTemplateSmtpUpdate,
		ReadContext:   resourceSlbTemplateSmtpRead,
		DeleteContext: resourceSlbTemplateSmtpDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"server_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"service_ready_msg": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"client_starttls_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"server_starttls_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
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
			"lf_to_crlf": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"error_code_to_client": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"client_domain_switching": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"switching_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
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
					},
				},
			},
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

func resourceSlbTemplateSmtpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateSmtp (Inside resourceSlbTemplateSmtpCreate) ")
		name1 := d.Get("name").(string)
		data := dataToSlbTemplateSmtp(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateSmtp --")
		d.SetId(name1)
		err := go_thunder.PostSlbTemplateSmtp(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateSmtpRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateSmtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbTemplateSmtp (Inside resourceSlbTemplateSmtpRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSlbTemplateSmtp(client.Token, name1, client.Host)
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

func resourceSlbTemplateSmtpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying SlbTemplateSmtp   (Inside resourceSlbTemplateSmtpUpdate) ")
		data := dataToSlbTemplateSmtp(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateSmtp ")
		err := go_thunder.PutSlbTemplateSmtp(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateSmtpRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateSmtpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateSmtpDelete) " + name1)
		err := go_thunder.DeleteSlbTemplateSmtp(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToSlbTemplateSmtp(d *schema.ResourceData) go_thunder.SlbTemplateSmtp {
	var vc go_thunder.SlbTemplateSmtp
	var c go_thunder.SlbTemplateSMTPInstance
	c.SlbTemplateSMTPInstanceName = d.Get("name").(string)
	c.SlbTemplateSMTPInstanceServerDomain = d.Get("server_domain").(string)
	c.SlbTemplateSMTPInstanceServiceReadyMsg = d.Get("service_ready_msg").(string)
	c.SlbTemplateSMTPInstanceClientStarttlsType = d.Get("client_starttls_type").(string)
	c.SlbTemplateSMTPInstanceServerStarttlsType = d.Get("server_starttls_type").(string)

	SlbTemplateSMTPInstanceCommandDisableCount := d.Get("command_disable.#").(int)
	c.SlbTemplateSMTPInstanceCommandDisableDisableType = make([]go_thunder.SlbTemplateSMTPInstanceCommandDisable, 0, SlbTemplateSMTPInstanceCommandDisableCount)

	for i := 0; i < SlbTemplateSMTPInstanceCommandDisableCount; i++ {
		var obj1 go_thunder.SlbTemplateSMTPInstanceCommandDisable
		prefix1 := fmt.Sprintf("command_disable.%d.", i)
		obj1.SlbTemplateSMTPInstanceCommandDisableDisableType = d.Get(prefix1 + "disable_type").(string)
		c.SlbTemplateSMTPInstanceCommandDisableDisableType = append(c.SlbTemplateSMTPInstanceCommandDisableDisableType, obj1)
	}

	c.SlbTemplateSMTPInstanceLFToCRLF = d.Get("lf_to_crlf").(int)
	c.SlbTemplateSMTPInstanceErrorCodeToClient = d.Get("error_code_to_client").(int)

	SlbTemplateSMTPInstanceClientDomainSwitchingCount := d.Get("client_domain_switching.#").(int)
	c.SlbTemplateSMTPInstanceClientDomainSwitchingSwitchingType = make([]go_thunder.SlbTemplateSMTPInstanceClientDomainSwitching, 0, SlbTemplateSMTPInstanceClientDomainSwitchingCount)

	for i := 0; i < SlbTemplateSMTPInstanceClientDomainSwitchingCount; i++ {
		var obj2 go_thunder.SlbTemplateSMTPInstanceClientDomainSwitching
		prefix2 := fmt.Sprintf("client_domain_switching.%d.", i)
		obj2.SlbTemplateSMTPInstanceClientDomainSwitchingSwitchingType = d.Get(prefix2 + "switching_type").(string)
		obj2.SlbTemplateSMTPInstanceClientDomainSwitchingMatchString = d.Get(prefix2 + "match_string").(string)
		obj2.SlbTemplateSMTPInstanceClientDomainSwitchingServiceGroup = d.Get(prefix2 + "service_group").(string)
		c.SlbTemplateSMTPInstanceClientDomainSwitchingSwitchingType = append(c.SlbTemplateSMTPInstanceClientDomainSwitchingSwitchingType, obj2)
	}

	var obj3 go_thunder.SlbTemplateSMTPInstanceTemplate
	prefix3 := "template.0."
	obj3.SlbTemplateSMTPInstanceTemplateLogging = d.Get(prefix3 + "logging").(string)

	c.SlbTemplateSMTPInstanceTemplateLogging = obj3

	c.SlbTemplateSMTPInstanceUserTag = d.Get("user_tag").(string)

	vc.SlbTemplateSMTPInstanceName = c
	return vc
}
