# GoVM
go实现的虚拟机

代码来自 《自己动手写java虚拟机》。这本书不错，内容比较详实，缺点是用go语言实现，对初学者不太友好，这本书并非Go语言入门书籍。<br/>

全书读下来收获颇多，这本书主要内容集中在Class文件执行，以及一些java的特性，对细节描述不多，对编译器和垃圾收集器没有任何描述。<br/>

我本人也看了一些hotspot的源码，感觉入门JVM，首先要对C++有比较深的了解，尤其是JVM实现多线程，需要用C++处理一些信号，其他书籍我自己也翻了一下，比较推荐这么几本：<br/>
1.周志明的 《深入理解java虚拟机》<br/>
2.java虚拟机规范 （这本书有1.7和1.8的版本，读者根据使用的虚拟机自行购买，可以说，《自己动手写java虚拟机》这本书就是虚拟机规范的go语言翻译）
3.实战java虚拟机<br/>

<br/>

代码已经基本都写完了，但是由于书上是按照知识体系来划分目录的，和实际逻辑上有出入，将来拉个分支，会把目录重新整合一下。<br/>

书中内容加上了许多我自己读写过程中，认识的不太全面或者不太懂的，或者亲身验证，加上了自己理解的地方，如果不对希望大家多多指正，提个issue。<br/>