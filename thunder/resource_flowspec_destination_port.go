package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFlowspecDestinationPort() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_flowspec_destination_port`: Configure Destination Port for a Flowspec\n\n__PLACEHOLDER__",
		CreateContext: resourceFlowspecDestinationPortCreate,
		UpdateContext: resourceFlowspecDestinationPortUpdate,
		ReadContext:   resourceFlowspecDestinationPortRead,
		DeleteContext: resourceFlowspecDestinationPortDelete,

		Schema: map[string]*schema.Schema{
			"port_attribute": {
				Type: schema.TypeString, Required: true, Description: "'eq': Match only packets on a given destination port; 'gt': Match only packets with a greater port number; 'lt': Match only packets with a lower port number; 'range': match only packets in the range of port numbers;",
			},
			"port_num": {
				Type: schema.TypeInt, Required: true, Description: "Specify the port number",
			},
			"port_num_end": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the port number",
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
func resourceFlowspecDestinationPortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecDestinationPortCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecDestinationPort(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFlowspecDestinationPortRead(ctx, d, meta)
	}
	return diags
}

func resourceFlowspecDestinationPortUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecDestinationPortUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecDestinationPort(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFlowspecDestinationPortRead(ctx, d, meta)
	}
	return diags
}
func resourceFlowspecDestinationPortDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecDestinationPortDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecDestinationPort(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFlowspecDestinationPortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecDestinationPortRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecDestinationPort(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFlowspecDestinationPort(d *schema.ResourceData) edpt.FlowspecDestinationPort {
	var ret edpt.FlowspecDestinationPort
	ret.Inst.PortAttribute = d.Get("port_attribute").(string)
	ret.Inst.PortNum = d.Get("port_num").(int)
	ret.Inst.PortNumEnd = d.Get("port_num_end").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
