---
title: Postgres environment variable configuration
sidebarTitle: "Postgres"
category: "envvars"
index: 6
---

## Postgres environment variable configuration

Environment variables available to the Nacelle process change the behavior of Postgres interactions.

---

{{< table table_class="table table-striped table-bordered" >}}
| Environment Variable            | Required | Default           | Description                                                                                          |
| ------------------------------- | -------- | ----------------- | ---------------------------------------------------------------------------------------------------- |
| DATABASE_URL                    | yes      |                   | The connection string of the remote database.                                                        |
| LOG_SQL_QUERIES                 |          | false             | Whether or not to log parameterized SQL queries.                                                     |
| MIGRATIONS_TABLE                |          | schema_migrations | The name of the migrations table.                                                                    |
| MIGRATIONS_SCHEMA_NAME          |          | default           | The name of the schema used during migrations.                                                       |
| FAIL_ON_NEWER_MIGRATION_VERSION |          | false             | If true, fail startup when the database migration version is newer than the known set of migrations. |
{{< /table >}}
