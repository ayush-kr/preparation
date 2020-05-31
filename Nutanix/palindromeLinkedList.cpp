/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 * 
 * Return if linked list is Palindrome
 */
#include<bits/stdc++.h>
class Solution {
        struct ListNode {
        int val;
        ListNode *next;
        ListNode() : val(0), next(nullptr) {}
        ListNode(int x) : val(x), next(nullptr) {}
        ListNode(int x, ListNode *next) : val(x), next(next) {}
  };
public:
    ListNode* reverseIt(ListNode *head){
        if(head == NULL)
            return head;
        ListNode* current = head;
        ListNode* nextNext;
        ListNode* previous = NULL;
        ListNode* next = current->next;
        while(next != NULL){
            nextNext = next->next;
            next->next = current;
            current->next = previous;
            previous = current;
            current = next;
            next = nextNext;
        }
        return current;
    }
    
    bool isPalindrome(ListNode* head) {
        ListNode* temp = head;
        int n = 0;
        while(temp != NULL){
            temp = temp->next;
            n++;
        }
        if (n<=1)
            return true;
        temp = head;
        int cou = 0;
        while(cou!=n/2){
            temp = temp->next;
            cou++;
        }
        ListNode* half;
        if (n%2 == 0)
            half = reverseIt(temp);
        else
            half = reverseIt(temp->next);
        
        while(head!= NULL && half!=NULL){
            if(head->val != half->val){
                return false;
            }
            head = head->next;
            half = half->next;
        }
        return true;
    }
};