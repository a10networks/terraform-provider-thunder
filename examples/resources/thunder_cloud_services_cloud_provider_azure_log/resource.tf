provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_cloud_services_cloud_provider_azure_log" "test" {
  action = "enable"
  customer_id = "b15e06fe-45d0-4822-8029-9ba14ac745b4"
  shared_key = "xuedNH81Yu0OZ7hFySQwZbS5uI+5Y9Aimp9E3RhqAvWdCYZQWn4GmvNvOvb8W02rCCoFt6oCHb40d9HStqRhWtQkypoPeTTqZ6zBkZ1oNTGog+LVNwH4r5uYtczmvi7x"
}
