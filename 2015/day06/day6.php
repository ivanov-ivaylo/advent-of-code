<?php

$input = file_get_contents("input.txt");

$input = str_replace(" through ", ":", $input);
$input = str_replace("turn on ", "^", $input);
$input = str_replace("turn off ", "v", $input);
$input = str_replace("toggle ", "t", $input);

$rows = explode("\n", $input);

function findNumberOfLights($grid) {
    $res = 0;
    for ($i = 0; $i < 1000; $i++) {
        for ($j = 0; $j < 1000; $j++) {
            if ($grid[$i][$j] != 0) {
                $res++;
            }
        }
    }
    return $res;
}

function findNumberOfBrightness($grid) {
    $res = 0;
    for ($i = 0; $i < 1000; $i++) {
        for ($j = 0; $j < 1000; $j++) {
            if ($grid[$i][$j] != 0) {
                $res = $res + $grid[$i][$j];
            }
        }
    }
    return $res;
}


function Part1($rows) {
    $grid = array();
    for ($i = 0; $i < 1000; $i++) {
        $rowBit = array();
        for ($j = 0; $j < 1000; $j++) {
            $rowBit[] = 0;
        }
        $grid[] = $rowBit;
    }

    for ($i = 0; $i < count($rows); $i++) {

        $parts = explode(":", $rows[$i]);
        if (count($parts) != 2) {
            echo "Error for row: " . $rows[$i];
            exit;
        }
        $coord1 = explode(",", substr($parts[0], 1)); 
        $coord2 = explode(",", $parts[1]);

        for ($x = $coord1[0]; $x <= $coord2[0]; $x++) {
            for ($y = $coord1[1]; $y <= $coord2[1]; $y++) {
                if (strpos($parts[0], '^') === 0) {
                    $grid[$x][$y] = 1;
                } else if (strpos($parts[0], 'v') === 0) {
                    $grid[$x][$y] = 0;
                } else if (strpos($parts[0], 't') === 0) {
                    if ($grid[$x][$y] == 0) {
                        $grid[$x][$y] = 1;
                    } else {
                        $grid[$x][$y] = 0;
                    }
                }
            }
        }

    }

    echo findNumberOfLights($grid);
}

function Part2($rows) {
    $grid = array();
    for ($i = 0; $i < 1000; $i++) {
        $rowBit = array();
        for ($j = 0; $j < 1000; $j++) {
            $rowBit[] = 0;
        }
        $grid[] = $rowBit;
    }
    
    for ($i = 0; $i < count($rows); $i++) {

        $parts = explode(":", $rows[$i]);
        if (count($parts) != 2) {
            echo "Error for row: " . $rows[$i];
            exit;
        }
        $coord1 = explode(",", substr($parts[0], 1)); 
        $coord2 = explode(",", $parts[1]);

        for ($x = $coord1[0]; $x <= $coord2[0]; $x++) {
            for ($y = $coord1[1]; $y <= $coord2[1]; $y++) {
                if (strpos($parts[0], '^') === 0) {
                    $grid[$x][$y]++;
                } else if (strpos($parts[0], 'v') === 0) {
                    $grid[$x][$y]--;
                    if ($grid[$x][$y] < 0) {
                        $grid[$x][$y] = 0;
                    }
                } else if (strpos($parts[0], 't') === 0) {
                    $grid[$x][$y] += 2;
                }
            }
        }
    }

    echo findNumberOfBrightness($grid);
}


Part1($rows);
echo "\n";
Part2($rows);