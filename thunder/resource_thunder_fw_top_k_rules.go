package thunder

//Thunder resource FwTopKRules

import (
	"context"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwTopKRules() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwTopKRulesCreate,
		UpdateContext: resourceFwTopKRulesUpdate,
		ReadContext:   resourceFwTopKRulesRead,
		DeleteContext: resourceFwTopKRulesDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwTopKRulesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating FwTopKRules (Inside resourceFwTopKRulesCreate) ")

		data := dataToFwTopKRules(d)
		logger.Println("[INFO] received formatted data from method data to FwTopKRules --")
		d.SetId(strconv.Itoa('1'))
		err := go_thunder.PostFwTopKRules(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwTopKRulesRead(ctx, d, meta)

	}
	return diags
}

func resourceFwTopKRulesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading FwTopKRules (Inside resourceFwTopKRulesRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwTopKRules(client.Token, client.Host)
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

func resourceFwTopKRulesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwTopKRulesRead(ctx, d, meta)
}

func resourceFwTopKRulesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwTopKRulesRead(ctx, d, meta)
}
func dataToFwTopKRules(d *schema.ResourceData) go_thunder.FwTopKRules {
	var vc go_thunder.FwTopKRules
	var c go_thunder.FwTopKRulesInstance

	vc.UUID = c
	return vc
}
