package thunder

//Thunder resource FwGtpV0

import (
	"context"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwGtpV0() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwGtpV0Create,
		UpdateContext: resourceFwGtpV0Update,
		ReadContext:   resourceFwGtpV0Read,
		DeleteContext: resourceFwGtpV0Delete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"gtpv0_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwGtpV0Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating FwGtpV0 (Inside resourceFwGtpV0Create) ")

		data := dataToFwGtpV0(d)
		logger.Println("[INFO] received formatted data from method data to FwGtpV0 --")
		d.SetId(strconv.Itoa('1'))
		err := go_thunder.PostFwGtpV0(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwGtpV0Read(ctx, d, meta)

	}
	return diags
}

func resourceFwGtpV0Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading FwGtpV0 (Inside resourceFwGtpV0Read)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwGtpV0(client.Token, client.Host)
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

func resourceFwGtpV0Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwGtpV0Read(ctx, d, meta)
}

func resourceFwGtpV0Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwGtpV0Read(ctx, d, meta)
}
func dataToFwGtpV0(d *schema.ResourceData) go_thunder.FwGtpV0 {
	var vc go_thunder.FwGtpV0
	var c go_thunder.FwGtpV0Instance
	c.Gtpv0Value = d.Get("gtpv0_value").(string)

	vc.UUID = c
	return vc
}
