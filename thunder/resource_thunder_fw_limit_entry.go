package thunder

//Thunder resource FwLimitEntry

import (
	"context"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwLimitEntry() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwLimitEntryCreate,
		UpdateContext: resourceFwLimitEntryUpdate,
		ReadContext:   resourceFwLimitEntryRead,
		DeleteContext: resourceFwLimitEntryDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwLimitEntryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating FwLimitEntry (Inside resourceFwLimitEntryCreate) ")

		data := dataToFwLimitEntry(d)
		logger.Println("[INFO] received formatted data from method data to FwLimitEntry --")
		d.SetId(strconv.Itoa('1'))
		err := go_thunder.PostFwLimitEntry(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwLimitEntryRead(ctx, d, meta)

	}
	return diags
}

func resourceFwLimitEntryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading FwLimitEntry (Inside resourceFwLimitEntryRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwLimitEntry(client.Token, client.Host)
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

func resourceFwLimitEntryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwLimitEntryRead(ctx, d, meta)
}

func resourceFwLimitEntryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwLimitEntryRead(ctx, d, meta)
}
func dataToFwLimitEntry(d *schema.ResourceData) go_thunder.FwLimitEntry {
	var vc go_thunder.FwLimitEntry
	var c go_thunder.FwLimitEntryInstance

	vc.UUID = c
	return vc
}
