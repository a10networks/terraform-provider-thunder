package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileAflexOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_aflex_oper`: Operational Status for the object aflex\n\n__PLACEHOLDER__",
		CreateContext: resourceFileAflexOperCreate,
		UpdateContext: resourceFileAflexOperUpdate,
		ReadContext:   resourceFileAflexOperRead,
		DeleteContext: resourceFileAflexOperDelete,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"file_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"file": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"syntax": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"vport": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"vport_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"vserver": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"port": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
									"events": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"event_type": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"total_executions": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"failures": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"aborts": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
									"waf_rule": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"waf_rule_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"waf_rule_total": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"waf_rule_failures": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"waf_rule_aborts": {
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
	}
}
func resourceFileAflexOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAflexOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAflexOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileAflexOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileAflexOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAflexOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAflexOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileAflexOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileAflexOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAflexOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAflexOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileAflexOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAflexOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAflexOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileAflexOperOper(d []interface{}) edpt.FileAflexOperOper {

	count1 := len(d)
	var ret edpt.FileAflexOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceFileAflexOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceFileAflexOperOperFileList(d []interface{}) []edpt.FileAflexOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.FileAflexOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileAflexOperOperFileList
		oi.File = in["file"].(string)
		oi.Syntax = in["syntax"].(string)
		oi.Vport = in["vport"].(string)
		oi.VportList = getSliceFileAflexOperOperFileListVportList(in["vport_list"].([]interface{}))
		oi.Events = getSliceFileAflexOperOperFileListEvents(in["events"].([]interface{}))
		oi.WafRule = getSliceFileAflexOperOperFileListWafRule(in["waf_rule"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFileAflexOperOperFileListVportList(d []interface{}) []edpt.FileAflexOperOperFileListVportList {

	count1 := len(d)
	ret := make([]edpt.FileAflexOperOperFileListVportList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileAflexOperOperFileListVportList
		oi.Vserver = in["vserver"].(string)
		oi.Port = in["port"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFileAflexOperOperFileListEvents(d []interface{}) []edpt.FileAflexOperOperFileListEvents {

	count1 := len(d)
	ret := make([]edpt.FileAflexOperOperFileListEvents, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileAflexOperOperFileListEvents
		oi.EventType = in["event_type"].(string)
		oi.TotalExecutions = in["total_executions"].(int)
		oi.Failures = in["failures"].(int)
		oi.Aborts = in["aborts"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFileAflexOperOperFileListWafRule(d []interface{}) []edpt.FileAflexOperOperFileListWafRule {

	count1 := len(d)
	ret := make([]edpt.FileAflexOperOperFileListWafRule, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileAflexOperOperFileListWafRule
		oi.WafRuleName = in["waf_rule_name"].(string)
		oi.WafRuleTotal = in["waf_rule_total"].(int)
		oi.WafRuleFailures = in["waf_rule_failures"].(int)
		oi.WafRuleAborts = in["waf_rule_aborts"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileAflexOper(d *schema.ResourceData) edpt.FileAflexOper {
	var ret edpt.FileAflexOper
	ret.Inst.Oper = getObjectFileAflexOperOper(d.Get("oper").([]interface{}))
	return ret
}
