package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceReportDebug() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_report_debug`: Report debugging options\n\n__PLACEHOLDER__",
		CreateContext: resourceReportDebugCreate,
		UpdateContext: resourceReportDebugUpdate,
		ReadContext:   resourceReportDebugRead,
		DeleteContext: resourceReportDebugDelete,

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
	}
}
func resourceReportDebugCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceReportDebugCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointReportDebug(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceReportDebugRead(ctx, d, meta)
	}
	return diags
}

func resourceReportDebugUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceReportDebugUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointReportDebug(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceReportDebugRead(ctx, d, meta)
	}
	return diags
}
func resourceReportDebugDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceReportDebugDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointReportDebug(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceReportDebugRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceReportDebugRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointReportDebug(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectReportDebugSflow1094(d []interface{}) edpt.ReportDebugSflow1094 {

	count1 := len(d)
	var ret edpt.ReportDebugSflow1094
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Parser = in["parser"].(int)
		ret.StatsOid = in["stats_oid"].(int)
	}
	return ret
}

func dataToEndpointReportDebug(d *schema.ResourceData) edpt.ReportDebug {
	var ret edpt.ReportDebug
	ret.Inst.Log = d.Get("log").(int)
	ret.Inst.Sflow = getObjectReportDebugSflow1094(d.Get("sflow").([]interface{}))
	return ret
}
