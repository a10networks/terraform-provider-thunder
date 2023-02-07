package thunder

//Thunder resource smpp

import (
	"context"
	"log"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateSmpp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSmppCreate,
		UpdateContext: resourceSmppUpdate,
		ReadContext:   resourceSmppRead,
		DeleteContext: resourceSmppDelete,

		Schema: map[string]*schema.Schema{
			"client_enquire_link": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
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
			"server_enquire_link": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"server_selection_per_request": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"server_enquire_link_val": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSmppCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Creating smpp (Inside resourceSmppCreate    " + name)
		v := dataToSmpp(name, d)
		d.SetId(name)
		err := go_thunder.PostSmpp(client.Token, v, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSmppRead(ctx, d, meta)
	}
	return diags
}

func resourceSmppRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading smpp (Inside resourceSmppRead)")

	if client.Host != "" {
		client := meta.(Thunder)

		name := d.Id()

		logger.Println("[INFO] Fetching smppp Read" + name)

		smpp, err := go_thunder.GetSmpp(client.Token, name, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if smpp == nil {
			logger.Println("[INFO] No smpp found " + name)
			d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceSmppUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Modifying smpp (Inside resourceSmppUpdate    " + name)
		v := dataToSmpp(name, d)
		d.SetId(name)
		err := go_thunder.PutSmpp(client.Token, name, v, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSmppRead(ctx, d, meta)
	}
	return diags
}

func resourceSmppDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger := util.GetLoggerInstance()

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting smpp (Inside resourceSmppDelete) " + name)

		err := go_thunder.DeleteSmpp(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete smpp  (%s) (%v)", name, err)
			return diags
		}
		d.SetId("")
		return nil
	}
	return nil
}

//utility method to instantiate smpp structure
func dataToSmpp(name string, d *schema.ResourceData) go_thunder.Smpp {
	var s go_thunder.Smpp

	var sInstance go_thunder.SmppInstance
	sInstance.Name = d.Get("name").(string)
	sInstance.ServerEnquireLink = d.Get("server_enquire_link").(int)
	sInstance.ServerSelectionPerRequest = d.Get("server_selection_per_request").(int)
	sInstance.ClientEnquireLink = d.Get("client_enquire_link").(int)
	sInstance.UserTag = d.Get("user_tag").(string)
	sInstance.ServerEnquireLinkVal = d.Get("server_enquire_link_val").(int)
	sInstance.User = d.Get("user").(string)
	sInstance.Password = d.Get("password").(string)
	sInstance.UUID = d.Get("uuid").(string)

	s.Name = sInstance

	return s
}
