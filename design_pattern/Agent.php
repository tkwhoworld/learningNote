<?php

abstract class baseAgent{
    abstract function printPage();
}

class IAgent extends baseAgent{
    public function printPage(){
        echo 'hello print';
    }
}

class BAgent extends baseAgent{
    public function printPage(){
        echo 'hello print2';
    }
}

class CAgent extends baseAgent{
    public function printPage(){
        echo 'hello print2';
    }
}

class Browser { //具体策略角色
    public function call($object) {
        return $object->PrintPage ();
    }
}

$bro = new Browser ();
echo $bro->call ( new IAgent () );