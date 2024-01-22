package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateSmtp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_smtp`: SMTP\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateSmtpCreate,
		UpdateContext: resourceSlbTemplateSmtpUpdate,
		ReadContext:   resourceSlbTemplateSmtpRead,
		DeleteContext: resourceSlbTemplateSmtpDelete,

		Schema: map[string]*schema.Schema{
			"client_domain_switching": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"switching_type": {
							Type: schema.TypeString, Optional: true, Description: "'contains': Specify domain name string if domain contains another string; 'ends-with': Specify domain name string if domain ends with another string; 'starts-with': Specify domain string if domain starts with another string;",
						},
						"match_string": {
							Type: schema.TypeString, Optional: true, Description: "Domain name string",
						},
						"service_group": {
							Type: schema.TypeString, Optional: true, Description: "Select service group (Service group name)",
						},
					},
				},
			},
			"client_starttls_type": {
				Type: schema.TypeString, Optional: true, Description: "'optional': STARTTLS is optional requirement; 'enforced': Must issue STARTTLS command before mail transaction;",
			},
			"command_disable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disable_type": {
							Type: schema.TypeString, Optional: true, Description: "'expn': Disable SMTP EXPN commands; 'turn': Disable SMTP TURN commands; 'vrfy': Disable SMTP VRFY commands;",
						},
					},
				},
			},
			"error_code_to_client": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Would transfer error code(554) to client, when getting it from connection establishing with real-server",
			},
			"lf_to_crlf": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Change the LF to CRLF for smtp end of line",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "SMTP Template Name",
			},
			"server_domain": {
				Type: schema.TypeString, Optional: true, Default: "mail-server-domain", Description: "Config the domain of the email servers (Server's domain, default is \"mail-server-domain\")",
			},
			"server_starttls_type": {
				Type: schema.TypeString, Optional: true, Description: "'optional': STARTTLS is optional requirement; 'enforced': Must issue STARTTLS command before mail transaction;",
			},
			"service_ready_msg": {
				Type: schema.TypeString, Optional: true, Default: "ESMTP mail service ready", Description: "Set SMTP service ready message (SMTP service ready message, default is \"ESMTP mail service ready\")",
			},
			"template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logging": {
							Type: schema.TypeString, Optional: true, Description: "Logging template (Logging Config name)",
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
func resourceSlbTemplateSmtpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateSmtpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateSmtp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateSmtpRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateSmtpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateSmtpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateSmtp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateSmtpRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateSmtpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateSmtpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateSmtp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateSmtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateSmtpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateSmtp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateSmtpClientDomainSwitching(d []interface{}) []edpt.SlbTemplateSmtpClientDomainSwitching {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateSmtpClientDomainSwitching, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateSmtpClientDomainSwitching
		oi.SwitchingType = in["switching_type"].(string)
		oi.MatchString = in["match_string"].(string)
		oi.ServiceGroup = in["service_group"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateSmtpCommandDisable(d []interface{}) []edpt.SlbTemplateSmtpCommandDisable {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateSmtpCommandDisable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateSmtpCommandDisable
		oi.DisableType = in["disable_type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateSmtpTemplate(d []interface{}) edpt.SlbTemplateSmtpTemplate {

	count1 := len(d)
	var ret edpt.SlbTemplateSmtpTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func dataToEndpointSlbTemplateSmtp(d *schema.ResourceData) edpt.SlbTemplateSmtp {
	var ret edpt.SlbTemplateSmtp
	ret.Inst.ClientDomainSwitching = getSliceSlbTemplateSmtpClientDomainSwitching(d.Get("client_domain_switching").([]interface{}))
	ret.Inst.ClientStarttlsType = d.Get("client_starttls_type").(string)
	ret.Inst.CommandDisable = getSliceSlbTemplateSmtpCommandDisable(d.Get("command_disable").([]interface{}))
	ret.Inst.ErrorCodeToClient = d.Get("error_code_to_client").(int)
	ret.Inst.LfToCrlf = d.Get("lf_to_crlf").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.ServerDomain = d.Get("server_domain").(string)
	ret.Inst.ServerStarttlsType = d.Get("server_starttls_type").(string)
	ret.Inst.ServiceReadyMsg = d.Get("service_ready_msg").(string)
	ret.Inst.Template = getObjectSlbTemplateSmtpTemplate(d.Get("template").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
