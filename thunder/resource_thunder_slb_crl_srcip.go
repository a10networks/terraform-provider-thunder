package thunder

// Thunder resource Slb CrlSrcip

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbCrlSrcip() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbCrlSrcipCreate,
		UpdateContext: resourceSlbCrlSrcipUpdate,
		ReadContext:   resourceSlbCrlSrcipRead,
		DeleteContext: resourceSlbCrlSrcipDelete,

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

func resourceSlbCrlSrcipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	logger.Println("[INFO] Creating crl-srcip (Inside resourceSlbCrlSrcipCreate)")

	if client.Host != "" {
		vc := dataToSlbCrlSrcip(d)
		d.SetId("1")
		err := go_thunder.PostSlbCrlSrcip(client.Token, vc, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbCrlSrcipRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbCrlSrcipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading crl-srcip (Inside resourceSlbCrlSrcipRead)")

	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {

		name := d.Id()

		vc, err := go_thunder.GetSlbCrlSrcip(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if vc == nil {
			logger.Println("[INFO] No crl-srcip found" + name)
			d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceSlbCrlSrcipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbCrlSrcipRead(ctx, d, meta)
}

func resourceSlbCrlSrcipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbCrlSrcipRead(ctx, d, meta)
}

//Utility method to instantiate CrlSrcip Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbCrlSrcip(d *schema.ResourceData) go_thunder.CrlSrcip {
	var vc go_thunder.CrlSrcip

	var c go_thunder.CrlSrcipInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable8, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_thunder.SamplingEnable8
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
