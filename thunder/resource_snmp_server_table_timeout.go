package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerTableTimeout() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_table_timeout`: Define SNMP table timeout\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerTableTimeoutCreate,
		UpdateContext: resourceSnmpServerTableTimeoutUpdate,
		ReadContext:   resourceSnmpServerTableTimeoutRead,
		DeleteContext: resourceSnmpServerTableTimeoutDelete,

		Schema: map[string]*schema.Schema{
			"table": {
				Type: schema.TypeString, Required: true, Description: "SNMP table name or table oid",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "timeout value",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSnmpServerTableTimeoutCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerTableTimeoutCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerTableTimeout(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerTableTimeoutRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerTableTimeoutUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerTableTimeoutUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerTableTimeout(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerTableTimeoutRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerTableTimeoutDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerTableTimeoutDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerTableTimeout(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerTableTimeoutRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerTableTimeoutRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerTableTimeout(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerTableTimeout(d *schema.ResourceData) edpt.SnmpServerTableTimeout {
	var ret edpt.SnmpServerTableTimeout
	ret.Inst.Table = d.Get("table").(string)
	ret.Inst.Timeout = d.Get("timeout").(int)
	//omit uuid
	return ret
}
