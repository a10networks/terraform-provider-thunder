package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugIpv6OspfPacket() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_ipv6_ospf_packet`: OSPFv3 packets\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugIpv6OspfPacketCreate,
		UpdateContext: resourceDebugIpv6OspfPacketUpdate,
		ReadContext:   resourceDebugIpv6OspfPacketRead,
		DeleteContext: resourceDebugIpv6OspfPacketDelete,

		Schema: map[string]*schema.Schema{
			"dd": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "OSPFv3 Database Description",
			},
			"detail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Detail information",
			},
			"hello": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "OSPFv3 Hello",
			},
			"ls_ack": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "OSPFv3 Link State Acknowledgment",
			},
			"ls_request": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "OSPFv3 Link State Request",
			},
			"ls_update": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "OSPFv3 Link State Update",
			},
			"recv": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Packet received",
			},
			"send": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Packet sent",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugIpv6OspfPacketCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfPacketCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfPacket(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugIpv6OspfPacketRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugIpv6OspfPacketUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfPacketUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfPacket(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugIpv6OspfPacketRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugIpv6OspfPacketDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfPacketDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfPacket(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugIpv6OspfPacketRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfPacketRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfPacket(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugIpv6OspfPacket(d *schema.ResourceData) edpt.DebugIpv6OspfPacket {
	var ret edpt.DebugIpv6OspfPacket
	ret.Inst.Dd = d.Get("dd").(int)
	ret.Inst.Detail = d.Get("detail").(int)
	ret.Inst.Hello = d.Get("hello").(int)
	ret.Inst.LsAck = d.Get("ls_ack").(int)
	ret.Inst.LsRequest = d.Get("ls_request").(int)
	ret.Inst.LsUpdate = d.Get("ls_update").(int)
	ret.Inst.Recv = d.Get("recv").(int)
	ret.Inst.Send = d.Get("send").(int)
	//omit uuid
	return ret
}
