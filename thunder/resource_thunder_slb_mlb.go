package thunder

//Thunder resource SlbMlb

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbMlb() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbMlbCreate,
		UpdateContext: resourceSlbMlbUpdate,
		ReadContext:   resourceSlbMlbRead,
		DeleteContext: resourceSlbMlbDelete,
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

func resourceSlbMlbCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbMlb (Inside resourceSlbMlbCreate) ")

		data := dataToSlbMlb(d)
		logger.Println("[INFO] received V from method data to SlbMlb --")
		d.SetId("1")
		err := go_thunder.PostSlbMlb(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbMlbRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbMlbRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbMlb (Inside resourceSlbMlbRead)")

	if client.Host != "" {

		data, err := go_thunder.GetSlbMlb(client.Token, client.Host)
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

func resourceSlbMlbUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbMlbRead(ctx, d, meta)
}

func resourceSlbMlbDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbMlbRead(ctx, d, meta)
}

func dataToSlbMlb(d *schema.ResourceData) go_thunder.Mlb {
	var vc go_thunder.Mlb
	var c go_thunder.MlbInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable44, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable44
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
