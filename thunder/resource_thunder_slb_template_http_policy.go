package thunder

//Thunder resource http policy

import (
	"context"
	"fmt"
	"log"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateHttpPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceHttpPolicyCreate,
		UpdateContext: resourceHttpPolicyUpdate,
		ReadContext:   resourceHttpPolicyRead,
		DeleteContext: resourceHttpPolicyDelete,

		Schema: map[string]*schema.Schema{
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"geo_location_match": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"geo_location": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"geo_location_service_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"geo_location_template_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"geo_location_template": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"http_policy_match": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"template": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"match_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"url_under_host": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"other_match_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"service_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"other_match_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"match_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"cookie_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceHttpPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Creating http policy (Inside resourceHttpPolicyCreate    " + name)
		v := dataToHttpPolicy(name, d)
		d.SetId(name)
		err := go_thunder.PostHttpPolicy(client.Token, v, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHttpPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceHttpPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	var diags diag.Diagnostics
	logger.Println("[INFO] Reading http policy (Inside resourceHttpPolicyRead)")

	if client.Host != "" {
		client := meta.(Thunder)

		name := d.Id()

		logger.Println("[INFO] Fetching htpp policy Read" + name)

		http_policy, err := go_thunder.GetHttpPolicy(client.Token, name, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if http_policy == nil {
			logger.Println("[INFO] No http policy found " + name)
			d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceHttpPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Modifying http policy (Inside resourceHttpPolicyUpdate    " + name)
		v := dataToHttpPolicy(name, d)
		d.SetId(name)
		err := go_thunder.PutHttpPolicy(client.Token, name, v, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHttpPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceHttpPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger := util.GetLoggerInstance()

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting http policy (Inside resourceHttpPolicyDelete) " + name)

		err := go_thunder.DeleteHttpPolicy(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete http policy  (%s) (%v)", name, err)
			return diags
		}
		d.SetId("")
		return nil
	}
	return nil
}

//utility method to instantiate http policy structure
func dataToHttpPolicy(name string, d *schema.ResourceData) go_thunder.HttpPolicy {
	var s go_thunder.HttpPolicy

	var sInstance go_thunder.HttpPolicyInstance

	sInstance.CookieName = d.Get("cookie_name").(string)
	sInstance.Name = d.Get("name").(string)
	sInstance.UUID = d.Get("uuid").(string)
	sInstance.UserTag = d.Get("user_tag").(string)

	match_cnt := d.Get("http_policy_match.#").(int)
	sInstance.MatchString = make([]go_thunder.HTTPPolicyMatch, 0, match_cnt)
	for i := 0; i < match_cnt; i++ {

		var match go_thunder.HTTPPolicyMatch
		prefix := fmt.Sprintf("http_policy_match.%d", i)
		match.Type = d.Get(prefix + ".type").(string)
		match.MatchString = d.Get(prefix + ".match_string").(string)
		match.TemplateName = d.Get(prefix + ".template_name").(string)
		match.ServiceGroup = d.Get(prefix + ".service_group").(string)
		match.URLUnderHost = d.Get(prefix + ".url_under_host").(string)
		match.OtherMatchString = d.Get(prefix + ".other_match_string").(string)
		match.Template = d.Get(prefix + ".template").(string)
		match.OtherMatchType = d.Get(prefix + ".other_match_type").(string)
		match.MatchType = d.Get(prefix + ".match_type").(string)

		sInstance.MatchString = append(sInstance.MatchString, match)
	}

	geo_match_cnt := d.Get("geo_location_match.#").(int)
	sInstance.GeoLocation = make([]go_thunder.GeoLocationMatch, 0, geo_match_cnt)
	for i := 0; i < geo_match_cnt; i++ {

		var match go_thunder.GeoLocationMatch
		prefix := fmt.Sprintf("geo_location_match.%d", i)
		match.GeoLocation = d.Get(prefix + ".geo_location").(string)
		match.GeoLocationServiceGroup = d.Get(prefix + ".geo_location_service_group").(string)
		match.GeoLocationTemplate = d.Get(prefix + ".geo_location_template").(string)
		match.GeoLocationTemplateName = d.Get(prefix + ".geo_location_template_name").(string)

		sInstance.GeoLocation = append(sInstance.GeoLocation, match)
	}

	s.Name = sInstance

	return s
}
