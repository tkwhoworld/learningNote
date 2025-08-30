<?php

class Rigester{
    public static $tree =[];
    public static function set($alias,$obj){
        self::$tree[$alias] = $obj;
    }

    public static function get($name){
        if(self::$tree[$name]){
            return self::$tree[$name];
        }
    }

    public static function clearObj($name){
        unset(self::$tree[$name]);
    }
}