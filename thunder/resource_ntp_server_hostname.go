package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNtpServerHostname() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ntp_server_hostname`: IPV4 address, IPV6 address or host name of NTP server\n\n__PLACEHOLDER__",
		CreateContext: resourceNtpServerHostnameCreate,
		UpdateContext: resourceNtpServerHostnameUpdate,
		ReadContext:   resourceNtpServerHostnameRead,
		DeleteContext: resourceNtpServerHostnameDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable this NTP server; 'disable': Disable this NTP server;",
			},
			"host_servername": {
				Type: schema.TypeString, Required: true, Description: "IPV4 address, IPV6 address or host name of NTP server(string1~63)",
			},
			"key": {
				Type: schema.TypeInt, Optional: true, Description: "Use trusted key to authenticate this NTP server (The trusted key number to use)",
			},
			"prefer": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set this NTP server as preferred",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceNtpServerHostnameCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNtpServerHostnameCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNtpServerHostname(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNtpServerHostnameRead(ctx, d, meta)
	}
	return diags
}

func resourceNtpServerHostnameUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNtpServerHostnameUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNtpServerHostname(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNtpServerHostnameRead(ctx, d, meta)
	}
	return diags
}
func resourceNtpServerHostnameDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNtpServerHostnameDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNtpServerHostname(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNtpServerHostnameRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNtpServerHostnameRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNtpServerHostname(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNtpServerHostname(d *schema.ResourceData) edpt.NtpServerHostname {
	var ret edpt.NtpServerHostname
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.HostServername = d.Get("host_servername").(string)
	ret.Inst.Key = d.Get("key").(int)
	ret.Inst.Prefer = d.Get("prefer").(int)
	//omit uuid
	return ret
}
