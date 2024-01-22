package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstInterfaceIpv6Port() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_interface_ipv6_port`: Configure port for local interface IPs\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstInterfaceIpv6PortCreate,
		UpdateContext: resourceDdosDstInterfaceIpv6PortUpdate,
		ReadContext:   resourceDdosDstInterfaceIpv6PortRead,
		DeleteContext: resourceDdosDstInterfaceIpv6PortDelete,

		Schema: map[string]*schema.Schema{
			"deny": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
			},
			"glid": {
				Type: schema.TypeString, Optional: true, Description: "Global limit ID",
			},
			"port_num": {
				Type: schema.TypeInt, Required: true, Description: "Port Number",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'tcp': tcp; 'udp': udp; 'http-probe': http port for interface health check;",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"addr": {
				Type: schema.TypeString, Required: true, Description: "Addr",
			},
		},
	}
}
func resourceDdosDstInterfaceIpv6PortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstInterfaceIpv6PortCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstInterfaceIpv6Port(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstInterfaceIpv6PortRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstInterfaceIpv6PortUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstInterfaceIpv6PortUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstInterfaceIpv6Port(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstInterfaceIpv6PortRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstInterfaceIpv6PortDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstInterfaceIpv6PortDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstInterfaceIpv6Port(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstInterfaceIpv6PortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstInterfaceIpv6PortRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstInterfaceIpv6Port(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDstInterfaceIpv6Port(d *schema.ResourceData) edpt.DdosDstInterfaceIpv6Port {
	var ret edpt.DdosDstInterfaceIpv6Port
	ret.Inst.Deny = d.Get("deny").(int)
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.PortNum = d.Get("port_num").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Addr = d.Get("addr").(string)
	return ret
}
