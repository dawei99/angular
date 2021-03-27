<?php
$x = 10; // 001
$y = 11; // 101

//printf("%b", $x);
//echo "<br>";
//printf("%b", $x >> 2);
//die;
//gd_info();
//p(getimagesize("./image.png"));
//p(image_type_to_extension(IMAGETYPE_PNG));
//p(image_type_to_mime_type(IMAGETYPE_PNG));
header ('Content-Type: image/png');
$i = imagecreatetruecolor(200,200);
$color = imagecolorallocate($i, 255, 10,20);
imagefill($i, 10, 10, $color);
imagejpeg($i);
die;
//$file = 'image.png';
//$image = imagecreatefromjpeg($file);
//$color = imagecolorallocate($image, 255,10,5);
//$colorWhite = imagecolorallocate($image, 0,0,0);

$img = "./sign.jpg";
list($srcWith, $srcHeight, $type) = getimagesize($img);
$width=100;
$height=30;

header ('Content-Type: image/png');

switch($type){
    case IMAGETYPE_PNG : $gdImage = imagecreatefrompng($img); break;
    case IMAGETYPE_JPEG : $gdImage = imagecreatefromjpeg($img); break;
    case IMAGETYPE_GIF : $gdImage = imagecreatefromgif($img); break;
    default: die('无法识别类型'); break;
}

$trueColor = imagecreatetruecolor($width, $height);
imagefill($trueColor,0,0, imagecolorallocate($gdImage,255, 255, 255));
imagecopyresampled($trueColor, $gdImage, 0, 0, 0,0, $width, $height, $srcWith, $srcHeight);
imagepng($trueColor);
imagedestroy($trueColor);
imagedestroy($gdImage);
die;


//$im = imagecreatetruecolor(500, 500);
//$im2 = imagecreatefromjpeg('image.png');
//$text_color = imagecolorallocate($im, 200, 140, 91);
//imagestring($im2, 1, 5, 5,  'A Simple Text String', $text_color);
//imagepng($im2);
//imagedestroy($im2);

function p($p){
    print_r($p);
    echo "\n";
}

