# DBA - Database Anonymizer

Tool for anonymizing database records

*work in progress*

## Installation

For now, you need Go and the [glide](https://github.com/Masterminds/glide) build to install dba. Downloadable binaries will be available soon.

``` shell
λ git clone github.com/rpip/dba
λ cd dba && make deps && make
λ ./bin/dba --help
usage: dba [<flags>] <conf>

Anonymize database records.

Flags:
  -h, --help     Show context-sensitive help (also try --help-long and --help-man).
  -v, --verbose  Verbose mode.
      --version  Show application version.

Args:
  <conf>  Config file.

```

## Usage

DBA uses the HCL config language from HashiCorp to describe database parameters and table updates. It's the same config language used in products like Terraform. It's quite flexible, easy to read and as such makes the config seem like a little DSL. It supports basic data types and ternary operations. You can [read more here](https://www.terraform.io/docs/configuration/index.html).

``` hcl
db "library" {
  type = "mysql"
  dsn = "dba:dba123@(:3306)/dbatest?charset=utf8&parseTime=True&loc=Local"
  verbose = true

  table "user" {
    // defaults to 'id'
    primary_key = "user_id"

    updates {
      first_name = "${first_name()}"
      last_name = "${first_name()}"
      username = "${username()}"
      bio = "${paragraph()}"
      age = "${digits_n(2) * 2}"
      gender = "${gender_abbrev()}"
      role = "${row.role == '0' ? 8 : 30}"
    }
  }

  table "product" {

    updates {
      name = "${product()}"
      price = "${digits_n(5)}"
      merchant = "${company()}"
      brand = "${brand()}"
    }

  }
}
```

```shell
λ ./bin/dba sample.conf.hcl
```

## TODO
- [] test suite
- [] more fake data generators, eg: date, time
- [] support sqlite, Postgres
- [] atomic operation:
   ```hcl
   atomic {
     field = value
     field = value
  }```
- [] Progress indictaor
- [] Batch updates: build a long query string to run all table updates in one go
- [] documentation

## Development

Pull requests are welcome. Please make sure the build succeeds and the test suite passes. Also see `HACKING.md` for info on the internals of dba.

You can also build a Docker image to test this:

```shell
λ make docker
λ docker run -i -t dba sample.conf.hcl
```
