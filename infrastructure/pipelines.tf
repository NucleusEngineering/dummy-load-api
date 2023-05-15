// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

resource "google_clouddeploy_target" "services_deploy_target" {
  location    = var.services_gke_config.location
  name        = "dummy-services-target"
  description = "Dummy Services Deploy Target"

  gke {
    cluster = data.google_container_cluster.services-gke.id
  }

  project          = var.project_id
  require_approval = false

  labels = {
    "environment" = var.resource_env_label
  }

  depends_on = [google_project_service.project]
}

resource "google_clouddeploy_delivery_pipeline" "services_pipeline" {
  location = var.services_gke_config.location
  name     = "dummy-services"

  description = "Dummy Services Pipeline"

  project = var.project_id

  labels = {
    "environment" = var.resource_env_label
  }

  serial_pipeline {
    stages {
      target_id = google_clouddeploy_target.services_deploy_target.target_id
    }
  }
}

resource "google_clouddeploy_target" "locust_services_deploy_target" {
  location    = var.services_gke_config.location
  name        = "locust-services-target"
  description = "Locust Services Deploy Target"

  gke {
    cluster = data.google_container_cluster.services-gke.id
  }

  project          = var.project_id
  require_approval = false

  labels = {
    "environment" = var.resource_env_label
  }

  depends_on = [google_project_service.project]
}

resource "google_clouddeploy_delivery_pipeline" "locust-services_pipeline" {
  location = var.services_gke_config.location
  name     = "locust-services"

  description = "Locust Services Pipeline"

  project = var.project_id

  labels = {
    "environment" = var.resource_env_label
  }

  serial_pipeline {
    stages {
      target_id = google_clouddeploy_target.locust_services_deploy_target.target_id
    }
  }
}
