provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_http_policy" "resourceSlbTemplateHttpPolicy" {
  name        = "http_policy"
  cookie_name = "cookie-two-three"
  user_tag    = "template_http_policy"
  http_policy_match {
    type          = "cookie"
    match_type    = "contains"
    match_string  = "www.example.com"
    service_group = "sg_http"
  }
  geo_location_match {
    geo_location               = "geo_test"
    geo_location_service_group = "sg_http"
  }
  geo_location_match {
    geo_location               = "geo_http_policy"
    geo_location_service_group = "sg_http"
  }
}