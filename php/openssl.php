<?php
$username = "root";
$password = "123456";
$arr = [$username, $password];
$x = openssl_encrypt(json_encode($arr), 'AES-128-ECB', 'ABC123', OPENSSL_RAW_DATA);
echo $x , PHP_EOL;
$token = base64_encode($x);

$y = openssl_decrypt(base64_decode($token), 'AES-128-ECB', 'ABC123', OPENSSL_RAW_DATA);
var_dump(json_decode($y));

