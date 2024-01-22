package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerEnableTrapsVrrpA() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_enable_traps_vrrp_a`: Enable VRRP-A group traps\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerEnableTrapsVrrpACreate,
		UpdateContext: resourceSnmpServerEnableTrapsVrrpAUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsVrrpARead,
		DeleteContext: resourceSnmpServerEnableTrapsVrrpADelete,

		Schema: map[string]*schema.Schema{
			"active": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable VRRP-A active trap",
			},
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all VRRP-A group traps",
			},
			"standby": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable VRRP-A standby trap",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSnmpServerEnableTrapsVrrpACreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsVrrpACreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsVrrpA(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsVrrpARead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerEnableTrapsVrrpAUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsVrrpAUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsVrrpA(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsVrrpARead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerEnableTrapsVrrpADelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsVrrpADelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsVrrpA(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerEnableTrapsVrrpARead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsVrrpARead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsVrrpA(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerEnableTrapsVrrpA(d *schema.ResourceData) edpt.SnmpServerEnableTrapsVrrpA {
	var ret edpt.SnmpServerEnableTrapsVrrpA
	ret.Inst.Active = d.Get("active").(int)
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.Standby = d.Get("standby").(int)
	//omit uuid
	return ret
}
