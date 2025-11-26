package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemSyslogTimeMsec() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_syslog_time_msec`: Enable syslog timestamp in millisecond\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemSyslogTimeMsecCreate,
		UpdateContext: resourceSystemSyslogTimeMsecUpdate,
		ReadContext:   resourceSystemSyslogTimeMsecRead,
		DeleteContext: resourceSystemSyslogTimeMsecDelete,

		Schema: map[string]*schema.Schema{
			"enable_flag": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
		},
	}
}
func resourceSystemSyslogTimeMsecCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSyslogTimeMsecCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSyslogTimeMsec(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSyslogTimeMsecRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemSyslogTimeMsecUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSyslogTimeMsecUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSyslogTimeMsec(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSyslogTimeMsecRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemSyslogTimeMsecDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSyslogTimeMsecDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSyslogTimeMsec(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemSyslogTimeMsecRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSyslogTimeMsecRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSyslogTimeMsec(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemSyslogTimeMsec(d *schema.ResourceData) edpt.SystemSyslogTimeMsec {
	var ret edpt.SystemSyslogTimeMsec
	ret.Inst.EnableFlag = d.Get("enable_flag").(int)
	return ret
}
