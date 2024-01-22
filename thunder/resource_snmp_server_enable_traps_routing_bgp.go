package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerEnableTrapsRoutingBgp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_enable_traps_routing_bgp`: Enable bgp traps\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerEnableTrapsRoutingBgpCreate,
		UpdateContext: resourceSnmpServerEnableTrapsRoutingBgpUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsRoutingBgpRead,
		DeleteContext: resourceSnmpServerEnableTrapsRoutingBgpDelete,

		Schema: map[string]*schema.Schema{
			"ax": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bgpestablishednotification": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable bgpEstablishedNotification traps",
						},
						"bgpbackwardtransnotification": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable bgpBackwardTransNotification traps",
						},
						"bgpprefixthresholdexceedednotification": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable bgpPrefixThresholdExceededNotification traps",
						},
						"bgpprefixthresholdclearnotification": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable bgpPrefixThresholdClearNotification traps",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"bgpbackwardtransnotification": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable bgpBackwardTransNotification traps",
			},
			"bgpestablishednotification": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable bgpEstablishedNotification traps",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSnmpServerEnableTrapsRoutingBgpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsRoutingBgpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsRoutingBgp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsRoutingBgpRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerEnableTrapsRoutingBgpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsRoutingBgpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsRoutingBgp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsRoutingBgpRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerEnableTrapsRoutingBgpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsRoutingBgpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsRoutingBgp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerEnableTrapsRoutingBgpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsRoutingBgpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsRoutingBgp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSnmpServerEnableTrapsRoutingBgpAx1478(d []interface{}) edpt.SnmpServerEnableTrapsRoutingBgpAx1478 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsRoutingBgpAx1478
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bgpestablishednotification = in["bgpestablishednotification"].(int)
		ret.Bgpbackwardtransnotification = in["bgpbackwardtransnotification"].(int)
		ret.Bgpprefixthresholdexceedednotification = in["bgpprefixthresholdexceedednotification"].(int)
		ret.Bgpprefixthresholdclearnotification = in["bgpprefixthresholdclearnotification"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointSnmpServerEnableTrapsRoutingBgp(d *schema.ResourceData) edpt.SnmpServerEnableTrapsRoutingBgp {
	var ret edpt.SnmpServerEnableTrapsRoutingBgp
	ret.Inst.Ax = getObjectSnmpServerEnableTrapsRoutingBgpAx1478(d.Get("ax").([]interface{}))
	ret.Inst.Bgpbackwardtransnotification = d.Get("bgpbackwardtransnotification").(int)
	ret.Inst.Bgpestablishednotification = d.Get("bgpestablishednotification").(int)
	//omit uuid
	return ret
}
