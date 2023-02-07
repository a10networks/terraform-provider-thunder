package thunder

//Thunder resource FwLocalLog

import (
	"context"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceFwLocalLog() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwLocalLogCreate,
		UpdateContext: resourceFwLocalLogCreate,
		ReadContext:   resourceFwLocalLogRead,
		DeleteContext: resourceFwLocalLogDelete,
		Schema: map[string]*schema.Schema{
			"local_logging": {
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
	}
}

func resourceFwLocalLogCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating FwLocalLog (Inside resourceFwLocalLogCreate) ")

		data := dataToFwLocalLog(d)
		logger.Println("[INFO] received formatted data from method data to FwLocalLog --")
		d.SetId("1")
		err := go_thunder.PostFwLocalLog(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwLocalLogRead(ctx, d, meta)

	}
	return diags
}

func resourceFwLocalLogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading FwLocalLog (Inside resourceFwLocalLogRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetFwLocalLog(client.Token, client.Host)
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

func resourceFwLocalLogUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwLocalLogRead(ctx, d, meta)
}

func resourceFwLocalLogDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwLocalLogRead(ctx, d, meta)
}
func dataToFwLocalLog(d *schema.ResourceData) go_thunder.FwLocalLog {
	var vc go_thunder.FwLocalLog
	var c go_thunder.FwLocalLogInstance
	c.FwLocalLogInstanceLocalLogging = d.Get("local_logging").(int)

	vc.FwLocalLogInstanceLocalLogging = c
	return vc
}
