<?php
error_reporting(E_ALL);
ini_set("display_errors", 1);


class myException extends Exception {
    public function __toString()
    {
        return "犯错了！！！";
    }
}

try {
    set_error_handler(function($no, $error, $file, $line){
        throw new Exception("no: $no\n err: $error\n file:$file \n line:$line\n");
    });
    //require_once "test.php";
    //throw new Exception('style');
    $a = $y;
} catch (myException $e) {
    echo $e;
} catch (Exception $e) {
    echo $e->getMessage();
} catch (Error $error) {
    print_r($error);
}
