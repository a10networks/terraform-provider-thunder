package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateGtpGeneralPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_template_gtp_general_policy`: Configure GTP General Policy\n\n__PLACEHOLDER__",
		CreateContext: resourceTemplateGtpGeneralPolicyCreate,
		UpdateContext: resourceTemplateGtpGeneralPolicyUpdate,
		ReadContext:   resourceTemplateGtpGeneralPolicyRead,
		DeleteContext: resourceTemplateGtpGeneralPolicyDelete,

		Schema: map[string]*schema.Schema{
			"handover_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "Tunnel Inactivity Timeout during Handover in minutes (default: 2 mins)",
			},
			"max_mesg_length_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'monitor': Forward failed packet; 'drop': drop packet failing check(Default);",
			},
			"maximum_message_length": {
				Type: schema.TypeInt, Optional: true, Default: 1500, Description: "Maximum message length for a GTP message in bytes",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of the GTP General Policy",
			},
			"tunnel_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 1440, Description: "Tunnel Inactivity Timeout in minutes (default: 1440 minutes)",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"v0_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'permit': Permit GTP-C version 0; 'drop': Drop GTP-C version 0(Default);",
			},
		},
	}
}
func resourceTemplateGtpGeneralPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpGeneralPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpGeneralPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateGtpGeneralPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceTemplateGtpGeneralPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpGeneralPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpGeneralPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateGtpGeneralPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceTemplateGtpGeneralPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpGeneralPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpGeneralPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTemplateGtpGeneralPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpGeneralPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpGeneralPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointTemplateGtpGeneralPolicy(d *schema.ResourceData) edpt.TemplateGtpGeneralPolicy {
	var ret edpt.TemplateGtpGeneralPolicy
	ret.Inst.HandoverTimeout = d.Get("handover_timeout").(int)
	ret.Inst.MaxMesgLengthAction = d.Get("max_mesg_length_action").(string)
	ret.Inst.MaximumMessageLength = d.Get("maximum_message_length").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.TunnelTimeout = d.Get("tunnel_timeout").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.V0Action = d.Get("v0_action").(string)
	return ret
}
