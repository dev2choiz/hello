variable "app_name" {
  description = "application name"
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
variable "front_cr_sub_domain" {
  description = "front cloud run sub domain"
}
variable "front_gke_sub_domain" {
  description = "front gke sub domain"
}

variable "certPrivateKeyPath" {
  description = "private key path"
  type = string
}

variable "certCrtPath" {
  description = "crt path"
  type = string
}

variable "kms_key_ring_hello" {
  description = "Hello Kms keyring"
}

variable "kms_key_hello" {
  description = "Hello kms key"
}

variable "dir_output" {
    description = "dir_output"
}
