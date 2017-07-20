# actor

组成：State,Behavior,Mailbox,Children,Supervisor Strategy
封装在Actor Reference中

### 1.Actor Reference

actor对象需要从外部屏蔽，actor通过Actor Reference表示外部，Actor Reference传递象。从不同的应用向Actor发送信息，Actor不向外部发送信息，外部无法掌控信息

### 2.State（状态）

actor对象通常会包含一些反映actor可能处于状态的变量，这些数据要避免其它actor的影响，Proto.Actor在概念上都有自己的轻量级线程，它完全与系统的其他部分屏蔽。这意味着，不必使用锁同步访问，您只需编写您的actor代码，而不必担心并发性,Proto.Actor将在一组真正的线程上运行一组actor，通常许多actor共享一个线程，并且一个actor的后续调用可能最终在不同的线程上被处理。 Proto.Actor确保这个实现细节不会影响处理演员状态的单线程。
内部状态对于演员的操作至关重要，所以状态不一致是致命的,actor失败重新启动会从头开始创建，持续接受数据在重新启动后重新播放来恢复到重启之前

### 3.Behavior(行为)

行为是指定义在该时间点对消息作出反应的动作的function,行为可能会随时间而变化，重新启动actor会将其行为重置为该初始行为

#### 4.Mailbox(邮箱)

actor的目的是处理消息，并将这些消息从其他actor（或从actor系统外部）发送给actor,连接发送者和接收者的部分是actor的邮箱,每一个actor有一个邮箱，线程之间分布参与者的明显随机性，从同一个演员发送多个消息到同一个目标，将以相同的顺序排列它们，由actor处理的邮件的顺序与排入队列的顺序相匹配，处理特殊优先级顺序的信息由队列的算法定义Proto.Actor与其他actor模型实现不同的一个重要特征是当前行为必须始终处理下一个出队消息，因此没有扫描下一个匹配邮件的邮箱

### 5.children(子actor)

创建子actor,actor本身就会自动监督他们，同时可以使用他们，实际的创建和终止操作以异步方式发生在幕后，

创建（Context.ActorOf（...））或停止（Context.Stop（child））

### 6.Supervisor Strategy(主管策略)

处理子actor故障的策略，actor一旦创建，strategy不能被更改，每个actor只有一个这样的策略，适用于各种子actor

### 7.终止

一旦演员终止，即以不重新启动处理的方式发生故障，将自动停止或由其主管停止，它将释放其资源，将所有剩余的邮件从其邮箱中排入系统的“死信邮箱”
邮箱通过actor reference被系统邮箱所代替