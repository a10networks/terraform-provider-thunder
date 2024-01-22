package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerEnableTrapsNetwork() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_enable_traps_network`: Enable network group traps\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerEnableTrapsNetworkCreate,
		UpdateContext: resourceSnmpServerEnableTrapsNetworkUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsNetworkRead,
		DeleteContext: resourceSnmpServerEnableTrapsNetworkDelete,

		Schema: map[string]*schema.Schema{
			"trunk_port_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable network trunk-port-threshold trap",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSnmpServerEnableTrapsNetworkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsNetworkCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsNetwork(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsNetworkRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerEnableTrapsNetworkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsNetworkUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsNetwork(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsNetworkRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerEnableTrapsNetworkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsNetworkDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsNetwork(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerEnableTrapsNetworkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsNetworkRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsNetwork(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerEnableTrapsNetwork(d *schema.ResourceData) edpt.SnmpServerEnableTrapsNetwork {
	var ret edpt.SnmpServerEnableTrapsNetwork
	ret.Inst.TrunkPortThreshold = d.Get("trunk_port_threshold").(int)
	//omit uuid
	return ret
}
