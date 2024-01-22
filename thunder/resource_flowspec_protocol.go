package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFlowspecProtocol() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_flowspec_protocol`: Configure Protocol for a Flowspec\n\n__PLACEHOLDER__",
		CreateContext: resourceFlowspecProtocolCreate,
		UpdateContext: resourceFlowspecProtocolUpdate,
		ReadContext:   resourceFlowspecProtocolRead,
		DeleteContext: resourceFlowspecProtocolDelete,

		Schema: map[string]*schema.Schema{
			"proto_attribute": {
				Type: schema.TypeString, Required: true, Description: "'eq': Match only packets on a given protocol; 'gt': Match only packets with a greater protocol number; 'lt': Match only packets with a lower protocol number; 'range': match only packets in the range of protocol numbers;",
			},
			"proto_num": {
				Type: schema.TypeInt, Required: true, Description: "Specify the protocol number(6 for TCP and 17 for UDP)",
			},
			"proto_num_end": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the protocol number",
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
func resourceFlowspecProtocolCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecProtocolCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecProtocol(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFlowspecProtocolRead(ctx, d, meta)
	}
	return diags
}

func resourceFlowspecProtocolUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecProtocolUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecProtocol(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFlowspecProtocolRead(ctx, d, meta)
	}
	return diags
}
func resourceFlowspecProtocolDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecProtocolDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecProtocol(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFlowspecProtocolRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecProtocolRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecProtocol(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFlowspecProtocol(d *schema.ResourceData) edpt.FlowspecProtocol {
	var ret edpt.FlowspecProtocol
	ret.Inst.ProtoAttribute = d.Get("proto_attribute").(string)
	ret.Inst.ProtoNum = d.Get("proto_num").(int)
	ret.Inst.ProtoNumEnd = d.Get("proto_num_end").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
