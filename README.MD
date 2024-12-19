# SMS Backend
---- 短信任务管理平台 ----  

项目使用了group语句 需要确保数据库的查询模式开启 用一下命令配置mysql
```
SET GLOBAL sql_mode=(SELECT REPLACE(@@sql_mode,'ONLY_FULL_GROUP_BY',''));
```
Background  
Overview  
Overall Architecture!
![SMS Project - SMS Management Controller.jpg](doc%2Fimg%2FSMS%20Project%20-%20SMS%20Management%20Controller.jpg)
![SMS Project - Frame 1.jpg](./doc/img/SMS%20Project%20-%20Frame%201.jpg)