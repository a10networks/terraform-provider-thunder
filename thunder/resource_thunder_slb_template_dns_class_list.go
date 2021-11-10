package thunder

//Thunder resource SlbTemplateDnsClassList

import (
	"context"
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceSlbTemplateDnsClassList() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbTemplateDnsClassListCreate,
		UpdateContext: resourceSlbTemplateDnsClassListUpdate,
		ReadContext:   resourceSlbTemplateDnsClassListRead,
		DeleteContext: resourceSlbTemplateDnsClassListDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"lid_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"lidnum": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"conn_rate_limit": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"per": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"over_limit_action": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"action_value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"lockout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"log": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"log_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"dns": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cache_action": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ttl": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"weight": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"honor_server_response_ttl": {
										Type:        schema.TypeInt,
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
				},
			},
		},
	}
}

func resourceSlbTemplateDnsClassListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateDnsClassList (Inside resourceSlbTemplateDnsClassListCreate) ")
		name1 := d.Get("name").(string)
		data := dataToSlbTemplateDnsClassList(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateDnsClassList --")
		d.SetId(name1)
		err := go_thunder.PostSlbTemplateDnsClassList(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateDnsClassListRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateDnsClassListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbTemplateDnsClassList (Inside resourceSlbTemplateDnsClassListRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSlbTemplateDnsClassList(client.Token, name1, client.Host)
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

func resourceSlbTemplateDnsClassListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbTemplateDnsClassListRead(ctx, d, meta)
}

func resourceSlbTemplateDnsClassListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbTemplateDnsClassListRead(ctx, d, meta)
}
func dataToSlbTemplateDnsClassList(d *schema.ResourceData) go_thunder.SlbTemplateDnsClassList {
	var vc go_thunder.SlbTemplateDnsClassList
	var c go_thunder.SlbTemplateDNSClassListInstance
	c.SlbTemplateDNSClassListInstanceName = d.Get("name").(string)

	SlbTemplateDNSClassListInstanceLidListCount := d.Get("lid_list.#").(int)
	c.SlbTemplateDNSClassListInstanceLidListLidnum = make([]go_thunder.SlbTemplateDNSClassListInstanceLidList, 0, SlbTemplateDNSClassListInstanceLidListCount)

	for i := 0; i < SlbTemplateDNSClassListInstanceLidListCount; i++ {
		var obj1 go_thunder.SlbTemplateDNSClassListInstanceLidList
		prefix1 := fmt.Sprintf("lid_list.%d.", i)
		obj1.SlbTemplateDNSClassListInstanceLidListLidnum = d.Get(prefix1 + "lidnum").(int)
		obj1.SlbTemplateDNSClassListInstanceLidListConnRateLimit = d.Get(prefix1 + "conn_rate_limit").(int)
		obj1.SlbTemplateDNSClassListInstanceLidListPer = d.Get(prefix1 + "per").(int)
		obj1.SlbTemplateDNSClassListInstanceLidListOverLimitAction = d.Get(prefix1 + "over_limit_action").(int)
		obj1.SlbTemplateDNSClassListInstanceLidListActionValue = d.Get(prefix1 + "action_value").(string)
		obj1.SlbTemplateDNSClassListInstanceLidListLockout = d.Get(prefix1 + "lockout").(int)
		obj1.SlbTemplateDNSClassListInstanceLidListLog = d.Get(prefix1 + "log").(int)
		obj1.SlbTemplateDNSClassListInstanceLidListLogInterval = d.Get(prefix1 + "log_interval").(int)

		var obj1_1 go_thunder.SlbTemplateDNSClassListInstanceLidListDNS
		prefix1_1 := prefix1 + "dns.0."
		obj1_1.SlbTemplateDNSClassListInstanceLidListDNSCacheAction = d.Get(prefix1_1 + "cache_action").(string)
		obj1_1.SlbTemplateDNSClassListInstanceLidListDNSTTL = d.Get(prefix1_1 + "ttl").(int)
		obj1_1.SlbTemplateDNSClassListInstanceLidListDNSWeight = d.Get(prefix1_1 + "weight").(int)
		obj1_1.SlbTemplateDNSClassListInstanceLidListDNSHonorServerResponseTTL = d.Get(prefix1_1 + "honor_server_response_ttl").(int)

		obj1.SlbTemplateDNSClassListInstanceLidListDNSCacheAction = obj1_1

		obj1.SlbTemplateDNSClassListInstanceLidListUserTag = d.Get(prefix1 + "user_tag").(string)
		c.SlbTemplateDNSClassListInstanceLidListLidnum = append(c.SlbTemplateDNSClassListInstanceLidListLidnum, obj1)
	}

	vc.SlbTemplateDNSClassListInstanceName = c
	return vc
}
