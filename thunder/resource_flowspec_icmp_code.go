package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFlowspecIcmpCode() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_flowspec_icmp_code`: Configure ICMP Code for a Flowspec\n\n__PLACEHOLDER__",
		CreateContext: resourceFlowspecIcmpCodeCreate,
		UpdateContext: resourceFlowspecIcmpCodeUpdate,
		ReadContext:   resourceFlowspecIcmpCodeRead,
		DeleteContext: resourceFlowspecIcmpCodeDelete,

		Schema: map[string]*schema.Schema{
			"code": {
				Type: schema.TypeInt, Required: true, Description: "Specify the ICMP Code",
			},
			"code_end": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the ICMP Code",
			},
			"icmp_code_attribute": {
				Type: schema.TypeString, Required: true, Description: "'eq': Match only packets on a given ICMP Code; 'gt': Match only packets with a greater ICMP Code; 'lt': Match only packets with a lower ICMP Code; 'range': match only packets in the range of ICMP Codes;",
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
func resourceFlowspecIcmpCodeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecIcmpCodeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecIcmpCode(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFlowspecIcmpCodeRead(ctx, d, meta)
	}
	return diags
}

func resourceFlowspecIcmpCodeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecIcmpCodeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecIcmpCode(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFlowspecIcmpCodeRead(ctx, d, meta)
	}
	return diags
}
func resourceFlowspecIcmpCodeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecIcmpCodeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecIcmpCode(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFlowspecIcmpCodeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecIcmpCodeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecIcmpCode(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFlowspecIcmpCode(d *schema.ResourceData) edpt.FlowspecIcmpCode {
	var ret edpt.FlowspecIcmpCode
	ret.Inst.Code = d.Get("code").(int)
	ret.Inst.CodeEnd = d.Get("code_end").(int)
	ret.Inst.IcmpCodeAttribute = d.Get("icmp_code_attribute").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
