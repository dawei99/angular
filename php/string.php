<?php
set_error_handler(function($code, $err, $file, $line){
    echo "-----------";
    echo "${file}:${line}【${err}】";
    echo "---------", PHP_EOL;
});
try {
    var_dump(strlen("abc是"));
    var_dump(mb_strlen("abc是"));
    var_dump(gethostbyname("www.baidu.com"));
    var_dump($_SERVER['REMOTE_ADDR']);

    var_dump(substr_count("asd", "f"));
    var_dump(ucwords("ABC", "B"));

    $x = 1;
    $y = 2;
    list($x, $y) = [$y , $x];
//    echo $x, PHP_EOL, $y;
} catch (Exception $e) {
    print_r($e);
} catch (Error $e) {
    print_r($e);
}
