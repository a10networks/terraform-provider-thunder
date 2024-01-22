package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbCommonCertPinningCandidateListFeedbackOptIn() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_common_cert_pinning_candidate_list_feedback_opt_in`: Opt-in to improve A10’s certificate-pinning sites list\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbCommonCertPinningCandidateListFeedbackOptInCreate,
		UpdateContext: resourceSlbCommonCertPinningCandidateListFeedbackOptInUpdate,
		ReadContext:   resourceSlbCommonCertPinningCandidateListFeedbackOptInRead,
		DeleteContext: resourceSlbCommonCertPinningCandidateListFeedbackOptInDelete,

		Schema: map[string]*schema.Schema{
			"daily": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Every day",
			},
			"day_time": {
				Type: schema.TypeString, Optional: true, Description: "Time of day to update (hh:mm) in 24 hour local time",
			},
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the feedback function",
			},
			"schedule": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "schedule the uploading time, default is daily 00:00",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port to connect",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"week_day": {
				Type: schema.TypeString, Optional: true, Description: "'Monday': Monday; 'Tuesday': Tuesday; 'Wednesday': Wednesday; 'Thursday': Thursday; 'Friday': Friday; 'Saturday': Saturday; 'Sunday': Sunday;",
			},
			"week_time": {
				Type: schema.TypeString, Optional: true, Description: "Time of day to update (hh:mm) in 24 hour local time",
			},
			"weekly": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Every week",
			},
		},
	}
}
func resourceSlbCommonCertPinningCandidateListFeedbackOptInCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonCertPinningCandidateListFeedbackOptInCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonCertPinningCandidateListFeedbackOptIn(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbCommonCertPinningCandidateListFeedbackOptInRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbCommonCertPinningCandidateListFeedbackOptInUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonCertPinningCandidateListFeedbackOptInUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonCertPinningCandidateListFeedbackOptIn(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbCommonCertPinningCandidateListFeedbackOptInRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbCommonCertPinningCandidateListFeedbackOptInDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonCertPinningCandidateListFeedbackOptInDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonCertPinningCandidateListFeedbackOptIn(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbCommonCertPinningCandidateListFeedbackOptInRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonCertPinningCandidateListFeedbackOptInRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonCertPinningCandidateListFeedbackOptIn(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbCommonCertPinningCandidateListFeedbackOptIn(d *schema.ResourceData) edpt.SlbCommonCertPinningCandidateListFeedbackOptIn {
	var ret edpt.SlbCommonCertPinningCandidateListFeedbackOptIn
	ret.Inst.Daily = d.Get("daily").(int)
	ret.Inst.DayTime = d.Get("day_time").(string)
	ret.Inst.Enable = d.Get("enable").(int)
	ret.Inst.Schedule = d.Get("schedule").(int)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	//omit uuid
	ret.Inst.WeekDay = d.Get("week_day").(string)
	ret.Inst.WeekTime = d.Get("week_time").(string)
	ret.Inst.Weekly = d.Get("weekly").(int)
	return ret
}
