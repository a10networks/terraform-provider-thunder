package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSrcPortTemplateDns() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_src_port_template_dns`: DNS template Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosSrcPortTemplateDnsCreate,
		UpdateContext: resourceDdosSrcPortTemplateDnsUpdate,
		ReadContext:   resourceDdosSrcPortTemplateDnsRead,
		DeleteContext: resourceDdosSrcPortTemplateDnsDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "",
			},
			"query_resolution_check": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"session_timeout_value": {
							Type: schema.TypeInt, Optional: true, Description: "max session timeout (secs) between DNS external server and Protected object",
						},
						"domain_lockup_action": {
							Type: schema.TypeString, Optional: true, Default: "default", Description: "'default': Default, No action for future connections; 'blacklist-src': Blacklist the external server for future connections;",
						},
						"big_response_size": {
							Type: schema.TypeInt, Optional: true, Description: "Max DNS response size (in Bytes)",
						},
						"big_response_action": {
							Type: schema.TypeString, Optional: true, Default: "default", Description: "'default': Default, No action for future connections; 'blacklist-src': Blacklist the external server for future connections;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
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
func resourceDdosSrcPortTemplateDnsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcPortTemplateDnsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcPortTemplateDns(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcPortTemplateDnsRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosSrcPortTemplateDnsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcPortTemplateDnsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcPortTemplateDns(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcPortTemplateDnsRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosSrcPortTemplateDnsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcPortTemplateDnsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcPortTemplateDns(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosSrcPortTemplateDnsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcPortTemplateDnsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcPortTemplateDns(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosSrcPortTemplateDnsQueryResolutionCheck294(d []interface{}) edpt.DdosSrcPortTemplateDnsQueryResolutionCheck294 {

	count1 := len(d)
	var ret edpt.DdosSrcPortTemplateDnsQueryResolutionCheck294
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionTimeoutValue = in["session_timeout_value"].(int)
		ret.DomainLockupAction = in["domain_lockup_action"].(string)
		ret.BigResponseSize = in["big_response_size"].(int)
		ret.BigResponseAction = in["big_response_action"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointDdosSrcPortTemplateDns(d *schema.ResourceData) edpt.DdosSrcPortTemplateDns {
	var ret edpt.DdosSrcPortTemplateDns
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.QueryResolutionCheck = getObjectDdosSrcPortTemplateDnsQueryResolutionCheck294(d.Get("query_resolution_check").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
