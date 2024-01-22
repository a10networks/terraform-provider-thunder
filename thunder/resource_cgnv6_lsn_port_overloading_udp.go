package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnPortOverloadingUdp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lsn_port_overloading_udp`: Set behavior for UDP only\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LsnPortOverloadingUdpCreate,
		UpdateContext: resourceCgnv6LsnPortOverloadingUdpUpdate,
		ReadContext:   resourceCgnv6LsnPortOverloadingUdpRead,
		DeleteContext: resourceCgnv6LsnPortOverloadingUdpDelete,

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
func resourceCgnv6LsnPortOverloadingUdpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnPortOverloadingUdpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnPortOverloadingUdp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnPortOverloadingUdpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LsnPortOverloadingUdpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnPortOverloadingUdpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnPortOverloadingUdp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnPortOverloadingUdpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LsnPortOverloadingUdpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnPortOverloadingUdpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnPortOverloadingUdp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LsnPortOverloadingUdpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnPortOverloadingUdpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnPortOverloadingUdp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6LsnPortOverloadingUdpPortList(d []interface{}) []edpt.Cgnv6LsnPortOverloadingUdpPortList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnPortOverloadingUdpPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnPortOverloadingUdpPortList
		oi.Port = in["port"].(int)
		oi.PortEnd = in["port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6LsnPortOverloadingUdp(d *schema.ResourceData) edpt.Cgnv6LsnPortOverloadingUdp {
	var ret edpt.Cgnv6LsnPortOverloadingUdp
	ret.Inst.PortList = getSliceCgnv6LsnPortOverloadingUdpPortList(d.Get("port_list").([]interface{}))
	//omit uuid
	return ret
}
