package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbPlayerIdGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_player_id_global`: Player-id global commands\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbPlayerIdGlobalCreate,
		UpdateContext: resourceSlbPlayerIdGlobalUpdate,
		ReadContext:   resourceSlbPlayerIdGlobalRead,
		DeleteContext: resourceSlbPlayerIdGlobalDelete,

		Schema: map[string]*schema.Schema{
			"abs_max_expiration": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Absolute max record expiration value (default 10 minutes) (Absolute max record expiration time in minutes, default 10)",
			},
			"enable_64bit_player_id": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable 64 bit player id check. Default is 32 bit",
			},
			"enforcement_timer": {
				Type: schema.TypeInt, Optional: true, Default: 480, Description: "Time to playerid enforcement after bootup (default 480 seconds) (Time to playerid enforcement in seconds, default 480)",
			},
			"force_passive": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forces the device to be in passive mode (Only stats and no packet drops)",
			},
			"min_expiration": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Minimum record expiration value (default 1 min) (Min record expiration time in minutes, default 1)",
			},
			"pkt_activity_expiration": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Packet activity record expiration value (default 5 minutes) (Packet activity record expiration time in minutes, default 5)",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'total_playerids_created': Playerid records created; 'total_playerids_deleted': Playerid records deleted; 'total_abs_max_age_outs': Playerid records max time aged out; 'total_pkt_activity_age_outs': Playerid records idle timeout; 'total_invalid_playerid_pkts': Invalid playerid packets; 'total_invalid_playerid_drops': Invalid playerid packet drops; 'total_valid_playerid_pkts': Valid playerid packets;",
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
func resourceSlbPlayerIdGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPlayerIdGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPlayerIdGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbPlayerIdGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbPlayerIdGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPlayerIdGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPlayerIdGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbPlayerIdGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbPlayerIdGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPlayerIdGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPlayerIdGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbPlayerIdGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPlayerIdGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPlayerIdGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbPlayerIdGlobalSamplingEnable(d []interface{}) []edpt.SlbPlayerIdGlobalSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbPlayerIdGlobalSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbPlayerIdGlobalSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbPlayerIdGlobal(d *schema.ResourceData) edpt.SlbPlayerIdGlobal {
	var ret edpt.SlbPlayerIdGlobal
	ret.Inst.AbsMaxExpiration = d.Get("abs_max_expiration").(int)
	ret.Inst.Enable64bitPlayerId = d.Get("enable_64bit_player_id").(int)
	ret.Inst.EnforcementTimer = d.Get("enforcement_timer").(int)
	ret.Inst.ForcePassive = d.Get("force_passive").(int)
	ret.Inst.MinExpiration = d.Get("min_expiration").(int)
	ret.Inst.PktActivityExpiration = d.Get("pkt_activity_expiration").(int)
	ret.Inst.SamplingEnable = getSliceSlbPlayerIdGlobalSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
