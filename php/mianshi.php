<?php
//fwrite(STDOUT, "pleace : ");
//$x = fread(STDIN, 1000);
//var_dump($x);


$x = [10,8,5,3,2,9,8];
print_r(selectSort($x));
print_r(bubbleSort($x));

// 选择
function selectSort($arr){
    static $result = [];
    $rodKey = 0;
    $rodVal = $arr[$rodKey];
    for ($i=0; $i<count($arr); $i++) {
        if ($arr[$i] > $rodVal) {
            list($rodKey, $rodVal) = [$i, $arr[$i]];
        }
    }
    array_splice($arr, $rodKey, 1);
    array_unshift($result, $rodVal);
    if ($arr) selectSort($arr);
    return $result;
}

// 冒泡
function bubbleSort($arr){
    for ($c=0;$c<count($arr);$c++){
       for ($i=0;$i<count($arr)-1;$i++){
            if ($arr[$i] > $arr[$i+1]) {
                $bak = $arr[$i+1];
                list($arr[$i+1], $arr[$i]) = [$arr[$i], $bak];
            }
        }
    }
    return $arr;
}

// 插入
