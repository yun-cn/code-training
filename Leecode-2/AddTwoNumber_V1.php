<?php
/**
* @param ListNode $l1
* @param ListNode $l2
* @return ListNode
*/

class ListNode {
    function __construct($val)
    {
        $this->val      = $val;
    }
    public $val     = 0;
    public $next    = null;
}
class Solution
{
    function addTwoNumbers(ListNode $l1, ListNode $l2)
    {
        function addTwoNumbers($l1, $l2) {
            $dummy_head = new ListNode(0);
            $cur = $dummy_head;
            $carry = $sum = 0;
            while ($l1 != null || $l2 != null) {
                $val1 = $l1 != null ? $l1->val : 0;
                $val2 = $l2 != null ? $l2->val : 0;
                $sum = $val1 + $val2 + $carry;
                $cur->next = new ListNode($sum % 10);
                $cur = $cur->next;
                $carry = intval($sum / 10);
                $l1 = $l1->next;
                $l2 = $l2->next;
            }

            if ($carry > 0) {
                $cur->next = new ListNode($carry);
            }

            return $dummy_head->next;
        }
    }
}
