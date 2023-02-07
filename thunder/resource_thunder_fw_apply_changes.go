package thunder

//Thunder resource FwApplyChanges

import (
	"context"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwApplyChanges() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwApplyChangesCreate,
		UpdateContext: resourceFwApplyChangesUpdate,
		ReadContext:   resourceFwApplyChangesRead,
		DeleteContext: resourceFwApplyChangesDelete,
		Schema: map[string]*schema.Schema{
			"forced": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwApplyChangesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating FwApplyChanges (Inside resourceFwApplyChangesCreate) ")

		data := dataToFwApplyChanges(d)
		logger.Println("[INFO] received formatted data from method data to FwApplyChanges --")
		d.SetId(strconv.Itoa('1'))
		err := go_thunder.PostFwApplyChanges(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwApplyChangesRead(ctx, d, meta)

	}
	return diags
}

func resourceFwApplyChangesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading FwApplyChanges (Inside resourceFwApplyChangesRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwApplyChanges(client.Token, client.Host)
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

func resourceFwApplyChangesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwApplyChangesRead(ctx, d, meta)
}

func resourceFwApplyChangesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwApplyChangesRead(ctx, d, meta)
}
func dataToFwApplyChanges(d *schema.ResourceData) go_thunder.FwApplyChanges {
	var vc go_thunder.FwApplyChanges
	var c go_thunder.FwApplyChangesInstance
	c.Forced = d.Get("forced").(int)

	vc.Forced = c
	return vc
}
