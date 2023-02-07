package thunder

//Thunder resource SlbSSLExpireCheck

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSSLExpireCheck() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbSSLExpireCheckCreate,
		UpdateContext: resourceSlbSSLExpireCheckUpdate,
		ReadContext:   resourceSlbSSLExpireCheckRead,
		DeleteContext: resourceSlbSSLExpireCheckDelete,
		Schema: map[string]*schema.Schema{
			"exception": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"certificate_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"expire_address1": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"before": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"interval_days": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ssl_expire_email_address": {
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

func resourceSlbSSLExpireCheckCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbSSLExpireCheck (Inside resourceSlbSSLExpireCheckCreate) ")

		data := dataToSlbSSLExpireCheck(d)
		logger.Println("[INFO] received SSL Expire Check from method data to SlbSSLExpireCheck --")
		d.SetId("1")
		err := go_thunder.PostSlbSSLExpireCheck(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbSSLExpireCheckRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbSSLExpireCheckRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbSSLExpireCheck (Inside resourceSlbSSLExpireCheckRead)")

	if client.Host != "" {
		data, err := go_thunder.GetSlbSSLExpireCheck(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			d.SetId("")
			return nil
		}
		return diags
	}
	return nil
}

func resourceSlbSSLExpireCheckUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbSSLExpireCheckRead(ctx, d, meta)
}

func resourceSlbSSLExpireCheckDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbSSLExpireCheckRead(ctx, d, meta)
}
func dataToSlbSSLExpireCheck(d *schema.ResourceData) go_thunder.SSLExpireCheck {
	var vc go_thunder.SSLExpireCheck
	var c go_thunder.SSLExpireCheckInstance
	c.SslExpireEmailAddress = d.Get("ssl_expire_email_address").(string)

	var obj1 go_thunder.Exception
	prefix := "exception.0."
	obj1.Action = d.Get(prefix + "action").(string)
	obj1.CertificateName = d.Get(prefix + "certificate_name").(string)
	c.Action = obj1

	c.ExpireAddress1 = d.Get("expire_address1").(string)
	c.IntervalDays = d.Get("interval_days").(int)
	c.Before = d.Get("before").(int)
	vc.UUID = c
	return vc
}
