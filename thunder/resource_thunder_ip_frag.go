package thunder

//Thunder resource IpFrag

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpFrag() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpFragCreate,
		UpdateContext: resourceIpFragUpdate,
		ReadContext:   resourceIpFragRead,
		DeleteContext: resourceIpFragDelete,
		Schema: map[string]*schema.Schema{
			"cpu_threshold": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"high": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"low": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"max_packets_per_reassembly": {
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
			"buff": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"max_reassembly_sessions": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceIpFragCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating IpFrag (Inside resourceIpFragCreate) ")

		data := dataToIpFrag(d)
		logger.Println("[INFO] received V from method data to IpFrag --")
		d.SetId("1")
		err := go_thunder.PostIpFrag(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceIpFragRead(ctx, d, meta)

	}
	return diags
}

func resourceIpFragRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading IpFrag (Inside resourceIpFragRead)")

	if client.Host != "" {
		data, err := go_thunder.GetIpFrag(client.Token, client.Host)
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

func resourceIpFragUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpFragRead(ctx, d, meta)
}

func resourceIpFragDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpFragRead(ctx, d, meta)
}
func dataToIpFrag(d *schema.ResourceData) go_thunder.Frag {
	var vc go_thunder.Frag
	var c go_thunder.FragInstance

	c.MaxReassemblySessions = d.Get("max_reassembly_sessions").(int)
	c.MaxPacketsPerReassembly = d.Get("max_packets_per_reassembly").(int)

	var obj1 go_thunder.CPUThreshold
	prefix := "cpu_threshold.0."
	obj1.Low = d.Get(prefix + "low").(int)
	obj1.High = d.Get(prefix + "high").(int)
	c.High = obj1

	c.Timeout = d.Get("timeout").(int)
	c.Buff = d.Get("buff").(int)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable47, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj2 go_thunder.SamplingEnable47
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj2.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj2)
	}
	vc.UUID = c
	return vc
}
