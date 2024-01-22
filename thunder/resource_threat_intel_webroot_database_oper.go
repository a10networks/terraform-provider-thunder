package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceThreatIntelWebrootDatabaseOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_threat_intel_webroot_database_oper`: Operational Status for the object webroot-database\n\n__PLACEHOLDER__",
		ReadContext: resourceThreatIntelWebrootDatabaseOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"size": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"version": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"last_update_time": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"next_update_time": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"connection_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"failure_reason": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"last_successful_connection": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"spam_sources": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"windows_exploits": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"web_attacks": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"botnets": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"scanners": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dos_attacks": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"reputation": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"phishing": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"proxy": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"mobile_threats": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"tor_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_entries": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceThreatIntelWebrootDatabaseOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceThreatIntelWebrootDatabaseOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointThreatIntelWebrootDatabaseOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ThreatIntelWebrootDatabaseOperOper := setObjectThreatIntelWebrootDatabaseOperOper(res)
		d.Set("oper", ThreatIntelWebrootDatabaseOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectThreatIntelWebrootDatabaseOperOper(ret edpt.DataThreatIntelWebrootDatabaseOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"name":                       ret.DtThreatIntelWebrootDatabaseOper.Oper.Name,
			"status":                     ret.DtThreatIntelWebrootDatabaseOper.Oper.Status,
			"size":                       ret.DtThreatIntelWebrootDatabaseOper.Oper.Size,
			"version":                    ret.DtThreatIntelWebrootDatabaseOper.Oper.Version,
			"last_update_time":           ret.DtThreatIntelWebrootDatabaseOper.Oper.LastUpdateTime,
			"next_update_time":           ret.DtThreatIntelWebrootDatabaseOper.Oper.NextUpdateTime,
			"connection_status":          ret.DtThreatIntelWebrootDatabaseOper.Oper.ConnectionStatus,
			"failure_reason":             ret.DtThreatIntelWebrootDatabaseOper.Oper.FailureReason,
			"last_successful_connection": ret.DtThreatIntelWebrootDatabaseOper.Oper.LastSuccessfulConnection,
			"spam_sources":               ret.DtThreatIntelWebrootDatabaseOper.Oper.SpamSources,
			"windows_exploits":           ret.DtThreatIntelWebrootDatabaseOper.Oper.WindowsExploits,
			"web_attacks":                ret.DtThreatIntelWebrootDatabaseOper.Oper.WebAttacks,
			"botnets":                    ret.DtThreatIntelWebrootDatabaseOper.Oper.Botnets,
			"scanners":                   ret.DtThreatIntelWebrootDatabaseOper.Oper.Scanners,
			"dos_attacks":                ret.DtThreatIntelWebrootDatabaseOper.Oper.DosAttacks,
			"reputation":                 ret.DtThreatIntelWebrootDatabaseOper.Oper.Reputation,
			"phishing":                   ret.DtThreatIntelWebrootDatabaseOper.Oper.Phishing,
			"proxy":                      ret.DtThreatIntelWebrootDatabaseOper.Oper.Proxy,
			"mobile_threats":             ret.DtThreatIntelWebrootDatabaseOper.Oper.MobileThreats,
			"tor_proxy":                  ret.DtThreatIntelWebrootDatabaseOper.Oper.TorProxy,
			"total_entries":              ret.DtThreatIntelWebrootDatabaseOper.Oper.TotalEntries,
		},
	}
}

func getObjectThreatIntelWebrootDatabaseOperOper(d []interface{}) edpt.ThreatIntelWebrootDatabaseOperOper {

	count1 := len(d)
	var ret edpt.ThreatIntelWebrootDatabaseOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
		ret.Status = in["status"].(string)
		ret.Size = in["size"].(string)
		ret.Version = in["version"].(int)
		ret.LastUpdateTime = in["last_update_time"].(string)
		ret.NextUpdateTime = in["next_update_time"].(string)
		ret.ConnectionStatus = in["connection_status"].(string)
		ret.FailureReason = in["failure_reason"].(string)
		ret.LastSuccessfulConnection = in["last_successful_connection"].(string)
		ret.SpamSources = in["spam_sources"].(int)
		ret.WindowsExploits = in["windows_exploits"].(int)
		ret.WebAttacks = in["web_attacks"].(int)
		ret.Botnets = in["botnets"].(int)
		ret.Scanners = in["scanners"].(int)
		ret.DosAttacks = in["dos_attacks"].(int)
		ret.Reputation = in["reputation"].(int)
		ret.Phishing = in["phishing"].(int)
		ret.Proxy = in["proxy"].(int)
		ret.MobileThreats = in["mobile_threats"].(int)
		ret.TorProxy = in["tor_proxy"].(int)
		ret.TotalEntries = in["total_entries"].(int)
	}
	return ret
}

func dataToEndpointThreatIntelWebrootDatabaseOper(d *schema.ResourceData) edpt.ThreatIntelWebrootDatabaseOper {
	var ret edpt.ThreatIntelWebrootDatabaseOper

	ret.Oper = getObjectThreatIntelWebrootDatabaseOperOper(d.Get("oper").([]interface{}))
	return ret
}
