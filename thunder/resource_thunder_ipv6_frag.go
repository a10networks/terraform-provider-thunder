package thunder

//Thunder resource Ipv6Frag

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6Frag() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpv6FragCreate,
		UpdateContext: resourceIpv6FragUpdate,
		ReadContext:   resourceIpv6FragRead,
		DeleteContext: resourceIpv6FragDelete,
		Schema: map[string]*schema.Schema{
			"frag_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
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

func resourceIpv6FragCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating Ipv6Frag (Inside resourceIpv6FragCreate) ")

		data := dataToIpv6Frag(d)
		logger.Println("[INFO] received V from method data to Ipv6Frag --")
		d.SetId("1")
		err := go_thunder.PostIpv6Frag(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceIpv6FragRead(ctx, d, meta)

	}
	return diags
}

func resourceIpv6FragRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading Ipv6Frag (Inside resourceIpv6FragRead)")

	if client.Host != "" {
		data, err := go_thunder.GetIpv6Frag(client.Token, client.Host)
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

func resourceIpv6FragUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpv6FragRead(ctx, d, meta)
}

func resourceIpv6FragDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpv6FragRead(ctx, d, meta)
}
func dataToIpv6Frag(d *schema.ResourceData) go_thunder.IPv6Frag {
	var vc go_thunder.IPv6Frag
	var c go_thunder.IPv6FragInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable49, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable49
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	c.FragTimeout = d.Get("frag_timeout").(int)

	vc.UUID = c
	return vc
}
