![image](https://user-images.githubusercontent.com/16795993/227710095-913f63a0-5969-4158-8e3a-48a055ab3401.png)

# 一个postgreSql的Cli工具
[![imi License](https://img.shields.io/badge/license-MIT-green)](https://github.com/xuejiazhi/pgii/blob/main/LICENSE)
简体中文 | [English](./readme.md) | [帮助文档](https://github.com/xuejiazhi/pgii/wiki/pgii-%E4%B8%AD%E6%96%87%E5%B8%AE%E5%8A%A9%E6%96%87%E6%A1%A3)<br/>
pgii 是一个PostgreSql cli的工具,对PostgreSql 在CMD或者,采用Golang进行开发,可以多平台下面编译使用：

- **跨平台**： 可以在多平台下编译，跨平台使用；
- **零学习成本**：类似于MySQL Cli的指令,对熟悉mysql操作的人上手快；
- **互动 Console**: 通过命令行 console。 
- **增加了交互式提示**: 可以提示要输入的命令
 

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

# 设置语言
### set language <cn|en>
***功能：***<br/>
>   使用些命令设置语言后,指令的头部显示可以是中文或英文

**使用方式：**<br/>
~~~C
pgii~[postgres/]# show db;
+--------+-------------+----------+----------+------------+----------+-----------+-----------+------------+------------+---------+
| #OID   | DBNAME      | AUTH     | ENCODING | LC_COLLATE | LC_CTYPE | ALLOWCONN | CONNLIMIT | LASTSYSOID | TABLESPACE | SIZE    |
| 13891  | template0   | postgres | UTF8     | C          | C        | false     |        -1 | 13891      | pg_default | 8385 kB |
+--------+-------------+----------+----------+------------+----------+-----------+-----------+------------+------------+---------+
| 280382 | benchmark   | postgres | UTF8     | C          | C        | true      |        -1 | 13891      | pg_default | 10 MB   |
+--------+-------------+----------+----------+------------+----------+-----------+-----------+------------+------------+---------+
pgii~[postgres/]# set language cn;
Set Language Success
pgii~[postgres/]# show db;
+--------+-------------+--------------+----------+------------+----------+----------+----------------+-----------------+------------+------------+
| #OID   | 数据库      | 数据库拥用者 | 字符编码 | LC_COLLATE | LC_CTYPE | 允许连接 | 最大并发连接数 | 最后一个系统OID | 表空间     | 数据库尺寸 |
+--------+-------------+--------------+----------+------------+----------+----------+----------------+-----------------+------------+------------+
| 13892  | postgres[✓] | postgres     | UTF8     | C          | C        | true     |             -1 | 13891           | pg_default | 8537 kB    |
+--------+-------------+--------------+----------+------------+----------+----------+----------------+-----------------+------------+------------+
| 16384  | clouddb     | postgres     | UTF8     | C          | C        | true     |             -1 | 13891           | pg_default | 25 MB      |
+--------+-------------+--------------+----------+------------+----------+----------+----------------+-----------------+------------+------------+
| 1      | template1   | postgres     | UTF8     | C          | C        | true     |             -1 | 13891           | pg_default | 8385 kB    |
+--------+-------------+--------------+----------+------------+----------+----------+----------------+-----------------+------------+------------+
| 13891  | template0   | postgres     | UTF8     | C          | C        | false    |             -1 | 13891           | pg_default | 8385 kB    |
+--------+-------------+--------------+----------+------------+----------+----------+----------------+-----------------+------------+------------+
| 280382 | benchmark   | postgres     | UTF8     | C          | C        | true     |             -1 | 13891           | pg_default | 10 MB      |
+--------+-------------+--------------+----------+------------+----------+----------+----------------+-----------------+------------+------------+
~~~

# 相关指令
## use 指令
### use <db|database> <dbName>
```bash
  功能：
    用于选择数据库,选中数据后，可以使用show db 或 show selectdb 查看当前选中的数据库
  用法：
    pgii~[postgres/]# use db benchmark
	pgii~[postgres/]# use database benchmark
      # Use Database Success!
    pgii~[benchmark/]#
```

### use <sc|schema> <schemaName>
```bash
  功能：
    用于选择数据库模式,选中模式后，可以使用show sc 或 show schema 查看当前选中的模式
  用法：
    pgii~[benchmark/]# use sc  public;
	pgii~[benchmark/]# use schema public;
     # Use Schema Success!
    pgii~[benchmark/public]#
```

## show 指令
### show <db|database>
```bash
  功能：
    用于查看数据库的相关信息,包括当前选中的库,以及库的大小
  用法
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
  功能：
    用于查看数据库的相关模式信息,包括当前选中的模式
  用法
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
  功能：
    用于查看数据库的相关表信息,使用filter,可以过滤TABLENAME包含value的记录，equal 为全等于
  用法
   pgii~[postgres/]# show table;
   pgii~[postgres/]# show tb;
+--------+-----------+------------+------------+-----------+-----------+
| SCHEMA | TABLENAME | TABLEOWNER | TABLESPACE | TABLESIZE | INDEXSIZE |
+--------+-----------+------------+------------+-----------+-----------+
| public | tags      | postgres   | <nil>      | 400 kB    | 264 kB    |
+--------+-----------+------------+------------+-----------+-----------+
| public | cpu       | postgres   | <nil>      | 32 kB     | 24 kB     |
+--------+-----------+------------+------------+-----------+-----------+
pgii~[benchmark/public]# show tb filter c
+--------+-----------+------------+------------+-----------+-----------+
| SCHEMA | TABLENAME | TABLEOWNER | TABLESPACE | TABLESIZE | INDEXSIZE |
+--------+-----------+------------+------------+-----------+-----------+
| public | cpu       | postgres   | <nil>      | <nil>     | <nil>     |
+--------+-----------+------------+------------+-----------+-----------+
```

### show <vw|view> [filter|equal] [value]
```bash
  功能：
    用于查看数据库的相关视图信息,使用filter,可以过滤VIEWNAME包含value的记录，equal 为全等于
  用法
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
  功能：
    用于查看数据库的相关触发器信息,使用filter,可以过滤触发器包含value的记录，equal 为全等于
  用法
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
  功能：
    用于查看数据库的相关版本信息
  用法
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
  功能：
    用于查看数据库的当前选中的database 和schema
  用法
   pgii~[benchmark/public]# show sd;
   pgii~[benchmark/public]# show selectdb;
     DataBase: benchmark ;Schema: public
```

### show <connection|conn>
***功能：***<br/>
> 用于查看链接的一些参数
> MAX_CONNECTION  最大连接数
> SUPERUSER_RESERVED_CONNECTIONS 超级用户保留的连接数
> REMAINING_CONNECTIONS 剩余连接数
> INUSE_CONNECTIONS 当前正使用的连接数

***方法：***<br/>
~~~C
pgii~[benchmark/public]# show connection;
+----------------+--------------------------------+-----------------------+-------------------+
| MAX_CONNECTION | SUPERUSER_RESERVED_CONNECTIONS | REMAINING_CONNECTIONS | INUSE_CONNECTIONS |
+----------------+--------------------------------+-----------------------+-------------------+
|            800 |                             13 |                   760 |                40 |
+----------------+--------------------------------+-----------------------+-------------------+
~~~

### show <process|proc> [all]|[pid start and end]
***功能：***<br/>
> 用于查看当前数据库中的会话:
> all: 如果没有带 all 参数,那么只显示当前选择数据库下面的会话，带了all显示所有的会话;
> pid: 列出pid 在 start 和 end 区间内的会话


***方法：***<br/>
~~~C
pgii~[postgres/]# show process;
+--------+---------------+-----------+---------------+-------------+--------------------------------------+--------+
| PID    | DATABASE_NAME | USER_NAME | CLIENT_ADDR   | CLIENT_PORT | APPLICATION_NAME                     | STATE  |
+--------+---------------+-----------+---------------+-------------+--------------------------------------+--------+
| 142668 | postgres      | postgres  | 10.161.55.214 | 56766       | DBeaver 22.0.2 - Main <postgres>     | idle   |
+--------+---------------+-----------+---------------+-------------+--------------------------------------+--------+
| 142669 | postgres      | postgres  | 10.161.55.214 | 56767       | DBeaver 22.0.2 - Metadata <postgres> | idle   |
+--------+---------------+-----------+---------------+-------------+--------------------------------------+--------+
| 206841 | postgres      | postgres  | 10.161.55.214 | 42969       |                                      | active |
+--------+---------------+-----------+---------------+-------------+--------------------------------------+--------+
pgii~[postgres/]# show process all; 
+--------+---------------+-----------+-----------------+-------------+--------------------------------------------+--------+
| PID    | DATABASE_NAME | USER_NAME | CLIENT_ADDR     | CLIENT_PORT | APPLICATION_NAME                           | STATE  |
+--------+---------------+-----------+-----------------+-------------+--------------------------------------------+--------+
| 88720  |               |           |                 |             |                                            |        |
+--------+---------------+-----------+-----------------+-------------+--------------------------------------------+--------+
| 117720 | clouddb       | postgres  | 100.123.237.153 | 38390       |                                            | idle   |
+--------+---------------+-----------+-----------------+-------------+--------------------------------------------+--------+
| 90141  | clouddb       | postgres  | 100.123.237.196 | 54220       |                                            | idle   |
+--------+---------------+-----------+-----------------+-------------+--------------------------------------------+--------+
| 142668 | postgres      | postgres  | 10.161.55.214   | 56766       | DBeaver 22.0.2 - Main <postgres>           | idle   |
+--------+---------------+-----------+-----------------+-------------+--------------------------------------------+--------+
| 142669 | postgres      | postgres  | 10.161.55.214   | 56767       | DBeaver 22.0.2 - Metadata <postgres>       | idle   |
+--------+---------------+-----------+-----------------+-------------+--------------------------------------------+--------+
| 142670 | clouddb       | postgres  | 10.161.55.214   | 56768       | DBeaver 22.0.2 - SQLEditor <Script-24.sql> | idle   |
+--------+---------------+-----------+-----------------+-------------+--------------------------------------------+--------+
| 142671 | clouddb       | postgres  | 10.161.55.214   | 56769       | DBeaver 22.0.2 - Main <clouddb>            | idle   |
+--------+---------------+-----------+-----------------+-------------+--------------------------------------------+--------+
| 142672 | clouddb       | postgres  | 10.161.55.214   | 56770       | DBeaver 22.0.2 - Metadata <clouddb>        | idle   |
+--------+---------------+-----------+-----------------+-------------+--------------------------------------------+--------+
| 206841 | postgres      | postgres  | 10.161.55.214   | 42969       |                                            | active |
+--------+---------------+-----------+-----------------+-------------+--------------------------------------------+--------+
| 33133  |               |           |                 |             |                                            |        |
+--------+---------------+-----------+-----------------+-------------+--------------------------------------------+--------+
| 33135  |               |           |                 |             |                                            |        |
+--------+---------------+-----------+-----------------+-------------+--------------------------------------------+--------+
pgii~[postgres/]# show process pid 1 and 88585; 
+-------+---------------+-----------+-----------------+-------------+------------------+-------+
| PID   | DATABASE_NAME | USER_NAME | CLIENT_ADDR     | CLIENT_PORT | APPLICATION_NAME | STATE |
+-------+---------------+-----------+-----------------+-------------+------------------+-------+
| 88585 | clouddb       | postgres  | 100.123.237.162 | 54040       |                  | idle  |
+-------+---------------+-----------+-----------------+-------------+------------------+-------+
| 88583 | clouddb       | postgres  | 100.123.237.150 | 58828       |                  | idle  |
+-------+---------------+-----------+-----------------+-------------+------------------+-------+
| 88581 |               | postgres  |                 |             |                  |       |
+-------+---------------+-----------+-----------------+-------------+------------------+-------+
| 88579 | clouddb       | postgres  | 100.123.237.196 | 35001       |                  | idle  |
+-------+---------------+-----------+-----------------+-------------+------------------+-------+
| 88578 | clouddb       | postgres  | 100.123.237.148 | 48182       |                  | idle  |
+-------+---------------+-----------+-----------------+-------------+------------------+-------+
| 88577 | clouddb       | postgres  | 100.123.237.149 | 44360       |                  | idle  |
+-------+---------------+-----------+-----------------+-------------+------------------+-------+
| 88576 | clouddb       | postgres  | 100.123.237.154 | 43840       |                  | idle  |
+-------+---------------+-----------+-----------------+-------------+------------------+-------+
| 88575 | clouddb       | postgres  | 100.123.237.164 | 57554       |                  | idle  |
+-------+---------------+-----------+-----------------+-------------+------------------+-------+
| 33135 |               |           |                 |             |                  |       |
+-------+---------------+-----------+-----------------+-------------+------------------+-------+
| 33133 |               |           |                 |             |                  |       |
+-------+---------------+-----------+-----------------+-------------+------------------+-------+
| 33134 |               |           |                 |             |                  |       |
+-------+---------------+-----------+-----------------+-------------+------------------+-------+
~~~


## desc 指令
### desc <tableName>
```bash
  功能：
     用于查看表结构
  用法
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

## size 指令
### size <db|database> <dbName>
```bash
  功能：
     用于查看数据库的大小
  用法
    pgii~[benchmark/public]# size database benchmark;
    pgii~[benchmark/public]# size db benchmark;
┌───────────┬─────────┐
│ DATABASE  │ SIZE    │
├───────────┼─────────┤
│ benchmark │ 3370 MB │
└───────────┴─────────┘
```

### size <tb|table> <tableName>
***功能：***<br/>
> 用于查看数据库表的大小

***方法：***<br/>
~~~C
  pgii~[benchmark/public]# size table cpu;
  pgii~[benchmark/public]# size tb cpu;
┌───────────┬───────┐
│ TABLENAME │ SIZE  │
├───────────┼───────┤
│ cpu       │ 32 kB │
└───────────┴───────┘
~~~

### size <tbsp|tablespace> <tableSpaceName>
***功能：***<br/>
>     用于查看表空间的大小

***方法：***<br/>
~~~C
pgii~[benchmark/public]# size tablespace pg_default;
+-----------------+-----------------+
| TABLESPACE_NAME | TABLESPACE_SIZE |
+-----------------+-----------------+
| pg_default      | 60 MB           |
+-----------------+-----------------+
~~~

## ddl 指令
### ddl <tb|table> <tableName>
```bash
  功能：
     用于查看表的ddl建表语句
  用法
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
  功能：
     用于查看模式的ddl建表语句
  用法
  pgii~[benchmark/public]# ddl schema public;
  pgii~[benchmark/public]# ddl sc public;
========= Create Schema Success ============
-- DROP SCHEMA public;
CREATE SCHEMA "public" AUTHORIZATION postgres;
```

### ddl <vw|view> <viewName>
```bash
  功能：
     用于查看视图的ddl建表语句
  用法
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
## kill 指令
### kill pid <pid>
***功能：***<br/>
>   关闭数据库中的会话，关闭后可以使用“show proc”查看是否关闭

***用法：***<br/>
~~~C
pgii~[postgres/]# kill pid 33134;
Kill Process Success,pid[33134]
~~~

## dump 指令
### dump <tb|table> <tableName>
```bash
  功能：
     用于dump一个表的备份文件，可以用于后续的恢复；
  用法
    pgii~[clouddb/common]# dump tb role;
      Dump Table Success
    ## linux下查看
    [root@localhost src]# ls *.pgi
      dump_table_role_time.pgi
```

### dump <sc|schema>
```bash
  功能：
     用于dump当前模式和下面的表的建模式语句和建表语句,并将表下面数据生成批量插入的T-SQL语句，生成一个pgi文件；
  用法
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

### dump <db|database>
***功能：***<br/>
>  用于dump当前选择的数据库下面,并生成创建库,模式,表的语句，并将表里的数据生成批量插入的数据，生成一个压缩的pgi文件；

***方法：***<br/>
~~~C
pgii~[clouddb/]# dump db;
>Dump DataBase Begin
Dump DataBase Struct Success            
----------------------------------------
>>Dump Schema Success[db_mcs1.com]      
>>>Dump Table Struct Success [pub_vpn]
 >>>>Dump Table Record Success ["db_mcs1.com".pub_vpn]
>>>Dump Table Struct Success [userprofile]
 >>>>Dump Table Record Success ["db_mcs1.com".userprofile]
>>>Dump Table Struct Success [pub_groupprofile]
 >>>>Dump Table Record Success ["db_mcs1.com".pub_groupprofile]
>>>Dump Table Struct Success [mcs_groupprofile]
 >>>>Dump Table Record Success ["db_mcs1.com".mcs_groupprofile]
>>>Dump Table Struct Success [pub_groupmember]
 >>>>Dump Table Record Success ["db_mcs1.com".pub_groupmember]
>>>Dump Table Struct Success [mcs_blacklist]
 >>>>Dump Table Record Success ["db_mcs1.com".mcs_blacklist]
>>>Dump Table Struct Success [mcs_groupareaconfig]
 >>>>Dump Table Record Success ["db_mcs1.com".mcs_groupareaconfig]
>>>Dump Table Struct Success [mcs_groupgpsareaconfig]
 >>>>Dump Table Record Success ["db_mcs1.com".mcs_groupgpsareaconfig]
>>>Dump Table Struct Success [pub_ldsvpn]
 >>>>Dump Table Record Success ["db_mcs1.com".pub_ldsvpn]
>>>Dump Table Struct Success [pub_record_user]
 >>>>Dump Table Record Success ["db_mcs1.com".pub_record_user]
>>>Dump Table Struct Success [pub_record_group]
 >>>>Dump Table Record Success ["db_mcs1.com".pub_record_group]
>>>Dump Table Struct Success [pub_record_vpn]
 >>>>Dump Table Record Success ["db_mcs1.com".pub_record_vpn]
>>>Dump Table Struct Success [mcs_user_area_authorize]
 >>>>Dump Table Record Success ["db_mcs1.com".mcs_user_area_authorize]
>>>Dump Table Struct Success [pub_ldsgmk]
 >>>>Dump Table Record Success ["db_mcs1.com".pub_ldsgmk]
>>>Dump Table Struct Success [pub_grouppatch_gmk]
 >>>>Dump Table Record Success ["db_mcs1.com".pub_grouppatch_gmk]
>>>Dump Table Struct Success [kms_random]
 >>>>Dump Table Record Success ["db_mcs1.com".kms_random]
>>>Dump Table Struct Success [callrecord]
 >>>>Dump Table Record Success ["db_mcs1.com".callrecord]
>>>Dump Table Struct Success [datarecord]
 >>>>Dump Table Record Success ["db_mcs1.com".datarecord]
>>>Dump Table Struct Success [t_frs_fileinfo]
 >>>>Dump Table Record Success ["db_mcs1.com".t_frs_fileinfo]
>>>Dump Table Struct Success [ue_register_backup]
 >>>>Dump Table Record Success ["db_mcs1.com".ue_register_backup]
>>>Dump Table Struct Success [mcs_user_contact_list]
 >>>>Dump Table Record Success ["db_mcs1.com".mcs_user_contact_list]
>>>Dump Table Struct Success [fa_info]
 >>>>Dump Table Record Success ["db_mcs1.com".fa_info]
>>>Dump Table Struct Success [fa_users]
 >>>>Dump Table Record Success ["db_mcs1.com".fa_users]
>>>Dump Table Struct Success [fa_history]
 >>>>Dump Table Record Success ["db_mcs1.com".fa_history]
>>>Dump Table Struct Success [user_roles]
 >>>>Dump Table Record Success ["db_mcs1.com".user_roles]
>>>Dump Table Struct Success [group_roles]
 >>>>Dump Table Record Success ["db_mcs1.com".group_roles]
>>>Dump Table Struct Success [pcall_role_pri]
 >>>>Dump Table Record Success ["db_mcs1.com".pcall_role_pri]
>>>Dump Table Struct Success [gcall_role_pri]
 >>>>Dump Table Record Success ["db_mcs1.com".gcall_role_pri]
>>>Dump Table Struct Success [mcs_user_common_sms_list]
 >>>>Dump Table Record Success ["db_mcs1.com".mcs_user_common_sms_list]
>>>Dump Table Struct Success [tetra_gid]
 >>>>Dump Table Record Success ["db_mcs1.com".tetra_gid]
>>>Dump Table Struct Success [service_priority]
 >>>>Dump Table Record Success ["db_mcs1.com".service_priority]
----------------------------------------
>>Dump Schema Success[public]
tips:Dump Cmd Table fail,no table in Schema!
----------------------------------------
>>Dump Schema Success[common]
>>>Dump Table Struct Success [app_cfg]
 >>>>Dump Table Record Success ["common".app_cfg]
>>>Dump Table Struct Success [gw_info]
 >>>>Dump Table Record Success ["common".gw_info]
>>>Dump Table Struct Success [listening_info]
 >>>>Dump Table Record Success ["common".listening_info]
>>>Dump Table Struct Success [mcs_user_area]
 >>>>Dump Table Record Success ["common".mcs_user_area]
>>>Dump Table Struct Success [tenant_info]
 >>>>Dump Table Record Success ["common".tenant_info]
>>>Dump Table Struct Success [pstn_info]
 >>>>Dump Table Record Success ["common".pstn_info]
>>>Dump Table Struct Success [terminal_statistic]
 >>>>Dump Table Record Success ["common".terminal_statistic]
>>>Dump Table Struct Success [tetra_cfg]
 >>>>Dump Table Record Success ["common".tetra_cfg]
----------------------------------------
## linux下查看
[root@localhost src]# ls *.pgi
dump_Database_clouddb_1680227905.pgi
~~~

### explain <T-SQL>
***功能：***<br/>
>  用于分析T-SQL执行计划

***方法：***<br/>
~~~C
pgii~[clouddb/db_mcs.com]# explain ANALYZE select * from userprofile;
+---------------------------------------------------------------------------------------------------------------+
| QUERY PLAN                                                                                                    |
+---------------------------------------------------------------------------------------------------------------+
| Seq Scan on userprofile  (cost=0.00..105.41 rows=2241 width=230) (actual time=0.024..0.705 rows=2241 loops=1) |
| Planning time: 0.131 ms                                                                                       |
| Execution time: 0.882 ms                                                                                      |
+---------------------------------------------------------------------------------------------------------------+
[Total: 3 Rows]  [RunTimes 3.50s]

pgii~[clouddb/db_mcs.com]# explain  select * from userprofile;        
+------------------------------------------------------------------+
| QUERY PLAN                                                       |
+------------------------------------------------------------------+
| Seq Scan on userprofile  (cost=0.00..105.41 rows=2241 width=230) |
+------------------------------------------------------------------+
[Total: 1 Rows]  [RunTimes 2.18s]
~~~


## TODO
-  dump database ✅
-  kill pid ✅
-  show process ✅
-  load table
-  load schema
-  load database
