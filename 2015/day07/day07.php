<?php

$input = file_get_contents("input.txt");

$parts = explode("\n", $input);


$mapData = array();

for ($i = 0; $i < count($parts); $i++) {
    $elems = explode(" -> ", $parts[$i]);
    
    $mapData[$elems[1]] = $elems[0];
}

function findValue($start, $mapData, $level) {
    
    echo $start . "/" . $level . "\n";
    if ($level < 0) {
        return 0;
    }
    $level--;
    //sleep(1);
    if (empty($start)) {
        echo "#### END: " . 0 . "\n";
        return 0;
    }
    if (is_numeric($start)) {
        echo "#### END: " . $start . "\n";
        return intval($start);
    }
    if (substr_count($start, " ") == 2) {
       
        $parts = explode(" ", $start);

        if (is_numeric($parts[0])) {
            echo "#### END: " . $parts[0] . "\n";
            $v1 = intval($parts[0]);
        } else {
            $v1 = findValue($mapData[$parts[0]], $mapData, $level);
        }

        if (is_numeric($parts[2])) {
            $v2 = intval($parts[2]);
        } else {
            $v2 = findValue($mapData[$parts[2]], $mapData, $level);
        }
        
        
        if ($parts[1] == "AND") {
            return $v1 & $v2;
        }
        if ($parts[1] == "OR") {
            return $v1 | $v2;
        }
        if ($parts[1] == "LSHIFT") {
            return $v1 << $v2;
        }
        if ($parts[1] == "RSHIFT") {
            return $v1 >> $v2;
        }
        return $v1;
    } else if (substr_count($start, " ") == 1) {
        $parts = explode(" ", $start);
        if ($parts[0] == "NOT") {
            if (is_numeric($parts[1])) {
                $val = intval($parts[1]);
            } else {
                $val = findValue($mapData[$parts[1]], $mapData, $level);
            }

            echo "#### END: " . ~$val . "\n";
            return ~$val;
        }
        
    } else {
        return findValue($mapData[$start] , $mapData, $level);
    }
}

echo findValue($mapData["a"], $mapData, 100);
