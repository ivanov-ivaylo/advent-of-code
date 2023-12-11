<?php

$input = file_get_contents("input.txt");

$parts = explode("\n", $input);

function computePermutations($array) {
    $result = [];

    $recurse = function($array, $start_i = 0) use (&$result, &$recurse) {
        if ($start_i === count($array)-1) {
            array_push($result, $array);
        }

        for ($i = $start_i; $i < count($array); $i++) {
            //Swap array value at $i and $start_i
            $t = $array[$i]; $array[$i] = $array[$start_i]; $array[$start_i] = $t;

            //Recurse
            $recurse($array, $start_i + 1);

            //Restore old order
            $t = $array[$i]; $array[$i] = $array[$start_i]; $array[$start_i] = $t;
        }
    };

    $recurse($array);

    return $result;
}

$distances = array();
$city = array();

for ($i = 0; $i < count($parts); $i++) {
    $row = explode(" to ", $parts[$i]);
    $c1 = $row[0];
    $row2 = explode(" = ", $row[1]);
    $c2 = $row2[0];
    $distance = intval($row2[1]);
    
    
    if (!in_array($c1, $city)) {
        $city [] = $c1;
    }
    if (!isset($distances[$c1])) {
        $distances[$c1] = array();
    }
    if (!in_array($c2, $city)) {
        $city [] = $c2;
    }
    if (!isset($distances[$c2])) {
        $distances[$c2] = array();
    }
    $distances[$c1][$c2] = $distance;
    $distances[$c2][$c1] = $distance;
}

$allCityWays = computePermutations($city);

$shortestDistance = 100000000;
$longestDistance = 0;
for ($i = 0; $i < count($allCityWays); $i++) {
    
    $currentDistance = 0;
    for ($j = 0; $j < count($allCityWays[$i]) - 1; $j++) {
        $c1 = $allCityWays[$i][$j];
        $c2 = $allCityWays[$i][$j+1];
        if (isset($distances[$c1][$c2])) {
            $currentDistance += $distances[$c1][$c2];
        } else {
            $currentDistance +=100000000;
        }
    }
    if ($currentDistance < $shortestDistance) {
        $shortestDistance = $currentDistance;
    }
    if ($currentDistance > $longestDistance) {
        $longestDistance = $currentDistance;
    }
}

echo $shortestDistance . "\n";
echo $longestDistance;
