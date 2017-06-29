## HACKING DBA

Below is a sample DBA config file.

```hcl
db "store" {
  type = "mysql"
  dsn = "dba:123456@(:3306)/dbastore?charset=utf8&parseTime=True&loc=Local"
  verbose = true

  table "user" {
    // defaults to 'id'
    primary_key = "user_id"

    updates {
      first_name = "${first_name()}"
      last_name = "${first_name()}"
      username = "${username()}"
      bio = "${paragraph()}"
      age = "${digits_n(2)}"
      gender = "${gender_abbrev()}"
    }
  }

  table "product" {
    updates {
      name = "${product()}"
      price = "${digits_n(3)}"
      merchant = "${company()}"
      brand = "${brand()}"
    }
  }
}


db "blog" {
  type = "mysql"
  dsn = "dba:123456@(:3306)/dbablog?charset=utf8&parseTime=True&loc=Local"
  verbose = true

  table "post" {
    title = "${title()}"
    body = "${paragraphs_n(5)}"
    published = true
  }
}
```

DBA parses the config above into an internal config format. Using a simple and declarative config format makes changes more explicit and maintainable as opposed to infering the table column types.

Below is a dump of the config data structure in Go.

```go
&dba.Config{
  EvalConfig: dba.EvalConfig{
    EvalConfig: (*hil.EvalConfig)(nil),
  },
  Databases: []*dba.Database{
    &dba.Database{
      meta: dba.meta{
        Name:     "store",
        Scope:    "database",
        Metadata: {
          "type":    "mysql",
          "dsn":     "dba:123456@(:3306)/dbastore?charset=utf8&parseTime=True&loc=Local",
          "verbose": true,
        },
      },
      Tables: []*dba.Table{
        &dba.Table{
          meta: dba.meta{
            Name:     "user",
            Scope:    "table",
            Metadata: {
              "primary_key": "user_id",
            },
          },
          Updates: {
            "gender":     "${gender_abbrev()}",
            "first_name": "${first_name()}",
            "last_name":  "${first_name()}",
            "username":   "${username()}",
            "bio":        "${paragraph()}",
            "age":        "${digits_n(2)}",
          },
        },
        &dba.Table{
          meta: dba.meta{
            Name:     "product",
            Scope:    "table",
            Metadata: map[string]interface {}{},
          },
          Updates: {
            "name":     "${product()}",
            "price":    "${digits_n(3)}",
            "merchant": "${company()}",
            "brand":    "${brand()}",
          },
        },
      },
      DB: (*sql.DB)(nil),
    },
    &dba.Database{
      meta: dba.meta{
        Name:     "blog",
        Scope:    "database",
        Metadata: {
          "verbose": true,
          "type":    "mysql",
          "dsn":     "dba:123456@(:3306)/dbablog?charset=utf8&parseTime=True&loc=Local",
        },
      },
      Tables: []*dba.Table{
        &dba.Table{
          meta: dba.meta{
            Name:     "post",
            Scope:    "table",
            Metadata: {
              "title":     "${title()}",
              "body":      "${paragraphs_n(5)}",
              "published": true,
            },
          },
          Updates: map[string]interface {}{},
        },
      },
      DB: (*sql.DB)(nil),
    },
  },
}
```

Each table implements the `Anonymizer` interface. As we iterate through the tables in the config, we call the the `Anonymize` function on each table object and it constructs an UPDATE SQL query string from the updates declared in the config. This SQL query is then run against the declared DB connection. That simple.

The lib folder contains the core of DBA. The goal is to enable usage of DBA as a library within other applications. The `main.go` script in the root folder contains the generated command line tool.
