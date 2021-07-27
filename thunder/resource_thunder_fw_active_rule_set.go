package thunder

//Thunder resource FwActiveRuleSet

import (
	"context"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwActiveRuleSet() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwActiveRuleSetCreate,
		UpdateContext: resourceFwActiveRuleSetUpdate,
		ReadContext:   resourceFwActiveRuleSetRead,
		DeleteContext: resourceFwActiveRuleSetDelete,
		Schema: map[string]*schema.Schema{
			"session_aging": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"override_nat_aging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
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
		},
	}
}

func resourceFwActiveRuleSetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating FwActiveRuleSet (Inside resourceFwActiveRuleSetCreate) ")

		data := dataToFwActiveRuleSet(d)
		logger.Println("[INFO] received formatted data from method data to FwActiveRuleSet --")
		d.SetId(strconv.Itoa('1'))
		err := go_thunder.PostFwActiveRuleSet(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwActiveRuleSetRead(ctx, d, meta)

	}
	return diags
}

func resourceFwActiveRuleSetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading FwActiveRuleSet (Inside resourceFwActiveRuleSetRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwActiveRuleSet(client.Token, client.Host)
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

func resourceFwActiveRuleSetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwActiveRuleSetRead(ctx, d, meta)
}

func resourceFwActiveRuleSetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwActiveRuleSetRead(ctx, d, meta)
}
func dataToFwActiveRuleSet(d *schema.ResourceData) go_thunder.FwActiveRuleSet {
	var vc go_thunder.FwActiveRuleSet
	var c go_thunder.FwActiveRuleSetInstance
	c.SessionAging = d.Get("session_aging").(string)
	c.OverrideNatAging = d.Get("override_nat_aging").(int)
	c.Name = d.Get("name").(string)

	vc.SessionAging = c
	return vc
}
