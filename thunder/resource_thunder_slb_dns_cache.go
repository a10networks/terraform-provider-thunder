package thunder

// Thunder resource Slb DNSCache

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbDNSCache() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbDNSCacheCreate,
		UpdateContext: resourceSlbDNSCacheUpdate,
		ReadContext:   resourceSlbDNSCacheRead,
		DeleteContext: resourceSlbDNSCacheDelete,

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

func resourceSlbDNSCacheCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	logger.Println("[INFO] Creating DNSCache (Inside resourceSlbDNSCacheCreate)")

	if client.Host != "" {
		vc := dataToSlbDNSCache(d)
		d.SetId("1")
		err := go_thunder.PostSlbDNSCache(client.Token, vc, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbDNSCacheRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbDNSCacheRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading DNSCache (Inside resourceSlbDNSCacheRead)")

	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {

		name := d.Id()

		vc, err := go_thunder.GetSlbDNSCache(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if vc == nil {
			logger.Println("[INFO] No DNSCache found" + name)
			d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceSlbDNSCacheUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbDNSCacheRead(ctx, d, meta)
}

func resourceSlbDNSCacheDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbDNSCacheRead(ctx, d, meta)
}

//Utility method to instantiate DNSCache Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbDNSCache(d *schema.ResourceData) go_thunder.DNSCache {
	var vc go_thunder.DNSCache

	var c go_thunder.DNSCacheInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable16, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_thunder.SamplingEnable16
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
