<?php

    class Solution_V2
    {
        /**
         * @param array $nums
         * @param Integer $target
         * @return array
         */

        public function twoSum (array $nums, $target)
        {
            $count = count($nums);
            $found = [];
            if ($count == 0) return [];

            for ($i = 0; $i < $count; $i ++) {
                $diff = $target - $nums[$i];

                if(array_key_exists($diff, $found)) {
                    return [$found[$diff], $i];
                }
                $found[$nums[$i]] = $i;
            }
            return [];
        }
    }

    $s      = new Solution_V2();
    $sum    = [2, 7, 11, 15];
    $target = 17;

    print_r($s->twoSum($sum, $target));
