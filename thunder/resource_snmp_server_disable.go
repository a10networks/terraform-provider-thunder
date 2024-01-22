package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerDisable() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_disable`: Disable SNMP service\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerDisableCreate,
		UpdateContext: resourceSnmpServerDisableUpdate,
		ReadContext:   resourceSnmpServerDisableRead,
		DeleteContext: resourceSnmpServerDisableDelete,

		Schema: map[string]*schema.Schema{
			"a10cmsubagent": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable a10cmsubagent",
			},
			"traps": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable all traps on this partition",
						},
						"snmp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable all snmp traps on this partition",
						},
						"gslb": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable all gslb traps on this partition",
						},
						"vrrp_a": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable all vrrp-a on this partition",
						},
						"slb": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable all slb traps on this partition",
						},
						"slb_change": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable all slb-change traps on this partition",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSnmpServerDisableCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerDisableCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerDisable(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerDisableRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerDisableUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerDisableUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerDisable(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerDisableRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerDisableDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerDisableDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerDisable(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerDisableRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerDisableRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerDisable(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSnmpServerDisableTraps1477(d []interface{}) edpt.SnmpServerDisableTraps1477 {

	count1 := len(d)
	var ret edpt.SnmpServerDisableTraps1477
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.All = in["all"].(int)
		ret.Snmp = in["snmp"].(int)
		ret.Gslb = in["gslb"].(int)
		ret.VrrpA = in["vrrp_a"].(int)
		ret.Slb = in["slb"].(int)
		ret.SlbChange = in["slb_change"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointSnmpServerDisable(d *schema.ResourceData) edpt.SnmpServerDisable {
	var ret edpt.SnmpServerDisable
	ret.Inst.A10cmsubagent = d.Get("a10cmsubagent").(int)
	ret.Inst.Traps = getObjectSnmpServerDisableTraps1477(d.Get("traps").([]interface{}))
	//omit uuid
	return ret
}
