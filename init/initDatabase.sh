
## 该脚本创建数据库
CREATE USER 'pengwl'@'%' IDENTIFIED BY '123456';
GRANT ALL ON *.* TO 'pengwl'@'%'

grant all privileges on *.* to pengwl@'%' identified by '123456';
flush privileges;