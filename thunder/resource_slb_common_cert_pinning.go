package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbCommonCertPinning() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_common_cert_pinning`: Configure the cert-pinning related settings\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbCommonCertPinningCreate,
		UpdateContext: resourceSlbCommonCertPinningUpdate,
		ReadContext:   resourceSlbCommonCertPinningRead,
		DeleteContext: resourceSlbCommonCertPinningDelete,

		Schema: map[string]*schema.Schema{
			"candidate_list_feedback_opt_in": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the feedback function",
						},
						"schedule": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "schedule the uploading time, default is daily 00:00",
						},
						"weekly": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Every week",
						},
						"week_day": {
							Type: schema.TypeString, Optional: true, Description: "'Monday': Monday; 'Tuesday': Tuesday; 'Wednesday': Wednesday; 'Thursday': Thursday; 'Friday': Friday; 'Saturday': Saturday; 'Sunday': Sunday;",
						},
						"week_time": {
							Type: schema.TypeString, Optional: true, Description: "Time of day to update (hh:mm) in 24 hour local time",
						},
						"daily": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Every day",
						},
						"day_time": {
							Type: schema.TypeString, Optional: true, Description: "Time of day to update (hh:mm) in 24 hour local time",
						},
						"use_mgmt_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port to connect",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ttl": {
				Type: schema.TypeInt, Optional: true, Default: 144, Description: "The ttl of local cert pinning candidate list, multiple of 10 minutes, default is 144 (1440 minutes)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbCommonCertPinningCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonCertPinningCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonCertPinning(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbCommonCertPinningRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbCommonCertPinningUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonCertPinningUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonCertPinning(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbCommonCertPinningRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbCommonCertPinningDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonCertPinningDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonCertPinning(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbCommonCertPinningRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonCertPinningRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonCertPinning(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbCommonCertPinningCandidateListFeedbackOptIn1406(d []interface{}) edpt.SlbCommonCertPinningCandidateListFeedbackOptIn1406 {

	count1 := len(d)
	var ret edpt.SlbCommonCertPinningCandidateListFeedbackOptIn1406
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		ret.Schedule = in["schedule"].(int)
		ret.Weekly = in["weekly"].(int)
		ret.WeekDay = in["week_day"].(string)
		ret.WeekTime = in["week_time"].(string)
		ret.Daily = in["daily"].(int)
		ret.DayTime = in["day_time"].(string)
		ret.UseMgmtPort = in["use_mgmt_port"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointSlbCommonCertPinning(d *schema.ResourceData) edpt.SlbCommonCertPinning {
	var ret edpt.SlbCommonCertPinning
	ret.Inst.CandidateListFeedbackOptIn = getObjectSlbCommonCertPinningCandidateListFeedbackOptIn1406(d.Get("candidate_list_feedback_opt_in").([]interface{}))
	ret.Inst.Ttl = d.Get("ttl").(int)
	//omit uuid
	return ret
}
