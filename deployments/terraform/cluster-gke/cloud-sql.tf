
resource "google_sql_database_instance" "cloud-sql-instance" {
    name   = "hello-db-instance"
    region = var.region
    database_version = "POSTGRES_13"
    settings {
        tier = "db-f1-micro"
    }

    deletion_protection  = "true"
}

resource "google_sql_database" "hello-db" {
    name     = "hello-db"
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
