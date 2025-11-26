package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCmUpdateFileNameRef() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_cm_update_file_name_ref`: update the bind name\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemCmUpdateFileNameRefCreate,
		UpdateContext: resourceSystemCmUpdateFileNameRefUpdate,
		ReadContext:   resourceSystemCmUpdateFileNameRefRead,
		DeleteContext: resourceSystemCmUpdateFileNameRefDelete,

		Schema: map[string]*schema.Schema{
			"dest_name": {
				Type: schema.TypeString, Optional: true, Description: "bind dest name",
			},
			"id1": {
				Type: schema.TypeInt, Optional: true, Description: "Specify unique Partition id",
			},
			"source_name": {
				Type: schema.TypeString, Optional: true, Description: "bind source name",
			},
		},
	}
}
func resourceSystemCmUpdateFileNameRefCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCmUpdateFileNameRefCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCmUpdateFileNameRef(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemCmUpdateFileNameRefRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemCmUpdateFileNameRefUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCmUpdateFileNameRefUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCmUpdateFileNameRef(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemCmUpdateFileNameRefRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemCmUpdateFileNameRefDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCmUpdateFileNameRefDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCmUpdateFileNameRef(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemCmUpdateFileNameRefRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCmUpdateFileNameRefRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCmUpdateFileNameRef(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemCmUpdateFileNameRef(d *schema.ResourceData) edpt.SystemCmUpdateFileNameRef {
	var ret edpt.SystemCmUpdateFileNameRef
	ret.Inst.Dest_name = d.Get("dest_name").(string)
	ret.Inst.Id1 = d.Get("id1").(int)
	ret.Inst.Source_name = d.Get("source_name").(string)
	return ret
}
