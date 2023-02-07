package thunder

//Thunder resource GlmSend

import (
	"context"

	go_thunder "github.com/go_thunder/thunder"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"util"
)

func resourceGlmSend() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGlmSendCreate,
		UpdateContext: resourceGlmSendUpdate,
		ReadContext:   resourceGlmSendRead,
		DeleteContext: resourceGlmSendDelete,
		Schema: map[string]*schema.Schema{
			"license_request": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceGlmSendCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating GlmSend (Inside resourceGlmSendCreate) ")

		data := dataToGlmSend(d)
		logger.Println("[INFO] received formatted data from method data to GlmSend --")
		d.SetId("1")
		err := go_thunder.PostGlmSend(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceGlmSendRead(ctx, d, meta)

	}
	return diags
}

func resourceGlmSendRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading GlmSend (Inside resourceGlmSendRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetGlmSend(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {

			logger.Println("[INFO] No data found ")
			return nil
		}
		return diags
	}
	return nil
}

func resourceGlmSendUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceGlmSendRead(ctx, d, meta)
}

func resourceGlmSendDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceGlmSendRead(ctx, d, meta)
}
func dataToGlmSend(d *schema.ResourceData) go_thunder.GlmSend {
	var vc go_thunder.GlmSend
	var c go_thunder.GlmSendInstance
	c.LicenseRequest = d.Get("license_request").(int)

	vc.LicenseRequest = c
	return vc
}
