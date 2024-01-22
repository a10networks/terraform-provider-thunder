package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingFacility() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_logging_facility`: Facility parameter for syslog messages\n\n__PLACEHOLDER__",
		CreateContext: resourceLoggingFacilityCreate,
		UpdateContext: resourceLoggingFacilityUpdate,
		ReadContext:   resourceLoggingFacilityRead,
		DeleteContext: resourceLoggingFacilityDelete,

		Schema: map[string]*schema.Schema{
			"facilityname": {
				Type: schema.TypeString, Optional: true, Default: "local0", Description: "'local0': Local use; 'local1': Local use; 'local2': Local use; 'local3': Local use; 'local4': Local use; 'local5': Local use; 'local6': Local use; 'local7': Local use;  (Facility parameter for syslog messages)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceLoggingFacilityCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingFacilityCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingFacility(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingFacilityRead(ctx, d, meta)
	}
	return diags
}

func resourceLoggingFacilityUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingFacilityUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingFacility(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingFacilityRead(ctx, d, meta)
	}
	return diags
}
func resourceLoggingFacilityDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingFacilityDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingFacility(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLoggingFacilityRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingFacilityRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingFacility(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointLoggingFacility(d *schema.ResourceData) edpt.LoggingFacility {
	var ret edpt.LoggingFacility
	ret.Inst.Facilityname = d.Get("facilityname").(string)
	//omit uuid
	return ret
}
