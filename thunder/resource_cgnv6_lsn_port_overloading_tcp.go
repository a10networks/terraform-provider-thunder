package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnPortOverloadingTcp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lsn_port_overloading_tcp`: Set behavior for TCP only\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LsnPortOverloadingTcpCreate,
		UpdateContext: resourceCgnv6LsnPortOverloadingTcpUpdate,
		ReadContext:   resourceCgnv6LsnPortOverloadingTcpRead,
		DeleteContext: resourceCgnv6LsnPortOverloadingTcpDelete,

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
func resourceCgnv6LsnPortOverloadingTcpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnPortOverloadingTcpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnPortOverloadingTcp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnPortOverloadingTcpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LsnPortOverloadingTcpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnPortOverloadingTcpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnPortOverloadingTcp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnPortOverloadingTcpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LsnPortOverloadingTcpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnPortOverloadingTcpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnPortOverloadingTcp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LsnPortOverloadingTcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnPortOverloadingTcpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnPortOverloadingTcp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6LsnPortOverloadingTcpPortList(d []interface{}) []edpt.Cgnv6LsnPortOverloadingTcpPortList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnPortOverloadingTcpPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnPortOverloadingTcpPortList
		oi.Port = in["port"].(int)
		oi.PortEnd = in["port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6LsnPortOverloadingTcp(d *schema.ResourceData) edpt.Cgnv6LsnPortOverloadingTcp {
	var ret edpt.Cgnv6LsnPortOverloadingTcp
	ret.Inst.PortList = getSliceCgnv6LsnPortOverloadingTcpPortList(d.Get("port_list").([]interface{}))
	//omit uuid
	return ret
}
