package thunder

//Thunder resource FwGtpInGtpFiltering

import (
	"context"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwGtpInGtpFiltering() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwGtpInGtpFilteringCreate,
		UpdateContext: resourceFwGtpInGtpFilteringUpdate,
		ReadContext:   resourceFwGtpInGtpFilteringRead,
		DeleteContext: resourceFwGtpInGtpFilteringDelete,
		Schema: map[string]*schema.Schema{
			"gtp_in_gtp_value": {
				Type:        schema.TypeString,
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

func resourceFwGtpInGtpFilteringCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating FwGtpInGtpFiltering (Inside resourceFwGtpInGtpFilteringCreate) ")

		data := dataToFwGtpInGtpFiltering(d)
		logger.Println("[INFO] received formatted data from method data to FwGtpInGtpFiltering --")
		d.SetId(strconv.Itoa('1'))
		err := go_thunder.PostFwGtpInGtpFiltering(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwGtpInGtpFilteringRead(ctx, d, meta)

	}
	return diags
}

func resourceFwGtpInGtpFilteringRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading FwGtpInGtpFiltering (Inside resourceFwGtpInGtpFilteringRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwGtpInGtpFiltering(client.Token, client.Host)
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

func resourceFwGtpInGtpFilteringUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwGtpInGtpFilteringRead(ctx, d, meta)
}

func resourceFwGtpInGtpFilteringDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwGtpInGtpFilteringRead(ctx, d, meta)
}
func dataToFwGtpInGtpFiltering(d *schema.ResourceData) go_thunder.FwGtpInGtpFiltering {
	var vc go_thunder.FwGtpInGtpFiltering
	var c go_thunder.FwGtpInGtpFilteringInstance
	c.GtpInGtpValue = d.Get("gtp_in_gtp_value").(string)

	vc.GtpInGtpValue = c
	return vc
}
