<?php
$wires = array();
$op = array('AND' => '&', 'OR' => '|', 'NOT' => '~', 'RSHIFT' => '>>', 'LSHIFT' => '<<');

foreach (file('input.txt', FILE_IGNORE_NEW_LINES) as $line) {
	list ($k, $v) = explode(' -> ', $line);
	$wires[$v] = $k;
}

function f($w) {
	global $wires;
	
	if (!isset($wires[$w])) return $w;
	if (strpos($wires[$w], ' ') !== false) {
		eval('$wires[$w] = (' . preg_replace_callback('#(([a-z0-9]+) )?([A-Z]+) ([a-z0-9]+)#', function ($p) {
			return f($p[2]) . $GLOBALS['op'][$p[3]] . f($p[4]);
		}, $wires[$w]) . ' & 65535);');
	}
	
	return f($wires[$w]);
}

$wires['b'] = 44430;
//part 2
//$wires['b'] = 3176;
echo f('a');