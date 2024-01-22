package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemRadiusServer() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_radius_server`: Configure system as a RADIUS server\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemRadiusServerCreate,
		UpdateContext: resourceSystemRadiusServerUpdate,
		ReadContext:   resourceSystemRadiusServerRead,
		DeleteContext: resourceSystemRadiusServerDelete,

		Schema: map[string]*schema.Schema{
			"accounting_interim_update": {
				Type: schema.TypeString, Optional: true, Default: "ignore", Description: "'ignore': Ignore (default); 'append-entry': Append the AVPs to existing entry; 'replace-entry': Replace the AVPs of existing entry;",
			},
			"accounting_on": {
				Type: schema.TypeString, Optional: true, Default: "ignore", Description: "'ignore': Ignore (default); 'delete-entries-using-attribute': Delete entries matching attribute in RADIUS Table;",
			},
			"accounting_start": {
				Type: schema.TypeString, Optional: true, Default: "append-entry", Description: "'ignore': Ignore; 'append-entry': Append the AVPs to existing entry (default); 'replace-entry': Replace the AVPs of existing entry;",
			},
			"accounting_stop": {
				Type: schema.TypeString, Optional: true, Default: "delete-entry", Description: "'ignore': Ignore; 'delete-entry': Delete the entry (default); 'delete-entry-and-sessions': Delete the entry and data sessions associated(CGN only);",
			},
			"attribute": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attribute_value": {
							Type: schema.TypeString, Optional: true, Description: "'inside-ipv6-prefix': Framed IPv6 Prefix; 'inside-ip': Inside IP address; 'inside-ipv6': Inside IPv6 address; 'imei': International Mobile Equipment Identity (IMEI); 'imsi': International Mobile Subscriber Identity (IMSI); 'msisdn': Mobile Subscriber Integrated Services Digital Network-Number (MSISDN); 'custom1': Customized attribute 1; 'custom2': Customized attribute 2; 'custom3': Customized attribute 3; 'custom4': Customized attribute 4; 'custom5': Customized attribute 5; 'custom6': Customized attribute 6;",
						},
						"prefix_length": {
							Type: schema.TypeString, Optional: true, Description: "'32': Prefix length 32; '48': Prefix length 48; '64': Prefix length 64; '80': Prefix length 80; '96': Prefix length 96; '112': Prefix length 112;",
						},
						"prefix_vendor": {
							Type: schema.TypeInt, Optional: true, Description: "RADIUS vendor attribute information (RADIUS vendor ID)",
						},
						"prefix_number": {
							Type: schema.TypeInt, Optional: true, Description: "RADIUS attribute number",
						},
						"name": {
							Type: schema.TypeString, Optional: true, Description: "Customized attribute name",
						},
						"value": {
							Type: schema.TypeString, Optional: true, Description: "'hexadecimal': Type of attribute value is hexadecimal;",
						},
						"custom_vendor": {
							Type: schema.TypeInt, Optional: true, Description: "RADIUS vendor attribute information (RADIUS vendor ID)",
						},
						"custom_number": {
							Type: schema.TypeInt, Optional: true, Description: "RADIUS attribute number",
						},
						"vendor": {
							Type: schema.TypeInt, Optional: true, Description: "RADIUS vendor attribute information (RADIUS vendor ID)",
						},
						"number": {
							Type: schema.TypeInt, Optional: true, Description: "RADIUS attribute number",
						},
					},
				},
			},
			"attribute_name": {
				Type: schema.TypeString, Optional: true, Description: "'msisdn': Clear using MSISDN; 'imei': Clear using IMEI; 'imsi': Clear using IMSI; 'custom1': Clear using CUSTOM1 attribute configured; 'custom2': Clear using CUSTOM2 attribute configured; 'custom3': Clear using CUSTOM3 attribute configured; 'custom4': Clear using CUSTOM4 attribute configured; 'custom5': Clear using CUSTOM5 attribute configured; 'custom6': Clear using CUSTOM6 attribute configured;",
			},
			"custom_attribute_name": {
				Type: schema.TypeString, Optional: true, Description: "Clear using customized attribute",
			},
			"disable_reply": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Toggle option for RADIUS reply packet(Default: Accounting response will be sent)",
			},
			"listen_port": {
				Type: schema.TypeInt, Optional: true, Default: 1813, Description: "Configure the listen port of RADIUS server (default 1813) (Port number)",
			},
			"remote": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_list_name": {
										Type: schema.TypeString, Optional: true, Description: "IP-list name",
									},
									"ip_list_secret": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure shared secret",
									},
									"ip_list_secret_string": {
										Type: schema.TypeString, Optional: true, Description: "The RADIUS secret",
									},
								},
							},
						},
					},
				},
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'msisdn-received': MSISDN Received; 'imei-received': IMEI Received; 'imsi-received': IMSI Received; 'custom-received': Custom attribute Received; 'radius-request-received': RADIUS Request Received; 'radius-request-dropped': RADIUS Request Dropped (Malformed Packet); 'request-bad-secret-dropped': RADIUS Request Bad Secret Dropped; 'request-no-key-vap-dropped': RADIUS Request No Key Attribute Dropped; 'request-malformed-dropped': RADIUS Request Malformed Dropped; 'request-ignored': RADIUS Request Ignored; 'radius-table-full': RADIUS Request Dropped (Table Full); 'secret-not-configured-dropped': RADIUS Secret Not Configured Dropped; 'ha-standby-dropped': HA Standby Dropped; 'ipv6-prefix-length-mismatch': Framed IPV6 Prefix Length Mismatch; 'invalid-key': Radius Request has Invalid Key Field; 'smp-created': RADIUS SMP Created; 'smp-deleted': RADIUS SMP Deleted; 'smp-mem-allocated': RADIUS SMP Memory Allocated; 'smp-mem-alloc-failed': RADIUS SMP Memory Allocation Failed; 'smp-mem-freed': RADIUS SMP Memory Freed; 'smp-in-rml': RADIUS SMP in RML; 'mem-allocated': RADIUS Memory Allocated; 'mem-alloc-failed': RADIUS Memory Allocation Failed; 'mem-freed': RADIUS Memory Freed; 'ha-sync-create-sent': HA Record Sync Create Sent; 'ha-sync-delete-sent': HA Record Sync Delete Sent; 'ha-sync-create-recv': HA Record Sync Create Received; 'ha-sync-delete-recv': HA Record Sync Delete Received; 'acct-on-filters-full': RADIUS Acct On Request Ignored(Filters Full); 'acct-on-dup-request': Duplicate RADIUS Acct On Request; 'ip-mismatch-delete': Radius Entry IP Mismatch Delete; 'ip-add-race-drop': Radius Entry IP Add Race Drop; 'ha-sync-no-key-vap-dropped': HA Record Sync No key dropped; 'inter-card-msg-fail-drop': Inter-Card Message Fail Drop; 'radius-packets-redirected': RADIUS packets redirected (SO); 'radius-packets-redirect-fail-dropped': RADIUS packets dropped due to redirect failure (SO); 'radius-packets-process-local': RADIUS packets processed locally without redirection (SO); 'radius-packets-dropped-not-lo': RADIUS packets dropped dest not loopback (SO); 'radius-inter-card-dup-redir': RADIUS packet dropped as redirected by other blade (SO);",
						},
					},
				},
			},
			"secret": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure shared secret",
			},
			"secret_string": {
				Type: schema.TypeString, Optional: true, Description: "The RADIUS secret",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vrid": {
				Type: schema.TypeInt, Optional: true, Description: "Join a VRRP-A failover group",
			},
		},
	}
}
func resourceSystemRadiusServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemRadiusServerCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemRadiusServer(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemRadiusServerRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemRadiusServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemRadiusServerUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemRadiusServer(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemRadiusServerRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemRadiusServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemRadiusServerDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemRadiusServer(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemRadiusServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemRadiusServerRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemRadiusServer(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemRadiusServerAttribute(d []interface{}) []edpt.SystemRadiusServerAttribute {

	count1 := len(d)
	ret := make([]edpt.SystemRadiusServerAttribute, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemRadiusServerAttribute
		oi.AttributeValue = in["attribute_value"].(string)
		oi.PrefixLength = in["prefix_length"].(string)
		oi.PrefixVendor = in["prefix_vendor"].(int)
		oi.PrefixNumber = in["prefix_number"].(int)
		oi.Name = in["name"].(string)
		oi.Value = in["value"].(string)
		oi.CustomVendor = in["custom_vendor"].(int)
		oi.CustomNumber = in["custom_number"].(int)
		oi.Vendor = in["vendor"].(int)
		oi.Number = in["number"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemRadiusServerRemote(d []interface{}) edpt.SystemRadiusServerRemote {

	count1 := len(d)
	var ret edpt.SystemRadiusServerRemote
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpList = getSliceSystemRadiusServerRemoteIpList(in["ip_list"].([]interface{}))
	}
	return ret
}

func getSliceSystemRadiusServerRemoteIpList(d []interface{}) []edpt.SystemRadiusServerRemoteIpList {

	count1 := len(d)
	ret := make([]edpt.SystemRadiusServerRemoteIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemRadiusServerRemoteIpList
		oi.IpListName = in["ip_list_name"].(string)
		oi.IpListSecret = in["ip_list_secret"].(int)
		oi.IpListSecretString = in["ip_list_secret_string"].(string)
		//omit ip_list_encrypted
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemRadiusServerSamplingEnable(d []interface{}) []edpt.SystemRadiusServerSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemRadiusServerSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemRadiusServerSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemRadiusServer(d *schema.ResourceData) edpt.SystemRadiusServer {
	var ret edpt.SystemRadiusServer
	ret.Inst.AccountingInterimUpdate = d.Get("accounting_interim_update").(string)
	ret.Inst.AccountingOn = d.Get("accounting_on").(string)
	ret.Inst.AccountingStart = d.Get("accounting_start").(string)
	ret.Inst.AccountingStop = d.Get("accounting_stop").(string)
	ret.Inst.Attribute = getSliceSystemRadiusServerAttribute(d.Get("attribute").([]interface{}))
	ret.Inst.AttributeName = d.Get("attribute_name").(string)
	ret.Inst.CustomAttributeName = d.Get("custom_attribute_name").(string)
	ret.Inst.DisableReply = d.Get("disable_reply").(int)
	//omit encrypted
	ret.Inst.ListenPort = d.Get("listen_port").(int)
	ret.Inst.Remote = getObjectSystemRadiusServerRemote(d.Get("remote").([]interface{}))
	ret.Inst.SamplingEnable = getSliceSystemRadiusServerSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.Secret = d.Get("secret").(int)
	ret.Inst.SecretString = d.Get("secret_string").(string)
	//omit uuid
	ret.Inst.Vrid = d.Get("vrid").(int)
	return ret
}
