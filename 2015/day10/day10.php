<?php

$input = file_get_contents("input.txt");

for ($i = 0; $i < 50; $i++) {
    
    $numCnt = 0;
    $numStr = "";
    $res = "";
    $len = strlen($input);
    for ($s = 0; $s < $len; $s++) {
        if ($numStr == "") {
            $numStr = $input[$s];
            $numCnt = 1;
        } else if ($numStr == $input[$s]) {
            $numCnt++;
        } else {
            $res .= $numCnt . $numStr;
            $numStr = $input[$s];
            $numCnt = 1;
        }
    }
    if ($numCnt > 0) {
        $res .= $numCnt . $numStr;
    }
    $input = $res;
    
    echo $i . " " . strlen($input) . "\n";
}


//echo $input . "\n";

echo strlen($input);