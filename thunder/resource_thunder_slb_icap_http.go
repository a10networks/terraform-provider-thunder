package thunder

//Thunder resource SlbIcapHTTP

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbIcapHTTP() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbIcapHTTPCreate,
		UpdateContext: resourceSlbIcapHTTPUpdate,
		ReadContext:   resourceSlbIcapHTTPRead,
		DeleteContext: resourceSlbIcapHTTPDelete,
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

func resourceSlbIcapHTTPCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbIcapHTTP (Inside resourceSlbIcapHTTPCreate) ")

		data := dataToSlbIcapHTTP(d)
		logger.Println("[INFO] received V from method data to SlbIcapHTTP --")
		d.SetId("1")
		err := go_thunder.PostSlbIcapHTTP(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbIcapHTTPRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbIcapHTTPRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbIcapHTTP (Inside resourceSlbIcapHTTPRead)")

	if client.Host != "" {
		data, err := go_thunder.GetSlbIcapHTTP(client.Token, client.Host)
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
func resourceSlbIcapHTTPUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbIcapHTTPRead(ctx, d, meta)
}

func resourceSlbIcapHTTPDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbIcapHTTPRead(ctx, d, meta)
}
func dataToSlbIcapHTTP(d *schema.ResourceData) go_thunder.IcapHTTP {
	var vc go_thunder.IcapHTTP
	var c go_thunder.IcapHTTPInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable40, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable40
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
