provider "google" {
  credentials = file("access.json")
  project     = var.project_id
  region      = var.region
}
