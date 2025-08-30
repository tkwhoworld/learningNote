<?php
class Subject{
    private $observer =[];
    public function register($observer){
        $this->observer[] = $observer;
    }

    public function trigger($data){
        foreach($this->observer as $key=>$observer){
            $observer->update($data);
        }
    }
}

class subscribe {
    public function update($data){
        echo 'the is update 1 hello'.PHP_EOL;
    }
}

class subscribe2 {
    public function update($data){
        echo 'the is update 2 hello'.PHP_EOL;
    }
}

//test
$a = new Subject();
$a->register(new subscribe());
$a->register(new subscribe2());
$a->trigger('hello');