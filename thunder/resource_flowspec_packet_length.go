package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFlowspecPacketLength() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_flowspec_packet_length`: Configure Packet Length for a Flowspec\n\n__PLACEHOLDER__",
		CreateContext: resourceFlowspecPacketLengthCreate,
		UpdateContext: resourceFlowspecPacketLengthUpdate,
		ReadContext:   resourceFlowspecPacketLengthRead,
		DeleteContext: resourceFlowspecPacketLengthDelete,

		Schema: map[string]*schema.Schema{
			"length": {
				Type: schema.TypeInt, Required: true, Description: "Specify the Packet Length",
			},
			"length_end": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the Packet Length",
			},
			"packet_length_attribute": {
				Type: schema.TypeString, Required: true, Description: "'eq': Match only packets on a given Packet Length; 'gt': Match only packets with a greater Packet Length; 'lt': Match only packets with a lower Packet Length; 'range': match only packets in the range of Packet Lengths;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceFlowspecPacketLengthCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecPacketLengthCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecPacketLength(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFlowspecPacketLengthRead(ctx, d, meta)
	}
	return diags
}

func resourceFlowspecPacketLengthUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecPacketLengthUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecPacketLength(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFlowspecPacketLengthRead(ctx, d, meta)
	}
	return diags
}
func resourceFlowspecPacketLengthDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecPacketLengthDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecPacketLength(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFlowspecPacketLengthRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecPacketLengthRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecPacketLength(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFlowspecPacketLength(d *schema.ResourceData) edpt.FlowspecPacketLength {
	var ret edpt.FlowspecPacketLength
	ret.Inst.Length = d.Get("length").(int)
	ret.Inst.LengthEnd = d.Get("length_end").(int)
	ret.Inst.PacketLengthAttribute = d.Get("packet_length_attribute").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
