package thunder

// Thunder resource Slb FastHttpProxy

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbFastHttpProxy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbFastHttpProxyCreate,
		UpdateContext: resourceSlbFastHttpProxyUpdate,
		ReadContext:   resourceSlbFastHttpProxyRead,
		DeleteContext: resourceSlbFastHttpProxyDelete,

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
						"counters2": {
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

func resourceSlbFastHttpProxyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	logger.Println("[INFO] Creating fast-http-proxy (Inside resourceSlbFastHttpProxyCreate)")

	if client.Host != "" {
		vc := dataToSlbFastHttpProxy(d)
		d.SetId("1")
		err := go_thunder.PostSlbFastHttpProxy(client.Token, vc, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbFastHttpProxyRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbFastHttpProxyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading fast-http-proxy (Inside resourceSlbFastHttpProxyRead)")

	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {

		name := d.Id()

		vc, err := go_thunder.GetSlbFastHttpProxy(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if vc == nil {
			logger.Println("[INFO] No fast-http-proxy found" + name)
			d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceSlbFastHttpProxyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbFastHttpProxyRead(ctx, d, meta)
}

func resourceSlbFastHttpProxyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbFastHttpProxyRead(ctx, d, meta)
}

//Utility method to instantiate FastHttpProxy Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbFastHttpProxy(d *schema.ResourceData) go_thunder.FastHttpProxy {
	var vc go_thunder.FastHttpProxy

	var c go_thunder.FastHttpProxyInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable18, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_thunder.SamplingEnable18
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		se.Counters2 = d.Get(prefix + ".counters2").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
