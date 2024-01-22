package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkVirtualWireGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_virtual_wire_global`: Virtual Wire Global Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkVirtualWireGlobalCreate,
		UpdateContext: resourceNetworkVirtualWireGlobalUpdate,
		ReadContext:   resourceNetworkVirtualWireGlobalRead,
		DeleteContext: resourceNetworkVirtualWireGlobalDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'vlan-update': VLAN update; 'mac-update': MAC update; 'vlan-pair-update': VLAN pair update;",
						},
					},
				},
			},
			"src_mac_learning": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Update source mac information for all interfaces",
			},
			"update_active_vlan": {
				Type: schema.TypeString, Optional: true, Default: "l3-packet", Description: "'all': all; 'l3-packet': l3-packet(default);",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vlan_update_period": {
				Type: schema.TypeInt, Optional: true, Default: 30, Description: "Update period in second (default is 30)",
			},
		},
	}
}
func resourceNetworkVirtualWireGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkVirtualWireGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkVirtualWireGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkVirtualWireGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkVirtualWireGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkVirtualWireGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceNetworkVirtualWireGlobalSamplingEnable(d []interface{}) []edpt.NetworkVirtualWireGlobalSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.NetworkVirtualWireGlobalSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkVirtualWireGlobalSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNetworkVirtualWireGlobal(d *schema.ResourceData) edpt.NetworkVirtualWireGlobal {
	var ret edpt.NetworkVirtualWireGlobal
	ret.Inst.SamplingEnable = getSliceNetworkVirtualWireGlobalSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SrcMacLearning = d.Get("src_mac_learning").(int)
	ret.Inst.UpdateActiveVlan = d.Get("update_active_vlan").(string)
	//omit uuid
	ret.Inst.VlanUpdatePeriod = d.Get("vlan_update_period").(int)
	return ret
}
