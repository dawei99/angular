<?php


class test
{
    public $x;
    public function __call($name, $arguments)
    {
        var_dump($name);
        var_dump($arguments);
    }

    public function __clone()
    {
        $this->x = 1;
    }

    public static function __callStatic($name, $arguments)
    {
        var_dump($name);
        var_dump($arguments);
    }

    public function __get($name)
    {
        return $this->$name ?? "没有";
    }
    public function __set($name, $value)
    {
        $this->$name = $value;
    }
    public function __invoke()
    {
        return 123;
    }

    public function __unset($name)
    {
        echo "把$name unset";
    }
    private $b;

    public function __isset($key)
    {
        return true;
    }
}

$x = new test();
$y = clone $x;
var_dump($x->wolao(["a", "b"]));
var_dump($y);
echo "-------------------" . PHP_EOL;
var_dump($y::wokao("sadfsaf"));
echo "-------------------" . PHP_EOL;
$y->x = 1000;
echo "-------------------" . PHP_EOL;
var_dump($y->x);
echo "-------------------" . PHP_EOL;
var_dump($y());
echo "-------------------" . PHP_EOL;
unset($y->b);
echo "-------------------" . PHP_EOL;
var_dump(isset($y->sadfsafd));
echo "-------------------" . PHP_EOL;
var_dump(empty($y->sadfsafd));
echo "-------------------" . PHP_EOL;

$x = [1,2,3];
array_walk($x, function(&$v, $k){
    $v = 1;
});
var_dump($x);
echo "-------------------" . PHP_EOL;
$x =  array_map(function($value){

    return "a";

}, $x);
var_dump($x);
echo "-------------------" . PHP_EOL;
var_dump(array_search("a", $x) !== false);

$x = "'style'\ ";
echo "-------------------" . PHP_EOL;
$x = addslashes($x);
echo "-------------------" . PHP_EOL;
echo stripslashes($x);
echo "-------------------" . PHP_EOL;
var_dump(mb_convert_encoding("的", "utf-8"));

echo "-------------------" . PHP_EOL;
var_dump(strrev("abc"));
echo "-------------------" . PHP_EOL;
$name = scandir('./')[7];
$pos = strpos($name, ",");
var_dump($pos);
var_dump(substr($name, $pos));
echo "-------------------" . PHP_EOL;
$host = "www.baidu.com";
var_dump(gethostbyname($host));
echo "-------------------" . PHP_EOL;
var_dump(strstr($host, "8"));
echo "-------------------" . PHP_EOL;
var_dump(str_replace("www", "***", $host));
var_dump(preg_replace("/(w{1,3})/", "###$1", $host));
echo "-------------------" . PHP_EOL;


