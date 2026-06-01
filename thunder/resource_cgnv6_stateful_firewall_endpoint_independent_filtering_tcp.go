package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6StatefulFirewallEndpointIndependentFilteringTcp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_stateful_firewall_endpoint_independent_filtering_tcp`: Set endpoint-independent filtering for TCP only\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6StatefulFirewallEndpointIndependentFilteringTcpCreate,
		UpdateContext: resourceCgnv6StatefulFirewallEndpointIndependentFilteringTcpUpdate,
		ReadContext:   resourceCgnv6StatefulFirewallEndpointIndependentFilteringTcpRead,
		DeleteContext: resourceCgnv6StatefulFirewallEndpointIndependentFilteringTcpDelete,

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
func resourceCgnv6StatefulFirewallEndpointIndependentFilteringTcpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallEndpointIndependentFilteringTcpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallEndpointIndependentFilteringTcp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallEndpointIndependentFilteringTcpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6StatefulFirewallEndpointIndependentFilteringTcpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallEndpointIndependentFilteringTcpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallEndpointIndependentFilteringTcp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallEndpointIndependentFilteringTcpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6StatefulFirewallEndpointIndependentFilteringTcpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallEndpointIndependentFilteringTcpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallEndpointIndependentFilteringTcp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6StatefulFirewallEndpointIndependentFilteringTcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallEndpointIndependentFilteringTcpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallEndpointIndependentFilteringTcp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6StatefulFirewallEndpointIndependentFilteringTcpPortList(d []interface{}) []edpt.Cgnv6StatefulFirewallEndpointIndependentFilteringTcpPortList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6StatefulFirewallEndpointIndependentFilteringTcpPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6StatefulFirewallEndpointIndependentFilteringTcpPortList
		oi.Port = in["port"].(int)
		oi.PortEnd = in["port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6StatefulFirewallEndpointIndependentFilteringTcp(d *schema.ResourceData) edpt.Cgnv6StatefulFirewallEndpointIndependentFilteringTcp {
	var ret edpt.Cgnv6StatefulFirewallEndpointIndependentFilteringTcp
	ret.Inst.PortList = getSliceCgnv6StatefulFirewallEndpointIndependentFilteringTcpPortList(d.Get("port_list").([]interface{}))
	//omit uuid
	return ret
}
