<?php
error_reporting(0);
ini_set("display_errors", false);

class myException extends Exception {
    public function __toString()
    {
        return "犯错了！！！";
    }
}
try {
    set_error_handler(function($no, $error, $file, $line){
//        throw new Exception("no: $no\n err: $error\n file:$file \n line:$line\n");
        echo "buhuo";
        return true;
    });
    require_once "test.php";
//    throw new Exception('style');
//    $a = $y;
//    throw  new myException("aaaaa");

} catch (myException $e) {
//    var_dump($e->getTrace());
//    var_dump($e->getMessage());
//    var_dump($e->getLine());
//    var_dump($e->getFile());
} catch (Exception $e) {
//    var_dump($e);
} catch (Error $e) {
//    var_dump($e);
}
