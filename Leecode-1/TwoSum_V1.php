<?php
  class Solution_V1
  {
      /**
       * @param array $nums
       * @param Integer $target
       * @return array
       */

      public function twoSum(array $nums, $target)
      {
          $count  =  count($nums);
          if ($count == 0) return [];

          for ($i = 0; $i < $count; $i++) {
              $value  = $nums[$i];

              if ($a = array_keys($nums, $target - $value)) {
                  return array($a[0], $i);
              }
          }
          return [];
      }
  }

  $s        = new Solution_V1();
  $sum      = [2,7,11,115];
  $target   = 17;

  print_r($s->twoSum($sum, $target));