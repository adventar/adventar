Create databases.

```
$ mysql -u root -e 'create database adventar_dev'
$ mysql -u root -e 'create database adventar_test'
```

Apply the `schema.sql` using [schemalex](https://github.com/schemalex/schemalex).

```
$ schemalex 'mysql://root@tcp(127.0.0.1:3306)/adventar_dev' schema.sql | mysql -u root adventar_dev
$ schemalex 'mysql://root@tcp(127.0.0.1:3306)/adventar_test' schema.sql | mysql -u root adventar_test
```
