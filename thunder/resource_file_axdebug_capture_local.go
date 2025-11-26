package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileAxdebugCaptureLocal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_axdebug_capture_local`: axdebug file information and management commands\n\n__PLACEHOLDER__",
		CreateContext: resourceFileAxdebugCaptureLocalCreate,
		UpdateContext: resourceFileAxdebugCaptureLocalUpdate,
		ReadContext:   resourceFileAxdebugCaptureLocalRead,
		DeleteContext: resourceFileAxdebugCaptureLocalDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'delete': delete;",
			},
			"file": {
				Type: schema.TypeString, Optional: true, Description: "axdebug local capture file name",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFileAxdebugCaptureLocalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAxdebugCaptureLocalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAxdebugCaptureLocal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileAxdebugCaptureLocalRead(ctx, d, meta)
	}
	return diags
}

func resourceFileAxdebugCaptureLocalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAxdebugCaptureLocalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAxdebugCaptureLocal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileAxdebugCaptureLocalRead(ctx, d, meta)
	}
	return diags
}
func resourceFileAxdebugCaptureLocalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAxdebugCaptureLocalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAxdebugCaptureLocal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileAxdebugCaptureLocalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAxdebugCaptureLocalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAxdebugCaptureLocal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFileAxdebugCaptureLocal(d *schema.ResourceData) edpt.FileAxdebugCaptureLocal {
	var ret edpt.FileAxdebugCaptureLocal
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.File = d.Get("file").(string)
	//omit uuid
	return ret
}
