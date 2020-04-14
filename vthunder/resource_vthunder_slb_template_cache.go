package vthunder

//vThunder resource SlbTemplateCache

import (
	"fmt"
	"log"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbTemplateCache() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbTemplateCacheCreate,
		Update: resourceSlbTemplateCacheUpdate,
		Read:   resourceSlbTemplateCacheRead,
		Delete: resourceSlbTemplateCacheDelete,
		Schema: map[string]*schema.Schema{
			"verify_host": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"accept_reload_req": {
				Type:        schema.TypeInt,
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
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"max_content_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"default_policy_nocache": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_insert_via": {
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
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
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
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uri_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"cache_action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"uri": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"logging": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"age": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_insert_age": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTemplateCacheCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateCache (Inside resourceSlbTemplateCacheCreate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateCache(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateCache --")
		d.SetId(name)
		go_vthunder.PostSlbTemplateCache(client.Token, data, client.Host)

		return resourceSlbTemplateCacheRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateCacheRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbTemplateCache (Inside resourceSlbTemplateCacheRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetSlbTemplateCache(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbTemplateCacheUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplateCache   (Inside resourceSlbTemplateCacheUpdate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateCache(d)
		logger.Println("[INFO] received V from method data to SlbTemplateCache ")
		d.SetId(name)
		go_vthunder.PutSlbTemplateCache(client.Token, name, data, client.Host)

		return resourceSlbTemplateCacheRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateCacheDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateCacheDelete) " + name)
		err := go_vthunder.DeleteSlbTemplateCache(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbTemplateCache(d *schema.ResourceData) go_vthunder.Cache {
	var vc go_vthunder.Cache
	var c go_vthunder.CacheInstance

	c.AcceptReloadReq = d.Get("accept_reload_req").(int)
	c.Name = d.Get("name").(string)
	c.DefaultPolicyNocache = d.Get("default_policy_nocache").(int)
	c.Age = d.Get("age").(int)
	c.DisableInsertVia = d.Get("disable_insert_via").(int)
	c.UserTag = d.Get("user_tag").(string)
	c.ReplacementPolicy = d.Get("replacement_policy").(string)
	c.DisableInsertAge = d.Get("disable_insert_age").(int)
	c.MaxContentSize = d.Get("max_content_size").(int)
	c.MaxCacheSize = d.Get("max_cache_size").(int)
	c.Logging = d.Get("logging").(string)
	c.RemoveCookies = d.Get("remove_cookies").(int)
	c.VerifyHost = d.Get("verify_host").(int)
	c.MinContentSize = d.Get("min_content_size").(int)
	LocalURIPolicyCount := d.Get("local_uri_policy.#").(int)
	c.LocalURI = make([]go_vthunder.LocalURIPolicy, 0, LocalURIPolicyCount)
	for i := 0; i < LocalURIPolicyCount; i++ {
		var lup go_vthunder.LocalURIPolicy
		prefix := fmt.Sprintf("local_uri_policy.%d", i)
		lup.LocalURI = d.Get(prefix + ".local_uri").(string)
		c.LocalURI = append(c.LocalURI, lup)
	}
	SamplingEnable4Count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable4, 0, SamplingEnable4Count)
	for i := 0; i < SamplingEnable4Count; i++ {
		var se go_vthunder.SamplingEnable4
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}
	URIPolicyCount := d.Get("uri_policy.#").(int)
	c.URI = make([]go_vthunder.URIPolicy, 0, URIPolicyCount)
	for i := 0; i < URIPolicyCount; i++ {
		var up go_vthunder.URIPolicy
		prefix := fmt.Sprintf("uri_policy.%d", i)
		up.CacheAction = d.Get(prefix + ".cache_action").(string)
		up.CacheValue = d.Get(prefix + ".cache_value").(int)
		up.Invalidate = d.Get(prefix + ".invalidate").(string)
		up.URI = d.Get(prefix + ".uri").(string)
		c.URI = append(c.URI, up)
	}

	vc.UUID = c
	return vc
}
