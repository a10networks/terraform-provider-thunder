package thunder

//Thunder resource SlbTemplateCache

import (
	"context"
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceSlbTemplateCache() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbTemplateCacheCreate,
		UpdateContext: resourceSlbTemplateCacheUpdate,
		ReadContext:   resourceSlbTemplateCacheRead,
		DeleteContext: resourceSlbTemplateCacheDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"accept_reload_req": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"age": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"default_policy_nocache": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_insert_age": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_insert_via": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"max_cache_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"min_content_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"max_content_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"local_uri_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"local_uri": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"uri_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uri": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"cache_action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"cache_value": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"invalidate": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"remove_cookies": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"replacement_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"logging": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"verify_host": {
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
			"sampling_enable": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"packet_capture_template": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTemplateCacheCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateCache (Inside resourceSlbTemplateCacheCreate) ")
		name1 := d.Get("name").(string)
		data := dataToSlbTemplateCache(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateCache --")
		d.SetId(name1)
		err := go_thunder.PostSlbTemplateCache(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateCacheRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateCacheRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbTemplateCache (Inside resourceSlbTemplateCacheRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSlbTemplateCache(client.Token, name1, client.Host)
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

func resourceSlbTemplateCacheUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying SlbTemplateCache   (Inside resourceSlbTemplateCacheUpdate) ")
		data := dataToSlbTemplateCache(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateCache ")
		err := go_thunder.PutSlbTemplateCache(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateCacheRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateCacheDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateCacheDelete) " + name1)
		err := go_thunder.DeleteSlbTemplateCache(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToSlbTemplateCache(d *schema.ResourceData) go_thunder.SlbTemplateCache {
	var vc go_thunder.SlbTemplateCache
	var c go_thunder.SlbTemplateCacheInstance
	c.SlbTemplateCacheInstanceName = d.Get("name").(string)
	c.SlbTemplateCacheInstanceAcceptReloadReq = d.Get("accept_reload_req").(int)
	c.SlbTemplateCacheInstanceAge = d.Get("age").(int)
	c.SlbTemplateCacheInstanceDefaultPolicyNocache = d.Get("default_policy_nocache").(int)
	c.SlbTemplateCacheInstanceDisableInsertAge = d.Get("disable_insert_age").(int)
	c.SlbTemplateCacheInstanceDisableInsertVia = d.Get("disable_insert_via").(int)
	c.SlbTemplateCacheInstanceMaxCacheSize = d.Get("max_cache_size").(int)
	c.SlbTemplateCacheInstanceMinContentSize = d.Get("min_content_size").(int)
	c.SlbTemplateCacheInstanceMaxContentSize = d.Get("max_content_size").(int)

	SlbTemplateCacheInstanceLocalURIPolicyCount := d.Get("local_uri_policy.#").(int)
	c.SlbTemplateCacheInstanceLocalURIPolicyLocalURI = make([]go_thunder.SlbTemplateCacheInstanceLocalURIPolicy, 0, SlbTemplateCacheInstanceLocalURIPolicyCount)

	for i := 0; i < SlbTemplateCacheInstanceLocalURIPolicyCount; i++ {
		var obj1 go_thunder.SlbTemplateCacheInstanceLocalURIPolicy
		prefix1 := fmt.Sprintf("local_uri_policy.%d.", i)
		obj1.SlbTemplateCacheInstanceLocalURIPolicyLocalURI = d.Get(prefix1 + "local_uri").(string)
		c.SlbTemplateCacheInstanceLocalURIPolicyLocalURI = append(c.SlbTemplateCacheInstanceLocalURIPolicyLocalURI, obj1)
	}

	SlbTemplateCacheInstanceURIPolicyCount := d.Get("uri_policy.#").(int)
	c.SlbTemplateCacheInstanceURIPolicyURI = make([]go_thunder.SlbTemplateCacheInstanceURIPolicy, 0, SlbTemplateCacheInstanceURIPolicyCount)

	for i := 0; i < SlbTemplateCacheInstanceURIPolicyCount; i++ {
		var obj2 go_thunder.SlbTemplateCacheInstanceURIPolicy
		prefix2 := fmt.Sprintf("uri_policy.%d.", i)
		obj2.SlbTemplateCacheInstanceURIPolicyURI = d.Get(prefix2 + "uri").(string)
		obj2.SlbTemplateCacheInstanceURIPolicyCacheAction = d.Get(prefix2 + "cache_action").(string)
		obj2.SlbTemplateCacheInstanceURIPolicyCacheValue = d.Get(prefix2 + "cache_value").(int)
		obj2.SlbTemplateCacheInstanceURIPolicyInvalidate = d.Get(prefix2 + "invalidate").(string)
		c.SlbTemplateCacheInstanceURIPolicyURI = append(c.SlbTemplateCacheInstanceURIPolicyURI, obj2)
	}

	c.SlbTemplateCacheInstanceRemoveCookies = d.Get("remove_cookies").(int)
	c.SlbTemplateCacheInstanceReplacementPolicy = d.Get("replacement_policy").(string)
	c.SlbTemplateCacheInstanceLogging = d.Get("logging").(string)
	c.SlbTemplateCacheInstanceVerifyHost = d.Get("verify_host").(int)
	c.SlbTemplateCacheInstanceUserTag = d.Get("user_tag").(string)

	SlbTemplateCacheInstanceSamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.SlbTemplateCacheInstanceSamplingEnableCounters1 = make([]go_thunder.SlbTemplateCacheInstanceSamplingEnable, 0, SlbTemplateCacheInstanceSamplingEnableCount)

	for i := 0; i < SlbTemplateCacheInstanceSamplingEnableCount; i++ {
		var obj3 go_thunder.SlbTemplateCacheInstanceSamplingEnable
		prefix3 := fmt.Sprintf("sampling_enable.%d.", i)
		obj3.SlbTemplateCacheInstanceSamplingEnableCounters1 = d.Get(prefix3 + "counters1").(string)
		c.SlbTemplateCacheInstanceSamplingEnableCounters1 = append(c.SlbTemplateCacheInstanceSamplingEnableCounters1, obj3)
	}

	c.SlbTemplateCacheInstancePacketCaptureTemplate = d.Get("packet_capture_template").(string)

	vc.SlbTemplateCacheInstanceName = c
	return vc
}
