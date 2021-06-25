package thunder

//Thunder resource WebCategoryProxyServer

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceWebCategoryProxyServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebCategoryProxyServerCreate,
		Update: resourceWebCategoryProxyServerUpdate,
		Read:   resourceWebCategoryProxyServerRead,
		Delete: resourceWebCategoryProxyServerDelete,
		Schema: map[string]*schema.Schema{
			"proxy_host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"http_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"https_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"auth_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"password": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"secret_string": {
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

func resourceWebCategoryProxyServerCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating WebCategoryProxyServer (Inside resourceWebCategoryProxyServerCreate) ")

		data := dataToWebCategoryProxyServer(d)
		logger.Println("[INFO] received formatted data from method data to WebCategoryProxyServer --")
		d.SetId("1")
		go_thunder.PostWebCategoryProxyServer(client.Token, data, client.Host)

		return resourceWebCategoryProxyServerRead(d, meta)

	}
	return nil
}

func resourceWebCategoryProxyServerRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading WebCategoryProxyServer (Inside resourceWebCategoryProxyServerRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetWebCategoryProxyServer(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return err
	}
	return nil
}

func resourceWebCategoryProxyServerUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceWebCategoryProxyServerRead(d, meta)
}

func resourceWebCategoryProxyServerDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceWebCategoryProxyServerRead(d, meta)
}
func dataToWebCategoryProxyServer(d *schema.ResourceData) go_thunder.WebCategoryProxyServer {
	var vc go_thunder.WebCategoryProxyServer
	var c go_thunder.WebCategoryProxyServerInstance
	c.ProxyHost = d.Get("proxy_host").(string)
	c.HTTPPort = d.Get("http_port").(int)
	c.HTTPSPort = d.Get("https_port").(int)
	c.AuthType = d.Get("auth_type").(string)
	c.Domain = d.Get("domain").(string)
	c.Username = d.Get("username").(string)
	c.Password = d.Get("password").(int)
	c.SecretString = d.Get("secret_string").(string)

	vc.ProxyHost = c
	return vc
}
