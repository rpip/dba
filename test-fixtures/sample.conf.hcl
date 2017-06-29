db "store" {
  type = "mysql"
  dsn = "root:@(:3306)/dbastore?charset=utf8&parseTime=True&loc=Local"
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
  dsn = "root:@(:3306)/dbablog?charset=utf8&parseTime=True&loc=Local"
  verbose = true

  table "post" {
    title = "${title()}"
    body = "${paragraphs_n(5)}"
    published = true
  }
}
