package thunder

//Thunder resource FwFullConeSession

import (
	"context"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwFullConeSession() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwFullConeSessionCreate,
		UpdateContext: resourceFwFullConeSessionUpdate,
		ReadContext:   resourceFwFullConeSessionRead,
		DeleteContext: resourceFwFullConeSessionDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwFullConeSessionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating FwFullConeSession (Inside resourceFwFullConeSessionCreate) ")

		data := dataToFwFullConeSession(d)
		logger.Println("[INFO] received formatted data from method data to FwFullConeSession --")
		d.SetId(strconv.Itoa('1'))
		err := go_thunder.PostFwFullConeSession(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwFullConeSessionRead(ctx, d, meta)

	}
	return diags
}

func resourceFwFullConeSessionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading FwFullConeSession (Inside resourceFwFullConeSessionRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwFullConeSession(client.Token, client.Host)
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

func resourceFwFullConeSessionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwFullConeSessionRead(ctx, d, meta)
}

func resourceFwFullConeSessionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwFullConeSessionRead(ctx, d, meta)
}
func dataToFwFullConeSession(d *schema.ResourceData) go_thunder.FwFullConeSession {
	var vc go_thunder.FwFullConeSession
	var c go_thunder.FwFullConeSessionInstance

	vc.UUID = c
	return vc
}
