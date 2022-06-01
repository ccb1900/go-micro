# 微服务

- sms
    - 阿里
    - 腾讯
    - local
- email
    - smtp
    - local
- 存储
    - 阿里
    - 腾讯
    - 七牛
    - s3
    - local
- 配置中心
- 注册中心
- 缓存
    - 文件
    - redis
- 平滑退出
- 熔断
- 限流
- 定时任务
- 支持 grpc
- 支持 web
- 日志
    - zap
    - logrus
- 链路追踪
    - zipkin
    - jaeger
- 数据库
    - mysql
    - postgresql
- 搜索
    - es
- 可观测性
- 消息队列
    - redis
    - rocketmq
    - rabbitmq
    - kafka
- 命令行
    - server
- 支付
    - 微信
    - 支付宝

## 功能

- 支持自定义命令
- 自定义缓存
- 自定义消息队列
- 自定义数据库
- grpc
- proto
- dockerfile

## 命令行一键生成项目

## 默认配置

## 完整配置

```yaml
# 是否开启
enabled:
  mq: false
  sms: false
  email: false
  cache: false
  db: false
# 日志配置
log:
  default: local
  drivers:
    - driver: zap
      debug: false
# 消息队列配置
mq:
  default: rabbitmq
  drivers:
    - driver: zap
      debug: false
# 数据库
db:
  default: rabbitmq
  drivers:
    - driver: zap
      debug: false
# 短信
sms:
  default: rabbitmq
  drivers:
    - driver: zap
      debug: false
email:
  default: rabbitmq
  drivers:
    - driver: zap
      debug: false
registry:
  default: rabbitmq
  drivers:
    - driver: zap
      debug: false
config:
  default: rabbitmq
  drivers:
    - driver: zap
      debug: false
cache:
  default: rabbitmq
  drivers:
    - driver: zap
      debug: false
```

## 内置微服务

- 邮件
- 短信
- 存储
- 缓存
- 支付

## 参考
