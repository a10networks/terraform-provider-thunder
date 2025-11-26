package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonitoredEntity() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_monitored_entity`: Display Monitoring entities\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityMonitoredEntityCreate,
		UpdateContext: resourceVisibilityMonitoredEntityUpdate,
		ReadContext:   resourceVisibilityMonitoredEntityRead,
		DeleteContext: resourceVisibilityMonitoredEntityDelete,

		Schema: map[string]*schema.Schema{
			"detail": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"debug": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"mon_topk": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sources": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"secondary": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mon_topk": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"sources": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"sessions": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVisibilityMonitoredEntityCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitoredEntityCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitoredEntity(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitoredEntityRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityMonitoredEntityUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitoredEntityUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitoredEntity(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitoredEntityRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityMonitoredEntityDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitoredEntityDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitoredEntity(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityMonitoredEntityRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitoredEntityRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitoredEntity(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityMonitoredEntityDetail2052(d []interface{}) edpt.VisibilityMonitoredEntityDetail2052 {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityDetail2052
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.Debug = getObjectVisibilityMonitoredEntityDetailDebug2053(in["debug"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityMonitoredEntityDetailDebug2053(d []interface{}) edpt.VisibilityMonitoredEntityDetailDebug2053 {

	var ret edpt.VisibilityMonitoredEntityDetailDebug2053
	return ret
}

func getObjectVisibilityMonitoredEntityMonTopk2054(d []interface{}) edpt.VisibilityMonitoredEntityMonTopk2054 {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityMonTopk2054
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.Sources = getObjectVisibilityMonitoredEntityMonTopkSources2055(in["sources"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityMonitoredEntityMonTopkSources2055(d []interface{}) edpt.VisibilityMonitoredEntityMonTopkSources2055 {

	var ret edpt.VisibilityMonitoredEntityMonTopkSources2055
	return ret
}

func getObjectVisibilityMonitoredEntitySecondary2056(d []interface{}) edpt.VisibilityMonitoredEntitySecondary2056 {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntitySecondary2056
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MonTopk = getObjectVisibilityMonitoredEntitySecondaryMonTopk2057(in["mon_topk"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityMonitoredEntitySecondaryMonTopk2057(d []interface{}) edpt.VisibilityMonitoredEntitySecondaryMonTopk2057 {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntitySecondaryMonTopk2057
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.Sources = getObjectVisibilityMonitoredEntitySecondaryMonTopkSources2058(in["sources"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityMonitoredEntitySecondaryMonTopkSources2058(d []interface{}) edpt.VisibilityMonitoredEntitySecondaryMonTopkSources2058 {

	var ret edpt.VisibilityMonitoredEntitySecondaryMonTopkSources2058
	return ret
}

func getObjectVisibilityMonitoredEntitySessions2059(d []interface{}) edpt.VisibilityMonitoredEntitySessions2059 {

	var ret edpt.VisibilityMonitoredEntitySessions2059
	return ret
}

func dataToEndpointVisibilityMonitoredEntity(d *schema.ResourceData) edpt.VisibilityMonitoredEntity {
	var ret edpt.VisibilityMonitoredEntity
	ret.Inst.Detail = getObjectVisibilityMonitoredEntityDetail2052(d.Get("detail").([]interface{}))
	ret.Inst.MonTopk = getObjectVisibilityMonitoredEntityMonTopk2054(d.Get("mon_topk").([]interface{}))
	ret.Inst.Secondary = getObjectVisibilityMonitoredEntitySecondary2056(d.Get("secondary").([]interface{}))
	ret.Inst.Sessions = getObjectVisibilityMonitoredEntitySessions2059(d.Get("sessions").([]interface{}))
	//omit uuid
	return ret
}
