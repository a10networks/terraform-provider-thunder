package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnEndpointIndependentFilteringUdp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lsn_endpoint_independent_filtering_udp`: Set behavior for UDP\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LsnEndpointIndependentFilteringUdpCreate,
		UpdateContext: resourceCgnv6LsnEndpointIndependentFilteringUdpUpdate,
		ReadContext:   resourceCgnv6LsnEndpointIndependentFilteringUdpRead,
		DeleteContext: resourceCgnv6LsnEndpointIndependentFilteringUdpDelete,

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
			"session_limit": {
				Type: schema.TypeInt, Optional: true, Default: 65535, Description: "Limit number of EIF sessions that can be created per port",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6LsnEndpointIndependentFilteringUdpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnEndpointIndependentFilteringUdpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnEndpointIndependentFilteringUdp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnEndpointIndependentFilteringUdpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LsnEndpointIndependentFilteringUdpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnEndpointIndependentFilteringUdpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnEndpointIndependentFilteringUdp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnEndpointIndependentFilteringUdpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LsnEndpointIndependentFilteringUdpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnEndpointIndependentFilteringUdpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnEndpointIndependentFilteringUdp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LsnEndpointIndependentFilteringUdpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnEndpointIndependentFilteringUdpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnEndpointIndependentFilteringUdp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6LsnEndpointIndependentFilteringUdpPortList(d []interface{}) []edpt.Cgnv6LsnEndpointIndependentFilteringUdpPortList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnEndpointIndependentFilteringUdpPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnEndpointIndependentFilteringUdpPortList
		oi.Port = in["port"].(int)
		oi.PortEnd = in["port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6LsnEndpointIndependentFilteringUdp(d *schema.ResourceData) edpt.Cgnv6LsnEndpointIndependentFilteringUdp {
	var ret edpt.Cgnv6LsnEndpointIndependentFilteringUdp
	ret.Inst.PortList = getSliceCgnv6LsnEndpointIndependentFilteringUdpPortList(d.Get("port_list").([]interface{}))
	ret.Inst.SessionLimit = d.Get("session_limit").(int)
	//omit uuid
	return ret
}
