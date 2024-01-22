package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpPrefixList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_prefix_list`: Configure Prefix-list\n\n__PLACEHOLDER__",
		CreateContext: resourceIpPrefixListCreate,
		UpdateContext: resourceIpPrefixListUpdate,
		ReadContext:   resourceIpPrefixListRead,
		DeleteContext: resourceIpPrefixListDelete,

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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Any prefix match. Same as \"0.0.0.0/0 le 32\"",
						},
						"ipaddr": {
							Type: schema.TypeString, Optional: true, Description: "IP prefix, e.g., 35.0.0.0/8",
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
func resourceIpPrefixListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpPrefixListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpPrefixList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpPrefixListRead(ctx, d, meta)
	}
	return diags
}

func resourceIpPrefixListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpPrefixListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpPrefixList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpPrefixListRead(ctx, d, meta)
	}
	return diags
}
func resourceIpPrefixListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpPrefixListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpPrefixList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpPrefixListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpPrefixListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpPrefixList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpPrefixListRules(d []interface{}) []edpt.IpPrefixListRules {

	count1 := len(d)
	ret := make([]edpt.IpPrefixListRules, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpPrefixListRules
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

func dataToEndpointIpPrefixList(d *schema.ResourceData) edpt.IpPrefixList {
	var ret edpt.IpPrefixList
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Rules = getSliceIpPrefixListRules(d.Get("rules").([]interface{}))
	//omit uuid
	return ret
}
