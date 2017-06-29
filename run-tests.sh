#!/bin/bash

set -ev

mysql -u root < test-fixtures/shop.sql

mysql -u root < test-fixtures/blog.sql

./dba test-fixtures/sample.conf.hcl
