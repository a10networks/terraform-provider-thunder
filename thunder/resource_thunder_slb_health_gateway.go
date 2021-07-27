package thunder

// Thunder resource Slb HealthGateway

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHealthGateway() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbHealthGatewayCreate,
		UpdateContext: resourceSlbHealthGatewayUpdate,
		ReadContext:   resourceSlbHealthGatewayRead,
		DeleteContext: resourceSlbHealthGatewayDelete,

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

func resourceSlbHealthGatewayCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	logger.Println("[INFO] Creating health_gateway (Inside resourceSlbHealthGatewayCreate)")

	if client.Host != "" {
		vc := dataToSlbHealthGateway(d)
		d.SetId("1")
		err := go_thunder.PostSlbHealthGateway(client.Token, vc, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbHealthGatewayRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbHealthGatewayRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading health_gateway (Inside resourceSlbHealthGatewayRead)")

	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {

		name := d.Id()

		vc, err := go_thunder.GetSlbHealthGateway(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if vc == nil {
			logger.Println("[INFO] No health_gateway found" + name)
			d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceSlbHealthGatewayUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbHealthGatewayRead(ctx, d, meta)
}

func resourceSlbHealthGatewayDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbHealthGatewayRead(ctx, d, meta)
}

//Utility method to instantiate HealthGateway Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbHealthGateway(d *schema.ResourceData) go_thunder.HealthGateway {
	var vc go_thunder.HealthGateway

	var c go_thunder.HealthGatewayInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable32, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_thunder.SamplingEnable32
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
