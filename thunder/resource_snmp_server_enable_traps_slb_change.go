package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerEnableTrapsSlbChange() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_enable_traps_slb_change`: Enable SLB change traps\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerEnableTrapsSlbChangeCreate,
		UpdateContext: resourceSnmpServerEnableTrapsSlbChangeUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsSlbChangeRead,
		DeleteContext: resourceSnmpServerEnableTrapsSlbChangeDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all system group traps",
			},
			"connection_resource_event": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system connection resource event trap",
			},
			"resource_usage_warning": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable partition resource usage warning trap",
			},
			"server": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable slb server create/delete trap",
			},
			"server_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable slb server port create/delete trap",
			},
			"ssl_cert_change": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SSL certificate change trap",
			},
			"ssl_cert_expire": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SSL certificate expiring trap",
			},
			"system_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable slb system threshold trap",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable slb vip create/delete trap",
			},
			"vip_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable slb vip-port create/delete trap",
			},
		},
	}
}
func resourceSnmpServerEnableTrapsSlbChangeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsSlbChangeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsSlbChange(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsSlbChangeRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerEnableTrapsSlbChangeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsSlbChangeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsSlbChange(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsSlbChangeRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerEnableTrapsSlbChangeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsSlbChangeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsSlbChange(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerEnableTrapsSlbChangeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsSlbChangeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsSlbChange(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerEnableTrapsSlbChange(d *schema.ResourceData) edpt.SnmpServerEnableTrapsSlbChange {
	var ret edpt.SnmpServerEnableTrapsSlbChange
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.ConnectionResourceEvent = d.Get("connection_resource_event").(int)
	ret.Inst.ResourceUsageWarning = d.Get("resource_usage_warning").(int)
	ret.Inst.Server = d.Get("server").(int)
	ret.Inst.ServerPort = d.Get("server_port").(int)
	ret.Inst.SslCertChange = d.Get("ssl_cert_change").(int)
	ret.Inst.SslCertExpire = d.Get("ssl_cert_expire").(int)
	ret.Inst.SystemThreshold = d.Get("system_threshold").(int)
	//omit uuid
	ret.Inst.Vip = d.Get("vip").(int)
	ret.Inst.VipPort = d.Get("vip_port").(int)
	return ret
}
