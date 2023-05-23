- 数据类型
 - string (set,get,incr) key -> value, 缓存,计数器,分布式锁
 - hash(hset,hget)  key -> (field,value), (field,value)  - 存储用户信息
 - list(lpop, rpop, lrange 队列) - 发布订阅,消息队列
 - set (sadd,sismember,sinter) - 无序集合,用于用户画像,共同爱好,共同好友,共同股票
 - sorted set(zadd) - O(logn) 有序集合,用于排行榜,积分
 