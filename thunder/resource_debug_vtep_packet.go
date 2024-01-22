package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugVtepPacket() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_vtep_packet`: Debug Overlay packet forwarding\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugVtepPacketCreate,
		UpdateContext: resourceDebugVtepPacketUpdate,
		ReadContext:   resourceDebugVtepPacketRead,
		DeleteContext: resourceDebugVtepPacketDelete,

		Schema: map[string]*schema.Schema{
			"dumy": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dummy",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugVtepPacketCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugVtepPacketCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugVtepPacket(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugVtepPacketRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugVtepPacketUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugVtepPacketUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugVtepPacket(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugVtepPacketRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugVtepPacketDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugVtepPacketDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugVtepPacket(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugVtepPacketRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugVtepPacketRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugVtepPacket(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugVtepPacket(d *schema.ResourceData) edpt.DebugVtepPacket {
	var ret edpt.DebugVtepPacket
	ret.Inst.Dumy = d.Get("dumy").(int)
	//omit uuid
	return ret
}
