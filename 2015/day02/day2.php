<?php
$input = file_get_contents("input.txt");



$parts = explode("\n", $input);


$result1 = 0;
$result2 = 0;
foreach ($parts as $part) {
    $dim = explode("x", $part);
    sort($dim, SORT_NUMERIC); 
    $a = intval($dim[0]);
    $b = intval($dim[1]);
    $c = intval($dim[2]);

    $area = (3 * $a * $b) + (2 * $a * $c) + (2 * $b * $c);
    $result1 += $area;    

    $result2 += 2*$a + 2*$b + ($a*$b*$c);
    
}




echo $result1;
echo "\n";
echo $result2;

?>