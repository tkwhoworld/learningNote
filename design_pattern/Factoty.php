<?php
class A{
    public function hello(){
        echo 'nihao woshi a'.PHP_EOL;
    }
}

class B{
    public function hello(){
        echo 'nihao woshi b'.PHP_EOL;
    }
}

class C{
    public function hello(){
        echo 'nihao woshi c'.PHP_EOL;
    }
};



class simpleFactory{
    public static function getInstance($type){
        switch ($type){
            case 1:
                return new A();
            case 2:
                return new B();
            case 3:
                return new C();
            default:
                 return new A();
        }
    }
}

$a=simpleFactory::getInstance(1);
$a->hello();
$b=simpleFactory::getInstance(2);
$b->hello();
$c=simpleFactory::getInstance(3);
$c->hello();