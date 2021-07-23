package thunder

//Thunder resource logging

import (
	"context"
	"log"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateLogging() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceLoggingCreate,
		UpdateContext: resourceLoggingUpdate,
		ReadContext:   resourceLoggingRead,
		DeleteContext: resourceLoggingDelete,

		Schema: map[string]*schema.Schema{
			"template_tcp_proxy_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"auto": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_tcp_proxy_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"pcre_mask": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"format": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"pool": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"keep_start": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"service_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_pool": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"local_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"pool_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"tcp_proxy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"keep_end": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"mask": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceLoggingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Creating Logging (Inside resourceLoggingCreate    " + name)
		v := dataToLogging(name, d)
		d.SetId(name)
		err := go_thunder.PostLogging(client.Token, v, client.Host)
if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingRead(ctx, d, meta)
	}
	return diags
}

func resourceLoggingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading Logging (Inside resourcLoggingRead)")

	if client.Host != "" {
		client := meta.(Thunder)


		name := d.Id()

		logger.Println("[INFO] Fetching Logging Read" + name)

		logging, err := go_thunder.GetLogging(client.Token, name, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if logging == nil {
			logger.Println("[INFO] No Logging found " + name)
			d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceLoggingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Modifying Logging (Inside resourceLoggingUpdate    " + name)
		v := dataToLogging(name, d)
		d.SetId(name)
		err := go_thunder.PutLogging(client.Token, name, v, client.Host)
if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingRead(ctx, d, meta)
	}
	return diags
}

func resourceLoggingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger := util.GetLoggerInstance()

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting Logging (Inside resourceLoggingDelete) " + name)

		err := go_thunder.DeleteLogging(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete Logging  (%s) (%v)", name, err)
			return diags
		}
		d.SetId("")
		return nil
	}
	return nil
}

//utility method to instantiate Logging structure
func dataToLogging(name string, d *schema.ResourceData) go_thunder.Logging {
	var s go_thunder.Logging

	var sInstance go_thunder.LoggingInstance

	sInstance.PoolShared = d.Get("pool_shared").(string)
	sInstance.Name = d.Get("name").(string)
	sInstance.Format = d.Get("format").(string)
	sInstance.Auto = d.Get("auto").(string)
	sInstance.KeepEnd = d.Get("keep_end").(int)
	sInstance.LocalLogging = d.Get("local_logging").(int)
	sInstance.Mask = d.Get("mask").(string)
	sInstance.TemplateTCPProxyShared = d.Get("template_tcp_proxy_shared").(string)
	sInstance.SharedPartitionTCPProxyTemplate = d.Get("shared_partition_tcp_proxy_template").(int)
	sInstance.KeepStart = d.Get("keep_start").(int)
	sInstance.ServiceGroup = d.Get("service_group").(string)
	sInstance.PcreMask = d.Get("pcre_mask").(string)
	sInstance.UserTag = d.Get("user_tag").(string)
	sInstance.TCPProxy = d.Get("tcp_proxy").(string)
	sInstance.SharedPartitionPool = d.Get("shared_partition_pool").(int)
	sInstance.Pool = d.Get("pool").(string)

	s.Name = sInstance

	return s
}
