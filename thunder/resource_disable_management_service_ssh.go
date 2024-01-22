package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDisableManagementServiceSsh() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_disable_management_service_ssh`: SSH service\n\n__PLACEHOLDER__",
		CreateContext: resourceDisableManagementServiceSshCreate,
		UpdateContext: resourceDisableManagementServiceSshUpdate,
		ReadContext:   resourceDisableManagementServiceSshRead,
		DeleteContext: resourceDisableManagementServiceSshDelete,

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
func resourceDisableManagementServiceSshCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDisableManagementServiceSshCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDisableManagementServiceSsh(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDisableManagementServiceSshRead(ctx, d, meta)
	}
	return diags
}

func resourceDisableManagementServiceSshUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDisableManagementServiceSshUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDisableManagementServiceSsh(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDisableManagementServiceSshRead(ctx, d, meta)
	}
	return diags
}
func resourceDisableManagementServiceSshDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDisableManagementServiceSshDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDisableManagementServiceSsh(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDisableManagementServiceSshRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDisableManagementServiceSshRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDisableManagementServiceSsh(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDisableManagementServiceSsh(d *schema.ResourceData) edpt.DisableManagementServiceSsh {
	var ret edpt.DisableManagementServiceSsh
	ret.Inst.Management = d.Get("management").(int)
	//omit uuid
	return ret
}
