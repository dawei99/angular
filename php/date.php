<?php
$now = strtotime("2021-04-02");
echo date("Y-m-d H:i:s", mktime(0,0,0,date('m'), date('d'), date('y'))) , PHP_EOL;
echo date("Y-m-d H:i:s", mktime(0,0,0,date('m'), date('d')+1, date('y'))  -1) , PHP_EOL;

$week = date('w') ? date('w', $now)-1 : 7-1;
$diff = 7 - date('w', $now);
echo date("Y-m-d H:i:s", mktime(0,0,0,date('m',$now), date('d',$now)-$week, date('y',$now))) , PHP_EOL;
echo date("Y-m-d H:i:s", mktime(0,0,0,date('m',$now), date('d',$now)+$diff, date('y',$now))) , PHP_EOL;




