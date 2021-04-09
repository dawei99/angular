<?php
$content = "abc收到是";
//iconv("utf8", "gbk", $content);
mb_convert_encoding($content, "utf8", "gbk");
var_dump($content);
var_dump(strrev("ABC"));