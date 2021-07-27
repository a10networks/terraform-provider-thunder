package thunder

//Thunder resource FwLocalLog

import (
	"context"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwLocalLog() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwLocalLogCreate,
		UpdateContext: resourceFwLocalLogUpdate,
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
		d.SetId(strconv.Itoa('1'))
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

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading FwLocalLog (Inside resourceFwLocalLogRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwLocalLog(client.Token, client.Host)
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

func resourceFwLocalLogUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwLocalLogRead(ctx, d, meta)
}

func resourceFwLocalLogDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwLocalLogRead(ctx, d, meta)
}
func dataToFwLocalLog(d *schema.ResourceData) go_thunder.FwLocalLog {
	var vc go_thunder.FwLocalLog
	var c go_thunder.FwLocalLogInstance
	c.LocalLogging = d.Get("local_logging").(int)

	vc.LocalLogging = c
	return vc
}
