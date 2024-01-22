package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAudit() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_audit`: Configure audit\n\n__PLACEHOLDER__",
		CreateContext: resourceAuditCreate,
		UpdateContext: resourceAuditUpdate,
		ReadContext:   resourceAuditRead,
		DeleteContext: resourceAuditDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Enable audit service",
			},
			"privilege": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable privilege command for audit service",
			},
			"size": {
				Type: schema.TypeInt, Optional: true, Description: "Config audit buffer size, default is 20,000 (Audit buffer size(in items), default is 20,000)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAuditCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAuditCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAudit(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAuditRead(ctx, d, meta)
	}
	return diags
}

func resourceAuditUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAuditUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAudit(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAuditRead(ctx, d, meta)
	}
	return diags
}
func resourceAuditDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAuditDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAudit(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAuditRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAuditRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAudit(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAudit(d *schema.ResourceData) edpt.Audit {
	var ret edpt.Audit
	ret.Inst.Enable = d.Get("enable").(int)
	ret.Inst.Privilege = d.Get("privilege").(int)
	ret.Inst.Size = d.Get("size").(int)
	//omit uuid
	return ret
}
