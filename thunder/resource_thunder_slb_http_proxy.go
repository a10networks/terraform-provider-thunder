package thunder

// Thunder resource Slb HttpProxy

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHttpProxy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbHttpProxyCreate,
		UpdateContext: resourceSlbHttpProxyUpdate,
		ReadContext:   resourceSlbHttpProxyRead,
		DeleteContext: resourceSlbHttpProxyDelete,

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

func resourceSlbHttpProxyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	logger.Println("[INFO] Creating http-proxy (Inside resourceSlbHttpProxyCreate)")

	if client.Host != "" {
		vc := dataToSlbHttpProxy(d)
		d.SetId("1")
		err := go_thunder.PostSlbHttpProxy(client.Token, vc, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbHttpProxyRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbHttpProxyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading http-proxy (Inside resourceSlbHttpProxyRead)")

	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {

		name := d.Id()

		vc, err := go_thunder.GetSlbHttpProxy(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if vc == nil {
			logger.Println("[INFO] No http-proxy found" + name)
			d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceSlbHttpProxyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbHttpProxyRead(ctx, d, meta)
}

func resourceSlbHttpProxyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbHttpProxyRead(ctx, d, meta)
}

//Utility method to instantiate HttpProxy Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbHttpProxy(d *schema.ResourceData) go_thunder.HttpProxy {
	var vc go_thunder.HttpProxy

	var c go_thunder.HttpProxyInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable35, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_thunder.SamplingEnable35
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		se.Counters2 = d.Get(prefix + ".counters2").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
