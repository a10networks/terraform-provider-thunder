package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthExternalDelete() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_external_delete`: Delete external health monitor script\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthExternalDeleteCreate,
		UpdateContext: resourceHealthExternalDeleteUpdate,
		ReadContext:   resourceHealthExternalDeleteRead,
		DeleteContext: resourceHealthExternalDeleteDelete,

		Schema: map[string]*schema.Schema{
			"file_name": {
				Type: schema.TypeString, Optional: true, Description: "External health monitor script file name",
			},
		},
	}
}
func resourceHealthExternalDeleteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthExternalDeleteCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthExternalDelete(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthExternalDeleteRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthExternalDeleteUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthExternalDeleteUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthExternalDelete(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthExternalDeleteRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthExternalDeleteDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthExternalDeleteDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthExternalDelete(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthExternalDeleteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthExternalDeleteRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthExternalDelete(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHealthExternalDelete(d *schema.ResourceData) edpt.HealthExternalDelete {
	var ret edpt.HealthExternalDelete
	ret.Inst.FileName = d.Get("file_name").(string)
	return ret
}
