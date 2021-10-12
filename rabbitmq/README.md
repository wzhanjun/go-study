## Go RabbitMQ

#### 

```
1. 简单模型
   生产者往队列添加消息，消费者从队列消费消息
2. worker-queue
   生产者往队列添加消息，有多个消费者消费消息
3. pub/sub
   交换机type==fanout
4. routing
   交换机type==direct, 指定routingkey
5. topic
   交换机type==topic, 指定topic路由routingkey
6. rpc

```

#### 参考

- https://blog.51cto.com/u_14286115/3326255
- https://zhuanlan.zhihu.com/p/79545722
- https://www.cnblogs.com/hello-/articles/10345021.html
- https://www.jianshu.com/p/04f443dcd8bd
- https://www.cnblogs.com/lonely-wolf/p/14225132.html
- https://segmentfault.com/a/1190000014862123
- https://github.com/php-amqplib/php-amqplib/issues/444#issuecomment-261169320