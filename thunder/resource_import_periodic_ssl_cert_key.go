package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceImportPeriodicSslCertKey() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_import_periodic_ssl_cert_key`: Bulk file\n\n__PLACEHOLDER__",
		CreateContext: resourceImportPeriodicSslCertKeyCreate,
		UpdateContext: resourceImportPeriodicSslCertKeyUpdate,
		ReadContext:   resourceImportPeriodicSslCertKeyRead,
		DeleteContext: resourceImportPeriodicSslCertKeyDelete,

		Schema: map[string]*schema.Schema{
			"period": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the period in second",
			},
			"remote_file": {
				Type: schema.TypeString, Optional: true, Description: "profile name for remote url",
			},
			"secured": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Mark keys as non-exportable",
			},
			"ssl_cert_key": {
				Type: schema.TypeString, Required: true, Description: "'bulk': import an archive file;",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceImportPeriodicSslCertKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicSslCertKeyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicSslCertKey(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportPeriodicSslCertKeyRead(ctx, d, meta)
	}
	return diags
}

func resourceImportPeriodicSslCertKeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicSslCertKeyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicSslCertKey(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportPeriodicSslCertKeyRead(ctx, d, meta)
	}
	return diags
}
func resourceImportPeriodicSslCertKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicSslCertKeyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicSslCertKey(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceImportPeriodicSslCertKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicSslCertKeyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicSslCertKey(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointImportPeriodicSslCertKey(d *schema.ResourceData) edpt.ImportPeriodicSslCertKey {
	var ret edpt.ImportPeriodicSslCertKey
	ret.Inst.Period = d.Get("period").(int)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.Secured = d.Get("secured").(int)
	ret.Inst.SslCertKey = d.Get("ssl_cert_key").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	//omit uuid
	return ret
}
