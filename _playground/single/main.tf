
resource "google_compute_forwarding_rule" "default" {
  project               = var.project
  name                  = var.name
  region                = var.region
  data                  = data.foo
  custom                = custom_type.bar
  hardcoded             = "Hello world"
  timeout_sec         = var.health_check["timeout_sec"]
  example_bracket = {
    
  }
}

