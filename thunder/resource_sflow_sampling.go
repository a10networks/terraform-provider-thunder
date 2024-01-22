package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSflowSampling() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sflow_sampling`: Configure sFlow sampling on specified interfaces\n\n__PLACEHOLDER__",
		CreateContext: resourceSflowSamplingCreate,
		UpdateContext: resourceSflowSamplingUpdate,
		ReadContext:   resourceSflowSamplingRead,
		DeleteContext: resourceSflowSamplingDelete,

		Schema: map[string]*schema.Schema{
			"eth_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"eth_start": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet interface to sample",
						},
						"eth_end": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet interface to sample",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ve_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ve_start": {
							Type: schema.TypeInt, Optional: true, Description: "VE interface to sample",
						},
						"ve_end": {
							Type: schema.TypeInt, Optional: true, Description: "VE interface to sample",
						},
					},
				},
			},
		},
	}
}
func resourceSflowSamplingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowSamplingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowSampling(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowSamplingRead(ctx, d, meta)
	}
	return diags
}

func resourceSflowSamplingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowSamplingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowSampling(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowSamplingRead(ctx, d, meta)
	}
	return diags
}
func resourceSflowSamplingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowSamplingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowSampling(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSflowSamplingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowSamplingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowSampling(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSflowSamplingEthList(d []interface{}) []edpt.SflowSamplingEthList {

	count1 := len(d)
	ret := make([]edpt.SflowSamplingEthList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SflowSamplingEthList
		oi.EthStart = in["eth_start"].(int)
		oi.EthEnd = in["eth_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSflowSamplingVeList(d []interface{}) []edpt.SflowSamplingVeList {

	count1 := len(d)
	ret := make([]edpt.SflowSamplingVeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SflowSamplingVeList
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSflowSampling(d *schema.ResourceData) edpt.SflowSampling {
	var ret edpt.SflowSampling
	ret.Inst.EthList = getSliceSflowSamplingEthList(d.Get("eth_list").([]interface{}))
	//omit uuid
	ret.Inst.VeList = getSliceSflowSamplingVeList(d.Get("ve_list").([]interface{}))
	return ret
}
