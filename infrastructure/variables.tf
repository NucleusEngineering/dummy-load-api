variable "region" {
   type		= string
   default	= "europe-west3"
}

variable "project_id" {
   type		= string
   default	= "dh-cloudsql-demos"
}

variable "vpc_name" {
   default	= "custom"
   type		= string
   description	= "Name of the VPC"
}

variable "authorized_source_ranges" {
   type		= list(string)
   description	= "Addresses or CIDR blocks which are allowed to connect to GKE API Server."
}

variable "gke_master_ipv4_cidr_block" {
  type    = string
  default = "172.23.0.0/28"
}

variable "gcp_project_services" {
  type        = list(any)
  description = "GCP Service APIs (<api>.googleapis.com) to enable for this project"
  default     = []
}

variable "services_gke_config" {
  type = object({
    cluster_name    = string
    location        = string
    resource_labels = map(string)
  })

  description = "Configuration specs for GKE Autopilot cluster that hosts all backend services"
}

variable "resource_env_label" {
  type        = string
  description = "Label/Tag to apply to resources"
}
