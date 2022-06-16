#include <gtest/gtest.h>

#include <utility>

#ifdef LEETCODE_LIST
struct ListNode
{
  int val;
  ListNode* next;
  ListNode()
    : val(0)
    , next(nullptr)
  {
  }
  ListNode(int x)
    : val(x)
    , next(nullptr)
  {
  }
  ListNode(int x, ListNode* next)
    : val(x)
    , next(next)
  {
  }
};

#endif

class Solution
{
public:
  ListNode* removeNthFromEnd(ListNode* head, int n)
  {
    auto *prev = head, *current = head, *next = head;
    for (int i = 0; i < n; ++i) {
      next = next->next;
    }

    if (next == nullptr) {
      head = head->next;
      delete prev;

      return head;
    }

    while (next != nullptr) {
      next = next->next;
      prev = current;
      current = current->next;
    }

    prev->next = current->next;
    delete current;

    return head;
  }
};

ListNode*
GenerateLinkedList(std::initializer_list<int> vals)
{
  ListNode* head = nullptr;
  for (auto it = std::rbegin(vals), end = std::rend(vals); it != end; ++it) {
    if (!head) {
      head = new ListNode(*it);
    } else {
      head = new ListNode(*it, head);
    }
  }
  return head;
}

bool
Equal(ListNode* lhs, ListNode* rhs)
{
  while (lhs != nullptr || rhs != nullptr) {
    if (lhs->val != rhs->val) {
      return false;
    }
    lhs = lhs->next;
    rhs = rhs->next;
  }

  if (lhs != nullptr || rhs != nullptr) {
    return false;
  }
  return true;
}

#if 1
TEST(Test, BasicTest)
{
  auto head = GenerateLinkedList({ 1, 2, 3, 4, 5 });
  auto expected = GenerateLinkedList({ 1, 2, 3, 5 });

  Solution s;
  head = s.removeNthFromEnd(head, 2);

  ASSERT_TRUE(Equal(head, expected));
}

TEST(Test, OneListValueTest)
{
  auto head = GenerateLinkedList({ 1 });
  auto expected = GenerateLinkedList({});

  Solution s;
  head = s.removeNthFromEnd(head, 1);

  ASSERT_TRUE(Equal(head, expected));
}

TEST(Test, MultipleValueTest)
{
  auto head = GenerateLinkedList({ 1, 2 });
  auto expected = GenerateLinkedList({ 1 });

  Solution s;
  head = s.removeNthFromEnd(head, 1);

  ASSERT_TRUE(Equal(head, expected));
}
#endif

int
main(int argc, char* argv[])
{
  ::testing::InitGoogleTest(&argc, argv);

  return RUN_ALL_TESTS();
}