package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthExternalCreate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_external_create`: Create external health monitor script\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthExternalCreateCreate,
		UpdateContext: resourceHealthExternalCreateUpdate,
		ReadContext:   resourceHealthExternalCreateRead,
		DeleteContext: resourceHealthExternalCreateDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString, Optional: true, Description: "Describe health external monitor script briefly",
			},
			"file_name": {
				Type: schema.TypeString, Optional: true, Description: "External health monitor script file name",
			},
		},
	}
}
func resourceHealthExternalCreateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthExternalCreateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthExternalCreate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthExternalCreateRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthExternalCreateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthExternalCreateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthExternalCreate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthExternalCreateRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthExternalCreateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthExternalCreateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthExternalCreate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthExternalCreateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthExternalCreateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthExternalCreate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHealthExternalCreate(d *schema.ResourceData) edpt.HealthExternalCreate {
	var ret edpt.HealthExternalCreate
	ret.Inst.Description = d.Get("description").(string)
	ret.Inst.FileName = d.Get("file_name").(string)
	return ret
}
