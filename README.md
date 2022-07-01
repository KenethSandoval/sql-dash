# TUIDB

_A database administrator only for consultation and change of user with privileges and information of tables of the db_


![Tab user](https://github.com/KenethSandoval/sql-dash/blob/main/docs/assets/user-tab.png)

![Tab table](https://github.com/KenethSandoval/sql-dash/blob/main/docs/assets/table-tab.png)
## Starting
_These instructions will allow you to get a working copy of the project on your local machine for development and testing purposes._

```sh
$ git clone https://github.com/KenethSandoval/sql-dash
```

_run docker for test_
```sh
$ make up
```

_make build compile_
```sh
$ make compile
```

_clean docker and binary sqldash_
```sh
$ make clean
```

## Configuration

A config file will be generated when you first run `tuidb`. Depending on your operating system it can be found in one of the following locations:

* macOS: ~/Library/Application\ Support/tuidb/config.yml
* Linux: ~/.config/tuidb/config.yml
* Windows: C:\Users\me\AppData\Roaming\tuidb\config.yml

It will include the following default settings:

```yml
settings:
    username: root
    password: root
    database: testdas
```
