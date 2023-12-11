<?php

$input = file_get_contents("input.txt");

$parts = explode("\n", $input);

$result = array();

$diff1 = 0;
$diff2 = 0;
for ($i = 0; $i < count($parts); $i++) {
    $org_len = strlen($parts[$i]);
    $abs_len = strlen(stripcslashes(rtrim(ltrim($parts[$i], '"'), '"')));
    $diff1 += ($org_len - $abs_len);

    $pattern = '"\'\\';
    $diff2 += (strlen(addcslashes( $parts[$i], $pattern)) + 2 - $org_len);
}

echo $diff1 . "\n";

echo $diff2;

