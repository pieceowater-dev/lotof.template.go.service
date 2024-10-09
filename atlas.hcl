data "external_schema" "gorm" {
  program = ["go", "run", "./internal/core/db/migrate.go"]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/16/dev"
  migration {
    dir = "file://internal/core/db/migrations"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}