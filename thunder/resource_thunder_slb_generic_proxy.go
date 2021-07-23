package thunder

// Thunder resource Slb GenericProxy

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbGenericProxy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbGenericProxyCreate,
		UpdateContext: resourceSlbGenericProxyUpdate,
		ReadContext:   resourceSlbGenericProxyRead,
		DeleteContext: resourceSlbGenericProxyDelete,

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

func resourceSlbGenericProxyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	logger.Println("[INFO] Creating generic-proxy (Inside resourceSlbGenericProxyCreate)")

	if client.Host != "" {
		vc := dataToSlbGenericProxy(d)
		d.SetId("1")
		err := go_thunder.PostSlbGenericProxy(client.Token, vc, client.Host)
if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbGenericProxyRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbGenericProxyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading generic-proxy (Inside resourceSlbGenericProxyRead)")

	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {

		name := d.Id()

		vc, err := go_thunder.GetSlbGenericProxy(client.Token, client.Host)
if err != nil {
			return diag.FromErr(err)
		}
		if vc == nil {
			logger.Println("[INFO] No generic-proxy found" + name)
			d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceSlbGenericProxyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbGenericProxyRead(ctx, d, meta)
}

func resourceSlbGenericProxyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbGenericProxyRead(ctx, d, meta)
}

//Utility method to instantiate GenericProxy Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbGenericProxy(d *schema.ResourceData) go_thunder.GenericProxy {
	var vc go_thunder.GenericProxy

	var c go_thunder.GenericProxyInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable23, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_thunder.SamplingEnable23
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
