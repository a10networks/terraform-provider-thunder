package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerEnableTrapsLsn() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_enable_traps_lsn`: Enable LSN group traps\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerEnableTrapsLsnCreate,
		UpdateContext: resourceSnmpServerEnableTrapsLsnUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsLsnRead,
		DeleteContext: resourceSnmpServerEnableTrapsLsnDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all LSN group traps",
			},
			"fixed_nat_port_mapping_file_change": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable LSN trap when fixed nat port mapping file change",
			},
			"max_ipport_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 64512, Description: "Maximum threshold",
			},
			"max_port_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 655350000, Description: "Maximum threshold",
			},
			"per_ip_port_usage_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable LSN trap when IP total port usage reaches the threshold (default 64512)",
			},
			"total_port_usage_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable LSN trap when NAT total port usage reaches the threshold (default 655350000)",
			},
			"traffic_exceeded": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable LSN trap when NAT pool reaches the threshold",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSnmpServerEnableTrapsLsnCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsLsnCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsLsn(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsLsnRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerEnableTrapsLsnUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsLsnUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsLsn(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsLsnRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerEnableTrapsLsnDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsLsnDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsLsn(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerEnableTrapsLsnRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsLsnRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsLsn(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerEnableTrapsLsn(d *schema.ResourceData) edpt.SnmpServerEnableTrapsLsn {
	var ret edpt.SnmpServerEnableTrapsLsn
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.FixedNatPortMappingFileChange = d.Get("fixed_nat_port_mapping_file_change").(int)
	ret.Inst.MaxIpportThreshold = d.Get("max_ipport_threshold").(int)
	ret.Inst.MaxPortThreshold = d.Get("max_port_threshold").(int)
	ret.Inst.PerIpPortUsageThreshold = d.Get("per_ip_port_usage_threshold").(int)
	ret.Inst.TotalPortUsageThreshold = d.Get("total_port_usage_threshold").(int)
	ret.Inst.TrafficExceeded = d.Get("traffic_exceeded").(int)
	//omit uuid
	return ret
}
