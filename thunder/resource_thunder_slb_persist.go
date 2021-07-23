package thunder

//Thunder resource Slbpersist

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbpersist() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbpersistCreate,
		UpdateContext: resourceSlbpersistUpdate,
		ReadContext:   resourceSlbpersistRead,
		DeleteContext: resourceSlbpersistDelete,
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

func resourceSlbpersistCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating Slbpersist (Inside resourceSlbpersistCreate) ")

		data := dataToSlbpersist(d)
		logger.Println("[INFO] received formatted data from method data to Slbpersist --")
		d.SetId("1")
		err := go_thunder.PostSlbpersist(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbpersistRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbpersistRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading Slbpersist (Inside resourceSlbpersistRead)")

	if client.Host != "" {

		data, err := go_thunder.GetSlbpersist(client.Token, client.Host)
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

func resourceSlbpersistUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbpersistRead(ctx, d, meta)
}

func resourceSlbpersistDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbpersistRead(ctx, d, meta)
}
func dataToSlbpersist(d *schema.ResourceData) go_thunder.SlbPersist {
	var vc go_thunder.SlbPersist
	var c go_thunder.SlbPersistInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable26, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SamplingEnable26
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
