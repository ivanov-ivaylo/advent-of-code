<?php

$input = file_get_contents("input.txt");

function hasIncreasingStraight($str) {
    // Define the pattern for one increasing straight of at least three letters
    $pattern = '/(?:abc|bcd|cde|def|efg|fgh|ghi|hij|ijk|jkl|klm|lmn|mno|nop|opq|pqr|qrs|rst|stu|tuv|uvw|vwx|wxy|xyz)/';

    // Check if the string contains the pattern
    return preg_match($pattern, strtolower($str)) !== 0;
}

function doesNotContainConfusingLetters($str) {
    // Define the array of forbidden letters
    $forbiddenLetters = ['i', 'o', 'l'];

    // Check if the string contains any forbidden letters
    foreach ($forbiddenLetters as $letter) {
        if (stripos($str, $letter) !== false) {
            return false;
        }
    }
    // If no forbidden letters are found, return true
    return true;
}

function hasTwoDifferentPairs($str) {
    $length = strlen($str);

    // Check for the first pair
    for ($i = 0; $i < $length - 1; $i++) {
        if ($str[$i] === $str[$i + 1]) {
            $pair1 = $str[$i];
            break;
        }
    }

    // If no first pair is found, return false
    if (!isset($pair1)) {
        return false;
    }

    // Check for the second pair
    for ($j = $i + 2; $j < $length - 1; $j++) {
        if ($str[$j] === $str[$j + 1] && $str[$j] !== $pair1) {
            return true;
        }
    }

    // If no second pair is found, return false
    return false;
}

function generateNextString($str) {
    $length = strlen($str);

    // Start from the rightmost character
    $i = $length - 1;

    // Increment the rightmost character
    $str[$i] = incrementChar($str[$i]);

    // Handle wrapping around for each character to the left
    while ($i > 0 && $str[$i] === 'a') {
        $i--;
        $str[$i] = incrementChar($str[$i]);
    }

    // If the leftmost character wraps around, prepend a new character
    if ($i === 0 && $str[$i] === 'a') {
        $str = 'a' . $str;
    }

    return $str;
}

function incrementChar($char) {
    // Increment a character, wrapping around from 'z' to 'a'
    if ($char === 'z') {
        return 'a';
    } else {
        return chr(ord($char) + 1);
    }
}

function isValidPassword($password) {
    
    return hasIncreasingStraight($password) && doesNotContainConfusingLetters($password) && hasTwoDifferentPairs($password);
}

$password = generateNextString($input);
while (!isValidPassword($password)) {
    $password = generateNextString($password);
}

$password2 = generateNextString($password);
while (!isValidPassword($password2)) {
    $password2 = generateNextString($password2);
}

echo $password . "\n";
echo $password2 . "\n";