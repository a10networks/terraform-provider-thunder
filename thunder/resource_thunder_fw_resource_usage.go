package thunder

//Thunder resource FwResourceUsage

import (
	"context"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwResourceUsage() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwResourceUsageCreate,
		UpdateContext: resourceFwResourceUsageUpdate,
		ReadContext:   resourceFwResourceUsageRead,
		DeleteContext: resourceFwResourceUsageDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwResourceUsageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating FwResourceUsage (Inside resourceFwResourceUsageCreate) ")

		data := dataToFwResourceUsage(d)
		logger.Println("[INFO] received formatted data from method data to FwResourceUsage --")
		d.SetId(strconv.Itoa('1'))
		err := go_thunder.PostFwResourceUsage(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwResourceUsageRead(ctx, d, meta)

	}
	return diags
}

func resourceFwResourceUsageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading FwResourceUsage (Inside resourceFwResourceUsageRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwResourceUsage(client.Token, client.Host)
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

func resourceFwResourceUsageUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwResourceUsageRead(ctx, d, meta)
}

func resourceFwResourceUsageDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwResourceUsageRead(ctx, d, meta)
}
func dataToFwResourceUsage(d *schema.ResourceData) go_thunder.FwResourceUsage {
	var vc go_thunder.FwResourceUsage
	var c go_thunder.FwResourceUsageInstance

	vc.UUID = c
	return vc
}
