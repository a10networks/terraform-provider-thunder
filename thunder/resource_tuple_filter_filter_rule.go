package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTupleFilterFilterRule() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_tuple_filter_filter_rule`: Specify filter matching rule\n\n__PLACEHOLDER__",
		CreateContext: resourceTupleFilterFilterRuleCreate,
		UpdateContext: resourceTupleFilterFilterRuleUpdate,
		ReadContext:   resourceTupleFilterFilterRuleRead,
		DeleteContext: resourceTupleFilterFilterRuleDelete,

		Schema: map[string]*schema.Schema{
			"dst_addr": {
				Type: schema.TypeString, Optional: true, Description: "Destination IPv4 address with prefix",
			},
			"dst_port": {
				Type: schema.TypeInt, Optional: true, Description: "Destination port",
			},
			"dst_v6_addr": {
				Type: schema.TypeString, Optional: true, Description: "Destination IPv6 address with prefix",
			},
			"dst_v6_port": {
				Type: schema.TypeInt, Optional: true, Description: "Destination port",
			},
			"id1": {
				Type: schema.TypeInt, Required: true, Description: "filter rule id",
			},
			"src_addr": {
				Type: schema.TypeString, Optional: true, Description: "Source IPv4 address with prefix",
			},
			"src_port": {
				Type: schema.TypeInt, Optional: true, Description: "Source port",
			},
			"src_v6_addr": {
				Type: schema.TypeString, Optional: true, Description: "Source IPv6 address with prefix",
			},
			"src_v6_port": {
				Type: schema.TypeInt, Optional: true, Description: "Source port",
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
func resourceTupleFilterFilterRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTupleFilterFilterRuleCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTupleFilterFilterRule(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTupleFilterFilterRuleRead(ctx, d, meta)
	}
	return diags
}

func resourceTupleFilterFilterRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTupleFilterFilterRuleUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTupleFilterFilterRule(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTupleFilterFilterRuleRead(ctx, d, meta)
	}
	return diags
}
func resourceTupleFilterFilterRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTupleFilterFilterRuleDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTupleFilterFilterRule(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTupleFilterFilterRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTupleFilterFilterRuleRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTupleFilterFilterRule(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointTupleFilterFilterRule(d *schema.ResourceData) edpt.TupleFilterFilterRule {
	var ret edpt.TupleFilterFilterRule
	ret.Inst.DstAddr = d.Get("dst_addr").(string)
	ret.Inst.DstPort = d.Get("dst_port").(int)
	ret.Inst.DstV6Addr = d.Get("dst_v6_addr").(string)
	ret.Inst.DstV6Port = d.Get("dst_v6_port").(int)
	ret.Inst.Id1 = d.Get("id1").(int)
	ret.Inst.SrcAddr = d.Get("src_addr").(string)
	ret.Inst.SrcPort = d.Get("src_port").(int)
	ret.Inst.SrcV6Addr = d.Get("src_v6_addr").(string)
	ret.Inst.SrcV6Port = d.Get("src_v6_port").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
