package thunder

//Thunder resource SlbSpdyProxy

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSpdyProxy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbSpdyProxyCreate,
		UpdateContext: resourceSlbSpdyProxyUpdate,
		ReadContext:   resourceSlbSpdyProxyRead,
		DeleteContext: resourceSlbSpdyProxyDelete,
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

func resourceSlbSpdyProxyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbSpdyProxy (Inside resourceSlbSpdyProxyCreate) ")

		data := dataToSlbSpdyProxy(d)
		logger.Println("[INFO] received formatted data from method data to SlbSpdyProxy --")
		d.SetId("1")
		err := go_thunder.PostSlbSpdyProxy(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbSpdyProxyRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbSpdyProxyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbSpdyProxy (Inside resourceSlbSpdyProxyRead)")

	if client.Host != "" {

		data, err := go_thunder.GetSlbSpdyProxy(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {

			return nil
		}
		return diags
	}
	return nil
}
func resourceSlbSpdyProxyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbSpdyProxyRead(ctx, d, meta)
}

func resourceSlbSpdyProxyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbSpdyProxyRead(ctx, d, meta)
}

func dataToSlbSpdyProxy(d *schema.ResourceData) go_thunder.SpdyProxy {
	var vc go_thunder.SpdyProxy
	var c go_thunder.SpdyProxyInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable14, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable14
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		obj1.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
