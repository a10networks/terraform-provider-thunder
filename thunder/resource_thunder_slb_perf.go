package thunder

//Thunder resource SlbPerf

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbPerf() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbPerfCreate,
		UpdateContext: resourceSlbPerfUpdate,
		ReadContext:   resourceSlbPerfRead,
		DeleteContext: resourceSlbPerfDelete,
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

func resourceSlbPerfCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbPerf (Inside resourceSlbPerfCreate) ")

		data := dataToSlbPerf(d)
		logger.Println("[INFO] received formatted data from method data to SlbPerf --")
		d.SetId("1")
		err := go_thunder.PostSlbPerf(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbPerfRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbPerfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbPerf (Inside resourceSlbPerfRead)")

	if client.Host != "" {
		name := d.Id()

		data, err := go_thunder.GetSlbPerf(client.Token, client.Host)
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

func resourceSlbPerfUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbPerfRead(ctx, d, meta)
}

func resourceSlbPerfDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbPerfRead(ctx, d, meta)
}

func dataToSlbPerf(d *schema.ResourceData) go_thunder.Perf {
	var vc go_thunder.Perf
	var c go_thunder.PerfInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable25, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable25
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
