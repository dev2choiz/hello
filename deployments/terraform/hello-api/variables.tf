variable "app_name" {
  description = "application name"
  default = "hello-api"
}

variable "project_id" {
  description = "project id"
}

variable "cluster_name" {
  description = "cluster name"
}

variable "region" {
  description = "region"
}

variable "zone" {
  description = "zone"
}

variable "gke_num_nodes" {
  description = "number of gke nodes"
}

variable "machine_type" {
  description = "type of node machines"
}

variable "disk_size" {
  description = "size of disk"
}

variable "namespace" {
  description = "cluster namespace"
}

variable "environment" {
  description = "environment"
}

variable "domain" {
  description = "domain"
}

variable "kms_key_ring_hello" {
  description = "Hello Kms keyring"
}

variable "kms_key_hello" {
  description = "Hello kms key"
}

variable "private_key_path" {
  description = "path to dump jenkins key"
  default = "/var/www/gcp/install/files"
}
