package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNgWafStatsStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ng_waf_stats_stats`: Statistics for the object stats\n\n__PLACEHOLDER__",
		ReadContext: resourceNgWafStatsStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "ng-waf name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"req_received": {
							Type: schema.TypeInt, Optional: true, Description: "Requests Received",
						},
						"req_forward": {
							Type: schema.TypeInt, Optional: true, Description: "Requests Forwarded",
						},
						"req_blocked": {
							Type: schema.TypeInt, Optional: true, Description: "Requests Blocked",
						},
						"req_marked": {
							Type: schema.TypeInt, Optional: true, Description: "Requests Marked",
						},
						"req_bypass": {
							Type: schema.TypeInt, Optional: true, Description: "Requests Bypassed",
						},
						"error": {
							Type: schema.TypeInt, Optional: true, Description: "Error Happened",
						},
						"useragent": {
							Type: schema.TypeInt, Optional: true, Description: "Attack Tooling",
						},
						"aws_ssrf": {
							Type: schema.TypeInt, Optional: true, Description: "AWS SSRF",
						},
						"backdoor": {
							Type: schema.TypeInt, Optional: true, Description: "Backdoor",
						},
						"cmdexe": {
							Type: schema.TypeInt, Optional: true, Description: "Command Execution",
						},
						"xss": {
							Type: schema.TypeInt, Optional: true, Description: "Cross Site Scripting",
						},
						"traversal": {
							Type: schema.TypeInt, Optional: true, Description: "Directory Traversal",
						},
						"graphql_depth": {
							Type: schema.TypeInt, Optional: true, Description: "GraphQL Max Depth",
						},
						"log4j_jndi": {
							Type: schema.TypeInt, Optional: true, Description: "Log4J JNDI",
						},
						"sqli": {
							Type: schema.TypeInt, Optional: true, Description: "SQL Injection",
						},
						"abnormalpath": {
							Type: schema.TypeInt, Optional: true, Description: "Abnormal Path",
						},
						"bhh": {
							Type: schema.TypeInt, Optional: true, Description: "Bad Hop Headers",
						},
						"block": {
							Type: schema.TypeInt, Optional: true, Description: "Blocked Requests",
						},
						"codei": {
							Type: schema.TypeInt, Optional: true, Description: "Code Injection PHP",
						},
						"datacenter": {
							Type: schema.TypeInt, Optional: true, Description: "Datacenter Traffic",
						},
						"doubleenc": {
							Type: schema.TypeInt, Optional: true, Description: "Double Encoding",
						},
						"duplicate_headers": {
							Type: schema.TypeInt, Optional: true, Description: "Duplicate Header Names",
						},
						"fbrowing": {
							Type: schema.TypeInt, Optional: true, Description: "Forceful Browsing",
						},
						"graphql_ide": {
							Type: schema.TypeInt, Optional: true, Description: "GraphQL IDE",
						},
						"graphql_introspection": {
							Type: schema.TypeInt, Optional: true, Description: "GraphQL Introspection",
						},
						"graphql_unused_variables": {
							Type: schema.TypeInt, Optional: true, Description: "GRAPHQL-UNUSED-VARIABLES",
						},
						"http403": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP 403 Errors",
						},
						"http404": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP 404 Errors",
						},
						"http429": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP 429 Errors",
						},
						"http4xx": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP 4xx Errors",
						},
						"http500": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP 500 Errors",
						},
						"http503": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP 503 Errors",
						},
						"http5xx": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP 5xx Errors",
						},
						"respsplit": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP Response Splitting",
						},
						"notutf8": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid Encoding",
						},
						"json_error": {
							Type: schema.TypeInt, Optional: true, Description: "JSON Encoding Error",
						},
						"malformed": {
							Type: schema.TypeInt, Optional: true, Description: "Malformed Data in the Request Body",
						},
						"sans": {
							Type: schema.TypeInt, Optional: true, Description: "Malicious IP Traffic",
						},
						"sigsci_ip": {
							Type: schema.TypeInt, Optional: true, Description: "Network Effect",
						},
						"no_ctype": {
							Type: schema.TypeInt, Optional: true, Description: "Missing Content-Type Request Header",
						},
						"noua": {
							Type: schema.TypeInt, Optional: true, Description: "No User Agent",
						},
						"nullbyte": {
							Type: schema.TypeInt, Optional: true, Description: "Null Byte",
						},
						"privatefile": {
							Type: schema.TypeInt, Optional: true, Description: "Private File",
						},
						"scanner": {
							Type: schema.TypeInt, Optional: true, Description: "Scanner",
						},
						"impostor": {
							Type: schema.TypeInt, Optional: true, Description: "SearchBot Impostor",
						},
						"site_flagged_ip": {
							Type: schema.TypeInt, Optional: true, Description: "Site Flagged IP",
						},
						"tornode": {
							Type: schema.TypeInt, Optional: true, Description: "Tor Traffic",
						},
						"weaktls": {
							Type: schema.TypeInt, Optional: true, Description: "Weak TLS",
						},
						"xml_error": {
							Type: schema.TypeInt, Optional: true, Description: "XML Encoding Error",
						},
						"other": {
							Type: schema.TypeInt, Optional: true, Description: "Unknown Attack",
						},
						"custom1": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 1",
						},
						"custom2": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 2",
						},
						"custom3": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 3",
						},
						"custom4": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 4",
						},
						"custom5": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 5",
						},
						"custom6": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 6",
						},
						"custom7": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 7",
						},
						"custom8": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 8",
						},
						"custom9": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 9",
						},
						"custom10": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 10",
						},
						"custom11": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 11",
						},
						"custom12": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 12",
						},
						"custom13": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 13",
						},
						"custom14": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 14",
						},
						"custom15": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 15",
						},
						"custom16": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 16",
						},
						"custom17": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 17",
						},
						"custom18": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 18",
						},
						"custom19": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 19",
						},
						"custom20": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 20",
						},
						"custom21": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 21",
						},
						"custom22": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 22",
						},
						"custom23": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 23",
						},
						"custom24": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 24",
						},
						"custom25": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 25",
						},
						"custom26": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 26",
						},
						"custom27": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 27",
						},
						"custom28": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 28",
						},
						"custom29": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 29",
						},
						"custom30": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 30",
						},
						"custom31": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 31",
						},
						"custom32": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 32",
						},
						"custom33": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 33",
						},
						"custom34": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 34",
						},
						"custom35": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 35",
						},
						"custom36": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 36",
						},
						"custom37": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 37",
						},
						"custom38": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 38",
						},
						"custom39": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 39",
						},
						"custom40": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 40",
						},
						"custom41": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 41",
						},
						"custom42": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 42",
						},
						"custom43": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 43",
						},
						"custom44": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 44",
						},
						"custom45": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 45",
						},
						"custom46": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 46",
						},
						"custom47": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 47",
						},
						"custom48": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 48",
						},
						"custom49": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 49",
						},
						"custom50": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 50",
						},
						"custom51": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 51",
						},
						"custom52": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 52",
						},
						"custom53": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 53",
						},
						"custom54": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 54",
						},
						"custom55": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 55",
						},
						"custom56": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 56",
						},
						"custom57": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 57",
						},
						"custom58": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 58",
						},
						"custom59": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 59",
						},
						"custom60": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 60",
						},
						"custom61": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 61",
						},
						"custom62": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 62",
						},
						"custom63": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 63",
						},
						"custom64": {
							Type: schema.TypeInt, Optional: true, Description: "Custom signal 64",
						},
					},
				},
			},
		},
	}
}

func resourceNgWafStatsStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNgWafStatsStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNgWafStatsStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NgWafStatsStatsStats := setObjectNgWafStatsStatsStats(res)
		d.Set("stats", NgWafStatsStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNgWafStatsStatsStats(ret edpt.DataNgWafStatsStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"req_received":             ret.DtNgWafStatsStats.Stats.ReqReceived,
			"req_forward":              ret.DtNgWafStatsStats.Stats.ReqForward,
			"req_blocked":              ret.DtNgWafStatsStats.Stats.ReqBlocked,
			"req_marked":               ret.DtNgWafStatsStats.Stats.ReqMarked,
			"req_bypass":               ret.DtNgWafStatsStats.Stats.ReqBypass,
			"error":                    ret.DtNgWafStatsStats.Stats.Error,
			"useragent":                ret.DtNgWafStatsStats.Stats.Useragent,
			"aws_ssrf":                 ret.DtNgWafStatsStats.Stats.AwsSsrf,
			"backdoor":                 ret.DtNgWafStatsStats.Stats.Backdoor,
			"cmdexe":                   ret.DtNgWafStatsStats.Stats.Cmdexe,
			"xss":                      ret.DtNgWafStatsStats.Stats.Xss,
			"traversal":                ret.DtNgWafStatsStats.Stats.Traversal,
			"graphql_depth":            ret.DtNgWafStatsStats.Stats.GraphqlDepth,
			"log4j_jndi":               ret.DtNgWafStatsStats.Stats.Log4jJndi,
			"sqli":                     ret.DtNgWafStatsStats.Stats.Sqli,
			"abnormalpath":             ret.DtNgWafStatsStats.Stats.Abnormalpath,
			"bhh":                      ret.DtNgWafStatsStats.Stats.Bhh,
			"block":                    ret.DtNgWafStatsStats.Stats.Block,
			"codei":                    ret.DtNgWafStatsStats.Stats.Codei,
			"datacenter":               ret.DtNgWafStatsStats.Stats.Datacenter,
			"doubleenc":                ret.DtNgWafStatsStats.Stats.Doubleenc,
			"duplicate_headers":        ret.DtNgWafStatsStats.Stats.DuplicateHeaders,
			"fbrowing":                 ret.DtNgWafStatsStats.Stats.Fbrowing,
			"graphql_ide":              ret.DtNgWafStatsStats.Stats.GraphqlIde,
			"graphql_introspection":    ret.DtNgWafStatsStats.Stats.GraphqlIntrospection,
			"graphql_unused_variables": ret.DtNgWafStatsStats.Stats.GraphqlUnusedVariables,
			"http403":                  ret.DtNgWafStatsStats.Stats.Http403,
			"http404":                  ret.DtNgWafStatsStats.Stats.Http404,
			"http429":                  ret.DtNgWafStatsStats.Stats.Http429,
			"http4xx":                  ret.DtNgWafStatsStats.Stats.Http4xx,
			"http500":                  ret.DtNgWafStatsStats.Stats.Http500,
			"http503":                  ret.DtNgWafStatsStats.Stats.Http503,
			"http5xx":                  ret.DtNgWafStatsStats.Stats.Http5xx,
			"respsplit":                ret.DtNgWafStatsStats.Stats.Respsplit,
			"notutf8":                  ret.DtNgWafStatsStats.Stats.Notutf8,
			"json_error":               ret.DtNgWafStatsStats.Stats.JsonError,
			"malformed":                ret.DtNgWafStatsStats.Stats.Malformed,
			"sans":                     ret.DtNgWafStatsStats.Stats.Sans,
			"sigsci_ip":                ret.DtNgWafStatsStats.Stats.SigsciIp,
			"no_ctype":                 ret.DtNgWafStatsStats.Stats.NoCtype,
			"noua":                     ret.DtNgWafStatsStats.Stats.Noua,
			"nullbyte":                 ret.DtNgWafStatsStats.Stats.Nullbyte,
			"privatefile":              ret.DtNgWafStatsStats.Stats.Privatefile,
			"scanner":                  ret.DtNgWafStatsStats.Stats.Scanner,
			"impostor":                 ret.DtNgWafStatsStats.Stats.Impostor,
			"site_flagged_ip":          ret.DtNgWafStatsStats.Stats.SiteFlaggedIp,
			"tornode":                  ret.DtNgWafStatsStats.Stats.Tornode,
			"weaktls":                  ret.DtNgWafStatsStats.Stats.Weaktls,
			"xml_error":                ret.DtNgWafStatsStats.Stats.XmlError,
			"other":                    ret.DtNgWafStatsStats.Stats.Other,
			"custom1":                  ret.DtNgWafStatsStats.Stats.Custom1,
			"custom2":                  ret.DtNgWafStatsStats.Stats.Custom2,
			"custom3":                  ret.DtNgWafStatsStats.Stats.Custom3,
			"custom4":                  ret.DtNgWafStatsStats.Stats.Custom4,
			"custom5":                  ret.DtNgWafStatsStats.Stats.Custom5,
			"custom6":                  ret.DtNgWafStatsStats.Stats.Custom6,
			"custom7":                  ret.DtNgWafStatsStats.Stats.Custom7,
			"custom8":                  ret.DtNgWafStatsStats.Stats.Custom8,
			"custom9":                  ret.DtNgWafStatsStats.Stats.Custom9,
			"custom10":                 ret.DtNgWafStatsStats.Stats.Custom10,
			"custom11":                 ret.DtNgWafStatsStats.Stats.Custom11,
			"custom12":                 ret.DtNgWafStatsStats.Stats.Custom12,
			"custom13":                 ret.DtNgWafStatsStats.Stats.Custom13,
			"custom14":                 ret.DtNgWafStatsStats.Stats.Custom14,
			"custom15":                 ret.DtNgWafStatsStats.Stats.Custom15,
			"custom16":                 ret.DtNgWafStatsStats.Stats.Custom16,
			"custom17":                 ret.DtNgWafStatsStats.Stats.Custom17,
			"custom18":                 ret.DtNgWafStatsStats.Stats.Custom18,
			"custom19":                 ret.DtNgWafStatsStats.Stats.Custom19,
			"custom20":                 ret.DtNgWafStatsStats.Stats.Custom20,
			"custom21":                 ret.DtNgWafStatsStats.Stats.Custom21,
			"custom22":                 ret.DtNgWafStatsStats.Stats.Custom22,
			"custom23":                 ret.DtNgWafStatsStats.Stats.Custom23,
			"custom24":                 ret.DtNgWafStatsStats.Stats.Custom24,
			"custom25":                 ret.DtNgWafStatsStats.Stats.Custom25,
			"custom26":                 ret.DtNgWafStatsStats.Stats.Custom26,
			"custom27":                 ret.DtNgWafStatsStats.Stats.Custom27,
			"custom28":                 ret.DtNgWafStatsStats.Stats.Custom28,
			"custom29":                 ret.DtNgWafStatsStats.Stats.Custom29,
			"custom30":                 ret.DtNgWafStatsStats.Stats.Custom30,
			"custom31":                 ret.DtNgWafStatsStats.Stats.Custom31,
			"custom32":                 ret.DtNgWafStatsStats.Stats.Custom32,
			"custom33":                 ret.DtNgWafStatsStats.Stats.Custom33,
			"custom34":                 ret.DtNgWafStatsStats.Stats.Custom34,
			"custom35":                 ret.DtNgWafStatsStats.Stats.Custom35,
			"custom36":                 ret.DtNgWafStatsStats.Stats.Custom36,
			"custom37":                 ret.DtNgWafStatsStats.Stats.Custom37,
			"custom38":                 ret.DtNgWafStatsStats.Stats.Custom38,
			"custom39":                 ret.DtNgWafStatsStats.Stats.Custom39,
			"custom40":                 ret.DtNgWafStatsStats.Stats.Custom40,
			"custom41":                 ret.DtNgWafStatsStats.Stats.Custom41,
			"custom42":                 ret.DtNgWafStatsStats.Stats.Custom42,
			"custom43":                 ret.DtNgWafStatsStats.Stats.Custom43,
			"custom44":                 ret.DtNgWafStatsStats.Stats.Custom44,
			"custom45":                 ret.DtNgWafStatsStats.Stats.Custom45,
			"custom46":                 ret.DtNgWafStatsStats.Stats.Custom46,
			"custom47":                 ret.DtNgWafStatsStats.Stats.Custom47,
			"custom48":                 ret.DtNgWafStatsStats.Stats.Custom48,
			"custom49":                 ret.DtNgWafStatsStats.Stats.Custom49,
			"custom50":                 ret.DtNgWafStatsStats.Stats.Custom50,
			"custom51":                 ret.DtNgWafStatsStats.Stats.Custom51,
			"custom52":                 ret.DtNgWafStatsStats.Stats.Custom52,
			"custom53":                 ret.DtNgWafStatsStats.Stats.Custom53,
			"custom54":                 ret.DtNgWafStatsStats.Stats.Custom54,
			"custom55":                 ret.DtNgWafStatsStats.Stats.Custom55,
			"custom56":                 ret.DtNgWafStatsStats.Stats.Custom56,
			"custom57":                 ret.DtNgWafStatsStats.Stats.Custom57,
			"custom58":                 ret.DtNgWafStatsStats.Stats.Custom58,
			"custom59":                 ret.DtNgWafStatsStats.Stats.Custom59,
			"custom60":                 ret.DtNgWafStatsStats.Stats.Custom60,
			"custom61":                 ret.DtNgWafStatsStats.Stats.Custom61,
			"custom62":                 ret.DtNgWafStatsStats.Stats.Custom62,
			"custom63":                 ret.DtNgWafStatsStats.Stats.Custom63,
			"custom64":                 ret.DtNgWafStatsStats.Stats.Custom64,
		},
	}
}

func getObjectNgWafStatsStatsStats(d []interface{}) edpt.NgWafStatsStatsStats {

	count1 := len(d)
	var ret edpt.NgWafStatsStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ReqReceived = in["req_received"].(int)
		ret.ReqForward = in["req_forward"].(int)
		ret.ReqBlocked = in["req_blocked"].(int)
		ret.ReqMarked = in["req_marked"].(int)
		ret.ReqBypass = in["req_bypass"].(int)
		ret.Error = in["error"].(int)
		ret.Useragent = in["useragent"].(int)
		ret.AwsSsrf = in["aws_ssrf"].(int)
		ret.Backdoor = in["backdoor"].(int)
		ret.Cmdexe = in["cmdexe"].(int)
		ret.Xss = in["xss"].(int)
		ret.Traversal = in["traversal"].(int)
		ret.GraphqlDepth = in["graphql_depth"].(int)
		ret.Log4jJndi = in["log4j_jndi"].(int)
		ret.Sqli = in["sqli"].(int)
		ret.Abnormalpath = in["abnormalpath"].(int)
		ret.Bhh = in["bhh"].(int)
		ret.Block = in["block"].(int)
		ret.Codei = in["codei"].(int)
		ret.Datacenter = in["datacenter"].(int)
		ret.Doubleenc = in["doubleenc"].(int)
		ret.DuplicateHeaders = in["duplicate_headers"].(int)
		ret.Fbrowing = in["fbrowing"].(int)
		ret.GraphqlIde = in["graphql_ide"].(int)
		ret.GraphqlIntrospection = in["graphql_introspection"].(int)
		ret.GraphqlUnusedVariables = in["graphql_unused_variables"].(int)
		ret.Http403 = in["http403"].(int)
		ret.Http404 = in["http404"].(int)
		ret.Http429 = in["http429"].(int)
		ret.Http4xx = in["http4xx"].(int)
		ret.Http500 = in["http500"].(int)
		ret.Http503 = in["http503"].(int)
		ret.Http5xx = in["http5xx"].(int)
		ret.Respsplit = in["respsplit"].(int)
		ret.Notutf8 = in["notutf8"].(int)
		ret.JsonError = in["json_error"].(int)
		ret.Malformed = in["malformed"].(int)
		ret.Sans = in["sans"].(int)
		ret.SigsciIp = in["sigsci_ip"].(int)
		ret.NoCtype = in["no_ctype"].(int)
		ret.Noua = in["noua"].(int)
		ret.Nullbyte = in["nullbyte"].(int)
		ret.Privatefile = in["privatefile"].(int)
		ret.Scanner = in["scanner"].(int)
		ret.Impostor = in["impostor"].(int)
		ret.SiteFlaggedIp = in["site_flagged_ip"].(int)
		ret.Tornode = in["tornode"].(int)
		ret.Weaktls = in["weaktls"].(int)
		ret.XmlError = in["xml_error"].(int)
		ret.Other = in["other"].(int)
		ret.Custom1 = in["custom1"].(int)
		ret.Custom2 = in["custom2"].(int)
		ret.Custom3 = in["custom3"].(int)
		ret.Custom4 = in["custom4"].(int)
		ret.Custom5 = in["custom5"].(int)
		ret.Custom6 = in["custom6"].(int)
		ret.Custom7 = in["custom7"].(int)
		ret.Custom8 = in["custom8"].(int)
		ret.Custom9 = in["custom9"].(int)
		ret.Custom10 = in["custom10"].(int)
		ret.Custom11 = in["custom11"].(int)
		ret.Custom12 = in["custom12"].(int)
		ret.Custom13 = in["custom13"].(int)
		ret.Custom14 = in["custom14"].(int)
		ret.Custom15 = in["custom15"].(int)
		ret.Custom16 = in["custom16"].(int)
		ret.Custom17 = in["custom17"].(int)
		ret.Custom18 = in["custom18"].(int)
		ret.Custom19 = in["custom19"].(int)
		ret.Custom20 = in["custom20"].(int)
		ret.Custom21 = in["custom21"].(int)
		ret.Custom22 = in["custom22"].(int)
		ret.Custom23 = in["custom23"].(int)
		ret.Custom24 = in["custom24"].(int)
		ret.Custom25 = in["custom25"].(int)
		ret.Custom26 = in["custom26"].(int)
		ret.Custom27 = in["custom27"].(int)
		ret.Custom28 = in["custom28"].(int)
		ret.Custom29 = in["custom29"].(int)
		ret.Custom30 = in["custom30"].(int)
		ret.Custom31 = in["custom31"].(int)
		ret.Custom32 = in["custom32"].(int)
		ret.Custom33 = in["custom33"].(int)
		ret.Custom34 = in["custom34"].(int)
		ret.Custom35 = in["custom35"].(int)
		ret.Custom36 = in["custom36"].(int)
		ret.Custom37 = in["custom37"].(int)
		ret.Custom38 = in["custom38"].(int)
		ret.Custom39 = in["custom39"].(int)
		ret.Custom40 = in["custom40"].(int)
		ret.Custom41 = in["custom41"].(int)
		ret.Custom42 = in["custom42"].(int)
		ret.Custom43 = in["custom43"].(int)
		ret.Custom44 = in["custom44"].(int)
		ret.Custom45 = in["custom45"].(int)
		ret.Custom46 = in["custom46"].(int)
		ret.Custom47 = in["custom47"].(int)
		ret.Custom48 = in["custom48"].(int)
		ret.Custom49 = in["custom49"].(int)
		ret.Custom50 = in["custom50"].(int)
		ret.Custom51 = in["custom51"].(int)
		ret.Custom52 = in["custom52"].(int)
		ret.Custom53 = in["custom53"].(int)
		ret.Custom54 = in["custom54"].(int)
		ret.Custom55 = in["custom55"].(int)
		ret.Custom56 = in["custom56"].(int)
		ret.Custom57 = in["custom57"].(int)
		ret.Custom58 = in["custom58"].(int)
		ret.Custom59 = in["custom59"].(int)
		ret.Custom60 = in["custom60"].(int)
		ret.Custom61 = in["custom61"].(int)
		ret.Custom62 = in["custom62"].(int)
		ret.Custom63 = in["custom63"].(int)
		ret.Custom64 = in["custom64"].(int)
	}
	return ret
}

func dataToEndpointNgWafStatsStats(d *schema.ResourceData) edpt.NgWafStatsStats {
	var ret edpt.NgWafStatsStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectNgWafStatsStatsStats(d.Get("stats").([]interface{}))
	return ret
}
