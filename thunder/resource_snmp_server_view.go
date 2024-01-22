package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerView() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_view`: Defines named \"view\" - a subset of the overall OID tree\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerViewCreate,
		UpdateContext: resourceSnmpServerViewUpdate,
		ReadContext:   resourceSnmpServerViewRead,
		DeleteContext: resourceSnmpServerViewDelete,

		Schema: map[string]*schema.Schema{
			"mask": {
				Type: schema.TypeString, Optional: true, Description: "Hex octets, separated by '.', mask of oid",
			},
			"oid": {
				Type: schema.TypeString, Required: true, Description: "MIB view family name or oid",
			},
			"type": {
				Type: schema.TypeString, Optional: true, Description: "'included': MIB family is included in the view; 'excluded': MIB family is excluded from the view;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"viewname": {
				Type: schema.TypeString, Required: true, Description: "Name of the view",
			},
		},
	}
}
func resourceSnmpServerViewCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerViewCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerView(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerViewRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerViewUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerViewUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerView(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerViewRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerViewDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerViewDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerView(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerViewRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerViewRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerView(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerView(d *schema.ResourceData) edpt.SnmpServerView {
	var ret edpt.SnmpServerView
	ret.Inst.Mask = d.Get("mask").(string)
	ret.Inst.Oid = d.Get("oid").(string)
	ret.Inst.Type = d.Get("type").(string)
	//omit uuid
	ret.Inst.Viewname = d.Get("viewname").(string)
	return ret
}
