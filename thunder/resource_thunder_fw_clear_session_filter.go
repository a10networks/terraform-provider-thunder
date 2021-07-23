package thunder

//Thunder resource FwClearSessionFilter

import (
	"context"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwClearSessionFilter() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwClearSessionFilterCreate,
		UpdateContext: resourceFwClearSessionFilterUpdate,
		ReadContext:   resourceFwClearSessionFilterRead,
		DeleteContext: resourceFwClearSessionFilterDelete,
		Schema: map[string]*schema.Schema{
			"status": {
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

func resourceFwClearSessionFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating FwClearSessionFilter (Inside resourceFwClearSessionFilterCreate) ")

		data := dataToFwClearSessionFilter(d)
		logger.Println("[INFO] received formatted data from method data to FwClearSessionFilter --")
		d.SetId(strconv.Itoa('1'))
		err := go_thunder.PostFwClearSessionFilter(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwClearSessionFilterRead(ctx, d, meta)

	}
	return diags
}

func resourceFwClearSessionFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading FwClearSessionFilter (Inside resourceFwClearSessionFilterRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwClearSessionFilter(client.Token, client.Host)
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

func resourceFwClearSessionFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwClearSessionFilterRead(ctx, d, meta)
}

func resourceFwClearSessionFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwClearSessionFilterRead(ctx, d, meta)
}
func dataToFwClearSessionFilter(d *schema.ResourceData) go_thunder.FwClearSessionFilter {
	var vc go_thunder.FwClearSessionFilter
	var c go_thunder.FwClearSessionFilterInstance
	c.Status = d.Get("status").(string)

	vc.Status = c
	return vc
}
