# pgii 备份导出与导入功能
pgii 可以做为是一个PostgreSql的一个导入导出工具,可以对百万级,千万数据的大表进行备份导入导出;

## postgreSql下生成测试数据
**表结构和插入数据的语句，5000000条**
~~~C
CREATE TABLE t_test(
    ID INT PRIMARY KEY   NOT NULL,
    NAME      TEXT  NOT NULL,
    AGE      INT   NOT NULL,
    ADDRESS    CHAR(50),
    SALARY     REAL
);

insert
    into
    t_test
select
    generate_series(1, 5000000) as key,
    repeat( chr(int4(random()* 26)+ 65), 4),
    (random()*(6 ^2))::integer,
    null,
    (random()*(10 ^4))::integer;
~~~

## 查看生成的数据
~~~C
pgi~[yc/pgii]# show tb;
+--------+-----------+------------+------------+-----------+-----------+
| SCHEMA | TABLENAME | TABLEOWNER | TABLESPACE | TABLESIZE | INDEXSIZE |
+--------+-----------+------------+------------+-----------+-----------+
| pgii   | t_test    | postgres   | <nil>      | 356 MB    | 107 MB    |
+--------+-----------+------------+------------+-----------+-----------+

pgi~[yc/pgii]# select count(*) from t_test;
+---------+
|   COUNT |
+---------+
| 5000000 |
+---------+
[Total: 1 Rows]  [RunTimes 0.20s]
~~~

## DUMP 指令
- **dump table**： 导出表的结构和数据，对大于50000条数据已上的表会进行分段处理,按50000条数据生成一个批量导入的文件✅;
- **dump schema**：可以导出选中模式的创建语句，选中模式下的表的结构与数据,对于单表大于50000条数据的表会进行分段处理,按50000条数据生成一个批量导入的文件✅;
- **dump database**：todo(开发中)

### dump <tb|table> <tableName>
***说明：***<br/>
>   将指定的表导出对应的pgi压缩文件

***用法：***<br/>
~~~C
pgi~[yc/pgii]# dump tb t_test;
 Dump Table Success [t_test].....
~~~
如下图,在命令所在的文件夹下面生成了一个 dump_table_pgii_1686301979 的文件夹,文件夹下面是t_test表生成的相关导出备份文件 
![image](https://github.com/xuejiazhi/pgii/assets/16795993/2a2a6490-19fe-4fb1-ad36-6a8042f38e63)

### dump <sc|schema>
***说明：***<br/>
>   将指定模式下的表和表数据等导出对应的pgi压缩文件,在使用这个指令时需要选中对应的模式,否则无法完成;

***用法：***<br/>
~~~C
pgi~[yc/pgii]# show tb;
+--------+-----------+------------+------------+-----------+-----------+
| SCHEMA | TABLENAME | TABLEOWNER | TABLESPACE | TABLESIZE | INDEXSIZE |
+--------+-----------+------------+------------+-----------+-----------+
| pgii   | t_test    | postgres   | <nil>      | 356 MB    | 107 MB    |
+--------+-----------+------------+------------+-----------+-----------+
| pgii   | t_user    | postgres   | <nil>      | 128 kB    | 40 kB     |
+--------+-----------+------------+------------+-----------+-----------+

pgi~[yc/pgii]# dump sc;
 Dump Schema Success [pgii]
 Dump Table Success [t_test].....
 Dump Table Success [t_user].....
~~~
如下图,在命令所在的文件夹下面生成了一个 dump_schema_pgii_1686302845 的文件夹,文件夹下面是pgii这个模式下面相关的表生成的pgi导出缩文件
 ![image](https://github.com/xuejiazhi/pgii/assets/16795993/0160ecb9-dd7c-4764-b151-a490d6c292c8)


## Load 指令
- **Load table**： 导出表的结构和数据，对大于50000条数据已上的表会进行分段处理,按50000条数据生成一个批量导入的文件✅;
- **Load schema**：可以导出选中模式的创建语句，选中模式下的表的结构与数据,对于单表大于50000条数据的表会进行分段处理,按50000条数据生成一个批量导入的文件✅;
- **Load database**：todo(开发中)

### load <tb|table> <table backup filepath>
***说明：***<br/>
>   将指定的表备份的文件路径下的表的结构与表数据导入到一个新的模式下面或者其数据库的模式下面,必须选中模式。

**用法：**<br/>
- **选中另外一个数据xc和模式pgii**
~~~C
pgi~[yc/]# use db xc;
  Use Database Success!
pgi~[xc/]# use sc pgii;
  Use Schema Success!
~~~
- **开始load先前dump tb出来的备份文件夹名称 dump_table_pgii_1686301979**
~~~C
pgi~[xc/pgii]# load tb dump_table_pgii_1686301979;
Load Table Success [dump_table_pgii_1686301979/tb_t_test.pgi] Affect Nums:50000
Load Table Success [dump_table_pgii_1686301979/tb_t_test_1.pgi] Affect Nums:50000
Load Table Success [dump_table_pgii_1686301979/tb_t_test_2.pgi] Affect Nums:50000
Load Table Success [dump_table_pgii_1686301979/tb_t_test_3.pgi] Affect Nums:50000
Load Table Success [dump_table_pgii_1686301979/tb_t_test_4.pgi] Affect Nums:50000
Load Table Success [dump_table_pgii_1686301979/tb_t_test_5.pgi] Affect Nums:50000
Load Table Success [dump_table_pgii_1686301979/tb_t_test_6.pgi] Affect Nums:50000
.
.
.
Load Table Success [dump_table_pgii_1686301979/tb_t_test_99.pgi] Affect Nums:50000
Load Table Success [dump_table_pgii_1686301979/tb_t_test_100.pgi] Affect Nums:0
~~~

- **查询表如下，表已经成功的导入过来了**
~~~C
pgi~[xc/pgii]# show tb;                    
+--------+-----------+------------+------------+-----------+-----------+
| SCHEMA | TABLENAME | TABLEOWNER | TABLESPACE | TABLESIZE | INDEXSIZE |
+--------+-----------+------------+------------+-----------+-----------+
| pgii   | t_test    | postgres   | <nil>      | 356 MB    | 107 MB    |
+--------+-----------+------------+------------+-----------+-----------+
pgi~[xc/pgii]# select count(*) from t_test;        
+---------+
|   COUNT |
+---------+
| 5000000 |
+---------+
[Total: 1 Rows]  [RunTimes 0.20s]
~~~

### load <sc|schema> <schema backup filepath>
***说明：***<br/>
>   将指定的模式备份的文件路径下的模式和所有表的结构与表数据导入到其它数据库下面,。

**用法：**<br/>
- **选中另外一个数据xc**
~~~C
pgi~[yc/]# use db xc;
  Use Database Success!

~~~
- **开始load先前dump sc出来的备份文件夹名称 dump_schema_pgii_1686302845**
~~~C
pgi~[xc/pgii]# load sc dump_schema_pgii_1686302845;
Load Table Success [dump_table_pgii_1686301979/tb_t_test.pgi] Affect Nums:50000
Load Table Success [dump_table_pgii_1686301979/tb_t_test_1.pgi] Affect Nums:50000
Load Table Success [dump_table_pgii_1686301979/tb_t_test_2.pgi] Affect Nums:50000
Load Table Success [dump_table_pgii_1686301979/tb_t_test_3.pgi] Affect Nums:50000
Load Table Success [dump_table_pgii_1686301979/tb_t_test_4.pgi] Affect Nums:50000
Load Table Success [dump_table_pgii_1686301979/tb_t_test_5.pgi] Affect Nums:50000
Load Table Success [dump_table_pgii_1686301979/tb_t_test_6.pgi] Affect Nums:50000
.
.
.
Load Table Success [dump_table_pgii_1686301979/tb_t_test_99.pgi] Affect Nums:50000
Load Table Success [dump_table_pgii_1686301979/tb_t_test_100.pgi] Affect Nums:0
Load Table Success [dump_schema_pgii_1686302845/tb_t_user.pgi] Affect Nums:1000
~~~

- **查询表如下，表已经成功的导入过来了**
~~~C
pgi~[xc/pgii]# show tb;                    
+--------+-----------+------------+------------+-----------+-----------+
| SCHEMA | TABLENAME | TABLEOWNER | TABLESPACE | TABLESIZE | INDEXSIZE |
+--------+-----------+------------+------------+-----------+-----------+
| pgii   | t_test    | postgres   | <nil>      | 356 MB    | 107 MB    |
+--------+-----------+------------+------------+-----------+-----------+
| pgii   | t_user    | postgres   | <nil>      | 128 kB    | 40 kB     |
+--------+-----------+------------+------------+-----------+-----------+

pgi~[xc/pgii]# select count(*) from t_test;        
+---------+
|   COUNT |
+---------+
| 5000000 |
+---------+
[Total: 1 Rows]  [RunTimes 0.20s]

pgi~[xc/pgii]# select count(*) from t_user;
+-------+
| COUNT |
+-------+
|  1000 |
+-------+
[Total: 1 Rows]  [RunTimes 0.10s]
~~~