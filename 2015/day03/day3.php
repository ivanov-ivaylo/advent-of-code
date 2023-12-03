<?php

$input = file_get_contents("input.txt");

$x = 0;
$y = 0;
$grid = array("" . $x . ":" . $y => 1);

for ($i = 0; $i < strlen($input); $i++) {
    if ($input[$i] == "^") {
        $y++;
    } else if ($input[$i] == ">") {
        $x++;
    } else if ($input[$i] == "<") {
        $x--;
    } else if ($input[$i] == "v") {
        $y--;
    }
    
    if (isset($grid["" . $x . ":" . $y])) {
        $grid["" . $x . ":" . $y] = $grid["" . $x . ":" . $y] + 1;
    } else {
        $grid["" . $x . ":" . $y] = 1;
    }
}
echo count($grid);


$x1 = 0;
$y1 = 0;
$x2 = 0;
$y2 = 0;
$grid2 = array("" . $x1 . ":" . $y1 => 1);

for ($i = 0; $i < strlen($input); $i++) {
    if ($i % 2 == 1) {
        if ($input[$i] == "^") {
            $y1++;
        } else if ($input[$i] == ">") {
            $x1++;
        } else if ($input[$i] == "<") {
            $x1--;
        } else if ($input[$i] == "v") {
            $y1--;
        }
        if (isset($grid2["" . $x1 . ":" . $y1])) {
            $grid2["" . $x1 . ":" . $y1] = $grid2["" . $x1 . ":" . $y1] + 1;
        } else {
            $grid2["" . $x1 . ":" . $y1] = 1;
        }
    }
    if ($i % 2 == 0) {
        if ($input[$i] == "^") {
            $y2++;
        } else if ($input[$i] == ">") {
            $x2++;
        } else if ($input[$i] == "<") {
            $x2--;
        } else if ($input[$i] == "v") {
            $y2--;
        }
        if (isset($grid2["" . $x2 . ":" . $y2])) {
            $grid2["" . $x2 . ":" . $y2] = $grid2["" . $x2 . ":" . $y2] + 1;
        } else {
            $grid2["" . $x2 . ":" . $y2] = 1;
        }
    }
    
}
echo "\n";
echo count($grid2);