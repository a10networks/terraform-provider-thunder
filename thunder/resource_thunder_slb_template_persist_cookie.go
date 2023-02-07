package thunder

//Thunder resource SlbTemplatePersistCookie

import (
	"context"
	"github.com/go_thunder/thunder"
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
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"dont_honor_conn_rules": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"expire": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"insert_always": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"encrypt_level": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"pass_phrase": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"cookie_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"prefix": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"samesite": {
				Type:        schema.TypeString,
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
			"secure": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"httponly": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"match_type": {
				Type:        schema.TypeInt,
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
			"scan_all_members": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"service_group": {
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
	logger.Println("[INFO] Reading SlbTemplatePersistCookie (Inside resourceSlbTemplatePersistCookieRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSlbTemplatePersistCookie(client.Token, name1, client.Host)
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
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplatePersistCookieDelete) " + name1)
		err := go_thunder.DeleteSlbTemplatePersistCookie(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToSlbTemplatePersistCookie(d *schema.ResourceData) go_thunder.SlbTemplatePersistCookie {
	var vc go_thunder.SlbTemplatePersistCookie
	var c go_thunder.SlbTemplatePersistCookieInstance
	c.SlbTemplatePersistCookieInstanceName = d.Get("name").(string)
	c.SlbTemplatePersistCookieInstanceDontHonorConnRules = d.Get("dont_honor_conn_rules").(int)
	c.SlbTemplatePersistCookieInstanceExpire = d.Get("expire").(int)
	c.SlbTemplatePersistCookieInstanceInsertAlways = d.Get("insert_always").(int)
	c.SlbTemplatePersistCookieInstanceEncryptLevel = d.Get("encrypt_level").(int)
	c.SlbTemplatePersistCookieInstancePassPhrase = d.Get("pass_phrase").(string)
	c.SlbTemplatePersistCookieInstanceCookieName = d.Get("cookie_name").(string)
	c.SlbTemplatePersistCookieInstancePrefix = d.Get("prefix").(string)
	c.SlbTemplatePersistCookieInstanceDomain = d.Get("domain").(string)
	c.SlbTemplatePersistCookieInstanceSamesite = d.Get("samesite").(string)
	c.SlbTemplatePersistCookieInstancePath = d.Get("path").(string)
	c.SlbTemplatePersistCookieInstancePassThru = d.Get("pass_thru").(int)
	c.SlbTemplatePersistCookieInstanceSecure = d.Get("secure").(int)
	c.SlbTemplatePersistCookieInstanceHttponly = d.Get("httponly").(int)
	c.SlbTemplatePersistCookieInstanceMatchType = d.Get("match_type").(int)
	c.SlbTemplatePersistCookieInstanceServer = d.Get("server").(int)
	c.SlbTemplatePersistCookieInstanceServerServiceGroup = d.Get("server_service_group").(int)
	c.SlbTemplatePersistCookieInstanceScanAllMembers = d.Get("scan_all_members").(int)
	c.SlbTemplatePersistCookieInstanceServiceGroup = d.Get("service_group").(int)
	c.SlbTemplatePersistCookieInstanceUserTag = d.Get("user_tag").(string)

	vc.SlbTemplatePersistCookieInstanceName = c
	return vc
}
