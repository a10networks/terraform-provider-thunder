package thunder

//Thunder resource SlbTemplatePersistCookie

import (
	"context"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"util"
)

func resourceSlbTemplatePersistCookie() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbTemplatePersistCookieCreate,
		UpdateContext: resourceSlbTemplatePersistCookieUpdate,
		ReadContext:   resourceSlbTemplatePersistCookieRead,
		DeleteContext: resourceSlbTemplatePersistCookieDelete,
		Schema: map[string]*schema.Schema{
			"pass_phrase": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"cookie_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"secure": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"encrypted": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"dont_honor_conn_rules": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"encrypt_level": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
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
			"server": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"server_service_group": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"service_group": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"expire": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"httponly": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"pass_thru": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"scan_all_members": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"insert_always": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"match_type": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTemplatePersistCookieCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplatePersistCookie (Inside resourceSlbTemplatePersistCookieCreate) ")
		name1 := d.Get("name").(string)
		data := dataToSlbTemplatePersistCookie(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplatePersistCookie --")
		d.SetId(name1)
		err := go_thunder.PostSlbTemplatePersistCookie(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplatePersistCookieRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplatePersistCookieRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbTemplatePersistCookie (Inside resourceSlbTemplatePersistCookieRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read SlbTemplatePersistCookie")
		data, err := go_thunder.GetSlbTemplatePersistCookie(client.Token, name1, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found SlbTemplatePersistCookie")
			return nil
		}
		return diags
	}
	return nil
}

func resourceSlbTemplatePersistCookieUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying SlbTemplatePersistCookie   (Inside resourceSlbTemplatePersistCookieUpdate) ")
		data := dataToSlbTemplatePersistCookie(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplatePersistCookie ")
		err := go_thunder.PutSlbTemplatePersistCookie(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplatePersistCookieRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplatePersistCookieDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplatePersistCookieDelete) ")
		err := go_thunder.DeleteSlbTemplatePersistCookie(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance SlbTemplatePersistCookie (%s) (%v)", name1, err)
			return diag.FromErr(err)
		}
		return nil
	}
	return diags
}

func dataToSlbTemplatePersistCookie(d *schema.ResourceData) go_thunder.SlbTemplatePersistCookie {
	var vc go_thunder.SlbTemplatePersistCookie
	var c go_thunder.SlbTemplatePersistCookieInstance
	c.PassPhrase = d.Get("pass_phrase").(string)
	c.Domain = d.Get("domain").(string)
	c.CookieName = d.Get("cookie_name").(string)
	c.Secure = d.Get("secure").(int)
	c.Encrypted = d.Get("encrypted").(string)
	c.DontHonorConnRules = d.Get("dont_honor_conn_rules").(int)
	c.EncryptLevel = d.Get("encrypt_level").(int)
	c.UserTag = d.Get("user_tag").(string)
	c.Server = d.Get("server").(int)
	c.ServerServiceGroup = d.Get("server_service_group").(int)
	c.ServiceGroup = d.Get("service_group").(int)
	c.Expire = d.Get("expire").(int)
	c.Httponly = d.Get("httponly").(int)
	c.Path = d.Get("path").(string)
	c.PassThru = d.Get("pass_thru").(int)
	c.ScanAllMembers = d.Get("scan_all_members").(int)
	c.InsertAlways = d.Get("insert_always").(int)
	c.MatchType = d.Get("match_type").(int)
	c.Name = d.Get("name").(string)

	vc.PassPhrase = c
	return vc
}
