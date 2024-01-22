package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFlowspecIcmpType() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_flowspec_icmp_type`: Configure ICMP Type for a Flowspec\n\n__PLACEHOLDER__",
		CreateContext: resourceFlowspecIcmpTypeCreate,
		UpdateContext: resourceFlowspecIcmpTypeUpdate,
		ReadContext:   resourceFlowspecIcmpTypeRead,
		DeleteContext: resourceFlowspecIcmpTypeDelete,

		Schema: map[string]*schema.Schema{
			"icmp_type_attribute": {
				Type: schema.TypeString, Required: true, Description: "'eq': Match only packets on a given ICMP Type; 'gt': Match only packets with a greater ICMP Type; 'lt': Match only packets with a lower ICMP Type; 'range': match only packets in the range of ICMP Types;",
			},
			"type": {
				Type: schema.TypeInt, Required: true, Description: "Specify the ICMP Type",
			},
			"type_end": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the ICMP Type",
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
func resourceFlowspecIcmpTypeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecIcmpTypeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecIcmpType(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFlowspecIcmpTypeRead(ctx, d, meta)
	}
	return diags
}

func resourceFlowspecIcmpTypeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecIcmpTypeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecIcmpType(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFlowspecIcmpTypeRead(ctx, d, meta)
	}
	return diags
}
func resourceFlowspecIcmpTypeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecIcmpTypeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecIcmpType(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFlowspecIcmpTypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecIcmpTypeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecIcmpType(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFlowspecIcmpType(d *schema.ResourceData) edpt.FlowspecIcmpType {
	var ret edpt.FlowspecIcmpType
	ret.Inst.IcmpTypeAttribute = d.Get("icmp_type_attribute").(string)
	ret.Inst.Type = d.Get("type").(int)
	ret.Inst.TypeEnd = d.Get("type_end").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
