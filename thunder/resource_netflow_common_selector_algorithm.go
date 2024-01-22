package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetflowCommonSelectorAlgorithm() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_netflow_common_selector_algorithm`: NetFlow/IPFIX Selector Algorithm Setting\n\n__PLACEHOLDER__",
		CreateContext: resourceNetflowCommonSelectorAlgorithmCreate,
		UpdateContext: resourceNetflowCommonSelectorAlgorithmUpdate,
		ReadContext:   resourceNetflowCommonSelectorAlgorithmRead,
		DeleteContext: resourceNetflowCommonSelectorAlgorithmDelete,

		Schema: map[string]*schema.Schema{
			"algorithm_name": {
				Type: schema.TypeString, Required: true, Description: "'random': random;",
			},
			"sampling_population": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Configure sampling population for random algorithm",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceNetflowCommonSelectorAlgorithmCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowCommonSelectorAlgorithmCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowCommonSelectorAlgorithm(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowCommonSelectorAlgorithmRead(ctx, d, meta)
	}
	return diags
}

func resourceNetflowCommonSelectorAlgorithmUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowCommonSelectorAlgorithmUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowCommonSelectorAlgorithm(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowCommonSelectorAlgorithmRead(ctx, d, meta)
	}
	return diags
}
func resourceNetflowCommonSelectorAlgorithmDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowCommonSelectorAlgorithmDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowCommonSelectorAlgorithm(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetflowCommonSelectorAlgorithmRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowCommonSelectorAlgorithmRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowCommonSelectorAlgorithm(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNetflowCommonSelectorAlgorithm(d *schema.ResourceData) edpt.NetflowCommonSelectorAlgorithm {
	var ret edpt.NetflowCommonSelectorAlgorithm
	ret.Inst.AlgorithmName = d.Get("algorithm_name").(string)
	ret.Inst.SamplingPopulation = d.Get("sampling_population").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
