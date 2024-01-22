package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsQueryTypeFilter() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_query_type_filter`: DNS query type filter list\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsQueryTypeFilterCreate,
		UpdateContext: resourceSlbTemplateDnsQueryTypeFilterUpdate,
		ReadContext:   resourceSlbTemplateDnsQueryTypeFilterRead,
		DeleteContext: resourceSlbTemplateDnsQueryTypeFilterDelete,

		Schema: map[string]*schema.Schema{
			"query_type": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"str_query_type": {
							Type: schema.TypeString, Optional: true, Description: "'A': Address record; 'AAAA': IPv6 Address record; 'CNAME': Canonical name record; 'MX': Mail exchange record; 'NS': Name server record; 'SRV': Service locator; 'PTR': PTR resource record; 'SOA': Start of authority record; 'TXT': Text record; 'ANY': All cached record;",
						},
						"num_query_type": {
							Type: schema.TypeInt, Optional: true, Description: "Other record type value",
						},
					},
				},
			},
			"query_type_action": {
				Type: schema.TypeString, Required: true, Description: "'allow': Allow only certain DNS query types; 'deny': Deny only certain DNS query types;",
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
func resourceSlbTemplateDnsQueryTypeFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsQueryTypeFilterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsQueryTypeFilter(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsQueryTypeFilterRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsQueryTypeFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsQueryTypeFilterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsQueryTypeFilter(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsQueryTypeFilterRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsQueryTypeFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsQueryTypeFilterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsQueryTypeFilter(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsQueryTypeFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsQueryTypeFilterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsQueryTypeFilter(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateDnsQueryTypeFilterQueryType(d []interface{}) []edpt.SlbTemplateDnsQueryTypeFilterQueryType {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsQueryTypeFilterQueryType, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsQueryTypeFilterQueryType
		oi.StrQueryType = in["str_query_type"].(string)
		oi.NumQueryType = in["num_query_type"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateDnsQueryTypeFilter(d *schema.ResourceData) edpt.SlbTemplateDnsQueryTypeFilter {
	var ret edpt.SlbTemplateDnsQueryTypeFilter
	ret.Inst.QueryType = getSliceSlbTemplateDnsQueryTypeFilterQueryType(d.Get("query_type").([]interface{}))
	ret.Inst.QueryTypeAction = d.Get("query_type_action").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
