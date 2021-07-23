package thunder

// Thunder resource Slb Fix

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbFix() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbFixCreate,
		UpdateContext: resourceSlbFixUpdate,
		ReadContext:   resourceSlbFixRead,
		DeleteContext: resourceSlbFixDelete,

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

func resourceSlbFixCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	logger.Println("[INFO] Creating fix (Inside resourceSlbFixCreate)")

	if client.Host != "" {
		vc := dataToSlbFix(d)
		d.SetId("1")
		err := go_thunder.PostSlbFix(client.Token, vc, client.Host)
if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbFixRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbFixRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading fix (Inside resourceSlbFixRead)")

	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {

		name := d.Id()

		vc, err := go_thunder.GetSlbFix(client.Token, client.Host)
if err != nil {
			return diag.FromErr(err)
		}
		if vc == nil {
			logger.Println("[INFO] No fix found" + name)
			d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceSlbFixUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbFixRead(ctx, d, meta)
}

func resourceSlbFixDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbFixRead(ctx, d, meta)
}

//Utility method to instantiate Fix Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbFix(d *schema.ResourceData) go_thunder.SlbFix {
	var vc go_thunder.SlbFix

	var c go_thunder.SlbFixInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable19, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_thunder.SamplingEnable19
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
