package thunder

//Thunder resource FwRadiusServer

import (
	"context"
	"fmt"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwRadiusServer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwRadiusServerCreate,
		UpdateContext: resourceFwRadiusServerUpdate,
		ReadContext:   resourceFwRadiusServerRead,
		DeleteContext: resourceFwRadiusServerDelete,
		Schema: map[string]*schema.Schema{
			"accounting_start": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"attribute_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"vrid": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"remote": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_list_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ip_list_encrypted": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ip_list_secret_string": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ip_list_secret": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"encrypted": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"accounting_interim_update": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"secret": {
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
			"accounting_stop": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"custom_attribute_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"attribute": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"prefix_number": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"prefix_length": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"prefix_vendor": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"number": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"custom_vendor": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"custom_number": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"vendor": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"attribute_value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"listen_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"accounting_on": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"secret_string": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwRadiusServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating FwRadiusServer (Inside resourceFwRadiusServerCreate) ")

		data := dataToFwRadiusServer(d)
		logger.Println("[INFO] received formatted data from method data to FwRadiusServer --")
		d.SetId(strconv.Itoa('1'))
		err := go_thunder.PostFwRadiusServer(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwRadiusServerRead(ctx, d, meta)

	}
	return diags
}

func resourceFwRadiusServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading FwRadiusServer (Inside resourceFwRadiusServerRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwRadiusServer(client.Token, client.Host)
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

func resourceFwRadiusServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwRadiusServerRead(ctx, d, meta)
}

func resourceFwRadiusServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwRadiusServerRead(ctx, d, meta)
}
func dataToFwRadiusServer(d *schema.ResourceData) go_thunder.FwRadiusServer {
	var vc go_thunder.FwRadiusServer
	var c go_thunder.FwRadiusServerInstance
	c.AccountingStart = d.Get("accounting_start").(string)
	c.AttributeName = d.Get("attribute_name").(string)
	c.Vrid = d.Get("vrid").(int)

	var obj1 go_thunder.FwRadiusServerRemote
	prefix := "remote.0."

	IpListCount := d.Get(prefix + "ip_list.#").(int)
	obj1.IPListName = make([]go_thunder.FwRadiusServerIPList, 0, IpListCount)

	for i := 0; i < IpListCount; i++ {
		var obj1_1 go_thunder.FwRadiusServerIPList
		prefix1 := prefix + fmt.Sprintf("ip_list.%d.", i)
		obj1_1.IPListName = d.Get(prefix1 + "ip_list_name").(string)
		obj1_1.IPListEncrypted = d.Get(prefix1 + "ip_list_encrypted").(string)
		obj1_1.IPListSecretString = d.Get(prefix1 + "ip_list_secret_string").(string)
		obj1_1.IPListSecret = d.Get(prefix1 + "ip_list_secret").(int)
		obj1.IPListName = append(obj1.IPListName, obj1_1)
	}

	c.IPList = obj1

	c.Encrypted = d.Get("encrypted").(string)
	c.AccountingInterimUpdate = d.Get("accounting_interim_update").(string)
	c.Secret = d.Get("secret").(int)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.FwRadiusServerSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.FwRadiusServerSamplingEnable
		prefix2 := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix2 + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	c.AccountingStop = d.Get("accounting_stop").(string)
	c.CustomAttributeName = d.Get("custom_attribute_name").(string)

	AttributeCount := d.Get("attribute.#").(int)
	c.PrefixNumber = make([]go_thunder.FwRadiusServerAttribute, 0, AttributeCount)

	for i := 0; i < AttributeCount; i++ {
		var obj1 go_thunder.FwRadiusServerAttribute
		prefix3 := fmt.Sprintf("attribute.%d.", i)
		obj1.PrefixNumber = d.Get(prefix3 + "prefix_number").(int)
		obj1.PrefixLength = d.Get(prefix3 + "prefix_length").(string)
		obj1.Name = d.Get(prefix3 + "name").(string)
		obj1.PrefixVendor = d.Get(prefix3 + "prefix_vendor").(int)
		obj1.Number = d.Get(prefix3 + "number").(int)
		obj1.Value = d.Get(prefix3 + "value").(string)
		obj1.CustomVendor = d.Get(prefix3 + "custom_vendor").(int)
		obj1.CustomNumber = d.Get(prefix3 + "custom_number").(int)
		obj1.Vendor = d.Get(prefix3 + "vendor").(int)
		obj1.AttributeValue = d.Get(prefix3 + "attribute_value").(string)
		c.PrefixNumber = append(c.PrefixNumber, obj1)
	}

	c.ListenPort = d.Get("listen_port").(int)
	c.AccountingOn = d.Get("accounting_on").(string)
	c.SecretString = d.Get("secret_string").(string)

	vc.AccountingStart = c
	return vc
}
