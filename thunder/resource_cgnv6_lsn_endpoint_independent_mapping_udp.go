package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnEndpointIndependentMappingUdp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lsn_endpoint_independent_mapping_udp`: Set behavior for UDP\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LsnEndpointIndependentMappingUdpCreate,
		UpdateContext: resourceCgnv6LsnEndpointIndependentMappingUdpUpdate,
		ReadContext:   resourceCgnv6LsnEndpointIndependentMappingUdpRead,
		DeleteContext: resourceCgnv6LsnEndpointIndependentMappingUdpDelete,

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
func resourceCgnv6LsnEndpointIndependentMappingUdpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnEndpointIndependentMappingUdpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnEndpointIndependentMappingUdp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnEndpointIndependentMappingUdpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LsnEndpointIndependentMappingUdpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnEndpointIndependentMappingUdpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnEndpointIndependentMappingUdp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnEndpointIndependentMappingUdpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LsnEndpointIndependentMappingUdpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnEndpointIndependentMappingUdpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnEndpointIndependentMappingUdp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LsnEndpointIndependentMappingUdpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnEndpointIndependentMappingUdpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnEndpointIndependentMappingUdp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6LsnEndpointIndependentMappingUdpPortList(d []interface{}) []edpt.Cgnv6LsnEndpointIndependentMappingUdpPortList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnEndpointIndependentMappingUdpPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnEndpointIndependentMappingUdpPortList
		oi.Port = in["port"].(int)
		oi.PortEnd = in["port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6LsnEndpointIndependentMappingUdp(d *schema.ResourceData) edpt.Cgnv6LsnEndpointIndependentMappingUdp {
	var ret edpt.Cgnv6LsnEndpointIndependentMappingUdp
	ret.Inst.PortList = getSliceCgnv6LsnEndpointIndependentMappingUdpPortList(d.Get("port_list").([]interface{}))
	//omit uuid
	return ret
}
