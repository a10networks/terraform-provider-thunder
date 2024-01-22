package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateSip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_sip`: SIP template Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateSipCreate,
		UpdateContext: resourceDdosZoneTemplateSipUpdate,
		ReadContext:   resourceDdosZoneTemplateSipRead,
		DeleteContext: resourceDdosZoneTemplateSipDelete,

		Schema: map[string]*schema.Schema{
			"dst": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sip_request_rate_limit": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dst_sip_rate_action_list_name": {
										Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
									},
									"dst_sip_rate_action": {
										Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets(Default); 'ignore': Take no action; 'reset': Reset (sip-tcp) client connection; 'blacklist-src': Blacklist-src;",
									},
									"method": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"invite_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"invite": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "INVITE method",
															},
															"dst_sip_invite_rate": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
														},
													},
												},
												"register_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"register": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "REGISTER method",
															},
															"dst_sip_register_rate": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
														},
													},
												},
												"options_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"options": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "OPTIONS method",
															},
															"dst_sip_options_rate": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
														},
													},
												},
												"bye_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bye": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "BYE method",
															},
															"dst_sip_bye_rate": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
														},
													},
												},
												"subscribe_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"subscribe": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "SUBSCRIBE method",
															},
															"dst_sip_subscribe_rate": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
														},
													},
												},
												"notify_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"notify": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "NOTIFY method",
															},
															"dst_sip_notify_rate": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
														},
													},
												},
												"refer_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"refer": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "REFER method",
															},
															"dst_sip_refer_rate": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
														},
													},
												},
												"message_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"message": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "MESSAGE method",
															},
															"dst_sip_message_rate": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
														},
													},
												},
												"update_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"update": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "UPDATE method",
															},
															"dst_sip_update_rate": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"filter_header_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sip_filter_name": {
							Type: schema.TypeString, Required: true, Description: "",
						},
						"sip_filter_header_seq": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence number",
						},
						"sip_header_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sip_filter_header_regex": {
										Type: schema.TypeString, Optional: true, Description: "Regex Expression",
									},
									"sip_filter_header_inverse_match": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
									},
								},
							},
						},
						"sip_filter_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
						},
						"sip_filter_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'blacklist-src': Blacklist-src; 'authenticate-src': Authenticate-src; 'reset': Reset client connection(for sip-tcp);",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"idle_timeout": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"idle_timeout_value": {
							Type: schema.TypeInt, Optional: true, Description: "Set the the idle timeout value for SIP-TCP connections",
						},
						"ignore_zero_payload": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Don't reset idle timer on packets with zero payload length from clients",
						},
						"idle_timeout_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
						},
						"idle_timeout_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'blacklist-src': Blacklist-src; 'reset': Reset (sip-tcp) client connection;",
						},
					},
				},
			},
			"malformed_sip": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"malformed_sip_check": {
							Type: schema.TypeString, Optional: true, Description: "'enable-check': Enable malformed SIP parameters;",
						},
						"malformed_sip_max_line_size": {
							Type: schema.TypeInt, Optional: true, Default: 32511, Description: "Set the maximum line size. Default value is 32511",
						},
						"malformed_sip_max_uri_length": {
							Type: schema.TypeInt, Optional: true, Default: 32511, Description: "Set the maximum uri size. Default value is 32511",
						},
						"malformed_sip_max_header_name_length": {
							Type: schema.TypeInt, Optional: true, Default: 63, Description: "Set the maximum header name length. Default value is 63",
						},
						"malformed_sip_max_header_value_length": {
							Type: schema.TypeInt, Optional: true, Default: 32511, Description: "Set the maximum header value length. Default value is 32511",
						},
						"malformed_sip_call_id_max_length": {
							Type: schema.TypeInt, Optional: true, Default: 32511, Description: "Set the maximum call-id length. Default value is 32511",
						},
						"malformed_sip_sdp_max_length": {
							Type: schema.TypeInt, Optional: true, Default: 32511, Description: "Set the maxinum SDP content length. Default value is 32511",
						},
						"malformed_sip_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
						},
						"malformed_sip_action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets (Default); 'reset': Reset (sip-tcp) client connection; 'blacklist-src': Blacklist-src;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"multi_pu_threshold_distribution": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"multi_pu_threshold_distribution_value": {
							Type: schema.TypeInt, Optional: true, Description: "Destination side rate limit only. Default: 0",
						},
						"multi_pu_threshold_distribution_disable": {
							Type: schema.TypeString, Optional: true, Description: "'disable': Destination side rate limit only. Default: Enable;",
						},
					},
				},
			},
			"sip_tmpl_name": {
				Type: schema.TypeString, Required: true, Description: "DDOS SIP Template Name",
			},
			"src": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sip_request_rate_limit": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"src_sip_rate_action_list_name": {
										Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
									},
									"src_sip_rate_action": {
										Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets(Default); 'ignore': Take no action; 'reset': Reset (sip-tcp) client connection; 'blacklist-src': Blacklist-src;",
									},
									"method": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"invite_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"invite": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "INVITE method",
															},
															"src_sip_invite_rate": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
														},
													},
												},
												"register_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"register": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "REGISTER method",
															},
															"src_sip_register_rate": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
														},
													},
												},
												"options_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"options": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "OPTIONS method",
															},
															"src_sip_options_rate": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
														},
													},
												},
												"bye_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bye": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "BYE method",
															},
															"src_sip_bye_rate": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
														},
													},
												},
												"subscribe_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"subscribe": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "SUBSCRIBE method",
															},
															"src_sip_subscribe_rate": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
														},
													},
												},
												"notify_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"notify": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "NOTIFY method",
															},
															"src_sip_notify_rate": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
														},
													},
												},
												"refer_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"refer": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "REFER method",
															},
															"src_sip_refer_rate": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
														},
													},
												},
												"message_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"message": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "MESSAGE method",
															},
															"src_sip_message_rate": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
														},
													},
												},
												"update_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"update": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "UPDATE method",
															},
															"src_sip_update_rate": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
														},
													},
												},
											},
										},
									},
								},
							},
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
func resourceDdosZoneTemplateSipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateSipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateSip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateSipRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateSipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateSipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateSip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateSipRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateSipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateSipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateSip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateSipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateSipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateSip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosZoneTemplateSipDst(d []interface{}) edpt.DdosZoneTemplateSipDst {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipDst
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SipRequestRateLimit = getObjectDdosZoneTemplateSipDstSipRequestRateLimit(in["sip_request_rate_limit"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateSipDstSipRequestRateLimit(d []interface{}) edpt.DdosZoneTemplateSipDstSipRequestRateLimit {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipDstSipRequestRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DstSipRateActionListName = in["dst_sip_rate_action_list_name"].(string)
		ret.DstSipRateAction = in["dst_sip_rate_action"].(string)
		ret.Method = getObjectDdosZoneTemplateSipDstSipRequestRateLimitMethod(in["method"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateSipDstSipRequestRateLimitMethod(d []interface{}) edpt.DdosZoneTemplateSipDstSipRequestRateLimitMethod {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipDstSipRequestRateLimitMethod
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.InviteCfg = getObjectDdosZoneTemplateSipDstSipRequestRateLimitMethodInviteCfg(in["invite_cfg"].([]interface{}))
		ret.RegisterCfg = getObjectDdosZoneTemplateSipDstSipRequestRateLimitMethodRegisterCfg(in["register_cfg"].([]interface{}))
		ret.OptionsCfg = getObjectDdosZoneTemplateSipDstSipRequestRateLimitMethodOptionsCfg(in["options_cfg"].([]interface{}))
		ret.ByeCfg = getObjectDdosZoneTemplateSipDstSipRequestRateLimitMethodByeCfg(in["bye_cfg"].([]interface{}))
		ret.SubscribeCfg = getObjectDdosZoneTemplateSipDstSipRequestRateLimitMethodSubscribeCfg(in["subscribe_cfg"].([]interface{}))
		ret.NotifyCfg = getObjectDdosZoneTemplateSipDstSipRequestRateLimitMethodNotifyCfg(in["notify_cfg"].([]interface{}))
		ret.ReferCfg = getObjectDdosZoneTemplateSipDstSipRequestRateLimitMethodReferCfg(in["refer_cfg"].([]interface{}))
		ret.MessageCfg = getObjectDdosZoneTemplateSipDstSipRequestRateLimitMethodMessageCfg(in["message_cfg"].([]interface{}))
		ret.UpdateCfg = getObjectDdosZoneTemplateSipDstSipRequestRateLimitMethodUpdateCfg(in["update_cfg"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateSipDstSipRequestRateLimitMethodInviteCfg(d []interface{}) edpt.DdosZoneTemplateSipDstSipRequestRateLimitMethodInviteCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipDstSipRequestRateLimitMethodInviteCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Invite = in["invite"].(int)
		ret.DstSipInviteRate = in["dst_sip_invite_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateSipDstSipRequestRateLimitMethodRegisterCfg(d []interface{}) edpt.DdosZoneTemplateSipDstSipRequestRateLimitMethodRegisterCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipDstSipRequestRateLimitMethodRegisterCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Register = in["register"].(int)
		ret.DstSipRegisterRate = in["dst_sip_register_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateSipDstSipRequestRateLimitMethodOptionsCfg(d []interface{}) edpt.DdosZoneTemplateSipDstSipRequestRateLimitMethodOptionsCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipDstSipRequestRateLimitMethodOptionsCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Options = in["options"].(int)
		ret.DstSipOptionsRate = in["dst_sip_options_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateSipDstSipRequestRateLimitMethodByeCfg(d []interface{}) edpt.DdosZoneTemplateSipDstSipRequestRateLimitMethodByeCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipDstSipRequestRateLimitMethodByeCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bye = in["bye"].(int)
		ret.DstSipByeRate = in["dst_sip_bye_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateSipDstSipRequestRateLimitMethodSubscribeCfg(d []interface{}) edpt.DdosZoneTemplateSipDstSipRequestRateLimitMethodSubscribeCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipDstSipRequestRateLimitMethodSubscribeCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Subscribe = in["subscribe"].(int)
		ret.DstSipSubscribeRate = in["dst_sip_subscribe_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateSipDstSipRequestRateLimitMethodNotifyCfg(d []interface{}) edpt.DdosZoneTemplateSipDstSipRequestRateLimitMethodNotifyCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipDstSipRequestRateLimitMethodNotifyCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Notify = in["notify"].(int)
		ret.DstSipNotifyRate = in["dst_sip_notify_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateSipDstSipRequestRateLimitMethodReferCfg(d []interface{}) edpt.DdosZoneTemplateSipDstSipRequestRateLimitMethodReferCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipDstSipRequestRateLimitMethodReferCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Refer = in["refer"].(int)
		ret.DstSipReferRate = in["dst_sip_refer_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateSipDstSipRequestRateLimitMethodMessageCfg(d []interface{}) edpt.DdosZoneTemplateSipDstSipRequestRateLimitMethodMessageCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipDstSipRequestRateLimitMethodMessageCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Message = in["message"].(int)
		ret.DstSipMessageRate = in["dst_sip_message_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateSipDstSipRequestRateLimitMethodUpdateCfg(d []interface{}) edpt.DdosZoneTemplateSipDstSipRequestRateLimitMethodUpdateCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipDstSipRequestRateLimitMethodUpdateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Update = in["update"].(int)
		ret.DstSipUpdateRate = in["dst_sip_update_rate"].(int)
	}
	return ret
}

func getSliceDdosZoneTemplateSipFilterHeaderList(d []interface{}) []edpt.DdosZoneTemplateSipFilterHeaderList {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateSipFilterHeaderList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateSipFilterHeaderList
		oi.SipFilterName = in["sip_filter_name"].(string)
		oi.SipFilterHeaderSeq = in["sip_filter_header_seq"].(int)
		oi.SipHeaderCfg = getObjectDdosZoneTemplateSipFilterHeaderListSipHeaderCfg(in["sip_header_cfg"].([]interface{}))
		oi.SipFilterActionListName = in["sip_filter_action_list_name"].(string)
		oi.SipFilterAction = in["sip_filter_action"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneTemplateSipFilterHeaderListSipHeaderCfg(d []interface{}) edpt.DdosZoneTemplateSipFilterHeaderListSipHeaderCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipFilterHeaderListSipHeaderCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SipFilterHeaderRegex = in["sip_filter_header_regex"].(string)
		ret.SipFilterHeaderInverseMatch = in["sip_filter_header_inverse_match"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateSipIdleTimeout(d []interface{}) edpt.DdosZoneTemplateSipIdleTimeout {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipIdleTimeout
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IdleTimeoutValue = in["idle_timeout_value"].(int)
		ret.IgnoreZeroPayload = in["ignore_zero_payload"].(int)
		ret.IdleTimeoutActionListName = in["idle_timeout_action_list_name"].(string)
		ret.IdleTimeoutAction = in["idle_timeout_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateSipMalformedSip316(d []interface{}) edpt.DdosZoneTemplateSipMalformedSip316 {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipMalformedSip316
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MalformedSipCheck = in["malformed_sip_check"].(string)
		ret.MalformedSipMaxLineSize = in["malformed_sip_max_line_size"].(int)
		ret.MalformedSipMaxUriLength = in["malformed_sip_max_uri_length"].(int)
		ret.MalformedSipMaxHeaderNameLength = in["malformed_sip_max_header_name_length"].(int)
		ret.MalformedSipMaxHeaderValueLength = in["malformed_sip_max_header_value_length"].(int)
		ret.MalformedSipCallIdMaxLength = in["malformed_sip_call_id_max_length"].(int)
		ret.MalformedSipSdpMaxLength = in["malformed_sip_sdp_max_length"].(int)
		ret.MalformedSipActionListName = in["malformed_sip_action_list_name"].(string)
		ret.MalformedSipAction = in["malformed_sip_action"].(string)
		//omit uuid
	}
	return ret
}

func getObjectDdosZoneTemplateSipMultiPuThresholdDistribution(d []interface{}) edpt.DdosZoneTemplateSipMultiPuThresholdDistribution {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipMultiPuThresholdDistribution
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MultiPuThresholdDistributionValue = in["multi_pu_threshold_distribution_value"].(int)
		ret.MultiPuThresholdDistributionDisable = in["multi_pu_threshold_distribution_disable"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateSipSrc(d []interface{}) edpt.DdosZoneTemplateSipSrc {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipSrc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SipRequestRateLimit = getObjectDdosZoneTemplateSipSrcSipRequestRateLimit(in["sip_request_rate_limit"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateSipSrcSipRequestRateLimit(d []interface{}) edpt.DdosZoneTemplateSipSrcSipRequestRateLimit {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipSrcSipRequestRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcSipRateActionListName = in["src_sip_rate_action_list_name"].(string)
		ret.SrcSipRateAction = in["src_sip_rate_action"].(string)
		ret.Method = getObjectDdosZoneTemplateSipSrcSipRequestRateLimitMethod(in["method"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateSipSrcSipRequestRateLimitMethod(d []interface{}) edpt.DdosZoneTemplateSipSrcSipRequestRateLimitMethod {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipSrcSipRequestRateLimitMethod
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.InviteCfg = getObjectDdosZoneTemplateSipSrcSipRequestRateLimitMethodInviteCfg(in["invite_cfg"].([]interface{}))
		ret.RegisterCfg = getObjectDdosZoneTemplateSipSrcSipRequestRateLimitMethodRegisterCfg(in["register_cfg"].([]interface{}))
		ret.OptionsCfg = getObjectDdosZoneTemplateSipSrcSipRequestRateLimitMethodOptionsCfg(in["options_cfg"].([]interface{}))
		ret.ByeCfg = getObjectDdosZoneTemplateSipSrcSipRequestRateLimitMethodByeCfg(in["bye_cfg"].([]interface{}))
		ret.SubscribeCfg = getObjectDdosZoneTemplateSipSrcSipRequestRateLimitMethodSubscribeCfg(in["subscribe_cfg"].([]interface{}))
		ret.NotifyCfg = getObjectDdosZoneTemplateSipSrcSipRequestRateLimitMethodNotifyCfg(in["notify_cfg"].([]interface{}))
		ret.ReferCfg = getObjectDdosZoneTemplateSipSrcSipRequestRateLimitMethodReferCfg(in["refer_cfg"].([]interface{}))
		ret.MessageCfg = getObjectDdosZoneTemplateSipSrcSipRequestRateLimitMethodMessageCfg(in["message_cfg"].([]interface{}))
		ret.UpdateCfg = getObjectDdosZoneTemplateSipSrcSipRequestRateLimitMethodUpdateCfg(in["update_cfg"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateSipSrcSipRequestRateLimitMethodInviteCfg(d []interface{}) edpt.DdosZoneTemplateSipSrcSipRequestRateLimitMethodInviteCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipSrcSipRequestRateLimitMethodInviteCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Invite = in["invite"].(int)
		ret.SrcSipInviteRate = in["src_sip_invite_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateSipSrcSipRequestRateLimitMethodRegisterCfg(d []interface{}) edpt.DdosZoneTemplateSipSrcSipRequestRateLimitMethodRegisterCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipSrcSipRequestRateLimitMethodRegisterCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Register = in["register"].(int)
		ret.SrcSipRegisterRate = in["src_sip_register_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateSipSrcSipRequestRateLimitMethodOptionsCfg(d []interface{}) edpt.DdosZoneTemplateSipSrcSipRequestRateLimitMethodOptionsCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipSrcSipRequestRateLimitMethodOptionsCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Options = in["options"].(int)
		ret.SrcSipOptionsRate = in["src_sip_options_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateSipSrcSipRequestRateLimitMethodByeCfg(d []interface{}) edpt.DdosZoneTemplateSipSrcSipRequestRateLimitMethodByeCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipSrcSipRequestRateLimitMethodByeCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bye = in["bye"].(int)
		ret.SrcSipByeRate = in["src_sip_bye_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateSipSrcSipRequestRateLimitMethodSubscribeCfg(d []interface{}) edpt.DdosZoneTemplateSipSrcSipRequestRateLimitMethodSubscribeCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipSrcSipRequestRateLimitMethodSubscribeCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Subscribe = in["subscribe"].(int)
		ret.SrcSipSubscribeRate = in["src_sip_subscribe_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateSipSrcSipRequestRateLimitMethodNotifyCfg(d []interface{}) edpt.DdosZoneTemplateSipSrcSipRequestRateLimitMethodNotifyCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipSrcSipRequestRateLimitMethodNotifyCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Notify = in["notify"].(int)
		ret.SrcSipNotifyRate = in["src_sip_notify_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateSipSrcSipRequestRateLimitMethodReferCfg(d []interface{}) edpt.DdosZoneTemplateSipSrcSipRequestRateLimitMethodReferCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipSrcSipRequestRateLimitMethodReferCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Refer = in["refer"].(int)
		ret.SrcSipReferRate = in["src_sip_refer_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateSipSrcSipRequestRateLimitMethodMessageCfg(d []interface{}) edpt.DdosZoneTemplateSipSrcSipRequestRateLimitMethodMessageCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipSrcSipRequestRateLimitMethodMessageCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Message = in["message"].(int)
		ret.SrcSipMessageRate = in["src_sip_message_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateSipSrcSipRequestRateLimitMethodUpdateCfg(d []interface{}) edpt.DdosZoneTemplateSipSrcSipRequestRateLimitMethodUpdateCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateSipSrcSipRequestRateLimitMethodUpdateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Update = in["update"].(int)
		ret.SrcSipUpdateRate = in["src_sip_update_rate"].(int)
	}
	return ret
}

func dataToEndpointDdosZoneTemplateSip(d *schema.ResourceData) edpt.DdosZoneTemplateSip {
	var ret edpt.DdosZoneTemplateSip
	ret.Inst.Dst = getObjectDdosZoneTemplateSipDst(d.Get("dst").([]interface{}))
	ret.Inst.FilterHeaderList = getSliceDdosZoneTemplateSipFilterHeaderList(d.Get("filter_header_list").([]interface{}))
	ret.Inst.IdleTimeout = getObjectDdosZoneTemplateSipIdleTimeout(d.Get("idle_timeout").([]interface{}))
	ret.Inst.MalformedSip = getObjectDdosZoneTemplateSipMalformedSip316(d.Get("malformed_sip").([]interface{}))
	ret.Inst.MultiPuThresholdDistribution = getObjectDdosZoneTemplateSipMultiPuThresholdDistribution(d.Get("multi_pu_threshold_distribution").([]interface{}))
	ret.Inst.SipTmplName = d.Get("sip_tmpl_name").(string)
	ret.Inst.Src = getObjectDdosZoneTemplateSipSrc(d.Get("src").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
