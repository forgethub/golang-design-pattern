# 【工厂模式】心得

## 这个模式核心是什么？其实讲的是什么？

定义一个创建对象的接口，但让子类决定实例化哪一个类。工厂方法模式允许类的实例化推迟到子类

## 给我什么启发？学到了什么？

写代码的过程要解耦，工厂方法模式通过将对象的创建过程抽象化，使得客户端代码与具体的产品类解耦。

## 我以前写过的什么代码场景也许可以用这个重构？大概思路？

1. 需要隐藏对象的创建细节的场景。
2. 需要根据不同格式数据读取配置文件
