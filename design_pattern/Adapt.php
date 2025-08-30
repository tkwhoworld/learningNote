<?php

interface IDataBase{
    public function connection();
    public function query();
    public function close();
}

class MySql implements IDataBase{
    public function connection(){

    }
    public function query(){

    }
    public function close(){

    }
}

class Mongo implements IDataBase{
    public function connection(){

    }
    public function query(){

    }
    public function close(){
        
    }
}

class redis implements IDataBase{
    public function connection(){

    }
    public function query(){

    }
    public function close(){
        
    }
}