package thunder

//Thunder resource FwSystemStatus

import (
	"context"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwSystemStatus() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwSystemStatusCreate,
		UpdateContext: resourceFwSystemStatusUpdate,
		ReadContext:   resourceFwSystemStatusRead,
		DeleteContext: resourceFwSystemStatusDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwSystemStatusCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating FwSystemStatus (Inside resourceFwSystemStatusCreate) ")

		data := dataToFwSystemStatus(d)
		logger.Println("[INFO] received formatted data from method data to FwSystemStatus --")
		d.SetId(strconv.Itoa('1'))
		err := go_thunder.PostFwSystemStatus(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwSystemStatusRead(ctx, d, meta)

	}
	return diags
}

func resourceFwSystemStatusRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading FwSystemStatus (Inside resourceFwSystemStatusRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwSystemStatus(client.Token, client.Host)
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

func resourceFwSystemStatusUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwSystemStatusRead(ctx, d, meta)
}

func resourceFwSystemStatusDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwSystemStatusRead(ctx, d, meta)
}
func dataToFwSystemStatus(d *schema.ResourceData) go_thunder.FwSystemStatus {
	var vc go_thunder.FwSystemStatus
	var c go_thunder.FwSystemStatusInstance

	vc.UUID = c
	return vc
}
