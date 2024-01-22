package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerEnableTrapsRoutingBgpAx() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_enable_traps_routing_bgp_ax`: Enable bgp ax traps\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerEnableTrapsRoutingBgpAxCreate,
		UpdateContext: resourceSnmpServerEnableTrapsRoutingBgpAxUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsRoutingBgpAxRead,
		DeleteContext: resourceSnmpServerEnableTrapsRoutingBgpAxDelete,

		Schema: map[string]*schema.Schema{
			"bgpbackwardtransnotification": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable bgpBackwardTransNotification traps",
			},
			"bgpestablishednotification": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable bgpEstablishedNotification traps",
			},
			"bgpprefixthresholdclearnotification": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable bgpPrefixThresholdClearNotification traps",
			},
			"bgpprefixthresholdexceedednotification": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable bgpPrefixThresholdExceededNotification traps",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSnmpServerEnableTrapsRoutingBgpAxCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsRoutingBgpAxCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsRoutingBgpAx(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsRoutingBgpAxRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerEnableTrapsRoutingBgpAxUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsRoutingBgpAxUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsRoutingBgpAx(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsRoutingBgpAxRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerEnableTrapsRoutingBgpAxDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsRoutingBgpAxDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsRoutingBgpAx(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerEnableTrapsRoutingBgpAxRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsRoutingBgpAxRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsRoutingBgpAx(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerEnableTrapsRoutingBgpAx(d *schema.ResourceData) edpt.SnmpServerEnableTrapsRoutingBgpAx {
	var ret edpt.SnmpServerEnableTrapsRoutingBgpAx
	ret.Inst.Bgpbackwardtransnotification = d.Get("bgpbackwardtransnotification").(int)
	ret.Inst.Bgpestablishednotification = d.Get("bgpestablishednotification").(int)
	ret.Inst.Bgpprefixthresholdclearnotification = d.Get("bgpprefixthresholdclearnotification").(int)
	ret.Inst.Bgpprefixthresholdexceedednotification = d.Get("bgpprefixthresholdexceedednotification").(int)
	//omit uuid
	return ret
}
