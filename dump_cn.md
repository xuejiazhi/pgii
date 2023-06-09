# pgii 备份导出与导入功能
pgii 可以做为是一个PostgreSql的一个导入导出工具,可以对千万数据的大表进行备份导入导出;

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
***FUNCTION：***<br/>
>   将指定的表导出对应的pgi压缩文件

***用法：***<br/>
~~~C
pgi~[yc/pgii]# dump tb t_test;
 Dump Table Success [t_test].....
~~~
