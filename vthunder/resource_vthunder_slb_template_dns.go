package vthunder

//vThunder resource TemplateDNS

import (
	"fmt"
	"log"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceTemplateDNS() *schema.Resource {
	return &schema.Resource{
		Create: resourceTemplateDNSCreate,
		Update: resourceTemplateDNSUpdate,
		Read:   resourceTemplateDNSRead,
		Delete: resourceTemplateDNSDelete,

		Schema: map[string]*schema.Schema{
			"drop": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"max_query_length": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"query_id_switch": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"period": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"response_rate_limiting": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"response_rate": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"slip_rate": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"filter_response_rate": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"window": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"enable_log": {
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
				},
			},
			"forward": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"dnssec_service_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"max_cache_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"max_cache_entry_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"disable_dns_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"enable_cache_sharing": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"redirect_to_tcp_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"class_list": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"lid_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"log_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"lockout": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"user_tag": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"log": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"over_limit_action": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"conn_rate_limit": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"action_value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"per": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"lidnum": {
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
							},
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"default_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceTemplateDNSCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating TemplateDNS (Inside resourceTemplateDNSCreate) ")
		name := d.Get("name").(string)
		data := dataToTemplateDNS(d)
		logger.Println("[INFO] received V from method data to TemplateDNS --")
		d.SetId(name)
		go_vthunder.PostTemplateDNS(client.Token, data, client.Host)

		return resourceTemplateDNSRead(d, meta)

	}
	return nil
}

func resourceTemplateDNSRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading TemplateDNS (Inside resourceTemplateDNSRead)")

	if client.Host != "" {

		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetTemplateDNS(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceTemplateDNSUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying TemplateDNS   (Inside resourceTemplateDNSUpdate) ")
		name := d.Get("name").(string)
		data := dataToTemplateDNS(d)
		logger.Println("[INFO] received V from method data to TemplateDNS ")
		d.SetId(name)
		go_vthunder.PutTemplateDNS(client.Token, name, data, client.Host)

		return resourceTemplateDNSRead(d, meta)

	}
	return nil
}

func resourceTemplateDNSDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceTemplateDNSDelete) " + name)
		err := go_vthunder.DeleteTemplateDNS(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToTemplateDNS(d *schema.ResourceData) go_vthunder.DNS {
	var vc go_vthunder.DNS
	var c go_vthunder.DNSInstance

	c.Name = d.Get("name").(string)
	c.DnssecServiceGroup = d.Get("dnssec_service_group").(string)
	var classListObj go_vthunder.ClassList

	lidListCount := d.Get("class_list.0.lid_list.#").(int)
	classListObj.ActionValue = make([]go_vthunder.LidList, 0, lidListCount)

	for i := 0; i < lidListCount; i++ {
		var cl go_vthunder.LidList
		prefix := fmt.Sprintf("class_list.0.lid_list.%d", i)
		cl.ActionValue = d.Get(prefix + ".action_value").(string)
		cl.ConnRateLimit = d.Get(prefix + ".conn_rate_limit").(int)
		cl.Lidnum = d.Get(prefix + ".lidnum").(int)
		cl.Lockout = d.Get(prefix + ".lockout").(int)
		cl.Log = d.Get(prefix + ".log").(int)
		cl.LogInterval = d.Get(prefix + ".log_interval").(int)
		cl.OverLimitAction = d.Get(prefix + ".over_limit_action").(int)
		cl.UserTag = d.Get(prefix + ".user_tag").(string)
		cl.Per = d.Get(prefix + ".per").(int)
		classListObj.ActionValue = append(classListObj.ActionValue, cl)
	}
	c.LidList = classListObj
	var RateLimiting go_vthunder.ResponseRateLimiting
	RateLimiting.Action = d.Get("response_rate_limiting.0.action").(string)
	RateLimiting.SlipRate = d.Get("response_rate_limiting.0.slip_rate").(int)
	RateLimiting.EnableLog = d.Get("response_rate_limiting.0.enable_log").(int)
	RateLimiting.FilterResponseRate = d.Get("response_rate_limiting.0.filter_response_rate").(int)
	RateLimiting.Window = d.Get("response_rate_limiting.0.window").(int)
	RateLimiting.ResponseRate = d.Get("response_rate_limiting.0.response_rate").(int)
	RateLimiting.UUID = d.Get("response_rate_limiting.0.uuid").(string)
	c.FilterResponseRate = RateLimiting
	c.DefaultPolicy = d.Get("default_policy").(string)
	c.DisableDNSTemplate = d.Get("disable_dns_template").(int)
	c.Drop = d.Get("drop").(int)
	c.EnableCacheSharing = d.Get("enable_cache_sharing").(int)
	c.Forward = d.Get("forward").(string)
	c.MaxCacheEntrySize = d.Get("max_cache_entry_size").(int)
	c.MaxCacheSize = d.Get("max_cache_size").(int)
	c.MaxQueryLength = d.Get("max_query_length").(int)
	c.Period = d.Get("period").(int)
	c.QueryIDSwitch = d.Get("query_id_switch").(int)
	c.RedirectToTCPPort = d.Get("redirect_to_tcp_port").(int)
	c.UserTag = d.Get("user_tag").(string)

	vc.Name = c
	return vc
}
