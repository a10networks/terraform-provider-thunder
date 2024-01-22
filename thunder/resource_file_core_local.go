package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileCoreLocal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_core_local`: file operation for core files\n\n__PLACEHOLDER__",
		CreateContext: resourceFileCoreLocalCreate,
		UpdateContext: resourceFileCoreLocalUpdate,
		ReadContext:   resourceFileCoreLocalRead,
		DeleteContext: resourceFileCoreLocalDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'export': export; 'delete': delete;",
			},
			"file": {
				Type: schema.TypeString, Optional: true, Description: "local core file name",
			},
			"slot": {
				Type: schema.TypeInt, Optional: true, Description: "specify slot id in chassis",
			},
		},
	}
}
func resourceFileCoreLocalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileCoreLocalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileCoreLocal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileCoreLocalRead(ctx, d, meta)
	}
	return diags
}

func resourceFileCoreLocalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileCoreLocalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileCoreLocal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileCoreLocalRead(ctx, d, meta)
	}
	return diags
}
func resourceFileCoreLocalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileCoreLocalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileCoreLocal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileCoreLocalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileCoreLocalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileCoreLocal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFileCoreLocal(d *schema.ResourceData) edpt.FileCoreLocal {
	var ret edpt.FileCoreLocal
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.Slot = d.Get("slot").(int)
	return ret
}
