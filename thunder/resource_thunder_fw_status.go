package thunder

//Thunder resource FwStatus

import (
	"context"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwStatus() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwStatusCreate,
		UpdateContext: resourceFwStatusUpdate,
		ReadContext:   resourceFwStatusRead,
		DeleteContext: resourceFwStatusDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwStatusCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating FwStatus (Inside resourceFwStatusCreate) ")

		data := dataToFwStatus(d)
		logger.Println("[INFO] received formatted data from method data to FwStatus --")
		d.SetId(strconv.Itoa('1'))
		err := go_thunder.PostFwStatus(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwStatusRead(ctx, d, meta)

	}
	return diags
}

func resourceFwStatusRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading FwStatus (Inside resourceFwStatusRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwStatus(client.Token, client.Host)
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

func resourceFwStatusUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwStatusRead(ctx, d, meta)
}

func resourceFwStatusDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwStatusRead(ctx, d, meta)
}
func dataToFwStatus(d *schema.ResourceData) go_thunder.FwStatus {
	var vc go_thunder.FwStatus
	var c go_thunder.FwStatusInstance

	vc.UUID = c
	return vc
}
