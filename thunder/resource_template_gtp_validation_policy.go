package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateGtpValidationPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_template_gtp_validation_policy`: Configure GTP Validation Policy\n\n__PLACEHOLDER__",
		CreateContext: resourceTemplateGtpValidationPolicyCreate,
		UpdateContext: resourceTemplateGtpValidationPolicyUpdate,
		ReadContext:   resourceTemplateGtpValidationPolicyRead,
		DeleteContext: resourceTemplateGtpValidationPolicyDelete,

		Schema: map[string]*schema.Schema{
			"anomaly_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'monitor': Forward failed packet; 'drop': drop check failed packet(default);",
			},
			"anomaly_checks": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable Anomaly Checks for GTP Messages; 'disable': Disable Anomaly Checks for GTP Messages;",
			},
			"anti_spoofing_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'monitor': Forward failed packet; 'drop': drop check failed packet(default);",
			},
			"anti_spoofing_check": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Spoofed IP Address Check for GTP-U Messages; 'disable': Disable Spoofed IP Address Check for GTP Messages;",
			},
			"crosslayer_corr_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'monitor': Forward failed packet; 'drop': drop check failed packet(default);",
			},
			"crosslayer_correlation": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Cross Layer Correlation for GTP Messages; 'disable': Disable Cross Layer Correlation for GTP Messages;",
			},
			"mandatory_ie_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'monitor': Forward failed packet; 'drop': drop check failed packet(default);",
			},
			"mandatory_ie_check": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable Mandatory IE Check for GTP Messages; 'disable': Disable Mandatory IE Check for GTP Messages;",
			},
			"msisdn_imsi_corr_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'monitor': Forward failed packet; 'drop': drop check failed packet(default);",
			},
			"msisdn_imsi_correlation": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Country Code Correlation in MSISDN and IMSI for GTP Messages; 'disable': Disable Country Code Correlation in MSISDN and IMSI for GTP Messages;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of the GTP Validation Policy",
			},
			"out_of_order_ie_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'monitor': Forward failed packet; 'drop': drop check failed packet(default);",
			},
			"out_of_order_ie_check": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Out of Order IE Check for GTP Messages; 'disable': Disable Out of Order IE Check for GTP Messages;",
			},
			"out_of_state_ie_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'monitor': Forward failed packet; 'drop': drop check failed packet(default);",
			},
			"out_of_state_ie_check": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Out of State IE Check for GTP Messages; 'disable': Disable Out of State IE Check for GTP Messages;",
			},
			"reserved_ie_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'monitor': Forward failed packet; 'drop': drop check failed packet(default);",
			},
			"reserved_ie_check": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable Reserved IE Check for GTP Messages; 'disable': Disable Reserved IE Check for GTP Messages;",
			},
			"sequence_num_corr_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'monitor': Forward failed packet; 'drop': drop check failed packet(default);",
			},
			"sequence_num_correlation": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Sequence number correlation between GTP Request and Responses; 'disable': Disable Sequence number correlation between GTP Request and Responses;",
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
func resourceTemplateGtpValidationPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpValidationPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpValidationPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateGtpValidationPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceTemplateGtpValidationPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpValidationPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpValidationPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateGtpValidationPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceTemplateGtpValidationPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpValidationPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpValidationPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTemplateGtpValidationPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpValidationPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpValidationPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointTemplateGtpValidationPolicy(d *schema.ResourceData) edpt.TemplateGtpValidationPolicy {
	var ret edpt.TemplateGtpValidationPolicy
	ret.Inst.AnomalyAction = d.Get("anomaly_action").(string)
	ret.Inst.AnomalyChecks = d.Get("anomaly_checks").(string)
	ret.Inst.AntiSpoofingAction = d.Get("anti_spoofing_action").(string)
	ret.Inst.AntiSpoofingCheck = d.Get("anti_spoofing_check").(string)
	ret.Inst.CrosslayerCorrAction = d.Get("crosslayer_corr_action").(string)
	ret.Inst.CrosslayerCorrelation = d.Get("crosslayer_correlation").(string)
	ret.Inst.MandatoryIeAction = d.Get("mandatory_ie_action").(string)
	ret.Inst.MandatoryIeCheck = d.Get("mandatory_ie_check").(string)
	ret.Inst.MsisdnImsiCorrAction = d.Get("msisdn_imsi_corr_action").(string)
	ret.Inst.MsisdnImsiCorrelation = d.Get("msisdn_imsi_correlation").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.OutOfOrderIeAction = d.Get("out_of_order_ie_action").(string)
	ret.Inst.OutOfOrderIeCheck = d.Get("out_of_order_ie_check").(string)
	ret.Inst.OutOfStateIeAction = d.Get("out_of_state_ie_action").(string)
	ret.Inst.OutOfStateIeCheck = d.Get("out_of_state_ie_check").(string)
	ret.Inst.ReservedIeAction = d.Get("reserved_ie_action").(string)
	ret.Inst.ReservedIeCheck = d.Get("reserved_ie_check").(string)
	ret.Inst.SequenceNumCorrAction = d.Get("sequence_num_corr_action").(string)
	ret.Inst.SequenceNumCorrelation = d.Get("sequence_num_correlation").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
