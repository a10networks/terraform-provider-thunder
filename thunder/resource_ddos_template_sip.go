package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateSip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_sip`: SIP template Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateSipCreate,
		UpdateContext: resourceDdosTemplateSipUpdate,
		ReadContext:   resourceDdosTemplateSipRead,
		DeleteContext: resourceDdosTemplateSipDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for sip connection; 'reset': Send RST for sip-tcp connection;",
			},
			"dst": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sip_request_rate_limit": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"method": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"invite_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dst_sip_invite_cfg_flag": {
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
															"dst_sip_register_cfg_flag": {
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
															"dst_sip_options_cfg_flag": {
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
															"dst_sip_bye_cfg_flag": {
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
															"dst_sip_subscribe_cfg_flag": {
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
															"dst_sip_notify_cfg_flag": {
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
															"dst_sip_refer_cfg_flag": {
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
															"dst_sip_message_cfg_flag": {
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
															"dst_sip_update_cfg_flag": {
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
						"sip_filter_header_seq": {
							Type: schema.TypeInt, Required: true, Description: "Sequence number",
						},
						"sip_filter_header_regex": {
							Type: schema.TypeString, Optional: true, Description: "Regex Expression",
						},
						"sip_filter_header_unmatched": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "action taken when it does not match",
						},
						"sip_filter_header_blacklist": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Also blacklist the source when action is taken",
						},
						"sip_filter_header_whitelist": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Whitelist the source after filter passes, packets are dropped until then",
						},
						"sip_filter_header_count_only": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Take no action and continue processing the next filter",
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
				Type: schema.TypeInt, Optional: true, Description: "Set the the idle timeout value for sip-tcp connections",
			},
			"ignore_zero_payload": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Don't reset idle timer on packets with zero payload length from clients",
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
									"method": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"invite_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"src_sip_invite_cfg_flag": {
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
															"src_sip_register_cfg_flag": {
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
															"src_sip_options_cfg_flag": {
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
															"src_sip_bye_cfg_flag": {
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
															"src_sip_subscribe_cfg_flag": {
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
															"src_sip_notify_cfg_flag": {
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
															"src_sip_refer_cfg_flag": {
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
															"src_sip_message_cfg_flag": {
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
															"src_sip_update_cfg_flag": {
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
func resourceDdosTemplateSipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateSipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateSip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateSipRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateSipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateSipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateSip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateSipRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateSipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateSipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateSip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateSipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateSipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateSip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosTemplateSipDst(d []interface{}) edpt.DdosTemplateSipDst {

	count1 := len(d)
	var ret edpt.DdosTemplateSipDst
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SipRequestRateLimit = getObjectDdosTemplateSipDstSipRequestRateLimit(in["sip_request_rate_limit"].([]interface{}))
	}
	return ret
}

func getObjectDdosTemplateSipDstSipRequestRateLimit(d []interface{}) edpt.DdosTemplateSipDstSipRequestRateLimit {

	count1 := len(d)
	var ret edpt.DdosTemplateSipDstSipRequestRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Method = getObjectDdosTemplateSipDstSipRequestRateLimitMethod(in["method"].([]interface{}))
	}
	return ret
}

func getObjectDdosTemplateSipDstSipRequestRateLimitMethod(d []interface{}) edpt.DdosTemplateSipDstSipRequestRateLimitMethod {

	count1 := len(d)
	var ret edpt.DdosTemplateSipDstSipRequestRateLimitMethod
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.InviteCfg = getObjectDdosTemplateSipDstSipRequestRateLimitMethodInviteCfg(in["invite_cfg"].([]interface{}))
		ret.RegisterCfg = getObjectDdosTemplateSipDstSipRequestRateLimitMethodRegisterCfg(in["register_cfg"].([]interface{}))
		ret.OptionsCfg = getObjectDdosTemplateSipDstSipRequestRateLimitMethodOptionsCfg(in["options_cfg"].([]interface{}))
		ret.ByeCfg = getObjectDdosTemplateSipDstSipRequestRateLimitMethodByeCfg(in["bye_cfg"].([]interface{}))
		ret.SubscribeCfg = getObjectDdosTemplateSipDstSipRequestRateLimitMethodSubscribeCfg(in["subscribe_cfg"].([]interface{}))
		ret.NotifyCfg = getObjectDdosTemplateSipDstSipRequestRateLimitMethodNotifyCfg(in["notify_cfg"].([]interface{}))
		ret.ReferCfg = getObjectDdosTemplateSipDstSipRequestRateLimitMethodReferCfg(in["refer_cfg"].([]interface{}))
		ret.MessageCfg = getObjectDdosTemplateSipDstSipRequestRateLimitMethodMessageCfg(in["message_cfg"].([]interface{}))
		ret.UpdateCfg = getObjectDdosTemplateSipDstSipRequestRateLimitMethodUpdateCfg(in["update_cfg"].([]interface{}))
	}
	return ret
}

func getObjectDdosTemplateSipDstSipRequestRateLimitMethodInviteCfg(d []interface{}) edpt.DdosTemplateSipDstSipRequestRateLimitMethodInviteCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateSipDstSipRequestRateLimitMethodInviteCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DstSipInviteCfgFlag = in["dst_sip_invite_cfg_flag"].(int)
		ret.DstSipInviteRate = in["dst_sip_invite_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateSipDstSipRequestRateLimitMethodRegisterCfg(d []interface{}) edpt.DdosTemplateSipDstSipRequestRateLimitMethodRegisterCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateSipDstSipRequestRateLimitMethodRegisterCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DstSipRegisterCfgFlag = in["dst_sip_register_cfg_flag"].(int)
		ret.DstSipRegisterRate = in["dst_sip_register_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateSipDstSipRequestRateLimitMethodOptionsCfg(d []interface{}) edpt.DdosTemplateSipDstSipRequestRateLimitMethodOptionsCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateSipDstSipRequestRateLimitMethodOptionsCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DstSipOptionsCfgFlag = in["dst_sip_options_cfg_flag"].(int)
		ret.DstSipOptionsRate = in["dst_sip_options_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateSipDstSipRequestRateLimitMethodByeCfg(d []interface{}) edpt.DdosTemplateSipDstSipRequestRateLimitMethodByeCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateSipDstSipRequestRateLimitMethodByeCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DstSipByeCfgFlag = in["dst_sip_bye_cfg_flag"].(int)
		ret.DstSipByeRate = in["dst_sip_bye_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateSipDstSipRequestRateLimitMethodSubscribeCfg(d []interface{}) edpt.DdosTemplateSipDstSipRequestRateLimitMethodSubscribeCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateSipDstSipRequestRateLimitMethodSubscribeCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DstSipSubscribeCfgFlag = in["dst_sip_subscribe_cfg_flag"].(int)
		ret.DstSipSubscribeRate = in["dst_sip_subscribe_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateSipDstSipRequestRateLimitMethodNotifyCfg(d []interface{}) edpt.DdosTemplateSipDstSipRequestRateLimitMethodNotifyCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateSipDstSipRequestRateLimitMethodNotifyCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DstSipNotifyCfgFlag = in["dst_sip_notify_cfg_flag"].(int)
		ret.DstSipNotifyRate = in["dst_sip_notify_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateSipDstSipRequestRateLimitMethodReferCfg(d []interface{}) edpt.DdosTemplateSipDstSipRequestRateLimitMethodReferCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateSipDstSipRequestRateLimitMethodReferCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DstSipReferCfgFlag = in["dst_sip_refer_cfg_flag"].(int)
		ret.DstSipReferRate = in["dst_sip_refer_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateSipDstSipRequestRateLimitMethodMessageCfg(d []interface{}) edpt.DdosTemplateSipDstSipRequestRateLimitMethodMessageCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateSipDstSipRequestRateLimitMethodMessageCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DstSipMessageCfgFlag = in["dst_sip_message_cfg_flag"].(int)
		ret.DstSipMessageRate = in["dst_sip_message_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateSipDstSipRequestRateLimitMethodUpdateCfg(d []interface{}) edpt.DdosTemplateSipDstSipRequestRateLimitMethodUpdateCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateSipDstSipRequestRateLimitMethodUpdateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DstSipUpdateCfgFlag = in["dst_sip_update_cfg_flag"].(int)
		ret.DstSipUpdateRate = in["dst_sip_update_rate"].(int)
	}
	return ret
}

func getSliceDdosTemplateSipFilterHeaderList(d []interface{}) []edpt.DdosTemplateSipFilterHeaderList {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateSipFilterHeaderList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateSipFilterHeaderList
		oi.SipFilterHeaderSeq = in["sip_filter_header_seq"].(int)
		oi.SipFilterHeaderRegex = in["sip_filter_header_regex"].(string)
		oi.SipFilterHeaderUnmatched = in["sip_filter_header_unmatched"].(int)
		oi.SipFilterHeaderBlacklist = in["sip_filter_header_blacklist"].(int)
		oi.SipFilterHeaderWhitelist = in["sip_filter_header_whitelist"].(int)
		oi.SipFilterHeaderCountOnly = in["sip_filter_header_count_only"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosTemplateSipMalformedSip298(d []interface{}) edpt.DdosTemplateSipMalformedSip298 {

	count1 := len(d)
	var ret edpt.DdosTemplateSipMalformedSip298
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MalformedSipCheck = in["malformed_sip_check"].(string)
		ret.MalformedSipMaxLineSize = in["malformed_sip_max_line_size"].(int)
		ret.MalformedSipMaxUriLength = in["malformed_sip_max_uri_length"].(int)
		ret.MalformedSipMaxHeaderNameLength = in["malformed_sip_max_header_name_length"].(int)
		ret.MalformedSipMaxHeaderValueLength = in["malformed_sip_max_header_value_length"].(int)
		ret.MalformedSipCallIdMaxLength = in["malformed_sip_call_id_max_length"].(int)
		ret.MalformedSipSdpMaxLength = in["malformed_sip_sdp_max_length"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosTemplateSipMultiPuThresholdDistribution(d []interface{}) edpt.DdosTemplateSipMultiPuThresholdDistribution {

	count1 := len(d)
	var ret edpt.DdosTemplateSipMultiPuThresholdDistribution
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MultiPuThresholdDistributionValue = in["multi_pu_threshold_distribution_value"].(int)
		ret.MultiPuThresholdDistributionDisable = in["multi_pu_threshold_distribution_disable"].(string)
	}
	return ret
}

func getObjectDdosTemplateSipSrc(d []interface{}) edpt.DdosTemplateSipSrc {

	count1 := len(d)
	var ret edpt.DdosTemplateSipSrc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SipRequestRateLimit = getObjectDdosTemplateSipSrcSipRequestRateLimit(in["sip_request_rate_limit"].([]interface{}))
	}
	return ret
}

func getObjectDdosTemplateSipSrcSipRequestRateLimit(d []interface{}) edpt.DdosTemplateSipSrcSipRequestRateLimit {

	count1 := len(d)
	var ret edpt.DdosTemplateSipSrcSipRequestRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Method = getObjectDdosTemplateSipSrcSipRequestRateLimitMethod(in["method"].([]interface{}))
	}
	return ret
}

func getObjectDdosTemplateSipSrcSipRequestRateLimitMethod(d []interface{}) edpt.DdosTemplateSipSrcSipRequestRateLimitMethod {

	count1 := len(d)
	var ret edpt.DdosTemplateSipSrcSipRequestRateLimitMethod
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.InviteCfg = getObjectDdosTemplateSipSrcSipRequestRateLimitMethodInviteCfg(in["invite_cfg"].([]interface{}))
		ret.RegisterCfg = getObjectDdosTemplateSipSrcSipRequestRateLimitMethodRegisterCfg(in["register_cfg"].([]interface{}))
		ret.OptionsCfg = getObjectDdosTemplateSipSrcSipRequestRateLimitMethodOptionsCfg(in["options_cfg"].([]interface{}))
		ret.ByeCfg = getObjectDdosTemplateSipSrcSipRequestRateLimitMethodByeCfg(in["bye_cfg"].([]interface{}))
		ret.SubscribeCfg = getObjectDdosTemplateSipSrcSipRequestRateLimitMethodSubscribeCfg(in["subscribe_cfg"].([]interface{}))
		ret.NotifyCfg = getObjectDdosTemplateSipSrcSipRequestRateLimitMethodNotifyCfg(in["notify_cfg"].([]interface{}))
		ret.ReferCfg = getObjectDdosTemplateSipSrcSipRequestRateLimitMethodReferCfg(in["refer_cfg"].([]interface{}))
		ret.MessageCfg = getObjectDdosTemplateSipSrcSipRequestRateLimitMethodMessageCfg(in["message_cfg"].([]interface{}))
		ret.UpdateCfg = getObjectDdosTemplateSipSrcSipRequestRateLimitMethodUpdateCfg(in["update_cfg"].([]interface{}))
	}
	return ret
}

func getObjectDdosTemplateSipSrcSipRequestRateLimitMethodInviteCfg(d []interface{}) edpt.DdosTemplateSipSrcSipRequestRateLimitMethodInviteCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateSipSrcSipRequestRateLimitMethodInviteCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcSipInviteCfgFlag = in["src_sip_invite_cfg_flag"].(int)
		ret.SrcSipInviteRate = in["src_sip_invite_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateSipSrcSipRequestRateLimitMethodRegisterCfg(d []interface{}) edpt.DdosTemplateSipSrcSipRequestRateLimitMethodRegisterCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateSipSrcSipRequestRateLimitMethodRegisterCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcSipRegisterCfgFlag = in["src_sip_register_cfg_flag"].(int)
		ret.SrcSipRegisterRate = in["src_sip_register_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateSipSrcSipRequestRateLimitMethodOptionsCfg(d []interface{}) edpt.DdosTemplateSipSrcSipRequestRateLimitMethodOptionsCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateSipSrcSipRequestRateLimitMethodOptionsCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcSipOptionsCfgFlag = in["src_sip_options_cfg_flag"].(int)
		ret.SrcSipOptionsRate = in["src_sip_options_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateSipSrcSipRequestRateLimitMethodByeCfg(d []interface{}) edpt.DdosTemplateSipSrcSipRequestRateLimitMethodByeCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateSipSrcSipRequestRateLimitMethodByeCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcSipByeCfgFlag = in["src_sip_bye_cfg_flag"].(int)
		ret.SrcSipByeRate = in["src_sip_bye_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateSipSrcSipRequestRateLimitMethodSubscribeCfg(d []interface{}) edpt.DdosTemplateSipSrcSipRequestRateLimitMethodSubscribeCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateSipSrcSipRequestRateLimitMethodSubscribeCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcSipSubscribeCfgFlag = in["src_sip_subscribe_cfg_flag"].(int)
		ret.SrcSipSubscribeRate = in["src_sip_subscribe_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateSipSrcSipRequestRateLimitMethodNotifyCfg(d []interface{}) edpt.DdosTemplateSipSrcSipRequestRateLimitMethodNotifyCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateSipSrcSipRequestRateLimitMethodNotifyCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcSipNotifyCfgFlag = in["src_sip_notify_cfg_flag"].(int)
		ret.SrcSipNotifyRate = in["src_sip_notify_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateSipSrcSipRequestRateLimitMethodReferCfg(d []interface{}) edpt.DdosTemplateSipSrcSipRequestRateLimitMethodReferCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateSipSrcSipRequestRateLimitMethodReferCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcSipReferCfgFlag = in["src_sip_refer_cfg_flag"].(int)
		ret.SrcSipReferRate = in["src_sip_refer_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateSipSrcSipRequestRateLimitMethodMessageCfg(d []interface{}) edpt.DdosTemplateSipSrcSipRequestRateLimitMethodMessageCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateSipSrcSipRequestRateLimitMethodMessageCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcSipMessageCfgFlag = in["src_sip_message_cfg_flag"].(int)
		ret.SrcSipMessageRate = in["src_sip_message_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateSipSrcSipRequestRateLimitMethodUpdateCfg(d []interface{}) edpt.DdosTemplateSipSrcSipRequestRateLimitMethodUpdateCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateSipSrcSipRequestRateLimitMethodUpdateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcSipUpdateCfgFlag = in["src_sip_update_cfg_flag"].(int)
		ret.SrcSipUpdateRate = in["src_sip_update_rate"].(int)
	}
	return ret
}

func dataToEndpointDdosTemplateSip(d *schema.ResourceData) edpt.DdosTemplateSip {
	var ret edpt.DdosTemplateSip
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Dst = getObjectDdosTemplateSipDst(d.Get("dst").([]interface{}))
	ret.Inst.FilterHeaderList = getSliceDdosTemplateSipFilterHeaderList(d.Get("filter_header_list").([]interface{}))
	ret.Inst.IdleTimeout = d.Get("idle_timeout").(int)
	ret.Inst.IgnoreZeroPayload = d.Get("ignore_zero_payload").(int)
	ret.Inst.MalformedSip = getObjectDdosTemplateSipMalformedSip298(d.Get("malformed_sip").([]interface{}))
	ret.Inst.MultiPuThresholdDistribution = getObjectDdosTemplateSipMultiPuThresholdDistribution(d.Get("multi_pu_threshold_distribution").([]interface{}))
	ret.Inst.SipTmplName = d.Get("sip_tmpl_name").(string)
	ret.Inst.Src = getObjectDdosTemplateSipSrc(d.Get("src").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
