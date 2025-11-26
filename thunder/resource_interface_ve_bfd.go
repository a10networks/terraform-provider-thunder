package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceVeBfd() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_ve_bfd`: Configure BFD (Bidirectional Forwarding Detection)\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceVeBfdCreate,
		UpdateContext: resourceInterfaceVeBfdUpdate,
		ReadContext:   resourceInterfaceVeBfdRead,
		DeleteContext: resourceInterfaceVeBfdDelete,

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
			"ifnum": {
				Type: schema.TypeString, Required: true, Description: "Ifnum",
			},
		},
	}
}
func resourceInterfaceVeBfdCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeBfdCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVeBfd(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceVeBfdRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceVeBfdUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeBfdUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVeBfd(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceVeBfdRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceVeBfdDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeBfdDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVeBfd(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceVeBfdRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeBfdRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVeBfd(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceVeBfdAuthentication(d []interface{}) edpt.InterfaceVeBfdAuthentication {

	count1 := len(d)
	var ret edpt.InterfaceVeBfdAuthentication
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyId = in["key_id"].(int)
		ret.Method = in["method"].(string)
		ret.Password = in["password"].(string)
		//omit encrypted
	}
	return ret
}

func getObjectInterfaceVeBfdIntervalCfg(d []interface{}) edpt.InterfaceVeBfdIntervalCfg {

	count1 := len(d)
	var ret edpt.InterfaceVeBfdIntervalCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interval = in["interval"].(int)
		ret.MinRx = in["min_rx"].(int)
		ret.Multiplier = in["multiplier"].(int)
	}
	return ret
}

func dataToEndpointInterfaceVeBfd(d *schema.ResourceData) edpt.InterfaceVeBfd {
	var ret edpt.InterfaceVeBfd
	ret.Inst.Authentication = getObjectInterfaceVeBfdAuthentication(d.Get("authentication").([]interface{}))
	ret.Inst.Demand = d.Get("demand").(int)
	ret.Inst.Echo = d.Get("echo").(int)
	ret.Inst.IntervalCfg = getObjectInterfaceVeBfdIntervalCfg(d.Get("interval_cfg").([]interface{}))
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
