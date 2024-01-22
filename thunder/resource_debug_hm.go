package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugHm() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_hm`: Debug health monitor\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugHmCreate,
		UpdateContext: resourceDebugHmUpdate,
		ReadContext:   resourceDebugHmRead,
		DeleteContext: resourceDebugHmDelete,

		Schema: map[string]*schema.Schema{
			"level": {
				Type: schema.TypeInt, Optional: true, Description: "Debug level (Level 1-3)",
			},
			"method_type": {
				Type: schema.TypeString, Optional: true, Description: "'icmp': ICMP type; 'tcp': TCP type; 'udp': UDP type; 'ftp': FTP type; 'http': HTTP type; 'snmp': SNMP type; 'smtp': SMTP type; 'dns': DNS type; 'dns-tcp': DNS TCP type; 'pop3': POP3 type; 'imap': IMAP type; 'sip': SIP type; 'sip-tcp': SIP TCP type; 'radius': RADIUS type; 'ldap': LDAP type; 'rtsp': RTSP type; 'kerberos-kdc': Kerberos KDC type; 'database': DATABASE type; 'external': EXTERNAL type; 'https': HTTPS type; 'ntp': NTP type; 'compound': Compound type;",
			},
			"pin_uid": {
				Type: schema.TypeInt, Optional: true, Description: "Debug Pin Unique Id",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugHmCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHmCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHm(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugHmRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugHmUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHmUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHm(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugHmRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugHmDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHmDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHm(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugHmRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugHmRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugHm(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugHm(d *schema.ResourceData) edpt.DebugHm {
	var ret edpt.DebugHm
	ret.Inst.Level = d.Get("level").(int)
	ret.Inst.MethodType = d.Get("method_type").(string)
	ret.Inst.PinUid = d.Get("pin_uid").(int)
	//omit uuid
	return ret
}
