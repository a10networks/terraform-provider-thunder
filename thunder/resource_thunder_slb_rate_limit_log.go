package thunder

//Thunder resource SlbRateLimitLog

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbRateLimitLog() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbRateLimitLogCreate,
		UpdateContext: resourceSlbRateLimitLogUpdate,
		ReadContext:   resourceSlbRateLimitLogRead,
		DeleteContext: resourceSlbRateLimitLogDelete,
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

func resourceSlbRateLimitLogCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbRateLimitLog (Inside resourceSlbRateLimitLogCreate) ")

		data := dataToSlbRateLimitLog(d)
		logger.Println("[INFO] received formatted data from method data to SlbRateLimitLog --")
		d.SetId("1")
		err := go_thunder.PostSlbRateLimitLog(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbRateLimitLogRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbRateLimitLogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbRateLimitLog (Inside resourceSlbRateLimitLogRead)")

	if client.Host != "" {

		data, err := go_thunder.GetSlbRateLimitLog(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {

			d.SetId("")
			return nil
		}
		return diags
	}
	return nil
}

func resourceSlbRateLimitLogUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbRateLimitLogRead(ctx, d, meta)
}

func resourceSlbRateLimitLogDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbRateLimitLogRead(ctx, d, meta)
}

func dataToSlbRateLimitLog(d *schema.ResourceData) go_thunder.RateLimitLog {
	var vc go_thunder.RateLimitLog
	var c go_thunder.RateLimitLogInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable30, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable30
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
