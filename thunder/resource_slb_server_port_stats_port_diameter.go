package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbServerPortStats() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_server_port_stats_port_diameter`: Statistics for the object port\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbServerPortStatsCreate,
		UpdateContext: resourceSlbServerPortStatsUpdate,
		ReadContext:   resourceSlbServerPortStatsRead,
		DeleteContext: resourceSlbServerPortStatsDelete,

		Schema: map[string]*schema.Schema{
			"port_number": {
				Type: schema.TypeInt, Required: true, Description: "Port Number",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'tcp': TCP Port; 'udp': UDP Port;",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_diameter": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"num": {
										Type: schema.TypeInt, Optional: true, Description: "Number",
									},
									"curr": {
										Type: schema.TypeInt, Optional: true, Description: "Current",
									},
									"total": {
										Type: schema.TypeInt, Optional: true, Description: "Total",
									},
									"svrsel_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Number of server selection failed",
									},
									"no_route": {
										Type: schema.TypeInt, Optional: true, Description: "Number of no routes",
									},
									"snat_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Number of snat failures",
									},
									"client_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Number of client failures",
									},
									"server_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Number of server failures",
									},
									"no_sess": {
										Type: schema.TypeInt, Optional: true, Description: "Number of no sessions",
									},
									"user_session": {
										Type: schema.TypeInt, Optional: true, Description: "Number of user sessions",
									},
									"acr_out": {
										Type: schema.TypeInt, Optional: true, Description: "Number of ACRs out",
									},
									"acr_in": {
										Type: schema.TypeInt, Optional: true, Description: "Number of ACRs in",
									},
									"aca_out": {
										Type: schema.TypeInt, Optional: true, Description: "Number of ACAs out",
									},
									"aca_in": {
										Type: schema.TypeInt, Optional: true, Description: "Number of ACAs in",
									},
									"cea_out": {
										Type: schema.TypeInt, Optional: true, Description: "Number of CEAs out",
									},
									"cea_in": {
										Type: schema.TypeInt, Optional: true, Description: "Number of CEAs in",
									},
									"cer_out": {
										Type: schema.TypeInt, Optional: true, Description: "Number of CERs out",
									},
									"cer_in": {
										Type: schema.TypeInt, Optional: true, Description: "Number of CERs in",
									},
									"dwr_out": {
										Type: schema.TypeInt, Optional: true, Description: "Number of DWRs out",
									},
									"dwr_in": {
										Type: schema.TypeInt, Optional: true, Description: "Number of DWRs in",
									},
									"dwa_out": {
										Type: schema.TypeInt, Optional: true, Description: "Number of DWAs out",
									},
									"dwa_in": {
										Type: schema.TypeInt, Optional: true, Description: "Number of DWAs in",
									},
									"str_out": {
										Type: schema.TypeInt, Optional: true, Description: "Number of STRs out",
									},
									"str_in": {
										Type: schema.TypeInt, Optional: true, Description: "Number of STRs in",
									},
									"sta_out": {
										Type: schema.TypeInt, Optional: true, Description: "Number of STAs out",
									},
									"sta_in": {
										Type: schema.TypeInt, Optional: true, Description: "Number of STAs in",
									},
									"asr_out": {
										Type: schema.TypeInt, Optional: true, Description: "Number of ASRs out",
									},
									"asr_in": {
										Type: schema.TypeInt, Optional: true, Description: "Number of ASRs in",
									},
									"asa_out": {
										Type: schema.TypeInt, Optional: true, Description: "Number of ASAs out",
									},
									"asa_in": {
										Type: schema.TypeInt, Optional: true, Description: "Number of ASAs in",
									},
									"other_out": {
										Type: schema.TypeInt, Optional: true, Description: "Number of other messages out",
									},
									"other_in": {
										Type: schema.TypeInt, Optional: true, Description: "Number of other messages in",
									},
									"cca_out": {
										Type: schema.TypeInt, Optional: true, Description: "Number of CCAs out",
									},
									"cca_in": {
										Type: schema.TypeInt, Optional: true, Description: "Number of CCAs in",
									},
									"ccr_out": {
										Type: schema.TypeInt, Optional: true, Description: "Number of CCRs out",
									},
									"ccr_in": {
										Type: schema.TypeInt, Optional: true, Description: "Number of CCRs in",
									},
									"dpr_out": {
										Type: schema.TypeInt, Optional: true, Description: "Number of DPRs out",
									},
									"dpr_in": {
										Type: schema.TypeInt, Optional: true, Description: "Number of DPRs in",
									},
									"dpa_out": {
										Type: schema.TypeInt, Optional: true, Description: "Number of DPAs out",
									},
									"dpa_in": {
										Type: schema.TypeInt, Optional: true, Description: "Number of DPAs in",
									},
								},
							},
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
func resourceSlbServerPortStatsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServerPortStatsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServerPortStats(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbServerPortStatsRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbServerPortStatsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServerPortStatsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServerPortStats(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbServerPortStatsRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbServerPortStatsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServerPortStatsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServerPortStats(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbServerPortStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServerPortStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServerPortStats(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbServerPortStatsStats(d []interface{}) edpt.SlbServerPortStatsStats {

	count1 := len(d)
	var ret edpt.SlbServerPortStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PortDiameter = getObjectSlbServerPortStatsStatsPortDiameter(in["port_diameter"].([]interface{}))
	}
	return ret
}

func getObjectSlbServerPortStatsStatsPortDiameter(d []interface{}) edpt.SlbServerPortStatsStatsPortDiameter {

	count1 := len(d)
	var ret edpt.SlbServerPortStatsStatsPortDiameter
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Num = in["num"].(int)
		ret.Curr = in["curr"].(int)
		ret.Total = in["total"].(int)
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.No_route = in["no_route"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Client_fail = in["client_fail"].(int)
		ret.Server_fail = in["server_fail"].(int)
		ret.No_sess = in["no_sess"].(int)
		ret.User_session = in["user_session"].(int)
		ret.Acr_out = in["acr_out"].(int)
		ret.Acr_in = in["acr_in"].(int)
		ret.Aca_out = in["aca_out"].(int)
		ret.Aca_in = in["aca_in"].(int)
		ret.Cea_out = in["cea_out"].(int)
		ret.Cea_in = in["cea_in"].(int)
		ret.Cer_out = in["cer_out"].(int)
		ret.Cer_in = in["cer_in"].(int)
		ret.Dwr_out = in["dwr_out"].(int)
		ret.Dwr_in = in["dwr_in"].(int)
		ret.Dwa_out = in["dwa_out"].(int)
		ret.Dwa_in = in["dwa_in"].(int)
		ret.Str_out = in["str_out"].(int)
		ret.Str_in = in["str_in"].(int)
		ret.Sta_out = in["sta_out"].(int)
		ret.Sta_in = in["sta_in"].(int)
		ret.Asr_out = in["asr_out"].(int)
		ret.Asr_in = in["asr_in"].(int)
		ret.Asa_out = in["asa_out"].(int)
		ret.Asa_in = in["asa_in"].(int)
		ret.Other_out = in["other_out"].(int)
		ret.Other_in = in["other_in"].(int)
		ret.Cca_out = in["cca_out"].(int)
		ret.Cca_in = in["cca_in"].(int)
		ret.Ccr_out = in["ccr_out"].(int)
		ret.Ccr_in = in["ccr_in"].(int)
		ret.Dpr_out = in["dpr_out"].(int)
		ret.Dpr_in = in["dpr_in"].(int)
		ret.Dpa_out = in["dpa_out"].(int)
		ret.Dpa_in = in["dpa_in"].(int)
	}
	return ret
}

func dataToEndpointSlbServerPortStats(d *schema.ResourceData) edpt.SlbServerPortStats {
	var ret edpt.SlbServerPortStats
	ret.Inst.PortNumber = d.Get("port_number").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Stats = getObjectSlbServerPortStatsStats(d.Get("stats").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
