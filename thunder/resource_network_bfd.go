package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkBfd() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_bfd`: Configure BFD (Bidirectional Forwarding Detection)\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkBfdCreate,
		UpdateContext: resourceNetworkBfdUpdate,
		ReadContext:   resourceNetworkBfdRead,
		DeleteContext: resourceNetworkBfdDelete,

		Schema: map[string]*schema.Schema{
			"echo": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable BFD Echo",
			},
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable BFD",
			},
			"interval_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interval": {
							Type: schema.TypeInt, Optional: true, Default: 800, Description: "Transmit interval between BFD packets (Milliseconds (default: 800))",
						},
						"min_rx": {
							Type: schema.TypeInt, Optional: true, Default: 800, Description: "Minimum receive interval capability (Milliseconds (default: 800))",
						},
						"multiplier": {
							Type: schema.TypeInt, Optional: true, Default: 4, Description: "Multiplier value used to compute holddown (value used to multiply the interval (default: 4))",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceNetworkBfdCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkBfdCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkBfd(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkBfdRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkBfdUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkBfdUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkBfd(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkBfdRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkBfdDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkBfdDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkBfd(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkBfdRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkBfdRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkBfd(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectNetworkBfdIntervalCfg(d []interface{}) edpt.NetworkBfdIntervalCfg {

	count1 := len(d)
	var ret edpt.NetworkBfdIntervalCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interval = in["interval"].(int)
		ret.MinRx = in["min_rx"].(int)
		ret.Multiplier = in["multiplier"].(int)
	}
	return ret
}

func dataToEndpointNetworkBfd(d *schema.ResourceData) edpt.NetworkBfd {
	var ret edpt.NetworkBfd
	ret.Inst.Echo = d.Get("echo").(int)
	ret.Inst.Enable = d.Get("enable").(int)
	ret.Inst.IntervalCfg = getObjectNetworkBfdIntervalCfg(d.Get("interval_cfg").([]interface{}))
	//omit uuid
	return ret
}
