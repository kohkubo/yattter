-- DB一覧
show databases;

-- DBを指定
use db_name;

-- テーブル一覧
show tables;

-- テーブルのカラム一覧
show columns from table_name;

-- テーブル内一覧
select * from table_name;
select * from account;

-- データ挿入
insert into table_name (column_name) values (value);
