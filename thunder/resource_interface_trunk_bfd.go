package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTrunkBfd() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_trunk_bfd`: Configure BFD (Bidirectional Forwarding Detection)\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceTrunkBfdCreate,
		UpdateContext: resourceInterfaceTrunkBfdUpdate,
		ReadContext:   resourceInterfaceTrunkBfdRead,
		DeleteContext: resourceInterfaceTrunkBfdDelete,

		Schema: map[string]*schema.Schema{
			"authentication": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key_id": {
							Type: schema.TypeInt, Optional: true, Description: "Key ID",
						},
						"method": {
							Type: schema.TypeString, Optional: true, Description: "'md5': Keyed MD5; 'meticulous-md5': Meticulous Keyed MD5; 'meticulous-sha1': Meticulous Keyed SHA1; 'sha1': Keyed SHA1; 'simple': Simple Password;",
						},
						"password": {
							Type: schema.TypeString, Optional: true, Description: "Key String",
						},
					},
				},
			},
			"demand": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Demand mode",
			},
			"echo": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable BFD Echo",
			},
			"interval_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interval": {
							Type: schema.TypeInt, Optional: true, Description: "Transmit interval between BFD packets (Milliseconds)",
						},
						"min_rx": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum receive interval capability (Milliseconds)",
						},
						"multiplier": {
							Type: schema.TypeInt, Optional: true, Description: "Multiplier value used to compute holddown (value used to multiply the interval)",
						},
					},
				},
			},
			"per_member_port": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"local_address": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 local-address",
						},
						"neighbor_address": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 neighbor address",
						},
						"ipv6_local": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 local-address",
						},
						"ipv6_nbr": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 neighbor-address",
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
			"ifnum": {
				Type: schema.TypeString, Required: true, Description: "Ifnum",
			},
		},
	}
}
func resourceInterfaceTrunkBfdCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkBfdCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkBfd(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkBfdRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceTrunkBfdUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkBfdUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkBfd(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkBfdRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceTrunkBfdDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkBfdDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkBfd(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceTrunkBfdRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkBfdRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkBfd(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceTrunkBfdAuthentication(d []interface{}) edpt.InterfaceTrunkBfdAuthentication {

	count1 := len(d)
	var ret edpt.InterfaceTrunkBfdAuthentication
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyId = in["key_id"].(int)
		ret.Method = in["method"].(string)
		ret.Password = in["password"].(string)
		//omit encrypted
	}
	return ret
}

func getObjectInterfaceTrunkBfdIntervalCfg(d []interface{}) edpt.InterfaceTrunkBfdIntervalCfg {

	count1 := len(d)
	var ret edpt.InterfaceTrunkBfdIntervalCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interval = in["interval"].(int)
		ret.MinRx = in["min_rx"].(int)
		ret.Multiplier = in["multiplier"].(int)
	}
	return ret
}

func getObjectInterfaceTrunkBfdPerMemberPort742(d []interface{}) edpt.InterfaceTrunkBfdPerMemberPort742 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkBfdPerMemberPort742
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LocalAddress = in["local_address"].(string)
		ret.NeighborAddress = in["neighbor_address"].(string)
		ret.Ipv6Local = in["ipv6_local"].(string)
		ret.Ipv6Nbr = in["ipv6_nbr"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointInterfaceTrunkBfd(d *schema.ResourceData) edpt.InterfaceTrunkBfd {
	var ret edpt.InterfaceTrunkBfd
	ret.Inst.Authentication = getObjectInterfaceTrunkBfdAuthentication(d.Get("authentication").([]interface{}))
	ret.Inst.Demand = d.Get("demand").(int)
	ret.Inst.Echo = d.Get("echo").(int)
	ret.Inst.IntervalCfg = getObjectInterfaceTrunkBfdIntervalCfg(d.Get("interval_cfg").([]interface{}))
	ret.Inst.PerMemberPort = getObjectInterfaceTrunkBfdPerMemberPort742(d.Get("per_member_port").([]interface{}))
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
