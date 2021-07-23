package thunder

//Thunder resource SlbRcCacheGlobal

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbRcCacheGlobal() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbRcCacheGlobalCreate,
		UpdateContext: resourceSlbRcCacheGlobalUpdate,
		ReadContext:   resourceSlbRcCacheGlobalRead,
		DeleteContext: resourceSlbRcCacheGlobalDelete,
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

func resourceSlbRcCacheGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbRcCacheGlobal (Inside resourceSlbRcCacheGlobalCreate) ")

		data := dataToSlbRcCacheGlobal(d)
		logger.Println("[INFO] received formatted data from method data to SlbRcCacheGlobal --")
		d.SetId("1")
		err := go_thunder.PostSlbRcCacheGlobal(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbRcCacheGlobalRead(ctx, d, meta)

	}
	return diags
}
func resourceSlbRcCacheGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbRcCacheGlobalRead(ctx, d, meta)
}

func resourceSlbRcCacheGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbRcCacheGlobalRead(ctx, d, meta)
}

func resourceSlbRcCacheGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbRcCacheGlobal (Inside resourceSlbRcCacheGlobalRead)")

	if client.Host != "" {

		data, err := go_thunder.GetSlbRcCacheGlobal(client.Token, client.Host)
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

func dataToSlbRcCacheGlobal(d *schema.ResourceData) go_thunder.RcCacheGlobal {
	var vc go_thunder.RcCacheGlobal
	var c go_thunder.RcCacheGlobalInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable31, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable31
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
