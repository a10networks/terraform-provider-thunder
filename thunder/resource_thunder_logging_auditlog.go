package thunder

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"util"
)

func resourceLoggingAuditlog() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceLoggingAuditlogCreate,
		UpdateContext: resourceLoggingAuditlogUpdate,
		ReadContext:   resourceLoggingAuditlogRead,
		DeleteContext: resourceLoggingAuditlogDelete,
		Schema: map[string]*schema.Schema{
			"shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Select shared partition",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"partition_name": {
				Type: schema.TypeString, Optional: true, Description: "Select partition name for logging",
				ValidateFunc: validation.StringLenBetween(1, 14),
			},
			"host6": {
				Type: schema.TypeString, Optional: true, Description: "Configure the auditlog host",
				ValidateFunc: validation.IsIPv6Address,
			},
			"host4": {
				Type: schema.TypeString, Optional: true, Description: "Configure the auditlog host",
				ValidateFunc: validation.StringLenBetween(1, 31),
			},
			"audit_facility": {
				Type: schema.TypeString, Optional: true, Description: "'local0': Local use; 'local1': Local use; 'local2': Local use; 'local3': Local use; 'local4': Local use; 'local5': Local use; 'local6': Local use; 'local7': Local use;  (Configure the facility of auditlog)",
				ValidateFunc: validation.StringInSlice([]string{"local0", "local1", "local2", "local3", "local4", "local5", "local6", "local7"}, false),
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Default: 514, Description: "Set remote audit log port number(Default 514)",
				ValidateFunc: validation.IntBetween(1, 32767),
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}

func resourceLoggingAuditlogCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("resourceLoggingAuditlogCreate()")
	client := meta.(Thunder)
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingAuditlog(d)
		d.SetId(obj.GetId())
		err := PostEndpointLoggingAuditlog(client.Token, obj, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingAuditlogRead(ctx, d, meta)
	}
	return diags
}

func resourceLoggingAuditlogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("resourceLoggingAuditlogRead()")
	client := meta.(Thunder)
	var diags diag.Diagnostics
	if client.Host != "" {
		_, err := GetEndpointLoggingAuditlog(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLoggingAuditlogUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("resourceLoggingAuditlogUpdate()")
	client := meta.(Thunder)
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingAuditlog(d)
		err := PutEndpointLoggingAuditlog(client.Token, obj, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingAuditlogRead(ctx, d, meta)
	}
	return diags
}

func resourceLoggingAuditlogDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("resourceLoggingAuditlogDelete()")
	client := meta.(Thunder)
	var diags diag.Diagnostics
	if client.Host != "" {
		err := DeleteEndpointLoggingAuditlog(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointLoggingAuditlog(d *schema.ResourceData) EndpointLoggingAuditlog {
	var ret EndpointLoggingAuditlog
	ret.Inst.Shared = d.Get("shared").(int)
	ret.Inst.PartitionName = d.Get("partition_name").(string)
	ret.Inst.Host6 = d.Get("host6").(string)
	ret.Inst.Host4 = d.Get("host4").(string)
	ret.Inst.AuditFacility = d.Get("audit_facility").(string)
	ret.Inst.Port = d.Get("port").(int)
	//omit uuid
	return ret
}
