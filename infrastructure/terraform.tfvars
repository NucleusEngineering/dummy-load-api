authorized_source_ranges    = ["0.0.0.0/0"]
vpc_name                    = "custom"
project_id                  = "dh-cloudsql-demos"
resource_env_label = "dummy-load-api"
services_gke_config = {
  cluster_name = "dummy-workload-cluster"
  location     = "europe-west3"
  resource_labels = {
    environment = "staging"
  }
}

# GCP APIs to activate
gcp_project_services = [
  "clouddeploy.googleapis.com",
  "cloudbuild.googleapis.com",
  "compute.googleapis.com",
  "container.googleapis.com",
  "gameservices.googleapis.com",
  "artifactregistry.googleapis.com",
  "spanner.googleapis.com",
  "secretmanager.googleapis.com",
  "servicenetworking.googleapis.com",
  "servicecontrol.googleapis.com",
  "run.googleapis.com",
  "orgpolicy.googleapis.com",
  "redis.googleapis.com",
  "run.googleapis.com",
  "iap.googleapis.com"
]
