package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingAuditlog() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_logging_auditlog`: Set send auditlog to syslog host\n\n__PLACEHOLDER__",
		CreateContext: resourceLoggingAuditlogCreate,
		UpdateContext: resourceLoggingAuditlogUpdate,
		ReadContext:   resourceLoggingAuditlogRead,
		DeleteContext: resourceLoggingAuditlogDelete,

		Schema: map[string]*schema.Schema{
			"audit_facility": {
				Type: schema.TypeString, Optional: true, Description: "'local0': Local use; 'local1': Local use; 'local2': Local use; 'local3': Local use; 'local4': Local use; 'local5': Local use; 'local6': Local use; 'local7': Local use;  (Configure the facility of auditlog)",
			},
			"host4": {
				Type: schema.TypeString, Optional: true, Description: "Configure the auditlog host",
			},
			"host6": {
				Type: schema.TypeString, Optional: true, Description: "Configure the auditlog host",
			},
			"partition_name": {
				Type: schema.TypeString, Optional: true, Description: "Select partition name for logging",
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Default: 514, Description: "Set remote audit log port number(Default 514)",
			},
			"shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Select shared partition",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceLoggingAuditlogCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingAuditlogCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingAuditlog(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingAuditlogRead(ctx, d, meta)
	}
	return diags
}

func resourceLoggingAuditlogUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingAuditlogUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingAuditlog(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingAuditlogRead(ctx, d, meta)
	}
	return diags
}
func resourceLoggingAuditlogDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingAuditlogDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingAuditlog(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLoggingAuditlogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingAuditlogRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingAuditlog(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointLoggingAuditlog(d *schema.ResourceData) edpt.LoggingAuditlog {
	var ret edpt.LoggingAuditlog
	ret.Inst.AuditFacility = d.Get("audit_facility").(string)
	ret.Inst.Host4 = d.Get("host4").(string)
	ret.Inst.Host6 = d.Get("host6").(string)
	ret.Inst.PartitionName = d.Get("partition_name").(string)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.Shared = d.Get("shared").(int)
	//omit uuid
	return ret
}
