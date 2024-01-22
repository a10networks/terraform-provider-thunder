package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpARestartPortList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vrrp_a_restart_port_list`: Ports to be restarted on standby system after transition\n\n__PLACEHOLDER__",
		CreateContext: resourceVrrpARestartPortListCreate,
		UpdateContext: resourceVrrpARestartPortListUpdate,
		ReadContext:   resourceVrrpARestartPortListRead,
		DeleteContext: resourceVrrpARestartPortListDelete,

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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vrid_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vrid_val": {
							Type: schema.TypeInt, Required: true, Description: "Specify ha VRRP-A vrid",
						},
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
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
		},
	}
}
func resourceVrrpARestartPortListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpARestartPortListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpARestartPortList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpARestartPortListRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpARestartPortListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpARestartPortListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpARestartPortList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpARestartPortListRead(ctx, d, meta)
	}
	return diags
}
func resourceVrrpARestartPortListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpARestartPortListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpARestartPortList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVrrpARestartPortListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpARestartPortListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpARestartPortList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceVrrpARestartPortListEthernetCfg(d []interface{}) []edpt.VrrpARestartPortListEthernetCfg {

	count1 := len(d)
	ret := make([]edpt.VrrpARestartPortListEthernetCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpARestartPortListEthernetCfg
		oi.FlapEthernetStart = in["flap_ethernet_start"].(int)
		oi.FlapEthernetEnd = in["flap_ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpARestartPortListVridList(d []interface{}) []edpt.VrrpARestartPortListVridList {

	count1 := len(d)
	ret := make([]edpt.VrrpARestartPortListVridList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpARestartPortListVridList
		oi.VridVal = in["vrid_val"].(int)
		oi.EthernetCfg = getSliceVrrpARestartPortListVridListEthernetCfg(in["ethernet_cfg"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpARestartPortListVridListEthernetCfg(d []interface{}) []edpt.VrrpARestartPortListVridListEthernetCfg {

	count1 := len(d)
	ret := make([]edpt.VrrpARestartPortListVridListEthernetCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpARestartPortListVridListEthernetCfg
		oi.FlapEthernetStart = in["flap_ethernet_start"].(int)
		oi.FlapEthernetEnd = in["flap_ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVrrpARestartPortList(d *schema.ResourceData) edpt.VrrpARestartPortList {
	var ret edpt.VrrpARestartPortList
	ret.Inst.EthernetCfg = getSliceVrrpARestartPortListEthernetCfg(d.Get("ethernet_cfg").([]interface{}))
	//omit uuid
	ret.Inst.VridList = getSliceVrrpARestartPortListVridList(d.Get("vrid_list").([]interface{}))
	return ret
}
