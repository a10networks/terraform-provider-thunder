package thunder

//Thunder resource SlbTemplateDiameter

import (
	"fmt"
	"log"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbTemplateDiameter() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbTemplateDiameterCreate,
		Update: resourceSlbTemplateDiameterUpdate,
		Read:   resourceSlbTemplateDiameterRead,
		Delete: resourceSlbTemplateDiameterDelete,
		Schema: map[string]*schema.Schema{
			"multiple_origin_host": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"vendor_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_to_latest_server": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"avp_code": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"origin_realm": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"dwr_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"avp_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"int32": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"int64": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"avp": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"mandatory": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"service_group_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"forward_unknown_session_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"load_balance_on_session_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"avp_string": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"terminate_on_cca_t": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dwr_up_retry": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"origin_host": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"origin_host_name": {
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
				},
			},
			"idle_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"message_code_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"message_code": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"customize_cea": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"product_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"session_age": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTemplateDiameterCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateDiameter (Inside resourceSlbTemplateDiameterCreate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateDiameter(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateDiameter --")
		d.SetId(name)
		go_thunder.PostSlbTemplateDiameter(client.Token, data, client.Host)

		return resourceSlbTemplateDiameterRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateDiameterRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbTemplateDiameter (Inside resourceSlbTemplateDiameterRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetSlbTemplateDiameter(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbTemplateDiameterUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplateDiameter   (Inside resourceSlbTemplateDiameterUpdate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateDiameter(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateDiameter ")
		d.SetId(name)
		go_thunder.PutSlbTemplateDiameter(client.Token, name, data, client.Host)

		return resourceSlbTemplateDiameterRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateDiameterDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateDiameterDelete) " + name)
		err := go_thunder.DeleteSlbTemplateDiameter(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbTemplateDiameter(d *schema.ResourceData) go_thunder.Diameter {
	var vc go_thunder.Diameter
	var c go_thunder.DiameterInstance

	c.AvpString = d.Get("avp_string").(string)
	c.TerminateOnCcaT = d.Get("terminate_on_cca_t").(int)
	c.ServiceGroupName = d.Get("service_group_name").(string)
	c.IdleTimeout = d.Get("idle_timeout").(int)
	c.CustomizeCea = d.Get("customize_cea").(int)
	c.ProductName = d.Get("product_name").(string)
	c.DwrUpRetry = d.Get("dwr_up_retry").(int)
	c.ForwardUnknownSessionID = d.Get("forward_unknown_session_id").(int)
	c.AvpCode = d.Get("avp_code").(int)
	c.VendorID = d.Get("vendor_id").(int)
	c.SessionAge = d.Get("session_age").(int)
	c.LoadBalanceOnSessionID = d.Get("load_balance_on_session_id").(int)
	c.Name = d.Get("name").(string)
	c.DwrTime = d.Get("dwr_time").(int)
	c.UserTag = d.Get("user_tag").(string)
	c.OriginRealm = d.Get("origin_realm").(string)
	c.MultipleOriginHost = d.Get("multiple_origin_host").(int)
	c.ForwardToLatestServer = d.Get("forward_to_latest_server").(int)

	MessageCodeListCount := d.Get("message_code_list.#").(int)
	c.MessageCode = make([]go_thunder.MessageCodeList, 0, MessageCodeListCount)
	for i := 0; i < MessageCodeListCount; i++ {
		var mcl go_thunder.MessageCodeList
		prefix := fmt.Sprintf("message_code_list.%d.", i)
		mcl.MessageCode = d.Get(prefix + "message_code").(int)
		c.MessageCode = append(c.MessageCode, mcl)
	}
	AvpListCount := d.Get("avp_list.#").(int)
	c.Avp = make([]go_thunder.AvpList, 0, AvpListCount)
	for i := 0; i < AvpListCount; i++ {
		var al go_thunder.AvpList
		prefix := fmt.Sprintf("avp_list.%d.", i)
		al.Int32 = d.Get(prefix + "int32").(int)
		al.Mandatory = d.Get(prefix + "mandatory").(int)
		al.String = d.Get(prefix + "string").(string)
		al.Avp = d.Get(prefix + "avp").(int)
		al.Int64 = d.Get(prefix + "int64").(int)
		c.Avp = append(c.Avp, al)
	}
	var oh go_thunder.OriginHost
	prefix := "origin_host.0."
	oh.OriginHostName = d.Get(prefix + "origin_host_name").(string)
	c.OriginHostName = oh
	vc.UUID = c
	return vc
}
