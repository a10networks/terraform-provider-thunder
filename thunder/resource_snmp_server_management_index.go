package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerManagementIndex() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_management_index`: Define index of management interface\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerManagementIndexCreate,
		UpdateContext: resourceSnmpServerManagementIndexUpdate,
		ReadContext:   resourceSnmpServerManagementIndexRead,
		DeleteContext: resourceSnmpServerManagementIndexDelete,

		Schema: map[string]*schema.Schema{
			"mgmt_index": {
				Type: schema.TypeInt, Optional: true, Description: "index number",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSnmpServerManagementIndexCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerManagementIndexCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerManagementIndex(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerManagementIndexRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerManagementIndexUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerManagementIndexUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerManagementIndex(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerManagementIndexRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerManagementIndexDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerManagementIndexDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerManagementIndex(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerManagementIndexRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerManagementIndexRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerManagementIndex(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerManagementIndex(d *schema.ResourceData) edpt.SnmpServerManagementIndex {
	var ret edpt.SnmpServerManagementIndex
	ret.Inst.MgmtIndex = d.Get("mgmt_index").(int)
	//omit uuid
	return ret
}
