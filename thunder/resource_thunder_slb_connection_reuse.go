package thunder

// Thunder resource Slb ConnectionReuse

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbConnectionReuse() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbConnectionReuseCreate,
		UpdateContext: resourceSlbConnectionReuseUpdate,
		ReadContext:   resourceSlbConnectionReuseRead,
		DeleteContext: resourceSlbConnectionReuseDelete,

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

func resourceSlbConnectionReuseCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	logger.Println("[INFO] Creating connection-reuse (Inside resourceSlbConnectionReuseCreate)")

	if client.Host != "" {
		vc := dataToSlbConnectionReuse(d)
		d.SetId("1")
		err := go_thunder.PostSlbConnectionReuse(client.Token, vc, client.Host)
if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbConnectionReuseRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbConnectionReuseRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading connection-reuse (Inside resourceSlbConnectionReuseRead)")

	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {

		name := d.Id()

		vc, err := go_thunder.GetSlbConnectionReuse(client.Token, client.Host)
if err != nil {
			return diag.FromErr(err)
		}
		if vc == nil {
			logger.Println("[INFO] No connection-reuse found" + name)
			d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceSlbConnectionReuseUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbConnectionReuseRead(ctx, d, meta)
}

func resourceSlbConnectionReuseDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbConnectionReuseRead(ctx, d, meta)
}

//Utility method to instantiate ConnectionReuse Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbConnectionReuse(d *schema.ResourceData) go_thunder.ConnectionReuse {
	var vc go_thunder.ConnectionReuse

	var c go_thunder.SlbConnectionReuseInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable7, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_thunder.SamplingEnable7
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
