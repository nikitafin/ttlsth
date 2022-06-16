/**
 * @file 2.1.remove_dups.cc
 * @author nikitafin
 *  (finogenov.nik@gmail.com)
 * @brief https://www.lintcode.com/problem/217/
 *
 * @date 2022-06-15
 *
 */

#include <unordered_set>

class ListNode
{
public:
  int val;
  ListNode *next;
  ListNode (int val)
  {
    this->val = val;
    this->next = nullptr;
  }
};

#ifndef WITHOUT_BUFFER
class Solution
{
private:
  std::unordered_set<decltype (ListNode::val)> m_nodes;

public:
  /**
   * @param head: The first node of linked list.
   * @return: Head node.
   */
  ListNode *
  removeDuplicates (ListNode *head)
  {
    ListNode *current = head, *previous = nullptr;
    while (current != nullptr)
      {
        if (m_nodes.count (current->val))
          {
            ListNode *to_delete = current;
            current = current->next;

            if (previous != nullptr)
              {
                previous->next = current;
              }

            delete to_delete;
          }
        else
          {
            m_nodes.insert (current->val);
            previous = current;
            current = current->next;
          }
      }

    return head;
  }
};

#else

class Solution
{
private:
  std::unordered_set<decltype (ListNode::val)> m_nodes;

public:
  ListNode *
  removeDuplicates (ListNode *head)
  {
    throw "implement me";
  }
};

#endif

int
main ()
{
  ListNode *head = new ListNode (1);

  ListNode *n1 = new ListNode (2);
  head->next = n1;

  ListNode *n2 = new ListNode (3);
  n1->next = n2;

  Solution s;
  ListNode *new_head = s.removeDuplicates (head);

  return 0;
}