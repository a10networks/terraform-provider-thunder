package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsRecursiveDnsResolutionLookupOrder() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_recursive_dns_resolution_lookup_order`: Recursive lookup address type order for each DNS query type\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsRecursiveDnsResolutionLookupOrderCreate,
		UpdateContext: resourceSlbTemplateDnsRecursiveDnsResolutionLookupOrderUpdate,
		ReadContext:   resourceSlbTemplateDnsRecursiveDnsResolutionLookupOrderRead,
		DeleteContext: resourceSlbTemplateDnsRecursiveDnsResolutionLookupOrderDelete,

		Schema: map[string]*schema.Schema{
			"query_type": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"str_query_type": {
							Type: schema.TypeString, Optional: true, Description: "'A': Address record; 'AAAA': IPv6 Address record; 'CNAME': Canonical name record; 'MX': Mail exchange record; 'NS': Name server record; 'SRV': Service locator; 'PTR': PTR resource record; 'SOA': Start of authority record; 'TXT': Text record; 'ANY': All cached record;",
						},
						"num_query_type": {
							Type: schema.TypeInt, Optional: true, Description: "Other query type value",
						},
						"order": {
							Type: schema.TypeString, Optional: true, Description: "'ipv4-precede-ipv6': Recursive lookup via IPv4 then IPv6; 'ipv6-precede-ipv4': Recursive lookup via IPv6 then IPv4;",
						},
					},
				},
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
func resourceSlbTemplateDnsRecursiveDnsResolutionLookupOrderCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsRecursiveDnsResolutionLookupOrderCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsRecursiveDnsResolutionLookupOrder(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsRecursiveDnsResolutionLookupOrderRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsRecursiveDnsResolutionLookupOrderUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsRecursiveDnsResolutionLookupOrderUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsRecursiveDnsResolutionLookupOrder(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsRecursiveDnsResolutionLookupOrderRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsRecursiveDnsResolutionLookupOrderDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsRecursiveDnsResolutionLookupOrderDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsRecursiveDnsResolutionLookupOrder(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsRecursiveDnsResolutionLookupOrderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsRecursiveDnsResolutionLookupOrderRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsRecursiveDnsResolutionLookupOrder(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType(d []interface{}) []edpt.SlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType
		oi.StrQueryType = in["str_query_type"].(string)
		oi.NumQueryType = in["num_query_type"].(int)
		oi.Order = in["order"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateDnsRecursiveDnsResolutionLookupOrder(d *schema.ResourceData) edpt.SlbTemplateDnsRecursiveDnsResolutionLookupOrder {
	var ret edpt.SlbTemplateDnsRecursiveDnsResolutionLookupOrder
	ret.Inst.QueryType = getSliceSlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType(d.Get("query_type").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
