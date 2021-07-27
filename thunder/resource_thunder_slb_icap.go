package thunder

//Thunder resource SlbIcapHTTP

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbIcap() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbIcapCreate,
		UpdateContext: resourceSlbIcapUpdate,
		ReadContext:   resourceSlbIcapRead,
		DeleteContext: resourceSlbIcapDelete,
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

func resourceSlbIcapCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbIcap (Inside resourceSlbIcapCreate) ")

		data := dataToSlbIcap(d)
		logger.Println("[INFO] received V from method data to SlbIcap --")
		d.SetId("1")
		err := go_thunder.PostSlbIcap(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbIcapRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbIcapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbIcap (Inside resourceSlbIcapRead)")

	if client.Host != "" {
		data, err := go_thunder.GetSlbIcap(client.Token, client.Host)
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
func resourceSlbIcapUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbIcapRead(ctx, d, meta)
}

func resourceSlbIcapDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbIcapRead(ctx, d, meta)
}
func dataToSlbIcap(d *schema.ResourceData) go_thunder.Icap {
	var vc go_thunder.Icap
	var c go_thunder.IcapInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable39, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable39
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
