package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6StatefulFirewallEndpointIndependentFilteringUdp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_stateful_firewall_endpoint_independent_filtering_udp`: Set endpoint-independent filtering for UDP only\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6StatefulFirewallEndpointIndependentFilteringUdpCreate,
		UpdateContext: resourceCgnv6StatefulFirewallEndpointIndependentFilteringUdpUpdate,
		ReadContext:   resourceCgnv6StatefulFirewallEndpointIndependentFilteringUdpRead,
		DeleteContext: resourceCgnv6StatefulFirewallEndpointIndependentFilteringUdpDelete,

		Schema: map[string]*schema.Schema{
			"port_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": {
							Type: schema.TypeInt, Optional: true, Description: "Single Destination Port or Port Range Start",
						},
						"port_end": {
							Type: schema.TypeInt, Optional: true, Description: "Port Range End",
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
func resourceCgnv6StatefulFirewallEndpointIndependentFilteringUdpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallEndpointIndependentFilteringUdpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallEndpointIndependentFilteringUdp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallEndpointIndependentFilteringUdpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6StatefulFirewallEndpointIndependentFilteringUdpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallEndpointIndependentFilteringUdpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallEndpointIndependentFilteringUdp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallEndpointIndependentFilteringUdpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6StatefulFirewallEndpointIndependentFilteringUdpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallEndpointIndependentFilteringUdpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallEndpointIndependentFilteringUdp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6StatefulFirewallEndpointIndependentFilteringUdpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallEndpointIndependentFilteringUdpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallEndpointIndependentFilteringUdp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6StatefulFirewallEndpointIndependentFilteringUdpPortList(d []interface{}) []edpt.Cgnv6StatefulFirewallEndpointIndependentFilteringUdpPortList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6StatefulFirewallEndpointIndependentFilteringUdpPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6StatefulFirewallEndpointIndependentFilteringUdpPortList
		oi.Port = in["port"].(int)
		oi.PortEnd = in["port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6StatefulFirewallEndpointIndependentFilteringUdp(d *schema.ResourceData) edpt.Cgnv6StatefulFirewallEndpointIndependentFilteringUdp {
	var ret edpt.Cgnv6StatefulFirewallEndpointIndependentFilteringUdp
	ret.Inst.PortList = getSliceCgnv6StatefulFirewallEndpointIndependentFilteringUdpPortList(d.Get("port_list").([]interface{}))
	//omit uuid
	return ret
}
