<?php
$input = file_get_contents("input.txt");

$rows = explode("\n", $input);

$data = array();

foreach ($rows as $row) {
    $data[] = explode(" ", $row);
}

function isStrictlyOrdered($numbers) {
    $increasing = true;
    $decreasing = true;

    for ($i = 1; $i < count($numbers); $i++) {
        $current = (int)$numbers[$i];
        $previous = (int)$numbers[$i - 1];
        $diff = abs($current - $previous);

        if ($diff < 1 || $diff > 3) {
            return false; // Difference is not within 1 to 3
        }

        if ($current <= $previous) {
            $increasing = false; // Not strictly increasing
        }

        if ($current >= $previous) {
            $decreasing = false; // Not strictly decreasing
        }
    }

    return $increasing || $decreasing;
}

function isStrictlyOrderedWithOneRemoval($numbers) {
    // Check if the sequence is already strictly ordered
    if (isStrictlyOrdered($numbers)) {
        return true;
    }

    // Check by removing one element at a time
    for ($i = 0; $i < count($numbers); $i++) {
        $modifiedSequence = $numbers;
        array_splice($modifiedSequence, $i, 1); // Remove one element
        if (isStrictlyOrdered($modifiedSequence)) {
            return true;
        }
    }

    return false; 
}

$result1 = 0;
$result2 = 0;

foreach ($data as $seq) {
    if (isStrictlyOrdered($seq)) {
        $result1++;
    }
    if (isStrictlyOrderedWithOneRemoval($seq)) {
        $result2++;
    }
    
}

echo $result1;
echo "\n";
echo $result2;