package thunder

import (
	"context"
	"errors"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSnmpServerDisableTraps() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_disable_traps`: Disable l3v partition SNMP traps\n\nThis resource is private-partition-only",
		CreateContext: resourceSnmpServerDisableTrapsCreate,
		UpdateContext: resourceSnmpServerDisableTrapsUpdate,
		ReadContext:   resourceSnmpServerDisableTrapsRead,
		DeleteContext: resourceSnmpServerDisableTrapsDelete,
		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable all traps on this partition",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"gslb": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable all gslb traps on this partition",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"slb": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable all slb traps on this partition",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"slb_change": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable all slb-change traps on this partition",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"snmp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable all snmp traps on this partition",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vrrp_a": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable all vrrp-a on this partition",
				ValidateFunc: validation.IntBetween(0, 1),
			},
		},
	}
}

func resourceSnmpServerDisableTrapsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerDisableTrapsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		err := runtimeCheckSnmpServerDisableTraps(client)
		if err != nil {
			return diag.FromErr(err)
		}
		obj := dataToEndpointSnmpServerDisableTraps(d)
		d.SetId(obj.GetId())
		err = obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerDisableTrapsRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerDisableTrapsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerDisableTrapsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		err := runtimeCheckSnmpServerDisableTraps(client)
		if err != nil {
			return diag.FromErr(err)
		}
		obj := edpt.SnmpServerDisableTraps{}
		err = obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerDisableTrapsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerDisableTrapsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		err := runtimeCheckSnmpServerDisableTraps(client)
		if err != nil {
			return diag.FromErr(err)
		}
		obj := dataToEndpointSnmpServerDisableTraps(d)
		err = obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerDisableTrapsRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerDisableTrapsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerDisableTrapsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		err := runtimeCheckSnmpServerDisableTraps(client)
		if err != nil {
			return diag.FromErr(err)
		}
		obj := edpt.SnmpServerDisableTraps{}
		err = obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerDisableTraps(d *schema.ResourceData) edpt.SnmpServerDisableTraps {
	var ret edpt.SnmpServerDisableTraps
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.Gslb = d.Get("gslb").(int)
	ret.Inst.Slb = d.Get("slb").(int)
	ret.Inst.SlbChange = d.Get("slb_change").(int)
	ret.Inst.Snmp = d.Get("snmp").(int)
	//omit uuid
	ret.Inst.VrrpA = d.Get("vrrp_a").(int)
	return ret
}

func runtimeCheckSnmpServerDisableTraps(th Thunder) error {
	if th.Partition == "shared" {
		return errors.New("This resource is private-partition-only")
	}
	return nil
}
