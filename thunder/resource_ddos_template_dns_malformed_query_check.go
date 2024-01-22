package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateDnsMalformedQueryCheck() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_dns_malformed_query_check`: DNS Malform Query check options\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateDnsMalformedQueryCheckCreate,
		UpdateContext: resourceDdosTemplateDnsMalformedQueryCheckUpdate,
		ReadContext:   resourceDdosTemplateDnsMalformedQueryCheckRead,
		DeleteContext: resourceDdosTemplateDnsMalformedQueryCheckDelete,

		Schema: map[string]*schema.Schema{
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
func resourceDdosTemplateDnsMalformedQueryCheckCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateDnsMalformedQueryCheckCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateDnsMalformedQueryCheck(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateDnsMalformedQueryCheckRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateDnsMalformedQueryCheckUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateDnsMalformedQueryCheckUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateDnsMalformedQueryCheck(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateDnsMalformedQueryCheckRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateDnsMalformedQueryCheckDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateDnsMalformedQueryCheckDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateDnsMalformedQueryCheck(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateDnsMalformedQueryCheckRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateDnsMalformedQueryCheckRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateDnsMalformedQueryCheck(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosTemplateDnsMalformedQueryCheck(d *schema.ResourceData) edpt.DdosTemplateDnsMalformedQueryCheck {
	var ret edpt.DdosTemplateDnsMalformedQueryCheck
	ret.Inst.NonQueryOpcodeCheck = d.Get("non_query_opcode_check").(string)
	ret.Inst.SkipMultiPacketCheck = d.Get("skip_multi_packet_check").(int)
	//omit uuid
	ret.Inst.ValidationType = d.Get("validation_type").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
