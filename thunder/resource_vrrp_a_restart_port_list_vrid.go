package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpARestartPortListVrid() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vrrp_a_restart_port_list_vrid`: Specify VRRP-A vrid\n\n__PLACEHOLDER__",
		CreateContext: resourceVrrpARestartPortListVridCreate,
		UpdateContext: resourceVrrpARestartPortListVridUpdate,
		ReadContext:   resourceVrrpARestartPortListVridRead,
		DeleteContext: resourceVrrpARestartPortListVridDelete,

		Schema: map[string]*schema.Schema{
			"ethernet_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"flap_ethernet_start": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet Port",
						},
						"flap_ethernet_end": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet Port",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vrid_val": {
				Type: schema.TypeInt, Required: true, Description: "Specify ha VRRP-A vrid",
			},
		},
	}
}
func resourceVrrpARestartPortListVridCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpARestartPortListVridCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpARestartPortListVrid(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpARestartPortListVridRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpARestartPortListVridUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpARestartPortListVridUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpARestartPortListVrid(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpARestartPortListVridRead(ctx, d, meta)
	}
	return diags
}
func resourceVrrpARestartPortListVridDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpARestartPortListVridDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpARestartPortListVrid(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVrrpARestartPortListVridRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpARestartPortListVridRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpARestartPortListVrid(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceVrrpARestartPortListVridEthernetCfg(d []interface{}) []edpt.VrrpARestartPortListVridEthernetCfg {

	count1 := len(d)
	ret := make([]edpt.VrrpARestartPortListVridEthernetCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpARestartPortListVridEthernetCfg
		oi.FlapEthernetStart = in["flap_ethernet_start"].(int)
		oi.FlapEthernetEnd = in["flap_ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVrrpARestartPortListVrid(d *schema.ResourceData) edpt.VrrpARestartPortListVrid {
	var ret edpt.VrrpARestartPortListVrid
	ret.Inst.EthernetCfg = getSliceVrrpARestartPortListVridEthernetCfg(d.Get("ethernet_cfg").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VridVal = d.Get("vrid_val").(int)
	return ret
}
