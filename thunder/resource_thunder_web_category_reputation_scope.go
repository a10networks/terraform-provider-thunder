package thunder

//Thunder resource WebCategoryReputationScope

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebCategoryReputationScope() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWebCategoryReputationScopeCreate,
		UpdateContext: resourceWebCategoryReputationScopeUpdate,
		ReadContext:   resourceWebCategoryReputationScopeRead,
		DeleteContext: resourceWebCategoryReputationScopeDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"greater_than": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"greater_trustworthy": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"greater_low_risk": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"greater_moderate_risk": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"greater_suspicious": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"greater_malicious": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"greater_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"less_than": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"less_trustworthy": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"less_low_risk": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"less_moderate_risk": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"less_suspicious": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"less_malicious": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"less_threshold": {
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
		},
	}
}

func resourceWebCategoryReputationScopeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating WebCategoryReputationScope (Inside resourceWebCategoryReputationScopeCreate) ")
		name1 := d.Get("name").(string)
		data := dataToWebCategoryReputationScope(d)
		logger.Println("[INFO] received formatted data from method data to WebCategoryReputationScope --")
		d.SetId(name1)
		err := go_thunder.PostWebCategoryReputationScope(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceWebCategoryReputationScopeRead(ctx, d, meta)

	}
	return diags
}

func resourceWebCategoryReputationScopeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading WebCategoryReputationScope (Inside resourceWebCategoryReputationScopeRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetWebCategoryReputationScope(client.Token, name1, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return nil
}

func resourceWebCategoryReputationScopeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying WebCategoryReputationScope   (Inside resourceWebCategoryReputationScopeUpdate) ")
		data := dataToWebCategoryReputationScope(d)
		logger.Println("[INFO] received formatted data from method data to WebCategoryReputationScope ")
		err := go_thunder.PutWebCategoryReputationScope(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceWebCategoryReputationScopeRead(ctx, d, meta)

	}
	return diags
}

func resourceWebCategoryReputationScopeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceWebCategoryReputationScopeDelete) " + name1)
		err := go_thunder.DeleteWebCategoryReputationScope(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diag.FromErr(err)
		}
		return nil
	}
	return diags
}

func dataToWebCategoryReputationScope(d *schema.ResourceData) go_thunder.WebCategoryReputationScope {
	var vc go_thunder.WebCategoryReputationScope
	var c go_thunder.WebCategoryReputationScopeInstance
	c.Name = d.Get("name").(string)

	var obj1 go_thunder.WebCategoryGreaterThan
	prefix1 := "greater_than.0."
	obj1.GreaterTrustworthy = d.Get(prefix1 + "greater_trustworthy").(int)
	obj1.GreaterLowRisk = d.Get(prefix1 + "greater_low_risk").(int)
	obj1.GreaterModerateRisk = d.Get(prefix1 + "greater_moderate_risk").(int)
	obj1.GreaterSuspicious = d.Get(prefix1 + "greater_suspicious").(int)
	obj1.GreaterMalicious = d.Get(prefix1 + "greater_malicious").(int)
	obj1.GreaterThreshold = d.Get(prefix1 + "greater_threshold").(int)

	c.GreaterTrustworthy = obj1

	var obj2 go_thunder.WebCategoryLessThan
	prefix2 := "less_than.0."
	obj2.LessTrustworthy = d.Get(prefix2 + "less_trustworthy").(int)
	obj2.LessLowRisk = d.Get(prefix2 + "less_low_risk").(int)
	obj2.LessModerateRisk = d.Get(prefix2 + "less_moderate_risk").(int)
	obj2.LessSuspicious = d.Get(prefix2 + "less_suspicious").(int)
	obj2.LessMalicious = d.Get(prefix2 + "less_malicious").(int)
	obj2.LessThreshold = d.Get(prefix2 + "less_threshold").(int)

	c.LessTrustworthy = obj2

	c.UserTag = d.Get("user_tag").(string)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.WebCategoryScopeSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj3 go_thunder.WebCategoryScopeSamplingEnable
		prefix3 := fmt.Sprintf("sampling_enable.%d.", i)
		obj3.Counters1 = d.Get(prefix3 + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj3)
	}

	vc.Name = c
	return vc
}
