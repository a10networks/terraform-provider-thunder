package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceZoneVlan() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_zone_vlan`: Vlan\n\n__PLACEHOLDER__",
		CreateContext: resourceZoneVlanCreate,
		UpdateContext: resourceZoneVlanUpdate,
		ReadContext:   resourceZoneVlanRead,
		DeleteContext: resourceZoneVlanDelete,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vlan_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan_start": {
							Type: schema.TypeInt, Optional: true, Description: "VLAN ID",
						},
						"vlan_end": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceZoneVlanCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceZoneVlanCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointZoneVlan(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceZoneVlanRead(ctx, d, meta)
	}
	return diags
}

func resourceZoneVlanUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceZoneVlanUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointZoneVlan(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceZoneVlanRead(ctx, d, meta)
	}
	return diags
}
func resourceZoneVlanDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceZoneVlanDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointZoneVlan(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceZoneVlanRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceZoneVlanRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointZoneVlan(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceZoneVlanVlanList(d []interface{}) []edpt.ZoneVlanVlanList {

	count1 := len(d)
	ret := make([]edpt.ZoneVlanVlanList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ZoneVlanVlanList
		oi.VlanStart = in["vlan_start"].(int)
		oi.VlanEnd = in["vlan_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointZoneVlan(d *schema.ResourceData) edpt.ZoneVlan {
	var ret edpt.ZoneVlan
	//omit uuid
	ret.Inst.VlanList = getSliceZoneVlanVlanList(d.Get("vlan_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
