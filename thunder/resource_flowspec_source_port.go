package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFlowspecSourcePort() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_flowspec_source_port`: Configure Source Port for a Flowspec\n\n__PLACEHOLDER__",
		CreateContext: resourceFlowspecSourcePortCreate,
		UpdateContext: resourceFlowspecSourcePortUpdate,
		ReadContext:   resourceFlowspecSourcePortRead,
		DeleteContext: resourceFlowspecSourcePortDelete,

		Schema: map[string]*schema.Schema{
			"port_attribute": {
				Type: schema.TypeString, Required: true, Description: "'eq': Match only packets on a given source port; 'gt': Match only packets with a greater port number; 'lt': Match only packets with a lower port number; 'range': match only packets in the range of port numbers;",
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
func resourceFlowspecSourcePortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecSourcePortCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecSourcePort(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFlowspecSourcePortRead(ctx, d, meta)
	}
	return diags
}

func resourceFlowspecSourcePortUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecSourcePortUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecSourcePort(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFlowspecSourcePortRead(ctx, d, meta)
	}
	return diags
}
func resourceFlowspecSourcePortDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecSourcePortDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecSourcePort(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFlowspecSourcePortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecSourcePortRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecSourcePort(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFlowspecSourcePort(d *schema.ResourceData) edpt.FlowspecSourcePort {
	var ret edpt.FlowspecSourcePort
	ret.Inst.PortAttribute = d.Get("port_attribute").(string)
	ret.Inst.PortNum = d.Get("port_num").(int)
	ret.Inst.PortNumEnd = d.Get("port_num_end").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
