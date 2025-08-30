<?php
/**
 * 单例模式，是一种常用的软件设计模式。在它的核心结构中只包含一个被称为单例的特殊类。通过单例模式可以保证系统中一个类只有一个实例。即一个类只有一个对象实例。
*要点主要有三个：
*1.一个类只能有一个对象；
*2.必须是自行创建这个类的对象；
*3，要想整个系统提供这一个对象；
*从具体实现角度来说，就是以下三点：
*一是单例模式的类只提供私有的构造函数，
*二是类定义中含有一个该类的静态私有对象，
*三是该类提供了一个静态的公有的函数用于创建或获取它本身的静态私有对象。
*当然还要有一个private的clone方法，防止克隆；
*优点
*一、实例控制
*单例模式会阻止其他对象实例化其自己的单例对象的副本，从而确保所有对象都访问唯一实例。
*二、灵活性
*因为类控制了实例化过程，所以类可以灵活更改实例化过程。
* 
*缺点
*一、开销
*虽然数量很少，但如果每次对象请求引用时都要检查是否存在类的实例，将仍然需要一些开销。可以通过使用静态初始化解决此问题。
*二、可能的开发混淆
*使用单例对象（尤其在类库中定义的对象）时，开发人员必须记住自己不能使用new关键字实例化对象。因为可能无法访问库源代码，因此应用程序开发人员可能会意外发现自己无法直接实例化此类。
*三、对象生存期
*不能解决删除单个对象的问题。在提供内存管理的语言中（例如基于.NET Framework的语言），只有单例类能够导致实例被取消分配，因为它包含对该实例的私有引用。在某些语言中（如 C++），其他类可以删除对象实例，但这样会导致单例类中出现悬浮引用
 */
class Singleton {
    private static $instance;
    private $args;

    private  function __construct($args){
        $this->args = $args;
    }

    private function __clone(){

    }

    public static function getInstance($args){
        if(!self::$instance instanceof self){
            self::$instance = new self($args);
        }
        return self::$instance;
    }

    public function getArgs(){
        return $this->args;
    }
}

//test

$a = Singleton::getInstance(['a'=>1]);
echo var_dump($a);