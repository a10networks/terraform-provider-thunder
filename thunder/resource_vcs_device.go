package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVcsDevice() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vcs_device`: VCS Device\n\n__PLACEHOLDER__",
		CreateContext: resourceVcsDeviceCreate,
		UpdateContext: resourceVcsDeviceUpdate,
		ReadContext:   resourceVcsDeviceRead,
		DeleteContext: resourceVcsDeviceDelete,

		Schema: map[string]*schema.Schema{
			"affinity_vrrp_a_vrid": {
				Type: schema.TypeInt, Optional: true, Description: "vrid-group",
			},
			"device": {
				Type: schema.TypeInt, Required: true, Description: "Device ID",
			},
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable",
			},
			"ethernet_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ethernet": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet (Ethernet interface number)",
						},
					},
				},
			},
			"management": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Management interface",
			},
			"priority": {
				Type: schema.TypeInt, Optional: true, Description: "Device priority",
			},
			"trunk_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trunk": {
							Type: schema.TypeInt, Optional: true, Description: "Trunk interface (Trunk interface number)",
						},
					},
				},
			},
			"ttl": {
				Type: schema.TypeInt, Optional: true, Default: 64, Description: "TTL of the VCS packet",
			},
			"unicast_port": {
				Type: schema.TypeInt, Optional: true, Default: 41216, Description: "Port used in unicast communication (Port number)",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ve_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ve": {
							Type: schema.TypeInt, Optional: true, Description: "VE interface (VE interface number)",
						},
					},
				},
			},
		},
	}
}
func resourceVcsDeviceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsDeviceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsDevice(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsDeviceRead(ctx, d, meta)
	}
	return diags
}

func resourceVcsDeviceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsDeviceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsDevice(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsDeviceRead(ctx, d, meta)
	}
	return diags
}
func resourceVcsDeviceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsDeviceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsDevice(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVcsDeviceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsDeviceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsDevice(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceVcsDeviceEthernetCfg(d []interface{}) []edpt.VcsDeviceEthernetCfg {

	count1 := len(d)
	ret := make([]edpt.VcsDeviceEthernetCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VcsDeviceEthernetCfg
		oi.Ethernet = in["ethernet"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVcsDeviceTrunkCfg(d []interface{}) []edpt.VcsDeviceTrunkCfg {

	count1 := len(d)
	ret := make([]edpt.VcsDeviceTrunkCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VcsDeviceTrunkCfg
		oi.Trunk = in["trunk"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVcsDeviceVeCfg(d []interface{}) []edpt.VcsDeviceVeCfg {

	count1 := len(d)
	ret := make([]edpt.VcsDeviceVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VcsDeviceVeCfg
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVcsDevice(d *schema.ResourceData) edpt.VcsDevice {
	var ret edpt.VcsDevice
	ret.Inst.AffinityVrrpAVrid = d.Get("affinity_vrrp_a_vrid").(int)
	ret.Inst.Device = d.Get("device").(int)
	ret.Inst.Enable = d.Get("enable").(int)
	ret.Inst.EthernetCfg = getSliceVcsDeviceEthernetCfg(d.Get("ethernet_cfg").([]interface{}))
	ret.Inst.Management = d.Get("management").(int)
	ret.Inst.Priority = d.Get("priority").(int)
	ret.Inst.TrunkCfg = getSliceVcsDeviceTrunkCfg(d.Get("trunk_cfg").([]interface{}))
	ret.Inst.Ttl = d.Get("ttl").(int)
	ret.Inst.UnicastPort = d.Get("unicast_port").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VeCfg = getSliceVcsDeviceVeCfg(d.Get("ve_cfg").([]interface{}))
	return ret
}
