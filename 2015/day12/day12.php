<?php

$input = file_get_contents("input.txt");

function extractNumbersFromJson($json) {
    $decodedData = json_decode($json, true);

    if ($decodedData === null) {
        // JSON decoding failed
        return [];
    }

    $numbers = [];
    extractNumbersRecursive($decodedData, $numbers);

    return $numbers;
}

function extractNumbersRecursive($data, &$numbers) {
    foreach ($data as $value) {
        if (is_array($value)) {
            // If the value is an array, recursively call the function
            extractNumbersRecursive($value, $numbers);
        } elseif (is_numeric($value)) {
            // If the value is numeric, add it to the list of numbers
            $numbers[] = $value;
        }
    }
}

$numbers = extractNumbersFromJson($input);

$res = 0;
for ($i = 0; $i < count($numbers); $i++) {
    $res += $numbers[$i];
}

echo $res . "\n";

function extractNumbersFromJson2($json) {
    $decodedData = json_decode($json);

    if ($decodedData === null) {
        // JSON decoding failed
        return [];
    }

    $numbers = [];
    extractNumbersRecursive2($decodedData, $numbers);

    return $numbers;
}

function extractNumbersRecursive2($data, &$numbers) {
  
    
    if (is_array($data)) {
        foreach ($data as $value) {
            extractNumbersRecursive2($value, $numbers);
        }
    } elseif (is_object($data)) {
        $hasRedValue = false;
        
        foreach ($data as $propertyValue) {
            
            if ($propertyValue === "red") {
                $hasRedValue = true;
                break;
            }
        }

        if (!$hasRedValue) {
            foreach ($data as $value) {
                extractNumbersRecursive2($value, $numbers);
            }
        }
    } elseif (is_numeric($data)) {
        $numbers[] = $data;
    } 
}


$numbers = extractNumbersFromJson2($input);

$res = 0;
for ($i = 0; $i < count($numbers); $i++) {
    $res += $numbers[$i];
}

echo $res . "\n";