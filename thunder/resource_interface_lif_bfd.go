package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceLifBfd() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_lif_bfd`: Configure BFD (Bidirectional Forwarding Detection)\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceLifBfdCreate,
		UpdateContext: resourceInterfaceLifBfdUpdate,
		ReadContext:   resourceInterfaceLifBfdRead,
		DeleteContext: resourceInterfaceLifBfdDelete,

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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ifname": {
				Type: schema.TypeString, Required: true, Description: "Ifname",
			},
		},
	}
}
func resourceInterfaceLifBfdCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifBfdCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifBfd(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLifBfdRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceLifBfdUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifBfdUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifBfd(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLifBfdRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceLifBfdDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifBfdDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifBfd(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceLifBfdRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifBfdRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifBfd(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceLifBfdAuthentication(d []interface{}) edpt.InterfaceLifBfdAuthentication {

	count1 := len(d)
	var ret edpt.InterfaceLifBfdAuthentication
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyId = in["key_id"].(int)
		ret.Method = in["method"].(string)
		ret.Password = in["password"].(string)
		//omit encrypted
	}
	return ret
}

func getObjectInterfaceLifBfdIntervalCfg(d []interface{}) edpt.InterfaceLifBfdIntervalCfg {

	count1 := len(d)
	var ret edpt.InterfaceLifBfdIntervalCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interval = in["interval"].(int)
		ret.MinRx = in["min_rx"].(int)
		ret.Multiplier = in["multiplier"].(int)
	}
	return ret
}

func dataToEndpointInterfaceLifBfd(d *schema.ResourceData) edpt.InterfaceLifBfd {
	var ret edpt.InterfaceLifBfd
	ret.Inst.Authentication = getObjectInterfaceLifBfdAuthentication(d.Get("authentication").([]interface{}))
	ret.Inst.Demand = d.Get("demand").(int)
	ret.Inst.Echo = d.Get("echo").(int)
	ret.Inst.IntervalCfg = getObjectInterfaceLifBfdIntervalCfg(d.Get("interval_cfg").([]interface{}))
	//omit uuid
	ret.Inst.Ifname = d.Get("ifname").(string)
	return ret
}
