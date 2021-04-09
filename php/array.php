
<?php
$x = ["name" => "xiaoming"];
//$x = array_map(function(&$v){
//    return "1";
//}, $x);

array_walk($x, function(&$v){
    $v = "1";
});
var_dump($x);

