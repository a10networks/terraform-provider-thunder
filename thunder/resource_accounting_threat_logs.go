package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAccountingThreatLogs() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_accounting_threat_logs`: Check for threat-logs from a specific user\n\n__PLACEHOLDER__",
		CreateContext: resourceAccountingThreatLogsCreate,
		UpdateContext: resourceAccountingThreatLogsUpdate,
		ReadContext:   resourceAccountingThreatLogsRead,
		DeleteContext: resourceAccountingThreatLogsDelete,

		Schema: map[string]*schema.Schema{
			"check": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Check if the system is under threat from a specific user",
			},
			"days": {
				Type: schema.TypeInt, Optional: true, Default: 30, Description: "Set min-days to go back",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAccountingThreatLogsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAccountingThreatLogsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAccountingThreatLogs(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAccountingThreatLogsRead(ctx, d, meta)
	}
	return diags
}

func resourceAccountingThreatLogsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAccountingThreatLogsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAccountingThreatLogs(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAccountingThreatLogsRead(ctx, d, meta)
	}
	return diags
}
func resourceAccountingThreatLogsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAccountingThreatLogsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAccountingThreatLogs(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAccountingThreatLogsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAccountingThreatLogsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAccountingThreatLogs(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAccountingThreatLogs(d *schema.ResourceData) edpt.AccountingThreatLogs {
	var ret edpt.AccountingThreatLogs
	ret.Inst.Check = d.Get("check").(int)
	ret.Inst.Days = d.Get("days").(int)
	//omit uuid
	return ret
}
