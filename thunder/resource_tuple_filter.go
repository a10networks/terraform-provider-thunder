package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTupleFilter() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_tuple_filter`: Specify a tuple filter\n\n__PLACEHOLDER__",
		CreateContext: resourceTupleFilterCreate,
		UpdateContext: resourceTupleFilterUpdate,
		ReadContext:   resourceTupleFilterRead,
		DeleteContext: resourceTupleFilterDelete,

		Schema: map[string]*schema.Schema{
			"filter_rule_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id1": {
							Type: schema.TypeInt, Required: true, Description: "filter rule id",
						},
						"src_addr": {
							Type: schema.TypeString, Optional: true, Description: "Source IPv4 address with prefix",
						},
						"dst_addr": {
							Type: schema.TypeString, Optional: true, Description: "Destination IPv4 address with prefix",
						},
						"src_port": {
							Type: schema.TypeInt, Optional: true, Description: "Source port",
						},
						"dst_port": {
							Type: schema.TypeInt, Optional: true, Description: "Destination port",
						},
						"src_v6_addr": {
							Type: schema.TypeString, Optional: true, Description: "Source IPv6 address with prefix",
						},
						"dst_v6_addr": {
							Type: schema.TypeString, Optional: true, Description: "Destination IPv6 address with prefix",
						},
						"src_v6_port": {
							Type: schema.TypeInt, Optional: true, Description: "Source port",
						},
						"dst_v6_port": {
							Type: schema.TypeInt, Optional: true, Description: "Destination port",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"filter_type": {
				Type: schema.TypeString, Required: true, Description: "'ipv4': IPv4 tuple filter; 'ipv6': IPv6 tuple filter;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Tuple filter name",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceTupleFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTupleFilterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTupleFilter(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTupleFilterRead(ctx, d, meta)
	}
	return diags
}

func resourceTupleFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTupleFilterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTupleFilter(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTupleFilterRead(ctx, d, meta)
	}
	return diags
}
func resourceTupleFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTupleFilterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTupleFilter(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTupleFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTupleFilterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTupleFilter(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceTupleFilterFilterRuleList(d []interface{}) []edpt.TupleFilterFilterRuleList {

	count1 := len(d)
	ret := make([]edpt.TupleFilterFilterRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.TupleFilterFilterRuleList
		oi.Id1 = in["id1"].(int)
		oi.SrcAddr = in["src_addr"].(string)
		oi.DstAddr = in["dst_addr"].(string)
		oi.SrcPort = in["src_port"].(int)
		oi.DstPort = in["dst_port"].(int)
		oi.SrcV6Addr = in["src_v6_addr"].(string)
		oi.DstV6Addr = in["dst_v6_addr"].(string)
		oi.SrcV6Port = in["src_v6_port"].(int)
		oi.DstV6Port = in["dst_v6_port"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointTupleFilter(d *schema.ResourceData) edpt.TupleFilter {
	var ret edpt.TupleFilter
	ret.Inst.FilterRuleList = getSliceTupleFilterFilterRuleList(d.Get("filter_rule_list").([]interface{}))
	ret.Inst.FilterType = d.Get("filter_type").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
