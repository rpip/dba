[![Build Status](https://travis-ci.org/rpip/dba.svg?branch=master)](https://travis-ci.org/rpip/dba)

# DBA - Database Anonymizer

#### Tool for anonymizing database records.

Sometimes you need a production data dump to run performance tests,  see how a production release goes and debug production issues. It is however not easy getting production data for many reasons such as privacy concerns for user data, tool to actually anonymize the data among others. Here comes, this data anonymization tool.

*work in progress*

## Installation

For now, you need Go and the [glide](https://github.com/Masterminds/glide) build tool on your system to get this running. Downloadable binaries will be available soon.

``` shell
λ git clone github.com/rpip/dba
λ cd dba && make deps && make
λ ./dba --help
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

DBA uses the HCL config language from HashiCorp to describe database parameters and table updates. It's the same config language used in products like Terraform. It's quite flexible, easy to read and as such makes the config seem like a little DSL. It supports basic data types and ternary operations. You can [learn more about HCL here](https://www.terraform.io/docs/configuration/index.html).

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
      admin = "${row.age > 18 ? 1 : 0}"
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
```

```shell
λ ./dba sample.conf.hcl
```

## Development

Pull requests are welcome. Please make sure the build succeeds and the test suite passes. Also see `HACKING.md` for info on the internals of dba.

You can also build a Docker image to test this:

```shell
λ make docker
λ docker run -i -t dba test-fixtures/sample.conf.hcl
```
