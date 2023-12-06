<?php

$input = file_get_contents("input.txt");

$commands = explode(",", $input);

$x = 0;
$y = 0;
$currOrientation = "N";
$grid = array();
for ($i = 0; $i < 1000; $i++) {
    $rowBit = array();
    for ($j = 0; $j < 1000; $j++) {
        $rowBit[] = 0;
    }
    $grid[] = $rowBit;
}

$grid[0][0] = 1;

for ($i = 0; $i < count($commands); $i++) {
    $com = trim($commands[$i]);
    $dir = $com[0];
    $num = intval(substr($com, 1));
    
    echo $currOrientation . " ";
    echo $dir . " ";
    echo $num . " ";

    switch($currOrientation) {
        case "N": 
            if ($dir == "R") {
                $x += $num;
                $currOrientation = "E";
                
            } else {
                $x -= $num;
                $currOrientation = "W";
            }
            break;
        case "E":
            if ($dir == "R") {
                $y -= $num;
                $currOrientation = "S";
            } else {
                $y += $num;
                $currOrientation = "N";
            }
            break;
        case "W":
            if ($dir == "R") {
                $y += $num;
                $currOrientation = "N";
            } else {
                $y -= $num;
                $currOrientation = "S";
            }
            break;
        case "S":
            if ($dir == "R") {
                $x -= $num;
                $currOrientation = "W";
            } else {
                $x += $num;
                $currOrientation = "E";
            }
            break;
            
    }
    
    echo $x . " ";
    echo $y . "\n";
}

echo abs($x) + abs($y);
