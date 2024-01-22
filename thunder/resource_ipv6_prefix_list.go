package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6PrefixList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ipv6_prefix_list`: Configure Prefix-list\n\n__PLACEHOLDER__",
		CreateContext: resourceIpv6PrefixListCreate,
		UpdateContext: resourceIpv6PrefixListUpdate,
		ReadContext:   resourceIpv6PrefixListRead,
		DeleteContext: resourceIpv6PrefixListDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name of a prefix list",
			},
			"rules": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"seq": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence number of an entry",
						},
						"description": {
							Type: schema.TypeString, Optional: true, Description: "Prefix-list specific description (Up to 80 characters describing this prefix-list)",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Description: "'deny': Specify packets to reject; 'permit': Specify packets to forward;",
						},
						"any": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Any prefix match. Same as \"::0/0 le 128\"",
						},
						"ipaddr": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 prefix, e.g., 3ffe::/16",
						},
						"ge": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum prefix length to be matched",
						},
						"le": {
							Type: schema.TypeInt, Optional: true, Description: "Maximum prefix length to be matched",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpv6PrefixListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6PrefixListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6PrefixList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6PrefixListRead(ctx, d, meta)
	}
	return diags
}

func resourceIpv6PrefixListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6PrefixListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6PrefixList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6PrefixListRead(ctx, d, meta)
	}
	return diags
}
func resourceIpv6PrefixListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6PrefixListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6PrefixList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpv6PrefixListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6PrefixListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6PrefixList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpv6PrefixListRules(d []interface{}) []edpt.Ipv6PrefixListRules {

	count1 := len(d)
	ret := make([]edpt.Ipv6PrefixListRules, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Ipv6PrefixListRules
		oi.Seq = in["seq"].(int)
		oi.Description = in["description"].(string)
		oi.Action = in["action"].(string)
		oi.Any = in["any"].(int)
		oi.Ipaddr = in["ipaddr"].(string)
		oi.Ge = in["ge"].(int)
		oi.Le = in["le"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpv6PrefixList(d *schema.ResourceData) edpt.Ipv6PrefixList {
	var ret edpt.Ipv6PrefixList
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Rules = getSliceIpv6PrefixListRules(d.Get("rules").([]interface{}))
	//omit uuid
	return ret
}
