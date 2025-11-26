package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateSslL4SslHandshakePolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_ssl_l4_ssl_handshake_policy`: SSL Handshake Policy Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateSslL4SslHandshakePolicyCreate,
		UpdateContext: resourceDdosTemplateSslL4SslHandshakePolicyUpdate,
		ReadContext:   resourceDdosTemplateSslL4SslHandshakePolicyRead,
		DeleteContext: resourceDdosTemplateSslL4SslHandshakePolicyDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'reset': Reset client connection; 'blacklist-src': Blacklist source IP;",
			},
			"cipher_suites_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Set cipher suites limit",
			},
			"client_extensions_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Set client extensions limit",
			},
			"clienthello_to_appdata_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "Set maximum timeout seconds from ClientHello to Application-Data",
			},
			"finished_to_appdata_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "Set maximum timeout seconds from Handshake finished to Application-Data",
			},
			"src_handshaking_conn_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Set handshaking connection limit",
			},
			"ssl_handshake_policy_action_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ssl_l4_tmpl_name": {
				Type: schema.TypeString, Required: true, Description: "SslL4TmplName",
			},
		},
	}
}
func resourceDdosTemplateSslL4SslHandshakePolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateSslL4SslHandshakePolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateSslL4SslHandshakePolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateSslL4SslHandshakePolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateSslL4SslHandshakePolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateSslL4SslHandshakePolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateSslL4SslHandshakePolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateSslL4SslHandshakePolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateSslL4SslHandshakePolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateSslL4SslHandshakePolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateSslL4SslHandshakePolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateSslL4SslHandshakePolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateSslL4SslHandshakePolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateSslL4SslHandshakePolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosTemplateSslL4SslHandshakePolicy(d *schema.ResourceData) edpt.DdosTemplateSslL4SslHandshakePolicy {
	var ret edpt.DdosTemplateSslL4SslHandshakePolicy
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.CipherSuitesLimit = d.Get("cipher_suites_limit").(int)
	ret.Inst.ClientExtensionsLimit = d.Get("client_extensions_limit").(int)
	ret.Inst.ClienthelloToAppdataTimeout = d.Get("clienthello_to_appdata_timeout").(int)
	ret.Inst.FinishedToAppdataTimeout = d.Get("finished_to_appdata_timeout").(int)
	ret.Inst.SrcHandshakingConnLimit = d.Get("src_handshaking_conn_limit").(int)
	ret.Inst.SslHandshakePolicyActionListName = d.Get("ssl_handshake_policy_action_list_name").(string)
	//omit uuid
	ret.Inst.SslL4TmplName = d.Get("ssl_l4_tmpl_name").(string)
	return ret
}
