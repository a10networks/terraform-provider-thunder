package thunder

//Thunder resource SlbSip

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSip() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbSipCreate,
		UpdateContext: resourceSlbSipUpdate,
		ReadContext:   resourceSlbSipRead,
		DeleteContext: resourceSlbSipDelete,
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

func resourceSlbSipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbSip (Inside resourceSlbSipCreate) ")

		data := dataToSlbSip(d)
		logger.Println("[INFO] received formatted data from method data to SlbSip --")
		d.SetId("1")
		err := go_thunder.PostSlbSip(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbSipRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbSipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbSip (Inside resourceSlbSipRead)")

	if client.Host != "" {

		name := d.Id()

		data, err := go_thunder.GetSlbSip(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No Slb Sip found" + name)
			d.SetId("")
			return nil
		}
		return diags
	}
	return nil
}

func resourceSlbSipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbSipRead(ctx, d, meta)
}

func resourceSlbSipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbSipRead(ctx, d, meta)
}

func dataToSlbSip(d *schema.ResourceData) go_thunder.SlbSip {
	var vc go_thunder.SlbSip
	var c go_thunder.SlbSipInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable10, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable10
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
