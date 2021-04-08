<?php
$x = ["21153","19159","21911",10];
$y = "group[1:2]group[4:5]group[10:6]";
$result = null;
var_dump(preg_match_all("/group\[(\d+):(\d+)\]/", $y, $result, PREG_SET_ORDER));
print_r($result);