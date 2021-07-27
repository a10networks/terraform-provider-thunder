package thunder

//Thunder resource SlbProxy

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbProxy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbProxyCreate,
		UpdateContext: resourceSlbProxyUpdate,
		ReadContext:   resourceSlbProxyRead,
		DeleteContext: resourceSlbProxyDelete,
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

func resourceSlbProxyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbProxy (Inside resourceSlbProxyCreate) ")

		data := dataToSlbProxy(d)
		logger.Println("[INFO] received formatted data from method data to SlbProxy --")
		d.SetId("1")
		err := go_thunder.PostSlbProxy(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbProxyRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbProxyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbProxy (Inside resourceSlbProxyRead)")

	if client.Host != "" {

		data, err := go_thunder.GetSlbProxy(client.Token, client.Host)
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

func resourceSlbProxyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbProxyRead(ctx, d, meta)
}

func resourceSlbProxyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbProxyRead(ctx, d, meta)
}

func dataToSlbProxy(d *schema.ResourceData) go_thunder.Proxy {
	var vc go_thunder.Proxy
	var c go_thunder.ProxyInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable29, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable29
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
