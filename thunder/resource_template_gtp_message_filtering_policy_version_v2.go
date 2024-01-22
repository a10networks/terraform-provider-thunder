package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateGtpMessageFilteringPolicyVersionV2() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_template_gtp_message_filtering_policy_version_v2`: Configure Message Filtering Policy for GTPv2 Control Messages\n\n__PLACEHOLDER__",
		CreateContext: resourceTemplateGtpMessageFilteringPolicyVersionV2Create,
		UpdateContext: resourceTemplateGtpMessageFilteringPolicyVersionV2Update,
		ReadContext:   resourceTemplateGtpMessageFilteringPolicyVersionV2Read,
		DeleteContext: resourceTemplateGtpMessageFilteringPolicyVersionV2Delete,

		Schema: map[string]*schema.Schema{
			"bearer_resource": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"change_notification": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"create_bearer": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"create_session": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"delete_bearer": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"delete_command": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"delete_pdn": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"delete_session": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"enable_disable_action": {
				Type: schema.TypeString, Required: true, Description: "'enable': Enable Message Filtering on version; 'disable': Disable Message Filtering on version;",
			},
			"message_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the Message Type",
			},
			"modify_bearer": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"modify_command": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"pgw_downlink_trigger": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"remote_ue_report": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"reserved_messages": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"resume": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"suspend": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"trace_session": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"update_bearer": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
			},
			"update_pdn": {
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
func resourceTemplateGtpMessageFilteringPolicyVersionV2Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpMessageFilteringPolicyVersionV2Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpMessageFilteringPolicyVersionV2(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateGtpMessageFilteringPolicyVersionV2Read(ctx, d, meta)
	}
	return diags
}

func resourceTemplateGtpMessageFilteringPolicyVersionV2Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpMessageFilteringPolicyVersionV2Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpMessageFilteringPolicyVersionV2(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateGtpMessageFilteringPolicyVersionV2Read(ctx, d, meta)
	}
	return diags
}
func resourceTemplateGtpMessageFilteringPolicyVersionV2Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpMessageFilteringPolicyVersionV2Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpMessageFilteringPolicyVersionV2(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTemplateGtpMessageFilteringPolicyVersionV2Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpMessageFilteringPolicyVersionV2Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpMessageFilteringPolicyVersionV2(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointTemplateGtpMessageFilteringPolicyVersionV2(d *schema.ResourceData) edpt.TemplateGtpMessageFilteringPolicyVersionV2 {
	var ret edpt.TemplateGtpMessageFilteringPolicyVersionV2
	ret.Inst.BearerResource = d.Get("bearer_resource").(string)
	ret.Inst.ChangeNotification = d.Get("change_notification").(string)
	ret.Inst.CreateBearer = d.Get("create_bearer").(string)
	ret.Inst.CreateSession = d.Get("create_session").(string)
	ret.Inst.DeleteBearer = d.Get("delete_bearer").(string)
	ret.Inst.DeleteCommand = d.Get("delete_command").(string)
	ret.Inst.DeletePdn = d.Get("delete_pdn").(string)
	ret.Inst.DeleteSession = d.Get("delete_session").(string)
	ret.Inst.EnableDisableAction = d.Get("enable_disable_action").(string)
	ret.Inst.MessageType = d.Get("message_type").(int)
	ret.Inst.ModifyBearer = d.Get("modify_bearer").(string)
	ret.Inst.ModifyCommand = d.Get("modify_command").(string)
	ret.Inst.PgwDownlinkTrigger = d.Get("pgw_downlink_trigger").(string)
	ret.Inst.RemoteUeReport = d.Get("remote_ue_report").(string)
	ret.Inst.ReservedMessages = d.Get("reserved_messages").(string)
	ret.Inst.Resume = d.Get("resume").(string)
	ret.Inst.Suspend = d.Get("suspend").(string)
	ret.Inst.TraceSession = d.Get("trace_session").(string)
	ret.Inst.UpdateBearer = d.Get("update_bearer").(string)
	ret.Inst.UpdatePdn = d.Get("update_pdn").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
