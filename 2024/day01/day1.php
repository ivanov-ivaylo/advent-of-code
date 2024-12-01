<?php

$input = file_get_contents("input.txt");

$rows = explode("\n", $input);

$left = array();
$right = array();
foreach ($rows as $row) {
    $parts = explode("   ", $row);
    
    $left[] = $parts[0];
    $right[] = $parts[1];
}

sort($left, SORT_NUMERIC);
sort($right, SORT_NUMERIC);


$result1 = 0;
for ($i = 0; $i < count($left); $i++) {
    $result1 += abs((int)$left[$i] - (int)$right[$i]);
}

echo $result1;

//part 2

$counts_right = array_count_values($right);

$result2 = 0;
for ($i = 0; $i < count($left); $i++) {
    
    $occurrences = isset($counts_right[$left[$i]]) ? $counts_right[$left[$i]] : 0;
    
    $result2 += $left[$i] * $occurrences;
}

echo "\n";
echo $result2;