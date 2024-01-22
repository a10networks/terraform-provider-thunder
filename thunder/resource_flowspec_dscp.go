package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFlowspecDscp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_flowspec_dscp`: Configure DSCP for a Flowspec\n\n__PLACEHOLDER__",
		CreateContext: resourceFlowspecDscpCreate,
		UpdateContext: resourceFlowspecDscpUpdate,
		ReadContext:   resourceFlowspecDscpRead,
		DeleteContext: resourceFlowspecDscpDelete,

		Schema: map[string]*schema.Schema{
			"dscp_attribute": {
				Type: schema.TypeString, Required: true, Description: "'eq': Match only packets on a given DSCP; 'gt': Match only packets with a greater DSCP; 'lt': Match only packets with a lower DSCP; 'range': match only packets in the range of DSCPs;",
			},
			"dscp_val": {
				Type: schema.TypeInt, Required: true, Description: "Specify the DSCP value",
			},
			"dscp_val_end": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the DSCP value",
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
func resourceFlowspecDscpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecDscpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecDscp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFlowspecDscpRead(ctx, d, meta)
	}
	return diags
}

func resourceFlowspecDscpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecDscpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecDscp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFlowspecDscpRead(ctx, d, meta)
	}
	return diags
}
func resourceFlowspecDscpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecDscpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecDscp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFlowspecDscpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecDscpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecDscp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFlowspecDscp(d *schema.ResourceData) edpt.FlowspecDscp {
	var ret edpt.FlowspecDscp
	ret.Inst.DscpAttribute = d.Get("dscp_attribute").(string)
	ret.Inst.DscpVal = d.Get("dscp_val").(int)
	ret.Inst.DscpValEnd = d.Get("dscp_val_end").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
