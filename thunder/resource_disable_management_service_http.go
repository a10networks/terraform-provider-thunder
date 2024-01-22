package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDisableManagementServiceHttp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_disable_management_service_http`: HTTP service\n\n__PLACEHOLDER__",
		CreateContext: resourceDisableManagementServiceHttpCreate,
		UpdateContext: resourceDisableManagementServiceHttpUpdate,
		ReadContext:   resourceDisableManagementServiceHttpRead,
		DeleteContext: resourceDisableManagementServiceHttpDelete,

		Schema: map[string]*schema.Schema{
			"management": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Management port",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDisableManagementServiceHttpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDisableManagementServiceHttpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDisableManagementServiceHttp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDisableManagementServiceHttpRead(ctx, d, meta)
	}
	return diags
}

func resourceDisableManagementServiceHttpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDisableManagementServiceHttpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDisableManagementServiceHttp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDisableManagementServiceHttpRead(ctx, d, meta)
	}
	return diags
}
func resourceDisableManagementServiceHttpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDisableManagementServiceHttpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDisableManagementServiceHttp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDisableManagementServiceHttpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDisableManagementServiceHttpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDisableManagementServiceHttp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDisableManagementServiceHttp(d *schema.ResourceData) edpt.DisableManagementServiceHttp {
	var ret edpt.DisableManagementServiceHttp
	ret.Inst.Management = d.Get("management").(int)
	//omit uuid
	return ret
}
