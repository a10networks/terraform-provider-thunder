package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_group`: Define a User Security Model group\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerGroupCreate,
		UpdateContext: resourceSnmpServerGroupUpdate,
		ReadContext:   resourceSnmpServerGroupRead,
		DeleteContext: resourceSnmpServerGroupDelete,

		Schema: map[string]*schema.Schema{
			"groupname": {
				Type: schema.TypeString, Required: true, Description: "Name of the group",
			},
			"read": {
				Type: schema.TypeString, Optional: true, Description: "specify a read view for the group (read view name)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"v3": {
				Type: schema.TypeString, Optional: true, Description: "'auth': group using the authNoPriv Security Level; 'noauth': group using the noAuthNoPriv Security Level; 'priv': group using SNMPv3 authPriv security level;",
			},
		},
	}
}
func resourceSnmpServerGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerGroup(d *schema.ResourceData) edpt.SnmpServerGroup {
	var ret edpt.SnmpServerGroup
	ret.Inst.Groupname = d.Get("groupname").(string)
	ret.Inst.Read = d.Get("read").(string)
	//omit uuid
	ret.Inst.V3 = d.Get("v3").(string)
	return ret
}
