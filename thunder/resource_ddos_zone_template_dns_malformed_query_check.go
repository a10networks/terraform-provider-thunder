package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateDnsMalformedQueryCheck() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_dns_malformed_query_check`: DNS Malform Query check options\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateDnsMalformedQueryCheckCreate,
		UpdateContext: resourceDdosZoneTemplateDnsMalformedQueryCheckUpdate,
		ReadContext:   resourceDdosZoneTemplateDnsMalformedQueryCheckRead,
		DeleteContext: resourceDdosZoneTemplateDnsMalformedQueryCheckDelete,

		Schema: map[string]*schema.Schema{
			"dns_malformed_query_action": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'blacklist-src': Blacklist-src; 'reset': Reset client connection;",
			},
			"dns_malformed_query_action_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
			},
			"non_query_opcode_check": {
				Type: schema.TypeString, Optional: true, Description: "'disable': When malform check is enabled, TPS always drops DNS query with non query opcode, this option disables this opcode check;",
			},
			"skip_multi_packet_check": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bypass DNS fragmented and TCP segmented Queries(Default: dropped)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"validation_type": {
				Type: schema.TypeString, Optional: true, Description: "'basic-header-check': Basic header validation for DNS TCP/UDP queries; 'extended-header-check': Extended header/query validation for DNS TCP/UDP queries; 'disable': Disable Malform query validation for DNS TCP/UDP;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceDdosZoneTemplateDnsMalformedQueryCheckCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateDnsMalformedQueryCheckCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateDnsMalformedQueryCheck(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateDnsMalformedQueryCheckRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateDnsMalformedQueryCheckUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateDnsMalformedQueryCheckUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateDnsMalformedQueryCheck(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateDnsMalformedQueryCheckRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateDnsMalformedQueryCheckDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateDnsMalformedQueryCheckDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateDnsMalformedQueryCheck(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateDnsMalformedQueryCheckRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateDnsMalformedQueryCheckRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateDnsMalformedQueryCheck(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosZoneTemplateDnsMalformedQueryCheck(d *schema.ResourceData) edpt.DdosZoneTemplateDnsMalformedQueryCheck {
	var ret edpt.DdosZoneTemplateDnsMalformedQueryCheck
	ret.Inst.DnsMalformedQueryAction = d.Get("dns_malformed_query_action").(string)
	ret.Inst.DnsMalformedQueryActionListName = d.Get("dns_malformed_query_action_list_name").(string)
	ret.Inst.NonQueryOpcodeCheck = d.Get("non_query_opcode_check").(string)
	ret.Inst.SkipMultiPacketCheck = d.Get("skip_multi_packet_check").(int)
	//omit uuid
	ret.Inst.ValidationType = d.Get("validation_type").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
