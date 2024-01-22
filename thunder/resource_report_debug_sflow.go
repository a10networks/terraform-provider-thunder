package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceReportDebugSflow() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_report_debug_sflow`: Report sFlow parser options\n\n__PLACEHOLDER__",
		CreateContext: resourceReportDebugSflowCreate,
		UpdateContext: resourceReportDebugSflowUpdate,
		ReadContext:   resourceReportDebugSflowRead,
		DeleteContext: resourceReportDebugSflowDelete,

		Schema: map[string]*schema.Schema{
			"parser": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logs by parser",
			},
			"stats_oid": {
				Type: schema.TypeInt, Optional: true, Description: "Specify stats-oid to dump raw packets, 0 to disable",
			},
		},
	}
}
func resourceReportDebugSflowCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceReportDebugSflowCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointReportDebugSflow(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceReportDebugSflowRead(ctx, d, meta)
	}
	return diags
}

func resourceReportDebugSflowUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceReportDebugSflowUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointReportDebugSflow(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceReportDebugSflowRead(ctx, d, meta)
	}
	return diags
}
func resourceReportDebugSflowDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceReportDebugSflowDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointReportDebugSflow(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceReportDebugSflowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceReportDebugSflowRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointReportDebugSflow(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointReportDebugSflow(d *schema.ResourceData) edpt.ReportDebugSflow {
	var ret edpt.ReportDebugSflow
	ret.Inst.Parser = d.Get("parser").(int)
	ret.Inst.StatsOid = d.Get("stats_oid").(int)
	return ret
}
