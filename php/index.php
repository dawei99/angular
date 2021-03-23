<?php
class myIterator implements Iterator , ArrayAccess, Countable {
    private $position = 0;
    private $array = array(
        "firstelement",
        "secondelement",
        "lastelement",
    );

    public function __construct() {
        $this->position = 0;
    }

    /**
     * 【Iterator接口】返回到迭代器的第一个元素
     */
    function rewind() {
        $this->position = 0;
    }
    /**
     * 【Iterator接口】返回当前元素
     */
    function current() {
        return $this->array[$this->position];
    }
    /**
     * 【Iterator接口】返回当前元素的键
     */
    function key() {
        return $this->position;
    }
    /**
     * 【Iterator接口】向前移动到下一个元素
     */
    function next() {
        ++$this->position;
    }
    /**
     * 【Iterator接口】检查当前位置是否有效
     */
    function valid() {
        return isset($this->array[$this->position]);
    }


    /**
     * 【ArrayAccess接口】当使用empty()、isset()判断元素是否为空时调用
     *   例如： if (empty($object)) echo "存在";
     */
    public function offsetExists ($offset) {
        return isset($this->array[$offset]);
    }

    /**
     * 【ArrayAccess接口】当获取数组元素时执行
     *   例如： echo $object[0];
     */
    public function offsetGet($offset){
        return $this->array[$offset];
    }
    /**
     * 【ArrayAccess接口】当设置数组元素时执行
     *   例如：$object[2] = 'new_value';
     */
    public function offsetSet($offset, $value){
        $this->array[$offset] = $value;
    }
    /**
     * 【ArrayAccess接口】当删除数组元素时执行
     *   例如：unset($object[2]);
     */
    public function offsetUnset($offset){
        array_splice($this->array,$offset,1);
    }


    /**
     * 【Countable接口】当删除数组元素时执行
     *   例如：echo count($object);
     */
    public function count()
    {
        return count($this->array);
    }
}

$it = new myIterator;

// 判断存在
var_dump(isset($it[1]));
var_dump(!empty($it[1]));

// 添加元素
$it[3] = 'new_value';

// 删除元素
unset($it[1]);

// 获取元素
var_dump($it[1]);
echo "长度",(count($it));

// 遍历元素
foreach($it as $key => $value) {
    var_dump($key, $value);
    echo "\n";
}

?>
