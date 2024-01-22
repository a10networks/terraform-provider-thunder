package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6OneToOneGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_one_to_one_global`: Set one-to-one NAT config parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6OneToOneGlobalCreate,
		UpdateContext: resourceCgnv6OneToOneGlobalUpdate,
		ReadContext:   resourceCgnv6OneToOneGlobalRead,
		DeleteContext: resourceCgnv6OneToOneGlobalDelete,

		Schema: map[string]*schema.Schema{
			"mapping_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Configure timeout for the one-to-one NAT mapping (Timeout in minutes (default: 10 minutes))",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'total-map-allocated': Total One-to-One Address Mapping Allocated; 'total-map-freed': Total One-to-One Address Mapping Freed; 'map-alloc-failure': One-to-One Address Mapping Allocation Failure; 'map-dbl-free': Address Mapping Double Free; 'alloc-map-race': Mapping Exists When Allocating Address Mapping; 'map-not-found': Mapping to be Released Not Found; 'ha-map-mismatch': HA Standby Mapping Mismatch; 'ha-select-addr-failure': HA Standby Allocate Address Failure; 'ha-alloc-map-conflicts': HA Standby Allocate Mapping Conflicts;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6OneToOneGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOneGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOneGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6OneToOneGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6OneToOneGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOneGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOneGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6OneToOneGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6OneToOneGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOneGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOneGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6OneToOneGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOneGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOneGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6OneToOneGlobalSamplingEnable(d []interface{}) []edpt.Cgnv6OneToOneGlobalSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6OneToOneGlobalSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6OneToOneGlobalSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6OneToOneGlobal(d *schema.ResourceData) edpt.Cgnv6OneToOneGlobal {
	var ret edpt.Cgnv6OneToOneGlobal
	ret.Inst.MappingTimeout = d.Get("mapping_timeout").(int)
	ret.Inst.SamplingEnable = getSliceCgnv6OneToOneGlobalSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
