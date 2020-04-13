package vthunder

//vThunder resource SlbTemplateSIP

import (
	"fmt"
	"log"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbTemplateSIP() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbTemplateSIPCreate,
		Update: resourceSlbTemplateSIPUpdate,
		Read:   resourceSlbTemplateSIPRead,
		Delete: resourceSlbTemplateSIPDelete,
		Schema: map[string]*schema.Schema{
			"insert_client_ip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"failed_server_selection": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"smp_call_id_rtp_session": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"client_keep_alive": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"exclude_translation": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"translation_value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"header_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"pstn_gw": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"acl_name_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"drop_when_client_fail": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"acl_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"dialog_aware": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"call_id_persist_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"client_request_header": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_request_header_erase": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"insert_condition_client_request": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"client_request_erase_all": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"client_request_header_insert": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"keep_server_ip_if_match_acl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"server_request_header": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"insert_condition_server_request": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_request_header_insert": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_request_erase_all": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"server_request_header_erase": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"failed_server_selection_message": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"failed_client_selection": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"alg_dest_nat": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"drop_when_server_fail": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"failed_client_selection_message": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"server_selection_per_request": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"alg_source_nat": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"service_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"server_keep_alive": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"server_response_header": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server_response_header_insert": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"insert_condition_server_response": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_response_header_erase": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_response_erase_all": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"client_response_header": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_response_header_insert": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"client_response_header_erase": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"client_response_erase_all": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"insert_condition_client_response": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbTemplateSIPCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateSIP (Inside resourceSlbTemplateSIPCreate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateSIP(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateSIP --")
		d.SetId(name)
		go_vthunder.PostSlbTemplateSIP(client.Token, data, client.Host)

		return resourceSlbTemplateSIPRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateSIPRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbTemplateSIP (Inside resourceSlbTemplateSIPRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetSlbTemplateSIP(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbTemplateSIPUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplateSIP   (Inside resourceSlbTemplateSIPUpdate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateSIP(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateSIP ")
		d.SetId(name)
		go_vthunder.PutSlbTemplateSIP(client.Token, name, data, client.Host)

		return resourceSlbTemplateSIPRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateSIPDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateSIPDelete) " + name)
		err := go_vthunder.DeleteSlbTemplateSIP(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbTemplateSIP(d *schema.ResourceData) go_vthunder.SIP {
	var vc go_vthunder.SIP
	var c go_vthunder.SIPInstance

	ServerRequestHeaderCount := d.Get("server_request_header.#").(int)
	c.ServerRequestHeaderInsert = make([]go_vthunder.ServerRequestHeader, 0, ServerRequestHeaderCount)

	for i := 0; i < ServerRequestHeaderCount; i++ {
		var obj1 go_vthunder.ServerRequestHeader
		prefix := fmt.Sprintf("server_request_header.%d.", i)
		obj1.ServerRequestHeaderInsert = d.Get(prefix + "server_request_header_insert").(string)
		obj1.ServerRequestEraseAll = d.Get(prefix + "server_request_erase_all").(int)
		obj1.InsertConditionServerRequest = d.Get(prefix + "insert_condition_server_request").(string)
		obj1.ServerRequestHeaderErase = d.Get(prefix + "server_request_header_erase").(string)
		c.ServerRequestHeaderInsert = append(c.ServerRequestHeaderInsert, obj1)
	}

	c.SmpCallIDRtpSession = d.Get("smp_call_id_rtp_session").(int)
	c.KeepServerIPIfMatchACL = d.Get("keep_server_ip_if_match_acl").(int)
	c.ClientKeepAlive = d.Get("client_keep_alive").(int)
	c.AlgSourceNat = d.Get("alg_source_nat").(int)

	ServerResponseHeaderCount := d.Get("server_response_header.#").(int)
	c.ServerResponseHeaderInsert = make([]go_vthunder.ServerResponseHeader, 0, ServerResponseHeaderCount)

	for i := 0; i < ServerResponseHeaderCount; i++ {
		var obj2 go_vthunder.ServerResponseHeader
		prefix := fmt.Sprintf("server_response_header.%d.", i)
		obj2.ServerResponseHeaderInsert = d.Get(prefix + "server_response_header_insert").(string)
		obj2.InsertConditionServerResponse = d.Get(prefix + "insert_condition_server_response").(string)
		obj2.ServerResponseHeaderErase = d.Get(prefix + "server_response_header_erase").(string)
		obj2.ServerResponseEraseAll = d.Get(prefix + "server_response_erase_all").(int)
		c.ServerResponseHeaderInsert = append(c.ServerResponseHeaderInsert, obj2)
	}

	c.ServerSelectionPerRequest = d.Get("server_selection_per_request").(int)

	ClientRequestHeaderCount := d.Get("client_request_header.#").(int)
	c.ClientRequestHeaderErase = make([]go_vthunder.ClientRequestHeader, 0, ClientRequestHeaderCount)

	for i := 0; i < ClientRequestHeaderCount; i++ {
		var obj3 go_vthunder.ClientRequestHeader
		prefix := fmt.Sprintf("client_request_header.%d.", i)
		obj3.ClientRequestHeaderErase = d.Get(prefix + "client_request_header_erase").(string)
		obj3.ClientRequestHeaderInsert = d.Get(prefix + "client_request_header_insert").(string)
		obj3.ClientRequestEraseAll = d.Get(prefix + "client_request_erase_all").(int)
		obj3.InsertConditionClientRequest = d.Get(prefix + "insert_condition_client_request").(string)
		c.ClientRequestHeaderErase = append(c.ClientRequestHeaderErase, obj3)
	}

	c.PstnGw = d.Get("pstn_gw").(string)
	c.ServiceGroup = d.Get("service_group").(string)
	c.InsertClientIP = d.Get("insert_client_ip").(int)
	c.FailedClientSelection = d.Get("failed_client_selection").(int)
	c.FailedClientSelectionMessage = d.Get("failed_client_selection_message").(string)
	c.CallIDPersistDisable = d.Get("call_id_persist_disable").(int)
	c.ACLID = d.Get("acl_id").(int)
	c.AlgDestNat = d.Get("alg_dest_nat").(int)
	c.ServerKeepAlive = d.Get("server_keep_alive").(int)

	ClientResponseHeaderCount := d.Get("client_response_header.#").(int)
	c.ClientResponseEraseAll = make([]go_vthunder.ClientResponseHeader, 0, ClientResponseHeaderCount)

	for i := 0; i < ClientResponseHeaderCount; i++ {
		var obj4 go_vthunder.ClientResponseHeader
		prefix := fmt.Sprintf("client_response_header.%d.", i)
		obj4.ClientResponseEraseAll = d.Get(prefix + "client_response_erase_all").(int)
		obj4.InsertConditionClientResponse = d.Get(prefix + "insert_condition_client_response").(string)
		obj4.ClientResponseHeaderErase = d.Get(prefix + "client_response_header_erase").(string)
		obj4.ClientResponseHeaderInsert = d.Get(prefix + "client_response_header_insert").(string)
		c.ClientResponseEraseAll = append(c.ClientResponseEraseAll, obj4)
	}

	c.FailedServerSelectionMessage = d.Get("failed_server_selection_message").(string)
	c.Name = d.Get("name").(string)

	ExcludeTranslationCount := d.Get("exclude_translation.#").(int)
	c.TranslationValue = make([]go_vthunder.ExcludeTranslation, 0, ExcludeTranslationCount)

	for i := 0; i < ExcludeTranslationCount; i++ {
		var obj5 go_vthunder.ExcludeTranslation
		prefix := fmt.Sprintf("exclude_translation.%d.", i)
		obj5.TranslationValue = d.Get(prefix + "translation_value").(string)
		obj5.HeaderString = d.Get(prefix + "header_string").(string)
		c.TranslationValue = append(c.TranslationValue, obj5)
	}

	c.Interval = d.Get("interval").(int)
	c.UserTag = d.Get("user_tag").(string)
	c.DialogAware = d.Get("dialog_aware").(int)
	c.FailedServerSelection = d.Get("failed_server_selection").(int)
	c.DropWhenClientFail = d.Get("drop_when_client_fail").(int)
	c.Timeout = d.Get("timeout").(int)
	c.DropWhenServerFail = d.Get("drop_when_server_fail").(int)
	c.ACLNameValue = d.Get("acl_name_value").(string)
	vc.UUID = c
	return vc
}
