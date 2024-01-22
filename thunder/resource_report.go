package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceReport() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_report`: Report configurations\n\n__PLACEHOLDER__",
		CreateContext: resourceReportCreate,
		UpdateContext: resourceReportUpdate,
		ReadContext:   resourceReportRead,
		DeleteContext: resourceReportDelete,

		Schema: map[string]*schema.Schema{
			"debug": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Report module's normal logs",
						},
						"sflow": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"parser": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logs by parser",
									},
									"stats_oid": {
										Type: schema.TypeInt, Optional: true, Description: "Specify stats-oid to dump raw packets, 0 to disable",
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
func resourceReportCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceReportCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointReport(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceReportRead(ctx, d, meta)
	}
	return diags
}

func resourceReportUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceReportUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointReport(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceReportRead(ctx, d, meta)
	}
	return diags
}
func resourceReportDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceReportDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointReport(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceReportRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceReportRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointReport(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectReportDebug1095(d []interface{}) edpt.ReportDebug1095 {

	count1 := len(d)
	var ret edpt.ReportDebug1095
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Log = in["log"].(int)
		ret.Sflow = getObjectReportDebugSflow1096(in["sflow"].([]interface{}))
	}
	return ret
}

func getObjectReportDebugSflow1096(d []interface{}) edpt.ReportDebugSflow1096 {

	count1 := len(d)
	var ret edpt.ReportDebugSflow1096
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Parser = in["parser"].(int)
		ret.StatsOid = in["stats_oid"].(int)
	}
	return ret
}

func dataToEndpointReport(d *schema.ResourceData) edpt.Report {
	var ret edpt.Report
	ret.Inst.Debug = getObjectReportDebug1095(d.Get("debug").([]interface{}))
	return ret
}
