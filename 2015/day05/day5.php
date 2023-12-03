<?php

$input = file_get_contents("input.txt");

$parts = explode("\n", $input);

$result1 = 0;
$vovels = array("a", "o", "i", "u", "e");
$forbidden = array("ab", "cd", "pq", "xy");
foreach ($parts as $part) {
    $vowelsCnt = 0;
    $hasDuplicate = false;
    $hasForbidden = false;
    $prevLetter = "";
    for ($i = 0; $i < strlen($part); $i++) {
        if (in_array($part[$i], $vovels)) {
            $vowelsCnt++;
        }
        if ($prevLetter == $part[$i]) {
            $hasDuplicate = true;
        }
        if ( in_array($prevLetter . $part[$i], $forbidden) && $prevLetter != "") {
            $hasForbidden = true;
        }
        $prevLetter = $part[$i];       
    }
    if ($vowelsCnt >= 3 && $hasDuplicate && !$hasForbidden) {
        $result1++;
    }
}

echo $result1;

//Part 2

function isNiceString($s) {
    // Check condition 1: Contains a pair of any two letters that appears at least twice
    $pairRepeated = false;
    $length = strlen($s);
    
    for ($i = 0; $i < $length - 1; $i++) {
        $pair = substr($s, $i, 2);
        if (substr_count($s, $pair) >= 2 && strpos($s, $pair, $i + 2) !== false) {
            $pairRepeated = true;
            break;
        }
    }

    // Check condition 2: Contains at least one letter which repeats with exactly one letter between them
    $repeatingLetter = false;
    for ($i = 0; $i < $length - 2; $i++) {
        if ($s[$i] == $s[$i + 2]) {
            $repeatingLetter = true;
            break;
        }
    }

    // The string is nice if both conditions are satisfied
    return $pairRepeated && $repeatingLetter;
}


$result2 = 0;
foreach ($parts as $part) {
    if (isNiceString($part)) {
        $result2++;
    }
}
echo "\n";
echo $result2;