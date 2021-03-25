<?php

gd_info();
p(getimagesize("./image.png"));
p(image_type_to_extension(IMAGETYPE_PNG));
p(image_type_to_mime_type(IMAGETYPE_PNG));

$file = 'php.jpg';
$image = imagecreatefrompng($file);
header('Content-type: ' . image_type_to_mime_type(IMAGETYPE_WBMP));
image2wbmp($image); // output the stream directly
imagedestroy($image);

function p($p){
    print_r($p);
    echo "\n";
}

