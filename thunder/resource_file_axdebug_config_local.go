package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileAxdebugConfigLocal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_axdebug_config_local`: axdebug file information and management commands\n\n__PLACEHOLDER__",
		CreateContext: resourceFileAxdebugConfigLocalCreate,
		UpdateContext: resourceFileAxdebugConfigLocalUpdate,
		ReadContext:   resourceFileAxdebugConfigLocalRead,
		DeleteContext: resourceFileAxdebugConfigLocalDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'export': export; 'delete': delete;",
			},
			"file": {
				Type: schema.TypeString, Optional: true, Description: "axdebug local config file name",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFileAxdebugConfigLocalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAxdebugConfigLocalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAxdebugConfigLocal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileAxdebugConfigLocalRead(ctx, d, meta)
	}
	return diags
}

func resourceFileAxdebugConfigLocalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAxdebugConfigLocalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAxdebugConfigLocal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileAxdebugConfigLocalRead(ctx, d, meta)
	}
	return diags
}
func resourceFileAxdebugConfigLocalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAxdebugConfigLocalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAxdebugConfigLocal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileAxdebugConfigLocalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAxdebugConfigLocalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAxdebugConfigLocal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFileAxdebugConfigLocal(d *schema.ResourceData) edpt.FileAxdebugConfigLocal {
	var ret edpt.FileAxdebugConfigLocal
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.File = d.Get("file").(string)
	//omit uuid
	return ret
}
