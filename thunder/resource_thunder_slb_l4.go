package thunder

//Thunder resource SlL4

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlL4() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbL4Create,
		UpdateContext: resourceSlbL4Update,
		ReadContext:   resourceSlbL4Read,
		DeleteContext: resourceSlbL4Delete,
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

func resourceSlbL4Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlL4 (Inside resourceSlL4Create) ")

		data := dataToSlL4(d)
		logger.Println("[INFO] received V from method data to SlL4 --")
		d.SetId("1")
		err := go_thunder.PostSlbL4(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbL4Read(ctx, d, meta)

	}
	return diags
}

func resourceSlbL4Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlL4 (Inside resourceSlL4Read)")

	if client.Host != "" {
		data, err := go_thunder.GetSlbL4(client.Token, client.Host)
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

func resourceSlbL4Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbL4Read(ctx, d, meta)
}

func resourceSlbL4Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbL4Read(ctx, d, meta)
}

func dataToSlL4(d *schema.ResourceData) go_thunder.L4 {
	var vc go_thunder.L4
	var c go_thunder.L4Instance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable42, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable42
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
