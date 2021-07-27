package thunder

//Thunder resource SlbSwitch

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSwitch() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbSwitchCreate,
		UpdateContext: resourceSlbSwitchUpdate,
		ReadContext:   resourceSlbSwitchRead,
		DeleteContext: resourceSlbSwitchDelete,
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

func resourceSlbSwitchCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbSwitch (Inside resourceSlbSwitchCreate) ")
		d.SetId("1")
		data := dataToSlbSwitch(d)
		logger.Println("[INFO] received formatted data from method data to SlbSwitch --")
		err := go_thunder.PostSlbSwitch(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbSwitchRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbSwitchRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbSwitch (Inside resourceSlbSwitchRead)")

	if client.Host != "" {
		data, err := go_thunder.GetSlbSwitch(client.Token, client.Host)
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

func resourceSlbSwitchUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbSwitchRead(ctx, d, meta)
}

func resourceSlbSwitchDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbSwitchRead(ctx, d, meta)
}

func dataToSlbSwitch(d *schema.ResourceData) go_thunder.Switch {
	var vc go_thunder.Switch
	var c go_thunder.SwitchInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable13, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable13
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		obj1.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
