package thunder

//Thunder resource FwTcpMssClamp

import (
	"context"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwTcpMssClamp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwTcpMssClampCreate,
		UpdateContext: resourceFwTcpMssClampUpdate,
		ReadContext:   resourceFwTcpMssClampRead,
		DeleteContext: resourceFwTcpMssClampDelete,
		Schema: map[string]*schema.Schema{
			"mss_subtract": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"mss_value": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"mss_clamp_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"min": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwTcpMssClampCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating FwTcpMssClamp (Inside resourceFwTcpMssClampCreate) ")

		data := dataToFwTcpMssClamp(d)
		logger.Println("[INFO] received formatted data from method data to FwTcpMssClamp --")
		d.SetId(strconv.Itoa('1'))
		err := go_thunder.PostFwTcpMssClamp(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwTcpMssClampRead(ctx, d, meta)

	}
	return diags
}

func resourceFwTcpMssClampRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading FwTcpMssClamp (Inside resourceFwTcpMssClampRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwTcpMssClamp(client.Token, client.Host)
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

func resourceFwTcpMssClampUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwTcpMssClampRead(ctx, d, meta)
}

func resourceFwTcpMssClampDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceFwTcpMssClampRead(ctx, d, meta)
}
func dataToFwTcpMssClamp(d *schema.ResourceData) go_thunder.FwTcpMssClamp {
	var vc go_thunder.FwTcpMssClamp
	var c go_thunder.FwTcpMssClampInstance
	c.MssSubtract = d.Get("mss_subtract").(int)
	c.MssValue = d.Get("mss_value").(int)
	c.MssClampType = d.Get("mss_clamp_type").(string)
	c.Min = d.Get("min").(int)

	vc.MssSubtract = c
	return vc
}
