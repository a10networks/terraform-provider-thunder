package thunder

//Thunder resource SlbSmpp

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSmpp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbSmppCreate,
		UpdateContext: resourceSlbSmppUpdate,
		ReadContext:   resourceSlbSmppRead,
		DeleteContext: resourceSlbSmppDelete,
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

func resourceSlbSmppCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbSmpp (Inside resourceSlbSmppCreate) ")

		data := dataToSlbSmpp(d)
		logger.Println("[INFO] received formatted data from method data to SlbSmpp --")
		d.SetId("1")
		err := go_thunder.PostSlbSmpp(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbSmppRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbSmppRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbSmpp (Inside resourceSlbSmppRead)")

	if client.Host != "" {

		data, err := go_thunder.GetSlbSmpp(client.Token, client.Host)
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

func resourceSlbSmppUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbSmppRead(ctx, d, meta)
}

func resourceSlbSmppDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbSmppRead(ctx, d, meta)
}

func dataToSlbSmpp(d *schema.ResourceData) go_thunder.SlbSmpp {
	var vc go_thunder.SlbSmpp
	var c go_thunder.SlbSmppInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable11, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable11
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
