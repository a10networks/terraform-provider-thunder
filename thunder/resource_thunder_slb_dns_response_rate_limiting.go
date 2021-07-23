package thunder

// Thunder resource Slb DNSResponseRateLimiting

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbDNSResponseRateLimiting() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbDNSResponseRateLimitingCreate,
		UpdateContext: resourceSlbDNSResponseRateLimitingUpdate,
		ReadContext:   resourceSlbDNSResponseRateLimitingRead,
		DeleteContext: resourceSlbDNSResponseRateLimitingDelete,

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

func resourceSlbDNSResponseRateLimitingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	logger.Println("[INFO] Creating DNS Response Rate Limiting (Inside resourceSlbDNSResponseRateLimitingCreate)")

	if client.Host != "" {
		vc := dataToSlbDNSResponseRateLimiting(d)
		d.SetId("1")
		err := go_thunder.PostSlbDNSResponseRateLimiting(client.Token, vc, client.Host)
if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbDNSResponseRateLimitingRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbDNSResponseRateLimitingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading DNS Response Rate Limiting (Inside resourceSlbDNSResponseRateLimitingRead)")

	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {

		name := d.Id()

		vc, err := go_thunder.GetSlbDNSResponseRateLimiting(client.Token, client.Host)
if err != nil {
			return diag.FromErr(err)
		}
		if vc == nil {
			logger.Println("[INFO] No DNS Response Rate Limiting found" + name)
			d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceSlbDNSResponseRateLimitingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbDNSResponseRateLimitingRead(ctx, d, meta)
}

func resourceSlbDNSResponseRateLimitingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbDNSResponseRateLimitingRead(ctx, d, meta)
}

//Utility method to instantiate DNSResponseRateLimiting Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbDNSResponseRateLimiting(d *schema.ResourceData) go_thunder.SlbDNSResponseRateLimiting {
	var vc go_thunder.SlbDNSResponseRateLimiting

	var c go_thunder.SlbDNSResponseRateLimitingInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable17, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_thunder.SamplingEnable17
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
