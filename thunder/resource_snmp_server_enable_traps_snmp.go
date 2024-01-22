package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerEnableTrapsSnmp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_enable_traps_snmp`: Enable SNMP group traps\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerEnableTrapsSnmpCreate,
		UpdateContext: resourceSnmpServerEnableTrapsSnmpUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsSnmpRead,
		DeleteContext: resourceSnmpServerEnableTrapsSnmpDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all SNMP group traps",
			},
			"linkdown": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SNMP link-down trap",
			},
			"linkup": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SNMP link-up trap",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSnmpServerEnableTrapsSnmpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsSnmpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsSnmp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsSnmpRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerEnableTrapsSnmpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsSnmpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsSnmp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsSnmpRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerEnableTrapsSnmpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsSnmpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsSnmp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerEnableTrapsSnmpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsSnmpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsSnmp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerEnableTrapsSnmp(d *schema.ResourceData) edpt.SnmpServerEnableTrapsSnmp {
	var ret edpt.SnmpServerEnableTrapsSnmp
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.Linkdown = d.Get("linkdown").(int)
	ret.Inst.Linkup = d.Get("linkup").(int)
	//omit uuid
	return ret
}
