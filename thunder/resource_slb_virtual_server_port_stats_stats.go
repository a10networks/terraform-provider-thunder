package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbVirtualServerPortStats55() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_virtual_server_port_stats_stats`: Statistics for the object port\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbVirtualServerPortStats55Read,

		Schema: map[string]*schema.Schema{
			"port_number": {
				Type: schema.TypeInt, Required: true, Description: "Port",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'tcp': TCP LB service; 'udp': UDP Port; 'others': for no tcp/udp protocol, do IP load balancing; 'diameter': diameter port; 'dns-tcp': DNS service over TCP; 'dns-udp': DNS service over UDP; 'fast-http': Fast HTTP Port; 'fix': FIX Port; 'ftp': File Transfer Protocol Port; 'ftp-proxy': ftp proxy port; 'http': HTTP Port; 'https': HTTPS port; 'imap': imap proxy port; 'mlb': Message based load balancing; 'mms': Microsoft Multimedia Service Port; 'mysql': mssql port; 'mssql': mssql; 'pop3': pop3 proxy port; 'radius': RADIUS Port; 'rtsp': Real Time Streaming Protocol Port; 'sip': Session initiation protocol over UDP; 'sip-tcp': Session initiation protocol over TCP; 'sips': Session initiation protocol over TLS; 'smpp-tcp': SMPP service over TCP; 'spdy': spdy port; 'spdys': spdys port; 'smtp': SMTP Port; 'mqtt': MQTT Port; 'mqtts': MQTTS Port; 'ssl-proxy': Generic SSL proxy; 'ssli': SSL insight; 'ssh': SSH Port; 'tcp-proxy': Generic TCP proxy; 'tftp': TFTP Port; 'fast-fix': Fast FIX port; 'http-over-quic': HTTP3-over-quic port;",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceSlbVirtualServerPortStats55Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats55Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats55(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbVirtualServerPortStats55Stats := setObjectSlbVirtualServerPortStats55Stats(res)
		d.Set("stats", SlbVirtualServerPortStats55Stats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbVirtualServerPortStats55Stats(ret edpt.DataSlbVirtualServerPortStats55) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"stats": setObjectSlbVirtualServerPortStats55StatsStats(ret.DtSlbVirtualServerPortStats55.Stats.Stats),
		},
	}
}

func setObjectSlbVirtualServerPortStats55StatsStats(d edpt.SlbVirtualServerPortStats55StatsStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["req_received"] = d.ReqReceived

	in["req_forward"] = d.ReqForward

	in["req_blocked"] = d.ReqBlocked

	in["req_marked"] = d.ReqMarked

	in["req_bypass"] = d.ReqBypass

	in["error"] = d.Error

	in["useragent"] = d.Useragent

	in["aws_ssrf"] = d.AwsSsrf

	in["backdoor"] = d.Backdoor

	in["cmdexe"] = d.Cmdexe

	in["xss"] = d.Xss

	in["traversal"] = d.Traversal

	in["graphql_depth"] = d.GraphqlDepth

	in["log4j_jndi"] = d.Log4jJndi

	in["sqli"] = d.Sqli

	in["abnormalpath"] = d.Abnormalpath

	in["bhh"] = d.Bhh

	in["block"] = d.Block

	in["codei"] = d.Codei

	in["datacenter"] = d.Datacenter

	in["doubleenc"] = d.Doubleenc

	in["duplicate_headers"] = d.DuplicateHeaders

	in["fbrowing"] = d.Fbrowing

	in["graphql_ide"] = d.GraphqlIde

	in["graphql_introspection"] = d.GraphqlIntrospection

	in["graphql_unused_variables"] = d.GraphqlUnusedVariables

	in["http403"] = d.Http403

	in["http404"] = d.Http404

	in["http429"] = d.Http429

	in["http4xx"] = d.Http4xx

	in["http500"] = d.Http500

	in["http503"] = d.Http503

	in["http5xx"] = d.Http5xx

	in["respsplit"] = d.Respsplit

	in["notutf8"] = d.Notutf8

	in["json_error"] = d.JsonError

	in["malformed"] = d.Malformed

	in["sans"] = d.Sans

	in["sigsci_ip"] = d.SigsciIp

	in["no_ctype"] = d.NoCtype

	in["noua"] = d.Noua

	in["nullbyte"] = d.Nullbyte

	in["privatefile"] = d.Privatefile

	in["scanner"] = d.Scanner

	in["impostor"] = d.Impostor

	in["site_flagged_ip"] = d.SiteFlaggedIp

	in["tornode"] = d.Tornode

	in["weaktls"] = d.Weaktls

	in["xml_error"] = d.XmlError

	in["other"] = d.Other

	in["custom1"] = d.Custom1

	in["custom2"] = d.Custom2

	in["custom3"] = d.Custom3

	in["custom4"] = d.Custom4

	in["custom5"] = d.Custom5

	in["custom6"] = d.Custom6

	in["custom7"] = d.Custom7

	in["custom8"] = d.Custom8

	in["custom9"] = d.Custom9

	in["custom10"] = d.Custom10

	in["custom11"] = d.Custom11

	in["custom12"] = d.Custom12

	in["custom13"] = d.Custom13

	in["custom14"] = d.Custom14

	in["custom15"] = d.Custom15

	in["custom16"] = d.Custom16

	in["custom17"] = d.Custom17

	in["custom18"] = d.Custom18

	in["custom19"] = d.Custom19

	in["custom20"] = d.Custom20

	in["custom21"] = d.Custom21

	in["custom22"] = d.Custom22

	in["custom23"] = d.Custom23

	in["custom24"] = d.Custom24

	in["custom25"] = d.Custom25

	in["custom26"] = d.Custom26

	in["custom27"] = d.Custom27

	in["custom28"] = d.Custom28

	in["custom29"] = d.Custom29

	in["custom30"] = d.Custom30

	in["custom31"] = d.Custom31

	in["custom32"] = d.Custom32

	in["custom33"] = d.Custom33

	in["custom34"] = d.Custom34

	in["custom35"] = d.Custom35

	in["custom36"] = d.Custom36

	in["custom37"] = d.Custom37

	in["custom38"] = d.Custom38

	in["custom39"] = d.Custom39

	in["custom40"] = d.Custom40

	in["custom41"] = d.Custom41

	in["custom42"] = d.Custom42

	in["custom43"] = d.Custom43

	in["custom44"] = d.Custom44

	in["custom45"] = d.Custom45

	in["custom46"] = d.Custom46

	in["custom47"] = d.Custom47

	in["custom48"] = d.Custom48

	in["custom49"] = d.Custom49

	in["custom50"] = d.Custom50

	in["custom51"] = d.Custom51

	in["custom52"] = d.Custom52

	in["custom53"] = d.Custom53

	in["custom54"] = d.Custom54

	in["custom55"] = d.Custom55

	in["custom56"] = d.Custom56

	in["custom57"] = d.Custom57

	in["custom58"] = d.Custom58

	in["custom59"] = d.Custom59

	in["custom60"] = d.Custom60

	in["custom61"] = d.Custom61

	in["custom62"] = d.Custom62

	in["custom63"] = d.Custom63

	in["custom64"] = d.Custom64
	result = append(result, in)
	return result
}

func getObjectSlbVirtualServerPortStats55Stats(d []interface{}) edpt.SlbVirtualServerPortStats55Stats {

	count1 := len(d)
	var ret edpt.SlbVirtualServerPortStats55Stats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Stats = getObjectSlbVirtualServerPortStats55StatsStats(in["stats"].([]interface{}))
	}
	return ret
}

func getObjectSlbVirtualServerPortStats55StatsStats(d []interface{}) edpt.SlbVirtualServerPortStats55StatsStats {

	count1 := len(d)
	var ret edpt.SlbVirtualServerPortStats55StatsStats
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

func dataToEndpointSlbVirtualServerPortStats55(d *schema.ResourceData) edpt.SlbVirtualServerPortStats55 {
	var ret edpt.SlbVirtualServerPortStats55

	ret.PortNumber = d.Get("port_number").(int)

	ret.Protocol = d.Get("protocol").(string)

	ret.Stats = getObjectSlbVirtualServerPortStats55Stats(d.Get("stats").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
