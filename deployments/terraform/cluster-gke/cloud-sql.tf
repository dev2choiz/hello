
resource "google_sql_database_instance" "cloud-sql-instance" {
    name   = var.sql_instance
    region = var.region
    database_version = "POSTGRES_13"
    settings {
        tier = "db-f1-micro"
        ip_configuration {
            require_ssl = true
        }
    }

    deletion_protection  = "true"
}

resource "google_sql_ssl_cert" "sql_client_cert" {
    common_name = "sql-client-cert"
    instance    = google_sql_database_instance.cloud-sql-instance.name
}

resource "kubernetes_secret" "sql_client_cert_secret" {
    metadata {
        name      = "sql-client-cert-secret"
        namespace = var.namespace
    }
    data = {
        "sql-client-cert.pem" = google_sql_ssl_cert.sql_client_cert.cert
    }
}

resource "local_file" "client_cert_local_file" {
    filename = "${var.dir_output}/tls/sql_client_cert.pem"
    sensitive_content = google_sql_ssl_cert.sql_client_cert.cert
}

resource "google_sql_database" "hello-db" {
    name     = var.db_name
    instance = google_sql_database_instance.cloud-sql-instance.name
}

resource "google_sql_user" "sql-user" {
    name     = var.sql_user
    password = var.sql_password
    instance = google_sql_database_instance.cloud-sql-instance.name
}
resource "kubernetes_secret" "sql-user-credentials" {
    metadata {
        name      = "sql-user-credentials"
        namespace = var.namespace
    }
    data = {
        "username" = var.sql_user
        "password" = google_sql_user.sql-user.password
    }
}
