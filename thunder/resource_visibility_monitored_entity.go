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

func getObjectVisibilityMonitoredEntityDetail1927(d []interface{}) edpt.VisibilityMonitoredEntityDetail1927 {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityDetail1927
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.Debug = getObjectVisibilityMonitoredEntityDetailDebug1928(in["debug"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityMonitoredEntityDetailDebug1928(d []interface{}) edpt.VisibilityMonitoredEntityDetailDebug1928 {

	var ret edpt.VisibilityMonitoredEntityDetailDebug1928
	return ret
}

func getObjectVisibilityMonitoredEntityMonTopk1929(d []interface{}) edpt.VisibilityMonitoredEntityMonTopk1929 {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityMonTopk1929
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.Sources = getObjectVisibilityMonitoredEntityMonTopkSources1930(in["sources"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityMonitoredEntityMonTopkSources1930(d []interface{}) edpt.VisibilityMonitoredEntityMonTopkSources1930 {

	var ret edpt.VisibilityMonitoredEntityMonTopkSources1930
	return ret
}

func getObjectVisibilityMonitoredEntitySecondary1931(d []interface{}) edpt.VisibilityMonitoredEntitySecondary1931 {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntitySecondary1931
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MonTopk = getObjectVisibilityMonitoredEntitySecondaryMonTopk1932(in["mon_topk"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityMonitoredEntitySecondaryMonTopk1932(d []interface{}) edpt.VisibilityMonitoredEntitySecondaryMonTopk1932 {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntitySecondaryMonTopk1932
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.Sources = getObjectVisibilityMonitoredEntitySecondaryMonTopkSources1933(in["sources"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityMonitoredEntitySecondaryMonTopkSources1933(d []interface{}) edpt.VisibilityMonitoredEntitySecondaryMonTopkSources1933 {

	var ret edpt.VisibilityMonitoredEntitySecondaryMonTopkSources1933
	return ret
}

func getObjectVisibilityMonitoredEntitySessions1934(d []interface{}) edpt.VisibilityMonitoredEntitySessions1934 {

	var ret edpt.VisibilityMonitoredEntitySessions1934
	return ret
}

func dataToEndpointVisibilityMonitoredEntity(d *schema.ResourceData) edpt.VisibilityMonitoredEntity {
	var ret edpt.VisibilityMonitoredEntity
	ret.Inst.Detail = getObjectVisibilityMonitoredEntityDetail1927(d.Get("detail").([]interface{}))
	ret.Inst.MonTopk = getObjectVisibilityMonitoredEntityMonTopk1929(d.Get("mon_topk").([]interface{}))
	ret.Inst.Secondary = getObjectVisibilityMonitoredEntitySecondary1931(d.Get("secondary").([]interface{}))
	ret.Inst.Sessions = getObjectVisibilityMonitoredEntitySessions1934(d.Get("sessions").([]interface{}))
	//omit uuid
	return ret
}
