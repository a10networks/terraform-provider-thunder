package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemMfaCertStore() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_mfa_cert_store`: Define a 2FA management client certificate store\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemMfaCertStoreCreate,
		UpdateContext: resourceSystemMfaCertStoreUpdate,
		ReadContext:   resourceSystemMfaCertStoreRead,
		DeleteContext: resourceSystemMfaCertStoreDelete,

		Schema: map[string]*schema.Schema{
			"cert_host": {
				Type: schema.TypeString, Optional: true, Description: "Configure certificate store host",
			},
			"cert_store_path": {
				Type: schema.TypeString, Optional: true, Description: "Configure certificate store path",
			},
			"passwd_string": {
				Type: schema.TypeString, Optional: true, Description: "Certificate store host password",
			},
			"protocol": {
				Type: schema.TypeString, Optional: true, Description: "'tftp': Use tftp for connection; 'ftp': Use ftp for connection; 'scp': Use scp for connection; 'http': Use http for connection; 'https': Use https for connection; 'sftp': Use sftp for connection;",
			},
			"username": {
				Type: schema.TypeString, Optional: true, Description: "Certificate store host username",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemMfaCertStoreCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMfaCertStoreCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMfaCertStore(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemMfaCertStoreRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemMfaCertStoreUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMfaCertStoreUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMfaCertStore(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemMfaCertStoreRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemMfaCertStoreDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMfaCertStoreDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMfaCertStore(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemMfaCertStoreRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMfaCertStoreRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMfaCertStore(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemMfaCertStore(d *schema.ResourceData) edpt.SystemMfaCertStore {
	var ret edpt.SystemMfaCertStore
	ret.Inst.CertHost = d.Get("cert_host").(string)
	ret.Inst.CertStorePath = d.Get("cert_store_path").(string)
	//omit encrypted
	ret.Inst.PasswdString = d.Get("passwd_string").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Username = d.Get("username").(string)
	//omit uuid
	return ret
}
