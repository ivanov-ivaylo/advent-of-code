<?php

$input = file_get_contents("input.txt");


for ($i = 0; $i < 10000000; $i++) {
    $key = $input . $i;
    
    $md5 = md5($key);
    
    //echo $key . "/" . $md5 . "\n";
    if (strpos($md5, '000000') === 0) {
        echo $i;
        break;
    } 
}