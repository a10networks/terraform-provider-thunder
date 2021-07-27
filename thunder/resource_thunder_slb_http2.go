package thunder

// Thunder resource Slb HTTP2

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHTTP2() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbHTTP2Create,
		UpdateContext: resourceSlbHTTP2Update,
		ReadContext:   resourceSlbHTTP2Read,
		DeleteContext: resourceSlbHTTP2Delete,

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

func resourceSlbHTTP2Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	logger.Println("[INFO] Creating HTTP2 (Inside resourceSlbHTTP2Create)")

	if client.Host != "" {
		vc := dataToSlbHTTP2(d)
		d.SetId("1")
		err := go_thunder.PostSlbHTTP2(client.Token, vc, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbHTTP2Read(ctx, d, meta)
	}
	return diags
}

func resourceSlbHTTP2Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading HTTP2 (Inside resourceSlbHTTP2Read)")

	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {

		name := d.Id()

		vc, err := go_thunder.GetSlbHTTP2(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if vc == nil {
			logger.Println("[INFO] No HTTP2 found" + name)
			d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceSlbHTTP2Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbHTTP2Read(ctx, d, meta)
}

func resourceSlbHTTP2Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbHTTP2Read(ctx, d, meta)
}

//Utility method to instantiate HTTP2 Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbHTTP2(d *schema.ResourceData) go_thunder.HTTP2 {
	var vc go_thunder.HTTP2

	var c go_thunder.HTTP2Instance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable34, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_thunder.SamplingEnable34
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
