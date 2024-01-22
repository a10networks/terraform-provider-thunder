package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateGtpLoggingPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_template_gtp_logging_policy`: Configure GTP Logging Policy\n\n__PLACEHOLDER__",
		CreateContext: resourceTemplateGtpLoggingPolicyCreate,
		UpdateContext: resourceTemplateGtpLoggingPolicyUpdate,
		ReadContext:   resourceTemplateGtpLoggingPolicyRead,
		DeleteContext: resourceTemplateGtpLoggingPolicyDelete,

		Schema: map[string]*schema.Schema{
			"log": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mandatory_ie_check": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Logging Packet Drop due to Missing Mandatory Information Element",
						},
						"out_of_state_ie_check": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Logging Packet Drop due to Out of State Information Elements",
						},
						"out_of_order_ie_check": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Logging Packet Drop due to Out of Order Information Elements for GTPv1-C",
						},
						"invalid_teid_check": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Logging Packet Drop due to Invalid Tunnel Endpoint Identifier",
						},
						"reserved_ie_check": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Logging Packet Drop due to Presence of Reserved Information Element",
						},
						"message_filtering": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Logging Packet Drop due to Message Filtering",
						},
						"apn_imsi_filtering": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Logging Packet Drop due to APN IMSI filtering",
						},
						"rat_type_filtering": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Logging Packet Drop due to RAT type filtering",
						},
						"msisdn_filtering": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Logging Packet Drop due to MSISDN Filtering",
						},
						"crosslayer_correlation": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Logging Packet Drop due to mismatch in IP address and GTP FTEID/GSN address",
						},
						"anti_spoofing_check": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Logging Packet Drop due to End User IP Address Spoofing",
						},
						"msisdn_imsi_correlation": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Logging Packet Drop due to mismatch in Country Code and Mobile Country Code",
						},
						"max_message_length_check": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Logging Packet Drop due to Max Message Length Filtering",
						},
						"gtp_in_gtp_filtering": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Logging Packet Drop due to GTP in GTP filtering",
						},
						"sequence_num_correlation": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Logging Packet Drop due to mismatch in Sequence number between GTP request and response",
						},
						"invalid_header_check": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Logging Packet Drop due to Invalid Header checks",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of the GTP Logging Policy",
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
func resourceTemplateGtpLoggingPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpLoggingPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpLoggingPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateGtpLoggingPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceTemplateGtpLoggingPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpLoggingPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpLoggingPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateGtpLoggingPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceTemplateGtpLoggingPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpLoggingPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpLoggingPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTemplateGtpLoggingPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpLoggingPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpLoggingPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectTemplateGtpLoggingPolicyLog(d []interface{}) edpt.TemplateGtpLoggingPolicyLog {

	count1 := len(d)
	var ret edpt.TemplateGtpLoggingPolicyLog
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MandatoryIeCheck = in["mandatory_ie_check"].(int)
		ret.OutOfStateIeCheck = in["out_of_state_ie_check"].(int)
		ret.OutOfOrderIeCheck = in["out_of_order_ie_check"].(int)
		ret.InvalidTeidCheck = in["invalid_teid_check"].(int)
		ret.ReservedIeCheck = in["reserved_ie_check"].(int)
		ret.MessageFiltering = in["message_filtering"].(int)
		ret.ApnImsiFiltering = in["apn_imsi_filtering"].(int)
		ret.RatTypeFiltering = in["rat_type_filtering"].(int)
		ret.MsisdnFiltering = in["msisdn_filtering"].(int)
		ret.CrosslayerCorrelation = in["crosslayer_correlation"].(int)
		ret.AntiSpoofingCheck = in["anti_spoofing_check"].(int)
		ret.MsisdnImsiCorrelation = in["msisdn_imsi_correlation"].(int)
		ret.MaxMessageLengthCheck = in["max_message_length_check"].(int)
		ret.GtpInGtpFiltering = in["gtp_in_gtp_filtering"].(int)
		ret.SequenceNumCorrelation = in["sequence_num_correlation"].(int)
		ret.InvalidHeaderCheck = in["invalid_header_check"].(int)
	}
	return ret
}

func dataToEndpointTemplateGtpLoggingPolicy(d *schema.ResourceData) edpt.TemplateGtpLoggingPolicy {
	var ret edpt.TemplateGtpLoggingPolicy
	ret.Inst.Log = getObjectTemplateGtpLoggingPolicyLog(d.Get("log").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
