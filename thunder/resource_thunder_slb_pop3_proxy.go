package thunder

//Thunder resource SlbPop3Proxy

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbPop3Proxy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbPop3ProxyCreate,
		UpdateContext: resourceSlbPop3ProxyUpdate,
		ReadContext:   resourceSlbPop3ProxyRead,
		DeleteContext: resourceSlbPop3ProxyDelete,
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

func resourceSlbPop3ProxyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbPop3Proxy (Inside resourceSlbPop3ProxyCreate) ")

		data := dataToSlbPop3Proxy(d)
		logger.Println("[INFO] received formatted data from method data to SlbPop3Proxy --")
		d.SetId("1")
		err := go_thunder.PostSlbPop3Proxy(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbPop3ProxyRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbPop3ProxyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbPop3Proxy (Inside resourceSlbPop3ProxyRead)")

	if client.Host != "" {

		data, err := go_thunder.GetSlbPop3Proxy(client.Token, client.Host)
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

func resourceSlbPop3ProxyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbPop3ProxyRead(ctx, d, meta)
}

func resourceSlbPop3ProxyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbPop3ProxyRead(ctx, d, meta)
}

func dataToSlbPop3Proxy(d *schema.ResourceData) go_thunder.Pop3Proxy {
	var vc go_thunder.Pop3Proxy
	var c go_thunder.Pop3ProxyInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable28, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable28
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
