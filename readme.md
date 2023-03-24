# A postgreSql Cli tool
[![imi License](https://img.shields.io/badge/license-MIT-green)](https://github.com/xuejiazhi/pgii/blob/main/LICENSE)
English | [简体中文](./readme_cn.md)| [Help Document](https://github.com/xuejiazhi/pgii/wiki/pgii--postgreSql-Cli--Help-document)

pgii is a PostgreSql cli tool. PostgreSql is developed in CMD or Golang and can be compiled for multiple platforms：

- **cross-platform**： Can be compiled under multiple platforms, cross-platform use；

- **Zero-cost learning**：Similar to the MySQL Cli command, familiar with the mysql operation of the people on the hand；

- **Interactive Console**: Through the console command line。 

# Login
**cmd**:
~~~C
  ./pgii [-h|--host] [-u|--user] [-p|--password] [-d|--db] [--port]<br/>
         [-h|--host]  Database address  # eg: -h localhost | --host=localhost<br/>
         [-u|--user] Database user  # eg: -u postgres | --user=postgres<br/>
         [-p|--password]  Database password  # eg: -p postgres | --password=postgres<br/>
         [-d|--db] select database Default:postgres # eg: -d postgres | --db=postgres<br/>
         [--port] Specified port # eg: --port=5432<br/>
~~~
  #### example
~~~C
  $ pgii -h 127.0.0.1 -u postgres -p 123456 
    Connect Pgsql Success Host 127.0.0.1
    PostgresSql Version: 14.5
    pgii~[postgres/]#
~~~

# Related instruction
## use instruction
### use <db|database> <dbName>
  ***FUNCTION：***<br/>
 > Used to select a database. After data is selected, you can run "show db" or "show selectdb" to view the selected database.<br/>

  ***USAGE：***<br/>
> pgii~[postgres/]# use db benchmark<br/>
> pgii~[postgres/]# use database benchmark
>>    Use Database Success!<br/>

### use <sc|schema> <schemaName>

***FUNCTION：***<br/>
> Used to select a database schema. After selecting a schema, you can run the "show sc" or "show schema" command to view the selected schema

***USAGE：***<br/>
>   pgii~[benchmark/]# use sc  public;
>	pgii~[benchmark/]# use schema public;
>     # Use Schema Success!

## show instruction
### show <db|database>
```bash
  FUNCTION：
    Used to view information about the database, including the currently selected library and the size of the library
  USAGE
   pgii~[postgres/]# show database;
   pgii~[postgres/]# show db;
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
  FUNCTION：
     Used to view schema information about the database, including the selected schema
  USAGE
   pgii~[postgres/]# show schema;
   pgii~[postgres/]# show sc;
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

### show <tb|table> [filter|equal] [value]
```bash
  FUNCTION：
    This command is used to view information about tables in the database. If "filter" is used, the TABLENAME records containing value are filtered,and "equal" is all equal
  USAGE
   pgii~[postgres/]# show table;
   pgii~[postgres/]# show tb;
+--------+-----------+------------+------------+
| SCHEMA | TABLENAME | TABLEOWNER | TABLESPACE |
+--------+-----------+------------+------------+
| public | tags      | postgres   | <nil>      |
| public | cpu       | postgres   | <nil>      |
+--------+-----------+------------+------------+
pgii~[benchmark/public]# show tb filter c
+--------+-----------+------------+------------+
| SCHEMA | TABLENAME | TABLEOWNER | TABLESPACE |
+--------+-----------+------------+------------+
| public | cpu       | postgres   | <nil>      |
+--------+-----------+------------+------------+
```

### show <vw|view> [filter|equal] [value]
```bash
  FUNCTION：
    This command is used to view view information about the database. The "filter" command is used to filter records whose VIEWNAME contains a value. "equal" is equal to all values
  USAGE
   pgii~[postgres/]# show view;
   pgii~[postgres/]# show vw;
+--------+----------+-----------+
| SCHEMA | VIEWNAME | VIEWOWNER |
+--------+----------+-----------+
| public | cpu_view | postgres  |
+--------+----------+-----------+
pgii~[benchmark/public]# show tb filter c
+--------+----------+-----------+
| SCHEMA | VIEWNAME | VIEWOWNER |
+--------+----------+-----------+
| public | cpu_view | postgres  |
+--------+----------+-----------+
```

### show <tg|trigger> [filter|equal] [value]
```bash
  FUNCTION：
    This command is used to view information about triggers in the database. If "filter" is used, you can filter the records that contain a value in the trigger. "equal" is all equal
  USAGE
   pgii~[postgres/]# show tg;
   pgii~[postgres/]# show trigger;
+-----------+--------+-------------------+--------------------+--------------------+--------------------+---------------+
| DATABASE  | SCHEMA | TRIGGER_NAME      | EVENT_MANIPULATION | EVENT_OBJECT_TABLE | ACTION_ORIENTATION | ACTION_TIMING |
+-----------+--------+-------------------+--------------------+--------------------+--------------------+---------------+
| benchmark | public | ts_insert_blocker | INSERT             | cpu                | ROW                | BEFORE        |
+-----------+--------+-------------------+--------------------+--------------------+--------------------+---------------+
pgii~[benchmark/public]# show trigger filter ts;
+-----------+--------+-------------------+--------------------+--------------------+--------------------+---------------+
| DATABASE  | SCHEMA | TRIGGER_NAME      | EVENT_MANIPULATION | EVENT_OBJECT_TABLE | ACTION_ORIENTATION | ACTION_TIMING |
+-----------+--------+-------------------+--------------------+--------------------+--------------------+---------------+
| benchmark | public | ts_insert_blocker | INSERT             | cpu                | ROW                | BEFORE        |
+-----------+--------+-------------------+--------------------+--------------------+--------------------+---------------+
```

### show <ver|version>
```bash
  FUNCTION：
    Used to view the version information about the database
  USAGE
   pgii~[postgres/]# show version;
   pgii~[postgres/]# show ver;
+-------------+---------+
| #           | VERSION |
+-------------+---------+
| PostgresSql | 14.5    |
+-------------+---------+
```

### show <sd|selectdb>
```bash
  FUNCTION：
    Used to view the currently selected database and schema for the database
  USAGE
   pgii~[benchmark/public]# show sd;
   pgii~[benchmark/public]# show selectdb;
     DataBase: benchmark ;Schema: public
```

## desc instruction
### desc <tableName>
```bash
  FUNCTION：
     Used to view the table structure
  USAGE
    pgii~[benchmark/public]# desc cpu;
+----+------------------+-------------+--------+--------+--------------+
| #  | COLUMN           | DATATYPE    | LENGTH | ISNULL | DEFAULTVALUE |
+----+------------------+-------------+--------+--------+--------------+
| 1  | time             | timestamptz | <nil>  | NO     | <nil>        |
| 2  | tags_id          | int4        | <nil>  | YES    | <nil>        |
| 3  | hostname         | text        | <nil>  | YES    | <nil>        |
| 4  | usage_user       | float8      | <nil>  | YES    | <nil>        |
| 5  | usage_system     | float8      | <nil>  | YES    | <nil>        |
| 6  | usage_idle       | float8      | <nil>  | YES    | <nil>        |
| 7  | usage_nice       | float8      | <nil>  | YES    | <nil>        |
| 8  | usage_iowait     | float8      | <nil>  | YES    | <nil>        |
| 9  | usage_irq        | float8      | <nil>  | YES    | <nil>        |
| 10 | usage_softirq    | float8      | <nil>  | YES    | <nil>        |
| 11 | usage_steal      | float8      | <nil>  | YES    | <nil>        |
| 12 | usage_guest      | float8      | <nil>  | YES    | <nil>        |
| 13 | usage_guest_nice | float8      | <nil>  | YES    | <nil>        |
| 14 | additional_tags  | jsonb       | <nil>  | YES    | <nil>        |
+----+------------------+-------------+--------+--------+--------------+
```

## size instruction
### size <db|database> <dbName>
```bash
  FUNCTION：
     Used to view the size of the database
  USAGE
    pgii~[benchmark/public]# size database benchmark;
    pgii~[benchmark/public]# size db benchmark;
┌───────────┬─────────┐
│ DATABASE  │ SIZE    │
├───────────┼─────────┤
│ benchmark │ 3370 MB │
└───────────┴─────────┘
```

### size <tb|table> <tableName>
```bash
  FUNCTION：
     Used to view the size of a database table
  USAGE
  pgii~[benchmark/public]# size table cpu;
  pgii~[benchmark/public]# size tb cpu;
┌───────────┬───────┐
│ TABLENAME │ SIZE  │
├───────────┼───────┤
│ cpu       │ 32 kB │
└───────────┴───────┘
```

## ddl instruction
### ddl <tb|table> <tableName>
```bash
  FUNCTION：
     A ddl construct clause for viewing the table
  USAGE
  pgii~[benchmark/public]# ddl table cpu;
  pgii~[benchmark/public]# ddl tb cpu;
========= Create Table Success ============
-- DROP Table;
-- DROP Table cpu;
CREATE TABLE "public".cpu (
    time timestamptz NOT NULL,
    tags_id int4 NULL,
    hostname text NULL,
    usage_user float8 NULL,
    usage_system float8 NULL,
    usage_idle float8 NULL,
    usage_nice float8 NULL,
    usage_iowait float8 NULL,
    usage_irq float8 NULL,
    usage_softirq float8 NULL,
    usage_steal float8 NULL,
    usage_guest float8 NULL,
    usage_guest_nice float8 NULL,
    additional_tags jsonb NULL
);
CREATE INDEX cpu_usage_user_time_idx ON public.cpu USING btree (usage_user, "time" DESC);
CREATE INDEX cpu_time_idx ON public.cpu USING btree ("time" DESC);
CREATE INDEX cpu_hostname_time_idx ON public.cpu USING btree (hostname, "time" DESC);
```

### ddl <sc|schema> <schemaName>
```bash
  FUNCTION：
     A ddl constructor clause for viewing the schema
  USAGE
  pgii~[benchmark/public]# ddl schema public;
  pgii~[benchmark/public]# ddl sc public;
========= Create Schema Success ============
-- DROP SCHEMA public;
CREATE SCHEMA "public" AUTHORIZATION postgres;
```

### ddl <vw|view> <viewName>
```bash
  FUNCTION：
     The ddl constructor clause for viewing the view
  USAGE
  pgii~[benchmark/public]# ddl view cpu_view;
  pgii~[benchmark/public]# ddl vw cpu_view;
========= Create View Success ============
 CREATE OR REPLACE VIEW "public".cpu_view
 AS SELECT cpu."time",
    cpu.tags_id,
    cpu.hostname,
    cpu.usage_user,
    cpu.usage_system,
    cpu.usage_idle,
    cpu.usage_nice,
    cpu.usage_iowait,
    cpu.usage_irq,
    cpu.usage_softirq,
    cpu.usage_steal,
    cpu.usage_guest,
    cpu.usage_guest_nice,
    cpu.additional_tags
   FROM cpu
  WHERE ((cpu."time" > '2023-03-01 08:00:00+08'::timestamp with time zone) AND (cpu.tags_id > 10) AND (cpu.tags_id < 1000) AND (cpu.usage_user = ANY (ARRAY[(21)::double precision, (22)::double precision, (23)::double precision, (24)::double precision, (25)::double precision, (26)::double precision, (27)::double precision, (28)::double precision, (29)::double precision])));
```
## dump 指令
### dump <tb|table> <tableName>
```bash
  FUNCTION：
     dump the backup file of a table, which can be used for subsequent restoration；
  USAGE
    pgii~[clouddb/common]# dump tb role;
      Dump Table Success
    ## View in linux
    [root@localhost src]# ls *.pgi
      dump_table_role_time.pgi
```

### dump <sc|schema>
```bash
  FUNCTION：
     dump schema creation statements and table creation statements of the current schema and the following tables, and generate T-SQL statements inserted in batches from the data in the table to generate a pgi file；
  USAGE
      pgii~[clouddb/db_mcs.com]# dump sc;
        Dump Schema Success [db_mcs.com]
        Dump Table Struct Success [dgna]
         ->Dump Table Record Success [dgna]
        Dump Table Struct Success [dgna_member]
         ->Dump Table Record Success [dgna_member]
        Dump Table Struct Success [syspatch_info]
         ->Dump Table Record Success [syspatch_info]
        Dump Table Struct Success [syspatch_member]
         ->Dump Table Record Success [syspatch_member]
        Dump Table Struct Success [predefineddgna_info]
         ->Dump Table Record Success [predefineddgna_info]
        Dump Table Struct Success [simulselect_info]
         ->Dump Table Record Success [simulselect_info]
        Dump Table Struct Success [simulselect_member]
         ->Dump Table Record Success [simulselect_member]
    ## linux下查看
    [root@localhost src]# ls *.pgi
       dump_schema_db_mcs.com.pgi
```