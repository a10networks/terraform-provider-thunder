package thunder

//Thunder resource SlbSportRateLimit

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSportRateLimit() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbSportRateLimitCreate,
		UpdateContext: resourceSlbSportRateLimitUpdate,
		ReadContext:   resourceSlbSportRateLimitRead,
		DeleteContext: resourceSlbSportRateLimitDelete,
		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbSportRateLimitCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbSportRateLimit (Inside resourceSlbSportRateLimitCreate) ")

		data := dataToSlbSportRateLimit(d)
		logger.Println("[INFO] received formatted data from method data to SlbSportRateLimit --")
		d.SetId("1")
		err := go_thunder.PostSlbSportRateLimit(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbSportRateLimitRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbSportRateLimitRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbSportRateLimit (Inside resourceSlbSportRateLimitRead)")

	if client.Host != "" {

		data, err := go_thunder.GetSlbSportRateLimit(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {

			return nil
		}
		return diags
	}
	return nil
}

func resourceSlbSportRateLimitUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbSportRateLimitRead(ctx, d, meta)
}

func resourceSlbSportRateLimitDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbSportRateLimitRead(ctx, d, meta)
}

func dataToSlbSportRateLimit(d *schema.ResourceData) go_thunder.SportRateLimit {
	var vc go_thunder.SportRateLimit
	var c go_thunder.SportRateLimitInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable15, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable15
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		obj1.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
