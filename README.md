# go-sweet-generator
# go-sweet与go-sweet-vue3的代码生成器
# go-sweet代码生成器已完成, go-sweet-vue3代码生成器正在开发中....

# 项目架构
```text
该代码生成器项目架构与SpringBoot项目类似
配置文件在resources目录下
需要将数据库配置改为正确的配置
然后将moduleName.common和moduleName.src改为你自己的模块名称
tableName是一个String列表 , 将你想要生成的表名放入其中
然后运行代码即可将代码结构生成到generatorFile目录下
```