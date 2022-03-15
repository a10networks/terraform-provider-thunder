package thunder

//Thunder resource Timezone

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTimezone() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTimezoneCreate,
		UpdateContext: resourceTimezoneUpdate,
		ReadContext:   resourceTimezoneRead,
		DeleteContext: resourceTimezoneDelete,
		Schema: map[string]*schema.Schema{
			"timezone_index_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"timezone_index": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"nodst": {
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
		},
	}
}

func resourceTimezoneCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating Timezone (Inside resourceTimezoneCreate) ")

		data := dataToTimezone(d)
		logger.Println("[INFO] received formatted data from method data to Timezone --")
		d.SetId("1")
		err := go_thunder.PostTimezone(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceTimezoneRead(ctx, d, meta)

	}
	return diags
}

func resourceTimezoneRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading Timezone (Inside resourceTimezoneRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetTimezone(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return diags
	}
	return diags
}

func resourceTimezoneUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Updating Timezone (Inside resourceTimezoneCreate) ")

		data := dataToTimezone(d)
		logger.Println("[INFO] received formatted data from method data to Timezone --")
		d.SetId("1")
		err := go_thunder.PutTimezone(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceTimezoneRead(ctx, d, meta)

	}
	return diags
}

func resourceTimezoneDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Deleting Timezone (Inside resourceTimezoneRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		err := go_thunder.DeleteTimezone(client.Token, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete Timezone")
			return diag.FromErr(err)
		}
		return nil
	}
	return diags
}

func dataToTimezone(d *schema.ResourceData) go_thunder.Timezone {
	var vc go_thunder.Timezone
	var c go_thunder.TimezoneInstance

	var obj1 go_thunder.TimezoneInstanceTimezoneIndexCfg
	prefix1 := "timezone_index_cfg.0."
	obj1.TimezoneInstanceTimezoneIndexCfgTimezoneIndex = d.Get(prefix1 + "timezone_index").(string)
	obj1.TimezoneInstanceTimezoneIndexCfgNodst = d.Get(prefix1 + "nodst").(int)

	c.TimezoneInstanceTimezoneIndexCfgTimezoneIndex = obj1

	vc.TimezoneInstanceTimezoneIndexCfg = c
	return vc
}
