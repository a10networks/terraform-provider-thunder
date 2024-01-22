package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugOspfPacket() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_ospf_packet`: OSPFv3 packets\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugOspfPacketCreate,
		UpdateContext: resourceDebugOspfPacketUpdate,
		ReadContext:   resourceDebugOspfPacketRead,
		DeleteContext: resourceDebugOspfPacketDelete,

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
func resourceDebugOspfPacketCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfPacketCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfPacket(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugOspfPacketRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugOspfPacketUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfPacketUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfPacket(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugOspfPacketRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugOspfPacketDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfPacketDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfPacket(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugOspfPacketRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfPacketRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfPacket(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugOspfPacket(d *schema.ResourceData) edpt.DebugOspfPacket {
	var ret edpt.DebugOspfPacket
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
