<?php
$curl = curl_init("http://127.0.0.1:2035/api/archiveTransit.php");
curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
curl_setopt($curl, CURLOPT_POST, 1);
$data = json_encode(['name' => '物资系统归档文件名称']);
curl_setopt($curl, CURLOPT_POSTFIELDS, $data);
$resp = curl_exec($curl);

echo($resp);
