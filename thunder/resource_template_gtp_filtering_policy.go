package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateGtpFilteringPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_template_gtp_filtering_policy`: Configure GTP Filtering Policy\n\n__PLACEHOLDER__",
		CreateContext: resourceTemplateGtpFilteringPolicyCreate,
		UpdateContext: resourceTemplateGtpFilteringPolicyUpdate,
		ReadContext:   resourceTemplateGtpFilteringPolicyRead,
		DeleteContext: resourceTemplateGtpFilteringPolicyDelete,

		Schema: map[string]*schema.Schema{
			"apn_imsi_filt_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'monitor': Forward packet; 'drop': drop packet (default);",
			},
			"apn_imsi_filtering": {
				Type: schema.TypeString, Optional: true, Description: "Specify the APN IMSI Inspection Policy",
			},
			"gtp_in_gtp_filt_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'monitor': Forward packet; 'drop': drop packet (default);",
			},
			"gtp_in_gtp_filtering": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable GTP-in-GTP Filtering; 'disable': Disable GTP-in-GTP Filtering;",
			},
			"message_filt_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'monitor': Forward packet; 'drop': drop packet (default);",
			},
			"message_filtering_policy_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify the Message Filtering Policy",
			},
			"msisdn_filt_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'monitor': Forward packet; 'drop': drop packet (default);",
			},
			"msisdn_filtering": {
				Type: schema.TypeString, Optional: true, Description: "Specify the MSISDN Inspection Policy",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of the GTP Filtering Policy",
			},
			"rat_type_filtering": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rat_type_value": {
							Type: schema.TypeString, Optional: true, Description: "'utran': RAT Type for 3G Networks; 'geran': RAT Type for 2G Networks; 'wlan': RAT Type for WLAN Networks; 'gan': RAT Type for GAN Networks; 'hspa-evolution': RAT Type for HSPA Networks; 'eutran': RAT Type for 4G Networks;",
						},
						"rat_type_filt_action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'monitor': Forward packet; 'drop': drop packet (default);",
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
func resourceTemplateGtpFilteringPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpFilteringPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpFilteringPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateGtpFilteringPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceTemplateGtpFilteringPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpFilteringPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpFilteringPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateGtpFilteringPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceTemplateGtpFilteringPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpFilteringPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpFilteringPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTemplateGtpFilteringPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpFilteringPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpFilteringPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceTemplateGtpFilteringPolicyRatTypeFiltering(d []interface{}) []edpt.TemplateGtpFilteringPolicyRatTypeFiltering {

	count1 := len(d)
	ret := make([]edpt.TemplateGtpFilteringPolicyRatTypeFiltering, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.TemplateGtpFilteringPolicyRatTypeFiltering
		oi.RatTypeValue = in["rat_type_value"].(string)
		oi.RatTypeFiltAction = in["rat_type_filt_action"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointTemplateGtpFilteringPolicy(d *schema.ResourceData) edpt.TemplateGtpFilteringPolicy {
	var ret edpt.TemplateGtpFilteringPolicy
	ret.Inst.ApnImsiFiltAction = d.Get("apn_imsi_filt_action").(string)
	ret.Inst.ApnImsiFiltering = d.Get("apn_imsi_filtering").(string)
	ret.Inst.GtpInGtpFiltAction = d.Get("gtp_in_gtp_filt_action").(string)
	ret.Inst.GtpInGtpFiltering = d.Get("gtp_in_gtp_filtering").(string)
	ret.Inst.MessageFiltAction = d.Get("message_filt_action").(string)
	ret.Inst.MessageFilteringPolicyName = d.Get("message_filtering_policy_name").(string)
	ret.Inst.MsisdnFiltAction = d.Get("msisdn_filt_action").(string)
	ret.Inst.MsisdnFiltering = d.Get("msisdn_filtering").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.RatTypeFiltering = getSliceTemplateGtpFilteringPolicyRatTypeFiltering(d.Get("rat_type_filtering").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
