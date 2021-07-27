package thunder

//Thunder resource SlbPassthrough

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbPassthrough() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbPassthroughCreate,
		UpdateContext: resourceSlbPassthroughUpdate,
		ReadContext:   resourceSlbPassthroughRead,
		DeleteContext: resourceSlbPassthroughDelete,
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

func resourceSlbPassthroughCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbPassthrough (Inside resourceSlbPassthroughCreate) ")

		data := dataToSlbPassthrough(d)
		logger.Println("[INFO] received formatted data from method data to SlbPassthrough --")
		d.SetId("1")
		err := go_thunder.PostSlbPassthrough(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbPassthroughRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbPassthroughRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbPassthrough (Inside resourceSlbPassthroughRead)")

	if client.Host != "" {
		name := d.Id()

		data, err := go_thunder.GetSlbPassthrough(client.Token, client.Host)
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
func resourceSlbPassthroughUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbPassthroughRead(ctx, d, meta)
}

func resourceSlbPassthroughDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbPassthroughRead(ctx, d, meta)
}

func dataToSlbPassthrough(d *schema.ResourceData) go_thunder.Passthrough {
	var vc go_thunder.Passthrough
	var c go_thunder.PassthroughInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable24, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable24
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
