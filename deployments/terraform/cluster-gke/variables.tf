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

variable "domain" {
  description = "domain"
}

variable "membership_id" {
  description = "membership_id"
}

variable "dir_output" {
    description = "dir_output"
}

// database
variable "sql_instance" {
    description = "sql_instance"
}
variable "db_name" {
    description = "db_name"
}

// secrets
variable "sql_user" {
  description = "sql_password"
}

variable "sql_password" {
  description = "sql_password"
}
