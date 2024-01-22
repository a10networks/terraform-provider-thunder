package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateGtpMessageFilteringPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_template_gtp_message_filtering_policy`: Configure GTP Message Filtering Policy\n\n__PLACEHOLDER__",
		CreateContext: resourceTemplateGtpMessageFilteringPolicyCreate,
		UpdateContext: resourceTemplateGtpMessageFilteringPolicyUpdate,
		ReadContext:   resourceTemplateGtpMessageFilteringPolicyRead,
		DeleteContext: resourceTemplateGtpMessageFilteringPolicyDelete,

		Schema: map[string]*schema.Schema{
			"interface_type": {
				Type: schema.TypeString, Optional: true, Description: "'roaming': Roaming Interface(S8/Gp); 'non-roaming': Non-Roaming Interface(S5/Gn);",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of the GTP Message Filtering Policy",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"version_v0": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_disable_action": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable Message Filtering on version; 'disable': Disable Message Filtering on version;",
						},
						"message_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the Message Type",
						},
						"create_pdp": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"update_pdp": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"delete_pdp": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"pdu_notification": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"gtp_pdu": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"create_aa_pdp": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"delete_aa_pdp": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"reserved_messages": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"version_v1": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_disable_action": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable Message Filtering on version; 'disable': Disable Message Filtering on version;",
						},
						"message_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the Message Type",
						},
						"create_pdp": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"update_pdp": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"delete_pdp": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"initiate_pdp": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"pdu_notification": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"ms_info_change": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"gtp_pdu": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"mbms_session": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"mbms_notification": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"mbms_registration": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"mbms_deregistration": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"create_mbms": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"delete_mbms": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"update_mbms": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"reserved_messages": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"version_v2": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_disable_action": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable Message Filtering on version; 'disable': Disable Message Filtering on version;",
						},
						"message_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the Message Type",
						},
						"change_notification": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"create_session": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"delete_session": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"modify_bearer": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"remote_ue_report": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"modify_command": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"delete_command": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"bearer_resource": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"create_bearer": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"update_bearer": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"delete_bearer": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"delete_pdn": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"update_pdn": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"suspend": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"resume": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"pgw_downlink_trigger": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"trace_session": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"reserved_messages": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Message Type; 'disable': Disable the Message Type;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
		},
	}
}
func resourceTemplateGtpMessageFilteringPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpMessageFilteringPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpMessageFilteringPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateGtpMessageFilteringPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceTemplateGtpMessageFilteringPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpMessageFilteringPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpMessageFilteringPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateGtpMessageFilteringPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceTemplateGtpMessageFilteringPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpMessageFilteringPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpMessageFilteringPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTemplateGtpMessageFilteringPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpMessageFilteringPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpMessageFilteringPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectTemplateGtpMessageFilteringPolicyVersionV01900(d []interface{}) edpt.TemplateGtpMessageFilteringPolicyVersionV01900 {

	count1 := len(d)
	var ret edpt.TemplateGtpMessageFilteringPolicyVersionV01900
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EnableDisableAction = in["enable_disable_action"].(string)
		ret.MessageType = in["message_type"].(int)
		ret.CreatePdp = in["create_pdp"].(string)
		ret.UpdatePdp = in["update_pdp"].(string)
		ret.DeletePdp = in["delete_pdp"].(string)
		ret.PduNotification = in["pdu_notification"].(string)
		ret.GtpPdu = in["gtp_pdu"].(string)
		ret.CreateAaPdp = in["create_aa_pdp"].(string)
		ret.DeleteAaPdp = in["delete_aa_pdp"].(string)
		ret.ReservedMessages = in["reserved_messages"].(string)
		//omit uuid
	}
	return ret
}

func getObjectTemplateGtpMessageFilteringPolicyVersionV11901(d []interface{}) edpt.TemplateGtpMessageFilteringPolicyVersionV11901 {

	count1 := len(d)
	var ret edpt.TemplateGtpMessageFilteringPolicyVersionV11901
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EnableDisableAction = in["enable_disable_action"].(string)
		ret.MessageType = in["message_type"].(int)
		ret.CreatePdp = in["create_pdp"].(string)
		ret.UpdatePdp = in["update_pdp"].(string)
		ret.DeletePdp = in["delete_pdp"].(string)
		ret.InitiatePdp = in["initiate_pdp"].(string)
		ret.PduNotification = in["pdu_notification"].(string)
		ret.MsInfoChange = in["ms_info_change"].(string)
		ret.GtpPdu = in["gtp_pdu"].(string)
		ret.MbmsSession = in["mbms_session"].(string)
		ret.MbmsNotification = in["mbms_notification"].(string)
		ret.MbmsRegistration = in["mbms_registration"].(string)
		ret.MbmsDeregistration = in["mbms_deregistration"].(string)
		ret.CreateMbms = in["create_mbms"].(string)
		ret.DeleteMbms = in["delete_mbms"].(string)
		ret.UpdateMbms = in["update_mbms"].(string)
		ret.ReservedMessages = in["reserved_messages"].(string)
		//omit uuid
	}
	return ret
}

func getObjectTemplateGtpMessageFilteringPolicyVersionV21902(d []interface{}) edpt.TemplateGtpMessageFilteringPolicyVersionV21902 {

	count1 := len(d)
	var ret edpt.TemplateGtpMessageFilteringPolicyVersionV21902
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EnableDisableAction = in["enable_disable_action"].(string)
		ret.MessageType = in["message_type"].(int)
		ret.ChangeNotification = in["change_notification"].(string)
		ret.CreateSession = in["create_session"].(string)
		ret.DeleteSession = in["delete_session"].(string)
		ret.ModifyBearer = in["modify_bearer"].(string)
		ret.RemoteUeReport = in["remote_ue_report"].(string)
		ret.ModifyCommand = in["modify_command"].(string)
		ret.DeleteCommand = in["delete_command"].(string)
		ret.BearerResource = in["bearer_resource"].(string)
		ret.CreateBearer = in["create_bearer"].(string)
		ret.UpdateBearer = in["update_bearer"].(string)
		ret.DeleteBearer = in["delete_bearer"].(string)
		ret.DeletePdn = in["delete_pdn"].(string)
		ret.UpdatePdn = in["update_pdn"].(string)
		ret.Suspend = in["suspend"].(string)
		ret.Resume = in["resume"].(string)
		ret.PgwDownlinkTrigger = in["pgw_downlink_trigger"].(string)
		ret.TraceSession = in["trace_session"].(string)
		ret.ReservedMessages = in["reserved_messages"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointTemplateGtpMessageFilteringPolicy(d *schema.ResourceData) edpt.TemplateGtpMessageFilteringPolicy {
	var ret edpt.TemplateGtpMessageFilteringPolicy
	ret.Inst.InterfaceType = d.Get("interface_type").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VersionV0 = getObjectTemplateGtpMessageFilteringPolicyVersionV01900(d.Get("version_v0").([]interface{}))
	ret.Inst.VersionV1 = getObjectTemplateGtpMessageFilteringPolicyVersionV11901(d.Get("version_v1").([]interface{}))
	ret.Inst.VersionV2 = getObjectTemplateGtpMessageFilteringPolicyVersionV21902(d.Get("version_v2").([]interface{}))
	return ret
}
