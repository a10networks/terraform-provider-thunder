package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateGtpMessageFilteringPolicyVersionV1() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_template_gtp_message_filtering_policy_version_v1`: Configure Message Filtering Policy for GTPv1 Control Messages,\n\n__PLACEHOLDER__",
		CreateContext: resourceTemplateGtpMessageFilteringPolicyVersionV1Create,
		UpdateContext: resourceTemplateGtpMessageFilteringPolicyVersionV1Update,
		ReadContext:   resourceTemplateGtpMessageFilteringPolicyVersionV1Read,
		DeleteContext: resourceTemplateGtpMessageFilteringPolicyVersionV1Delete,

		Schema: map[string]*schema.Schema{
			"create_mbms": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"create_pdp": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"delete_mbms": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
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
			"initiate_pdp": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"mbms_deregistration": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"mbms_notification": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"mbms_registration": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"mbms_session": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"message_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the Message Type",
			},
			"ms_info_change": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"pdu_notification": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"reserved_messages": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"update_mbms": {
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
func resourceTemplateGtpMessageFilteringPolicyVersionV1Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpMessageFilteringPolicyVersionV1Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpMessageFilteringPolicyVersionV1(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateGtpMessageFilteringPolicyVersionV1Read(ctx, d, meta)
	}
	return diags
}

func resourceTemplateGtpMessageFilteringPolicyVersionV1Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpMessageFilteringPolicyVersionV1Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpMessageFilteringPolicyVersionV1(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateGtpMessageFilteringPolicyVersionV1Read(ctx, d, meta)
	}
	return diags
}
func resourceTemplateGtpMessageFilteringPolicyVersionV1Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpMessageFilteringPolicyVersionV1Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpMessageFilteringPolicyVersionV1(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTemplateGtpMessageFilteringPolicyVersionV1Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpMessageFilteringPolicyVersionV1Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpMessageFilteringPolicyVersionV1(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointTemplateGtpMessageFilteringPolicyVersionV1(d *schema.ResourceData) edpt.TemplateGtpMessageFilteringPolicyVersionV1 {
	var ret edpt.TemplateGtpMessageFilteringPolicyVersionV1
	ret.Inst.CreateMbms = d.Get("create_mbms").(string)
	ret.Inst.CreatePdp = d.Get("create_pdp").(string)
	ret.Inst.DeleteMbms = d.Get("delete_mbms").(string)
	ret.Inst.DeletePdp = d.Get("delete_pdp").(string)
	ret.Inst.EnableDisableAction = d.Get("enable_disable_action").(string)
	ret.Inst.GtpPdu = d.Get("gtp_pdu").(string)
	ret.Inst.InitiatePdp = d.Get("initiate_pdp").(string)
	ret.Inst.MbmsDeregistration = d.Get("mbms_deregistration").(string)
	ret.Inst.MbmsNotification = d.Get("mbms_notification").(string)
	ret.Inst.MbmsRegistration = d.Get("mbms_registration").(string)
	ret.Inst.MbmsSession = d.Get("mbms_session").(string)
	ret.Inst.MessageType = d.Get("message_type").(int)
	ret.Inst.MsInfoChange = d.Get("ms_info_change").(string)
	ret.Inst.PduNotification = d.Get("pdu_notification").(string)
	ret.Inst.ReservedMessages = d.Get("reserved_messages").(string)
	ret.Inst.UpdateMbms = d.Get("update_mbms").(string)
	ret.Inst.UpdatePdp = d.Get("update_pdp").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
