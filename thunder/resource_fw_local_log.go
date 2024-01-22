package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwLocalLog() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_local_log`: Enable local-log for Application Firewall\n\n__PLACEHOLDER__",
		CreateContext: resourceFwLocalLogCreate,
		UpdateContext: resourceFwLocalLogUpdate,
		ReadContext:   resourceFwLocalLogRead,
		DeleteContext: resourceFwLocalLogDelete,

		Schema: map[string]*schema.Schema{
			"local_logging": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable local logging",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFwLocalLogCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwLocalLogCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwLocalLog(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwLocalLogRead(ctx, d, meta)
	}
	return diags
}

func resourceFwLocalLogUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwLocalLogUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwLocalLog(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwLocalLogRead(ctx, d, meta)
	}
	return diags
}
func resourceFwLocalLogDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwLocalLogDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwLocalLog(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwLocalLogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwLocalLogRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwLocalLog(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFwLocalLog(d *schema.ResourceData) edpt.FwLocalLog {
	var ret edpt.FwLocalLog
	ret.Inst.LocalLogging = d.Get("local_logging").(int)
	//omit uuid
	return ret
}
