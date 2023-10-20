package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGslbTemplateSnmp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_template_snmp`: Specify SNMP template\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbTemplateSnmpCreate,
		UpdateContext: resourceGslbTemplateSnmpUpdate,
		ReadContext:   resourceGslbTemplateSnmpRead,
		DeleteContext: resourceGslbTemplateSnmpDelete,
		Schema: map[string]*schema.Schema{
			"auth_key": {
				Type: schema.TypeString, Optional: true, Description: "Specify authentication key (Specify key)",
				ValidateFunc: validation.StringLenBetween(1, 127),
			},
			"auth_proto": {
				Type: schema.TypeString, Optional: true, Default: "md5", Description: "'sha': SHA; 'md5': MD5;",
				ValidateFunc: validation.StringInSlice([]string{"sha", "md5"}, false),
			},
			"community": {
				Type: schema.TypeString, Optional: true, Description: "Specify community for version 2c (Community name)",
				ValidateFunc: validation.StringLenBetween(1, 127),
			},
			"context_engine_id": {
				Type: schema.TypeString, Optional: true, Description: "Specify context engine ID",
				ValidateFunc: validation.StringLenBetween(1, 127),
			},
			"context_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify context name",
				ValidateFunc: validation.StringLenBetween(1, 127),
			},
			"host": {
				Type: schema.TypeString, Optional: true, Description: "Specify host (Host name or ip address)",
				ValidateFunc: validation.StringLenBetween(1, 127),
			},
			"interface": {
				Type: schema.TypeInt, Optional: true, Description: "Specify Interface ID",
				ValidateFunc: validation.IntBetween(0, 2147483647),
			},
			"interval": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify interval, default is 3 (Interval, unit: second, default is 3)",
				ValidateFunc: validation.IntBetween(1, 999),
			},
			"oid": {
				Type: schema.TypeString, Optional: true, Description: "Specify OID",
				ValidateFunc: validation.StringLenBetween(1, 127),
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Default: 161, Description: "Specify port, default is 161 (Port Number, default is 161)",
				ValidateFunc: validation.IntBetween(1, 65535),
			},
			"priv_key": {
				Type: schema.TypeString, Optional: true, Description: "Specify privacy key (Specify key)",
				ValidateFunc: validation.StringLenBetween(1, 127),
			},
			"priv_proto": {
				Type: schema.TypeString, Optional: true, Default: "des", Description: "'aes': AES; 'des': DES;",
				ValidateFunc: validation.StringInSlice([]string{"aes", "des"}, false),
			},
			"security_engine_id": {
				Type: schema.TypeString, Optional: true, Description: "Specify security engine ID",
				ValidateFunc: validation.StringLenBetween(1, 127),
			},
			"security_level": {
				Type: schema.TypeString, Optional: true, Default: "no-auth", Description: "'no-auth': No authentication; 'auth-no-priv': Authentication, but no privacy; 'auth-priv': Authentication and privacy;",
				ValidateFunc: validation.StringInSlice([]string{"no-auth", "auth-no-priv", "auth-priv"}, false),
			},
			"snmp_name": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "Specify name of snmp template",
				ValidateFunc: validation.StringLenBetween(1, 63),
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
				ValidateFunc: validation.StringLenBetween(1, 127),
			},
			"username": {
				Type: schema.TypeString, Optional: true, Description: "Specify username (User name)",
				ValidateFunc: validation.StringLenBetween(1, 63),
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"version": {
				Type: schema.TypeString, Optional: true, Default: "v3", Description: "'v1': Version 1; 'v2c': Version 2c; 'v3': Version 3;",
				ValidateFunc: validation.StringInSlice([]string{"v1", "v2c", "v3"}, false),
			},
		},
	}
}

func resourceGslbTemplateSnmpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbTemplateSnmpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbTemplateSnmp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbTemplateSnmpRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbTemplateSnmpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbTemplateSnmpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbTemplateSnmp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbTemplateSnmpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbTemplateSnmpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbTemplateSnmp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbTemplateSnmpRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbTemplateSnmpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbTemplateSnmpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbTemplateSnmp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointGslbTemplateSnmp(d *schema.ResourceData) edpt.GslbTemplateSnmp {
	var ret edpt.GslbTemplateSnmp
	ret.Inst.AuthKey = d.Get("auth_key").(string)
	ret.Inst.AuthProto = d.Get("auth_proto").(string)
	ret.Inst.Community = d.Get("community").(string)
	ret.Inst.ContextEngineId = d.Get("context_engine_id").(string)
	ret.Inst.ContextName = d.Get("context_name").(string)
	ret.Inst.Host = d.Get("host").(string)
	ret.Inst.Interface = d.Get("interface").(int)
	ret.Inst.Interval = d.Get("interval").(int)
	ret.Inst.Oid = d.Get("oid").(string)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.PrivKey = d.Get("priv_key").(string)
	ret.Inst.PrivProto = d.Get("priv_proto").(string)
	ret.Inst.SecurityEngineId = d.Get("security_engine_id").(string)
	ret.Inst.SecurityLevel = d.Get("security_level").(string)
	ret.Inst.SnmpName = d.Get("snmp_name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	ret.Inst.Username = d.Get("username").(string)
	//omit uuid
	ret.Inst.Version = d.Get("version").(string)
	return ret
}
