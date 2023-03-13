一种postgreSql的Cli工具

pgii 是一个PostgreSql cli的工具,对PostgreSql 在CMD或者,采用Golang进行开发,可以多平台下面编译使用：

- **跨平台**： 可以在多平台下编译，跨平台使用；

- **零学习成本**：类似于MySQL Cli的指令,对熟悉mysql操作的人上手快；

- **互动 Console**: 通过命令行 console。 

# 登录
**cmd**: 
```bash
  ./pgii [-h|--host] [-u|--user] [-p|--password] [-d|--db] [--port]
      [-h|--host] postgresql 数据库地址  # eg: -h localhost | --host=localhost
      [-u|--user] 数据库用户名  # eg: -u postgres | --user=postgres
      [-p|--password]  数据库密码  # eg: -p postgres | --password=postgres
      [-d|--db] 选择的数据库 默认为postgres # eg: -d postgres | --db=postgres
      [--port] 指定的端口 # eg: --port=5432

  #### 示例
  $ pgii -h 127.0.0.1 -u postgres -p 123456 
    Connect Pgsql Success Host 127.0.0.1
    PostgresSql Version: 14.5
    pgii~[postgres/]#
```

# 相关指令
## show 指令
### show <db|database>
```bash
  功能：
    用于查看数据库的相关信息,包括当前选中的库,以及库的大小
  用法
   pgii~[postgres/]# show database
   pgii~[postgres/]# show db
+-------+-------------+----------+----------+------------+----------+-----------+-----------+------------+------------+---------+
| #OID  | DBNAME      | AUTH     | ENCODING | LC_COLLATE | LC_CTYPE | ALLOWCONN | CONNLIMIT | LASTSYSOID | TABLESPACE | SIZE    |
+-------+-------------+----------+----------+------------+----------+-----------+-----------+------------+------------+---------+
| 13892 | postgres[✓] | postgres | UTF8     | C          | C        | true      |        -1 | 13891      | pg_default | 8537 kB |
| 16384 | clouddb     | postgres | UTF8     | C          | C        | true      |        -1 | 13891      | pg_default | 18 MB   |
| 1     | template1   | postgres | UTF8     | C          | C        | true      |        -1 | 13891      | pg_default | 8385 kB |
| 13891 | template0   | postgres | UTF8     | C          | C        | false     |        -1 | 13891      | pg_default | 8385 kB |
| 91966 | benchmark   | postgres | UTF8     | C          | C        | true      |        -1 | 13891      | pg_default | 3370 MB |
+-------+-------------+----------+----------+------------+----------+-----------+-----------+------------+------------+---------+
```

### show <sc|schema>
```bash
  功能：
    用于查看数据库的相关模式信息,包括当前选中的模式
  用法
   pgii~[postgres/]# show schema
   pgii~[postgres/]# show sc
┌───────┬──────────────────────────┬──────────┬─────────────────────────────────────┐
│ #OID  │ SCHEMANAME               │ OWNER    │ ACL                                 │
├───────┼──────────────────────────┼──────────┼─────────────────────────────────────┤
│ 99    │ pg_toast                 │ postgres │ <nil>                               │
│ 11    │ pg_catalog               │ postgres │ {postgres=UC/postgres,=U/postgres}  │
│ 2200  │ public[✓]                │ postgres │ {postgres=UC/postgres,=UC/postgres} │
│ 13526 │ information_schema       │ postgres │ {postgres=UC/postgres,=U/postgres}  │
│ 91989 │ _timescaledb_cache       │ postgres │ {postgres=UC/postgres,=U/postgres}  │
│ 91987 │ _timescaledb_catalog     │ postgres │ {postgres=UC/postgres,=U/postgres}  │
│ 91988 │ _timescaledb_internal    │ postgres │ {postgres=UC/postgres,=U/postgres}  │
│ 91990 │ _timescaledb_config      │ postgres │ {postgres=UC/postgres,=U/postgres}  │
│ 91992 │ timescaledb_information  │ postgres │ {postgres=UC/postgres,=U/postgres}  │
│ 91991 │ timescaledb_experimental │ postgres │ {postgres=UC/postgres,=U/postgres}  │
└───────┴──────────────────────────┴──────────┴─────────────────────────────────────┘
```

### show <tb|table>
```bash
  功能：
    用于查看数据库的相关表信息
  用法
   pgii~[postgres/]# show table
   pgii~[postgres/]# show tb
+--------+-----------+------------+------------+
| SCHEMA | TABLENAME | TABLEOWNER | TABLESPACE |
+--------+-----------+------------+------------+
| public | tags      | postgres   | <nil>      |
| public | cpu       | postgres   | <nil>      |
+--------+-----------+------------+------------+
```

### show <ver|version>
```bash
  功能：
    用于查看数据库的相关版本信息
  用法
   pgii~[postgres/]# show version
   pgii~[postgres/]# show ver
+-------------+---------+
| #           | VERSION |
+-------------+---------+
| PostgresSql | 14.5    |
+-------------+---------+
```

### show <sd|selectdb>
```bash
  功能：
    用于查看数据库的当前选中的database 和schema
  用法
   pgii~[benchmark/public]# show sd
   pgii~[benchmark/public]# show selectdb
     DataBase: benchmark ;Schema: public
```