package thunder

// Thunder resource Slb DNS

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbDNS() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbDNSCreate,
		UpdateContext: resourceSlbDNSUpdate,
		ReadContext:   resourceSlbDNSRead,
		DeleteContext: resourceSlbDNSDelete,

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

func resourceSlbDNSCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	logger.Println("[INFO] Creating dns (Inside resourceSlbDNSCreate)")

	if client.Host != "" {
		vc := dataToSlbDNS(d)
		d.SetId("1")
		err := go_thunder.PostSlbDNS(client.Token, vc, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbDNSRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbDNSRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading dns (Inside resourceSlbDNSRead)")

	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {

		name := d.Id()

		vc, err := go_thunder.GetSlbDNS(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if vc == nil {
			logger.Println("[INFO] No dns found" + name)
			d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceSlbDNSUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbDNSRead(ctx, d, meta)
}

func resourceSlbDNSDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbDNSRead(ctx, d, meta)
}

//Utility method to instantiate DNS Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbDNS(d *schema.ResourceData) go_thunder.SlbDNS {
	var vc go_thunder.SlbDNS

	var c go_thunder.SlbDNSInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable9, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_thunder.SamplingEnable9
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
