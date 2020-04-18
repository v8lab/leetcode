
#include <stdio.h>
#include <stdlib.h>

struct ListNode {
    int val;
    struct ListNode *next;
};

struct ListNode *addTwoNumbers(struct ListNode *list1, struct ListNode *list2)
{
    int cry = 0;
    struct ListNode *head = NULL;
    struct ListNode *tail = NULL;

    while (list1 != NULL || list2 != NULL || cry > 0) {
        int val = cry;

        if (list1 != NULL) {
            val += list1->val;
            list1 = list1->next;
        }

        if (list2 != NULL) {
            val += list2->val;
            list2 = list2->next;
        }

        cry = 0;
        while (val >= 10) {
            cry += 1;
            val -= 10;
        }

        struct ListNode *node = malloc(sizeof(*node));
        if (node == NULL) {
            return NULL; // TODO: free list
        }
        node->val = val;
        node->next = NULL;

        if (head == NULL) {
            head = node;
            tail = node;
        } else {
            tail->next = node;
            tail = node;
        }
    }

    return head;
}
