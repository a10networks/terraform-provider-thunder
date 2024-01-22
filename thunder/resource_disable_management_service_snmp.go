package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDisableManagementServiceSnmp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_disable_management_service_snmp`: SNMP service\n\n__PLACEHOLDER__",
		CreateContext: resourceDisableManagementServiceSnmpCreate,
		UpdateContext: resourceDisableManagementServiceSnmpUpdate,
		ReadContext:   resourceDisableManagementServiceSnmpRead,
		DeleteContext: resourceDisableManagementServiceSnmpDelete,

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
func resourceDisableManagementServiceSnmpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDisableManagementServiceSnmpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDisableManagementServiceSnmp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDisableManagementServiceSnmpRead(ctx, d, meta)
	}
	return diags
}

func resourceDisableManagementServiceSnmpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDisableManagementServiceSnmpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDisableManagementServiceSnmp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDisableManagementServiceSnmpRead(ctx, d, meta)
	}
	return diags
}
func resourceDisableManagementServiceSnmpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDisableManagementServiceSnmpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDisableManagementServiceSnmp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDisableManagementServiceSnmpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDisableManagementServiceSnmpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDisableManagementServiceSnmp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDisableManagementServiceSnmp(d *schema.ResourceData) edpt.DisableManagementServiceSnmp {
	var ret edpt.DisableManagementServiceSnmp
	ret.Inst.Management = d.Get("management").(int)
	//omit uuid
	return ret
}
