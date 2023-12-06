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
    sleep(1);
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
        
        //echo "PARTS: " . $parts[0] . " " . $parts[1] . " " . $parts[2] . "\n";
        
        $v1 = "";
        if (is_numeric($parts[0])) {
            echo "#### END: " . $parts[0] . "\n";
            $v1 = intval($parts[0]);
        } else if (!empty($mapData[$parts[0]])) {
            $v1 = findValue($mapData[$parts[0]], $mapData, $level);
        }
        
        $v2 = "";
        if (is_numeric($parts[2])) {
            
            $v2 = intval($parts[2]);
            echo "#### END: " . $parts[2] . "/" . $v2 . "/" . $v1 .  "\n";
        } else if (!empty($mapData[$parts[2]])) {
            $v2 = findValue($mapData[$parts[2]], $mapData, $level);
        }
        
        
        if ($parts[1] == "AND") {
            echo "#### END: " . $v1 & $v2 . "\n";
            return $v1 & $v2;
        }
        if ($parts[1] == "OR") {
            echo "#### END: " . $v1 | $v2 . "\n";
            return $v1 | $v2;
        }
        if ($parts[1] == "LSHIFT") {
            echo "#### END: " . $v1 << $v2 . "\n";
            return $v1 << $v2;
        }
        if ($parts[1] == "RSHIFT") {
            echo "#### END: " . $v1 >> $v2 . "\n";
            return $v1 >> $v2;
        }
        echo "#### END: " . $v1 . "\n";
        return $v1;
    }if (substr_count($start, " ") == 1) {
        $parts = explode(" ", $start);
        if ($parts[0] == "NOT") {
            
            $val = "";
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
