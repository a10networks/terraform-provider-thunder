package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateGtpMessageFilteringPolicyVersionV0() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_template_gtp_message_filtering_policy_version_v0`: Configure Message Filtering Policy for GTPv0 Control Messages,\n\n__PLACEHOLDER__",
		CreateContext: resourceTemplateGtpMessageFilteringPolicyVersionV0Create,
		UpdateContext: resourceTemplateGtpMessageFilteringPolicyVersionV0Update,
		ReadContext:   resourceTemplateGtpMessageFilteringPolicyVersionV0Read,
		DeleteContext: resourceTemplateGtpMessageFilteringPolicyVersionV0Delete,

		Schema: map[string]*schema.Schema{
			"create_aa_pdp": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"create_pdp": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"delete_aa_pdp": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"delete_pdp": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"enable_disable_action": {
				Type: schema.TypeString, Required: true, Description: "'enable': Enable Message Filtering on version; 'disable': Disable Message Filtering on version;",
			},
			"gtp_pdu": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"message_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the Message Type",
			},
			"pdu_notification": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"reserved_messages": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"update_pdp": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
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
func resourceTemplateGtpMessageFilteringPolicyVersionV0Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpMessageFilteringPolicyVersionV0Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpMessageFilteringPolicyVersionV0(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateGtpMessageFilteringPolicyVersionV0Read(ctx, d, meta)
	}
	return diags
}

func resourceTemplateGtpMessageFilteringPolicyVersionV0Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpMessageFilteringPolicyVersionV0Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpMessageFilteringPolicyVersionV0(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateGtpMessageFilteringPolicyVersionV0Read(ctx, d, meta)
	}
	return diags
}
func resourceTemplateGtpMessageFilteringPolicyVersionV0Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpMessageFilteringPolicyVersionV0Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpMessageFilteringPolicyVersionV0(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTemplateGtpMessageFilteringPolicyVersionV0Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpMessageFilteringPolicyVersionV0Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpMessageFilteringPolicyVersionV0(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointTemplateGtpMessageFilteringPolicyVersionV0(d *schema.ResourceData) edpt.TemplateGtpMessageFilteringPolicyVersionV0 {
	var ret edpt.TemplateGtpMessageFilteringPolicyVersionV0
	ret.Inst.CreateAaPdp = d.Get("create_aa_pdp").(string)
	ret.Inst.CreatePdp = d.Get("create_pdp").(string)
	ret.Inst.DeleteAaPdp = d.Get("delete_aa_pdp").(string)
	ret.Inst.DeletePdp = d.Get("delete_pdp").(string)
	ret.Inst.EnableDisableAction = d.Get("enable_disable_action").(string)
	ret.Inst.GtpPdu = d.Get("gtp_pdu").(string)
	ret.Inst.MessageType = d.Get("message_type").(int)
	ret.Inst.PduNotification = d.Get("pdu_notification").(string)
	ret.Inst.ReservedMessages = d.Get("reserved_messages").(string)
	ret.Inst.UpdatePdp = d.Get("update_pdp").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
