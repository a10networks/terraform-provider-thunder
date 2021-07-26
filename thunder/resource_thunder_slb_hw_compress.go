package thunder

// Thunder resource Slb HwCompress

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHwCompress() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbHwCompressCreate,
		UpdateContext: resourceSlbHwCompressUpdate,
		ReadContext:   resourceSlbHwCompressRead,
		DeleteContext: resourceSlbHwCompressDelete,

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
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}

}

func resourceSlbHwCompressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	logger.Println("[INFO] Creating hw-compress (Inside resourceSlbHwCompressCreate)")

	if client.Host != "" {
		vc := dataToSlbHwCompress(d)
		d.SetId("1")
		err := go_thunder.PostSlbHwCompress(client.Token, vc, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbHwCompressRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbHwCompressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading hw-compress (Inside resourceSlbHwCompressRead)")

	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {

		name := d.Id()

		vc, err := go_thunder.GetSlbHwCompress(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if vc == nil {
			logger.Println("[INFO] No hw-compress found" + name)
			d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceSlbHwCompressUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbHwCompressRead(ctx, d, meta)
}

func resourceSlbHwCompressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbHwCompressRead(ctx, d, meta)
}

//Utility method to instantiate HwCompress Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbHwCompress(d *schema.ResourceData) go_thunder.HwCompress {
	var vc go_thunder.HwCompress

	var c go_thunder.HwCompressInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable36, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_thunder.SamplingEnable36
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
